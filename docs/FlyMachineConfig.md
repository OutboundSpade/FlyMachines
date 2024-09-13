# FlyMachineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDestroy** | Pointer to **bool** | Optional boolean telling the Machine to destroy itself once it’s complete (default false) | [optional] 
**Checks** | Pointer to [**map[string]FlyMachineCheck**](FlyMachineCheck.md) |  | [optional] 
**DisableMachineAutostart** | Pointer to **bool** | Deprecated: use Service.Autostart instead | [optional] 
**Dns** | Pointer to [**FlyDNSConfig**](FlyDNSConfig.md) |  | [optional] 
**Env** | Pointer to **map[string]string** | An object filled with key/value pairs to be set as environment variables | [optional] 
**Files** | Pointer to [**[]FlyFile**](FlyFile.md) |  | [optional] 
**Guest** | Pointer to [**FlyMachineGuest**](FlyMachineGuest.md) |  | [optional] 
**Image** | Pointer to **string** | The docker image to run | [optional] 
**Init** | Pointer to [**FlyMachineInit**](FlyMachineInit.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Metrics** | Pointer to [**FlyMachineMetrics**](FlyMachineMetrics.md) |  | [optional] 
**Mounts** | Pointer to [**[]FlyMachineMount**](FlyMachineMount.md) |  | [optional] 
**Processes** | Pointer to [**[]FlyMachineProcess**](FlyMachineProcess.md) |  | [optional] 
**Restart** | Pointer to [**FlyMachineRestart**](FlyMachineRestart.md) |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]FlyMachineService**](FlyMachineService.md) |  | [optional] 
**Size** | Pointer to **string** | Deprecated: use Guest instead | [optional] 
**Standbys** | Pointer to **[]string** | Standbys enable a machine to be a standby for another. In the event of a hardware failure, the standby machine will be started. | [optional] 
**Statics** | Pointer to [**[]FlyStatic**](FlyStatic.md) |  | [optional] 
**StopConfig** | Pointer to [**FlyStopConfig**](FlyStopConfig.md) |  | [optional] 

## Methods

### NewFlyMachineConfig

`func NewFlyMachineConfig() *FlyMachineConfig`

NewFlyMachineConfig instantiates a new FlyMachineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineConfigWithDefaults

`func NewFlyMachineConfigWithDefaults() *FlyMachineConfig`

NewFlyMachineConfigWithDefaults instantiates a new FlyMachineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDestroy

`func (o *FlyMachineConfig) GetAutoDestroy() bool`

GetAutoDestroy returns the AutoDestroy field if non-nil, zero value otherwise.

### GetAutoDestroyOk

`func (o *FlyMachineConfig) GetAutoDestroyOk() (*bool, bool)`

GetAutoDestroyOk returns a tuple with the AutoDestroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDestroy

`func (o *FlyMachineConfig) SetAutoDestroy(v bool)`

SetAutoDestroy sets AutoDestroy field to given value.

### HasAutoDestroy

`func (o *FlyMachineConfig) HasAutoDestroy() bool`

HasAutoDestroy returns a boolean if a field has been set.

### GetChecks

`func (o *FlyMachineConfig) GetChecks() map[string]FlyMachineCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *FlyMachineConfig) GetChecksOk() (*map[string]FlyMachineCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *FlyMachineConfig) SetChecks(v map[string]FlyMachineCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *FlyMachineConfig) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetDisableMachineAutostart

`func (o *FlyMachineConfig) GetDisableMachineAutostart() bool`

GetDisableMachineAutostart returns the DisableMachineAutostart field if non-nil, zero value otherwise.

### GetDisableMachineAutostartOk

`func (o *FlyMachineConfig) GetDisableMachineAutostartOk() (*bool, bool)`

GetDisableMachineAutostartOk returns a tuple with the DisableMachineAutostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMachineAutostart

`func (o *FlyMachineConfig) SetDisableMachineAutostart(v bool)`

SetDisableMachineAutostart sets DisableMachineAutostart field to given value.

### HasDisableMachineAutostart

`func (o *FlyMachineConfig) HasDisableMachineAutostart() bool`

HasDisableMachineAutostart returns a boolean if a field has been set.

### GetDns

`func (o *FlyMachineConfig) GetDns() FlyDNSConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *FlyMachineConfig) GetDnsOk() (*FlyDNSConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *FlyMachineConfig) SetDns(v FlyDNSConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *FlyMachineConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetEnv

`func (o *FlyMachineConfig) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FlyMachineConfig) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FlyMachineConfig) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *FlyMachineConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetFiles

`func (o *FlyMachineConfig) GetFiles() []FlyFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FlyMachineConfig) GetFilesOk() (*[]FlyFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FlyMachineConfig) SetFiles(v []FlyFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FlyMachineConfig) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetGuest

