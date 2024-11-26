// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsK8sPluginConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementK8sPluginConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsK8sPluginConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookK8sPluginConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementK8sPluginConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

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

func decode_K8sPluginConfig(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookK8sPluginConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_K8sPluginConfig(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_K8sPluginConfig(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_K8sPluginConfig(val, result))
}

func testDecodeRaw_K8sPluginConfig(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_K8sPluginConfig(vStringSlice, result))
}

func TestK8sPluginConfig_GetPFlagSet(t *testing.T) {
	val := K8sPluginConfig{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestK8sPluginConfig_SetFlags(t *testing.T) {
	actual := K8sPluginConfig{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_inject-finalizer", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("inject-finalizer", testValue)
			if vBool, err := cmdFlags.GetBool("inject-finalizer"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vBool), &actual.InjectFinalizer)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-cpus", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultK8sConfig.DefaultCPURequest.String()

			cmdFlags.Set("default-cpus", testValue)
			if vString, err := cmdFlags.GetString("default-cpus"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.DefaultCPURequest)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-memory", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultK8sConfig.DefaultMemoryRequest.String()

			cmdFlags.Set("default-memory", testValue)
			if vString, err := cmdFlags.GetString("default-memory"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.DefaultMemoryRequest)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_scheduler-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("scheduler-name", testValue)
			if vString, err := cmdFlags.GetString("scheduler-name"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.SchedulerName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.name", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.name"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.NamePrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.image", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.image", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.image"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.Image)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.default-input-path", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.default-input-path", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.default-input-path"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.DefaultInputDataPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.default-output-path", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.default-output-path", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.default-output-path"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.DefaultOutputPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.input-vol-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.input-vol-name", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.input-vol-name"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.InputVolumeName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.output-vol-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.output-vol-name", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.output-vol-name"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.OutputVolumeName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.cpu", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.cpu", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.cpu"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.CPU)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.memory", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.memory", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.memory"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.Memory)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_co-pilot.storage", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("co-pilot.storage", testValue)
			if vString, err := cmdFlags.GetString("co-pilot.storage"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.CoPilot.Storage)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_delete-resource-on-finalize", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("delete-resource-on-finalize", testValue)
			if vBool, err := cmdFlags.GetBool("delete-resource-on-finalize"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vBool), &actual.DeleteResourceOnFinalize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-pod-template-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("default-pod-template-name", testValue)
			if vString, err := cmdFlags.GetString("default-pod-template-name"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.DefaultPodTemplateName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_default-pod-template-resync", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultK8sConfig.DefaultPodTemplateResync.String()

			cmdFlags.Set("default-pod-template-resync", testValue)
			if vString, err := cmdFlags.GetString("default-pod-template-resync"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vString), &actual.DefaultPodTemplateResync)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_send-object-events", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("send-object-events", testValue)
			if vBool, err := cmdFlags.GetBool("send-object-events"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vBool), &actual.SendObjectEvents)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_update-base-backoff-duration", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("update-base-backoff-duration", testValue)
			if vInt, err := cmdFlags.GetInt("update-base-backoff-duration"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vInt), &actual.UpdateBaseBackoffDuration)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_update-backoff-retries", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("update-backoff-retries", testValue)
			if vInt, err := cmdFlags.GetInt("update-backoff-retries"); err == nil {
				testDecodeJson_K8sPluginConfig(t, fmt.Sprintf("%v", vInt), &actual.UpdateBackoffRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
