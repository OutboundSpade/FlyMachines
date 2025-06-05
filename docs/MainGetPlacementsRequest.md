# MainGetPlacementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**FlyMachineGuest**](FlyMachineGuest.md) | Resource requirements for the Machine to simulate. Defaults to a performance-1x machine | [optional] 
**Count** | Pointer to **int64** | Number of machines to simulate placement. Defaults to 0, which returns the org-specific limit for each region. | [optional] 
**OrgSlug** | **string** |  | 
**Region** | Pointer to **string** | Region expression for placement as a comma-delimited set of regions or aliases. Defaults to \&quot;[region],any\&quot;, to prefer the API endpoint&#39;s local region with any other region as fallback. | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 
**VolumeSizeBytes** | Pointer to **int64** |  | [optional] 
**Weights** | Pointer to **map[string]int64** | Optional weights to override default placement preferences. | [optional] 

## Methods

### NewMainGetPlacementsRequest

`func NewMainGetPlacementsRequest(orgSlug string, ) *MainGetPlacementsRequest`

NewMainGetPlacementsRequest instantiates a new MainGetPlacementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainGetPlacementsRequestWithDefaults

`func NewMainGetPlacementsRequestWithDefaults() *MainGetPlacementsRequest`

NewMainGetPlacementsRequestWithDefaults instantiates a new MainGetPlacementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *MainGetPlacementsRequest) GetCompute() FlyMachineGuest`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *MainGetPlacementsRequest) GetComputeOk() (*FlyMachineGuest, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *MainGetPlacementsRequest) SetCompute(v FlyMachineGuest)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *MainGetPlacementsRequest) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetCount

`func (o *MainGetPlacementsRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MainGetPlacementsRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MainGetPlacementsRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MainGetPlacementsRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrgSlug

`func (o *MainGetPlacementsRequest) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *MainGetPlacementsRequest) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *MainGetPlacementsRequest) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetRegion

`func (o *MainGetPlacementsRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MainGetPlacementsRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MainGetPlacementsRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MainGetPlacementsRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVolumeName

`func (o *MainGetPlacementsRequest) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *MainGetPlacementsRequest) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *MainGetPlacementsRequest) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *MainGetPlacementsRequest) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeSizeBytes

`func (o *MainGetPlacementsRequest) GetVolumeSizeBytes() int64`

GetVolumeSizeBytes returns the VolumeSizeBytes field if non-nil, zero value otherwise.

### GetVolumeSizeBytesOk

`func (o *MainGetPlacementsRequest) GetVolumeSizeBytesOk() (*int64, bool)`

GetVolumeSizeBytesOk returns a tuple with the VolumeSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeBytes

`func (o *MainGetPlacementsRequest) SetVolumeSizeBytes(v int64)`

SetVolumeSizeBytes sets VolumeSizeBytes field to given value.

### HasVolumeSizeBytes

`func (o *MainGetPlacementsRequest) HasVolumeSizeBytes() bool`

HasVolumeSizeBytes returns a boolean if a field has been set.

### GetWeights

`func (o *MainGetPlacementsRequest) GetWeights() map[string]int64`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *MainGetPlacementsRequest) GetWeightsOk() (*map[string]int64, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *MainGetPlacementsRequest) SetWeights(v map[string]int64)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *MainGetPlacementsRequest) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


