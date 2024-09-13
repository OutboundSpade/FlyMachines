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

// checks if the FlyMachineGuest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyMachineGuest{}

// FlyMachineGuest struct for FlyMachineGuest
type FlyMachineGuest struct {
	CpuKind *string `json:"cpu_kind,omitempty"`
	Cpus *int32 `json:"cpus,omitempty"`
	GpuKind *string `json:"gpu_kind,omitempty"`
	Gpus *int32 `json:"gpus,omitempty"`
	HostDedicationId *string `json:"host_dedication_id,omitempty"`
	KernelArgs []string `json:"kernel_args,omitempty"`
	MemoryMb *int32 `json:"memory_mb,omitempty"`
}

// NewFlyMachineGuest instantiates a new FlyMachineGuest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyMachineGuest() *FlyMachineGuest {
	this := FlyMachineGuest{}
	return &this
}

// NewFlyMachineGuestWithDefaults instantiates a new FlyMachineGuest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyMachineGuestWithDefaults() *FlyMachineGuest {
	this := FlyMachineGuest{}
	return &this
}

// GetCpuKind returns the CpuKind field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetCpuKind() string {
	if o == nil || IsNil(o.CpuKind) {
		var ret string
		return ret
	}
	return *o.CpuKind
}

// GetCpuKindOk returns a tuple with the CpuKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetCpuKindOk() (*string, bool) {
	if o == nil || IsNil(o.CpuKind) {
		return nil, false
	}
	return o.CpuKind, true
}

// HasCpuKind returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasCpuKind() bool {
	if o != nil && !IsNil(o.CpuKind) {
		return true
	}

	return false
}

// SetCpuKind gets a reference to the given string and assigns it to the CpuKind field.
func (o *FlyMachineGuest) SetCpuKind(v string) {
	o.CpuKind = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetCpus() int32 {
	if o == nil || IsNil(o.Cpus) {
		var ret int32
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetCpusOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int32 and assigns it to the Cpus field.
func (o *FlyMachineGuest) SetCpus(v int32) {
	o.Cpus = &v
}

// GetGpuKind returns the GpuKind field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetGpuKind() string {
	if o == nil || IsNil(o.GpuKind) {
		var ret string
		return ret
	}
	return *o.GpuKind
}

// GetGpuKindOk returns a tuple with the GpuKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetGpuKindOk() (*string, bool) {
	if o == nil || IsNil(o.GpuKind) {
		return nil, false
	}
	return o.GpuKind, true
}

// HasGpuKind returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasGpuKind() bool {
	if o != nil && !IsNil(o.GpuKind) {
		return true
	}

	return false
}

// SetGpuKind gets a reference to the given string and assigns it to the GpuKind field.
func (o *FlyMachineGuest) SetGpuKind(v string) {
	o.GpuKind = &v
}

// GetGpus returns the Gpus field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetGpus() int32 {
	if o == nil || IsNil(o.Gpus) {
		var ret int32
		return ret
	}
	return *o.Gpus
}

// GetGpusOk returns a tuple with the Gpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetGpusOk() (*int32, bool) {
	if o == nil || IsNil(o.Gpus) {
		return nil, false
	}
	return o.Gpus, true
}

// HasGpus returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasGpus() bool {
	if o != nil && !IsNil(o.Gpus) {
		return true
	}

	return false
}

// SetGpus gets a reference to the given int32 and assigns it to the Gpus field.
func (o *FlyMachineGuest) SetGpus(v int32) {
	o.Gpus = &v
}

// GetHostDedicationId returns the HostDedicationId field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetHostDedicationId() string {
	if o == nil || IsNil(o.HostDedicationId) {
		var ret string
		return ret
	}
	return *o.HostDedicationId
}

// GetHostDedicationIdOk returns a tuple with the HostDedicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetHostDedicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostDedicationId) {
		return nil, false
	}
	return o.HostDedicationId, true
}

// HasHostDedicationId returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasHostDedicationId() bool {
	if o != nil && !IsNil(o.HostDedicationId) {
		return true
	}

	return false
}

// SetHostDedicationId gets a reference to the given string and assigns it to the HostDedicationId field.
func (o *FlyMachineGuest) SetHostDedicationId(v string) {
	o.HostDedicationId = &v
}

// GetKernelArgs returns the KernelArgs field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetKernelArgs() []string {
	if o == nil || IsNil(o.KernelArgs) {
		var ret []string
		return ret
	}
	return o.KernelArgs
}

// GetKernelArgsOk returns a tuple with the KernelArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetKernelArgsOk() ([]string, bool) {
	if o == nil || IsNil(o.KernelArgs) {
		return nil, false
	}
	return o.KernelArgs, true
}

// HasKernelArgs returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasKernelArgs() bool {
	if o != nil && !IsNil(o.KernelArgs) {
		return true
	}

	return false
}

// SetKernelArgs gets a reference to the given []string and assigns it to the KernelArgs field.
func (o *FlyMachineGuest) SetKernelArgs(v []string) {
	o.KernelArgs = v
}

// GetMemoryMb returns the MemoryMb field value if set, zero value otherwise.
func (o *FlyMachineGuest) GetMemoryMb() int32 {
	if o == nil || IsNil(o.MemoryMb) {
		var ret int32
		return ret
	}
	return *o.MemoryMb
}

// GetMemoryMbOk returns a tuple with the MemoryMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineGuest) GetMemoryMbOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryMb) {
		return nil, false
	}
	return o.MemoryMb, true
}

// HasMemoryMb returns a boolean if a field has been set.
func (o *FlyMachineGuest) HasMemoryMb() bool {
	if o != nil && !IsNil(o.MemoryMb) {
		return true
	}

	return false
}

// SetMemoryMb gets a reference to the given int32 and assigns it to the MemoryMb field.
func (o *FlyMachineGuest) SetMemoryMb(v int32) {
	o.MemoryMb = &v
}

func (o FlyMachineGuest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyMachineGuest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuKind) {
		toSerialize["cpu_kind"] = o.CpuKind
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.GpuKind) {
		toSerialize["gpu_kind"] = o.GpuKind
	}
	if !IsNil(o.Gpus) {
		toSerialize["gpus"] = o.Gpus
	}
	if !IsNil(o.HostDedicationId) {
		toSerialize["host_dedication_id"] = o.HostDedicationId
	}
	if !IsNil(o.KernelArgs) {
		toSerialize["kernel_args"] = o.KernelArgs
	}
	if !IsNil(o.MemoryMb) {
		toSerialize["memory_mb"] = o.MemoryMb
	}
	return toSerialize, nil
}

type NullableFlyMachineGuest struct {
	value *FlyMachineGuest
	isSet bool
}

func (v NullableFlyMachineGuest) Get() *FlyMachineGuest {
	return v.value
}

func (v *NullableFlyMachineGuest) Set(val *FlyMachineGuest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyMachineGuest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyMachineGuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyMachineGuest(val *FlyMachineGuest) *NullableFlyMachineGuest {
	return &NullableFlyMachineGuest{value: val, isSet: true}
}

func (v NullableFlyMachineGuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyMachineGuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

