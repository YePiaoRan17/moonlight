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

package kind

import (
	"fmt"

	"github.com/ping-cloudnative/moonlight/pkg/common/errors"
)

type Email struct {
	SMTPHost     string `json:"smtpHost"`
	SMTPUser     string `json:"smtpUser"`
	SMTPPassword string `json:"smtpPassword"`
	SMTPPort     int64  `json:"smtpPort"`
	SMTPIsSSL    bool   `json:"smtpIsSSL"`
}

func (email *Email) Validate() error {
	if email.SMTPHost == "" {
		return errors.NewMissingParameterError("smtpHost")
	}
	if email.SMTPUser == "" {
		return errors.NewMissingParameterError("smtpUser")
	}
	if email.SMTPPassword == "" {
		return errors.NewMissingParameterError("smtpPassword")
	}
	if email.SMTPPort < 0 || email.SMTPPort > 65535 {
		return fmt.Errorf("invalidate parameter %d", email.SMTPPort)
	}
	return nil
}
