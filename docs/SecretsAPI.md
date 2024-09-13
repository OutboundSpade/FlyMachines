# \SecretsAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretCreate**](SecretsAPI.md#SecretCreate) | **Post** /apps/{app_name}/secrets/{secret_label}/type/{secret_type} | Create Secret
[**SecretDelete**](SecretsAPI.md#SecretDelete) | **Delete** /apps/{app_name}/secrets/{secret_label} | Destroy Secret
[**SecretGenerate**](SecretsAPI.md#SecretGenerate) | **Post** /apps/{app_name}/secrets/{secret_label}/type/{secret_type}/generate | Generate Secret
[**SecretsList**](SecretsAPI.md#SecretsList) | **Get** /apps/{app_name}/secrets | List App secrets



## SecretCreate

> SecretCreate(ctx).Request(request).Execute()

Create Secret

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
	request := *openapiclient.NewCreateSecretRequest() // CreateSecretRequest | secret body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretCreate(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateSecretRequest**](CreateSecretRequest.md) | secret body | 

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


## SecretDelete

> SecretDelete(ctx, appName, secretLabel).Execute()

Destroy Secret

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
	secretLabel := "secretLabel_example" // string | App Secret Label

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretDelete(context.Background(), appName, secretLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretLabel** | **string** | App Secret Label | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretDeleteRequest struct via the builder pattern


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


## SecretGenerate

> SecretGenerate(ctx).Execute()

Generate Secret

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretGenerate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretGenerateRequest struct via the builder pattern


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


## SecretsList

> []ListSecret SecretsList(ctx).Execute()

List App secrets

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsList`: []ListSecret
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsListRequest struct via the builder pattern


### Return type

[**[]ListSecret**](ListSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

