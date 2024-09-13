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

// checks if the FlyDNSConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlyDNSConfig{}

// FlyDNSConfig struct for FlyDNSConfig
type FlyDNSConfig struct {
	DnsForwardRules []FlyDnsForwardRule `json:"dns_forward_rules,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	HostnameFqdn *string `json:"hostname_fqdn,omitempty"`
	Nameservers []string `json:"nameservers,omitempty"`
	Options []FlyDnsOption `json:"options,omitempty"`
	Searches []string `json:"searches,omitempty"`
	SkipRegistration *bool `json:"skip_registration,omitempty"`
}

// NewFlyDNSConfig instantiates a new FlyDNSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlyDNSConfig() *FlyDNSConfig {
	this := FlyDNSConfig{}
	return &this
}

// NewFlyDNSConfigWithDefaults instantiates a new FlyDNSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlyDNSConfigWithDefaults() *FlyDNSConfig {
	this := FlyDNSConfig{}
	return &this
}

// GetDnsForwardRules returns the DnsForwardRules field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetDnsForwardRules() []FlyDnsForwardRule {
	if o == nil || IsNil(o.DnsForwardRules) {
		var ret []FlyDnsForwardRule
		return ret
	}
	return o.DnsForwardRules
}

// GetDnsForwardRulesOk returns a tuple with the DnsForwardRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetDnsForwardRulesOk() ([]FlyDnsForwardRule, bool) {
	if o == nil || IsNil(o.DnsForwardRules) {
		return nil, false
	}
	return o.DnsForwardRules, true
}

// HasDnsForwardRules returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasDnsForwardRules() bool {
	if o != nil && !IsNil(o.DnsForwardRules) {
		return true
	}

	return false
}

// SetDnsForwardRules gets a reference to the given []FlyDnsForwardRule and assigns it to the DnsForwardRules field.
func (o *FlyDNSConfig) SetDnsForwardRules(v []FlyDnsForwardRule) {
	o.DnsForwardRules = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FlyDNSConfig) SetHostname(v string) {
	o.Hostname = &v
}

// GetHostnameFqdn returns the HostnameFqdn field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetHostnameFqdn() string {
	if o == nil || IsNil(o.HostnameFqdn) {
		var ret string
		return ret
	}
	return *o.HostnameFqdn
}

// GetHostnameFqdnOk returns a tuple with the HostnameFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetHostnameFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameFqdn) {
		return nil, false
	}
	return o.HostnameFqdn, true
}

// HasHostnameFqdn returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasHostnameFqdn() bool {
	if o != nil && !IsNil(o.HostnameFqdn) {
		return true
	}

	return false
}

// SetHostnameFqdn gets a reference to the given string and assigns it to the HostnameFqdn field.
func (o *FlyDNSConfig) SetHostnameFqdn(v string) {
	o.HostnameFqdn = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetNameservers() []string {
	if o == nil || IsNil(o.Nameservers) {
		var ret []string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *FlyDNSConfig) SetNameservers(v []string) {
	o.Nameservers = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetOptions() []FlyDnsOption {
	if o == nil || IsNil(o.Options) {
		var ret []FlyDnsOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetOptionsOk() ([]FlyDnsOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []FlyDnsOption and assigns it to the Options field.
func (o *FlyDNSConfig) SetOptions(v []FlyDnsOption) {
	o.Options = v
}

// GetSearches returns the Searches field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetSearches() []string {
	if o == nil || IsNil(o.Searches) {
		var ret []string
		return ret
	}
	return o.Searches
}

// GetSearchesOk returns a tuple with the Searches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetSearchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Searches) {
		return nil, false
	}
	return o.Searches, true
}

// HasSearches returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasSearches() bool {
	if o != nil && !IsNil(o.Searches) {
		return true
	}

	return false
}

// SetSearches gets a reference to the given []string and assigns it to the Searches field.
func (o *FlyDNSConfig) SetSearches(v []string) {
	o.Searches = v
}

// GetSkipRegistration returns the SkipRegistration field value if set, zero value otherwise.
func (o *FlyDNSConfig) GetSkipRegistration() bool {
	if o == nil || IsNil(o.SkipRegistration) {
		var ret bool
		return ret
	}
	return *o.SkipRegistration
}

// GetSkipRegistrationOk returns a tuple with the SkipRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlyDNSConfig) GetSkipRegistrationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipRegistration) {
		return nil, false
	}
	return o.SkipRegistration, true
}

// HasSkipRegistration returns a boolean if a field has been set.
func (o *FlyDNSConfig) HasSkipRegistration() bool {
	if o != nil && !IsNil(o.SkipRegistration) {
		return true
	}

	return false
}

// SetSkipRegistration gets a reference to the given bool and assigns it to the SkipRegistration field.
func (o *FlyDNSConfig) SetSkipRegistration(v bool) {
	o.SkipRegistration = &v
}

func (o FlyDNSConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlyDNSConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnsForwardRules) {
		toSerialize["dns_forward_rules"] = o.DnsForwardRules
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.HostnameFqdn) {
		toSerialize["hostname_fqdn"] = o.HostnameFqdn
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Searches) {
		toSerialize["searches"] = o.Searches
	}
	if !IsNil(o.SkipRegistration) {
		toSerialize["skip_registration"] = o.SkipRegistration
	}
	return toSerialize, nil
}

type NullableFlyDNSConfig struct {
	value *FlyDNSConfig
	isSet bool
}

func (v NullableFlyDNSConfig) Get() *FlyDNSConfig {
	return v.value
}

func (v *NullableFlyDNSConfig) Set(val *FlyDNSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFlyDNSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFlyDNSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlyDNSConfig(val *FlyDNSConfig) *NullableFlyDNSConfig {
	return &NullableFlyDNSConfig{value: val, isSet: true}
}

func (v NullableFlyDNSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlyDNSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

