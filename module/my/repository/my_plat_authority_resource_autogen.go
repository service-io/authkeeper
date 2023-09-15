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

// iMyPlatAuthorityResourceAutoGen 该接口自动生成, 请勿修改
type iMyPlatAuthorityResourceAutoGen interface {
	// SelectOneByConfig Use config service to execute(select statement)
	SelectOneByConfig(iris.ConfigService[entity.MyPlatAuthorityResource]) *entity.MyPlatAuthorityResource

	// SelectManyByConfig Use config service to execute(select statement)
	SelectManyByConfig(iris.ConfigService[entity.MyPlatAuthorityResource]) []*entity.MyPlatAuthorityResource

	// SelectPageByConfig Use config service to execute(select statement)
	SelectPageByConfig(iris.ConfigService[entity.MyPlatAuthorityResource]) ([]*entity.MyPlatAuthorityResource, int64)

	// InsertByConfig Use config service to execute(insert statement)
	InsertByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthorityResource]) bool

	// UpdateByConfig Use config service to execute(update statement)
	UpdateByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthorityResource]) bool

	// DeleteByConfig Use config service to execute(delete statement)
	DeleteByConfig(*sql.Tx, iris.ConfigService[entity.MyPlatAuthorityResource]) bool

	// SelectByID Query and return the related entries of id(p key)
	SelectByID(int64) *entity.MyPlatAuthorityResource

	// SelectByIDs Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one, otherwise call SelectByID
	SelectByIDs(...int64) []*entity.MyPlatAuthorityResource

	// BatchSelectByID Query and return the related entries of id(p key list).
	// Call BatchSelectByID if the total number of id is more than one
	BatchSelectByID([]int64) []*entity.MyPlatAuthorityResource

	// SelectByAuthorityID Query by Authority ID
	SelectByAuthorityID(int64) []*entity.MyPlatAuthorityResource

	// SelectByResourceID Query by Resource ID
	SelectByResourceID(int64) []*entity.MyPlatAuthorityResource

	// SelectByPlatID Query by Plat ID
	SelectByPlatID(int64) []*entity.MyPlatAuthorityResource

	// SelectByTenantID Query by Tenant ID
	SelectByTenantID(int64) []*entity.MyPlatAuthorityResource

	// SelectAuthorityByResourceID Query Authority by Resource ID
	SelectAuthorityByResourceID(int64) []*entity.MyPlatAuthority

	// SelectResourceByAuthorityID Query Resource by Authority ID
	SelectResourceByAuthorityID(int64) []*entity.MyPlatResource

	// SelectAllWithPage Query all entries by page
	SelectAllWithPage(int64, int64) ([]*entity.MyPlatAuthorityResource, int64)

	// Insert Matches all columns and insert into the table
	Insert(*sql.Tx, *entity.MyPlatAuthorityResource) int64

	// InsertNonNil Matches non-nil columns and insert into the table
	InsertNonNil(*sql.Tx, *entity.MyPlatAuthorityResource) int64

	// InsertWithFunc Matches all columns matching this function and insert into the table
	InsertWithFunc(*sql.Tx, *entity.MyPlatAuthorityResource, func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) int64

	// BatchInsert Batch insert all columns of table into table
	BatchInsert(*sql.Tx, []*entity.MyPlatAuthorityResource) []int64

	// BatchInsertNonNil Batch insert non-nil columns(matches first element of list) of table into table
	BatchInsertNonNil(*sql.Tx, []*entity.MyPlatAuthorityResource) []int64

	// BatchInsertWithFunc Batch insert all columns matching this function(matches first element of list) of table into table
	BatchInsertWithFunc(*sql.Tx, []*entity.MyPlatAuthorityResource, func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) []int64

	// DeleteByID Delete the related entries of id(p key)
	DeleteByID(*sql.Tx, int64) bool

	// DeleteByIDs Delete the related entries of id(p key list)
	DeleteByIDs(*sql.Tx, ...int64) bool

	// BatchDeleteByID Batch delete the related entries of id(p key list)
	BatchDeleteByID(*sql.Tx, []int64) bool

	// UpdateByID Update all columns of the related entries of id(p key)
	UpdateByID(*sql.Tx, *entity.MyPlatAuthorityResource) bool

	// UpdateNonNilByID Update non-nil columns of the related entries of id(p key)
	UpdateNonNilByID(*sql.Tx, *entity.MyPlatAuthorityResource) bool

	// UpdateByIDWithFunc Update all columns matching this function of the related entries of id(p key)
	UpdateByIDWithFunc(*sql.Tx, *entity.MyPlatAuthorityResource, func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) bool

	// BatchUpdateByIDWithFunc Batch update all columns matching this function of the related entries of id(p key)
	BatchUpdateByIDWithFunc(*sql.Tx, []*entity.MyPlatAuthorityResource, func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) bool
}

// myPlatAuthorityResourceAutoGen 该结构体自动生成, 请勿修改
type myPlatAuthorityResourceAutoGen struct {
	ctx *gin.Context
}

