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

// checks if the UpdateMachineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMachineRequest{}

// UpdateMachineRequest struct for UpdateMachineRequest
type UpdateMachineRequest struct {
	// An object defining the Machine configuration
	Config *FlyMachineConfig `json:"config,omitempty"`
	CurrentVersion *string `json:"current_version,omitempty"`
	LeaseTtl *int32 `json:"lease_ttl,omitempty"`
	Lsvd *bool `json:"lsvd,omitempty"`
	// Unique name for this Machine. If omitted, one is generated for you
	Name *string `json:"name,omitempty"`
	// The target region. Omitting this param launches in the same region as your WireGuard peer connection (somewhere near you).
	Region *string `json:"region,omitempty"`
	SkipLaunch *bool `json:"skip_launch,omitempty"`
	SkipServiceRegistration *bool `json:"skip_service_registration,omitempty"`
}

// NewUpdateMachineRequest instantiates a new UpdateMachineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMachineRequest() *UpdateMachineRequest {
	this := UpdateMachineRequest{}
	return &this
}

// NewUpdateMachineRequestWithDefaults instantiates a new UpdateMachineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMachineRequestWithDefaults() *UpdateMachineRequest {
	this := UpdateMachineRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetConfig() FlyMachineConfig {
	if o == nil || IsNil(o.Config) {
		var ret FlyMachineConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetConfigOk() (*FlyMachineConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given FlyMachineConfig and assigns it to the Config field.
func (o *UpdateMachineRequest) SetConfig(v FlyMachineConfig) {
	o.Config = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetCurrentVersion() string {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetCurrentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *UpdateMachineRequest) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetLeaseTtl returns the LeaseTtl field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetLeaseTtl() int32 {
	if o == nil || IsNil(o.LeaseTtl) {
		var ret int32
		return ret
	}
	return *o.LeaseTtl
}

// GetLeaseTtlOk returns a tuple with the LeaseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetLeaseTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.LeaseTtl) {
		return nil, false
	}
	return o.LeaseTtl, true
}

// HasLeaseTtl returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasLeaseTtl() bool {
	if o != nil && !IsNil(o.LeaseTtl) {
		return true
	}

	return false
}

// SetLeaseTtl gets a reference to the given int32 and assigns it to the LeaseTtl field.
func (o *UpdateMachineRequest) SetLeaseTtl(v int32) {
	o.LeaseTtl = &v
}

// GetLsvd returns the Lsvd field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetLsvd() bool {
	if o == nil || IsNil(o.Lsvd) {
		var ret bool
		return ret
	}
	return *o.Lsvd
}

// GetLsvdOk returns a tuple with the Lsvd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetLsvdOk() (*bool, bool) {
	if o == nil || IsNil(o.Lsvd) {
		return nil, false
	}
	return o.Lsvd, true
}

// HasLsvd returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasLsvd() bool {
	if o != nil && !IsNil(o.Lsvd) {
		return true
	}

	return false
}

// SetLsvd gets a reference to the given bool and assigns it to the Lsvd field.
func (o *UpdateMachineRequest) SetLsvd(v bool) {
	o.Lsvd = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateMachineRequest) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UpdateMachineRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSkipLaunch returns the SkipLaunch field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetSkipLaunch() bool {
	if o == nil || IsNil(o.SkipLaunch) {
		var ret bool
		return ret
	}
	return *o.SkipLaunch
}

// GetSkipLaunchOk returns a tuple with the SkipLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetSkipLaunchOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipLaunch) {
		return nil, false
	}
	return o.SkipLaunch, true
}

// HasSkipLaunch returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasSkipLaunch() bool {
	if o != nil && !IsNil(o.SkipLaunch) {
		return true
	}

	return false
}

// SetSkipLaunch gets a reference to the given bool and assigns it to the SkipLaunch field.
func (o *UpdateMachineRequest) SetSkipLaunch(v bool) {
	o.SkipLaunch = &v
}

// GetSkipServiceRegistration returns the SkipServiceRegistration field value if set, zero value otherwise.
func (o *UpdateMachineRequest) GetSkipServiceRegistration() bool {
	if o == nil || IsNil(o.SkipServiceRegistration) {
		var ret bool
		return ret
	}
	return *o.SkipServiceRegistration
}

// GetSkipServiceRegistrationOk returns a tuple with the SkipServiceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMachineRequest) GetSkipServiceRegistrationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipServiceRegistration) {
		return nil, false
	}
	return o.SkipServiceRegistration, true
}

// HasSkipServiceRegistration returns a boolean if a field has been set.
func (o *UpdateMachineRequest) HasSkipServiceRegistration() bool {
	if o != nil && !IsNil(o.SkipServiceRegistration) {
		return true
	}

	return false
}

// SetSkipServiceRegistration gets a reference to the given bool and assigns it to the SkipServiceRegistration field.
func (o *UpdateMachineRequest) SetSkipServiceRegistration(v bool) {
	o.SkipServiceRegistration = &v
}

func (o UpdateMachineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMachineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CurrentVersion) {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if !IsNil(o.LeaseTtl) {
		toSerialize["lease_ttl"] = o.LeaseTtl
	}
	if !IsNil(o.Lsvd) {
		toSerialize["lsvd"] = o.Lsvd
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SkipLaunch) {
		toSerialize["skip_launch"] = o.SkipLaunch
	}
	if !IsNil(o.SkipServiceRegistration) {
		toSerialize["skip_service_registration"] = o.SkipServiceRegistration
	}
	return toSerialize, nil
}

type NullableUpdateMachineRequest struct {
	value *UpdateMachineRequest
	isSet bool
}

func (v NullableUpdateMachineRequest) Get() *UpdateMachineRequest {
	return v.value
}

func (v *NullableUpdateMachineRequest) Set(val *UpdateMachineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMachineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMachineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMachineRequest(val *UpdateMachineRequest) *NullableUpdateMachineRequest {
	return &NullableUpdateMachineRequest{value: val, isSet: true}
}

func (v NullableUpdateMachineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMachineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


