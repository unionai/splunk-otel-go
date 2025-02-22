// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splunksqlx_test

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/unionai/splunk-otel-go/instrumentation/github.com/jmoiron/sqlx/splunksqlx"
)

func ExampleOpen() {
	// This assumes the instrumented driver,
	// "github.com/unionai/splunk-otel-go/instrumentation/github.com/jackc/pgx/splunkpgx",
	// is imported. That will ensure the driver and the instrumentation setup
	// for the driver are registered with the appropriate packages.
	db, err := splunksqlx.Open("pgx", "postgres://localhost/db")
	if err != nil {
		log.Print(err)
		return
	}
	defer db.Close()

	// All calls through the sqlx API are now traced.
	query, args, err := sqlx.In("SELECT * FROM users WHERE level IN (?);", []int{4, 6, 7})
	if err != nil {
		log.Print(err)
		return
	}
	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...) //nolint:sqlclosecheck // False positive: https://github.com/ryanrolds/sqlclosecheck/issues/35.
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()
	/* ... */
}
