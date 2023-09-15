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

// iMyPlatResourceAutoGen 该接口自动生成, 请勿修改
type iMyPlatResourceAutoGen interface {
	// SelectOneByConfig Use config service to execute(select statement)
	SelectOneByConfig(iris.ConfigService[entity.MyPlatResource]) *entity.MyPlatResource

	// SelectManyByConfig Use config service to execute(select statement)
	SelectManyByConfig(iris.ConfigService[entity.MyPlatResource]) []*entity.MyPlatResource

	// SelectPageByConfig Use config service to execute(select statement)
	SelectPageByConfig(iris.ConfigService[entity.MyPlatResource]) ([]*entity.MyPlatResource, int64)

	// InsertByConfig Use config service to execute(insert statement)
	InsertByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatResource]) bool

	// UpdateByConfig Use config service to execute(update statement)
	UpdateByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatResource]) bool

	// DeleteByConfig Use config service to execute(delete statement)
	DeleteByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatResource]) bool

	// SelectByID Query and return the related entries of id(p key)
	SelectByID(int64) *entity.MyPlatResource

	// SelectByIDs Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID
	SelectByIDs(...int64) []*entity.MyPlatResource

	// BatchSelectByID Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one
	BatchSelectByID([]int64) []*entity.MyPlatResource

	// SelectByPlatID Query by Plat ID
	SelectByPlatID(int64) []*entity.MyPlatResource

	// SelectByTitle Query by Title
	SelectByTitle(string) []*entity.MyPlatResource

	// SelectByName Query by Name
	SelectByName(string) []*entity.MyPlatResource

	// SelectByIcon Query by Icon
	SelectByIcon(string) []*entity.MyPlatResource

	// SelectByPermission Query by Permission
	SelectByPermission(string) []*entity.MyPlatResource

	// SelectByPath Query by Path
	SelectByPath(string) []*entity.MyPlatResource

	// SelectByRouter Query by Router
	SelectByRouter(string) []*entity.MyPlatResource

	// SelectBySort Query by Sort
	SelectBySort(string) []*entity.MyPlatResource

	// SelectAllWithPage Query all entries by page
	SelectAllWithPage(int64, int64) ([]*entity.MyPlatResource, int64)

	// Insert Matches all columns and insert into the table
	Insert(*sql.Tx, *entity.MyPlatResource) int64

	// InsertNonNil Matches non-nil columns and insert into the table
	InsertNonNil(*sql.Tx, *entity.MyPlatResource) int64

	// InsertWithFunc Matches all columns matching this function and insert into the table
	InsertWithFunc(*sql.Tx, *entity.MyPlatResource, func(*iris.Column[entity.MyPlatResource], any) bool) int64

	// BatchInsert Batch insert all columns of table into table
	BatchInsert(*sql.Tx, []*entity.MyPlatResource) []int64

	// BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table
	BatchInsertNonNil(*sql.Tx, []*entity.MyPlatResource) []int64

	// BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table
	BatchInsertWithFunc(*sql.Tx, []*entity.MyPlatResource, func(*iris.Column[entity.MyPlatResource], any) bool) []int64

	// DeleteByID Delete the related entries of id(p key)
	DeleteByID(*sql.Tx, int64) bool

	// DeleteByIDs Delete the related entries of id(p key list)
	DeleteByIDs(*sql.Tx, ...int64) bool

	// BatchDeleteByID Batch delete the related entries of id(p key list)
	BatchDeleteByID(*sql.Tx, []int64) bool

	// UpdateByID Update all columns of the related entries of id(p key)
	UpdateByID(*sql.Tx, *entity.MyPlatResource) bool

	// UpdateNonNilByID Update non-nil columns of the related entries of id(p key)
	UpdateNonNilByID(*sql.Tx, *entity.MyPlatResource) bool

	// UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)
	UpdateByIDWithFunc(*sql.Tx, *entity.MyPlatResource, func(*iris.Column[entity.MyPlatResource], any) bool) bool

	// BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)
	BatchUpdateByIDWithFunc(*sql.Tx, []*entity.MyPlatResource, func(*iris.Column[entity.MyPlatResource], any) bool) bool
}

// myPlatResourceAutoGen 该结构体自动生成, 请勿修改
type myPlatResourceAutoGen struct {
	ctx *gin.Context
}

