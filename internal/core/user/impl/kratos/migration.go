// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kratos

import (
	"github.com/sirupsen/logrus"

	"github.com/ping-cloudnative/moonlight/internal/core/legacy/conf"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpclient"
)

func UserMigration(req OryKratosCreateIdentitiyRequest) (string, error) {
	return CreateUser(req)
}

func MigrationReady() bool {
	var rsp OryKratosReadyResponse
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(conf.OryKratosPrivateAddr()).
		Path("/health/ready").
		Do().JSON(&rsp)
	if err != nil {
		logrus.Errorf("get kratos user info error: %v", err)
		return false
	}
	if !r.IsOK() {
		logrus.Errorf("get kratos user info error, statusCode: %d, err: %s", r.StatusCode(), r.Body())
		return false
	}
	return rsp.Status == "ok"
}
