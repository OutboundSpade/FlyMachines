# EncryptSecretkeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedData** | Pointer to **[]int64** |  | [optional] 
**Plaintext** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewEncryptSecretkeyRequest

`func NewEncryptSecretkeyRequest() *EncryptSecretkeyRequest`

NewEncryptSecretkeyRequest instantiates a new EncryptSecretkeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptSecretkeyRequestWithDefaults

`func NewEncryptSecretkeyRequestWithDefaults() *EncryptSecretkeyRequest`

NewEncryptSecretkeyRequestWithDefaults instantiates a new EncryptSecretkeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedData

`func (o *EncryptSecretkeyRequest) GetAssociatedData() []int64`

GetAssociatedData returns the AssociatedData field if non-nil, zero value otherwise.

### GetAssociatedDataOk

`func (o *EncryptSecretkeyRequest) GetAssociatedDataOk() (*[]int64, bool)`

GetAssociatedDataOk returns a tuple with the AssociatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedData

`func (o *EncryptSecretkeyRequest) SetAssociatedData(v []int64)`

SetAssociatedData sets AssociatedData field to given value.

### HasAssociatedData

`func (o *EncryptSecretkeyRequest) HasAssociatedData() bool`

HasAssociatedData returns a boolean if a field has been set.

### GetPlaintext

`func (o *EncryptSecretkeyRequest) GetPlaintext() []int64`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *EncryptSecretkeyRequest) GetPlaintextOk() (*[]int64, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *EncryptSecretkeyRequest) SetPlaintext(v []int64)`

SetPlaintext sets Plaintext field to given value.

### HasPlaintext

`func (o *EncryptSecretkeyRequest) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


