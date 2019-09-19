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
	"github.com/jmoiron/sqlx"
	// Build go-sqlite3 support into the application too
	_ "github.com/mattn/go-sqlite3"
)

// Database wraps the underlying database connection and helps with
// the storage side of things.
type Database struct {
	conn *sqlx.DB
}

// AssignRequest is used to serialise the basic license request.
type AssignRequest struct {
	AccountID string // The ID for the license request
	LicenseID string // The ID for the license being requested
}

// CreateRequest is used for creating new licenses
type CreateRequest struct {
	LicenseID   string // The ID of the license
	Description string // Description of the license
	MaxUsers    int    // Maximum number of users for the license
}

// SchemaSqlite3 is the default schema for our Sqlite3 impl
const SchemaSqlite3 = `
CREATE TABLE IF NOT EXISTS LICENSE_USERS (
    UUID text UNIQUE NOT NULL,
    ACCOUNT_ID text NOT NULL,
    LICENSE_ID text NOT NULL 
);

CREATE TABLE IF NOT EXISTS LICENSE_SPECS (
    ID text UNIQUE NOT NULL,
    DESC text NULL,
    MAX_USERS INT
);
`

// NewDatabase will attempt to open and initialise the database using
// the given Config instance.
func NewDatabase(config *Config) (*Database, error) {
	conn, err := sqlx.Connect(config.Database.Driver, config.Database.Name)
	if err != nil {
		return nil, err
	}

	// Ensure we have a schema
	_, err = conn.Exec(SchemaSqlite3)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Database{
		conn: conn,
	}, nil
}

// Close will shutdown any resources associate with the database
func (d *Database) Close() {
	if d.conn != nil {
		d.conn.Close()
	}
}

// Assign will attempt to claim the given license request,
// and return the UUID for the allocation.
func (d *Database) Assign(req AssignRequest) (string, error) {
	return "", errors.New("Not yet implemented")
}

// InsertLicense will insert a new license if possible
func (d *Database) InsertLicense(req CreateRequest) error {
	return errors.New("Not yet implemented")
}
