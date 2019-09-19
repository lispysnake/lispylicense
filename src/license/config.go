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
	"errors"
	"fmt"
	toml "github.com/pelletier/go-toml"
)

const (
	// DefaultConfigPath is the global default location to find a config for our service
	DefaultConfigPath = "config.toml"
)

// Config contains the service-wide configuration
type Config struct {

	// Database configuration
	Database struct {
		Driver string
		Name   string
	}

	Email struct {
		Method   string
		From     string `toml:"from"`
		SMTPUser string `toml:"smtp_user"`
		SMTPPort int    `toml:"smtp_port"`
		SMTPPass string `toml:"smtp_pass"`
		SMTPHost string `toml:"smtp_host"`
	}
}

// Ensure we have valid email setup
func (c *Config) validateEmail() error {
	if c.Email.Method == "" || c.Email.Method == "none" {
		return nil
	}
	if c.Email.Method != "smtp" {
		return fmt.Errorf("unsupported method '%v'", c.Email.Method)
	}
	if c.Email.SMTPUser == "" {
		return errors.New("missing smtp_user")
	}
	if c.Email.SMTPPass == "" {
		return errors.New("missing smtp_pass")
	}
	if c.Email.SMTPHost == "" {
		return errors.New("missing smtp_host")
	}
	return nil
}

// ShouldEmail will return true if we should email the new licenses
func (c *Config) ShouldEmail() bool {
	if c.Email.Method == "" || c.Email.Method == "none" {
		return false
	}
	return true
}

// NewConfig will return a new Config object preseeded from the
// default configuration path.
func NewConfig(path string) (*Config, error) {
	config := &Config{}
	config.Database.Driver = "sqlite3"
	config.Database.Name = ":memory:"

	config.Email.From = "no-reply@localhost"
	config.Email.Method = ""
	config.Email.SMTPPort = 587

	t, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}
	if err = t.Unmarshal(config); err != nil {
		return nil, err
	}

	if config.Email.Method != "" {
		if err := config.validateEmail(); err != nil {
			return nil, err
		}
	}

	return config, err
}
