# ReadsGetCapacityPerRegionRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**GatewayAvailable** | Pointer to **bool** |  | [optional] 
**GeoRegion** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RequiresPaidPlan** | Pointer to **bool** |  | [optional] 

## Methods

### NewReadsGetCapacityPerRegionRow

`func NewReadsGetCapacityPerRegionRow() *ReadsGetCapacityPerRegionRow`

NewReadsGetCapacityPerRegionRow instantiates a new ReadsGetCapacityPerRegionRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadsGetCapacityPerRegionRowWithDefaults

`func NewReadsGetCapacityPerRegionRowWithDefaults() *ReadsGetCapacityPerRegionRow`

NewReadsGetCapacityPerRegionRowWithDefaults instantiates a new ReadsGetCapacityPerRegionRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *ReadsGetCapacityPerRegionRow) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ReadsGetCapacityPerRegionRow) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ReadsGetCapacityPerRegionRow) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ReadsGetCapacityPerRegionRow) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCode

`func (o *ReadsGetCapacityPerRegionRow) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReadsGetCapacityPerRegionRow) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReadsGetCapacityPerRegionRow) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReadsGetCapacityPerRegionRow) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetGatewayAvailable

`func (o *ReadsGetCapacityPerRegionRow) GetGatewayAvailable() bool`

GetGatewayAvailable returns the GatewayAvailable field if non-nil, zero value otherwise.

### GetGatewayAvailableOk

`func (o *ReadsGetCapacityPerRegionRow) GetGatewayAvailableOk() (*bool, bool)`

GetGatewayAvailableOk returns a tuple with the GatewayAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAvailable

`func (o *ReadsGetCapacityPerRegionRow) SetGatewayAvailable(v bool)`

SetGatewayAvailable sets GatewayAvailable field to given value.

### HasGatewayAvailable

`func (o *ReadsGetCapacityPerRegionRow) HasGatewayAvailable() bool`

HasGatewayAvailable returns a boolean if a field has been set.

### GetGeoRegion

`func (o *ReadsGetCapacityPerRegionRow) GetGeoRegion() string`

GetGeoRegion returns the GeoRegion field if non-nil, zero value otherwise.

### GetGeoRegionOk

`func (o *ReadsGetCapacityPerRegionRow) GetGeoRegionOk() (*string, bool)`

GetGeoRegionOk returns a tuple with the GeoRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRegion

`func (o *ReadsGetCapacityPerRegionRow) SetGeoRegion(v string)`

SetGeoRegion sets GeoRegion field to given value.

### HasGeoRegion

`func (o *ReadsGetCapacityPerRegionRow) HasGeoRegion() bool`

HasGeoRegion returns a boolean if a field has been set.

### GetLatitude

`func (o *ReadsGetCapacityPerRegionRow) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ReadsGetCapacityPerRegionRow) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ReadsGetCapacityPerRegionRow) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ReadsGetCapacityPerRegionRow) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *ReadsGetCapacityPerRegionRow) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ReadsGetCapacityPerRegionRow) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ReadsGetCapacityPerRegionRow) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ReadsGetCapacityPerRegionRow) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *ReadsGetCapacityPerRegionRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReadsGetCapacityPerRegionRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReadsGetCapacityPerRegionRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReadsGetCapacityPerRegionRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequiresPaidPlan

`func (o *ReadsGetCapacityPerRegionRow) GetRequiresPaidPlan() bool`

GetRequiresPaidPlan returns the RequiresPaidPlan field if non-nil, zero value otherwise.

### GetRequiresPaidPlanOk

`func (o *ReadsGetCapacityPerRegionRow) GetRequiresPaidPlanOk() (*bool, bool)`

GetRequiresPaidPlanOk returns a tuple with the RequiresPaidPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPaidPlan

`func (o *ReadsGetCapacityPerRegionRow) SetRequiresPaidPlan(v bool)`

SetRequiresPaidPlan sets RequiresPaidPlan field to given value.

### HasRequiresPaidPlan

`func (o *ReadsGetCapacityPerRegionRow) HasRequiresPaidPlan() bool`

HasRequiresPaidPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


