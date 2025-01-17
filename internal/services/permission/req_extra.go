// Copyright 2025 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package permission

import (
	"github.com/go-the-way/uniperm/internal/models"
	"github.com/go-the-way/uniperm/internal/services/base"
)

func (req *AddReq) Check() (err error) {
	if req.ParentId > 0 {
		return base.CheckPermissionIsNotButton(req.ParentId)
	}
	return
}

func (req *UpdateReq) Check() (err error) { return base.CheckPermissionExist(req.Id) }

func (req *DeleteReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckPermissionExist(req.Id) },
		func() (err error) { return base.CheckPermissionHaveNoSubPerms(req.Id) },
		func() (err error) { return base.CheckPermissionRefRole(req.Id) },
	)
}

func (req *AddReq) transform() *models.Permission {
	return &models.Permission{Name: req.Name, Route: req.Route, ParentId: req.ParentId, IsButton: req.IsButton}
}

func (req *UpdateReq) transform() map[string]any {
	return map[string]any{"name": req.Name, "route": req.Route}
}
