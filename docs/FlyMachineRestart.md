# FlyMachineRestart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuBidPrice** | Pointer to **float32** | GPU bid price for spot Machines. | [optional] 
**MaxRetries** | Pointer to **int32** | When policy is on-failure, the maximum number of times to attempt to restart the Machine before letting it stop. | [optional] 
**Policy** | Pointer to **string** | * no - Never try to restart a Machine automatically when its main process exits, whether that’s on purpose or on a crash. * always - Always restart a Machine automatically and never let it enter a stopped state, even when the main process exits cleanly. * on-failure - Try up to MaxRetries times to automatically restart the Machine if it exits with a non-zero exit code. Default when no explicit policy is set, and for Machines with schedules. * spot-price - Starts the Machine only when there is capacity and the spot price is less than or equal to the bid price. | [optional] 

## Methods

### NewFlyMachineRestart

`func NewFlyMachineRestart() *FlyMachineRestart`

NewFlyMachineRestart instantiates a new FlyMachineRestart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineRestartWithDefaults

`func NewFlyMachineRestartWithDefaults() *FlyMachineRestart`

NewFlyMachineRestartWithDefaults instantiates a new FlyMachineRestart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuBidPrice

`func (o *FlyMachineRestart) GetGpuBidPrice() float32`

GetGpuBidPrice returns the GpuBidPrice field if non-nil, zero value otherwise.

### GetGpuBidPriceOk

`func (o *FlyMachineRestart) GetGpuBidPriceOk() (*float32, bool)`

GetGpuBidPriceOk returns a tuple with the GpuBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuBidPrice

`func (o *FlyMachineRestart) SetGpuBidPrice(v float32)`

SetGpuBidPrice sets GpuBidPrice field to given value.

### HasGpuBidPrice

`func (o *FlyMachineRestart) HasGpuBidPrice() bool`

HasGpuBidPrice returns a boolean if a field has been set.

### GetMaxRetries

`func (o *FlyMachineRestart) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *FlyMachineRestart) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *FlyMachineRestart) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *FlyMachineRestart) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetPolicy

`func (o *FlyMachineRestart) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *FlyMachineRestart) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *FlyMachineRestart) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *FlyMachineRestart) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


