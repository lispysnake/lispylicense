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
	"fmt"
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
	UUID      string `db:"uuid"`       // UUID (set internally)
	AccountID string `db:"account_id"` // The ID for the license request
	LicenseID string `db:"license_id"` // The ID for the license being requested
}

// CreateRequest is used for creating new licenses
type CreateRequest struct {
	LicenseID   string `db:"license_id"`  // The ID of the license
	Description string `db:"description"` // Description of the license
	MaxUsers    int    `db:"max_users"`   // Maximum number of users for the license
}

// Info type is used to validate a License entry and find out
// various facts about it.
type Info struct {
	Description    string  `json:"description"`     // Description of the license
	CurrentUsers   int     `json:"current_users"`   // Curent number of users for this license
	MaxUsers       int     `json:"max_users"`       // Maximum number of users for this license
	RemainingUsers int     `json:"remaining_users"` // Remaining number of users for this license
	FillRatio      float64 `json:"fill_ratio"`      // How 'full' our sales are
}

type licenseStore struct {
	ID       string `db:"ID"`
	Desc     string `db:"DESC"`
	MaxUsers int    `db:"MAX_USERS"`
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

// GetInfo will return an Info type populated with information on the
// given license. If the license doesn't exist, then no info will be
// returned, and the user cannot be assigned the license.
func (d *Database) GetInfo(id string) (*Info, error) {
	var count int
	var req licenseStore

	// Grab the current number of licenses
	q := "SELECT COUNT(*) FROM LICENSE_USERS WHERE LICENSE_ID = ?;"
	err := d.conn.Get(&count, q, id)
	if err != nil {
		return nil, err
	}

	q = "SELECT * FROM LICENSE_SPECS WHERE ID = ?;"
	err = d.conn.Get(&req, q, id)
	if err != nil {
		return nil, err
	}

	info := &Info{
		Description:    req.Desc,
		CurrentUsers:   count,
		MaxUsers:       req.MaxUsers,
		RemainingUsers: 0,
		FillRatio:      0.0,
	}

	// Remaining user count
	if req.MaxUsers > 0 {
		info.RemainingUsers = req.MaxUsers - count
		info.FillRatio = float64(info.CurrentUsers) / float64(info.MaxUsers)
	}

	return info, nil
}

// Assign will attempt to claim the given license request,
// and return the UUID for the allocation.
func (d *Database) Assign(req AssignRequest) (string, error) {
	info, err := d.GetInfo(req.LicenseID)
	if err != nil {
		return "", err
	}

	if info.MaxUsers > 0 && info.RemainingUsers < 1 {
		return "", fmt.Errorf("Cannot assign user to license '%v' as maxUser count of '%v' would be exceeded", req.LicenseID, info.MaxUsers)
	}

	uuid, err := NewUUID()
	if err != nil {
		return "", err
	}
	req.UUID = uuid
	q := "INSERT  INTO LICENSE_USERS (UUID, ACCOUNT_ID, LICENSE_ID) VALUES (:uuid, :account_id, :license_id)"
	_, err = d.conn.NamedExec(q, &req)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

// InsertLicense will insert a new license if possible
func (d *Database) InsertLicense(req CreateRequest) error {
	q := "INSERT INTO LICENSE_SPECS (ID, DESC, MAX_USERS) VALUES (:license_id, :description, :max_users)"
	_, err := d.conn.NamedExec(q, &req)
	return err
}
