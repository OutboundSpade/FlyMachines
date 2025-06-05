# PlacementRegionPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **int32** | Hint on the number of machines in this region can be created concurrently. Equal to the number of unique hosts selected for placement. | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewPlacementRegionPlacement

`func NewPlacementRegionPlacement() *PlacementRegionPlacement`

NewPlacementRegionPlacement instantiates a new PlacementRegionPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementRegionPlacementWithDefaults

`func NewPlacementRegionPlacementWithDefaults() *PlacementRegionPlacement`

NewPlacementRegionPlacementWithDefaults instantiates a new PlacementRegionPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *PlacementRegionPlacement) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *PlacementRegionPlacement) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *PlacementRegionPlacement) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *PlacementRegionPlacement) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetCount

`func (o *PlacementRegionPlacement) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PlacementRegionPlacement) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PlacementRegionPlacement) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PlacementRegionPlacement) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRegion

`func (o *PlacementRegionPlacement) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PlacementRegionPlacement) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PlacementRegionPlacement) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PlacementRegionPlacement) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


