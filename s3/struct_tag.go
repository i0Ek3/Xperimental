package s3

import (
	"reflect"
)

type TagType struct {
	field1 bool   "Whethter true or false."
	field2 string "Things' name."
	field3 int    "Stuffs' price."
}

func mtag(tt TagType) bool {
	i := 0
	if _, err := refTag(tt, i); err == nil {
		return true
	}
	return false
}

func refTag(tt TagType, idx int) (reflect.StructTag, error) {
	ttType := reflect.TypeOf(tt)
	idxField := ttType.Field(idx)
	return idxField.Tag, nil
}
