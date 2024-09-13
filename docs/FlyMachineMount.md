# FlyMachineMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddSizeGb** | Pointer to **int32** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**ExtendThresholdPercent** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SizeGb** | Pointer to **int32** |  | [optional] 
**SizeGbLimit** | Pointer to **int32** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewFlyMachineMount

`func NewFlyMachineMount() *FlyMachineMount`

NewFlyMachineMount instantiates a new FlyMachineMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineMountWithDefaults

`func NewFlyMachineMountWithDefaults() *FlyMachineMount`

NewFlyMachineMountWithDefaults instantiates a new FlyMachineMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddSizeGb

`func (o *FlyMachineMount) GetAddSizeGb() int32`

GetAddSizeGb returns the AddSizeGb field if non-nil, zero value otherwise.

### GetAddSizeGbOk

`func (o *FlyMachineMount) GetAddSizeGbOk() (*int32, bool)`

GetAddSizeGbOk returns a tuple with the AddSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSizeGb

`func (o *FlyMachineMount) SetAddSizeGb(v int32)`

SetAddSizeGb sets AddSizeGb field to given value.

### HasAddSizeGb

`func (o *FlyMachineMount) HasAddSizeGb() bool`

HasAddSizeGb returns a boolean if a field has been set.

### GetEncrypted

`func (o *FlyMachineMount) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FlyMachineMount) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FlyMachineMount) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *FlyMachineMount) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetExtendThresholdPercent

`func (o *FlyMachineMount) GetExtendThresholdPercent() int32`

GetExtendThresholdPercent returns the ExtendThresholdPercent field if non-nil, zero value otherwise.

### GetExtendThresholdPercentOk

`func (o *FlyMachineMount) GetExtendThresholdPercentOk() (*int32, bool)`

GetExtendThresholdPercentOk returns a tuple with the ExtendThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendThresholdPercent

`func (o *FlyMachineMount) SetExtendThresholdPercent(v int32)`

SetExtendThresholdPercent sets ExtendThresholdPercent field to given value.

### HasExtendThresholdPercent

`func (o *FlyMachineMount) HasExtendThresholdPercent() bool`

HasExtendThresholdPercent returns a boolean if a field has been set.

### GetName

`func (o *FlyMachineMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlyMachineMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlyMachineMount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlyMachineMount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *FlyMachineMount) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FlyMachineMount) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FlyMachineMount) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FlyMachineMount) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSizeGb

`func (o *FlyMachineMount) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *FlyMachineMount) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *FlyMachineMount) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *FlyMachineMount) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetSizeGbLimit

`func (o *FlyMachineMount) GetSizeGbLimit() int32`

GetSizeGbLimit returns the SizeGbLimit field if non-nil, zero value otherwise.

### GetSizeGbLimitOk

`func (o *FlyMachineMount) GetSizeGbLimitOk() (*int32, bool)`

GetSizeGbLimitOk returns a tuple with the SizeGbLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGbLimit

`func (o *FlyMachineMount) SetSizeGbLimit(v int32)`

SetSizeGbLimit sets SizeGbLimit field to given value.

### HasSizeGbLimit

`func (o *FlyMachineMount) HasSizeGbLimit() bool`

HasSizeGbLimit returns a boolean if a field has been set.

### GetVolume

`func (o *FlyMachineMount) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *FlyMachineMount) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *FlyMachineMount) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *FlyMachineMount) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


