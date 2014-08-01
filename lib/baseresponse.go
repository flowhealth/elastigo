// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elastigo

import (
	"encoding/json"
	"strconv"
)

type BaseResponse struct {
	Ok      bool        `json:"ok"`
	Index   string      `json:"_index,omitempty"`
	Type    string      `json:"_type,omitempty"`
	Id      string      `json:"_id,omitempty"`
	Source  interface{} `json:"_source,omitempty"` // depends on the schema you've defined
	Version int         `json:"_version,omitempty"`
	Found   bool        `json:"found,omitempty"`
	Exists  bool        `json:"exists,omitempty"`
	Matches []string    `json:"matches,omitempty"` // percolate matches
}

// StatusInt is required because /_optimize, at least, returns its status as
// strings instead of integers.
type StatusInt int

func (self *StatusInt) UnmarshalJSON(b []byte) error {
	s := ""
	if json.Unmarshal(b, &s) == nil {
		if i, err := strconv.Atoi(s); err == nil {
			*self = StatusInt(i)
		}
	}
	i := 0
	err := json.Unmarshal(b, &i)
	if err == nil {
		*self = StatusInt(i)
	}
	return err
}

func (self *StatusInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(*self)
}

type Status struct {
	Total      StatusInt `json:"total"`
	Successful StatusInt `json:"successful"`
	Failed     StatusInt `json:"failed"`
}

type ExtendedStatus struct {
	Ok           bool   `json:"ok"`
	ShardsStatus Status `json:"_shards"`
}

type Match struct {
	OK          bool         `json:"ok"`
	Matches     []string     `json:"matches"`
	Explanation *Explanation `json:"explanation,omitempty"`
}

type Explanation struct {
	Value       float32        `json:"value"`
	Description string         `json:"description"`
	Details     []*Explanation `json:"details,omitempty"`
}

func ScrollDuration(duration string) string {
	scrollString := ""
	if duration != "" {
		scrollString = "&scroll=" + duration
	}
	return scrollString
}

// http://www.elasticsearch.org/guide/reference/api/search/search-type/
