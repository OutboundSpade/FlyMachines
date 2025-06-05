# DecryptSecretkeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedData** | Pointer to **[]int64** |  | [optional] 
**Ciphertext** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewDecryptSecretkeyRequest

`func NewDecryptSecretkeyRequest() *DecryptSecretkeyRequest`

NewDecryptSecretkeyRequest instantiates a new DecryptSecretkeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptSecretkeyRequestWithDefaults

`func NewDecryptSecretkeyRequestWithDefaults() *DecryptSecretkeyRequest`

NewDecryptSecretkeyRequestWithDefaults instantiates a new DecryptSecretkeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedData

`func (o *DecryptSecretkeyRequest) GetAssociatedData() []int64`

GetAssociatedData returns the AssociatedData field if non-nil, zero value otherwise.

### GetAssociatedDataOk

`func (o *DecryptSecretkeyRequest) GetAssociatedDataOk() (*[]int64, bool)`

GetAssociatedDataOk returns a tuple with the AssociatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedData

`func (o *DecryptSecretkeyRequest) SetAssociatedData(v []int64)`

SetAssociatedData sets AssociatedData field to given value.

### HasAssociatedData

`func (o *DecryptSecretkeyRequest) HasAssociatedData() bool`

HasAssociatedData returns a boolean if a field has been set.

### GetCiphertext

`func (o *DecryptSecretkeyRequest) GetCiphertext() []int64`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *DecryptSecretkeyRequest) GetCiphertextOk() (*[]int64, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *DecryptSecretkeyRequest) SetCiphertext(v []int64)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *DecryptSecretkeyRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


