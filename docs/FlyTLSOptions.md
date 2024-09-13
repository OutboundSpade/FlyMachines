# FlyTLSOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpn** | Pointer to **[]string** |  | [optional] 
**DefaultSelfSigned** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFlyTLSOptions

`func NewFlyTLSOptions() *FlyTLSOptions`

NewFlyTLSOptions instantiates a new FlyTLSOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyTLSOptionsWithDefaults

`func NewFlyTLSOptionsWithDefaults() *FlyTLSOptions`

NewFlyTLSOptionsWithDefaults instantiates a new FlyTLSOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpn

`func (o *FlyTLSOptions) GetAlpn() []string`

GetAlpn returns the Alpn field if non-nil, zero value otherwise.

### GetAlpnOk

`func (o *FlyTLSOptions) GetAlpnOk() (*[]string, bool)`

GetAlpnOk returns a tuple with the Alpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpn

`func (o *FlyTLSOptions) SetAlpn(v []string)`

SetAlpn sets Alpn field to given value.

### HasAlpn

`func (o *FlyTLSOptions) HasAlpn() bool`

HasAlpn returns a boolean if a field has been set.

### GetDefaultSelfSigned

`func (o *FlyTLSOptions) GetDefaultSelfSigned() bool`

GetDefaultSelfSigned returns the DefaultSelfSigned field if non-nil, zero value otherwise.

### GetDefaultSelfSignedOk

`func (o *FlyTLSOptions) GetDefaultSelfSignedOk() (*bool, bool)`

GetDefaultSelfSignedOk returns a tuple with the DefaultSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSelfSigned

`func (o *FlyTLSOptions) SetDefaultSelfSigned(v bool)`

SetDefaultSelfSigned sets DefaultSelfSigned field to given value.

### HasDefaultSelfSigned

`func (o *FlyTLSOptions) HasDefaultSelfSigned() bool`

HasDefaultSelfSigned returns a boolean if a field has been set.

### GetVersions

`func (o *FlyTLSOptions) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *FlyTLSOptions) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *FlyTLSOptions) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *FlyTLSOptions) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


