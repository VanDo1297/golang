package utilsgorm

import (
	"reflect"

	"github.com/go-errors/errors"
	"gorm.io/gorm"
)

func FindObject(returnType interface{}, db *gorm.DB, where string, args ...interface{}) (interface{}, error) {
	t := reflect.TypeOf(returnType)
	v := reflect.New(t)
	dbResult := db.Where(where, args...).First(v.Interface())
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(dbResult.Error, 0)
	}
	return v.Interface(), nil
}

func ScanObject(returnType interface{}, db *gorm.DB, query string, args ...interface{}) (interface{}, error) {
	t := reflect.TypeOf(returnType)
	v := reflect.New(t)
	dbResult := db.Raw(query, args...).Scan(v.Interface())
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(dbResult.Error, 0)
	}
	return v.Interface(), nil
}

func FindObjects(returnType interface{}, db *gorm.DB, where string, args ...interface{}) (interface{}, error) {
	t := reflect.TypeOf(returnType)
	v := reflect.New(t)
	elemType := reflect.TypeOf(v.Interface())
	elemSlice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 10)

	val := reflect.ValueOf(elemSlice.Interface())
	vp := reflect.New(val.Type())
	vp.Elem().Set(val)

	dbResult := db.Where(where, args...).Find(vp.Interface())
	if dbResult.Error != nil {
		return nil, errors.Wrap(dbResult.Error, 0)
	}
	return vp.Elem().Interface(), nil
}

func ScanObjects(returnType interface{}, db *gorm.DB, query string, args ...interface{}) (interface{}, error) {
	t := reflect.TypeOf(returnType)
	v := reflect.New(t)
	elemType := reflect.TypeOf(v.Interface())
	elemSlice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 10)

	val := reflect.ValueOf(elemSlice.Interface())
	vp := reflect.New(val.Type())
	vp.Elem().Set(val)

	dbResult := db.Raw(query, args...).Scan(vp.Interface())
	if dbResult.Error != nil {
		return nil, errors.Wrap(dbResult.Error, 0)
	}
	return vp.Elem().Interface(), nil
}
