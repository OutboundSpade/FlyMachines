# FlyExecHealthcheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | The command to run to check the health of the container (e.g. [\&quot;cat\&quot;, \&quot;/tmp/healthy\&quot;]) | [optional] 

## Methods

### NewFlyExecHealthcheck

`func NewFlyExecHealthcheck() *FlyExecHealthcheck`

NewFlyExecHealthcheck instantiates a new FlyExecHealthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyExecHealthcheckWithDefaults

`func NewFlyExecHealthcheckWithDefaults() *FlyExecHealthcheck`

NewFlyExecHealthcheckWithDefaults instantiates a new FlyExecHealthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *FlyExecHealthcheck) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *FlyExecHealthcheck) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *FlyExecHealthcheck) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *FlyExecHealthcheck) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


