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
}

// NewConfig will return a new Config object preseeded from the
// default configuration path.
func NewConfig(path string) (*Config, error) {
	config := &Config{}
	config.Database.Driver = "sqlite3"
	config.Database.Name = ":memory:"

	t, err := toml.LoadFile(DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	if err = t.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, err
}
