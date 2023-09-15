// Package repository
// @author tabuyos
// @since 2023/9/12
// @description repository
package repository

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"metis/helper"
	"metis/helper/database"
	"metis/helper/recorderx"
	"metis/old/model/3.0/entity"
	iris2 "metis/old/model/3.0/iris"
)

// iAccountAutoGen 该接口自动生成, 请勿修改
type iAccountAutoGen interface {
	// SelectOneByConfig Use config service to execute(select statement)
	SelectOneByConfig(iris2.ConfigService[entity.Account]) *entity.Account
	// SelectManyByConfig Use config service to execute(select statement)
	SelectManyByConfig(iris2.ConfigService[entity.Account]) []*entity.Account
	// SelectPageByConfig Use config service to execute(select statement)
	SelectPageByConfig(iris2.ConfigService[entity.Account]) ([]*entity.Account, int64)
	// InsertByConfig Use config service to execute(insert statement)
	InsertByConfig(*sql.Tx, iris2.ConfigService[entity.Account]) bool
	// UpdateByConfig Use config service to execute(update statement)
	UpdateByConfig(*sql.Tx, iris2.ConfigService[entity.Account]) bool
	// DeleteByConfig Use config service to execute(delete statement)
	DeleteByConfig(*sql.Tx, iris2.ConfigService[entity.Account]) bool

	// SelectByID Query and return the related entries of id(p key)
	SelectByID(int64) *entity.Account
	// SelectByIDs Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID
	SelectByIDs(...int64) []*entity.Account
	// BatchSelectByID Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one
	BatchSelectByID([]int64) []*entity.Account

	// SelectByXXXID Query and return the related entries of XXX id(f key)
	SelectByXXXID(int64) []*entity.Account

	// SelectByXXXStr Query and return the related entries of XXX(col type is string in golang), use 'LIKE' statement
	SelectByXXXStr(string) []*entity.Account

	// SelectAllWithPage Query all entries by page
	SelectAllWithPage(int64, int64) ([]*entity.Account, int64)

	// Insert Matches all columns and insert into the table
	Insert(*sql.Tx, *entity.Account) int64
	// InsertNonNil Matches non-nil columns and insert into the table
	InsertNonNil(*sql.Tx, *entity.Account) int64
	// InsertWithFunc Matches all columns matching this function and insert into the table
	InsertWithFunc(*sql.Tx, *entity.Account, func(*iris2.Column[entity.Account], any) bool) int64
	// BatchInsert Batch insert all columns of table into table
	BatchInsert(*sql.Tx, []*entity.Account) []int64
	// BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table
	BatchInsertNonNil(*sql.Tx, []*entity.Account) []int64
	// BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table
	BatchInsertWithFunc(*sql.Tx, []*entity.Account, func(*iris2.Column[entity.Account], any) bool) []int64

	// DeleteByID Delete the related entries of id(p key)
	DeleteByID(*sql.Tx, int64) bool
	// DeleteByIDs Delete the related entries of id(p key list)
	DeleteByIDs(*sql.Tx, ...int64) bool
	// BatchDeleteByID Batch delete the related entries of id(p key list)
	BatchDeleteByID(*sql.Tx, []int64) bool

	// UpdateByID Update all columns of the related entries of id(p key)
	UpdateByID(*sql.Tx, *entity.Account) bool
	// UpdateNonNilByID Update non-nil columns of the related entries of id(p key)
	UpdateNonNilByID(*sql.Tx, *entity.Account) bool
	// UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)
	UpdateByIDWithFunc(*sql.Tx, *entity.Account, func(*iris2.Column[entity.Account], any) bool) bool
	// BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)
	BatchUpdateByIDWithFunc(*sql.Tx, []*entity.Account, func(*iris2.Column[entity.Account], any) bool) bool
}

// accountAutoGen 该结构体自动生成, 请勿修改
type accountAutoGen struct {
	ctx *gin.Context
}

