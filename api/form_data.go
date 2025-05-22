/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// # 多部分表单数据的正文负载
type MultipartFormDataBody struct {
	Fields        []FormField    // 普通表单字段列表
	Files         []FormFile     // 文件表单字段列表
	FileValidator *FileValidator // 可选的文件校验器
}

// 填充多部分表单数据的正文，包括普通字段和文件。
func (mfd MultipartFormDataBody) Prepare(writer *multipart.Writer) error {
	for _, ff := range mfd.Fields {
		if err := writer.WriteField(ff.Name, ff.Value); err != nil {
			return err
		}
	}

	for _, ff := range mfd.Files {
		if err := ff.process(mfd.FileValidator, writer); err != nil {
			return err
		}
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// # 普通表单字段
type FormField struct {
	Name  string // 字段名称
	Value string // 字段值
}

// # 文件表单字段
type FormFile struct {
	FieldName string      // 字段名称
	FileName  string      // 文件名（如果 FileData 能获取到，该字段可选）
	FileData  interface{} // 文件路径或文件数据流（如果是文件数据流，请在上传完毕后自行关闭）
}

// 处理文件表单字段，将其内容写入 writer。
func (ff FormFile) process(fv *FileValidator, writer *multipart.Writer) error {
	var (
		fileReader io.Reader
		fieldName  = ff.FieldName
		fileName   = ff.FileName
	)

	switch file := ff.FileData.(type) {
	case string:
		if file == "" {
			return fmt.Errorf("%q path is not set", fieldName)
		}
		// 处理文件路径，支持 "~" 表示用户目录
		if strings.HasPrefix(file, "~") {
			if homeDir, err := os.UserHomeDir(); err != nil {
				return err
			} else {
				file = strings.Replace(file, "~", homeDir, 1)
			}
		}
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		fileReader = f
		if fileName == "" {
			fileName = filepath.Base(file)
		}
		if fieldName == "" {
			fieldName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
		}
	case *os.File:
		if file == nil {
			return fmt.Errorf("nil %q File", fieldName)
		}
		fileReader = file
		if fileName == "" {
			fileName = filepath.Base(file.Name())
		}
		if fieldName == "" {
			fieldName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
		}
	case io.Reader:
		if file == nil {
			return fmt.Errorf("nil %q Reader", fieldName)
		}
		fileReader = file
		if fileName == "" {
			fileName = "file" // unexpected, maybe use a random name?
		}
		if fieldName == "" {
			fieldName = "file" // unexpected, maybe use a random name?
		}
	default:
		if file == nil {
			return fmt.Errorf("nil %q", fieldName)
		}
		return fmt.Errorf("unsupported file type: %T", file)
	}

	if err := fv.Validate(fileReader, fileName); err != nil {
		return err
	}

	fileWriter, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fileWriter, fileReader); err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// # 文件校验器
//
// 用于校验文件的大小、MIME 类型和扩展名等。
type FileValidator struct {
	MaxSize      int64    // 最大文件大小（字节）
	AllowedMimes []string // 允许的文件 MIME 类型（如 "image/png"、"text/plain" 等）
	AllowedExts  []string // 允许的文件扩展名（如 ".png"、".txt" 等）
}

// 校验文件合法性。
func (fv *FileValidator) Validate(fileReader io.Reader, fileName string) error {
	if fv == nil {
		return nil
	}

	// 校验文件大小
	if fv.MaxSize > 0 {
		if fileStat, ok := fileReader.(interface{ Size() int64 }); ok {
			if fileStat.Size() == 0 {
				return fmt.Errorf("%q is empty", fileName)
			}
			// 直接获取文件大小
			if fileStat.Size() > fv.MaxSize {
				return fmt.Errorf("%q size %d exceeds limit of %d bytes", fileName, fileStat.Size(), fv.MaxSize)
			}
		} else {
			// 手动计算文件大小
			fileSize, err := io.Copy(io.Discard, fileReader)
			if err != nil {
				return err
			}
			if fileSize == 0 {
				return fmt.Errorf("%q is empty", fileName)
			}
			if fileSize > fv.MaxSize {
				return fmt.Errorf("%q size %d exceeds limit of %d bytes", fileName, fileSize, fv.MaxSize)
			}

			// 重置 fileReader 的读取位置
			if seeker, ok := fileReader.(io.Seeker); ok {
				_, err = seeker.Seek(0, io.SeekStart)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("%q does not support seeking", fileName)
			}
		}
	}

	// 校验文件 MIME 类型
	if len(fv.AllowedMimes) > 0 {
		buf := make([]byte, 512) // 读取文件头的前 512 字节用于检测类型
		if _, err := fileReader.Read(buf); err != nil {
			return err
		}

		// 检测文件 MIME 类型
		mimeType := http.DetectContentType(buf)
		baseMimeType := getBaseMimeType(mimeType)
		isAllowed := false
		for _, allowedMime := range fv.AllowedMimes {
			if mimeType == allowedMime || baseMimeType == allowedMime {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			return fmt.Errorf("%q MIME %q is not allowed", fileName, mimeType)
		}

		// 重置 fileReader 的读取位置
		if seeker, ok := fileReader.(io.Seeker); ok {
			_, err := seeker.Seek(0, io.SeekStart)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("%q does not support seeking", fileName)
		}
	}

	// 校验文件扩展名
	if len(fv.AllowedExts) > 0 {
		ext := strings.ToLower(filepath.Ext(fileName))
		isAllowed := false
		for _, allowedExt := range fv.AllowedExts {
			if ext == allowedExt {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			return fmt.Errorf("%q ext %q is not allowed", fileName, ext)
		}
	}

	return nil
}

// MIME 类型正则表达式
var mimeRegexp = regexp.MustCompile(`^(.*?)(;.*)?$`)

// 提取 MIME 类型的基础部分
func getBaseMimeType(mimeType string) string {
	// 用正则去掉参数部分，例如 `;charset=utf-8`。
	matches := mimeRegexp.FindStringSubmatch(mimeType)
	if len(matches) > 1 {
		return matches[1] // 返回 MIME 类型的基础部分
	}
	return mimeType
}
