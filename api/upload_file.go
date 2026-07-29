// Copyright 2026 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// # 上传文件
//
// 表示待上传的本地文件或内存内容。请使用 UploadFileFromPath、UploadFileFromFile、
// UploadFileFromBytes 或 UploadFileFromReadSeeker 创建。
//
// 资源关闭说明：
//   - UploadFileFromPath：上传完成后由 SDK 自动关闭；
//   - UploadFileFromFile / UploadFileFromReadSeeker：由调用方自行关闭；
//   - UploadFileFromBytes：无需关闭。
//
// 由 SDK 自动关闭的 UploadFile 在一次上传后不可再次使用；如需重试，请重新创建。
type UploadFile struct {
	Name   string        // 文件名（multipart 中的 filename）
	Reader io.ReadSeeker // 文件内容
	Size   int64         // 文件大小（字节）
	closer io.Closer
}

// Close 关闭由本对象托管的底层资源。可重复调用。
func (uf *UploadFile) Close() error {
	if uf == nil || uf.closer == nil {
		return nil
	}
	c := uf.closer
	uf.closer = nil
	return c.Close()
}

func (uf *UploadFile) writeTo(w *multipart.Writer, field string) error {
	if uf == nil {
		return errors.New("api: UploadFile is nil")
	}
	if uf.Reader == nil {
		return fmt.Errorf("api: UploadFile %q has nil Reader", uf.Name)
	}
	if field == "" {
		return errors.New("api: multipart file field name is empty")
	}
	name := uf.Name
	if name == "" {
		name = "file"
	}

	if _, err := uf.Reader.Seek(0, io.SeekStart); err != nil {
		return err
	}

	part, err := w.CreateFormFile(field, name)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, uf.Reader)
	return err
}

// UploadFileFromPath 从本地路径创建 UploadFile。
//
// 支持以 "~" 开头的用户主目录路径。创建后由 SDK 在上传完成时自动关闭文件。
func UploadFileFromPath(path string) (*UploadFile, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, errors.New("api: file path is empty")
	}

	expanded, err := expandHomeDir(path)
	if err != nil {
		return nil, err
	}

	fd, err := os.Open(expanded)
	if err != nil {
		return nil, err
	}

	info, err := fd.Stat()
	if err != nil {
		_ = fd.Close()
		return nil, err
	}
	if info.IsDir() {
		_ = fd.Close()
		return nil, fmt.Errorf("api: %q is a directory", expanded)
	}

	return &UploadFile{
		Name:   filepath.Base(expanded),
		Reader: fd,
		Size:   info.Size(),
		closer: fd,
	}, nil
}

// UploadFileFromFile 使用已打开的 *os.File 创建 UploadFile。
//
// 文件仍由调用方负责关闭。
func UploadFileFromFile(f *os.File) (*UploadFile, error) {
	if f == nil {
		return nil, errors.New("api: os.File is nil")
	}

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &UploadFile{
		Name:   filepath.Base(f.Name()),
		Reader: f,
		Size:   info.Size(),
		closer: nil,
	}, nil
}

// UploadFileFromBytes 从内存字节创建 UploadFile，无需手动关闭。
func UploadFileFromBytes(name string, data []byte) *UploadFile {
	if name == "" {
		name = "file"
	}
	return &UploadFile{
		Name:   name,
		Reader: bytes.NewReader(data),
		Size:   int64(len(data)),
		closer: nil,
	}
}

// UploadFileFromReadSeeker 使用任意 io.ReadSeeker 创建 UploadFile。
//
// 若 size <= 0，将尝试自动探测大小。底层 Reader 仍由调用方负责关闭。
func UploadFileFromReadSeeker(name string, r io.ReadSeeker, size int64) (*UploadFile, error) {
	if r == nil {
		return nil, errors.New("api: ReadSeeker is nil")
	}
	if name == "" {
		name = "file"
	}

	if size <= 0 {
		if detected, err := detectReadSeekerSize(r); err == nil {
			size = detected
		}
	}

	return &UploadFile{
		Name:   name,
		Reader: r,
		Size:   size,
		closer: nil,
	}, nil
}

func expandHomeDir(path string) (string, error) {
	if path == "~" || strings.HasPrefix(path, "~/") || strings.HasPrefix(path, `~\`) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if path == "~" {
			return homeDir, nil
		}
		return filepath.Join(homeDir, path[2:]), nil
	}
	return path, nil
}

func detectReadSeekerSize(r io.ReadSeeker) (int64, error) {
	cur, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	end, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	if _, err = r.Seek(cur, io.SeekStart); err != nil {
		return 0, err
	}
	return end, nil
}
