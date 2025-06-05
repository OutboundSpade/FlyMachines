# VerifySecretkeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plaintext** | Pointer to **[]int32** |  | [optional] 
**Signature** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewVerifySecretkeyRequest

`func NewVerifySecretkeyRequest() *VerifySecretkeyRequest`

NewVerifySecretkeyRequest instantiates a new VerifySecretkeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifySecretkeyRequestWithDefaults

`func NewVerifySecretkeyRequestWithDefaults() *VerifySecretkeyRequest`

NewVerifySecretkeyRequestWithDefaults instantiates a new VerifySecretkeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaintext

`func (o *VerifySecretkeyRequest) GetPlaintext() []int32`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *VerifySecretkeyRequest) GetPlaintextOk() (*[]int32, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *VerifySecretkeyRequest) SetPlaintext(v []int32)`

SetPlaintext sets Plaintext field to given value.

### HasPlaintext

`func (o *VerifySecretkeyRequest) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.

### GetSignature

`func (o *VerifySecretkeyRequest) GetSignature() []int32`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifySecretkeyRequest) GetSignatureOk() (*[]int32, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifySecretkeyRequest) SetSignature(v []int32)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *VerifySecretkeyRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


