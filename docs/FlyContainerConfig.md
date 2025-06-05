# FlyContainerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmd** | Pointer to **[]string** | CmdOverride is used to override the default command of the image. | [optional] 
**DependsOn** | Pointer to [**[]FlyContainerDependency**](FlyContainerDependency.md) | DependsOn can be used to define dependencies between containers. The container will only be started after all of its dependent conditions have been satisfied. | [optional] 
**Entrypoint** | Pointer to **[]string** | EntrypointOverride is used to override the default entrypoint of the image. | [optional] 
**Env** | Pointer to **map[string]string** | ExtraEnv is used to add additional environment variables to the container. | [optional] 
**EnvFrom** | Pointer to [**[]FlyEnvFrom**](FlyEnvFrom.md) | EnvFrom can be provided to set environment variables from machine fields. | [optional] 
**Exec** | Pointer to **[]string** | Image Config overrides - these fields are used to override the image configuration. If not provided, the image configuration will be used. ExecOverride is used to override the default command of the image. | [optional] 
**Files** | Pointer to [**[]FlyFile**](FlyFile.md) | Files are files that will be written to the container file system. | [optional] 
**Healthchecks** | Pointer to [**[]FlyContainerHealthcheck**](FlyContainerHealthcheck.md) | Healthchecks determine the health of your containers. Healthchecks can use HTTP, TCP or an Exec command. | [optional] 
**Image** | Pointer to **string** | Image is the docker image to run. | [optional] 
**Name** | Pointer to **string** | Name is used to identify the container in the machine. | [optional] 
**Restart** | Pointer to [**FlyMachineRestart**](FlyMachineRestart.md) | Restart is used to define the restart policy for the container. NOTE: spot-price is not supported for containers. | [optional] 
**Secrets** | Pointer to [**[]FlyMachineSecret**](FlyMachineSecret.md) | Secrets can be provided at the process level to explicitly indicate which secrets should be used for the process. If not provided, the secrets provided at the machine level will be used. | [optional] 
**Stop** | Pointer to [**FlyStopConfig**](FlyStopConfig.md) | Stop is used to define the signal and timeout for stopping the container. | [optional] 
**User** | Pointer to **string** | UserOverride is used to override the default user of the image. | [optional] 

## Methods

### NewFlyContainerConfig

`func NewFlyContainerConfig() *FlyContainerConfig`

NewFlyContainerConfig instantiates a new FlyContainerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyContainerConfigWithDefaults

`func NewFlyContainerConfigWithDefaults() *FlyContainerConfig`

NewFlyContainerConfigWithDefaults instantiates a new FlyContainerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmd

`func (o *FlyContainerConfig) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *FlyContainerConfig) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *FlyContainerConfig) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *FlyContainerConfig) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetDependsOn

`func (o *FlyContainerConfig) GetDependsOn() []FlyContainerDependency`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *FlyContainerConfig) GetDependsOnOk() (*[]FlyContainerDependency, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *FlyContainerConfig) SetDependsOn(v []FlyContainerDependency)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *FlyContainerConfig) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetEntrypoint

`func (o *FlyContainerConfig) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *FlyContainerConfig) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *FlyContainerConfig) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *FlyContainerConfig) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *FlyContainerConfig) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FlyContainerConfig) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FlyContainerConfig) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *FlyContainerConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvFrom

`func (o *FlyContainerConfig) GetEnvFrom() []FlyEnvFrom`

GetEnvFrom returns the EnvFrom field if non-nil, zero value otherwise.

### GetEnvFromOk

`func (o *FlyContainerConfig) GetEnvFromOk() (*[]FlyEnvFrom, bool)`

GetEnvFromOk returns a tuple with the EnvFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvFrom

`func (o *FlyContainerConfig) SetEnvFrom(v []FlyEnvFrom)`

SetEnvFrom sets EnvFrom field to given value.

### HasEnvFrom

`func (o *FlyContainerConfig) HasEnvFrom() bool`

HasEnvFrom returns a boolean if a field has been set.

### GetExec

`func (o *FlyContainerConfig) GetExec() []string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *FlyContainerConfig) GetExecOk() (*[]string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *FlyContainerConfig) SetExec(v []string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *FlyContainerConfig) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFiles

`func (o *FlyContainerConfig) GetFiles() []FlyFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FlyContainerConfig) GetFilesOk() (*[]FlyFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FlyContainerConfig) SetFiles(v []FlyFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FlyContainerConfig) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetHealthchecks

`func (o *FlyContainerConfig) GetHealthchecks() []FlyContainerHealthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *FlyContainerConfig) GetHealthchecksOk() (*[]FlyContainerHealthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *FlyContainerConfig) SetHealthchecks(v []FlyContainerHealthcheck)`

SetHealthchecks sets Healthchecks field to given value.

### HasHealthchecks

`func (o *FlyContainerConfig) HasHealthchecks() bool`

HasHealthchecks returns a boolean if a field has been set.

### GetImage

`func (o *FlyContainerConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FlyContainerConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FlyContainerConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FlyContainerConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *FlyContainerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlyContainerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlyContainerConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlyContainerConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestart

`func (o *FlyContainerConfig) GetRestart() FlyMachineRestart`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *FlyContainerConfig) GetRestartOk() (*FlyMachineRestart, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *FlyContainerConfig) SetRestart(v FlyMachineRestart)`

SetRestart sets Restart field to given value.

### HasRestart

`func (o *FlyContainerConfig) HasRestart() bool`

HasRestart returns a boolean if a field has been set.

### GetSecrets

`func (o *FlyContainerConfig) GetSecrets() []FlyMachineSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FlyContainerConfig) GetSecretsOk() (*[]FlyMachineSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FlyContainerConfig) SetSecrets(v []FlyMachineSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *FlyContainerConfig) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetStop

`func (o *FlyContainerConfig) GetStop() FlyStopConfig`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *FlyContainerConfig) GetStopOk() (*FlyStopConfig, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *FlyContainerConfig) SetStop(v FlyStopConfig)`

SetStop sets Stop field to given value.

### HasStop

`func (o *FlyContainerConfig) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetUser

`func (o *FlyContainerConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FlyContainerConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FlyContainerConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FlyContainerConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


