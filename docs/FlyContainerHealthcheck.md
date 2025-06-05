# FlyContainerHealthcheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**FlyExecHealthcheck**](FlyExecHealthcheck.md) |  | [optional] 
**FailureThreshold** | Pointer to **int64** | The number of times the check must fail before considering the container unhealthy. | [optional] 
**GracePeriod** | Pointer to **int64** | The time in seconds to wait after a container starts before checking its health. | [optional] 
**Http** | Pointer to [**FlyHTTPHealthcheck**](FlyHTTPHealthcheck.md) |  | [optional] 
**Interval** | Pointer to **int64** | The time in seconds between executing the defined check. | [optional] 
**Kind** | Pointer to [**FlyContainerHealthcheckKind**](FlyContainerHealthcheckKind.md) | Kind of healthcheck (readiness, liveness) | [optional] 
**Name** | Pointer to **string** | The name of the check. Must be unique within the container. | [optional] 
**SuccessThreshold** | Pointer to **int64** | The number of times the check must succeeed before considering the container healthy. | [optional] 
**Tcp** | Pointer to [**FlyTCPHealthcheck**](FlyTCPHealthcheck.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The time in seconds to wait for the check to complete. | [optional] 
**Unhealthy** | Pointer to [**FlyUnhealthyPolicy**](FlyUnhealthyPolicy.md) | Unhealthy policy that determines what action to take if a container is deemed unhealthy | [optional] 

## Methods

### NewFlyContainerHealthcheck

`func NewFlyContainerHealthcheck() *FlyContainerHealthcheck`

NewFlyContainerHealthcheck instantiates a new FlyContainerHealthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyContainerHealthcheckWithDefaults

`func NewFlyContainerHealthcheckWithDefaults() *FlyContainerHealthcheck`

NewFlyContainerHealthcheckWithDefaults instantiates a new FlyContainerHealthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *FlyContainerHealthcheck) GetExec() FlyExecHealthcheck`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *FlyContainerHealthcheck) GetExecOk() (*FlyExecHealthcheck, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *FlyContainerHealthcheck) SetExec(v FlyExecHealthcheck)`

SetExec sets Exec field to given value.

### HasExec

`func (o *FlyContainerHealthcheck) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *FlyContainerHealthcheck) GetFailureThreshold() int64`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *FlyContainerHealthcheck) GetFailureThresholdOk() (*int64, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *FlyContainerHealthcheck) SetFailureThreshold(v int64)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *FlyContainerHealthcheck) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetGracePeriod

`func (o *FlyContainerHealthcheck) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *FlyContainerHealthcheck) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *FlyContainerHealthcheck) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *FlyContainerHealthcheck) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetHttp

`func (o *FlyContainerHealthcheck) GetHttp() FlyHTTPHealthcheck`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *FlyContainerHealthcheck) GetHttpOk() (*FlyHTTPHealthcheck, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *FlyContainerHealthcheck) SetHttp(v FlyHTTPHealthcheck)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *FlyContainerHealthcheck) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetInterval

`func (o *FlyContainerHealthcheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *FlyContainerHealthcheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *FlyContainerHealthcheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *FlyContainerHealthcheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKind

`func (o *FlyContainerHealthcheck) GetKind() FlyContainerHealthcheckKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlyContainerHealthcheck) GetKindOk() (*FlyContainerHealthcheckKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlyContainerHealthcheck) SetKind(v FlyContainerHealthcheckKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FlyContainerHealthcheck) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *FlyContainerHealthcheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlyContainerHealthcheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlyContainerHealthcheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlyContainerHealthcheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *FlyContainerHealthcheck) GetSuccessThreshold() int64`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *FlyContainerHealthcheck) GetSuccessThresholdOk() (*int64, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *FlyContainerHealthcheck) SetSuccessThreshold(v int64)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *FlyContainerHealthcheck) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetTcp

`func (o *FlyContainerHealthcheck) GetTcp() FlyTCPHealthcheck`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *FlyContainerHealthcheck) GetTcpOk() (*FlyTCPHealthcheck, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *FlyContainerHealthcheck) SetTcp(v FlyTCPHealthcheck)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *FlyContainerHealthcheck) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetTimeout

`func (o *FlyContainerHealthcheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FlyContainerHealthcheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FlyContainerHealthcheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FlyContainerHealthcheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUnhealthy

`func (o *FlyContainerHealthcheck) GetUnhealthy() FlyUnhealthyPolicy`

GetUnhealthy returns the Unhealthy field if non-nil, zero value otherwise.

### GetUnhealthyOk

`func (o *FlyContainerHealthcheck) GetUnhealthyOk() (*FlyUnhealthyPolicy, bool)`

GetUnhealthyOk returns a tuple with the Unhealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthy

`func (o *FlyContainerHealthcheck) SetUnhealthy(v FlyUnhealthyPolicy)`

SetUnhealthy sets Unhealthy field to given value.

### HasUnhealthy

`func (o *FlyContainerHealthcheck) HasUnhealthy() bool`

HasUnhealthy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


