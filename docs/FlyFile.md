# FlyFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestPath** | Pointer to **string** | GuestPath is the path on the machine where the file will be written and must be an absolute path. For example: /full/path/to/file.json | [optional] 
**Mode** | Pointer to **int32** | Mode bits used to set permissions on this file as accepted by chmod(2). | [optional] 
**RawValue** | Pointer to **string** | The base64 encoded string of the file contents. | [optional] 
**SecretName** | Pointer to **string** | The name of the secret that contains the base64 encoded file contents. | [optional] 

## Methods

### NewFlyFile

`func NewFlyFile() *FlyFile`

NewFlyFile instantiates a new FlyFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyFileWithDefaults

`func NewFlyFileWithDefaults() *FlyFile`

NewFlyFileWithDefaults instantiates a new FlyFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestPath

`func (o *FlyFile) GetGuestPath() string`

GetGuestPath returns the GuestPath field if non-nil, zero value otherwise.

### GetGuestPathOk

`func (o *FlyFile) GetGuestPathOk() (*string, bool)`

GetGuestPathOk returns a tuple with the GuestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPath

`func (o *FlyFile) SetGuestPath(v string)`

SetGuestPath sets GuestPath field to given value.

### HasGuestPath

`func (o *FlyFile) HasGuestPath() bool`

HasGuestPath returns a boolean if a field has been set.

### GetMode

`func (o *FlyFile) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FlyFile) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FlyFile) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FlyFile) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRawValue

`func (o *FlyFile) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *FlyFile) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *FlyFile) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *FlyFile) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### GetSecretName

`func (o *FlyFile) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *FlyFile) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *FlyFile) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *FlyFile) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


