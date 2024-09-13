# FlyMachineSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvVar** | Pointer to **string** | EnvVar is required and is the name of the environment variable that will be set from the secret. It must be a valid environment variable name. | [optional] 
**Name** | Pointer to **string** | Name is optional and when provided is used to reference a secret name where the EnvVar is different from what was set as the secret name. | [optional] 

## Methods

### NewFlyMachineSecret

`func NewFlyMachineSecret() *FlyMachineSecret`

NewFlyMachineSecret instantiates a new FlyMachineSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineSecretWithDefaults

`func NewFlyMachineSecretWithDefaults() *FlyMachineSecret`

NewFlyMachineSecretWithDefaults instantiates a new FlyMachineSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvVar

`func (o *FlyMachineSecret) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *FlyMachineSecret) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *FlyMachineSecret) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *FlyMachineSecret) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetName

`func (o *FlyMachineSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlyMachineSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlyMachineSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlyMachineSecret) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


