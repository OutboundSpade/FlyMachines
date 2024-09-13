# FlyHTTPOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compress** | Pointer to **bool** |  | [optional] 
**H2Backend** | Pointer to **bool** |  | [optional] 
**HeadersReadTimeout** | Pointer to **int32** |  | [optional] 
**IdleTimeout** | Pointer to **int32** |  | [optional] 
**Response** | Pointer to [**FlyHTTPResponseOptions**](FlyHTTPResponseOptions.md) |  | [optional] 

## Methods

### NewFlyHTTPOptions

`func NewFlyHTTPOptions() *FlyHTTPOptions`

NewFlyHTTPOptions instantiates a new FlyHTTPOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyHTTPOptionsWithDefaults

`func NewFlyHTTPOptionsWithDefaults() *FlyHTTPOptions`

NewFlyHTTPOptionsWithDefaults instantiates a new FlyHTTPOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompress

`func (o *FlyHTTPOptions) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *FlyHTTPOptions) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *FlyHTTPOptions) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *FlyHTTPOptions) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetH2Backend

`func (o *FlyHTTPOptions) GetH2Backend() bool`

GetH2Backend returns the H2Backend field if non-nil, zero value otherwise.

### GetH2BackendOk

`func (o *FlyHTTPOptions) GetH2BackendOk() (*bool, bool)`

GetH2BackendOk returns a tuple with the H2Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH2Backend

`func (o *FlyHTTPOptions) SetH2Backend(v bool)`

SetH2Backend sets H2Backend field to given value.

### HasH2Backend

`func (o *FlyHTTPOptions) HasH2Backend() bool`

HasH2Backend returns a boolean if a field has been set.

### GetHeadersReadTimeout

`func (o *FlyHTTPOptions) GetHeadersReadTimeout() int32`

GetHeadersReadTimeout returns the HeadersReadTimeout field if non-nil, zero value otherwise.

### GetHeadersReadTimeoutOk

`func (o *FlyHTTPOptions) GetHeadersReadTimeoutOk() (*int32, bool)`

GetHeadersReadTimeoutOk returns a tuple with the HeadersReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersReadTimeout

`func (o *FlyHTTPOptions) SetHeadersReadTimeout(v int32)`

SetHeadersReadTimeout sets HeadersReadTimeout field to given value.

### HasHeadersReadTimeout

`func (o *FlyHTTPOptions) HasHeadersReadTimeout() bool`

HasHeadersReadTimeout returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *FlyHTTPOptions) GetIdleTimeout() int32`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *FlyHTTPOptions) GetIdleTimeoutOk() (*int32, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *FlyHTTPOptions) SetIdleTimeout(v int32)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *FlyHTTPOptions) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetResponse

`func (o *FlyHTTPOptions) GetResponse() FlyHTTPResponseOptions`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FlyHTTPOptions) GetResponseOk() (*FlyHTTPResponseOptions, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FlyHTTPOptions) SetResponse(v FlyHTTPResponseOptions)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FlyHTTPOptions) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


