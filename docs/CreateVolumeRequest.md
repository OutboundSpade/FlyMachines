# CreateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**FlyMachineGuest**](FlyMachineGuest.md) |  | [optional] 
**ComputeImage** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**Fstype** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RequireUniqueZone** | Pointer to **bool** |  | [optional] 
**SizeGb** | Pointer to **int32** |  | [optional] 
**SnapshotId** | Pointer to **string** | restore from snapshot | [optional] 
**SnapshotRetention** | Pointer to **int32** |  | [optional] 
**SourceVolumeId** | Pointer to **string** | fork from remote volume | [optional] 

## Methods

### NewCreateVolumeRequest

`func NewCreateVolumeRequest() *CreateVolumeRequest`

NewCreateVolumeRequest instantiates a new CreateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeRequestWithDefaults

`func NewCreateVolumeRequestWithDefaults() *CreateVolumeRequest`

NewCreateVolumeRequestWithDefaults instantiates a new CreateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *CreateVolumeRequest) GetCompute() FlyMachineGuest`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *CreateVolumeRequest) GetComputeOk() (*FlyMachineGuest, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *CreateVolumeRequest) SetCompute(v FlyMachineGuest)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *CreateVolumeRequest) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetComputeImage

`func (o *CreateVolumeRequest) GetComputeImage() string`

GetComputeImage returns the ComputeImage field if non-nil, zero value otherwise.

### GetComputeImageOk

`func (o *CreateVolumeRequest) GetComputeImageOk() (*string, bool)`

GetComputeImageOk returns a tuple with the ComputeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeImage

`func (o *CreateVolumeRequest) SetComputeImage(v string)`

SetComputeImage sets ComputeImage field to given value.

### HasComputeImage

`func (o *CreateVolumeRequest) HasComputeImage() bool`

HasComputeImage returns a boolean if a field has been set.

### GetEncrypted

`func (o *CreateVolumeRequest) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *CreateVolumeRequest) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *CreateVolumeRequest) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *CreateVolumeRequest) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFstype

`func (o *CreateVolumeRequest) GetFstype() string`

GetFstype returns the Fstype field if non-nil, zero value otherwise.

### GetFstypeOk

`func (o *CreateVolumeRequest) GetFstypeOk() (*string, bool)`

GetFstypeOk returns a tuple with the Fstype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstype

`func (o *CreateVolumeRequest) SetFstype(v string)`

SetFstype sets Fstype field to given value.

### HasFstype

`func (o *CreateVolumeRequest) HasFstype() bool`

HasFstype returns a boolean if a field has been set.

### GetName

`func (o *CreateVolumeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVolumeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CreateVolumeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateVolumeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateVolumeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateVolumeRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequireUniqueZone

`func (o *CreateVolumeRequest) GetRequireUniqueZone() bool`

GetRequireUniqueZone returns the RequireUniqueZone field if non-nil, zero value otherwise.

### GetRequireUniqueZoneOk

`func (o *CreateVolumeRequest) GetRequireUniqueZoneOk() (*bool, bool)`

GetRequireUniqueZoneOk returns a tuple with the RequireUniqueZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireUniqueZone

`func (o *CreateVolumeRequest) SetRequireUniqueZone(v bool)`

SetRequireUniqueZone sets RequireUniqueZone field to given value.

### HasRequireUniqueZone

`func (o *CreateVolumeRequest) HasRequireUniqueZone() bool`

HasRequireUniqueZone returns a boolean if a field has been set.

### GetSizeGb

`func (o *CreateVolumeRequest) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *CreateVolumeRequest) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *CreateVolumeRequest) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *CreateVolumeRequest) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetSnapshotId

`func (o *CreateVolumeRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreateVolumeRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CreateVolumeRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *CreateVolumeRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotRetention

`func (o *CreateVolumeRequest) GetSnapshotRetention() int32`

GetSnapshotRetention returns the SnapshotRetention field if non-nil, zero value otherwise.

### GetSnapshotRetentionOk

`func (o *CreateVolumeRequest) GetSnapshotRetentionOk() (*int32, bool)`

GetSnapshotRetentionOk returns a tuple with the SnapshotRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetention

`func (o *CreateVolumeRequest) SetSnapshotRetention(v int32)`

SetSnapshotRetention sets SnapshotRetention field to given value.

### HasSnapshotRetention

`func (o *CreateVolumeRequest) HasSnapshotRetention() bool`

HasSnapshotRetention returns a boolean if a field has been set.

### GetSourceVolumeId

`func (o *CreateVolumeRequest) GetSourceVolumeId() string`

GetSourceVolumeId returns the SourceVolumeId field if non-nil, zero value otherwise.

### GetSourceVolumeIdOk

`func (o *CreateVolumeRequest) GetSourceVolumeIdOk() (*string, bool)`

GetSourceVolumeIdOk returns a tuple with the SourceVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeId

`func (o *CreateVolumeRequest) SetSourceVolumeId(v string)`

SetSourceVolumeId sets SourceVolumeId field to given value.

### HasSourceVolumeId

`func (o *CreateVolumeRequest) HasSourceVolumeId() bool`

HasSourceVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


