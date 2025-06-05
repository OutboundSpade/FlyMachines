# AppSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]AppSecret**](AppSecret.md) |  | [optional] 

## Methods

### NewAppSecrets

`func NewAppSecrets() *AppSecrets`

NewAppSecrets instantiates a new AppSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSecretsWithDefaults

`func NewAppSecretsWithDefaults() *AppSecrets`

NewAppSecretsWithDefaults instantiates a new AppSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *AppSecrets) GetSecrets() []AppSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *AppSecrets) GetSecretsOk() (*[]AppSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *AppSecrets) SetSecrets(v []AppSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *AppSecrets) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


