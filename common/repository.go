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

package common

import (
	"encoding/json"
	"log"
	"os"
)

type Repository struct {
	FullName  *string
	IsPrivate bool
	//Links     *Links
	Name  *string
	Owner *Owner
	//Project *Project
	Scm     *string
	Type    *string
	Uuid    *string
	Website *string
}

func InitRepository() *Repository {
	result := Repository{Owner: InitOwner()}
	return &result
}

func (r *Repository) UnmarshalJSON(b []byte) error {
	requiredKeys := []string{"type", "name", "full_name", "uuid", "links", "project", "website", "owner", "scm", "is_private"}
	values := []**string{&r.Type, &r.Name, &r.FullName, &r.Uuid, nil, nil, &r.Website, nil, &r.Scm, nil}

	var data map[string]json.RawMessage
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for index, key := range requiredKeys {
		if _, ok := data[key]; ok {
			switch key {
			case "is_private":
				var isPrivate bool
				if err := json.Unmarshal(data[key], &isPrivate); err != nil {
					return err
				}
				r.IsPrivate = isPrivate
			case "links":
				// TODO
			case "owner":
				var owner Owner
				if err := json.Unmarshal(data[key], &owner); err != nil {
					return err
				}
				r.Owner = &owner
			case "project":
				// TODO
			default:
				var value string
				if err := json.Unmarshal(data[key], &value); err != nil {
					return err
				}
				*values[index] = &value
			}
		} else {
			if os.Getenv("debug") == "true" {
				log.Println("bitbucket[push]: required key missing")
			}
			return MissingKeyError{Key: key}
		}
	}
	return nil
}
