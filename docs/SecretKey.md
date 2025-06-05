# SecretKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **[]int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretKey

`func NewSecretKey() *SecretKey`

NewSecretKey instantiates a new SecretKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretKeyWithDefaults

`func NewSecretKeyWithDefaults() *SecretKey`

NewSecretKeyWithDefaults instantiates a new SecretKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *SecretKey) GetPublicKey() []int32`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SecretKey) GetPublicKeyOk() (*[]int32, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SecretKey) SetPublicKey(v []int32)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SecretKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetType

`func (o *SecretKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecretKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


