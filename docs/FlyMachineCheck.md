# FlyMachineCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GracePeriod** | Pointer to [**FlyDuration**](FlyDuration.md) | The time to wait after a VM starts before checking its health | [optional] 
**Headers** | Pointer to [**[]FlyMachineHTTPHeader**](FlyMachineHTTPHeader.md) |  | [optional] 
**Interval** | Pointer to [**FlyDuration**](FlyDuration.md) | The time between connectivity checks | [optional] 
**Kind** | Pointer to **string** | Kind of the check (informational, readiness) | [optional] 
**Method** | Pointer to **string** | For http checks, the HTTP method to use to when making the request | [optional] 
**Path** | Pointer to **string** | For http checks, the path to send the request to | [optional] 
**Port** | Pointer to **int32** | The port to connect to, often the same as internal_port | [optional] 
**Protocol** | Pointer to **string** | For http checks, whether to use http or https | [optional] 
**Timeout** | Pointer to [**FlyDuration**](FlyDuration.md) | The maximum time a connection can take before being reported as failing its health check | [optional] 
**TlsServerName** | Pointer to **string** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] 
**TlsSkipVerify** | Pointer to **bool** | For http checks with https protocol, whether or not to verify the TLS certificate | [optional] 
**Type** | Pointer to **string** | tcp or http | [optional] 

## Methods

### NewFlyMachineCheck

`func NewFlyMachineCheck() *FlyMachineCheck`

NewFlyMachineCheck instantiates a new FlyMachineCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineCheckWithDefaults

`func NewFlyMachineCheckWithDefaults() *FlyMachineCheck`

NewFlyMachineCheckWithDefaults instantiates a new FlyMachineCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGracePeriod

`func (o *FlyMachineCheck) GetGracePeriod() FlyDuration`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *FlyMachineCheck) GetGracePeriodOk() (*FlyDuration, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *FlyMachineCheck) SetGracePeriod(v FlyDuration)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *FlyMachineCheck) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetHeaders

`func (o *FlyMachineCheck) GetHeaders() []FlyMachineHTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FlyMachineCheck) GetHeadersOk() (*[]FlyMachineHTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FlyMachineCheck) SetHeaders(v []FlyMachineHTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FlyMachineCheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetInterval

`func (o *FlyMachineCheck) GetInterval() FlyDuration`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *FlyMachineCheck) GetIntervalOk() (*FlyDuration, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *FlyMachineCheck) SetInterval(v FlyDuration)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *FlyMachineCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKind

`func (o *FlyMachineCheck) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlyMachineCheck) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlyMachineCheck) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FlyMachineCheck) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMethod

`func (o *FlyMachineCheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FlyMachineCheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FlyMachineCheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FlyMachineCheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *FlyMachineCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FlyMachineCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FlyMachineCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FlyMachineCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *FlyMachineCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FlyMachineCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FlyMachineCheck) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FlyMachineCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *FlyMachineCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FlyMachineCheck) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FlyMachineCheck) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FlyMachineCheck) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *FlyMachineCheck) GetTimeout() FlyDuration`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FlyMachineCheck) GetTimeoutOk() (*FlyDuration, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FlyMachineCheck) SetTimeout(v FlyDuration)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FlyMachineCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTlsServerName

`func (o *FlyMachineCheck) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *FlyMachineCheck) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *FlyMachineCheck) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *FlyMachineCheck) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *FlyMachineCheck) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *FlyMachineCheck) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *FlyMachineCheck) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *FlyMachineCheck) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.

### GetType

`func (o *FlyMachineCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlyMachineCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlyMachineCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlyMachineCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


