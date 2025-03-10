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

package db

import (
	"github.com/go-the-way/uniperm/internal/models"
	"gorm.io/gorm"
)

type PaginationFunc func(db *gorm.DB, page, limit int, count *int64, list any) (err error)

var (
	gdb      *gorm.DB
	pageFunc PaginationFunc
)

func SetDB(db *gorm.DB)                           { gdb = db }
func GetDB() *gorm.DB                             { return gdb }
func SetPagination(paginationFunc PaginationFunc) { pageFunc = paginationFunc }
func GetPagination() PaginationFunc               { return pageFunc }

func AutoMigrate() (err error) {
	return gdb.AutoMigrate(
		new(models.User),
		new(models.Role),
		new(models.RolePermission),
		new(models.Permission),
	)
}
