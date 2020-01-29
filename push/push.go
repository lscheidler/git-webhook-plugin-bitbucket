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

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lscheidler/git-webhook-plugin-bitbucket/common"
)

type Push struct {
	actor      *common.Owner
	repository *common.Repository
	Changes    []Change
	data       []byte
}

func (p *Push) UnmarshalJSON(b []byte) error {
	requiredKeys := []string{"actor", "repository", "push"}

	p.data = b

	var data map[string]json.RawMessage
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for _, key := range requiredKeys {
		if _, ok := data[key]; ok {
			switch key {
			case "actor":
				var actor common.Owner
				if err := json.Unmarshal(data[key], &actor); err != nil {
					return err
				}
				p.actor = &actor
			case "repository":
				var repository common.Repository
				if err := json.Unmarshal(data[key], &repository); err != nil {
					return err
				}
				p.repository = &repository
			case "push":
				var push map[string]json.RawMessage
				if err := json.Unmarshal(data[key], &push); err != nil {
					return err
				}

				var changes []Change
				if err := json.Unmarshal(push["changes"], &changes); err != nil {
					return err
				}
				p.Changes = changes
			}
		} else {
			if os.Getenv("debug") == "true" {
				log.Println("bitbucket[push]: required key missing")
			}
			return common.MissingKeyError{Key: key}
		}
	}
	return nil
}

func Init(data []byte) (*Push, error) {
	var result Push
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}

func (p *Push) Data() []byte {
	return p.data
}

func (p *Push) Name() string {
	return "push"
}

func (p *Push) RepositoryFullName() string {
	return *p.repository.FullName
}

func (p *Push) BranchNames() []string {
	branches := []string{}
	for _, v := range p.Changes {
		branches = append(branches, v.New.Name)
	}
	return branches
}
