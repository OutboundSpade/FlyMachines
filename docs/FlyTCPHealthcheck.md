# FlyTCPHealthcheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | The port to connect to, often the same as internal_port | [optional] 

## Methods

### NewFlyTCPHealthcheck

`func NewFlyTCPHealthcheck() *FlyTCPHealthcheck`

NewFlyTCPHealthcheck instantiates a new FlyTCPHealthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyTCPHealthcheckWithDefaults

`func NewFlyTCPHealthcheckWithDefaults() *FlyTCPHealthcheck`

NewFlyTCPHealthcheckWithDefaults instantiates a new FlyTCPHealthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *FlyTCPHealthcheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FlyTCPHealthcheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FlyTCPHealthcheck) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FlyTCPHealthcheck) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


