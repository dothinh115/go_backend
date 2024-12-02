package orm

import (
	"project/internal/database"
	"reflect"

	"gorm.io/gorm"
)

type ormInterface interface {
	Create(createValue interface{}, mainValue interface{}) error
	Update(updatedValue interface{}, mainValue interface{}) error
}

type ormStruct struct {
	DB *gorm.DB
}

func (os *ormStruct) Create(createValue interface{}, mainValue interface{}) error {
	createValueReflect := reflect.ValueOf(createValue).Elem()
	mainValueReflect := reflect.ValueOf(mainValue).Elem()
	preloadFields := make([]string, 0)

	for i := 0; i < createValueReflect.NumField(); i++ {
		field := createValueReflect.Type().Field(i).Name
		value := createValueReflect.Field(i)

		mainField, found := mainValueReflect.Type().FieldByName(field)
		if !found {
			continue
		}
		if mainField.Type.Kind() == reflect.Ptr && mainField.Type.Elem().Kind() == reflect.Struct {
			currentStructPtr := reflect.New(mainField.Type.Elem())
			currentStructPtr.Elem().FieldByName("ID").Set(reflect.ValueOf(value.Interface()))
			mainValueReflect.FieldByName(field).Set(currentStructPtr)
			preloadFields = append(preloadFields, field)
		} else if mainField.Type.Kind() == reflect.Ptr && mainField.Type.Elem().Kind() == reflect.Slice {
			elemType := mainField.Type.Elem()
			if elemType.Elem().Kind() == reflect.Struct {
				slice := reflect.MakeSlice(elemType, 0, value.Len())
				for j := 0; j < value.Len(); j++ {
					currentStructPtr := reflect.New(elemType.Elem())
					currentStructPtr.Elem().FieldByName("ID").Set(reflect.ValueOf(value.Index(j).Interface()))
					slice = reflect.Append(slice, currentStructPtr.Elem())
				}
				slicePtr := reflect.New(slice.Type())
				slicePtr.Elem().Set(slice)
				mainValueReflect.FieldByName(field).Set(slicePtr)
				preloadFields = append(preloadFields, field)
			} else {
				mainValueReflect.FieldByName(field).Set(reflect.ValueOf(value.Interface()))
			}
		} else {
			mainValueReflect.FieldByName(field).Set(reflect.ValueOf(value.Interface()))
		}
	}
	if err := os.DB.Create(mainValue).Error; err != nil {
		return err
	}
	preload := os.DB
	for _, field := range preloadFields {
		preload = preload.Preload(field)
	}
	if err := preload.Where("id = ?", mainValueReflect.FieldByName("ID").Interface()).First(mainValue).Error; err != nil {
		return err
	}

	return nil
}

func (os *ormStruct) Update(updatedValue interface{}, mainValue interface{}) error {
	updatedValueReflect := reflect.ValueOf(updatedValue).Elem()
	mainValueReflect := reflect.ValueOf(mainValue).Elem()
	preloadFields := make([]string, 0)
	relations := make(map[string][]interface{})

	for i := 0; i < updatedValueReflect.NumField(); i++ {
		field := updatedValueReflect.Type().Field(i).Name
		value := updatedValueReflect.Field(i)

		if value.IsZero() {
			continue
		}

		mainField, found := mainValueReflect.Type().FieldByName(field)
		if !found {
			continue
		}

		if mainField.Type.Kind() == reflect.Ptr && mainField.Type.Elem().Kind() == reflect.Struct {
			currentStructPtr := reflect.New(mainField.Type.Elem())
			currentStructPtr.Elem().FieldByName("ID").Set(reflect.ValueOf(value.Elem().Interface()))
			mainValueReflect.FieldByName(field).Set(currentStructPtr)
			preloadFields = append(preloadFields, field)
		} else if mainField.Type.Kind() == reflect.Ptr && mainField.Type.Elem().Kind() == reflect.Slice {
			elemType := mainField.Type.Elem() // []models.Category
			if elemType.Elem().Kind() == reflect.Struct {
				if _, found := relations[field]; !found {
					relations[field] = make([]interface{}, 0)
				}
				for j := 0; j < value.Elem().Len(); j++ {
					currentStructPtr := reflect.New(elemType.Elem())
					currentStructPtr.Elem().FieldByName("ID").Set(reflect.ValueOf(value.Elem().Index(j).Interface()))
					relations[field] = append(relations[field], currentStructPtr.Elem().Interface())
				}
				preloadFields = append(preloadFields, field)
			} else {
				mainValueReflect.FieldByName(field).Set(reflect.ValueOf(value.Elem().Interface()))
			}
		} else {
			mainValueReflect.FieldByName(field).Set(reflect.ValueOf(value.Elem().Interface()))
		}

	}
	db := os.DB
	if len(relations) > 0 {
		for relation, entities := range relations {
			copyDb := db.Preload(relation)

			fieldReflect := mainValueReflect.FieldByName(relation)
			fieldType := fieldReflect.Type()
			if fieldType.Kind() == reflect.Ptr == fieldReflect.IsNil() {
				fieldReflect.Set(reflect.New(fieldType.Elem()))
			}
			fieldSlice := reflect.MakeSlice(fieldType.Elem(), 0, len(entities))
			for _, entity := range entities {
				fieldSlice = reflect.Append(fieldSlice, reflect.ValueOf(entity))
			}
			fieldSlicePtr := reflect.New(fieldSlice.Type())
			fieldSlicePtr.Elem().Set(fieldSlice)
			if err := copyDb.Model(mainValue).Association(relation).Replace(fieldSlicePtr.Interface()); err != nil {
				return err
			}
		}
	}

	if err := db.Updates(mainValue).Error; err != nil {
		return err
	}
	for _, field := range preloadFields {
		db = db.Preload(field)
	}
	if err := db.Where("id = ?", mainValueReflect.FieldByName("ID").Interface()).First(mainValue).Error; err != nil {
		return err
	}

	return nil
}

func Service() ormInterface {
	return &ormStruct{
		DB: database.GetDb(),
	}
}