func (ag *myPlatAuthorityResourceAutoGen) SelectOneByConfig(config iris.ConfigService[entity.MyPlatAuthorityResource]) *entity.MyPlatAuthorityResource {
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
	eto := helper.Row(row, func() (*entity.MyPlatAuthorityResource, []any) {
		eto := entity.NewMyPlatAuthorityResource()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return eto
}

func (ag *myPlatAuthorityResourceAutoGen) SelectManyByConfig(config iris.ConfigService[entity.MyPlatAuthorityResource]) []*entity.MyPlatAuthorityResource {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatAuthorityResource, []any) {
		eto := entity.NewMyPlatAuthorityResource()
		mappers := evalInfo.MapperRows(eto)
		return eto, mappers
	})
	return ets
}

func (ag *myPlatAuthorityResourceAutoGen) SelectPageByConfig(config iris.ConfigService[entity.MyPlatAuthorityResource]) ([]*entity.MyPlatAuthorityResource, int64) {
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
	ets := helper.Rows(rows, func() (*entity.MyPlatAuthorityResource, []any) {
		eto := entity.NewMyPlatAuthorityResource()
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

func (ag *myPlatAuthorityResourceAutoGen) InsertByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthorityResource]) bool {
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

func (ag *myPlatAuthorityResourceAutoGen) UpdateByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthorityResource]) bool {
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

func (ag *myPlatAuthorityResourceAutoGen) DeleteByConfig(tx *sql.Tx, config iris.ConfigService[entity.MyPlatAuthorityResource]) bool {
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

func (ag *myPlatAuthorityResourceAutoGen) SelectByID(id int64) *entity.MyPlatAuthorityResource {
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.SelectOneByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectByIDs(ids ...int64) []*entity.MyPlatAuthorityResource {
	return ag.BatchSelectByID(ids)
}

func (ag *myPlatAuthorityResourceAutoGen) BatchSelectByID(ids []int64) []*entity.MyPlatAuthorityResource {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectByAuthorityID(id int64) []*entity.MyPlatAuthorityResource {
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.AuthorityIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectByResourceID(id int64) []*entity.MyPlatAuthorityResource {
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.ResourceIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectByPlatID(id int64) []*entity.MyPlatAuthorityResource {
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.PlatIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectByTenantID(id int64) []*entity.MyPlatAuthorityResource {
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Where(config.TenantIdCol().EQ(id)).Eval()
	})
	return ag.SelectManyByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) SelectAuthorityByResourceID(resourceID int64) []*entity.MyPlatAuthority {
	recorder := recorderx.FetchRecorder(ag.ctx)
	rightConfig := entity.NewMyPlatResource()
	shipConfig := entity.NewMyPlatAuthorityResource()
	leftConfig := entity.NewMyPlatAuthority()
	leftConfig.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthority]) {
		leftTable := leftConfig.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(leftConfig.IDCol().Decorate(leftTable.Decorate), shipConfig.AuthorityIdCol().Decorate(shipTable.Decorate))
		rightTable := rightConfig.Table()
		rightTable.LeftJoin().OnEQ(rightConfig.IDCol().Decorate(rightTable.Decorate), shipConfig.ResourceIdCol().Decorate(shipTable.Decorate))
		eval.Select(leftConfig.Asterisk(leftTable.Decorate)...).From(leftTable.Ref(shipTable, rightTable)).Where(shipConfig.ResourceIdCol().Decorate(shipTable.Decorate).EQ(resourceID)).Eval()
	})
	evalInfo := leftConfig.Evaluator().EvalInfo()
	if evalInfo == nil {
		return nil
	}
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()
	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	recorder.MaybePanic(err)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	rows, err := stmt.QueryContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	myPlatAuthoritys := helper.Rows(rows, func() (*entity.MyPlatAuthority, []any) {
		myPlatAuthority := entity.NewMyPlatAuthority()
		mappers := evalInfo.MapperRows(myPlatAuthority)
		return myPlatAuthority, mappers
	})
	return myPlatAuthoritys
}

func (ag *myPlatAuthorityResourceAutoGen) SelectResourceByAuthorityID(authorityID int64) []*entity.MyPlatResource {
	recorder := recorderx.FetchRecorder(ag.ctx)
	rightConfig := entity.NewMyPlatAuthority()
	shipConfig := entity.NewMyPlatAuthorityResource()
	leftConfig := entity.NewMyPlatResource()
	leftConfig.Configure(func(eval *iris.Evaluator[entity.MyPlatResource]) {
		leftTable := leftConfig.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(leftConfig.IDCol().Decorate(leftTable.Decorate), shipConfig.ResourceIdCol().Decorate(shipTable.Decorate))
		rightTable := rightConfig.Table()
		rightTable.LeftJoin().OnEQ(rightConfig.IDCol().Decorate(rightTable.Decorate), shipConfig.AuthorityIdCol().Decorate(shipTable.Decorate))
		eval.Select(leftConfig.Asterisk(leftTable.Decorate)...).From(leftTable.Ref(shipTable, rightTable)).Where(shipConfig.AuthorityIdCol().Decorate(shipTable.Decorate).EQ(authorityID)).Eval()
	})
	evalInfo := leftConfig.Evaluator().EvalInfo()
	if evalInfo == nil {
		return nil
	}
	execSQL := evalInfo.SQL()
	values := evalInfo.Values()
	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	recorder.MaybePanic(err)
	defer helper.DeferClose(stmt, recorder.MaybePanic)
	rows, err := stmt.QueryContext(ag.getDBCtx(), values...)
	recorder.MaybePanic(err)
	myPlatResources := helper.Rows(rows, func() (*entity.MyPlatResource, []any) {
		myPlatResource := entity.NewMyPlatResource()
		mappers := evalInfo.MapperRows(myPlatResource)
		return myPlatResource, mappers
	})
	return myPlatResources
}

func (ag *myPlatAuthorityResourceAutoGen) SelectAllWithPage(page int64, size int64) ([]*entity.MyPlatAuthorityResource, int64) {
	if page < 1 {
		return nil, 0
	}
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(size, (page-1)*size).Eval()
	})
	return ag.SelectPageByConfig(config)
}

