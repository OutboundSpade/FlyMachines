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

// checks if the Machine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Machine{}

// Machine struct for Machine
type Machine struct {
	Checks []CheckStatus `json:"checks,omitempty"`
	Config *FlyMachineConfig `json:"config,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Events []MachineEvent `json:"events,omitempty"`
	HostStatus *string `json:"host_status,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageRef *ImageRef `json:"image_ref,omitempty"`
	IncompleteConfig *FlyMachineConfig `json:"incomplete_config,omitempty"`
	// InstanceID is unique for each version of the machine
	InstanceId *string `json:"instance_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Nonce is only every returned on machine creation if a lease_duration was provided.
	Nonce *string `json:"nonce,omitempty"`
	// PrivateIP is the internal 6PN address of the machine.
	PrivateIp *string `json:"private_ip,omitempty"`
	Region *string `json:"region,omitempty"`
	State *string `json:"state,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewMachine instantiates a new Machine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachine() *Machine {
	this := Machine{}
	return &this
}

// NewMachineWithDefaults instantiates a new Machine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineWithDefaults() *Machine {
	this := Machine{}
	return &this
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *Machine) GetChecks() []CheckStatus {
	if o == nil || IsNil(o.Checks) {
		var ret []CheckStatus
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetChecksOk() ([]CheckStatus, bool) {
	if o == nil || IsNil(o.Checks) {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *Machine) HasChecks() bool {
	if o != nil && !IsNil(o.Checks) {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []CheckStatus and assigns it to the Checks field.
func (o *Machine) SetChecks(v []CheckStatus) {
	o.Checks = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Machine) GetConfig() FlyMachineConfig {
	if o == nil || IsNil(o.Config) {
		var ret FlyMachineConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetConfigOk() (*FlyMachineConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Machine) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given FlyMachineConfig and assigns it to the Config field.
func (o *Machine) SetConfig(v FlyMachineConfig) {
	o.Config = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Machine) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Machine) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Machine) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Machine) GetEvents() []MachineEvent {
	if o == nil || IsNil(o.Events) {
		var ret []MachineEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetEventsOk() ([]MachineEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Machine) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []MachineEvent and assigns it to the Events field.
func (o *Machine) SetEvents(v []MachineEvent) {
	o.Events = v
}

// GetHostStatus returns the HostStatus field value if set, zero value otherwise.
func (o *Machine) GetHostStatus() string {
	if o == nil || IsNil(o.HostStatus) {
		var ret string
		return ret
	}
	return *o.HostStatus
}

// GetHostStatusOk returns a tuple with the HostStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetHostStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HostStatus) {
		return nil, false
	}
	return o.HostStatus, true
}

// HasHostStatus returns a boolean if a field has been set.
func (o *Machine) HasHostStatus() bool {
	if o != nil && !IsNil(o.HostStatus) {
		return true
	}

	return false
}

// SetHostStatus gets a reference to the given string and assigns it to the HostStatus field.
func (o *Machine) SetHostStatus(v string) {
	o.HostStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Machine) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Machine) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Machine) SetId(v string) {
	o.Id = &v
}

// GetImageRef returns the ImageRef field value if set, zero value otherwise.
func (o *Machine) GetImageRef() ImageRef {
	if o == nil || IsNil(o.ImageRef) {
		var ret ImageRef
		return ret
	}
	return *o.ImageRef
}

// GetImageRefOk returns a tuple with the ImageRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetImageRefOk() (*ImageRef, bool) {
	if o == nil || IsNil(o.ImageRef) {
		return nil, false
	}
	return o.ImageRef, true
}

// HasImageRef returns a boolean if a field has been set.
func (o *Machine) HasImageRef() bool {
	if o != nil && !IsNil(o.ImageRef) {
		return true
	}

	return false
}

// SetImageRef gets a reference to the given ImageRef and assigns it to the ImageRef field.
func (o *Machine) SetImageRef(v ImageRef) {
	o.ImageRef = &v
}

// GetIncompleteConfig returns the IncompleteConfig field value if set, zero value otherwise.
func (o *Machine) GetIncompleteConfig() FlyMachineConfig {
	if o == nil || IsNil(o.IncompleteConfig) {
		var ret FlyMachineConfig
		return ret
	}
	return *o.IncompleteConfig
}

// GetIncompleteConfigOk returns a tuple with the IncompleteConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetIncompleteConfigOk() (*FlyMachineConfig, bool) {
	if o == nil || IsNil(o.IncompleteConfig) {
		return nil, false
	}
	return o.IncompleteConfig, true
}

// HasIncompleteConfig returns a boolean if a field has been set.
func (o *Machine) HasIncompleteConfig() bool {
	if o != nil && !IsNil(o.IncompleteConfig) {
		return true
	}

	return false
}

// SetIncompleteConfig gets a reference to the given FlyMachineConfig and assigns it to the IncompleteConfig field.
func (o *Machine) SetIncompleteConfig(v FlyMachineConfig) {
	o.IncompleteConfig = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *Machine) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *Machine) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *Machine) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Machine) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Machine) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Machine) SetName(v string) {
	o.Name = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Machine) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Machine) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Machine) SetNonce(v string) {
	o.Nonce = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *Machine) GetPrivateIp() string {
	if o == nil || IsNil(o.PrivateIp) {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetPrivateIpOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *Machine) HasPrivateIp() bool {
	if o != nil && !IsNil(o.PrivateIp) {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *Machine) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Machine) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Machine) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Machine) SetRegion(v string) {
	o.Region = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Machine) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Machine) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Machine) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Machine) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Machine) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Machine) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Machine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Machine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checks) {
		toSerialize["checks"] = o.Checks
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.HostStatus) {
		toSerialize["host_status"] = o.HostStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageRef) {
		toSerialize["image_ref"] = o.ImageRef
	}
	if !IsNil(o.IncompleteConfig) {
		toSerialize["incomplete_config"] = o.IncompleteConfig
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.PrivateIp) {
		toSerialize["private_ip"] = o.PrivateIp
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableMachine struct {
	value *Machine
	isSet bool
}

func (v NullableMachine) Get() *Machine {
	return v.value
}

func (v *NullableMachine) Set(val *Machine) {
	v.value = val
	v.isSet = true
}

func (v NullableMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachine(val *Machine) *NullableMachine {
	return &NullableMachine{value: val, isSet: true}
}

func (v NullableMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


