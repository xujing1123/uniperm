// Copyright 2025 unilog Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package role

import "github.com/go-the-way/unilog"

func (req *UpdatePermReq) LogName() (name string) { return "修改角色权限" }
func (req *AddReq) LogName() (name string)        { return "新增角色" }
func (req *UpdateReq) LogName() (name string)     { return "修改角色" }
func (req *DeleteReq) LogName() (name string)     { return "删除角色" }
func (req *EnableReq) LogName() (name string)     { return "启用角色" }
func (req *DisableReq) LogName() (name string)    { return "禁用角色" }

func (req *UpdatePermReq) LogFields() (fields unilog.FieldSlice) { return unilog.GetFieldsFromTag(req) }
func (req *AddReq) LogFields() (fields unilog.FieldSlice)        { return unilog.GetFieldsFromTag(req) }
func (req *UpdateReq) LogFields() (fields unilog.FieldSlice)     { return unilog.GetFieldsFromTag(req) }
func (req *DeleteReq) LogFields() (fields unilog.FieldSlice)     { return unilog.GetFieldsFromTag(req) }
func (req *EnableReq) LogFields() (fields unilog.FieldSlice)     { return unilog.GetFieldsFromTag(req) }
func (req *DisableReq) LogFields() (fields unilog.FieldSlice)    { return unilog.GetFieldsFromTag(req) }
