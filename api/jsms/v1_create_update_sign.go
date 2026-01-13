// Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

package jsms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"unicode/utf8"

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

	if err := validateSignParam(param); err != nil {
		return nil, err
	}

	body, err := buildSignMultipartBody(param)
	if err != nil {
		return nil, err
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
	// 【必填】签名类型，填写数字代号即可。
	//
	// 支持的类型值：
	//  - 1：公司名称全称或简称（需要营业执照图片）；
	//  - 3：APP 应用全称或简称（需要 ICP 备案 APP 截图）；
	//  - 6：品牌名称全称或简称（需要商标注册证图片）；
	//  - 7：其他类型（需要第三方授权委托书图片）。
	Type int `json:"type"`
	// 【必填】法人姓名，会进行验证。
	LegalPersonName string `json:"legalPersonName"`
	// 【必填】法人身份证号码，必须为有效的 18 位身份证号码。
	LegalPersonIDNumber string `json:"legalPersonIdNumber"`
	// 【必填】法人身份证图片，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	IDCardImage interface{} `json:"idCardImage"`
	// 【必填】统一社会信用代码，必须为有效的 18 位统一社会信用代码。
	SocialCreditCode string `json:"socialCreditCode"`
	// 【可选】第三方公司名称。
	ThirdPartyCompanyName string `json:"thirdPartyCompanyName,omitempty"`
	// 【条件必填】营业执照图片，Type 为 1 时必填，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	BusinessLicenseImage interface{} `json:"businessLicenseImage,omitempty"`
	// 【条件必填】ICP 备案 APP 截图，Type 为 3 时必填，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	ICPAppScreenshot interface{} `json:"icpAppScreenshot,omitempty"`
	// 【条件必填】商标注册证图片，Type 为 6 时必填，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	TrademarkImage interface{} `json:"trademarkImage,omitempty"`
	// 【条件必填】第三方授权委托书图片，Type 为 7 时必填，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	ThirdPartyAuthImage interface{} `json:"thirdPartyAuthImage,omitempty"`
	// 【可选】其他相关图片，文件支持 PNG、JPG、JPEG 格式，且大小不超过 2M。
	OtherImage interface{} `json:"otherImage,omitempty"`
	// 【可选】申请说明，请简略描述您的业务使用场景，不超过 100 个字。
	Remark string `json:"remark,omitempty"`
}

var (
	ErrNilParam            = errors.New("`param` cannot be nil")
	ErrInvalidSignLength   = errors.New("`sign` must be 2–8 characters")
	ErrInvalidSignType     = errors.New("invalid sign `type`, must be 1, 3, 6, or 7")
	ErrMissingLegalInfo    = errors.New("legal person info incomplete")
	ErrIDCardImageRequired = errors.New("`idCardImage` is required")
	ErrRemarkTooLong       = errors.New("`remark` too long, must be ≤100 characters")
	ErrMissingTypeFile     = errors.New("required image missing for sign `type`")
)

func validateSignParam(p *SignCreateParam) error {
	if p == nil {
		return ErrNilParam
	}

	if l := utf8.RuneCountInString(p.Sign); l < 2 || l > 8 {
		return ErrInvalidSignLength
	}

	if !isValidSignType(p.Type) {
		return ErrInvalidSignType
	}

	if p.LegalPersonName == "" ||
		p.LegalPersonIDNumber == "" ||
		p.SocialCreditCode == "" {
		return ErrMissingLegalInfo
	}

	if p.IDCardImage == nil {
		return ErrIDCardImageRequired
	}

	if utf8.RuneCountInString(p.Remark) > 100 {
		return ErrRemarkTooLong
	}

	if requiresExtraFile(p.Type) && getExtraFile(p) == nil {
		return ErrMissingTypeFile
	}

	return nil
}

func isValidSignType(t int) bool {
	switch t {
	case 1, 3, 6, 7:
		return true
	default:
		return false
	}
}

func requiresExtraFile(t int) bool {
	switch t {
	case 1, 3, 6, 7:
		return true
	default:
		return false
	}
}

func getExtraFile(p *SignCreateParam) interface{} {
	switch p.Type {
	case 1:
		return p.BusinessLicenseImage
	case 3:
		return p.ICPAppScreenshot
	case 6:
		return p.TrademarkImage
	case 7:
		return p.ThirdPartyAuthImage
	default:
		return nil
	}
}

func getExtraFileFieldName(t int) string {
	switch t {
	case 1:
		return "businessLicenseImage"
	case 3:
		return "icpAppScreenshot"
	case 6:
		return "trademarkImage"
	case 7:
		return "thirdPartyAuthImage"
	default:
		return ""
	}
}

func buildSignMultipartBody(p *SignCreateParam) (api.MultipartFormDataBody, error) {
	fields := make([]api.FormField, 0, 8)

	appendField := func(name, val string) {
		if val != "" {
			fields = append(fields, api.FormField{
				Name:  name,
				Value: val,
			})
		}
	}

	appendField("sign", p.Sign)
	appendField("type", strconv.Itoa(p.Type))
	appendField("legalPersonName", p.LegalPersonName)
	appendField("legalPersonIdNumber", p.LegalPersonIDNumber)
	appendField("socialCreditCode", p.SocialCreditCode)
	appendField("thirdPartyCompanyName", p.ThirdPartyCompanyName)
	appendField("remark", p.Remark)

	body := api.MultipartFormDataBody{
		Fields: fields,
		Files: []api.FormFile{
			{
				FieldName: "idCardImage",
				FileData:  p.IDCardImage,
			},
		},
		FileValidator: &api.FileValidator{
			MaxSize:      2 * 1024 * 1024, // 2MB
			AllowedMimes: []string{"image/jpeg", "image/png"},
			AllowedExts:  []string{".jpg", ".jpeg", ".png"},
		},
	}

	if extraFile := getExtraFile(p); extraFile != nil {
		body.Files = append(body.Files, api.FormFile{
			FieldName: getExtraFileFieldName(p.Type),
			FileData:  extraFile,
		})
	}

	return body, nil
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