`func (o *FlyMachineConfig) GetGuest() FlyMachineGuest`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *FlyMachineConfig) GetGuestOk() (*FlyMachineGuest, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *FlyMachineConfig) SetGuest(v FlyMachineGuest)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *FlyMachineConfig) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetImage

`func (o *FlyMachineConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FlyMachineConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FlyMachineConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FlyMachineConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInit

`func (o *FlyMachineConfig) GetInit() FlyMachineInit`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *FlyMachineConfig) GetInitOk() (*FlyMachineInit, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *FlyMachineConfig) SetInit(v FlyMachineInit)`

SetInit sets Init field to given value.

### HasInit

`func (o *FlyMachineConfig) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetMetadata

`func (o *FlyMachineConfig) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlyMachineConfig) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlyMachineConfig) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlyMachineConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetrics

`func (o *FlyMachineConfig) GetMetrics() FlyMachineMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *FlyMachineConfig) GetMetricsOk() (*FlyMachineMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *FlyMachineConfig) SetMetrics(v FlyMachineMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *FlyMachineConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMounts

`func (o *FlyMachineConfig) GetMounts() []FlyMachineMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *FlyMachineConfig) GetMountsOk() (*[]FlyMachineMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *FlyMachineConfig) SetMounts(v []FlyMachineMount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *FlyMachineConfig) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetProcesses

`func (o *FlyMachineConfig) GetProcesses() []FlyMachineProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *FlyMachineConfig) GetProcessesOk() (*[]FlyMachineProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *FlyMachineConfig) SetProcesses(v []FlyMachineProcess)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *FlyMachineConfig) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetRestart

`func (o *FlyMachineConfig) GetRestart() FlyMachineRestart`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *FlyMachineConfig) GetRestartOk() (*FlyMachineRestart, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *FlyMachineConfig) SetRestart(v FlyMachineRestart)`

SetRestart sets Restart field to given value.

### HasRestart

`func (o *FlyMachineConfig) HasRestart() bool`

HasRestart returns a boolean if a field has been set.

### GetSchedule

`func (o *FlyMachineConfig) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *FlyMachineConfig) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *FlyMachineConfig) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *FlyMachineConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServices

`func (o *FlyMachineConfig) GetServices() []FlyMachineService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *FlyMachineConfig) GetServicesOk() (*[]FlyMachineService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *FlyMachineConfig) SetServices(v []FlyMachineService)`

SetServices sets Services field to given value.

### HasServices

`func (o *FlyMachineConfig) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSize

`func (o *FlyMachineConfig) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FlyMachineConfig) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FlyMachineConfig) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *FlyMachineConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStandbys

`func (o *FlyMachineConfig) GetStandbys() []string`

GetStandbys returns the Standbys field if non-nil, zero value otherwise.

### GetStandbysOk

`func (o *FlyMachineConfig) GetStandbysOk() (*[]string, bool)`

GetStandbysOk returns a tuple with the Standbys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbys

`func (o *FlyMachineConfig) SetStandbys(v []string)`

SetStandbys sets Standbys field to given value.

### HasStandbys

`func (o *FlyMachineConfig) HasStandbys() bool`

HasStandbys returns a boolean if a field has been set.

### GetStatics

`func (o *FlyMachineConfig) GetStatics() []FlyStatic`

GetStatics returns the Statics field if non-nil, zero value otherwise.

### GetStaticsOk

`func (o *FlyMachineConfig) GetStaticsOk() (*[]FlyStatic, bool)`

GetStaticsOk returns a tuple with the Statics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatics

`func (o *FlyMachineConfig) SetStatics(v []FlyStatic)`

SetStatics sets Statics field to given value.

### HasStatics

`func (o *FlyMachineConfig) HasStatics() bool`

HasStatics returns a boolean if a field has been set.

### GetStopConfig

`func (o *FlyMachineConfig) GetStopConfig() FlyStopConfig`

GetStopConfig returns the StopConfig field if non-nil, zero value otherwise.

### GetStopConfigOk

`func (o *FlyMachineConfig) GetStopConfigOk() (*FlyStopConfig, bool)`

GetStopConfigOk returns a tuple with the StopConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopConfig

`func (o *FlyMachineConfig) SetStopConfig(v FlyStopConfig)`

SetStopConfig sets StopConfig field to given value.

### HasStopConfig

`func (o *FlyMachineConfig) HasStopConfig() bool`

HasStopConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


