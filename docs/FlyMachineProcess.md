# FlyMachineProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmd** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**EnvFrom** | Pointer to [**[]FlyEnvFrom**](FlyEnvFrom.md) | EnvFrom can be provided to set environment variables from machine fields. | [optional] 
**Exec** | Pointer to **[]string** |  | [optional] 
**IgnoreAppSecrets** | Pointer to **bool** | IgnoreAppSecrets can be set to true to ignore the secrets for the App the Machine belongs to and only use the secrets provided at the process level. The default/legacy behavior is to use the secrets provided at the App level. | [optional] 
**Secrets** | Pointer to [**[]FlyMachineSecret**](FlyMachineSecret.md) | Secrets can be provided at the process level to explicitly indicate which secrets should be used for the process. If not provided, the secrets provided at the machine level will be used. | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewFlyMachineProcess

`func NewFlyMachineProcess() *FlyMachineProcess`

NewFlyMachineProcess instantiates a new FlyMachineProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineProcessWithDefaults

`func NewFlyMachineProcessWithDefaults() *FlyMachineProcess`

NewFlyMachineProcessWithDefaults instantiates a new FlyMachineProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmd

`func (o *FlyMachineProcess) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *FlyMachineProcess) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *FlyMachineProcess) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *FlyMachineProcess) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetEntrypoint

`func (o *FlyMachineProcess) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *FlyMachineProcess) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *FlyMachineProcess) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *FlyMachineProcess) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *FlyMachineProcess) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FlyMachineProcess) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FlyMachineProcess) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *FlyMachineProcess) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvFrom

`func (o *FlyMachineProcess) GetEnvFrom() []FlyEnvFrom`

GetEnvFrom returns the EnvFrom field if non-nil, zero value otherwise.

### GetEnvFromOk

`func (o *FlyMachineProcess) GetEnvFromOk() (*[]FlyEnvFrom, bool)`

GetEnvFromOk returns a tuple with the EnvFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvFrom

`func (o *FlyMachineProcess) SetEnvFrom(v []FlyEnvFrom)`

SetEnvFrom sets EnvFrom field to given value.

### HasEnvFrom

`func (o *FlyMachineProcess) HasEnvFrom() bool`

HasEnvFrom returns a boolean if a field has been set.

### GetExec

`func (o *FlyMachineProcess) GetExec() []string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *FlyMachineProcess) GetExecOk() (*[]string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *FlyMachineProcess) SetExec(v []string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *FlyMachineProcess) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetIgnoreAppSecrets

`func (o *FlyMachineProcess) GetIgnoreAppSecrets() bool`

GetIgnoreAppSecrets returns the IgnoreAppSecrets field if non-nil, zero value otherwise.

### GetIgnoreAppSecretsOk

`func (o *FlyMachineProcess) GetIgnoreAppSecretsOk() (*bool, bool)`

GetIgnoreAppSecretsOk returns a tuple with the IgnoreAppSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAppSecrets

`func (o *FlyMachineProcess) SetIgnoreAppSecrets(v bool)`

SetIgnoreAppSecrets sets IgnoreAppSecrets field to given value.

### HasIgnoreAppSecrets

`func (o *FlyMachineProcess) HasIgnoreAppSecrets() bool`

HasIgnoreAppSecrets returns a boolean if a field has been set.

### GetSecrets

`func (o *FlyMachineProcess) GetSecrets() []FlyMachineSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FlyMachineProcess) GetSecretsOk() (*[]FlyMachineSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FlyMachineProcess) SetSecrets(v []FlyMachineSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *FlyMachineProcess) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetUser

`func (o *FlyMachineProcess) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FlyMachineProcess) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FlyMachineProcess) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FlyMachineProcess) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


