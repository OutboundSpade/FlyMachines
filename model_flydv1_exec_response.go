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

// checks if the Flydv1ExecResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Flydv1ExecResponse{}

// Flydv1ExecResponse struct for Flydv1ExecResponse
type Flydv1ExecResponse struct {
	ExitCode *int32 `json:"exit_code,omitempty"`
	ExitSignal *int32 `json:"exit_signal,omitempty"`
	Stderr *string `json:"stderr,omitempty"`
	Stdout *string `json:"stdout,omitempty"`
}

// NewFlydv1ExecResponse instantiates a new Flydv1ExecResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlydv1ExecResponse() *Flydv1ExecResponse {
	this := Flydv1ExecResponse{}
	return &this
}

// NewFlydv1ExecResponseWithDefaults instantiates a new Flydv1ExecResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlydv1ExecResponseWithDefaults() *Flydv1ExecResponse {
	this := Flydv1ExecResponse{}
	return &this
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *Flydv1ExecResponse) GetExitCode() int32 {
	if o == nil || IsNil(o.ExitCode) {
		var ret int32
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flydv1ExecResponse) GetExitCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ExitCode) {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *Flydv1ExecResponse) HasExitCode() bool {
	if o != nil && !IsNil(o.ExitCode) {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given int32 and assigns it to the ExitCode field.
func (o *Flydv1ExecResponse) SetExitCode(v int32) {
	o.ExitCode = &v
}

// GetExitSignal returns the ExitSignal field value if set, zero value otherwise.
func (o *Flydv1ExecResponse) GetExitSignal() int32 {
	if o == nil || IsNil(o.ExitSignal) {
		var ret int32
		return ret
	}
	return *o.ExitSignal
}

// GetExitSignalOk returns a tuple with the ExitSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flydv1ExecResponse) GetExitSignalOk() (*int32, bool) {
	if o == nil || IsNil(o.ExitSignal) {
		return nil, false
	}
	return o.ExitSignal, true
}

// HasExitSignal returns a boolean if a field has been set.
func (o *Flydv1ExecResponse) HasExitSignal() bool {
	if o != nil && !IsNil(o.ExitSignal) {
		return true
	}

	return false
}

// SetExitSignal gets a reference to the given int32 and assigns it to the ExitSignal field.
func (o *Flydv1ExecResponse) SetExitSignal(v int32) {
	o.ExitSignal = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise.
func (o *Flydv1ExecResponse) GetStderr() string {
	if o == nil || IsNil(o.Stderr) {
		var ret string
		return ret
	}
	return *o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flydv1ExecResponse) GetStderrOk() (*string, bool) {
	if o == nil || IsNil(o.Stderr) {
		return nil, false
	}
	return o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *Flydv1ExecResponse) HasStderr() bool {
	if o != nil && !IsNil(o.Stderr) {
		return true
	}

	return false
}

// SetStderr gets a reference to the given string and assigns it to the Stderr field.
func (o *Flydv1ExecResponse) SetStderr(v string) {
	o.Stderr = &v
}

// GetStdout returns the Stdout field value if set, zero value otherwise.
func (o *Flydv1ExecResponse) GetStdout() string {
	if o == nil || IsNil(o.Stdout) {
		var ret string
		return ret
	}
	return *o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flydv1ExecResponse) GetStdoutOk() (*string, bool) {
	if o == nil || IsNil(o.Stdout) {
		return nil, false
	}
	return o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *Flydv1ExecResponse) HasStdout() bool {
	if o != nil && !IsNil(o.Stdout) {
		return true
	}

	return false
}

// SetStdout gets a reference to the given string and assigns it to the Stdout field.
func (o *Flydv1ExecResponse) SetStdout(v string) {
	o.Stdout = &v
}

func (o Flydv1ExecResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Flydv1ExecResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExitCode) {
		toSerialize["exit_code"] = o.ExitCode
	}
	if !IsNil(o.ExitSignal) {
		toSerialize["exit_signal"] = o.ExitSignal
	}
	if !IsNil(o.Stderr) {
		toSerialize["stderr"] = o.Stderr
	}
	if !IsNil(o.Stdout) {
		toSerialize["stdout"] = o.Stdout
	}
	return toSerialize, nil
}

type NullableFlydv1ExecResponse struct {
	value *Flydv1ExecResponse
	isSet bool
}

func (v NullableFlydv1ExecResponse) Get() *Flydv1ExecResponse {
	return v.value
}

func (v *NullableFlydv1ExecResponse) Set(val *Flydv1ExecResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlydv1ExecResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlydv1ExecResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlydv1ExecResponse(val *Flydv1ExecResponse) *NullableFlydv1ExecResponse {
	return &NullableFlydv1ExecResponse{value: val, isSet: true}
}

func (v NullableFlydv1ExecResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlydv1ExecResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


