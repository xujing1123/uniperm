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

package role

import (
	"github.com/go-the-way/unilog"
	"github.com/go-the-way/uniperm/internal/services/base"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by"` // 排序

		Id          uint   `form:"id"`           // 角色id
		Name        string `form:"name"`         // 角色名称
		Description string `form:"description"`  // 角色描述
		Type        string `form:"type"`         // 角色类型
		State       byte   `form:"state"`        // 状态：1启用2禁用
		CreateTime1 string `form:"create_time1"` // 创建时间
		CreateTime2 string `form:"create_time2"` // 创建时间
		UpdateTime1 string `form:"update_time1"` // 修改时间
		UpdateTime2 string `form:"update_time2"` // 修改时间
	}
	IdReq struct {
		Id uint `validate:"min(1,角色Id不能为空)" form:"id" json:"id" log:"角色Id"`
	}
	GetReq        IdReq
	GetPermReq    IdReq
	UpdatePermReq struct {
		IdReq       `validate:"valid(T)"`
		Permissions []uint `validate:"minlength(1,权限Id不能为空)" json:"permissions" log:"权限Id"`
		unilog.UC
		Callback func(req UpdatePermReq)
	}
	AddReq struct {
		Name        string `validate:"minlength(1,角色名称不能为空) maxlength(50,角色名称长度不能超过50)" json:"name" log:"角色名称"` // 角色名称
		Description string `validate:"maxlength(200,角色描述长度不能超过200)" json:"description" log:"角色描述"`              // 角色描述
		Type        string `validate:"maxlength(50,角色类型长度不能超过50)" json:"type" log:"角色类型"`                       // 角色类型
		unilog.UC
		Callback func(req AddReq)
	}
	UpdateReq struct {
		IdReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
		unilog.UC
		Callback func(req UpdateReq)
	}
	DeleteReq struct {
		IdReq
		unilog.UC
		Callback func(req DeleteReq)
	}
	EnableReq struct {
		IdReq `validate:"valid(T)"`
		unilog.UC
		Callback func(req EnableReq)
	}
	DisableReq struct {
		IdReq `validate:"valid(T)"`
		unilog.UC
		Callback func(req DisableReq)
	}
)
