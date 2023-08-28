package entity

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

type User struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
	Age  *uint8  `json:"age"`

	BaseEntity[User]
}

func (e *User) Pkey() *FS[User] {
	return e.FID()
}

func (e *User) FID() *FS[User] {
	return OfFS(
		"id", func(r *User) any {
			return &r.ID
		},
	)
}

func (e *User) FName() *FS[User] {
	return OfFS(
		"name", func(r *User) any {
			return &r.Name
		},
	)
}

func (e *User) FAge() *FS[User] {
	return OfFS(
		"age", func(r *User) any {
			return &r.Age
		},
	)
}

func (e *User) TableName() string {
	return "user"
}

func (e *User) Alias() string {
	if e.BaseEntity.qualifier {
		return "user"
	}
	return e.BaseEntity.alias
}

func (e *User) AliasQualifier() string {
	if len(e.Alias()) == 0 {
		return ""
	}
	return e.Alias() + "."
}

func (e *User) SetAlias(alias string) {
	e.BaseEntity.alias = alias
}

func (e *User) AllFS() []*FS[User] {
	fss := []*FS[User]{e.Pkey(), e.FName(), e.FAge()}
	return fss
}

func (e *User) SelectFS() []*FS[User] {
	if len(e.BaseEntity.selects) == 0 {
		panic("not found select fields")
	}
	return e.BaseEntity.selects
}

func (e *User) Insert(fss ...*FS[User]) *User {
	e.BaseEntity.inserts = fss
	return e
}

func (e *User) InsertFields() (fields []string, placeholders []string) {
	for _, insert := range e.BaseEntity.inserts {
		fields = append(fields, insert.Name)
		placeholders = append(placeholders, "?")
	}
	return
}

func (e *User) Update(fss ...*FS[User]) *User {
	e.BaseEntity.updates = fss
	return e
}

func (e *User) UpdateFields() (fields []string) {
	for _, update := range e.BaseEntity.updates {
		fields = append(fields, update.Name+" = ?")
	}
	return
}

func (e *User) WithValues(values ...any) {
	e.BaseEntity.values = values
}

func (e *User) Select(fss ...*FS[User]) {
	e.BaseEntity.selects = fss
}

func (e *User) SelectFields() (fields []string) {
	for _, sel := range e.BaseEntity.selects {
		fields = append(fields, sel.Name)
	}
	return
}

func (e *User) ReturnFields() (fields []func(*User) any) {
	for _, sel := range e.BaseEntity.selects {
		fields = append(fields, sel.FFn)
	}
	return
}

func (e *User) WithWhere(pred *Pred) {
	e.BaseEntity.pred = pred
}

func (e *User) Where() (plain string, values []any) {
	plain, values = e.BaseEntity.pred.Render()
	return
}

func (e *User) WithOrder(orders ...*Order) {
	e.BaseEntity.orders = orders
}

func (e *User) Order() (plain []string) {
	for _, order := range e.BaseEntity.orders {
		k := order.Col
		if order.Asc {
			k += " asc"
		} else {
			k += " desc"
		}
		plain = append(plain, k)
	}
	return
}

func (e *User) CalcInsertFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any) {
	var fn = func(k string, f any) bool {
		return f != nil
	}

	if len(fns) > 0 {
		fn = fns[0]
	}

	if fn(e.FID().Name, e.ID) {
		fields = append(fields, e.FID().Name)
		values = append(values, e.ID)
	}

	if fn(e.FName().Name, e.Name) {
		fields = append(fields, e.FName().Name)
		values = append(values, e.Name)
	}

	if fn(e.FAge().Name, e.Age) {
		fields = append(fields, e.FAge().Name)
		values = append(values, e.Age)
	}

	return
}

func (e *User) CalcInsertFields(fns ...func(string, any) bool) (fields []string) {
	fields, _ = e.CalcInsertFieldsAndValues(fns...)
	return
}

func (e *User) CalcInsertValues(fns ...func(string, any) bool) (values []any) {
	_, values = e.CalcInsertFieldsAndValues(fns...)
	return
}

func (e *User) CalcUpdateFieldsAndValues(fns ...func(string, any) bool) (fields []string, values []any) {
	var fn = func(k string, f any) bool {
		if k == e.Pkey().Name {
			return false
		}
		return f != nil
	}

	if len(fns) > 0 {
		fn = fns[0]
	}

	if fn(e.FID().Name, e.ID) {
		fields = append(fields, e.FID().Name+" = ?")
		values = append(values, e.ID)
	}

	if fn(e.FName().Name, e.Name) {
		fields = append(fields, e.FName().Name+" = ?")
		values = append(values, e.Name)
	}

	if fn(e.FAge().Name, e.Age) {
		fields = append(fields, e.FAge().Name+" = ?")
		values = append(values, e.Age)
	}

	return fields, values
}

func (e *User) CalcUpdateFields(fns ...func(string, any) bool) (fields []string) {
	fields, _ = e.CalcUpdateFieldsAndValues(fns...)
	return
}

func (e *User) CalcUpdateValues(fns ...func(string, any) bool) (values []any) {
	_, values = e.CalcUpdateFieldsAndValues(fns...)
	return
}
