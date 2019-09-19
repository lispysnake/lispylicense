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
)

// A Manager instace is responsible for the DB connection, and assigning
// license users as well as the input of new licenses.
type Manager struct {
	config   *Config
	database *Database
}

// NewManager will create a new manager instance and start the ball rolling
func NewManager() (*Manager, error) {
	var err error
	manager := &Manager{}

	// Get our configuration going
	manager.config, err = NewConfig(DefaultConfigPath)
	if err != nil {
		return nil, err
	}

	// Ensure we have a database now.
	manager.database, err = NewDatabase(manager.config)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

// Close will close any underlying connections and resources
func (m *Manager) Close() {
	if m.database != nil {
		m.database.Close()
	}
}

// CreateLicense will attempt to create the new named license within the
// database for subscription by users.
func (m *Manager) CreateLicense(id string, maxUsers int, desc string) error {
	return errors.New("Not yet implemented")
}

// AssignLicense will attempt to assign the license_id to account_id, returning
// the UUID for the new subscription if successful.
func (m *Manager) AssignLicense(accountID, licenseID string) (string, error) {
	return "", errors.New("Not yet implemented")
}
