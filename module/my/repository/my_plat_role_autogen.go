package repository

import (
	"context"
	"database/sql"
	"deepsea/config/constant"
	"deepsea/helper"
	"deepsea/helper/database"
	"deepsea/helper/recorderx"
	"deepsea/model/entity"
	"deepsea/model/iris"
	"github.com/gin-gonic/gin"
)

// iMyPlatRoleAutoGen 该接口自动生成, 请勿修改
type iMyPlatRoleAutoGen interface {
	// SelectOneByConfig Use config service to execute(select statement)
	SelectOneByConfig(iris.ConfigService[entity.MyPlatRole]) *entity.MyPlatRole

	// SelectManyByConfig Use config service to execute(select statement)
	SelectManyByConfig(iris.ConfigService[entity.MyPlatRole]) []*entity.MyPlatRole

	// SelectPageByConfig Use config service to execute(select statement)
	SelectPageByConfig(iris.ConfigService[entity.MyPlatRole]) ([]*entity.MyPlatRole, int64)

	// InsertByConfig Use config service to execute(insert statement)
	InsertByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatRole]) bool

	// UpdateByConfig Use config service to execute(update statement)
	UpdateByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatRole]) bool

	// DeleteByConfig Use config service to execute(delete statement)
	DeleteByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatRole]) bool

	// SelectByID Query and return the related entries of id(p key)
	SelectByID(int64) *entity.MyPlatRole

	// SelectByIDs Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID
	SelectByIDs(...int64) []*entity.MyPlatRole

	// BatchSelectByID Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one
	BatchSelectByID([]int64) []*entity.MyPlatRole

	// SelectByPlatID Query by Plat ID
	SelectByPlatID(int64) []*entity.MyPlatRole

	// SelectByTenantID Query by Tenant ID
	SelectByTenantID(int64) []*entity.MyPlatRole

	// SelectByName Query by Name
	SelectByName(string) []*entity.MyPlatRole

	// SelectAllWithPage Query all entries by page
	SelectAllWithPage(int64, int64) ([]*entity.MyPlatRole, int64)

	// Insert Matches all columns and insert into the table
	Insert(*sql.Tx, *entity.MyPlatRole) int64

	// InsertNonNil Matches non-nil columns and insert into the table
	InsertNonNil(*sql.Tx, *entity.MyPlatRole) int64

	// InsertWithFunc Matches all columns matching this function and insert into the table
	InsertWithFunc(*sql.Tx, *entity.MyPlatRole, func(*iris.Column[entity.MyPlatRole], any) bool) int64

	// BatchInsert Batch insert all columns of table into table
	BatchInsert(*sql.Tx, []*entity.MyPlatRole) []int64

	// BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table
	BatchInsertNonNil(*sql.Tx, []*entity.MyPlatRole) []int64

	// BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table
	BatchInsertWithFunc(*sql.Tx, []*entity.MyPlatRole, func(*iris.Column[entity.MyPlatRole], any) bool) []int64

	// DeleteByID Delete the related entries of id(p key)
	DeleteByID(*sql.Tx, int64) bool

	// DeleteByIDs Delete the related entries of id(p key list)
	DeleteByIDs(*sql.Tx, ...int64) bool

	// BatchDeleteByID Batch delete the related entries of id(p key list)
	BatchDeleteByID(*sql.Tx, []int64) bool

	// UpdateByID Update all columns of the related entries of id(p key)
	UpdateByID(*sql.Tx, *entity.MyPlatRole) bool

	// UpdateNonNilByID Update non-nil columns of the related entries of id(p key)
	UpdateNonNilByID(*sql.Tx, *entity.MyPlatRole) bool

	// UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)
	UpdateByIDWithFunc(*sql.Tx, *entity.MyPlatRole, func(*iris.Column[entity.MyPlatRole], any) bool) bool

	// BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)
	BatchUpdateByIDWithFunc(*sql.Tx, []*entity.MyPlatRole, func(*iris.Column[entity.MyPlatRole], any) bool) bool
}

// myPlatRoleAutoGen 该结构体自动生成, 请勿修改
type myPlatRoleAutoGen struct {
	ctx *gin.Context
}

