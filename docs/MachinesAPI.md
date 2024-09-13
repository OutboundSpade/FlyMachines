# \MachinesAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachinesCordon**](MachinesAPI.md#MachinesCordon) | **Post** /apps/{app_name}/machines/{machine_id}/cordon | Cordon Machine
[**MachinesCreate**](MachinesAPI.md#MachinesCreate) | **Post** /apps/{app_name}/machines | Create Machine
[**MachinesCreateLease**](MachinesAPI.md#MachinesCreateLease) | **Post** /apps/{app_name}/machines/{machine_id}/lease | Create Lease
[**MachinesDelete**](MachinesAPI.md#MachinesDelete) | **Delete** /apps/{app_name}/machines/{machine_id} | Destroy Machine
[**MachinesDeleteMetadata**](MachinesAPI.md#MachinesDeleteMetadata) | **Delete** /apps/{app_name}/machines/{machine_id}/metadata/{key} | Delete Metadata
[**MachinesExec**](MachinesAPI.md#MachinesExec) | **Post** /apps/{app_name}/machines/{machine_id}/exec | Execute Command
[**MachinesList**](MachinesAPI.md#MachinesList) | **Get** /apps/{app_name}/machines | List Machines
[**MachinesListEvents**](MachinesAPI.md#MachinesListEvents) | **Get** /apps/{app_name}/machines/{machine_id}/events | List Events
[**MachinesListProcesses**](MachinesAPI.md#MachinesListProcesses) | **Get** /apps/{app_name}/machines/{machine_id}/ps | List Processes
[**MachinesListVersions**](MachinesAPI.md#MachinesListVersions) | **Get** /apps/{app_name}/machines/{machine_id}/versions | List Versions
[**MachinesReleaseLease**](MachinesAPI.md#MachinesReleaseLease) | **Delete** /apps/{app_name}/machines/{machine_id}/lease | Release Lease
[**MachinesRestart**](MachinesAPI.md#MachinesRestart) | **Post** /apps/{app_name}/machines/{machine_id}/restart | Restart Machine
[**MachinesShow**](MachinesAPI.md#MachinesShow) | **Get** /apps/{app_name}/machines/{machine_id} | Get Machine
[**MachinesShowLease**](MachinesAPI.md#MachinesShowLease) | **Get** /apps/{app_name}/machines/{machine_id}/lease | Get Lease
[**MachinesShowMetadata**](MachinesAPI.md#MachinesShowMetadata) | **Get** /apps/{app_name}/machines/{machine_id}/metadata | Get Metadata
[**MachinesSignal**](MachinesAPI.md#MachinesSignal) | **Post** /apps/{app_name}/machines/{machine_id}/signal | Signal Machine
[**MachinesStart**](MachinesAPI.md#MachinesStart) | **Post** /apps/{app_name}/machines/{machine_id}/start | Start Machine
[**MachinesStop**](MachinesAPI.md#MachinesStop) | **Post** /apps/{app_name}/machines/{machine_id}/stop | Stop Machine
[**MachinesSuspend**](MachinesAPI.md#MachinesSuspend) | **Post** /apps/{app_name}/machines/{machine_id}/suspend | Suspend Machine
[**MachinesUncordon**](MachinesAPI.md#MachinesUncordon) | **Post** /apps/{app_name}/machines/{machine_id}/uncordon | Uncordon Machine
[**MachinesUpdate**](MachinesAPI.md#MachinesUpdate) | **Post** /apps/{app_name}/machines/{machine_id} | Update Machine
[**MachinesUpdateMetadata**](MachinesAPI.md#MachinesUpdateMetadata) | **Post** /apps/{app_name}/machines/{machine_id}/metadata/{key} | Update Metadata
[**MachinesWait**](MachinesAPI.md#MachinesWait) | **Get** /apps/{app_name}/machines/{machine_id}/wait | Wait for State



## MachinesCordon

> MachinesCordon(ctx, appName, machineId).Execute()

Cordon Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesCordon(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesCordon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesCordonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesCreate

> Machine MachinesCreate(ctx, appName).Request(request).Execute()

Create Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	request := *openapiclient.NewCreateMachineRequest() // CreateMachineRequest | Create machine request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesCreate(context.Background(), appName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesCreate`: Machine
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateMachineRequest**](CreateMachineRequest.md) | Create machine request | 

### Return type

[**Machine**](Machine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesCreateLease

> Lease MachinesCreateLease(ctx, appName, machineId).Request(request).FlyMachineLeaseNonce(flyMachineLeaseNonce).Execute()

Create Lease



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	request := *openapiclient.NewCreateLeaseRequest() // CreateLeaseRequest | Request body
	flyMachineLeaseNonce := "flyMachineLeaseNonce_example" // string | Existing lease nonce to refresh by ttl, empty or non-existent to create a new lease (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesCreateLease(context.Background(), appName, machineId).Request(request).FlyMachineLeaseNonce(flyMachineLeaseNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesCreateLease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesCreateLease`: Lease
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesCreateLease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesCreateLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateLeaseRequest**](CreateLeaseRequest.md) | Request body | 
 **flyMachineLeaseNonce** | **string** | Existing lease nonce to refresh by ttl, empty or non-existent to create a new lease | 

### Return type

[**Lease**](Lease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesDelete

> MachinesDelete(ctx, appName, machineId).Force(force).Execute()

Destroy Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	force := true // bool | Force kill the machine if it's running (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesDelete(context.Background(), appName, machineId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Force kill the machine if it&#39;s running | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesDeleteMetadata

> MachinesDeleteMetadata(ctx, appName, machineId, key).Execute()

Delete Metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	key := "key_example" // string | Metadata Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesDeleteMetadata(context.Background(), appName, machineId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesDeleteMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 
**key** | **string** | Metadata Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesDeleteMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesExec

> Flydv1ExecResponse MachinesExec(ctx, appName, machineId).Request(request).Execute()

Execute Command



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	request := *openapiclient.NewMachineExecRequest() // MachineExecRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesExec(context.Background(), appName, machineId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesExec`: Flydv1ExecResponse
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesExec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**MachineExecRequest**](MachineExecRequest.md) | Request body | 

### Return type

[**Flydv1ExecResponse**](Flydv1ExecResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesList

> []Machine MachinesList(ctx, appName).IncludeDeleted(includeDeleted).Region(region).Execute()

List Machines



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	includeDeleted := true // bool | Include deleted machines (optional)
	region := "region_example" // string | Region filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesList(context.Background(), appName).IncludeDeleted(includeDeleted).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesList`: []Machine
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **bool** | Include deleted machines | 
 **region** | **string** | Region filter | 

### Return type

[**[]Machine**](Machine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesListEvents

> []MachineEvent MachinesListEvents(ctx, appName, machineId).Execute()

List Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesListEvents(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesListEvents`: []MachineEvent
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesListEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MachineEvent**](MachineEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesListProcesses

> []ProcessStat MachinesListProcesses(ctx, appName, machineId).SortBy(sortBy).Order(order).Execute()

List Processes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	sortBy := "sortBy_example" // string | Sort by (optional)
	order := "order_example" // string | Order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesListProcesses(context.Background(), appName, machineId).SortBy(sortBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesListProcesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesListProcesses`: []ProcessStat
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesListProcesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesListProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | Sort by | 
 **order** | **string** | Order | 

### Return type

[**[]ProcessStat**](ProcessStat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesListVersions

> []MachineVersion MachinesListVersions(ctx, appName, machineId).Execute()

List Versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesListVersions(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesListVersions`: []MachineVersion
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MachineVersion**](MachineVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesReleaseLease

> MachinesReleaseLease(ctx, appName, machineId).FlyMachineLeaseNonce(flyMachineLeaseNonce).Execute()

Release Lease



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	flyMachineLeaseNonce := "flyMachineLeaseNonce_example" // string | Existing lease nonce

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesReleaseLease(context.Background(), appName, machineId).FlyMachineLeaseNonce(flyMachineLeaseNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesReleaseLease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesReleaseLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flyMachineLeaseNonce** | **string** | Existing lease nonce | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesRestart

> MachinesRestart(ctx, appName, machineId).Timeout(timeout).Signal(signal).Execute()

Restart Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	timeout := "timeout_example" // string | Restart timeout as a Go duration string or number of seconds (optional)
	signal := "signal_example" // string | Unix signal name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesRestart(context.Background(), appName, machineId).Timeout(timeout).Signal(signal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **string** | Restart timeout as a Go duration string or number of seconds | 
 **signal** | **string** | Unix signal name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesShow

> Machine MachinesShow(ctx, appName, machineId).Execute()

Get Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesShow(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesShow`: Machine
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Machine**](Machine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesShowLease

> Lease MachinesShowLease(ctx, appName, machineId).Execute()

Get Lease



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesShowLease(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesShowLease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesShowLease`: Lease
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesShowLease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesShowLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Lease**](Lease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesShowMetadata

> map[string]string MachinesShowMetadata(ctx, appName, machineId).Execute()

Get Metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesShowMetadata(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesShowMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesShowMetadata`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesShowMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesShowMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesSignal

> MachinesSignal(ctx, appName, machineId).Request(request).Execute()

Signal Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	request := *openapiclient.NewSignalRequest() // SignalRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesSignal(context.Background(), appName, machineId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesSignal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesSignalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SignalRequest**](SignalRequest.md) | Request body | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesStart

> MachinesStart(ctx, appName, machineId).Execute()

Start Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesStart(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesStop

> MachinesStop(ctx, appName, machineId).Request(request).Execute()

Stop Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	request := *openapiclient.NewStopRequest() // StopRequest | Optional request body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesStop(context.Background(), appName, machineId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**StopRequest**](StopRequest.md) | Optional request body | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesSuspend

> MachinesSuspend(ctx, appName, machineId).Execute()

Suspend Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesSuspend(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesUncordon

> MachinesUncordon(ctx, appName, machineId).Execute()

Uncordon Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesUncordon(context.Background(), appName, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesUncordon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesUncordonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesUpdate

> Machine MachinesUpdate(ctx, appName, machineId).Request(request).Execute()

Update Machine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	request := *openapiclient.NewUpdateMachineRequest() // UpdateMachineRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.MachinesUpdate(context.Background(), appName, machineId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MachinesUpdate`: Machine
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.MachinesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateMachineRequest**](UpdateMachineRequest.md) | Request body | 

### Return type

[**Machine**](Machine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesUpdateMetadata

> MachinesUpdateMetadata(ctx, appName, machineId, key).Execute()

Update Metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	key := "key_example" // string | Metadata Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesUpdateMetadata(context.Background(), appName, machineId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesUpdateMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 
**key** | **string** | Metadata Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesUpdateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachinesWait

> MachinesWait(ctx, appName, machineId).InstanceId(instanceId).Timeout(timeout).State(state).Execute()

Wait for State



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func main() {
	appName := "appName_example" // string | Fly App Name
	machineId := "machineId_example" // string | Machine ID
	instanceId := "instanceId_example" // string | 26-character Machine version ID (optional)
	timeout := int32(56) // int32 | wait timeout. default 60s (optional)
	state := "state_example" // string | desired state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.MachinesWait(context.Background(), appName, machineId).InstanceId(instanceId).Timeout(timeout).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.MachinesWait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**machineId** | **string** | Machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachinesWaitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceId** | **string** | 26-character Machine version ID | 
 **timeout** | **int32** | wait timeout. default 60s | 
 **state** | **string** | desired state | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

