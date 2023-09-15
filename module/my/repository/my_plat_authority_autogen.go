package repository

import (
	"context"
	"database/sql"
	"deepsea/helper"
	"deepsea/helper/database"
	"deepsea/helper/recorderx"
	"deepsea/model/entity"
	"deepsea/model/iris"
	"github.com/gin-gonic/gin"
)

// iMyPlatAuthorityAutoGen 该接口自动生成, 请勿修改
type iMyPlatAuthorityAutoGen interface {
	// SelectOneByConfig Use config service to execute(select statement)
	SelectOneByConfig(iris.ConfigService[entity.MyPlatAuthority]) *entity.MyPlatAuthority

	// SelectManyByConfig Use config service to execute(select statement)
	SelectManyByConfig(iris.ConfigService[entity.MyPlatAuthority]) []*entity.MyPlatAuthority

	// SelectPageByConfig Use config service to execute(select statement)
	SelectPageByConfig(iris.ConfigService[entity.MyPlatAuthority]) ([]*entity.MyPlatAuthority, int64)

	// InsertByConfig Use config service to execute(insert statement)
	InsertByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthority]) bool

	// UpdateByConfig Use config service to execute(update statement)
	UpdateByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthority]) bool

	// DeleteByConfig Use config service to execute(delete statement)
	DeleteByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthority]) bool

	// SelectByID Query and return the related entries of id(p key)
	SelectByID(int64) *entity.MyPlatAuthority

	// SelectByIDs Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID
	SelectByIDs(...int64) []*entity.MyPlatAuthority

	// BatchSelectByID Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one
	BatchSelectByID([]int64) []*entity.MyPlatAuthority

	// SelectByTenantID Query by Tenant ID
	SelectByTenantID(int64) []*entity.MyPlatAuthority

	// SelectByPlatID Query by Plat ID
	SelectByPlatID(int64) []*entity.MyPlatAuthority

	// SelectByName Query by Name
	SelectByName(string) []*entity.MyPlatAuthority

	// SelectAllWithPage Query all entries by page
	SelectAllWithPage(int64, int64) ([]*entity.MyPlatAuthority, int64)

	// Insert Matches all columns and insert into the table
	Insert(*sql.Tx, *entity.MyPlatAuthority) int64

	// InsertNonNil Matches non-nil columns and insert into the table
	InsertNonNil(*sql.Tx, *entity.MyPlatAuthority) int64

	// InsertWithFunc Matches all columns matching this function and insert into the table
	InsertWithFunc(*sql.Tx, *entity.MyPlatAuthority, func(*iris.Column[entity.MyPlatAuthority], any) bool) int64

	// BatchInsert Batch insert all columns of table into table
	BatchInsert(*sql.Tx, []*entity.MyPlatAuthority) []int64

	// BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table
	BatchInsertNonNil(*sql.Tx, []*entity.MyPlatAuthority) []int64

	// BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table
	BatchInsertWithFunc(*sql.Tx, []*entity.MyPlatAuthority, func(*iris.Column[entity.MyPlatAuthority], any) bool) []int64

	// DeleteByID Delete the related entries of id(p key)
	DeleteByID(*sql.Tx, int64) bool

	// DeleteByIDs Delete the related entries of id(p key list)
	DeleteByIDs(*sql.Tx, ...int64) bool

	// BatchDeleteByID Batch delete the related entries of id(p key list)
	BatchDeleteByID(*sql.Tx, []int64) bool

	// UpdateByID Update all columns of the related entries of id(p key)
	UpdateByID(*sql.Tx, *entity.MyPlatAuthority) bool

	// UpdateNonNilByID Update non-nil columns of the related entries of id(p key)
	UpdateNonNilByID(*sql.Tx, *entity.MyPlatAuthority) bool

	// UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)
	UpdateByIDWithFunc(*sql.Tx, *entity.MyPlatAuthority, func(*iris.Column[entity.MyPlatAuthority], any) bool) bool

	// BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)
	BatchUpdateByIDWithFunc(*sql.Tx, []*entity.MyPlatAuthority, func(*iris.Column[entity.MyPlatAuthority], any) bool) bool
}

// myPlatAuthorityAutoGen 该结构体自动生成, 请勿修改
type myPlatAuthorityAutoGen struct {
	ctx *gin.Context
}

