package helpers

import (
	"reflect"
)

func MapFromTo[TEntity, TModel comparable](from TEntity, to *TModel) {
	entity := reflect.ValueOf(from)

	for i := 0; i < entity.NumField(); i++ {
		name := entity.Type().Field(i).Name
		value := entity.Field(i)

		result := reflect.ValueOf(to).Elem().FieldByName(name)
		if result.IsValid() && result.CanSet() {
			result.Set(value)
		}
	}
}
