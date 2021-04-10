// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package snippet

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
)

// Search list snippets
func Search(ctx *context.APIContext) {
	search(ctx, ctx.User)
}

func search(ctx *context.APIContext, u *models.User) {
	response := struct {
		Ok   bool          `json:"ok"`
		Data []interface{} `json:"data"`
	}{
		Ok: true,
	}

	ctx.JSON(http.StatusOK, response)
}
