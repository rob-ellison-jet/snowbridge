// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.

package core

// Target describes the interface for how to push the data pulled from the source
type Target interface {
	Write(events []*Event) (*TargetWriteResult, error)
	Close()
}
