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

// checks if the ListAppsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAppsResponse{}

// ListAppsResponse struct for ListAppsResponse
type ListAppsResponse struct {
	Apps []ListApp `json:"apps,omitempty"`
	TotalApps *int32 `json:"total_apps,omitempty"`
}

// NewListAppsResponse instantiates a new ListAppsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAppsResponse() *ListAppsResponse {
	this := ListAppsResponse{}
	return &this
}

// NewListAppsResponseWithDefaults instantiates a new ListAppsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAppsResponseWithDefaults() *ListAppsResponse {
	this := ListAppsResponse{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ListAppsResponse) GetApps() []ListApp {
	if o == nil || IsNil(o.Apps) {
		var ret []ListApp
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponse) GetAppsOk() ([]ListApp, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ListAppsResponse) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []ListApp and assigns it to the Apps field.
func (o *ListAppsResponse) SetApps(v []ListApp) {
	o.Apps = v
}

// GetTotalApps returns the TotalApps field value if set, zero value otherwise.
func (o *ListAppsResponse) GetTotalApps() int32 {
	if o == nil || IsNil(o.TotalApps) {
		var ret int32
		return ret
	}
	return *o.TotalApps
}

// GetTotalAppsOk returns a tuple with the TotalApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponse) GetTotalAppsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalApps) {
		return nil, false
	}
	return o.TotalApps, true
}

// HasTotalApps returns a boolean if a field has been set.
func (o *ListAppsResponse) HasTotalApps() bool {
	if o != nil && !IsNil(o.TotalApps) {
		return true
	}

	return false
}

// SetTotalApps gets a reference to the given int32 and assigns it to the TotalApps field.
func (o *ListAppsResponse) SetTotalApps(v int32) {
	o.TotalApps = &v
}

func (o ListAppsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAppsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.TotalApps) {
		toSerialize["total_apps"] = o.TotalApps
	}
	return toSerialize, nil
}

type NullableListAppsResponse struct {
	value *ListAppsResponse
	isSet bool
}

func (v NullableListAppsResponse) Get() *ListAppsResponse {
	return v.value
}

func (v *NullableListAppsResponse) Set(val *ListAppsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAppsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAppsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAppsResponse(val *ListAppsResponse) *NullableListAppsResponse {
	return &NullableListAppsResponse{value: val, isSet: true}
}

func (v NullableListAppsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAppsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


