// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package snippet

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplCreate base.TplName = "snippet/create"
)

// Create render creating repository page
func Create(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("new_snippet")

	ctx.HTML(200, tplCreate)
}
