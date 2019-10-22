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

type Owner struct {
	Type *string
	//Nickname    *string
	DisplayName *string
	Uuid        *string
	//Links       *Links
}

func InitOwner() *Owner {
	result := Owner{}
	return &result
}

func (o *Owner) UnmarshalJSON(b []byte) error {
	// TODO nickname is not available in repository owner
	requiredKeys := []string{"type", "display_name", "uuid", "links"}
	values := []**string{&o.Type, &o.DisplayName, &o.Uuid, nil}

	var data map[string]json.RawMessage
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for index, key := range requiredKeys {
		if _, ok := data[key]; ok {
			switch key {
			case "links":
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
