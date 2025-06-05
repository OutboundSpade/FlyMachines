# FlyHTTPHealthcheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]FlyMachineHTTPHeader**](FlyMachineHTTPHeader.md) | Additional headers to send with the request | [optional] 
**Method** | Pointer to **string** | The HTTP method to use to when making the request | [optional] 
**Path** | Pointer to **string** | The path to send the request to | [optional] 
**Port** | Pointer to **int64** | The port to connect to, often the same as internal_port | [optional] 
**Scheme** | Pointer to [**FlyContainerHealthcheckScheme**](FlyContainerHealthcheckScheme.md) | Whether to use http or https | [optional] 
**TlsServerName** | Pointer to **string** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] 
**TlsSkipVerify** | Pointer to **bool** | If the protocol is https, whether or not to verify the TLS certificate | [optional] 

## Methods

### NewFlyHTTPHealthcheck

`func NewFlyHTTPHealthcheck() *FlyHTTPHealthcheck`

NewFlyHTTPHealthcheck instantiates a new FlyHTTPHealthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyHTTPHealthcheckWithDefaults

`func NewFlyHTTPHealthcheckWithDefaults() *FlyHTTPHealthcheck`

NewFlyHTTPHealthcheckWithDefaults instantiates a new FlyHTTPHealthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *FlyHTTPHealthcheck) GetHeaders() []FlyMachineHTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FlyHTTPHealthcheck) GetHeadersOk() (*[]FlyMachineHTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FlyHTTPHealthcheck) SetHeaders(v []FlyMachineHTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FlyHTTPHealthcheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *FlyHTTPHealthcheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FlyHTTPHealthcheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FlyHTTPHealthcheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FlyHTTPHealthcheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *FlyHTTPHealthcheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FlyHTTPHealthcheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FlyHTTPHealthcheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FlyHTTPHealthcheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *FlyHTTPHealthcheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FlyHTTPHealthcheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FlyHTTPHealthcheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *FlyHTTPHealthcheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetScheme

`func (o *FlyHTTPHealthcheck) GetScheme() FlyContainerHealthcheckScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *FlyHTTPHealthcheck) GetSchemeOk() (*FlyContainerHealthcheckScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *FlyHTTPHealthcheck) SetScheme(v FlyContainerHealthcheckScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *FlyHTTPHealthcheck) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetTlsServerName

`func (o *FlyHTTPHealthcheck) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *FlyHTTPHealthcheck) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *FlyHTTPHealthcheck) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *FlyHTTPHealthcheck) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *FlyHTTPHealthcheck) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *FlyHTTPHealthcheck) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *FlyHTTPHealthcheck) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *FlyHTTPHealthcheck) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


