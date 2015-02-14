// Package envcfg provides environment variable mapping to structs.
//
// Can be used to read configuration parameters from the environment.
//
// Fields for which environment variables can be found are overwritten, otherwise they are left to their previous
// value.
//
// Can be used, for example, after gcfg to override settings provided in a configuration file.
package envcfg

import (
	"errors"
	"reflect"
)

const (
	TAG     = "env"
	ABS_TAG = "absenv"
	SEP     = "_"
)

type node struct {
	parent     *node
	value      *reflect.Value
	properties *reflect.StructField
}

var ErrInvalidConfigStruct = errors.New("invalid parameter: must map to a struct")

func ReadInto(cfgStruct interface{}) (interface{}, []error) {
	s := reflect.ValueOf(cfgStruct).Elem()

	if s.Kind() != reflect.Struct {
		return nil, []error{ErrInvalidConfigStruct}
	}

	_, errs := setStructFields(node{nil, &s, nil})

	return cfgStruct, errs
}
