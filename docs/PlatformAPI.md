# \PlatformAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformPlacementsPost**](PlatformAPI.md#PlatformPlacementsPost) | **Post** /platform/placements | Get Placements
[**PlatformRegionsGet**](PlatformAPI.md#PlatformRegionsGet) | **Get** /platform/regions | Get Regions



## PlatformPlacementsPost

> MainGetPlacementsResponse PlatformPlacementsPost(ctx).Request(request).Execute()

Get Placements



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
	request := *openapiclient.NewMainGetPlacementsRequest("personal") // MainGetPlacementsRequest | Get placements request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformPlacementsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPlacementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformPlacementsPost`: MainGetPlacementsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformPlacementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPlacementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MainGetPlacementsRequest**](MainGetPlacementsRequest.md) | Get placements request | 

### Return type

[**MainGetPlacementsResponse**](MainGetPlacementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformRegionsGet

> MainRegionResponse PlatformRegionsGet(ctx).Size(size).CpuKind(cpuKind).MemoryMb(memoryMb).Cpus(cpus).Gpus(gpus).GpuKind(gpuKind).Execute()

Get Regions



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
	size := "size_example" // string | guest machine size preset. default performance-1x (optional)
	cpuKind := "cpuKind_example" // string | guest CPU kind (optional)
	memoryMb := int32(56) // int32 | guest memory in megabytes (optional)
	cpus := int32(56) // int32 | guest CPU count (optional)
	gpus := int32(56) // int32 | guest GPU count (optional)
	gpuKind := "gpuKind_example" // string | guest GPU kind (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformRegionsGet(context.Background()).Size(size).CpuKind(cpuKind).MemoryMb(memoryMb).Cpus(cpus).Gpus(gpus).GpuKind(gpuKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformRegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformRegionsGet`: MainRegionResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformRegionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformRegionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **string** | guest machine size preset. default performance-1x | 
 **cpuKind** | **string** | guest CPU kind | 
 **memoryMb** | **int32** | guest memory in megabytes | 
 **cpus** | **int32** | guest CPU count | 
 **gpus** | **int32** | guest GPU count | 
 **gpuKind** | **string** | guest GPU kind | 

### Return type

[**MainRegionResponse**](MainRegionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

