// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package workflowexecutionconfig

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (AttrFetchConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (AttrFetchConfig) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (AttrFetchConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in AttrFetchConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg AttrFetchConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("AttrFetchConfig", pflag.ExitOnError)
	cmdFlags.StringVar(&DefaultFetchConfig.AttrFile, fmt.Sprintf("%v%v", prefix, "attrFile"), DefaultFetchConfig.AttrFile, "attribute file name to be used for generating attribute for the resource type.")
	cmdFlags.BoolVar(&DefaultFetchConfig.Gen, fmt.Sprintf("%v%v", prefix, "gen"), DefaultFetchConfig.Gen, "generates an empty workflow execution config file with conformance to the api format.")
	return cmdFlags
}
