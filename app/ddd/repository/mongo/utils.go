package mongo

import (
	"fmt"
	"net/url"
	"reflect"
)

func encodeValue(val string) string {
	if val == "" {
		return val
	}
	v := url.Values{}
	v.Set("", val)
	return v.Encode()[1:]
}

// merges two structs, where b's values take precendence over a's values (b's values will be kept over a's if each field has a value)
func merge(result, merge interface{}) error {
	aValOf := reflect.ValueOf(result)

	if aValOf.Kind() == reflect.Ptr {
		aValOf = aValOf.Elem()
	}

	if !aValOf.CanSet() {
		return fmt.Errorf("参数不可写")
	}

	bValOf := reflect.ValueOf(merge)

	if aValOf.Type() != bValOf.Type() {
		return fmt.Errorf("参数类型不同")
	}

	for i := 0; i < aValOf.NumField(); i++ {
		if bValOf.Field(i).IsZero() {
			continue
		}
		aValOf.Field(i).Set(bValOf.Field(i))
	}

	return nil
}
