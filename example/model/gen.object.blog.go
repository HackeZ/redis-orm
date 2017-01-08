package model

import (
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
	"time"
)

var (
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type Blog struct {
	Id        int32     `db:"id"`
	UserId    int32     `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Status    int32     `db:"status"`
	Readed    int32     `db:"readed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

//! object function

func (obj *Blog) GetNameSpace() string {
	return "model"
}

func (obj *Blog) GetClassName() string {
	return "Blog"
}

func (obj *Blog) GetTableName() string {
	return "blogs"
}

func (obj *Blog) GetColumns() []string {
	columns := []string{
		"`id`",
		"`user_id`",
		"`title`",
		"`content`",
		"`status`",
		"`readed`",
		"`created_at`",
		"`updated_at`",
	}
	return columns
}

//! uniques

//! indexes

type UserIdOfBlogIndex struct {
	UserId int32
	offset int
	limit  int
}

func (u *UserIdOfBlogIndex) Key() string {
	strs := []string{
		"UserId",
		fmt.Sprint(u.UserId),
	}
	return fmt.Sprintf("index:%s", strings.Join(strs, ":"))
}

func (u *UserIdOfBlogIndex) SQLFormat() string {
	conditions := []string{
		"user_id = ?",
	}
	return fmt.Sprintf("%s %s", strings.Join(conditions, " AND "), orm.OffsetLimit(u.offset, u.limit))
}

func (u *UserIdOfBlogIndex) SQLParams() []interface{} {
	return []interface{}{
		u.UserId,
	}
}

func (u *UserIdOfBlogIndex) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *UserIdOfBlogIndex) Limit(n int) {
	u.limit = n
}

func (u *UserIdOfBlogIndex) Offset(n int) {
	u.offset = n
}

//! ranges

//! orders
