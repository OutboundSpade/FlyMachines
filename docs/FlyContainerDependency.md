# FlyContainerDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**FlyContainerDependencyCondition**](FlyContainerDependencyCondition.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewFlyContainerDependency

`func NewFlyContainerDependency() *FlyContainerDependency`

NewFlyContainerDependency instantiates a new FlyContainerDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyContainerDependencyWithDefaults

`func NewFlyContainerDependencyWithDefaults() *FlyContainerDependency`

NewFlyContainerDependencyWithDefaults instantiates a new FlyContainerDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *FlyContainerDependency) GetCondition() FlyContainerDependencyCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *FlyContainerDependency) GetConditionOk() (*FlyContainerDependencyCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *FlyContainerDependency) SetCondition(v FlyContainerDependencyCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *FlyContainerDependency) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetName

`func (o *FlyContainerDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlyContainerDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlyContainerDependency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlyContainerDependency) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


