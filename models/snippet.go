// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "code.gitea.io/gitea/modules/timeutil"

const (
	visibilityPublic  = "public"
	visibilityLimited = "limited"
	visibilityPrivate = "private"
)

type Snippet struct {
	ID         string `xorm:"INDEX"`
	OwnerID    int64  `xorm:"INDEX"`
	Owner      *User  `xorm:"-"`
	Title      string
	Visibility string
	Revision   int64
	NumStars   int

	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`
}

func (snippet *Snippet) IsVisibilityPublic() bool {
	return visibilityPublic == snippet.Visibility
}

func (snippet *Snippet) IsVisibilityLimited() bool {
	return visibilityLimited == snippet.Visibility
}

func (snippet *Snippet) IsVisibilityPrivate() bool {
	return visibilityPrivate == snippet.Visibility
}