func (ag *accountAutoGen) SelectOneByConfig(config iris2.ConfigService[entity.Account]) *entity.Account {
	if config != nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	recorder.MaybePanic(err)
	row := stmt.QueryRowContext(ag.getDBCtx(), values...)
	account := helper.Row(row, func() (*entity.Account, []any) {
		account := entity.NewAccount()
		mappers := evalInfo.MapperRows(account)
		return account, mappers
	})
	return account
}

func (ag *accountAutoGen) SelectManyByConfig(config iris2.ConfigService[entity.Account]) []*entity.Account {
	if config != nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	recorder.MaybePanic(err)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	rows, err := stmt.QueryContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	accounts := helper.Rows(rows, func() (*entity.Account, []any) {
		account := entity.NewAccount()
		mappers := evalInfo.MapperRows(account)
		return account, mappers
	})

	return accounts
}

func (ag *accountAutoGen) SelectPageByConfig(config iris2.ConfigService[entity.Account]) ([]*entity.Account, int64) {
	if config != nil {
		return nil, 0
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil, 0
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	tx, err := database.FetchDB().Begin()
	recorder.MaybePanic(err)
	defer helper.HandleTx(tx, recorder.MaybePanic)
	stmt, err := tx.Prepare(execSQL)
	recorder.MaybePanic(err)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	rows, err := stmt.QueryContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	accounts := helper.Rows(rows, func() (*entity.Account, []any) {
		account := entity.NewAccount()
		mappers := evalInfo.MapperRows(account)
		return account, mappers
	})
	if evalInfo.Pageable() {
		totalSQL := evalInfo.TotalSQL()
		stmt, err := tx.Prepare(totalSQL)
		defer helper.DeferClose(stmt, recorder.MaybePanic)
		recorder.MaybePanic(err)
		row := stmt.QueryRowContext(ag.getDBCtx(), values...)
		total := helper.Row(row, func() (**int64, []any) {
			var r *int64
			var cs = []any{&r}
			return &r, cs
		})
		if *total == nil {
			return accounts, 0
		}
		return accounts, **total
	}
	return accounts, 0
}

func (ag *accountAutoGen) InsertByConfig(tx *sql.Tx, config iris2.ConfigService[entity.Account]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()
	self := config.Self()

	stmt, err := tx.Prepare(execSQL)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	recorder.MaybePanic(err)
	result, err := stmt.ExecContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	if self.ID != nil {
		return true
	}
	id, err := result.LastInsertId()
	recorder.MaybePanic(err)
	if id == 0 {
		panic("插入失败")
	}
	self.ID = &id
	return true
}

func (ag *accountAutoGen) UpdateByConfig(tx *sql.Tx, config iris2.ConfigService[entity.Account]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	stmt, err := tx.Prepare(execSQL)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	recorder.MaybePanic(err)
	result, err := stmt.ExecContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	_, err = result.RowsAffected()
	recorder.MaybePanic(err)
	return true
}

func (ag *accountAutoGen) DeleteByConfig(tx *sql.Tx, config iris2.ConfigService[entity.Account]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.DefaultRecorder()
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	stmt, err := tx.Prepare(execSQL)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	recorder.MaybePanic(err)
	result, err := stmt.ExecContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	_, err = result.RowsAffected()
	recorder.MaybePanic(err)
	return true
}

func (ag *accountAutoGen) SelectByID(id int64) *entity.Account {
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.SelectOneByConfig(config)
}

func (ag *accountAutoGen) SelectByIDs(ids ...int64) []*entity.Account {
	return ag.BatchSelectByID(ids)
}

func (ag *accountAutoGen) BatchSelectByID(ids []int64) []*entity.Account {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *accountAutoGen) SelectByXXXID(xxxID int64) []*entity.Account {
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(xxxID)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *accountAutoGen) SelectByXXXStr(str string) []*entity.Account {
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().Like(str)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *accountAutoGen) SelectAllWithPage(page int64, size int64) ([]*entity.Account, int64) {
	if page < 1 {
		return nil, 0
	}
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(size, (page-1)*size).Eval()
	})
	return ag.SelectPageByConfig(config)
}