func (ag *myPlatAuthorityResourceAutoGen) Insert(tx *sql.Tx, eto *entity.MyPlatAuthorityResource) int64 {
	return ag.InsertWithFunc(tx, eto, func(*iris.Column[entity.MyPlatAuthorityResource], any) bool {
		return true
	})
}

func (ag *myPlatAuthorityResourceAutoGen) InsertNonNil(tx *sql.Tx, eto *entity.MyPlatAuthorityResource) int64 {
	return ag.InsertWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthorityResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityResourceAutoGen) InsertWithFunc(tx *sql.Tx, eto *entity.MyPlatAuthorityResource, fn func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) int64 {
	if eto.Evaluator() != nil {
		if !ag.InsertByConfig(tx, eto) {
			return 0
		}
		return *eto.ID
	}
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return 0
	}
	return *config.ID
}

func (ag *myPlatAuthorityResourceAutoGen) BatchInsert(tx *sql.Tx, ets []*entity.MyPlatAuthorityResource) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatAuthorityResource], v any) bool {
		return true
	})
}

func (ag *myPlatAuthorityResourceAutoGen) BatchInsertNonNil(tx *sql.Tx, ets []*entity.MyPlatAuthorityResource) []int64 {
	return ag.BatchInsertWithFunc(tx, ets, func(c *iris.Column[entity.MyPlatAuthorityResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityResourceAutoGen) BatchInsertWithFunc(tx *sql.Tx, ets []*entity.MyPlatAuthorityResource, fn func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) []int64 {
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
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Insert(selfishs...).Into(config.Table()).Values(values...).Eval()
	})
	if !ag.InsertByConfig(tx, config) {
		return nil
	}
	return ids
}

func (ag *myPlatAuthorityResourceAutoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.DeleteByIDs(tx, id)
}

func (ag *myPlatAuthorityResourceAutoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

func (ag *myPlatAuthorityResourceAutoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	values := make([]any, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	config := entity.NewMyPlatAuthorityResource()
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.Delete().From(config.Table()).Where(config.IDCol().IN(values...)).Eval()
	})
	return ag.DeleteByConfig(tx, config)
}

func (ag *myPlatAuthorityResourceAutoGen) UpdateByID(tx *sql.Tx, eto *entity.MyPlatAuthorityResource) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthorityResource], v any) bool {
		return true
	})
}

func (ag *myPlatAuthorityResourceAutoGen) UpdateNonNilByID(tx *sql.Tx, eto *entity.MyPlatAuthorityResource) bool {
	return ag.UpdateByIDWithFunc(tx, eto, func(c *iris.Column[entity.MyPlatAuthorityResource], v any) bool {
		return helper.IsNonNil(v)
	})
}

func (ag *myPlatAuthorityResourceAutoGen) UpdateByIDWithFunc(tx *sql.Tx, eto *entity.MyPlatAuthorityResource, fn func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) bool {
	if eto.Evaluator() != nil {
		return ag.UpdateByConfig(tx, eto)
	}
	id := *eto.ID
	eto.ID = nil
	config := eto
	selfishs, values := eto.ColumnAndValue(fn)
	config.Configure(func(eval *iris.Evaluator[entity.MyPlatAuthorityResource]) {
		eval.UpdateRef(config.Table(), selfishs...).SetValues(values...).Where(config.IDCol().EQ(id)).Eval()
	})
	return ag.UpdateByConfig(tx, config)
}

func (ag *myPlatAuthorityResourceAutoGen) BatchUpdateByIDWithFunc(tx *sql.Tx, ets []*entity.MyPlatAuthorityResource, fn func(*iris.Column[entity.MyPlatAuthorityResource], any) bool) bool {
	for _, eto := range ets {
		ag.UpdateByIDWithFunc(tx, eto, fn)
	}
	return true
}

func (ag *myPlatAuthorityResourceAutoGen) getDBCtx() context.Context {
	if ag.ctx == nil {
		return context.Background()
	}
	return context.WithValue(context.Background(), constant.TraceIdKey, ag.ctx.GetString(constant.TraceIdKey))
}
