// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package configs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsCacheServiceConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementCacheServiceConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsCacheServiceConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookCacheServiceConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementCacheServiceConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_CacheServiceConfig(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookCacheServiceConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_CacheServiceConfig(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_CacheServiceConfig(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_CacheServiceConfig(val, result))
}

func testDecodeRaw_CacheServiceConfig(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_CacheServiceConfig(vStringSlice, result))
}

func TestCacheServiceConfig_GetPFlagSet(t *testing.T) {
	val := CacheServiceConfig{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestCacheServiceConfig_SetFlags(t *testing.T) {
	actual := CacheServiceConfig{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_storage-prefix", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("storage-prefix", testValue)
			if vString, err := cmdFlags.GetString("storage-prefix"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.StoragePrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metrics-scope", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metrics-scope", testValue)
			if vString, err := cmdFlags.GetString("metrics-scope"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.MetricsScope)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_profiler-port", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("profiler-port", testValue)
			if vInt, err := cmdFlags.GetInt("profiler-port"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vInt), &actual.ProfilerPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_heartbeat-grace-period-multiplier", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("heartbeat-grace-period-multiplier", testValue)
			if vInt, err := cmdFlags.GetInt("heartbeat-grace-period-multiplier"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vInt), &actual.HeartbeatGracePeriodMultiplier)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_max-reservation-heartbeat", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.MaxReservationHeartbeat.String()

			cmdFlags.Set("max-reservation-heartbeat", testValue)
			if vString, err := cmdFlags.GetString("max-reservation-heartbeat"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.MaxReservationHeartbeat)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_data-store-type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("data-store-type", testValue)
			if vString, err := cmdFlags.GetString("data-store-type"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.DataStoreType)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_reservation-data-store-type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("reservation-data-store-type", testValue)
			if vString, err := cmdFlags.GetString("reservation-data-store-type"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.ReservationDataStoreType)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxInlineSizeBytes", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("maxInlineSizeBytes", testValue)
			if vInt64, err := cmdFlags.GetInt64("maxInlineSizeBytes"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vInt64), &actual.MaxInlineSizeBytes)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_aws-region", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("aws-region", testValue)
			if vString, err := cmdFlags.GetString("aws-region"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.AwsRegion)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redis-address", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redis-address", testValue)
			if vString, err := cmdFlags.GetString("redis-address"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.RedisAddress)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redis-username", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redis-username", testValue)
			if vString, err := cmdFlags.GetString("redis-username"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.RedisUsername)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_redis-password", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("redis-password", testValue)
			if vString, err := cmdFlags.GetString("redis-password"); err == nil {
				testDecodeJson_CacheServiceConfig(t, fmt.Sprintf("%v", vString), &actual.RedisPassword)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
