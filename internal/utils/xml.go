package utils

import (
	"encoding/xml"
	"errors"
	"reflect"
)

func UnmarshalXML(b []byte, v interface{}, tag reflect.StructTag, topLevel bool, requiredFields []string) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("v must be a pointer")
	}

	return xml.Unmarshal(b, v)
}