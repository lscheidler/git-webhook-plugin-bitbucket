/*
Copyright 2019 Lars Eric Scheidler

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

package push

type Change struct {
	New       State       `json:"new"` // Hash
	Old       State       `json:"old"` // Hash
	Forced    bool        `json:"forced"`
	Links     interface{} `json:"links"` // Array
	Created   bool        `json:"created"`
	Commits   interface{} `json:"commits"` // Array
	Truncated bool        `json:"truncated"`
	Closed    bool        `json:"closed"`
}

type State struct {
	Type                 string      `json:"type"`
	Name                 string      `json:"name"`
	Links                interface{} `json:"links"`
	DefaultMergeStrategy string      `json:"default_merge_strategy"`
	MergeStrategies      []string    `json:"merge_strategies"`
	Target               interface{} `json:"target"`
}
