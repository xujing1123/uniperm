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

package uniperm

import (
	"github.com/go-the-way/uniperm/internal/db"
	"github.com/go-the-way/uniperm/internal/services/permission"
	"github.com/go-the-way/uniperm/internal/services/role"
	"github.com/go-the-way/uniperm/internal/services/user"
)

var (
	SetDB         = db.SetDB
	SetPagination = db.SetPagination
	AutoMigrate   = db.AutoMigrate
)

var (
	UserGetPage        = user.GetPage
	UserGet            = user.Get
	UserGetPerm        = user.GetPerm
	UserGetPermButton  = user.GetPermButton
	UserAdd            = user.Add
	UserUpdate         = user.Update
	UserUpdatePassword = user.UpdatePassword
	UserUpdateRole     = user.UpdateRole
	UserDelete         = user.Delete
	UserEnable         = user.Enable
	UserDisable        = user.Disable
	UserLogin          = user.Login
	UserLogout         = user.Logout
)

var (
	RoleGetPage    = role.GetPage
	RoleGet        = role.Get
	RoleGetPerm    = role.GetPerm
	RoleUpdatePerm = role.UpdatePerm
	RoleAdd        = role.Add
	RoleUpdate     = role.Update
	RoleDelete     = role.Delete
	RoleEnable     = role.Enable
	RoleDisable    = role.Disable
)

var (
	PermissionTree   = permission.Tree
	PermissionGet    = permission.Get
	PermissionAdd    = permission.Add
	PermissionUpdate = permission.Update
	PermissionDelete = permission.Delete
)
