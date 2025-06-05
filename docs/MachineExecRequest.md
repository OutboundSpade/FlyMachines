# MachineExecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmd** | Pointer to **string** | Deprecated: use Command instead | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Stdin** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewMachineExecRequest

`func NewMachineExecRequest() *MachineExecRequest`

NewMachineExecRequest instantiates a new MachineExecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineExecRequestWithDefaults

`func NewMachineExecRequestWithDefaults() *MachineExecRequest`

NewMachineExecRequestWithDefaults instantiates a new MachineExecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmd

`func (o *MachineExecRequest) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *MachineExecRequest) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *MachineExecRequest) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *MachineExecRequest) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetCommand

`func (o *MachineExecRequest) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *MachineExecRequest) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *MachineExecRequest) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *MachineExecRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetContainer

`func (o *MachineExecRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *MachineExecRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *MachineExecRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *MachineExecRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetStdin

`func (o *MachineExecRequest) GetStdin() string`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *MachineExecRequest) GetStdinOk() (*string, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *MachineExecRequest) SetStdin(v string)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *MachineExecRequest) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetTimeout

`func (o *MachineExecRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MachineExecRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MachineExecRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MachineExecRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