func (ag *myPlatRoleAutoGen) SelectOneByConfig(config iris.ConfigService[entity.MyPlatRole]) *entity.MyPlatRole {
	if config == nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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
	eto := helper.Row(row, func() (*entity.MyPlatRole, []any) {
		eto := entity.NewMyPlatRole()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return eto
}

func (ag *myPlatRoleAutoGen) SelectManyByConfig(config iris.ConfigService[entity.MyPlatRole]) []*entity.MyPlatRole {
	if config == nil {
		return nil
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatRole, []any) {
		eto := entity.NewMyPlatRole()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return ets
}

func (ag *myPlatRoleAutoGen) SelectPageByConfig(config iris.ConfigService[entity.MyPlatRole]) ([]*entity.MyPlatRole, int64) {
	if config == nil {
		return nil, 0
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatRole, []any) {
		eto := entity.NewMyPlatRole()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	if evalInfo.Pageable() {
		totalSQL := evalInfo.TotalSQL()
		stmt, err := tx.Prepare(totalSQL)
		defer helper.DeferClose(stmt, recorder.MaybePanic)
		recorder.MaybePanic(err)
		row := stmt.QueryRowContext(ag.getDBCtx(), values[:len(values)-2]...)
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

func (ag *myPlatRoleAutoGen) InsertByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatRole]) bool {
	if config == nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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

func (ag *myPlatRoleAutoGen) UpdateByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatRole]) bool {
	if config == nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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

func (ag *myPlatRoleAutoGen) DeleteByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatRole]) bool {
	if config == nil {
		return false
	}
	evaluator := config.Evaluator()
	if evaluator == nil {
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

func (ag *myPlatRoleAutoGen) SelectByID(id int64) *entity.MyPlatRole {
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.SelectOneByConfig(config)
}

func (ag *myPlatRoleAutoGen) SelectByIDs(ids ...int64) []*entity.MyPlatRole {
	return ag.BatchSelectByID(ids)
}

func (ag *myPlatRoleAutoGen) BatchSelectByID(ids []int64) []*entity.MyPlatRole {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatRoleAutoGen) SelectByPlatID(id int64) []*entity.MyPlatRole {
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PlatIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatRoleAutoGen) SelectByTenantID(id int64) []*entity.MyPlatRole {
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.TenantIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatRoleAutoGen) SelectByName(val string) []*entity.MyPlatRole {
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.NameCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatRoleAutoGen) SelectAllWithPage(page int64, size int64) ([]*entity.MyPlatRole, int64) {
	if page < 1 {
		return nil, 0
	}
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(size, (page-1)*size).Eval()
	})
	return ag.SelectPageByConfig(config)
}

func (ag *myPlatRoleAutoGen) Insert(tx *sql.Tx, eto *entity.MyPlatRole) int64 {
	return ag.InsertWithFunc(tx, eto, func(*iris.Column[entity.MyPlatRole], any) bool {
		return true
	})
}

func (ag *myPlatRoleAutoGen) InsertNonNil(tx *sql.Tx, eto *entity.MyPlatRole) int64 {
	return ag.InsertWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatRole], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatRoleAutoGen) InsertWithFunc(tx *sql.Tx, eto *entity.MyPlatRole, fn func(*iris.Column[entity.MyPlatRole], any) bool) int64 {
	if eto.Evaluator() != nil {
		if !ag.InsertByConfig(tx, eto) {
			return 0
		}
		return *eto.ID
	}
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return 0
	}
	return *config.ID
}

func (ag *myPlatRoleAutoGen) BatchInsert(tx *sql.Tx, ets []*entity.MyPlatRole) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatRole], v any) bool {
		return true
	})
}

func (ag *myPlatRoleAutoGen) BatchInsertNonNil(tx *sql.Tx, ets []*entity.MyPlatRole) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatRole], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatRoleAutoGen) BatchInsertWithFunc(tx *sql.Tx, ets []*entity.MyPlatRole, fn func(*iris.Column[entity.MyPlatRole], any) bool) []int64 {
	if len(ets) == 0 {
		return nil
	}
	eto := ets[0]
	ids := make([]int64, len(ets))
	for i, e := range ets {
		ids[i] = *e.ID
	}
	if eto.Evaluator() != nil {
		if !ag.InsertByConfig(tx, eto) {
			return nil
		}
		return ids
	}
	values := make([]any, 0)
	config := eto
	selfishs, _ := eto.ColumnAndValue(fn)
	for _, e := range ets {
		_, snipValues := e.ColumnAndValue(fn)
		values = append(values, snipValues...)
	}
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return nil
	}
	return ids
}

func (ag *myPlatRoleAutoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.DeleteByIDs(tx, id)
}

func (ag *myPlatRoleAutoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

func (ag *myPlatRoleAutoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatRole()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.Delete().From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.DeleteByConfig(tx, config)
}

func (ag *myPlatRoleAutoGen) UpdateByID(tx *sql.Tx, eto *entity.MyPlatRole) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatRole], v any) bool {
		return true
	})
}

func (ag *myPlatRoleAutoGen) UpdateNonNilByID(tx *sql.Tx, eto *entity.MyPlatRole) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatRole], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatRoleAutoGen) UpdateByIDWithFunc(tx *sql.Tx, eto *entity.MyPlatRole, fn func(*iris.Column[entity.MyPlatRole], any) bool) bool {
	if eto.Evaluator() != nil {
		return ag.UpdateByConfig(tx, eto)
	}
	id := *eto.ID
	eto.ID = nil
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatRole]) {
		eval.UpdateRef(config.Table(), selfishs...).SetValues(values...).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.UpdateByConfig(tx, config)
}

func (ag *myPlatRoleAutoGen) BatchUpdateByIDWithFunc(tx *sql.Tx, ets []*entity.MyPlatRole, fn func(*iris.Column[entity.MyPlatRole], any) bool) bool {
	for _, eto := range ets {
		ag.UpdateByIDWithFunc(tx, eto, fn)
	}
	return true
}

func (ag *myPlatRoleAutoGen) getDBCtx() context.Context {
	if ag.ctx == nil {
		return context.Background()
	}
	return context.WithValue(context.Background(), constant.TraceIdKey, ag.ctx.GetString(constant.TraceIdKey))
}
