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
	"fmt"
	"log"
	"os"
)

type MissingKeyError struct {
	Key string
}

func (m MissingKeyError) Error() string {
	return fmt.Sprintf("Key %s is missing.")
}

func ProjectValid(data map[string]interface{}) bool {
	return MustExist(data, []string{"type", "name", "uuid", "links", "key"})
}

func MustExist(data map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, ok := data[key]; !ok {
			if os.Getenv("debug") == "true" {
				log.Println("bitbucket: key", key, "is missing")
			}
			return false
		}
	}
	return true
}
