# UpdateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBackupEnabled** | Pointer to **bool** |  | [optional] 
**SnapshotRetention** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateVolumeRequest

`func NewUpdateVolumeRequest() *UpdateVolumeRequest`

NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeRequestWithDefaults

`func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest`

NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBackupEnabled

`func (o *UpdateVolumeRequest) GetAutoBackupEnabled() bool`

GetAutoBackupEnabled returns the AutoBackupEnabled field if non-nil, zero value otherwise.

### GetAutoBackupEnabledOk

`func (o *UpdateVolumeRequest) GetAutoBackupEnabledOk() (*bool, bool)`

GetAutoBackupEnabledOk returns a tuple with the AutoBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBackupEnabled

`func (o *UpdateVolumeRequest) SetAutoBackupEnabled(v bool)`

SetAutoBackupEnabled sets AutoBackupEnabled field to given value.

### HasAutoBackupEnabled

`func (o *UpdateVolumeRequest) HasAutoBackupEnabled() bool`

HasAutoBackupEnabled returns a boolean if a field has been set.

### GetSnapshotRetention

`func (o *UpdateVolumeRequest) GetSnapshotRetention() int32`

GetSnapshotRetention returns the SnapshotRetention field if non-nil, zero value otherwise.

### GetSnapshotRetentionOk

`func (o *UpdateVolumeRequest) GetSnapshotRetentionOk() (*int32, bool)`

GetSnapshotRetentionOk returns a tuple with the SnapshotRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetention

`func (o *UpdateVolumeRequest) SetSnapshotRetention(v int32)`

SetSnapshotRetention sets SnapshotRetention field to given value.

### HasSnapshotRetention

`func (o *UpdateVolumeRequest) HasSnapshotRetention() bool`

HasSnapshotRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