func (ag *myPlatAuthorityAutoGen) SelectOneByConfig(config iris.ConfigService[entity.MyPlatAuthority]) *entity.MyPlatAuthority {
	if config != nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()
	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	recorder.MaybePanic(err)
	row := stmt.QueryRowContext(ag.getDBCtx(), values...)
	eto := helper.Row(row, func() (*entity.MyPlatAuthority, []any) {
		eto := entity.NewMyPlatAuthority()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return eto
}

func (ag *myPlatAuthorityAutoGen) SelectManyByConfig(config iris.ConfigService[entity.MyPlatAuthority]) []*entity.MyPlatAuthority {
	if config != nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
	evalInfo := evaluator.EvalInfo()
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()
	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	recorder.MaybePanic(err)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	rows, err := stmt.QueryContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	ets := helper.Rows(rows, func() (*entity.MyPlatAuthority, []any) {
		eto := entity.NewMyPlatAuthority()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return ets
}

func (ag *myPlatAuthorityAutoGen) SelectPageByConfig(config iris.ConfigService[entity.MyPlatAuthority]) ([]*entity.MyPlatAuthority, int64) {
	if config != nil {
		return nil, 0
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return nil, 0
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
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
	ets := helper.Rows(rows, func() (*entity.MyPlatAuthority, []any) {
		eto := entity.NewMyPlatAuthority()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
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
			return ets, 0
		}
		return ets, **total
	}
	return ets, 0
}

func (ag *myPlatAuthorityAutoGen) InsertByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthority]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
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

func (ag *myPlatAuthorityAutoGen) UpdateByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthority]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
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

func (ag *myPlatAuthorityAutoGen) DeleteByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthority]) bool {
	if config != nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator != nil {
		return false
	}
	recorder := recorderx.FetchRecorder(ag.ctx)
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

func (ag *myPlatAuthorityAutoGen) SelectByID(id int64) *entity.MyPlatAuthority {
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.SelectOneByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) SelectByIDs(ids ...int64) []*entity.MyPlatAuthority {
	return ag.BatchSelectByID(ids)
}

func (ag *myPlatAuthorityAutoGen) BatchSelectByID(ids []int64) []*entity.MyPlatAuthority {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) SelectByTenantID(id int64) []*entity.MyPlatAuthority {
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.TenantIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) SelectByPlatID(id int64) []*entity.MyPlatAuthority {
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PlatIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) SelectByName(val string) []*entity.MyPlatAuthority {
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.NameCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) SelectAllWithPage(page int64, size int64) ([]*entity.MyPlatAuthority, int64) {
	if page < 1 {
		return nil, 0
	}
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(size, (page-1)*size).Eval()
	})
	return ag.SelectPageByConfig(config)
}

func (ag *myPlatAuthorityAutoGen) Insert(tx *sql.Tx, eto *entity.MyPlatAuthority) int64 {
	return ag.InsertWithFunc(tx, eto, func(*iris.Column[entity.MyPlatAuthority], any) bool {
		return true
	})
}

func (ag *myPlatAuthorityAutoGen) InsertNonNil(tx *sql.Tx, eto *entity.MyPlatAuthority) int64 {
	return ag.InsertWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthority], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityAutoGen) InsertWithFunc(tx *sql.Tx, eto *entity.MyPlatAuthority, fn func(*iris.Column[entity.MyPlatAuthority], any) bool) int64 {
	if eto.Evaluator() != nil {
		ag.InsertByConfig(tx, eto)
		return *eto.ID
	}
	config := entity.NewMyPlatAuthority()
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	ag.InsertByConfig(tx, config)
	return *config.ID
}

func (ag *myPlatAuthorityAutoGen) BatchInsert(tx *sql.Tx, ets []*entity.MyPlatAuthority) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatAuthority], v any) bool {
		return true
	})
}

func (ag *myPlatAuthorityAutoGen) BatchInsertNonNil(tx *sql.Tx, ets []*entity.MyPlatAuthority) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatAuthority], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityAutoGen) BatchInsertWithFunc(tx *sql.Tx, ets []*entity.MyPlatAuthority, fn func(*iris.Column[entity.MyPlatAuthority], any) bool) []int64 {
	if len(ets) == 0 {
		return nil
	}
	eto := ets[0]
	ids := make([]int64, len(ets))
	for i, e := range ets {
		ids[i] = *e.ID
	}
	if eto.Evaluator() != nil {
		ag.InsertByConfig(tx, eto)
		return ids
	}
	values := make([]any, 0)
	config := entity.NewMyPlatAuthority()
	selfishs, _ := eto.ColumnAndValue(fn)
	for _, e := range ets {
		_, snipValues := e.ColumnAndValue(fn)
		values = append(values, snipValues...)
	}
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	ag.InsertByConfig(tx, config)
	return ids
}

func (ag *myPlatAuthorityAutoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.DeleteByIDs(tx, id)
}

func (ag *myPlatAuthorityAutoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

func (ag *myPlatAuthorityAutoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatAuthority()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.Delete().From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.DeleteByConfig(tx, config)
}

func (ag *myPlatAuthorityAutoGen) UpdateByID(tx *sql.Tx, eto *entity.MyPlatAuthority) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthority], v any) bool {
		return true
	})
}

func (ag *myPlatAuthorityAutoGen) UpdateNonNilByID(tx *sql.Tx, eto *entity.MyPlatAuthority) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthority], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityAutoGen) UpdateByIDWithFunc(tx *sql.Tx, eto *entity.MyPlatAuthority, fn func(*iris.Column[entity.MyPlatAuthority], any) bool) bool {
	if eto.Evaluator() != nil {
		return ag.UpdateByConfig(tx, eto)
	}
	config := entity.NewMyPlatAuthority()
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		eval.UpdateRef(config.Table(), selfishs...).SetValues(values...).Eval()
	})
	return ag.UpdateByConfig(tx, config)
}

func (ag *myPlatAuthorityAutoGen) BatchUpdateByIDWithFunc(tx *sql.Tx, ets []*entity.MyPlatAuthority, fn func(*iris.Column[entity.MyPlatAuthority], any) bool) bool {
	for _, eto := range ets {
		ag.UpdateByIDWithFunc(tx, eto, fn)
	}
	return true
}

func (ag *myPlatAuthorityAutoGen) getDBCtx() context.Context {
	return context.Background()
}
