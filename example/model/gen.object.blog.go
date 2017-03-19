package model

import (
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
	"time"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type Blog struct {
	Id        int32     `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Status    int32     `db:"status"`
	Readed    int32     `db:"readed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type _BlogMgr struct {
}

var BlogMgr *_BlogMgr

func (m *_BlogMgr) NewBlog() *Blog {
	return &Blog{}
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
		"`title`",
		"`content`",
		"`status`",
		"`readed`",
		"`created_at`",
		"`updated_at`",
	}
	return columns
}

func (obj *Blog) GetPrimaryKey() PrimaryKey {
	pk := BlogMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

//! primary key

type IdOfBlogPK struct {
	Id int32
}

func (m *_BlogMgr) NewPrimaryKey() *IdOfBlogPK {
	return &IdOfBlogPK{}
}

func (u *IdOfBlogPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfBlogPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vId, ok := kv["Id"]
	if !ok {
		return fmt.Errorf("key (%s) without (Id) field", key)
	}
	if err := orm.StringScan(vId, &(u.Id)); err != nil {
		return err
	}
	return nil
}

func (u *IdOfBlogPK) SQLFormat() string {
	conditions := []string{
		"id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfBlogPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfBlogPK) Columns() []string {
	return []string{
		"`id`",
	}
}

//! uniques

type IdOfBlogUK struct {
	Id int32
}

func (u *IdOfBlogUK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfBlogUK) SQLFormat(limit bool) string {
	conditions := []string{
		"id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfBlogUK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfBlogUK) SQLLimit() int {
	return 1
}

func (u *IdOfBlogUK) Limit(n int) {
}

func (u *IdOfBlogUK) Offset(n int) {
}

func (u *IdOfBlogUK) UKRelation() UniqueRelation {
	return nil
}

//! indexes

type StatusOfBlogIDX struct {
	Status int32
	offset int
	limit  int
}

func (u *StatusOfBlogIDX) Key() string {
	strs := []string{
		"Status",
		fmt.Sprint(u.Status),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *StatusOfBlogIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"status = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *StatusOfBlogIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Status,
	}
}

func (u *StatusOfBlogIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *StatusOfBlogIDX) Limit(n int) {
	u.limit = n
}

func (u *StatusOfBlogIDX) Offset(n int) {
	u.offset = n
}

func (u *StatusOfBlogIDX) IDXRelation() IndexRelation {
	return nil
}

//! ranges

type IdOfBlogRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfBlogRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfBlogRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfBlogRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfBlogRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.beginOp()))
		}
		if u.IdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Id", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Id", u.revert))
}

func (u *IdOfBlogRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			params = append(params, u.IdBegin)
		}
		if u.IdEnd != -1 {
			params = append(params, u.IdEnd)
		}
	}
	return params
}

func (u *IdOfBlogRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfBlogRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfBlogRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfBlogRNG) Begin() int64 {
	start := u.IdBegin
	if start == -1 || start == 0 {
		start = 0
	}
	if start > 0 {
		if !u.includeBegin {
			start = start + 1
		}
	}
	return start
}

func (u *IdOfBlogRNG) End() int64 {
	stop := u.IdEnd
	if stop == 0 || stop == -1 {
		stop = -1
	}
	if stop > 0 {
		if !u.includeBegin {
			stop = stop - 1
		}
	}
	return stop
}

func (u *IdOfBlogRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfBlogRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfBlogRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfBlogRNG) RNGRelation() RangeRelation {
	return nil
}

type _BlogMySQLMgr struct {
	*orm.MySQLStore
}

func BlogMySQLMgr() *_BlogMySQLMgr {
	return &_BlogMySQLMgr{_mysql_store}
}

func NewBlogMySQLMgr(cf *MySQLConfig) (*_BlogMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_BlogMySQLMgr{store}, nil
}

func (m *_BlogMySQLMgr) Search(where string, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	if where != "" {
		where = " WHERE " + where
	}
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), where)
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return m.queryCount(where, args...)
}

func (m *_BlogMySQLMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result Blog
		err = rows.Scan(&(result.Id), &(result.Title), &(result.Content), &(result.Status), &(result.Readed), &CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)
		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
func (m *_BlogMySQLMgr) Fetch(pk PrimaryKey) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (m *_BlogMySQLMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*Blog, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*IdOfBlogPK).Id))
	}
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog find record not found")
}

func (m *_BlogMySQLMgr) FindOneFetch(unique Unique) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("none record")
}

