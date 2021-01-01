// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020-2021 Snowplow Analytics Ltd. All rights reserved.

package source

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
	"os"
	"time"

	"github.com/snowplow-devops/stream-replicator/internal/common"
	"github.com/snowplow-devops/stream-replicator/internal/models"
	"github.com/snowplow-devops/stream-replicator/internal/source/sourceiface"
)

// PubSubSource holds a new client for reading messages from PubSub
type PubSubSource struct {
	projectID      string
	client         *pubsub.Client
	subscriptionID string

	log *log.Entry

	// cancel function to be used to halt reading
	cancel context.CancelFunc
}

// NewPubSubSource creates a new client for reading messages from PubSub
func NewPubSubSource(projectID string, subscriptionID string, serviceAccountB64 string) (*PubSubSource, error) {
	if serviceAccountB64 != "" {
		targetFile, err := common.GetGCPServiceAccountFromBase64(serviceAccountB64)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to store GCP Service Account JSON file")
		}
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", targetFile)
	}

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create PubSub client")
	}

	return &PubSubSource{
		projectID:      projectID,
		client:         client,
		subscriptionID: subscriptionID,
		log:            log.WithFields(log.Fields{"source": "pubsub", "cloud": "GCP", "project": projectID, "subscription": subscriptionID}),
	}, nil
}

// Read will pull messages from the noted PubSub topic forever
func (ps *PubSubSource) Read(sf *sourceiface.SourceFunctions) error {
	ctx := context.Background()

	ps.log.Info("Reading messages from subscription ...")

	sub := ps.client.Subscription(ps.subscriptionID)
	cctx, cancel := context.WithCancel(ctx)

	// Store reference to cancel
	ps.cancel = cancel

	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		timePulled := time.Now().UTC()

		ps.log.Debugf("Read message with ID: %s", msg.ID)
		ackFunc := func() {
			ps.log.Debugf("Ack'ing message with ID: %s", msg.ID)
			msg.Ack()
		}

		timeCreated := msg.PublishTime.UTC()
		messages := []*models.Message{
			{
				Data:         msg.Data,
				PartitionKey: uuid.NewV4().String(),
				AckFunc:      ackFunc,
				TimeCreated:  timeCreated,
				TimePulled:   timePulled,
			},
		}
		err := sf.WriteToTarget(messages)
		if err != nil {
			ps.log.WithFields(log.Fields{"error": err}).Error(err)
		}
	})

	if err != nil {
		return errors.Wrap(err, "Failed to read from PubSub topic")
	}
	return nil
}

// Stop will halt the reader processing more events
func (ps *PubSubSource) Stop() {
	if ps.cancel != nil {
		ps.log.Warn("Cancelling PubSub receive ...")
		ps.cancel()
	}
	ps.cancel = nil
}