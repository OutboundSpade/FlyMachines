# CreateOIDCTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOIDCTokenRequest

`func NewCreateOIDCTokenRequest() *CreateOIDCTokenRequest`

NewCreateOIDCTokenRequest instantiates a new CreateOIDCTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCTokenRequestWithDefaults

`func NewCreateOIDCTokenRequestWithDefaults() *CreateOIDCTokenRequest`

NewCreateOIDCTokenRequestWithDefaults instantiates a new CreateOIDCTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *CreateOIDCTokenRequest) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *CreateOIDCTokenRequest) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *CreateOIDCTokenRequest) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *CreateOIDCTokenRequest) HasAud() bool`

HasAud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


