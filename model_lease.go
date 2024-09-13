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

// checks if the Lease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lease{}

// Lease struct for Lease
type Lease struct {
	// Description or reason for the Lease.
	Description *string `json:"description,omitempty"`
	// ExpiresAt is the unix timestamp in UTC to denote when the Lease will no longer be valid.
	ExpiresAt *int32 `json:"expires_at,omitempty"`
	// Nonce is the unique ID autogenerated and associated with the Lease.
	Nonce *string `json:"nonce,omitempty"`
	// Owner is the user identifier which acquired the Lease.
	Owner *string `json:"owner,omitempty"`
	// Machine version
	Version *string `json:"version,omitempty"`
}

// NewLease instantiates a new Lease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLease() *Lease {
	this := Lease{}
	return &this
}

// NewLeaseWithDefaults instantiates a new Lease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaseWithDefaults() *Lease {
	this := Lease{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Lease) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Lease) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Lease) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Lease) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetExpiresAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Lease) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int32 and assigns it to the ExpiresAt field.
func (o *Lease) SetExpiresAt(v int32) {
	o.ExpiresAt = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Lease) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Lease) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Lease) SetNonce(v string) {
	o.Nonce = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Lease) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Lease) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Lease) SetOwner(v string) {
	o.Owner = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Lease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Lease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Lease) SetVersion(v string) {
	o.Version = &v
}

func (o Lease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableLease struct {
	value *Lease
	isSet bool
}

func (v NullableLease) Get() *Lease {
	return v.value
}

func (v *NullableLease) Set(val *Lease) {
	v.value = val
	v.isSet = true
}

func (v NullableLease) IsSet() bool {
	return v.isSet
}

func (v *NullableLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLease(val *Lease) *NullableLease {
	return &NullableLease{value: val, isSet: true}
}

func (v NullableLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


