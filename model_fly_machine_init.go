/*
Machines API

This site hosts documentation generated from the Fly.io Machines API OpenAPI specification. Visit our complete [Machines API docs](https://fly.io/docs/machines/api/) for how to get started, more information about each endpoint, parameter descriptions, and examples.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package machines

import (
	"encoding/json"
)

// checks if the FlyMachineInit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyMachineInit{}

// FlyMachineInit struct for FlyMachineInit
type FlyMachineInit struct {
	Cmd []string `json:"cmd,omitempty"`
	Entrypoint []string `json:"entrypoint,omitempty"`
	Exec []string `json:"exec,omitempty"`
	KernelArgs []string `json:"kernel_args,omitempty"`
	SwapSizeMb *int32 `json:"swap_size_mb,omitempty"`
	Tty *bool `json:"tty,omitempty"`
}

// NewFlyMachineInit instantiates a new FlyMachineInit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyMachineInit() *FlyMachineInit {
	this := FlyMachineInit{}
	return &this
}

// NewFlyMachineInitWithDefaults instantiates a new FlyMachineInit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyMachineInitWithDefaults() *FlyMachineInit {
	this := FlyMachineInit{}
	return &this
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *FlyMachineInit) GetCmd() []string {
	if o == nil || IsNil(o.Cmd) {
		var ret []string
		return ret
	}
	return o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetCmdOk() ([]string, bool) {
	if o == nil || IsNil(o.Cmd) {
		return nil, false
	}
	return o.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *FlyMachineInit) HasCmd() bool {
	if o != nil && !IsNil(o.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the Cmd field.
func (o *FlyMachineInit) SetCmd(v []string) {
	o.Cmd = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *FlyMachineInit) GetEntrypoint() []string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret []string
		return ret
	}
	return o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetEntrypointOk() ([]string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *FlyMachineInit) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given []string and assigns it to the Entrypoint field.
func (o *FlyMachineInit) SetEntrypoint(v []string) {
	o.Entrypoint = v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *FlyMachineInit) GetExec() []string {
	if o == nil || IsNil(o.Exec) {
		var ret []string
		return ret
	}
	return o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetExecOk() ([]string, bool) {
	if o == nil || IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *FlyMachineInit) HasExec() bool {
	if o != nil && !IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given []string and assigns it to the Exec field.
func (o *FlyMachineInit) SetExec(v []string) {
	o.Exec = v
}

// GetKernelArgs returns the KernelArgs field value if set, zero value otherwise.
func (o *FlyMachineInit) GetKernelArgs() []string {
	if o == nil || IsNil(o.KernelArgs) {
		var ret []string
		return ret
	}
	return o.KernelArgs
}

// GetKernelArgsOk returns a tuple with the KernelArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetKernelArgsOk() ([]string, bool) {
	if o == nil || IsNil(o.KernelArgs) {
		return nil, false
	}
	return o.KernelArgs, true
}

// HasKernelArgs returns a boolean if a field has been set.
func (o *FlyMachineInit) HasKernelArgs() bool {
	if o != nil && !IsNil(o.KernelArgs) {
		return true
	}

	return false
}

// SetKernelArgs gets a reference to the given []string and assigns it to the KernelArgs field.
func (o *FlyMachineInit) SetKernelArgs(v []string) {
	o.KernelArgs = v
}

// GetSwapSizeMb returns the SwapSizeMb field value if set, zero value otherwise.
func (o *FlyMachineInit) GetSwapSizeMb() int32 {
	if o == nil || IsNil(o.SwapSizeMb) {
		var ret int32
		return ret
	}
	return *o.SwapSizeMb
}

// GetSwapSizeMbOk returns a tuple with the SwapSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetSwapSizeMbOk() (*int32, bool) {
	if o == nil || IsNil(o.SwapSizeMb) {
		return nil, false
	}
	return o.SwapSizeMb, true
}

// HasSwapSizeMb returns a boolean if a field has been set.
func (o *FlyMachineInit) HasSwapSizeMb() bool {
	if o != nil && !IsNil(o.SwapSizeMb) {
		return true
	}

	return false
}

// SetSwapSizeMb gets a reference to the given int32 and assigns it to the SwapSizeMb field.
func (o *FlyMachineInit) SetSwapSizeMb(v int32) {
	o.SwapSizeMb = &v
}

// GetTty returns the Tty field value if set, zero value otherwise.
func (o *FlyMachineInit) GetTty() bool {
	if o == nil || IsNil(o.Tty) {
		var ret bool
		return ret
	}
	return *o.Tty
}

// GetTtyOk returns a tuple with the Tty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineInit) GetTtyOk() (*bool, bool) {
	if o == nil || IsNil(o.Tty) {
		return nil, false
	}
	return o.Tty, true
}

// HasTty returns a boolean if a field has been set.
func (o *FlyMachineInit) HasTty() bool {
	if o != nil && !IsNil(o.Tty) {
		return true
	}

	return false
}

// SetTty gets a reference to the given bool and assigns it to the Tty field.
func (o *FlyMachineInit) SetTty(v bool) {
	o.Tty = &v
}

func (o FlyMachineInit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyMachineInit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if !IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !IsNil(o.KernelArgs) {
		toSerialize["kernel_args"] = o.KernelArgs
	}
	if !IsNil(o.SwapSizeMb) {
		toSerialize["swap_size_mb"] = o.SwapSizeMb
	}
	if !IsNil(o.Tty) {
		toSerialize["tty"] = o.Tty
	}
	return toSerialize, nil
}

type NullableFlyMachineInit struct {
	value *FlyMachineInit
	isSet bool
}

func (v NullableFlyMachineInit) Get() *FlyMachineInit {
	return v.value
}

func (v *NullableFlyMachineInit) Set(val *FlyMachineInit) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyMachineInit) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyMachineInit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyMachineInit(val *FlyMachineInit) *NullableFlyMachineInit {
	return &NullableFlyMachineInit{value: val, isSet: true}
}

func (v NullableFlyMachineInit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyMachineInit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


