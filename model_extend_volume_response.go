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

// checks if the ExtendVolumeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendVolumeResponse{}

// ExtendVolumeResponse struct for ExtendVolumeResponse
type ExtendVolumeResponse struct {
	NeedsRestart *bool `json:"needs_restart,omitempty"`
	Volume *Volume `json:"volume,omitempty"`
}

// NewExtendVolumeResponse instantiates a new ExtendVolumeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendVolumeResponse() *ExtendVolumeResponse {
	this := ExtendVolumeResponse{}
	return &this
}

// NewExtendVolumeResponseWithDefaults instantiates a new ExtendVolumeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendVolumeResponseWithDefaults() *ExtendVolumeResponse {
	this := ExtendVolumeResponse{}
	return &this
}

// GetNeedsRestart returns the NeedsRestart field value if set, zero value otherwise.
func (o *ExtendVolumeResponse) GetNeedsRestart() bool {
	if o == nil || IsNil(o.NeedsRestart) {
		var ret bool
		return ret
	}
	return *o.NeedsRestart
}

// GetNeedsRestartOk returns a tuple with the NeedsRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendVolumeResponse) GetNeedsRestartOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsRestart) {
		return nil, false
	}
	return o.NeedsRestart, true
}

// HasNeedsRestart returns a boolean if a field has been set.
func (o *ExtendVolumeResponse) HasNeedsRestart() bool {
	if o != nil && !IsNil(o.NeedsRestart) {
		return true
	}

	return false
}

// SetNeedsRestart gets a reference to the given bool and assigns it to the NeedsRestart field.
func (o *ExtendVolumeResponse) SetNeedsRestart(v bool) {
	o.NeedsRestart = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ExtendVolumeResponse) GetVolume() Volume {
	if o == nil || IsNil(o.Volume) {
		var ret Volume
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendVolumeResponse) GetVolumeOk() (*Volume, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ExtendVolumeResponse) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given Volume and assigns it to the Volume field.
func (o *ExtendVolumeResponse) SetVolume(v Volume) {
	o.Volume = &v
}

func (o ExtendVolumeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendVolumeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NeedsRestart) {
		toSerialize["needs_restart"] = o.NeedsRestart
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableExtendVolumeResponse struct {
	value *ExtendVolumeResponse
	isSet bool
}

func (v NullableExtendVolumeResponse) Get() *ExtendVolumeResponse {
	return v.value
}

func (v *NullableExtendVolumeResponse) Set(val *ExtendVolumeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendVolumeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendVolumeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendVolumeResponse(val *ExtendVolumeResponse) *NullableExtendVolumeResponse {
	return &NullableExtendVolumeResponse{value: val, isSet: true}
}

func (v NullableExtendVolumeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendVolumeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


