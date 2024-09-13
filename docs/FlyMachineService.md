# FlyMachineService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autostart** | Pointer to **bool** |  | [optional] 
**Autostop** | Pointer to **string** | Accepts a string (new format) or a boolean (old format). For backward compatibility with older clients, the API continues to use booleans for \&quot;off\&quot; and \&quot;stop\&quot; in responses. * \&quot;off\&quot; or false - Do not autostop the Machine. * \&quot;stop\&quot; or true - Automatically stop the Machine. * \&quot;suspend\&quot; - Automatically suspend the Machine, falling back to a full stop if this is not possible. | [optional] 
**Checks** | Pointer to [**[]FlyMachineCheck**](FlyMachineCheck.md) |  | [optional] 
**Concurrency** | Pointer to [**FlyMachineServiceConcurrency**](FlyMachineServiceConcurrency.md) |  | [optional] 
**ForceInstanceDescription** | Pointer to **string** |  | [optional] 
**ForceInstanceKey** | Pointer to **string** |  | [optional] 
**InternalPort** | Pointer to **int32** |  | [optional] 
**MinMachinesRunning** | Pointer to **int32** |  | [optional] 
**Ports** | Pointer to [**[]FlyMachinePort**](FlyMachinePort.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewFlyMachineService

`func NewFlyMachineService() *FlyMachineService`

NewFlyMachineService instantiates a new FlyMachineService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyMachineServiceWithDefaults

`func NewFlyMachineServiceWithDefaults() *FlyMachineService`

NewFlyMachineServiceWithDefaults instantiates a new FlyMachineService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutostart

`func (o *FlyMachineService) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *FlyMachineService) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *FlyMachineService) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *FlyMachineService) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetAutostop

`func (o *FlyMachineService) GetAutostop() string`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *FlyMachineService) GetAutostopOk() (*string, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *FlyMachineService) SetAutostop(v string)`

SetAutostop sets Autostop field to given value.

### HasAutostop

`func (o *FlyMachineService) HasAutostop() bool`

HasAutostop returns a boolean if a field has been set.

### GetChecks

`func (o *FlyMachineService) GetChecks() []FlyMachineCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *FlyMachineService) GetChecksOk() (*[]FlyMachineCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *FlyMachineService) SetChecks(v []FlyMachineCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *FlyMachineService) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetConcurrency

`func (o *FlyMachineService) GetConcurrency() FlyMachineServiceConcurrency`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *FlyMachineService) GetConcurrencyOk() (*FlyMachineServiceConcurrency, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *FlyMachineService) SetConcurrency(v FlyMachineServiceConcurrency)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *FlyMachineService) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetForceInstanceDescription

`func (o *FlyMachineService) GetForceInstanceDescription() string`

GetForceInstanceDescription returns the ForceInstanceDescription field if non-nil, zero value otherwise.

### GetForceInstanceDescriptionOk

`func (o *FlyMachineService) GetForceInstanceDescriptionOk() (*string, bool)`

GetForceInstanceDescriptionOk returns a tuple with the ForceInstanceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstanceDescription

`func (o *FlyMachineService) SetForceInstanceDescription(v string)`

SetForceInstanceDescription sets ForceInstanceDescription field to given value.

### HasForceInstanceDescription

`func (o *FlyMachineService) HasForceInstanceDescription() bool`

HasForceInstanceDescription returns a boolean if a field has been set.

### GetForceInstanceKey

`func (o *FlyMachineService) GetForceInstanceKey() string`

GetForceInstanceKey returns the ForceInstanceKey field if non-nil, zero value otherwise.

### GetForceInstanceKeyOk

`func (o *FlyMachineService) GetForceInstanceKeyOk() (*string, bool)`

GetForceInstanceKeyOk returns a tuple with the ForceInstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstanceKey

`func (o *FlyMachineService) SetForceInstanceKey(v string)`

SetForceInstanceKey sets ForceInstanceKey field to given value.

### HasForceInstanceKey

`func (o *FlyMachineService) HasForceInstanceKey() bool`

HasForceInstanceKey returns a boolean if a field has been set.

### GetInternalPort

`func (o *FlyMachineService) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *FlyMachineService) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *FlyMachineService) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.

### HasInternalPort

`func (o *FlyMachineService) HasInternalPort() bool`

HasInternalPort returns a boolean if a field has been set.

### GetMinMachinesRunning

`func (o *FlyMachineService) GetMinMachinesRunning() int32`

GetMinMachinesRunning returns the MinMachinesRunning field if non-nil, zero value otherwise.

### GetMinMachinesRunningOk

`func (o *FlyMachineService) GetMinMachinesRunningOk() (*int32, bool)`

GetMinMachinesRunningOk returns a tuple with the MinMachinesRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMachinesRunning

`func (o *FlyMachineService) SetMinMachinesRunning(v int32)`

SetMinMachinesRunning sets MinMachinesRunning field to given value.

### HasMinMachinesRunning

`func (o *FlyMachineService) HasMinMachinesRunning() bool`

HasMinMachinesRunning returns a boolean if a field has been set.

### GetPorts

`func (o *FlyMachineService) GetPorts() []FlyMachinePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FlyMachineService) GetPortsOk() (*[]FlyMachinePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FlyMachineService) SetPorts(v []FlyMachinePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FlyMachineService) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProtocol

`func (o *FlyMachineService) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FlyMachineService) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FlyMachineService) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FlyMachineService) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


