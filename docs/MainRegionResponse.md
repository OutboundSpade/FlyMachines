# MainRegionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]ReadsGetCapacityPerRegionRow**](ReadsGetCapacityPerRegionRow.md) |  | [optional] 

## Methods

### NewMainRegionResponse

`func NewMainRegionResponse() *MainRegionResponse`

NewMainRegionResponse instantiates a new MainRegionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainRegionResponseWithDefaults

`func NewMainRegionResponseWithDefaults() *MainRegionResponse`

NewMainRegionResponseWithDefaults instantiates a new MainRegionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *MainRegionResponse) GetRegions() []ReadsGetCapacityPerRegionRow`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *MainRegionResponse) GetRegionsOk() (*[]ReadsGetCapacityPerRegionRow, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *MainRegionResponse) SetRegions(v []ReadsGetCapacityPerRegionRow)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *MainRegionResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


