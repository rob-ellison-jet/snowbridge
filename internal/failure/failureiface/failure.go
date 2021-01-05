// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020-2021 Snowplow Analytics Ltd. All rights reserved.

package failureiface

import (
	"github.com/snowplow-devops/stream-replicator/internal/models"
)

// Failure describes the interface for where to push data that
// cannot be sent to its desired target and therefore should no longer be retried.
//
// This can be for:
// 1. Oversized records
type Failure interface {
	WriteOversized(maximumAllowedSizeBytes int, messages []*models.Message) (*models.TargetWriteResult, error)
	Open()
	Close()
}
