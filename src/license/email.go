//
// Copyright © 2019 Lispy Snake, Ltd.
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
//

package license

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

type emailTemplate struct {
	LicenseDescription string
	LicenseKey         string
	LicenseID          string
	AccountID          string
}

// Email is used to render and send HTML emails in batch
type Email struct {
	req      *AssignRequest
	info     *Info // Associated license info.
	rendered *bytes.Buffer
}

// NewEmail will create a new email for the given assignment
func NewEmail(req *AssignRequest, info *Info) (*Email, error) {
	templ := template.New("email_template.html")
	t, err := templ.ParseFiles("email_template.html")

	if err != nil {
		return nil, err
	}

	tr := &emailTemplate{
		LicenseDescription: info.Description,
		LicenseID:          req.LicenseID,
		AccountID:          req.AccountID,
		LicenseKey:         req.UUID,
	}

	b := &bytes.Buffer{}
	if err = t.Execute(b, tr); err != nil {
		return nil, err
	}

	return &Email{
		info:     info,
		rendered: b,
	}, nil
}

// Send will use the configuration to send an email.
func (e *Email) Send(config *Config, from string, to []string) error {
	txt := fmt.Sprintf("From: %v\r\n", from)
	txt += fmt.Sprintf("To: %v\r\n", strings.Join(to, ", "))
	txt += "Subject: Lispy Snake Ltd License\r\n"
	txt += "MIME-version: 1.0\r\n"
	txt += "Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	txt += e.rendered.String()
	auth := smtp.PlainAuth("", from, config.Email.SMTPPass, config.Email.SMTPHost)
	addr := fmt.Sprintf("%s:%v", config.Email.SMTPHost, config.Email.SMTPPort)
	err := smtp.SendMail(addr, auth, from, to, []byte(txt))
	return err
}
