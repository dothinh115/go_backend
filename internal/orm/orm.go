package orm

import (
	"project/internal/database"
	"reflect"

	"gorm.io/gorm"
)

type ormInterface interface {
	Create(createValue interface{}, mainValue interface{})
}

type ormStruct struct {
	DB *gorm.DB
}

func (os *ormStruct) Create(createValue interface{}, mainValue interface{}) {
	createValueReflect := reflect.ValueOf(createValue).Elem()
	mainValueReflect := reflect.ValueOf(mainValue).Elem()

	for i := 0; i < createValueReflect.NumField(); i++ {
		field := createValueReflect.Type().Field(i).Name
		value := createValueReflect.Field(i).Interface()

		mainField, found := mainValueReflect.Type().FieldByName(field)
		if !found {
			continue
		}
		if mainField.Type.Kind() == reflect.Struct {
			currentStructPtr := reflect.New(mainField.Type).Interface()
			if err := os.DB.Where("id = ?", value).First(&currentStructPtr).Error; err == nil {
				mainValueReflect.FieldByName(field).Set(reflect.ValueOf(currentStructPtr).Elem())
			}
		} else if mainField.Type.Kind() == reflect.Slice {
			elemType := mainField.Type.Elem()

			if elemType.Kind() == reflect.Struct {
				sliceType := reflect.SliceOf(elemType)
				slice := reflect.New(sliceType).Interface()
				if err := os.DB.Where("id IN ?", value).Find(slice).Error; err == nil {
					mainValueReflect.FieldByName(field).Set(reflect.ValueOf(slice).Elem())
				}
			}
		} else {
			mainValueReflect.FieldByName(field).Set(reflect.ValueOf(value))
		}
	}
}

func Service() ormInterface {
	return &ormStruct{
		DB: database.GetDb(),
	}
}
