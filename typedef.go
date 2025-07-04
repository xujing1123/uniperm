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
	"github.com/go-the-way/uniperm/internal/models"
	"github.com/go-the-way/uniperm/internal/services/permission"
	"github.com/go-the-way/uniperm/internal/services/role"
	"github.com/go-the-way/uniperm/internal/services/user"
)

type PaginationFunc = db.PaginationFunc

type (
	User           = models.User
	Role           = models.Role
	Permission     = models.Permission
	RolePermission = models.RolePermission
)

type (
	UserGetPageReq        = user.GetPageReq
	UserIdReq             = user.IdReq
	UserGetReq            = user.GetReq
	UserGetPermReq        = user.GetPermReq
	UserGetPermButtonReq  = user.GetPermButtonReq
	UserAddReq            = user.AddReq
	UserUpdateReq         = user.UpdateReq
	UserUpdatePasswordReq = user.UpdatePasswordReq
	UserUpdateRoleReq     = user.UpdateRoleReq
	UserDeleteReq         = user.DeleteReq
	UserEnableReq         = user.EnableReq
	UserDisableReq        = user.DisableReq
	UserLoginReqCallback  = user.LoginReqCallback
	UserLoginReq          = user.LoginReq
	UserLogoutReq         = user.LogoutReq

	UserGetPageResp       = user.GetPageResp
	UserGetResp           = user.GetResp
	UserGetPermResp       = user.GetPermResp
	UserGetPermRespRoute  = user.GetPermRespRoute
	UserGetPermButtonResp = user.GetPermButtonResp
	UserLoginResp         = user.LoginResp
)

type (
	RoleGetPageReq    = role.GetPageReq
	RoleIdReq         = role.IdReq
	RoleGetReq        = role.GetReq
	RoleGetPermReq    = role.GetPermReq
	RoleUpdatePermReq = role.UpdatePermReq
	RoleAddReq        = role.AddReq
	RoleUpdateReq     = role.UpdateReq
	RoleDeleteReq     = role.DeleteReq
	RoleEnableReq     = role.EnableReq
	RoleDisableReq    = role.DisableReq

	RoleGetPageResp = role.GetPageResp
	RoleGetResp     = role.GetResp
	RoleGetPermResp = role.GetPermResp
)

type (
	PermissionTreeReq   = permission.TreeReq
	PermissionIdReq     = permission.IdReq
	PermissionGetReq    = permission.GetReq
	PermissionAddReq    = permission.AddReq
	PermissionUpdateReq = permission.UpdateReq
	PermissionDeleteReq = permission.DeleteReq

	PermissionTreeResp = permission.TreeResp
	PermissionGetResp  = permission.GetResp
)
