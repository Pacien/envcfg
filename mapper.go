package envcfg

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func getEnvName(n node) string {
	name := n.properties.Tag.Get(TAG)
	if name == "" {
		name = n.properties.Name
	}

	abs, _ := strconv.ParseBool(n.properties.Tag.Get(ABS_TAG))
	if !abs && n.parent != nil {
		if n.parent.properties != nil {
			parentName := getEnvName(*n.parent)
			name = parentName + SEP + name
		}
	}

	return strings.ToUpper(name)
}

func setValue(n node, v string) (node, error) {
	switch n.value.Kind() {
	case reflect.String:
		n.value.SetString(v)

	case reflect.Bool:
		boolVal, err := strconv.ParseBool(v)
		if err != nil {
			return n, err
		}
		n.value.SetBool(boolVal)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(v, 0, n.value.Type().Bits())
		if err != nil {
			return n, err
		}
		n.value.SetInt(intVal)

	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(v, n.value.Type().Bits())
		if err != nil {
			return n, err
		}
		n.value.SetFloat(floatVal)
	}

	return n, nil
}

func setFieldValue(n node) (node, error) {
	if n.value.Kind() == reflect.Struct {
		setStructFields(n)
		return n, nil
	}

	v := os.Getenv(getEnvName(n))
	if v != "" {
		return setValue(n, v)
	}

	return n, nil
}

func setStructFields(n node) (node, []error) {
	t := n.value.Type()
	errs := []error{}

	for i := 0; i < n.value.NumField(); i++ {
		v := n.value.Field(i)
		p := t.Field(i)

		_, err := setFieldValue(node{&n, &v, &p})
		if err != nil {
			errs = append(errs, err)
		}
	}

	return n, errs
}
