// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package mysql

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"Back_End/model"
)

func newTeacher(db *gorm.DB, opts ...gen.DOOption) teacher {
	_teacher := teacher{}

	_teacher.teacherDo.UseDB(db, opts...)
	_teacher.teacherDo.UseModel(&model.Teacher{})

	tableName := _teacher.teacherDo.TableName()
	_teacher.ALL = field.NewAsterisk(tableName)
	_teacher.TeacherID = field.NewUint(tableName, "teacher_id")
	_teacher.TeacherName = field.NewString(tableName, "teacher_name")
	_teacher.DepartmentName = field.NewString(tableName, "department_name")
	_teacher.Phone = field.NewString(tableName, "phone")
	_teacher.Password = field.NewString(tableName, "password")

	_teacher.fillFieldMap()

	return _teacher
}

type teacher struct {
	teacherDo

	ALL            field.Asterisk
	TeacherID      field.Uint
	TeacherName    field.String
	DepartmentName field.String
	Phone          field.String
	Password       field.String

	fieldMap map[string]field.Expr
}

func (t teacher) Table(newTableName string) *teacher {
	t.teacherDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teacher) As(alias string) *teacher {
	t.teacherDo.DO = *(t.teacherDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teacher) updateTableName(table string) *teacher {
	t.ALL = field.NewAsterisk(table)
	t.TeacherID = field.NewUint(table, "teacher_id")
	t.TeacherName = field.NewString(table, "teacher_name")
	t.DepartmentName = field.NewString(table, "department_name")
	t.Phone = field.NewString(table, "phone")
	t.Password = field.NewString(table, "password")

	t.fillFieldMap()

	return t
}

func (t *teacher) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teacher) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["teacher_id"] = t.TeacherID
	t.fieldMap["teacher_name"] = t.TeacherName
	t.fieldMap["department_name"] = t.DepartmentName
	t.fieldMap["phone"] = t.Phone
	t.fieldMap["password"] = t.Password
}

func (t teacher) clone(db *gorm.DB) teacher {
	t.teacherDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teacher) replaceDB(db *gorm.DB) teacher {
	t.teacherDo.ReplaceDB(db)
	return t
}

type teacherDo struct{ gen.DO }

type ITeacherDo interface {
	gen.SubQuery
	Debug() ITeacherDo
	WithContext(ctx context.Context) ITeacherDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITeacherDo
	WriteDB() ITeacherDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITeacherDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITeacherDo
	Not(conds ...gen.Condition) ITeacherDo
	Or(conds ...gen.Condition) ITeacherDo
	Select(conds ...field.Expr) ITeacherDo
	Where(conds ...gen.Condition) ITeacherDo
	Order(conds ...field.Expr) ITeacherDo
	Distinct(cols ...field.Expr) ITeacherDo
	Omit(cols ...field.Expr) ITeacherDo
	Join(table schema.Tabler, on ...field.Expr) ITeacherDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITeacherDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITeacherDo
	Group(cols ...field.Expr) ITeacherDo
	Having(conds ...gen.Condition) ITeacherDo
	Limit(limit int) ITeacherDo
	Offset(offset int) ITeacherDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITeacherDo
	Unscoped() ITeacherDo
	Create(values ...*model.Teacher) error
	CreateInBatches(values []*model.Teacher, batchSize int) error
	Save(values ...*model.Teacher) error
	First() (*model.Teacher, error)
	Take() (*model.Teacher, error)
	Last() (*model.Teacher, error)
	Find() ([]*model.Teacher, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Teacher, err error)
	FindInBatches(result *[]*model.Teacher, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Teacher) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITeacherDo
	Assign(attrs ...field.AssignExpr) ITeacherDo
	Joins(fields ...field.RelationField) ITeacherDo
	Preload(fields ...field.RelationField) ITeacherDo
	FirstOrInit() (*model.Teacher, error)
	FirstOrCreate() (*model.Teacher, error)
	FindByPage(offset int, limit int) (result []*model.Teacher, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITeacherDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t teacherDo) Debug() ITeacherDo {
	return t.withDO(t.DO.Debug())
}

func (t teacherDo) WithContext(ctx context.Context) ITeacherDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teacherDo) ReadDB() ITeacherDo {
	return t.Clauses(dbresolver.Read)
}

func (t teacherDo) WriteDB() ITeacherDo {
	return t.Clauses(dbresolver.Write)
}

func (t teacherDo) Session(config *gorm.Session) ITeacherDo {
	return t.withDO(t.DO.Session(config))
}

func (t teacherDo) Clauses(conds ...clause.Expression) ITeacherDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teacherDo) Returning(value interface{}, columns ...string) ITeacherDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teacherDo) Not(conds ...gen.Condition) ITeacherDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teacherDo) Or(conds ...gen.Condition) ITeacherDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teacherDo) Select(conds ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teacherDo) Where(conds ...gen.Condition) ITeacherDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teacherDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITeacherDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teacherDo) Order(conds ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teacherDo) Distinct(cols ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teacherDo) Omit(cols ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teacherDo) Join(table schema.Tabler, on ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teacherDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teacherDo) RightJoin(table schema.Tabler, on ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teacherDo) Group(cols ...field.Expr) ITeacherDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teacherDo) Having(conds ...gen.Condition) ITeacherDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teacherDo) Limit(limit int) ITeacherDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teacherDo) Offset(offset int) ITeacherDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teacherDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITeacherDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teacherDo) Unscoped() ITeacherDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teacherDo) Create(values ...*model.Teacher) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teacherDo) CreateInBatches(values []*model.Teacher, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teacherDo) Save(values ...*model.Teacher) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teacherDo) First() (*model.Teacher, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Teacher), nil
	}
}

func (t teacherDo) Take() (*model.Teacher, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Teacher), nil
	}
}

func (t teacherDo) Last() (*model.Teacher, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Teacher), nil
	}
}

func (t teacherDo) Find() ([]*model.Teacher, error) {
	result, err := t.DO.Find()
	return result.([]*model.Teacher), err
}

func (t teacherDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Teacher, err error) {
	buf := make([]*model.Teacher, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teacherDo) FindInBatches(result *[]*model.Teacher, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teacherDo) Attrs(attrs ...field.AssignExpr) ITeacherDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teacherDo) Assign(attrs ...field.AssignExpr) ITeacherDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teacherDo) Joins(fields ...field.RelationField) ITeacherDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teacherDo) Preload(fields ...field.RelationField) ITeacherDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teacherDo) FirstOrInit() (*model.Teacher, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Teacher), nil
	}
}

func (t teacherDo) FirstOrCreate() (*model.Teacher, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Teacher), nil
	}
}

func (t teacherDo) FindByPage(offset int, limit int) (result []*model.Teacher, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t teacherDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teacherDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teacherDo) Delete(models ...*model.Teacher) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teacherDo) withDO(do gen.Dao) *teacherDo {
	t.DO = *do.(*gen.DO)
	return t
}
