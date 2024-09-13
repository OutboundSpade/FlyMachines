# FlyStopConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signal** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to [**FlyDuration**](FlyDuration.md) |  | [optional] 

## Methods

### NewFlyStopConfig

`func NewFlyStopConfig() *FlyStopConfig`

NewFlyStopConfig instantiates a new FlyStopConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyStopConfigWithDefaults

`func NewFlyStopConfigWithDefaults() *FlyStopConfig`

NewFlyStopConfigWithDefaults instantiates a new FlyStopConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignal

`func (o *FlyStopConfig) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *FlyStopConfig) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *FlyStopConfig) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *FlyStopConfig) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTimeout

`func (o *FlyStopConfig) GetTimeout() FlyDuration`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FlyStopConfig) GetTimeoutOk() (*FlyDuration, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FlyStopConfig) SetTimeout(v FlyDuration)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FlyStopConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


