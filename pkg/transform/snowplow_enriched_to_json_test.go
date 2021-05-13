// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020-2021 Snowplow Analytics Ltd. All rights reserved.

package transform

import (
	"testing"

	"github.com/snowplow-devops/stream-replicator/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestSpEnrichedToJson(t *testing.T) {
	assert := assert.New(t)

	var messageGood = models.Message{
		Data:         snowplowTsv1,
		PartitionKey: "some-key",
	}

	var messageBad = models.Message{
		Data:         nonSnowplowString,
		PartitionKey: "some-key4",
	}

	var expectedGood = models.Message{
		Data:              snowplowJson1,
		PartitionKey:      "some-key",
		IntermediateState: spTsv1Parsed,
	}

	// Simple success case
	transformSuccess, failure := SpEnrichedToJson(&messageGood)

	assert.Equal(&expectedGood, transformSuccess)
	assert.Nil(failure)

	// Simple failure case
	success, transformFailure := SpEnrichedToJson(&messageBad)

	// Not matching equivalence of whole object because error stacktrace makes it unfeasible. Doing each component part instead.
	assert.Equal("Cannot parse tsv event - wrong number of fields provided: 4", transformFailure.GetError().Error())
	assert.Equal([]byte("not	a	snowplow	event"), transformFailure.Data)
	assert.Equal("some-key4", transformFailure.PartitionKey)
	// Failure in this case is in parsing to IntermediateState, so none expected in output
	assert.Nil(transformFailure.IntermediateState)
	assert.Nil(success)

	// Check that the input has not been altered
	assert.Nil(messageGood.GetError())

	// Nuanced success case
	// Test to assert behaviour when there's an incompatible IntermediateState in the input
	incompatibleIntermediateMessage := models.Message{
		Data:              snowplowTsv1,
		PartitionKey:      "some-key",
		IntermediateState: "Incompatible intermediate state",
	}

	// When we have some incompatible IntermediateState, expected behaviour is to replace it with this transformation's IntermediateState
	transformSuccess2, failure2 := SpEnrichedToJson(&incompatibleIntermediateMessage)

	assert.Equal(&expectedGood, transformSuccess2)
	assert.Nil(failure2)
}
