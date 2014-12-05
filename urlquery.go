package urlquery

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

var StructTag = "url"

func Marshal(ops interface{}) url.Values {
	if ops == nil {
		return url.Values{}
	}
	value := reflect.ValueOf(ops)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return url.Values{}
	}
	items := url.Values(map[string][]string{})
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if field.PkgPath != "" {
			continue
		}
		key := field.Tag.Get(StructTag)
		if key == "" {
			key = strings.ToLower(field.Name)
		}
		v := value.Field(i)
		switch v.Kind() {
		case reflect.Bool:
			if v.Bool() {
				items.Add(key, "true")
			} else {
				items.Add(key, "false")
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v.Int() > 0 {
				items.Add(key, strconv.FormatInt(v.Int(), 10))
			}
		case reflect.Float32, reflect.Float64:
			if v.Float() > 0 {
				items.Add(key, strconv.FormatFloat(v.Float(), 'f', -1, 64))
			}
		case reflect.String:
			if v.String() != "" {
				items.Add(key, v.String())
			}
		}
	}
	return items
}

func Unmarshal(vals url.Values, ops interface{}) error {
	value := reflect.ValueOf(ops)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("urlquery: Passed interface must be a pointer to a struct")
	}
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("urlquery: Passed interface must be a pointer to a struct")
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if field.PkgPath != "" {
			continue
		}
		key := field.Tag.Get(StructTag)
		if key == "" {
			key = field.Name
		}
		urlVal := vals.Get(key)
		if urlVal == "" {
			urlVal = vals.Get(strings.ToLower(key))
		}
		if urlVal == "" {
			continue
		}
		v := value.Field(i)
		switch v.Kind() {
		case reflect.Bool:
			switch urlVal {
			case "true", "1":
				v.SetBool(true)
			case "false", "0":
				v.SetBool(false)
			default:
				return fmt.Errorf("urlquery: Expected boolean, got '%s'", urlVal)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.Atoi(urlVal)
			if err != nil {
				return fmt.Errorf("urlquery: %s", err)
			}
			v.SetInt(int64(intVal))
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(urlVal, 64)
			if err != nil {
				return fmt.Errorf("urlquery: %s", err)
			}
			v.SetFloat(floatVal)
		case reflect.String:
			v.SetString(urlVal)
		}
	}
	return nil
}
