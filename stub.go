// Copyright 2025 icmpacket Author. All Rights Reserved.
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

var (
	_ = SetDB
	_ = SetPagination
	_ = AutoMigrate

	_ = UserGetPage
	_ = UserGet
	_ = UserGetPerm
	_ = UserGetPermButton
	_ = UserAdd
	_ = UserUpdate
	_ = UserUpdatePassword
	_ = UserUpdateRole
	_ = UserDelete
	_ = UserEnable
	_ = UserDisable
	_ = UserLogin
	_ = UserLogout

	_ = RoleGetPage
	_ = RoleGet
	_ = RoleGetPerm
	_ = RoleUpdatePerm
	_ = RoleAdd
	_ = RoleUpdate
	_ = RoleDelete
	_ = RoleEnable
	_ = RoleDisable

	_ = PermissionTree
	_ = PermissionGet
	_ = PermissionAdd
	_ = PermissionUpdate
	_ = PermissionDelete
)
