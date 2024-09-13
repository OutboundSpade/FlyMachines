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

// checks if the SignalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalRequest{}

// SignalRequest struct for SignalRequest
type SignalRequest struct {
	Signal *string `json:"signal,omitempty"`
}

// NewSignalRequest instantiates a new SignalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalRequest() *SignalRequest {
	this := SignalRequest{}
	return &this
}

// NewSignalRequestWithDefaults instantiates a new SignalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalRequestWithDefaults() *SignalRequest {
	this := SignalRequest{}
	return &this
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *SignalRequest) GetSignal() string {
	if o == nil || IsNil(o.Signal) {
		var ret string
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalRequest) GetSignalOk() (*string, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *SignalRequest) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given string and assigns it to the Signal field.
func (o *SignalRequest) SetSignal(v string) {
	o.Signal = &v
}

func (o SignalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	return toSerialize, nil
}

type NullableSignalRequest struct {
	value *SignalRequest
	isSet bool
}

func (v NullableSignalRequest) Get() *SignalRequest {
	return v.value
}

func (v *NullableSignalRequest) Set(val *SignalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalRequest(val *SignalRequest) *NullableSignalRequest {
	return &NullableSignalRequest{value: val, isSet: true}
}

func (v NullableSignalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

