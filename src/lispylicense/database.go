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

package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Database wraps the underlying database connection and helps with
// the storage side of things.
type Database struct {
	conn *sqlx.DB
}

// NewDatabase will attempt to open and initialise the database using
// the given Config instance.
func NewDatabase(config *Config) (*Database, error) {
	conn, err := sqlx.Connect(config.Database.Driver, config.Database.Name)
	if err != nil {
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
