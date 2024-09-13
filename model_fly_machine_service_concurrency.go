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

// checks if the FlyMachineServiceConcurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyMachineServiceConcurrency{}

// FlyMachineServiceConcurrency struct for FlyMachineServiceConcurrency
type FlyMachineServiceConcurrency struct {
	HardLimit *int32 `json:"hard_limit,omitempty"`
	SoftLimit *int32 `json:"soft_limit,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewFlyMachineServiceConcurrency instantiates a new FlyMachineServiceConcurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyMachineServiceConcurrency() *FlyMachineServiceConcurrency {
	this := FlyMachineServiceConcurrency{}
	return &this
}

// NewFlyMachineServiceConcurrencyWithDefaults instantiates a new FlyMachineServiceConcurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyMachineServiceConcurrencyWithDefaults() *FlyMachineServiceConcurrency {
	this := FlyMachineServiceConcurrency{}
	return &this
}

// GetHardLimit returns the HardLimit field value if set, zero value otherwise.
func (o *FlyMachineServiceConcurrency) GetHardLimit() int32 {
	if o == nil || IsNil(o.HardLimit) {
		var ret int32
		return ret
	}
	return *o.HardLimit
}

// GetHardLimitOk returns a tuple with the HardLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineServiceConcurrency) GetHardLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.HardLimit) {
		return nil, false
	}
	return o.HardLimit, true
}

// HasHardLimit returns a boolean if a field has been set.
func (o *FlyMachineServiceConcurrency) HasHardLimit() bool {
	if o != nil && !IsNil(o.HardLimit) {
		return true
	}

	return false
}

// SetHardLimit gets a reference to the given int32 and assigns it to the HardLimit field.
func (o *FlyMachineServiceConcurrency) SetHardLimit(v int32) {
	o.HardLimit = &v
}

// GetSoftLimit returns the SoftLimit field value if set, zero value otherwise.
func (o *FlyMachineServiceConcurrency) GetSoftLimit() int32 {
	if o == nil || IsNil(o.SoftLimit) {
		var ret int32
		return ret
	}
	return *o.SoftLimit
}

// GetSoftLimitOk returns a tuple with the SoftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineServiceConcurrency) GetSoftLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.SoftLimit) {
		return nil, false
	}
	return o.SoftLimit, true
}

// HasSoftLimit returns a boolean if a field has been set.
func (o *FlyMachineServiceConcurrency) HasSoftLimit() bool {
	if o != nil && !IsNil(o.SoftLimit) {
		return true
	}

	return false
}

// SetSoftLimit gets a reference to the given int32 and assigns it to the SoftLimit field.
func (o *FlyMachineServiceConcurrency) SetSoftLimit(v int32) {
	o.SoftLimit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlyMachineServiceConcurrency) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineServiceConcurrency) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlyMachineServiceConcurrency) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlyMachineServiceConcurrency) SetType(v string) {
	o.Type = &v
}

func (o FlyMachineServiceConcurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyMachineServiceConcurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardLimit) {
		toSerialize["hard_limit"] = o.HardLimit
	}
	if !IsNil(o.SoftLimit) {
		toSerialize["soft_limit"] = o.SoftLimit
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFlyMachineServiceConcurrency struct {
	value *FlyMachineServiceConcurrency
	isSet bool
}

func (v NullableFlyMachineServiceConcurrency) Get() *FlyMachineServiceConcurrency {
	return v.value
}

func (v *NullableFlyMachineServiceConcurrency) Set(val *FlyMachineServiceConcurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyMachineServiceConcurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyMachineServiceConcurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyMachineServiceConcurrency(val *FlyMachineServiceConcurrency) *NullableFlyMachineServiceConcurrency {
	return &NullableFlyMachineServiceConcurrency{value: val, isSet: true}
}

func (v NullableFlyMachineServiceConcurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyMachineServiceConcurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

