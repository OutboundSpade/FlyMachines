# \VolumesAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeSnapshot**](VolumesAPI.md#CreateVolumeSnapshot) | **Post** /apps/{app_name}/volumes/{volume_id}/snapshots | Create Snapshot
[**VolumeDelete**](VolumesAPI.md#VolumeDelete) | **Delete** /apps/{app_name}/volumes/{volume_id} | Destroy Volume
[**VolumesCreate**](VolumesAPI.md#VolumesCreate) | **Post** /apps/{app_name}/volumes | Create Volume
[**VolumesExtend**](VolumesAPI.md#VolumesExtend) | **Put** /apps/{app_name}/volumes/{volume_id}/extend | Extend Volume
[**VolumesGetById**](VolumesAPI.md#VolumesGetById) | **Get** /apps/{app_name}/volumes/{volume_id} | Get Volume
[**VolumesList**](VolumesAPI.md#VolumesList) | **Get** /apps/{app_name}/volumes | List Volumes
[**VolumesListSnapshots**](VolumesAPI.md#VolumesListSnapshots) | **Get** /apps/{app_name}/volumes/{volume_id}/snapshots | List Snapshots
[**VolumesUpdate**](VolumesAPI.md#VolumesUpdate) | **Put** /apps/{app_name}/volumes/{volume_id} | Update Volume



## CreateVolumeSnapshot

> CreateVolumeSnapshot(ctx, appName, volumeId).Execute()

Create Snapshot



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
	volumeId := "volumeId_example" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumesAPI.CreateVolumeSnapshot(context.Background(), appName, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.CreateVolumeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeSnapshotRequest struct via the builder pattern


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


## VolumeDelete

> Volume VolumeDelete(ctx, appName, volumeId).Execute()

Destroy Volume



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
	volumeId := "volumeId_example" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeDelete(context.Background(), appName, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeDelete`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesCreate

> Volume VolumesCreate(ctx, appName).Request(request).Execute()

Create Volume



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
	request := *openapiclient.NewCreateVolumeRequest() // CreateVolumeRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesCreate(context.Background(), appName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesCreate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateVolumeRequest**](CreateVolumeRequest.md) | Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesExtend

> ExtendVolumeResponse VolumesExtend(ctx, appName, volumeId).Request(request).Execute()

Extend Volume



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
	volumeId := "volumeId_example" // string | Volume ID
	request := *openapiclient.NewExtendVolumeRequest() // ExtendVolumeRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesExtend(context.Background(), appName, volumeId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesExtend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesExtend`: ExtendVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesExtend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ExtendVolumeRequest**](ExtendVolumeRequest.md) | Request body | 

### Return type

[**ExtendVolumeResponse**](ExtendVolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesGetById

> Volume VolumesGetById(ctx, appName, volumeId).Execute()

Get Volume



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
	volumeId := "volumeId_example" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesGetById(context.Background(), appName, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesGetById`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesList

> []Volume VolumesList(ctx, appName).Execute()

List Volumes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesList(context.Background(), appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesListSnapshots

> []VolumeSnapshot VolumesListSnapshots(ctx, appName, volumeId).Execute()

List Snapshots



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
	volumeId := "volumeId_example" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesListSnapshots(context.Background(), appName, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesListSnapshots`: []VolumeSnapshot
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesListSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]VolumeSnapshot**](VolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesUpdate

> Volume VolumesUpdate(ctx, appName, volumeId).Request(request).Execute()

Update Volume



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
	volumeId := "volumeId_example" // string | Volume ID
	request := *openapiclient.NewUpdateVolumeRequest() // UpdateVolumeRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesUpdate(context.Background(), appName, volumeId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md) | Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

