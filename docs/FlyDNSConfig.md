# FlyDNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsForwardRules** | Pointer to [**[]FlyDnsForwardRule**](FlyDnsForwardRule.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**HostnameFqdn** | Pointer to **string** |  | [optional] 
**Nameservers** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to [**[]FlyDnsOption**](FlyDnsOption.md) |  | [optional] 
**Searches** | Pointer to **[]string** |  | [optional] 
**SkipRegistration** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlyDNSConfig

`func NewFlyDNSConfig() *FlyDNSConfig`

NewFlyDNSConfig instantiates a new FlyDNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyDNSConfigWithDefaults

`func NewFlyDNSConfigWithDefaults() *FlyDNSConfig`

NewFlyDNSConfigWithDefaults instantiates a new FlyDNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsForwardRules

`func (o *FlyDNSConfig) GetDnsForwardRules() []FlyDnsForwardRule`

GetDnsForwardRules returns the DnsForwardRules field if non-nil, zero value otherwise.

### GetDnsForwardRulesOk

`func (o *FlyDNSConfig) GetDnsForwardRulesOk() (*[]FlyDnsForwardRule, bool)`

GetDnsForwardRulesOk returns a tuple with the DnsForwardRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsForwardRules

`func (o *FlyDNSConfig) SetDnsForwardRules(v []FlyDnsForwardRule)`

SetDnsForwardRules sets DnsForwardRules field to given value.

### HasDnsForwardRules

`func (o *FlyDNSConfig) HasDnsForwardRules() bool`

HasDnsForwardRules returns a boolean if a field has been set.

### GetHostname

`func (o *FlyDNSConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FlyDNSConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FlyDNSConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FlyDNSConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnameFqdn

`func (o *FlyDNSConfig) GetHostnameFqdn() string`

GetHostnameFqdn returns the HostnameFqdn field if non-nil, zero value otherwise.

### GetHostnameFqdnOk

`func (o *FlyDNSConfig) GetHostnameFqdnOk() (*string, bool)`

GetHostnameFqdnOk returns a tuple with the HostnameFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameFqdn

`func (o *FlyDNSConfig) SetHostnameFqdn(v string)`

SetHostnameFqdn sets HostnameFqdn field to given value.

### HasHostnameFqdn

`func (o *FlyDNSConfig) HasHostnameFqdn() bool`

HasHostnameFqdn returns a boolean if a field has been set.

### GetNameservers

`func (o *FlyDNSConfig) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *FlyDNSConfig) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *FlyDNSConfig) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *FlyDNSConfig) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOptions

`func (o *FlyDNSConfig) GetOptions() []FlyDnsOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FlyDNSConfig) GetOptionsOk() (*[]FlyDnsOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FlyDNSConfig) SetOptions(v []FlyDnsOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FlyDNSConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSearches

`func (o *FlyDNSConfig) GetSearches() []string`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *FlyDNSConfig) GetSearchesOk() (*[]string, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *FlyDNSConfig) SetSearches(v []string)`

SetSearches sets Searches field to given value.

### HasSearches

`func (o *FlyDNSConfig) HasSearches() bool`

HasSearches returns a boolean if a field has been set.

### GetSkipRegistration

`func (o *FlyDNSConfig) GetSkipRegistration() bool`

GetSkipRegistration returns the SkipRegistration field if non-nil, zero value otherwise.

### GetSkipRegistrationOk

`func (o *FlyDNSConfig) GetSkipRegistrationOk() (*bool, bool)`

GetSkipRegistrationOk returns a tuple with the SkipRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRegistration

`func (o *FlyDNSConfig) SetSkipRegistration(v bool)`

SetSkipRegistration sets SkipRegistration field to given value.

### HasSkipRegistration

`func (o *FlyDNSConfig) HasSkipRegistration() bool`

HasSkipRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


