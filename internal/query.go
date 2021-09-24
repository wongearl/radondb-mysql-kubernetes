/*
Copyright 2021 RadonDB.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"strings"
	"errors"
)

// Query contains a escaped query string with variables marked with a question mark (?) and a slice
// of positional arguments.
type Query struct {
	escapedQuery string
	args         []interface{}
}

// String representation of the query.
func (q *Query) String() string {
	return q.escapedQuery
}

// Args is used in test.
func (q *Query) Args() []interface{} {
	return q.args
}

// NewQuery returns a new Query object.
func NewQuery(q string, args ...interface{}) Query {
	if q == "" {
		internalLog.Error(errors.New("SQLError"), "sql cannot be empty")
	}

	if !strings.HasSuffix(q, ";") {
		q += ";"
	}

	return Query{
		escapedQuery: q,
		args:         args,
	}
}
