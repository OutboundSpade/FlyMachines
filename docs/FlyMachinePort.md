# FlyMachinePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPort** | Pointer to **int32** |  | [optional] 
**ForceHttps** | Pointer to **bool** |  | [optional] 
**Handlers** | Pointer to **[]string** |  | [optional] 
**HttpOptions** | Pointer to [**FlyHTTPOptions**](FlyHTTPOptions.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProxyProtoOptions** | Pointer to [**FlyProxyProtoOptions**](FlyProxyProtoOptions.md) |  | [optional] 
**StartPort** | Pointer to **int32** |  | [optional] 
**TlsOptions** | Pointer to [**FlyTLSOptions**](FlyTLSOptions.md) |  | [optional] 

## Methods

### NewFlyMachinePort

`func NewFlyMachinePort() *FlyMachinePort`

NewFlyMachinePort instantiates a new FlyMachinePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachinePortWithDefaults

`func NewFlyMachinePortWithDefaults() *FlyMachinePort`

NewFlyMachinePortWithDefaults instantiates a new FlyMachinePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPort

`func (o *FlyMachinePort) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *FlyMachinePort) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *FlyMachinePort) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *FlyMachinePort) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetForceHttps

`func (o *FlyMachinePort) GetForceHttps() bool`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *FlyMachinePort) GetForceHttpsOk() (*bool, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *FlyMachinePort) SetForceHttps(v bool)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *FlyMachinePort) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetHandlers

`func (o *FlyMachinePort) GetHandlers() []string`

GetHandlers returns the Handlers field if non-nil, zero value otherwise.

### GetHandlersOk

`func (o *FlyMachinePort) GetHandlersOk() (*[]string, bool)`

GetHandlersOk returns a tuple with the Handlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlers

`func (o *FlyMachinePort) SetHandlers(v []string)`

SetHandlers sets Handlers field to given value.

### HasHandlers

`func (o *FlyMachinePort) HasHandlers() bool`

HasHandlers returns a boolean if a field has been set.

### GetHttpOptions

`func (o *FlyMachinePort) GetHttpOptions() FlyHTTPOptions`

GetHttpOptions returns the HttpOptions field if non-nil, zero value otherwise.

### GetHttpOptionsOk

`func (o *FlyMachinePort) GetHttpOptionsOk() (*FlyHTTPOptions, bool)`

GetHttpOptionsOk returns a tuple with the HttpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOptions

`func (o *FlyMachinePort) SetHttpOptions(v FlyHTTPOptions)`

SetHttpOptions sets HttpOptions field to given value.

### HasHttpOptions

`func (o *FlyMachinePort) HasHttpOptions() bool`

HasHttpOptions returns a boolean if a field has been set.

### GetPort

`func (o *FlyMachinePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FlyMachinePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FlyMachinePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FlyMachinePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProxyProtoOptions

`func (o *FlyMachinePort) GetProxyProtoOptions() FlyProxyProtoOptions`

GetProxyProtoOptions returns the ProxyProtoOptions field if non-nil, zero value otherwise.

### GetProxyProtoOptionsOk

`func (o *FlyMachinePort) GetProxyProtoOptionsOk() (*FlyProxyProtoOptions, bool)`

GetProxyProtoOptionsOk returns a tuple with the ProxyProtoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtoOptions

`func (o *FlyMachinePort) SetProxyProtoOptions(v FlyProxyProtoOptions)`

SetProxyProtoOptions sets ProxyProtoOptions field to given value.

### HasProxyProtoOptions

`func (o *FlyMachinePort) HasProxyProtoOptions() bool`

HasProxyProtoOptions returns a boolean if a field has been set.

### GetStartPort

`func (o *FlyMachinePort) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *FlyMachinePort) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *FlyMachinePort) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *FlyMachinePort) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.

### GetTlsOptions

`func (o *FlyMachinePort) GetTlsOptions() FlyTLSOptions`

GetTlsOptions returns the TlsOptions field if non-nil, zero value otherwise.

### GetTlsOptionsOk

`func (o *FlyMachinePort) GetTlsOptionsOk() (*FlyTLSOptions, bool)`

GetTlsOptionsOk returns a tuple with the TlsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptions

`func (o *FlyMachinePort) SetTlsOptions(v FlyTLSOptions)`

SetTlsOptions sets TlsOptions field to given value.

### HasTlsOptions

`func (o *FlyMachinePort) HasTlsOptions() bool`

HasTlsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