func (ag *accountAutoGen) Insert(tx *sql.Tx, account *entity.Account) int64 {
	return ag.InsertWithFunc(tx, account, func(*iris2.Column[entity.Account], any) bool {
		return true
	})
}

func (ag *accountAutoGen) InsertNonNil(tx *sql.Tx, account *entity.Account) int64 {
	return ag.InsertWithFunc(tx, account, func(c *iris2.Column[entity.Account], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *accountAutoGen) InsertWithFunc(tx *sql.Tx, account *entity.Account, fn func(*iris2.Column[entity.Account], any) bool) int64 {
	if account.Evaluator() != nil {
		ag.InsertByConfig(tx, account)
		return *account.ID
	}
	config := entity.NewAccount()
	selfishs, values := account.ColumnAndValue(fn)
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Insert(selfishs...).Insert(config.Table()).Values(values...).Eval()
	})
	ag.InsertByConfig(tx, config)
	return *config.ID
}

func (ag *accountAutoGen) BatchInsert(tx *sql.Tx, accounts []*entity.Account) []int64 {
	return ag.BatchInsertWithFunc(tx, accounts, func(c *iris2.Column[entity.Account], v any) bool {
		return true
	})
}

func (ag *accountAutoGen) BatchInsertNonNil(tx *sql.Tx, accounts []*entity.Account) []int64 {
	return ag.BatchInsertWithFunc(tx, accounts, func(c *iris2.Column[entity.Account], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *accountAutoGen) BatchInsertWithFunc(tx *sql.Tx, accounts []*entity.Account, fn func(*iris2.Column[entity.Account], any) bool) []int64 {
	if len(accounts) == 0 {
		return nil
	}
	account := accounts[0]
	ids := make([]int64, len(accounts))
	for i, e := range accounts {
		ids[i] = *e.ID
	}
	if account.Evaluator() != nil {
		ag.InsertByConfig(tx, account)
		return ids
	}
	values := make([]any, 0)
	config := entity.NewAccount()
	selfishs, _ := account.ColumnAndValue(fn)
	for _, e := range accounts {
		_, snipValues := e.ColumnAndValue(fn)
		values = append(values, snipValues...)
	}

	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Insert(selfishs...).Insert(config.Table()).Values(values...).Eval()
	})

	ag.InsertByConfig(tx, config)

	return ids
}

func (ag *accountAutoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.DeleteByIDs(tx, id)
}

func (ag *accountAutoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

func (ag *accountAutoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.Delete().From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.DeleteByConfig(tx, config)
}

func (ag *accountAutoGen) UpdateByID(tx *sql.Tx, account *entity.Account) bool {
	return ag.UpdateByIDWithFunc(tx, account, func(c *iris2.Column[entity.Account], v any) bool {
		return true
	})
}

func (ag *accountAutoGen) UpdateNonNilByID(tx *sql.Tx, account *entity.Account) bool {
	return ag.UpdateByIDWithFunc(tx, account, func(c *iris2.Column[entity.Account], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *accountAutoGen) UpdateByIDWithFunc(tx *sql.Tx, account *entity.Account, fn func(*iris2.Column[entity.Account], any) bool) bool {
	if account.Evaluator() != nil {
		return ag.UpdateByConfig(tx, account)
	}
	config := entity.NewAccount()
	selfishs, values := account.ColumnAndValue(fn)
	config.Configure(func(eval *iris2.Evaluator[entity.Account]) {
		eval.UpdateRef(config.Table(), selfishs...).SetValues(values...).Eval()
	})
	return ag.UpdateByConfig(tx, config)
}

func (ag *accountAutoGen) BatchUpdateByIDWithFunc(tx *sql.Tx, accounts []*entity.Account, fn func(*iris2.Column[entity.Account], any) bool) bool {
	for _, account := range accounts {
		ag.UpdateByIDWithFunc(tx, account, fn)
	}
	return true
}

// getDBCtx 获取 DB 的初始上下文
func (ag *accountAutoGen) getDBCtx() context.Context {
	return context.Background()
}

var _ iAccountAutoGen = (*accountAutoGen)(nil)
