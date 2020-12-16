// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.

package core

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
	"os"
	"strings"
)

// PubSubTarget holds a new client for writing events to Google PubSub
type PubSubTarget struct {
	ProjectID string
	Client    *pubsub.Client
	TopicName string
}

// NewPubSubTarget creates a new client for writing events to Google PubSub
func NewPubSubTarget(projectID string, topicName string, serviceAccountB64 string) (*PubSubTarget, error) {
	if serviceAccountB64 != "" {
		sDec, err := base64.StdEncoding.DecodeString(serviceAccountB64)
		if err != nil {
			return nil, fmt.Errorf("Could not Base64 decode service account: %s", err.Error())
		}

		targetFile := fmt.Sprintf("/tmp/stream-replicator-service-account-%s.json", uuid.NewV4().String())

		f, err := os.Create(targetFile)
		if err != nil {
			return nil, fmt.Errorf("Could not create target file '%s' for service account: %s", targetFile, err.Error())
		}
		defer f.Close()

		_, err2 := f.WriteString(string(sDec))
		if err2 != nil {
			return nil, fmt.Errorf("Could not write decoded service account to target file '%s': %s", targetFile, err.Error())
		}

		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", targetFile)
	}

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %s", err.Error())
	}

	return &PubSubTarget{
		ProjectID: projectID,
		Client:    client,
		TopicName: topicName,
	}, nil
}

// Write pushes all events to the required target
func (ps *PubSubTarget) Write(events []*Event) error {
	ctx := context.Background()

	topic := ps.Client.Topic(ps.TopicName)
	defer topic.Stop()

	var results []*pubsub.PublishResult

	log.Infof("Writing %d records to target topic '%s' in project %s ...", len(events), ps.TopicName, ps.ProjectID)

	for _, event := range events {
		msg := &pubsub.Message{
			Data: event.Data,
		}

		r := topic.Publish(ctx, msg)
		results = append(results, r)
	}

	successes := 0
	failures := 0
	var errstrings []string

	for _, r := range results {
		_, err := r.Get(ctx)

		if err != nil {
			errstrings = append(errstrings, err.Error())
			failures++
		} else {
			successes++
		}
	}

	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	log.Infof("Successfully wrote %d / %d records to topic '%s' in project %s", successes, len(events), ps.TopicName, ps.ProjectID)

	return nil
}