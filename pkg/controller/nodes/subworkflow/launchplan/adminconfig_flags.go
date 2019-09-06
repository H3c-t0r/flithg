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
func (AdminConfig) elemValueOrNil(v interface{}) interface{} {
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

func (AdminConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in AdminConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg AdminConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("AdminConfig", pflag.ExitOnError)
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "tps"), defaultAdminConfig.TPS, "The maximum number of transactions per second to flyte admin from this client.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "burst"), defaultAdminConfig.Burst, "Maximum burst for throttle")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "cacheSize"), defaultAdminConfig.MaxCacheSize, "Maximum cache in terms of number of items stored.")
	return cmdFlags
}
