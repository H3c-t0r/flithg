// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package catalog

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
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "type"), defaultConfig.Type, " Catalog Implementation to use")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "endpoint"), defaultConfig.Endpoint, " Endpoint for catalog service")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "cache-endpoint"), defaultConfig.CacheEndpoint, " Endpoint for cache service")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "insecure"), defaultConfig.Insecure, " Use insecure grpc connection")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "max-cache-age"), defaultConfig.MaxCacheAge.String(), " Cache entries past this age will incur cache miss. 0 means cache never expires")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "use-admin-auth"), defaultConfig.UseAdminAuth, " Use the same gRPC credentials option as the flyteadmin client")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "max-retries"), defaultConfig.MaxRetries, "Max number of gRPC retries")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "max-per-retry-timeout"), defaultConfig.MaxPerRetryTimeout.String(), "gRPC per retry timeout. O means no timeout.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "backoff-scalar"), defaultConfig.BackOffScalar, "gRPC backoff scalar in milliseconds")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "inline-cache"), defaultConfig.InlineCache, " Attempt to use in-line cache")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "default-service-config"), defaultConfig.DefaultServiceConfig, " Set the default service config for the catalog gRPC client")
	return cmdFlags
}
