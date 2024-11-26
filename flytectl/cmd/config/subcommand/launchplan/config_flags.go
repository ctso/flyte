// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package launchplan

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
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

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.StringVar(&DefaultConfig.ExecFile, fmt.Sprintf("%v%v", prefix, "execFile"), DefaultConfig.ExecFile, "execution file name to be used for generating execution spec of a single launchplan.")
	cmdFlags.StringVar(&DefaultConfig.Version, fmt.Sprintf("%v%v", prefix, "version"), DefaultConfig.Version, "version of the launchplan to be fetched.")
	cmdFlags.BoolVar(&DefaultConfig.Latest, fmt.Sprintf("%v%v", prefix, "latest"), DefaultConfig.Latest, " flag to indicate to fetch the latest version,  version flag will be ignored in this case")
	cmdFlags.StringVar(&DefaultConfig.Filter.FieldSelector, fmt.Sprintf("%v%v", prefix, "filter.fieldSelector"), DefaultConfig.Filter.FieldSelector, "Specifies the Field selector")
	cmdFlags.StringVar(&DefaultConfig.Filter.SortBy, fmt.Sprintf("%v%v", prefix, "filter.sortBy"), DefaultConfig.Filter.SortBy, "Specifies which field to sort results ")
	cmdFlags.Int32Var(&DefaultConfig.Filter.Limit, fmt.Sprintf("%v%v", prefix, "filter.limit"), DefaultConfig.Filter.Limit, "Specifies the limit")
	cmdFlags.BoolVar(&DefaultConfig.Filter.Asc, fmt.Sprintf("%v%v", prefix, "filter.asc"), DefaultConfig.Filter.Asc, "Specifies the sorting order. By default flytectl sort result in descending order")
	cmdFlags.Int32Var(&DefaultConfig.Filter.Page, fmt.Sprintf("%v%v", prefix, "filter.page"), DefaultConfig.Filter.Page, "Specifies the page number,  in case there are multiple pages of results")
	cmdFlags.StringVar(&DefaultConfig.Workflow, fmt.Sprintf("%v%v", prefix, "workflow"), DefaultConfig.Workflow, "name of the workflow for which the launchplans need to be fetched.")
	return cmdFlags
}
