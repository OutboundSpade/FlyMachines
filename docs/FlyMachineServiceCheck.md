# FlyMachineServiceCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GracePeriod** | Pointer to [**FlyDuration**](FlyDuration.md) | The time to wait after a VM starts before checking its health | [optional] 
**Headers** | Pointer to [**[]FlyMachineHTTPHeader**](FlyMachineHTTPHeader.md) |  | [optional] 
**Interval** | Pointer to [**FlyDuration**](FlyDuration.md) | The time between connectivity checks | [optional] 
**Method** | Pointer to **string** | For http checks, the HTTP method to use to when making the request | [optional] 
**Path** | Pointer to **string** | For http checks, the path to send the request to | [optional] 
**Port** | Pointer to **int32** | The port to connect to, often the same as internal_port | [optional] 
**Protocol** | Pointer to **string** | For http checks, whether to use http or https | [optional] 
**Timeout** | Pointer to [**FlyDuration**](FlyDuration.md) | The maximum time a connection can take before being reported as failing its health check | [optional] 
**TlsServerName** | Pointer to **string** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] 
**TlsSkipVerify** | Pointer to **bool** | For http checks with https protocol, whether or not to verify the TLS certificate | [optional] 
**Type** | Pointer to **string** | tcp or http | [optional] 

## Methods

### NewFlyMachineServiceCheck

`func NewFlyMachineServiceCheck() *FlyMachineServiceCheck`

NewFlyMachineServiceCheck instantiates a new FlyMachineServiceCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineServiceCheckWithDefaults

`func NewFlyMachineServiceCheckWithDefaults() *FlyMachineServiceCheck`

NewFlyMachineServiceCheckWithDefaults instantiates a new FlyMachineServiceCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGracePeriod

`func (o *FlyMachineServiceCheck) GetGracePeriod() FlyDuration`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *FlyMachineServiceCheck) GetGracePeriodOk() (*FlyDuration, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *FlyMachineServiceCheck) SetGracePeriod(v FlyDuration)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *FlyMachineServiceCheck) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetHeaders

`func (o *FlyMachineServiceCheck) GetHeaders() []FlyMachineHTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FlyMachineServiceCheck) GetHeadersOk() (*[]FlyMachineHTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FlyMachineServiceCheck) SetHeaders(v []FlyMachineHTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FlyMachineServiceCheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetInterval

`func (o *FlyMachineServiceCheck) GetInterval() FlyDuration`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *FlyMachineServiceCheck) GetIntervalOk() (*FlyDuration, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *FlyMachineServiceCheck) SetInterval(v FlyDuration)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *FlyMachineServiceCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMethod

`func (o *FlyMachineServiceCheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FlyMachineServiceCheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FlyMachineServiceCheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FlyMachineServiceCheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *FlyMachineServiceCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FlyMachineServiceCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FlyMachineServiceCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FlyMachineServiceCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *FlyMachineServiceCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FlyMachineServiceCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FlyMachineServiceCheck) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FlyMachineServiceCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FlyMachineServiceCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FlyMachineServiceCheck) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FlyMachineServiceCheck) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FlyMachineServiceCheck) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *FlyMachineServiceCheck) GetTimeout() FlyDuration`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FlyMachineServiceCheck) GetTimeoutOk() (*FlyDuration, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FlyMachineServiceCheck) SetTimeout(v FlyDuration)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FlyMachineServiceCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTlsServerName

`func (o *FlyMachineServiceCheck) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *FlyMachineServiceCheck) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *FlyMachineServiceCheck) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *FlyMachineServiceCheck) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *FlyMachineServiceCheck) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *FlyMachineServiceCheck) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *FlyMachineServiceCheck) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *FlyMachineServiceCheck) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.

### GetType

`func (o *FlyMachineServiceCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlyMachineServiceCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlyMachineServiceCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlyMachineServiceCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


