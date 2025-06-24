/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

package jsms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 创建签名
//   - 功能说明：创建短信签名。
//   - 调用地址：POST `/v1/sign`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E5%88%9B%E5%BB%BA%E7%AD%BE%E5%90%8D-api
func (s *apiv1) CreateSign(ctx context.Context, param *SignCreateParam) (*SignCreateResult, error) {
	return s.sign(ctx, 0, param, "create")
}

// # 修改签名
//   - 功能说明：修改审核不通过的签名，并再次提交审核。
//   - 调用地址：POST `/v1/sign/{signID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E4%BF%AE%E6%94%B9%E7%AD%BE%E5%90%8Dapi
func (s *apiv1) UpdateSign(ctx context.Context, signID int, param *SignUpdateParam) (*SignUpdateResult, error) {
	return s.sign(ctx, signID, param, "update")
}

// 创建或修改签名，createOrUpdate 可选值：create、update。
func (s *apiv1) sign(ctx context.Context, signID int, param *SignCreateParam, createOrUpdate string) (*SignCreateResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	var fields []api.FormField

	l := len(param.Sign)
	if l < 2 || l > 8 {
		return nil, errors.New("`sign` must be 2-8 characters")
	}
	fields = append(fields, api.FormField{Name: "sign", Value: param.Sign})

	t := param.Type
	if t < 1 || t > 6 {
		return nil, errors.New("invalid `type`, must be 1-6")
	}
	fields = append(fields, api.FormField{Name: "type", Value: strconv.Itoa(t)})

	l = len(param.Remark)
	if l > 100 {
		return nil, errors.New("`remark` too long, must be ≤100 characters")
	} else if l > 0 {
		fields = append(fields, api.FormField{Name: "remark", Value: param.Remark})
	}

	body := api.MultipartFormDataBody{Fields: fields}

	images := reflect.ValueOf(param.Images)
	if images.Kind() != reflect.Invalid {
		body.FileValidator = &api.FileValidator{
			MaxSize:      2 * 1024 * 1024, // 2MB
			AllowedMimes: []string{"image/jpeg", "image/png"},
			AllowedExts:  []string{".jpg", ".jpeg", ".png"},
		}
		if images.Kind() == reflect.Slice || images.Kind() == reflect.Array {
			for i := 0; i < images.Len(); i++ {
				image := images.Index(i).Interface()
				body.Files = append(body.Files, api.FormFile{FieldName: "images", FileData: image})
			}
		} else {
			body.Files = []api.FormFile{{FieldName: "images", FileData: images.Interface()}}
		}
	}

	url := s.host + "/v1/sign"
	if createOrUpdate == "update" {
		url += "/" + strconv.Itoa(signID)
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    url,
		Auth:   s.auth,
		Body:   body,
	}
	resp, err := s.client.FormRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SignCreateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SignCreateParam struct {
	// 【必填】签名内容。
	//  - 仅支持中英文、数字组成，长度为 2~8 个字；
	//  - 无须添加【】、[]、（）等符号；
	//  - 范例：极光推送
	Sign string `json:"sign"`
	// 【可选】签名类型，填写数字代号即可。
	//  - 1、公司名称全称或简称：需提供签名所属主体的营业执照复印件图片及对应法人代表的身份证正反面复印件图片，均需加盖公章；
	//  - 2、工信部备案的网站全称或简称：需提供签名所属的已备案的 ICP 备案截图复印件图片、主办单位的营业执照复印件图片及对应法人代表的身份证正反面复印件图片，均需加盖公章；
	//  - 3、APP 应用名称或简称：需提供签名所属的任意应用商店的下载链接、APP 软著证明复印件图片及开发者的营业执照复印件图片、对应法人代表的身份证正反面复印件图片，均需加盖公章；
	//  - 4、公众号小程序全称或简称：需提供签名所属的公众号小程序含主体的页面截图、开发者主体营业执照复印件图片、对应法人代表的身份证正反面复印件图片，均需加盖公章；
	//  - 5、商标名称全称或简称：需提供签名所属商标注册证书复印件图片及商标主体营业执照复印件图片、对应法人代表身份证正反面复印件图片，均需加盖公章；
	//  - 6、其他：申请的签名与所属主体不一致或涉及第三方权益时，需提供第三方授权委托书、第三方签名相关资质（详见类型 1-5），[授权委托书格式参考文档]。
	// 注意：
	//  - 如果用户 A 创建自用签名，签名归属主体属于 A，则不涉及第三方权益；
	//  - 如果用户 A 替公司 B 创建签名，签名归属主体属于公司 B，则涉及第三方权益，需要公司 B 给予 A 授权委托书。公司 B 为授权方，A 为被授权方，并需要加盖公司 B 的公章。
	// [授权委托书格式参考文档]: https://shimo.im/docs/vqqd6wgwhXCtHjC3
	Type int `json:"type,omitempty"`
	// 【可选】上传签名相关的资质证件图片（文件支持 PNG、JPG、JPEG 格式，且每个大小不超过 2M）。
	Images interface{} `json:"images,omitempty"`
	// 【可选】简略描述您的业务使用场景，不超过 100 个字。
	Remark string `json:"remark,omitempty"`
}

type SignCreateResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	SignID        int            `json:"sign_id,omitempty"` // 签名 ID
}

func (rs *SignCreateResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

type (
	SignUpdateParam  = SignCreateParam
	SignUpdateResult = SignCreateResult
)
