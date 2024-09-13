# FlyEnvFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvVar** | Pointer to **string** | EnvVar is required and is the name of the environment variable that will be set from the secret. It must be a valid environment variable name. | [optional] 
**FieldRef** | Pointer to **string** | FieldRef selects a field of the Machine: supports id, version, app_name, private_ip, region, image. | [optional] 

## Methods

### NewFlyEnvFrom

`func NewFlyEnvFrom() *FlyEnvFrom`

NewFlyEnvFrom instantiates a new FlyEnvFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyEnvFromWithDefaults

`func NewFlyEnvFromWithDefaults() *FlyEnvFrom`

NewFlyEnvFromWithDefaults instantiates a new FlyEnvFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvVar

`func (o *FlyEnvFrom) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *FlyEnvFrom) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *FlyEnvFrom) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *FlyEnvFrom) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetFieldRef

`func (o *FlyEnvFrom) GetFieldRef() string`

GetFieldRef returns the FieldRef field if non-nil, zero value otherwise.

### GetFieldRefOk

`func (o *FlyEnvFrom) GetFieldRefOk() (*string, bool)`

GetFieldRefOk returns a tuple with the FieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRef

`func (o *FlyEnvFrom) SetFieldRef(v string)`

SetFieldRef sets FieldRef field to given value.

### HasFieldRef

`func (o *FlyEnvFrom) HasFieldRef() bool`

HasFieldRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


