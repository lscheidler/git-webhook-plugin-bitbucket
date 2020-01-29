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

package bitbucket

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"

	"github.com/lscheidler/git-webhook-plugin"
	"github.com/lscheidler/git-webhook-plugin-bitbucket/push"
)

const (
	BASEURL = "git@bitbucket.org"
)

type Bitbucket struct {
	attributes map[string]string
	event      gitWebhookPlugin.Event
}

func Init() *Bitbucket {
	return &Bitbucket{}
}

func (b *Bitbucket) Attributes() map[string]string {
	return b.attributes
}

func (b *Bitbucket) Event() gitWebhookPlugin.Event {
	return b.event
}

func (b *Bitbucket) IsBitbucketWebhookRequest(headers map[string]string) bool {
	attributes := make(map[string]string)
	headerKeys := []string{"x-attempt-number", "x-event-key", "x-hook-uuid", "x-request-uuid"}
	for _, header := range headerKeys {
		if value, ok := headers[header]; !ok {
			if os.Getenv("debug") == "true" {
				log.Println("bitbucket: header", header, "is missing")
			}
			return false
		} else {
			attributes[header] = value
		}
	}
	b.attributes = attributes
	return true
}

func (b *Bitbucket) EventType() string {
	return b.event.Name()
}

func (b *Bitbucket) GitBranches() []string {
	return b.event.BranchNames()
}

func (b *Bitbucket) GitUrl() string {
	return BASEURL + ":" + b.event.RepositoryFullName() + ".git"
}

func (b *Bitbucket) Valid(request events.ALBTargetGroupRequest) bool {
	var body []byte
	var err error
	if request.IsBase64Encoded {
		body, err = base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			log.Println("Decoding body with base64 failed.", err)
			return false
		}
	} else {
		body = []byte(request.Body)
	}
	return b.ValidBody(request.Headers["x-event-key"], body)
}

func (b *Bitbucket) ValidBody(eventType string, body []byte) bool {
	if json.Valid(body) {
		if eventType == "repo:push" {
			event, err := push.Init(body)
			if err != nil {
				if os.Getenv("debug") == "true" {
					log.Println("bitbucket: not a valid push request")
				}
				return false
			}
			b.event = event
		} else {
			log.Printf("bitbucket event %s not known\n", eventType)
			return false
		}
		//b.data = result
		return true
	}
	log.Println("JSON not valid")
	log.Println(body)
	return false
}
