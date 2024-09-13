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

// checks if the FlyMachineCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyMachineCheck{}

// FlyMachineCheck An optional object that defines one or more named checks. The key for each check is the check name.
type FlyMachineCheck struct {
	// The time to wait after a VM starts before checking its health
	GracePeriod *FlyDuration `json:"grace_period,omitempty"`
	Headers []FlyMachineHTTPHeader `json:"headers,omitempty"`
	// The time between connectivity checks
	Interval *FlyDuration `json:"interval,omitempty"`
	// Kind of the check (informational, readiness)
	Kind *string `json:"kind,omitempty"`
	// For http checks, the HTTP method to use to when making the request
	Method *string `json:"method,omitempty"`
	// For http checks, the path to send the request to
	Path *string `json:"path,omitempty"`
	// The port to connect to, often the same as internal_port
	Port *int32 `json:"port,omitempty"`
	// For http checks, whether to use http or https
	Protocol *string `json:"protocol,omitempty"`
	// The maximum time a connection can take before being reported as failing its health check
	Timeout *FlyDuration `json:"timeout,omitempty"`
	// If the protocol is https, the hostname to use for TLS certificate validation
	TlsServerName *string `json:"tls_server_name,omitempty"`
	// For http checks with https protocol, whether or not to verify the TLS certificate
	TlsSkipVerify *bool `json:"tls_skip_verify,omitempty"`
	// tcp or http
	Type *string `json:"type,omitempty"`
}

// NewFlyMachineCheck instantiates a new FlyMachineCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyMachineCheck() *FlyMachineCheck {
	this := FlyMachineCheck{}
	return &this
}

// NewFlyMachineCheckWithDefaults instantiates a new FlyMachineCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyMachineCheckWithDefaults() *FlyMachineCheck {
	this := FlyMachineCheck{}
	return &this
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetGracePeriod() FlyDuration {
	if o == nil || IsNil(o.GracePeriod) {
		var ret FlyDuration
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetGracePeriodOk() (*FlyDuration, bool) {
	if o == nil || IsNil(o.GracePeriod) {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasGracePeriod() bool {
	if o != nil && !IsNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given FlyDuration and assigns it to the GracePeriod field.
func (o *FlyMachineCheck) SetGracePeriod(v FlyDuration) {
	o.GracePeriod = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetHeaders() []FlyMachineHTTPHeader {
	if o == nil || IsNil(o.Headers) {
		var ret []FlyMachineHTTPHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetHeadersOk() ([]FlyMachineHTTPHeader, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []FlyMachineHTTPHeader and assigns it to the Headers field.
func (o *FlyMachineCheck) SetHeaders(v []FlyMachineHTTPHeader) {
	o.Headers = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetInterval() FlyDuration {
	if o == nil || IsNil(o.Interval) {
		var ret FlyDuration
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetIntervalOk() (*FlyDuration, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given FlyDuration and assigns it to the Interval field.
func (o *FlyMachineCheck) SetInterval(v FlyDuration) {
	o.Interval = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *FlyMachineCheck) SetKind(v string) {
	o.Kind = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *FlyMachineCheck) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FlyMachineCheck) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *FlyMachineCheck) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FlyMachineCheck) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetTimeout() FlyDuration {
	if o == nil || IsNil(o.Timeout) {
		var ret FlyDuration
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetTimeoutOk() (*FlyDuration, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given FlyDuration and assigns it to the Timeout field.
func (o *FlyMachineCheck) SetTimeout(v FlyDuration) {
	o.Timeout = &v
}

// GetTlsServerName returns the TlsServerName field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetTlsServerName() string {
	if o == nil || IsNil(o.TlsServerName) {
		var ret string
		return ret
	}
	return *o.TlsServerName
}

// GetTlsServerNameOk returns a tuple with the TlsServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetTlsServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.TlsServerName) {
		return nil, false
	}
	return o.TlsServerName, true
}

// HasTlsServerName returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasTlsServerName() bool {
	if o != nil && !IsNil(o.TlsServerName) {
		return true
	}

	return false
}

// SetTlsServerName gets a reference to the given string and assigns it to the TlsServerName field.
func (o *FlyMachineCheck) SetTlsServerName(v string) {
	o.TlsServerName = &v
}

// GetTlsSkipVerify returns the TlsSkipVerify field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetTlsSkipVerify() bool {
	if o == nil || IsNil(o.TlsSkipVerify) {
		var ret bool
		return ret
	}
	return *o.TlsSkipVerify
}

// GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetTlsSkipVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsSkipVerify) {
		return nil, false
	}
	return o.TlsSkipVerify, true
}

// HasTlsSkipVerify returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasTlsSkipVerify() bool {
	if o != nil && !IsNil(o.TlsSkipVerify) {
		return true
	}

	return false
}

// SetTlsSkipVerify gets a reference to the given bool and assigns it to the TlsSkipVerify field.
func (o *FlyMachineCheck) SetTlsSkipVerify(v bool) {
	o.TlsSkipVerify = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlyMachineCheck) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyMachineCheck) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlyMachineCheck) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlyMachineCheck) SetType(v string) {
	o.Type = &v
}

func (o FlyMachineCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyMachineCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GracePeriod) {
		toSerialize["grace_period"] = o.GracePeriod
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.TlsServerName) {
		toSerialize["tls_server_name"] = o.TlsServerName
	}
	if !IsNil(o.TlsSkipVerify) {
		toSerialize["tls_skip_verify"] = o.TlsSkipVerify
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFlyMachineCheck struct {
	value *FlyMachineCheck
	isSet bool
}

func (v NullableFlyMachineCheck) Get() *FlyMachineCheck {
	return v.value
}

func (v *NullableFlyMachineCheck) Set(val *FlyMachineCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyMachineCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyMachineCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyMachineCheck(val *FlyMachineCheck) *NullableFlyMachineCheck {
	return &NullableFlyMachineCheck{value: val, isSet: true}
}

func (v NullableFlyMachineCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyMachineCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


