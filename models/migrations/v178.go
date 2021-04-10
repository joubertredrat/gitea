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
	type Snippetao struct {
		ID          int64 `xorm:"pk autoincr"`
		OwnerID     int64 `xorm:"INDEX"`
		Name        string
		Title       string
		Visibility  string
		Revision    int64
		NumStars    int64
		CreatedUnix time.Time `xorm:"INDEX created"`
		UpdatedUnix time.Time `xorm:"INDEX updated"`
	}

	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	if err := sess.Sync2(new(Snippetao)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}

	return sess.Commit()
}
