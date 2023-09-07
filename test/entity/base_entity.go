package entity

import (
	"strconv"
	"strings"
)

type InfoService[T any] interface {
	Pkey() *FS[T]

	LogicDeleted() (bool, *FS[T])

	TableName() string

	Alias() string
	AliasQualifier()
	SetAlias(alias string)

	AllFS() []*FS[T]

	CalcInsertFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any)
	CalcInsertFields(fns ...func(string, any) bool) (fields []string)
	CalcInsertValues(fns ...func(string, any) bool) (values []any)
	CalcUpdateFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any)
	CalcUpdateFields(fns ...func(string, any) bool) (fields []string)
	CalcUpdateValues(fns ...func(string, any) bool) (values []any)
}

type ConfigService[T any] interface {
	Configure(configFn func(entity *BaseEntity[T]))
	BEI() *BaseEntity[T]
}

type OperateService[T any] interface {
	Insert(fss ...*FS[T]) *BaseEntity[T]
	InsertFields() (fields []string, placeholders []string)
	InsertSql(times ...int) (string, []any)

	Update(fss ...*FS[T]) *BaseEntity[T]
	UpdateFields() (fields []string)
	UpdateSql() (string, []any)

	WithValues(values ...any)

	WithWhere(pred *Pred)
	Where() (plain string, values []any)

	WithPageable(limit, offset int64)
	Pageable() (bool, int64, int64)

	WithOrder(orders ...*Order)
	Order() (plain []string)

	Select(fss ...*FS[T])
	SelectFS() []*FS[T]
	SelectSql() (pageSql, countSql string, values []any, mappers []func(*T) any)

	Values() []any
}

type BaseEntity[T any] struct {
	inserts      []*FS[T]
	updates      []*FS[T]
	selects      []*FS[T]
	pred         *Pred
	orders       []*Order
	values       []any
	table        string
	buf          strings.Builder
	alias        string
	pageable     bool
	limit        int64
	offset       int64
	qualifier    bool
	logicDeleted bool
	deletedKey   string
	deletedVal   string
}

type FS[T any] struct {
	Name string
	FFn  func(*T) any
}

func OfFS[T any](name string, fn func(*T) any) *FS[T] {
	return &FS[T]{
		Name: name,
		FFn:  fn,
	}
}

func (be *BaseEntity[T]) Insert(fss ...*FS[T]) *BaseEntity[T] {
	be.inserts = fss
	return be
}

func (be *BaseEntity[T]) InsertFields() (fields []string, placeholders []string) {
	for _, insert := range be.inserts {
		fields = append(fields, insert.Name)
		placeholders = append(placeholders, "?")
	}
	return
}

func (be *BaseEntity[T]) InsertSql(times ...int) (sql string, value []any) {
	c := 1
	if len(times) > 0 {
		c = times[0]
	}
	fields, placeholders := be.InsertFields()
	be.buf.WriteString("INSERT INTO ")
	be.buf.WriteString(be.table)
	be.buf.WriteString("(")
	be.buf.WriteString(strings.Join(fields, ", "))
	be.buf.WriteString(") VALUES ")
	ps := make([]string, c)
	for i := 0; i < c; i++ {
		ps[i] = "(" + strings.Join(placeholders, ", ") + ")"
	}
	be.buf.WriteString(strings.Join(ps, ", "))
	be.buf.WriteString(";")

	return be.buf.String(), be.values
}

func (be *BaseEntity[T]) Update(fss ...*FS[T]) *BaseEntity[T] {
	be.updates = fss
	return be
}

func (be *BaseEntity[T]) UpdateFields() (fields []string) {
	for _, update := range be.updates {
		fields = append(fields, update.Name+" = ?")
	}
	return
}

func (be *BaseEntity[T]) UpdateSql() (sql string, value []any) {
	fields := be.UpdateFields()
	plain, values := be.Where()
	be.buf.WriteString("UPDATE ")
	be.buf.WriteString(be.table)
	be.buf.WriteString(" SET ")
	be.buf.WriteString(strings.Join(fields, ", "))
	if len(plain) > 0 {
		be.buf.WriteString(" WHERE ")
		be.buf.WriteString(plain)
	}
	if be.logicDeleted {
		be.buf.WriteString(" AND ")
		be.buf.WriteString(be.deletedKey)
		be.buf.WriteString(" = ")
		be.buf.WriteString(be.deletedVal)
	}
	be.buf.WriteString(";")

	return be.buf.String(), append(be.values, values)
}

func (be *BaseEntity[T]) WithValues(values ...any) {
	be.values = append(be.values, values...)
}

func (be *BaseEntity[T]) WithWhere(pred *Pred) {
	be.pred = pred
}

func (be *BaseEntity[T]) Where() (plain string, values []any) {
	plain, values = be.pred.Render()
	return
}

func (be *BaseEntity[T]) WithPageable(limit, offset int64) {
	be.pageable = true
	be.limit = limit
	be.offset = offset
}

func (be *BaseEntity[T]) Pageable() (bool, int64, int64) {
	return be.pageable, be.limit, be.offset
}

func (be *BaseEntity[T]) WithOrder(orders ...*Order) {
	be.orders = append(be.orders, orders...)
}

func (be *BaseEntity[T]) Order() (plain []string) {
	for _, order := range be.orders {
		k := order.Col
		if order.Asc {
			k += " ASC"
		} else {
			k += " DESC"
		}
		plain = append(plain, k)
	}
	return
}

func (be *BaseEntity[T]) Select(fss ...*FS[T]) {
	be.selects = append(be.selects, fss...)
}

func (be *BaseEntity[T]) SelectFS() []*FS[T] {
	if len(be.selects) == 0 {
		panic("not found select fields")
	}
	return be.selects
}

func (be *BaseEntity[T]) SelectSql() (pageSql, countSql string, values []any, mapping []func(*T) any) {
	fss := be.SelectFS()
	fields := make([]string, len(fss))
	mappers := make([]func(*T) any, len(fss))
	for i, fs := range fss {
		fields[i] = fs.Name
		mappers[i] = fs.FFn
	}
	plain, values := be.Where()
	order := be.Order()
	pageable, limit, offset := be.Pageable()
	be.buf.WriteString("SELECT ")
	be.buf.WriteString(strings.Join(fields, ", "))
	be.buf.WriteString(" FROM ")
	be.buf.WriteString(be.table)
	if len(plain) > 0 {
		be.buf.WriteString(" WHERE ")
		be.buf.WriteString(plain)
	}
	if be.logicDeleted {
		be.buf.WriteString(" AND ")
		be.buf.WriteString(be.deletedKey)
		be.buf.WriteString(" = ")
		be.buf.WriteString(be.deletedVal)
	}
	if len(order) > 0 {
		be.buf.WriteString(" ORDER BY ")
		be.buf.WriteString(strings.Join(order, ", "))
	}
	if pageable {
		countSql = be.buf.String()
		be.buf.WriteString(" LIMIT ")
		be.buf.WriteString(strconv.FormatInt(limit, 10))
		be.buf.WriteString(" OFFSET ")
		be.buf.WriteString(strconv.FormatInt(offset, 10))
	}
	be.buf.WriteString(";")

	return be.buf.String(), countSql, values, mappers
}

func (be *BaseEntity[T]) Values() []any {
	return be.values
}
