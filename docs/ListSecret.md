# ListSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Publickey** | Pointer to **[]int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewListSecret

`func NewListSecret() *ListSecret`

NewListSecret instantiates a new ListSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecretWithDefaults

`func NewListSecretWithDefaults() *ListSecret`

NewListSecretWithDefaults instantiates a new ListSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ListSecret) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListSecret) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListSecret) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListSecret) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPublickey

`func (o *ListSecret) GetPublickey() []int32`

GetPublickey returns the Publickey field if non-nil, zero value otherwise.

### GetPublickeyOk

`func (o *ListSecret) GetPublickeyOk() (*[]int32, bool)`

GetPublickeyOk returns a tuple with the Publickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublickey

`func (o *ListSecret) SetPublickey(v []int32)`

SetPublickey sets Publickey field to given value.

### HasPublickey

`func (o *ListSecret) HasPublickey() bool`

HasPublickey returns a boolean if a field has been set.

### GetType

`func (o *ListSecret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSecret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSecret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListSecret) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


