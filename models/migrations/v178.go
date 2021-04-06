// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

// createSnippetTable create snippet table for snippet feature
func createSnippetTable(x *xorm.Engine) error {
	type Snippet struct {
		ID          string `xorm:"INDEX"`
		OwnerID     int64  `xorm:"INDEX"`
		Title       string
		Visibility  string
		Revision    int64
		NumStars    int
		CreatedUnix time.Time `xorm:"INDEX created"`
		UpdatedUnix time.Time `xorm:"INDEX updated"`
	}

	if err := x.Sync(new(Snippet)); err != nil {
		return fmt.Errorf("Error creating snippet table: %v", err)
	}

	return nil
}