func (ag *myPlatResourceAutoGen) SelectOneByConfig(config iris.ConfigService[entity.MyPlatResource]) *entity.MyPlatResource {
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
	eto := helper.Row(row, func() (*entity.MyPlatResource, []any) {
		eto := entity.NewMyPlatResource()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return eto
}

func (ag *myPlatResourceAutoGen) SelectManyByConfig(config iris.ConfigService[entity.MyPlatResource]) []*entity.MyPlatResource {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatResource, []any) {
		eto := entity.NewMyPlatResource()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return ets
}

func (ag *myPlatResourceAutoGen) SelectPageByConfig(config iris.ConfigService[entity.MyPlatResource]) ([]*entity.MyPlatResource, int64) {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatResource, []any) {
		eto := entity.NewMyPlatResource()
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

func (ag *myPlatResourceAutoGen) InsertByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatResource]) bool {
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

func (ag *myPlatResourceAutoGen) UpdateByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatResource]) bool {
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

func (ag *myPlatResourceAutoGen) DeleteByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatResource]) bool {
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

func (ag *myPlatResourceAutoGen) SelectByID(id int64) *entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.SelectOneByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByIDs(ids ...int64) []*entity.MyPlatResource {
	return ag.BatchSelectByID(ids)
}

func (ag *myPlatResourceAutoGen) BatchSelectByID(ids []int64) []*entity.MyPlatResource {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByPlatID(id int64) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PlatIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByTitle(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.TitleCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByName(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.NameCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByIcon(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IconCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByPermission(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PermissionCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByPath(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PathCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectByRouter(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.RouterCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectBySort(val string) []*entity.MyPlatResource {
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.SortCol().Like(val)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatResourceAutoGen) SelectAllWithPage(page int64, size int64) ([]*entity.MyPlatResource, int64) {
	if page < 1 {
		return nil, 0
	}
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(size, (page-1)*size).Eval()
	})
	return ag.SelectPageByConfig(config)
}

func (ag *myPlatResourceAutoGen) Insert(tx *sql.Tx, eto *entity.MyPlatResource) int64 {
	return ag.InsertWithFunc(tx, eto, func(*iris.Column[entity.MyPlatResource], any) bool {
		return true
	})
}

func (ag *myPlatResourceAutoGen) InsertNonNil(tx *sql.Tx, eto *entity.MyPlatResource) int64 {
	return ag.InsertWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatResourceAutoGen) InsertWithFunc(tx *sql.Tx, eto *entity.MyPlatResource, fn func(*iris.Column[entity.MyPlatResource], any) bool) int64 {
	if eto.Evaluator() != nil {
		if !ag.InsertByConfig(tx, eto) {
			return 0
		}
		return *eto.ID
	}
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return 0
	}
	return *config.ID
}

func (ag *myPlatResourceAutoGen) BatchInsert(tx *sql.Tx, ets []*entity.MyPlatResource) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatResource], v any) bool {
		return true
	})
}

func (ag *myPlatResourceAutoGen) BatchInsertNonNil(tx *sql.Tx, ets []*entity.MyPlatResource) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatResourceAutoGen) BatchInsertWithFunc(tx *sql.Tx, ets []*entity.MyPlatResource, fn func(*iris.Column[entity.MyPlatResource], any) bool) []int64 {
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
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return nil
	}
	return ids
}

func (ag *myPlatResourceAutoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.DeleteByIDs(tx, id)
}

func (ag *myPlatResourceAutoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

func (ag *myPlatResourceAutoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.Delete().From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.DeleteByConfig(tx, config)
}

func (ag *myPlatResourceAutoGen) UpdateByID(tx *sql.Tx, eto *entity.MyPlatResource) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatResource], v any) bool {
		return true
	})
}

func (ag *myPlatResourceAutoGen) UpdateNonNilByID(tx *sql.Tx, eto *entity.MyPlatResource) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatResourceAutoGen) UpdateByIDWithFunc(tx *sql.Tx, eto *entity.MyPlatResource, fn func(*iris.Column[entity.MyPlatResource], any) bool) bool {
	if eto.Evaluator() != nil {
		return ag.UpdateByConfig(tx, eto)
	}
	id := *eto.ID
	eto.ID = nil
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		eval.UpdateRef(config.Table(), selfishs...).SetValues(values...).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.UpdateByConfig(tx, config)
}

func (ag *myPlatResourceAutoGen) BatchUpdateByIDWithFunc(tx *sql.Tx, ets []*entity.MyPlatResource, fn func(*iris.Column[entity.MyPlatResource], any) bool) bool {
	for _, eto := range ets {
		ag.UpdateByIDWithFunc(tx, eto, fn)
	}
	return true
}

func (ag *myPlatResourceAutoGen) getDBCtx() context.Context {
	if ag.ctx == nil {
		return context.Background()
	}
	return context.WithValue(context.Background(), constant.TraceIdKey, ag.ctx.GetString(constant.TraceIdKey))
}
