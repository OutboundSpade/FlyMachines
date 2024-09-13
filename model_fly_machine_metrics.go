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

// checks if the FlyMachineMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyMachineMetrics{}

// FlyMachineMetrics struct for FlyMachineMetrics
type FlyMachineMetrics struct {
	Path *string `json:"path,omitempty"`
	Port *int32 `json:"port,omitempty"`
}

// NewFlyMachineMetrics instantiates a new FlyMachineMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyMachineMetrics() *FlyMachineMetrics {
	this := FlyMachineMetrics{}
	return &this
}

// NewFlyMachineMetricsWithDefaults instantiates a new FlyMachineMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyMachineMetricsWithDefaults() *FlyMachineMetrics {
	this := FlyMachineMetrics{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FlyMachineMetrics) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineMetrics) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FlyMachineMetrics) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FlyMachineMetrics) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FlyMachineMetrics) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineMetrics) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FlyMachineMetrics) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *FlyMachineMetrics) SetPort(v int32) {
	o.Port = &v
}

func (o FlyMachineMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyMachineMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableFlyMachineMetrics struct {
	value *FlyMachineMetrics
	isSet bool
}

func (v NullableFlyMachineMetrics) Get() *FlyMachineMetrics {
	return v.value
}

func (v *NullableFlyMachineMetrics) Set(val *FlyMachineMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyMachineMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyMachineMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyMachineMetrics(val *FlyMachineMetrics) *NullableFlyMachineMetrics {
	return &NullableFlyMachineMetrics{value: val, isSet: true}
}

func (v NullableFlyMachineMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyMachineMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


