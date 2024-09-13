# Flydv1ExecResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int32** |  | [optional] 
**ExitSignal** | Pointer to **int32** |  | [optional] 
**Stderr** | Pointer to **string** |  | [optional] 
**Stdout** | Pointer to **string** |  | [optional] 

## Methods

### NewFlydv1ExecResponse

`func NewFlydv1ExecResponse() *Flydv1ExecResponse`

NewFlydv1ExecResponse instantiates a new Flydv1ExecResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlydv1ExecResponseWithDefaults

`func NewFlydv1ExecResponseWithDefaults() *Flydv1ExecResponse`

NewFlydv1ExecResponseWithDefaults instantiates a new Flydv1ExecResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *Flydv1ExecResponse) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Flydv1ExecResponse) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Flydv1ExecResponse) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Flydv1ExecResponse) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExitSignal

`func (o *Flydv1ExecResponse) GetExitSignal() int32`

GetExitSignal returns the ExitSignal field if non-nil, zero value otherwise.

### GetExitSignalOk

`func (o *Flydv1ExecResponse) GetExitSignalOk() (*int32, bool)`

GetExitSignalOk returns a tuple with the ExitSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitSignal

`func (o *Flydv1ExecResponse) SetExitSignal(v int32)`

SetExitSignal sets ExitSignal field to given value.

### HasExitSignal

`func (o *Flydv1ExecResponse) HasExitSignal() bool`

HasExitSignal returns a boolean if a field has been set.

### GetStderr

`func (o *Flydv1ExecResponse) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Flydv1ExecResponse) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Flydv1ExecResponse) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *Flydv1ExecResponse) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *Flydv1ExecResponse) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Flydv1ExecResponse) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Flydv1ExecResponse) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *Flydv1ExecResponse) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