func (m *_BlogMySQLMgr) Find(index Index) ([]PrimaryKey, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_BlogMySQLMgr) FindFetch(index Index) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_BlogMySQLMgr) Range(scope Range) ([]PrimaryKey, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeFetch(scope Range) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeRevertFetch(scope Range) ([]*Blog, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_BlogMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := BlogMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := BlogMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (m *_BlogMySQLMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `blogs` %s", where)
	rows, err := m.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("Blog query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return 0, err
		}
		break
	}
	return count, nil
}

//! tx write
type _BlogMySQLTx struct {
	*orm.MySQLTx
	err          error
	rowsAffected int64
}

func (m *_BlogMySQLMgr) BeginTx(tx *orm.MySQLTx) (*_BlogMySQLTx, error) {
	ux := tx
	if ux == nil {
		tx, err := m.MySQLStore.BeginTx()
		if err != nil {
			return nil, err
		}
		ux = tx
	}
	return &_BlogMySQLTx{ux, nil, 0}, nil
}

func (tx *_BlogMySQLTx) BatchCreate(objs []*Blog) error {
	if len(objs) == 0 {
		return nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*7)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(7, "?"), ",")))
		values = append(values, obj.Id)
		values = append(values, obj.Title)
		values = append(values, obj.Content)
		values = append(values, obj.Status)
		values = append(values, obj.Readed)
		values = append(values, orm.TimeFormat(obj.CreatedAt))
		values = append(values, orm.TimeFormat(obj.UpdatedAt))
	}
	query := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES %s", strings.Join(objs[0].GetColumns(), ","), strings.Join(params, ","))
	result, err := tx.Exec(query, values...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) BatchDelete(objs []*Blog) error {
	for _, obj := range objs {
		if err := tx.Delete(obj); err != nil {
			return err
		}
	}
	return nil
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (tx *_BlogMySQLTx) UpdateBySQL(set, where string, args ...interface{}) error {
	query := fmt.Sprintf("UPDATE `blogs` SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE `blogs` SET %s WHERE %s", set, where)
	}
	result, err := tx.Exec(query, args...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Create(obj *Blog) error {
	params := orm.NewStringSlice(7, "?")
	q := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 7)
	values = append(values, obj.Id)
	values = append(values, obj.Title)
	values = append(values, obj.Content)
	values = append(values, obj.Status)
	values = append(values, obj.Readed)
	values = append(values, orm.TimeFormat(obj.CreatedAt))
	values = append(values, orm.TimeFormat(obj.UpdatedAt))
	result, err := tx.Exec(q, values...)
	if err != nil {
		tx.err = err
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		tx.err = err
		return err
	}
	obj.Id = int32(lastInsertId)
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Update(obj *Blog) error {
	columns := []string{
		"`title` = ?",
		"`content` = ?",
		"`status` = ?",
		"`readed` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE `blogs` SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 7-1)
	values = append(values, obj.Title)
	values = append(values, obj.Content)
	values = append(values, obj.Status)
	values = append(values, obj.Readed)
	values = append(values, orm.TimeFormat(obj.CreatedAt))
	values = append(values, orm.TimeFormat(obj.UpdatedAt))
	values = append(values, pk.SQLParams()...)

	result, err := tx.Exec(q, values...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Save(obj *Blog) error {
	err := tx.Update(obj)
	if err != nil {
		return err
	}
	if tx.rowsAffected > 0 {
		return nil
	}
	return tx.Create(obj)
}

func (tx *_BlogMySQLTx) Delete(obj *Blog) error {
	pk := obj.GetPrimaryKey()
	return tx.DeleteByPrimaryKey(pk)
}

func (tx *_BlogMySQLTx) DeleteByPrimaryKey(pk PrimaryKey) error {
	q := fmt.Sprintf("DELETE FROM `blogs` %s", pk.SQLFormat())
	result, err := tx.Exec(q, pk.SQLParams()...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) DeleteBySQL(where string, args ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM `blogs`")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM `blogs` WHERE %s", where)
	}
	result, err := tx.Exec(query, args...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Close() error {
	if tx.err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//! tx read
func (tx *_BlogMySQLTx) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.err = err
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	tx.err = fmt.Errorf("Blog find record not found")
	return nil, tx.err
}

func (tx *_BlogMySQLTx) FindOneFetch(unique Unique) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("none record")
}

func (tx *_BlogMySQLTx) Find(index Index) ([]PrimaryKey, error) {
	return tx.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_BlogMySQLTx) FindFetch(index Index) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) FindCount(index Index) (int64, error) {
	return tx.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (tx *_BlogMySQLTx) Range(scope Range) ([]PrimaryKey, error) {
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeFetch(scope Range) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) RangeCount(scope Range) (int64, error) {
	return tx.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeRevertFetch(scope Range) ([]*Blog, error) {
	scope.Revert(true)
	return tx.RangeFetch(scope)
}

func (tx *_BlogMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := BlogMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := BlogMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (tx *_BlogMySQLTx) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `blogs` %s", where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.err = err
		return 0, fmt.Errorf("Blog query limit error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			tx.err = err
			return 0, err
		}
		break
	}

	return count, nil
}

func (tx *_BlogMySQLTx) Fetch(pk PrimaryKey) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := tx.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (tx *_BlogMySQLTx) FetchByIds(ids []interface{}) ([]*Blog, error) {
	if len(ids) == 0 {
		return []*Blog{}, nil
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) Search(where string, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	if where != "" {
		where = " WHERE " + where
	}
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), where)
	objs, err := tx.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return tx.queryCount(where, args...)
}

func (tx *_BlogMySQLTx) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := tx.Query(q, args...)
	if err != nil {
		tx.err = err
		return nil, fmt.Errorf("Blog fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result Blog
		err = rows.Scan(&(result.Id), &(result.Title), &(result.Content), &(result.Status), &(result.Readed), &CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)
		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		tx.err = err
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
