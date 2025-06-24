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

package adapter

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type RestyClient struct {
	client *resty.Client
}

func NewRestyClient() *RestyClient {
	return &RestyClient{resty.NewWithClient(http.DefaultClient)}
}

func (r *RestyClient) Do(req *http.Request) (*http.Response, error) {
	rreq := r.client.R()
	rreq.Method = req.Method
	rreq.URL = req.URL.String()
	rreq.Header = req.Header
	rreq.Body = req.Body

	rresp, err := rreq.Send()
	if err != nil {
		return nil, err
	}

	resp := rresp.RawResponse
	resp.ContentLength = rresp.Size()
	resp.Body = io.NopCloser(bytes.NewReader(rresp.Body()))
	return resp, nil
}
