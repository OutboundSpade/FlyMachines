# \SecretsAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretCreate**](SecretsAPI.md#SecretCreate) | **Post** /apps/{app_name}/secrets/{secret_name} | Create or update Secret
[**SecretDelete**](SecretsAPI.md#SecretDelete) | **Delete** /apps/{app_name}/secrets/{secret_name} | Delete an app secret
[**SecretGet**](SecretsAPI.md#SecretGet) | **Get** /apps/{app_name}/secrets/{secret_name} | Get an app secret
[**SecretkeyDecrypt**](SecretsAPI.md#SecretkeyDecrypt) | **Post** /apps/{app_name}/secretkeys/{secret_name}/decrypt | Decrypt with a secret key
[**SecretkeyDelete**](SecretsAPI.md#SecretkeyDelete) | **Delete** /apps/{app_name}/secretkeys/{secret_name} | Delete an app&#39;s secret key
[**SecretkeyEncrypt**](SecretsAPI.md#SecretkeyEncrypt) | **Post** /apps/{app_name}/secretkeys/{secret_name}/encrypt | Encrypt with a secret key
[**SecretkeyGenerate**](SecretsAPI.md#SecretkeyGenerate) | **Post** /apps/{app_name}/secretkeys/{secret_name}/generate | Generate a random secret key
[**SecretkeyGet**](SecretsAPI.md#SecretkeyGet) | **Get** /apps/{app_name}/secretkeys/{secret_name} | Get an app&#39;s secret key
[**SecretkeySet**](SecretsAPI.md#SecretkeySet) | **Post** /apps/{app_name}/secretkeys/{secret_name} | Create or update a secret key
[**SecretkeySign**](SecretsAPI.md#SecretkeySign) | **Post** /apps/{app_name}/secretkeys/{secret_name}/sign | Sign with a secret key
[**SecretkeyVerify**](SecretsAPI.md#SecretkeyVerify) | **Post** /apps/{app_name}/secretkeys/{secret_name}/verify | Verify with a secret key
[**SecretkeysList**](SecretsAPI.md#SecretkeysList) | **Get** /apps/{app_name}/secretkeys | List secret keys belonging to an app
[**SecretsList**](SecretsAPI.md#SecretsList) | **Get** /apps/{app_name}/secrets | List app secrets belonging to an app



## SecretCreate

> SetAppSecretResponse SecretCreate(ctx, appName, secretName).Request(request).Execute()

Create or update Secret

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
	secretName := "secretName_example" // string | App secret name
	request := *openapiclient.NewSetAppSecretRequest() // SetAppSecretRequest | Create app secret request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretCreate(context.Background(), appName, secretName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretCreate`: SetAppSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | App secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SetAppSecretRequest**](SetAppSecretRequest.md) | Create app secret request | 

### Return type

[**SetAppSecretResponse**](SetAppSecretResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretDelete

> SecretDelete(ctx, appName, secretName).Execute()

Delete an app secret

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
	secretName := "secretName_example" // string | App secret name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretDelete(context.Background(), appName, secretName).Execute()
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
**secretName** | **string** | App secret name | 

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


## SecretGet

> AppSecret SecretGet(ctx, appName, secretName).Version(version).ShowSecrets(showSecrets).Execute()

Get an app secret

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
	secretName := "secretName_example" // string | App secret name
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)
	showSecrets := true // bool | Show the secret value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretGet(context.Background(), appName, secretName).Version(version).ShowSecrets(showSecrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretGet`: AppSecret
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | App secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 
 **showSecrets** | **bool** | Show the secret value. | 

### Return type

[**AppSecret**](AppSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeyDecrypt

> DecryptSecretkeyResponse SecretkeyDecrypt(ctx, appName, secretName).Request(request).Version(version).Execute()

Decrypt with a secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewDecryptSecretkeyRequest() // DecryptSecretkeyRequest | Decrypt with secret key request
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeyDecrypt(context.Background(), appName, secretName).Request(request).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyDecrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeyDecrypt`: DecryptSecretkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeyDecrypt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyDecryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**DecryptSecretkeyRequest**](DecryptSecretkeyRequest.md) | Decrypt with secret key request | 
 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 

### Return type

[**DecryptSecretkeyResponse**](DecryptSecretkeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeyDelete

> SecretkeyDelete(ctx, appName, secretName).Execute()

Delete an app's secret key

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
	secretName := "secretName_example" // string | Secret key name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretkeyDelete(context.Background(), appName, secretName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyDeleteRequest struct via the builder pattern


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


## SecretkeyEncrypt

> EncryptSecretkeyResponse SecretkeyEncrypt(ctx, appName, secretName).Request(request).Version(version).Execute()

Encrypt with a secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewEncryptSecretkeyRequest() // EncryptSecretkeyRequest | Encrypt with secret key request
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeyEncrypt(context.Background(), appName, secretName).Request(request).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeyEncrypt`: EncryptSecretkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeyEncrypt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**EncryptSecretkeyRequest**](EncryptSecretkeyRequest.md) | Encrypt with secret key request | 
 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 

### Return type

[**EncryptSecretkeyResponse**](EncryptSecretkeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeyGenerate

> SetSecretkeyResponse SecretkeyGenerate(ctx, appName, secretName).Request(request).Execute()

Generate a random secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewSetSecretkeyRequest() // SetSecretkeyRequest | generate secret key request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeyGenerate(context.Background(), appName, secretName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeyGenerate`: SetSecretkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeyGenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SetSecretkeyRequest**](SetSecretkeyRequest.md) | generate secret key request | 

### Return type

[**SetSecretkeyResponse**](SetSecretkeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeyGet

> SecretKey SecretkeyGet(ctx, appName, secretName).Version(version).Execute()

Get an app's secret key

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
	secretName := "secretName_example" // string | Secret key name
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeyGet(context.Background(), appName, secretName).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeyGet`: SecretKey
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 

### Return type

[**SecretKey**](SecretKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeySet

> SetSecretkeyResponse SecretkeySet(ctx, appName, secretName).Request(request).Execute()

Create or update a secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewSetSecretkeyRequest() // SetSecretkeyRequest | Create secret key request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeySet(context.Background(), appName, secretName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeySet`: SetSecretkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SetSecretkeyRequest**](SetSecretkeyRequest.md) | Create secret key request | 

### Return type

[**SetSecretkeyResponse**](SetSecretkeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeySign

> SignSecretkeyResponse SecretkeySign(ctx, appName, secretName).Request(request).Version(version).Execute()

Sign with a secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewSignSecretkeyRequest() // SignSecretkeyRequest | Sign with secret key request
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeySign(context.Background(), appName, secretName).Request(request).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeySign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeySign`: SignSecretkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeySign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeySignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SignSecretkeyRequest**](SignSecretkeyRequest.md) | Sign with secret key request | 
 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 

### Return type

[**SignSecretkeyResponse**](SignSecretkeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretkeyVerify

> SecretkeyVerify(ctx, appName, secretName).Request(request).Version(version).Execute()

Verify with a secret key

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
	secretName := "secretName_example" // string | Secret key name
	request := *openapiclient.NewVerifySecretkeyRequest() // VerifySecretkeyRequest | Verify with secret key request
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretkeyVerify(context.Background(), appName, secretName).Request(request).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeyVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**secretName** | **string** | Secret key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeyVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**VerifySecretkeyRequest**](VerifySecretkeyRequest.md) | Verify with secret key request | 
 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 

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


## SecretkeysList

> SecretKeys SecretkeysList(ctx, appName).Version(version).Types(types).Execute()

List secret keys belonging to an app

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
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)
	types := "types_example" // string | Comma-seperated list of secret keys to list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretkeysList(context.Background(), appName).Version(version).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretkeysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretkeysList`: SecretKeys
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretkeysList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretkeysListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 
 **types** | **string** | Comma-seperated list of secret keys to list | 

### Return type

[**SecretKeys**](SecretKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsList

> AppSecrets SecretsList(ctx, appName).Version(version).ShowSecrets(showSecrets).Execute()

List app secrets belonging to an app

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
	version := "version_example" // string | Minimum secrets version to return. Returned when setting a new secret (optional)
	showSecrets := true // bool | Show the secret values. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretsList(context.Background(), appName).Version(version).ShowSecrets(showSecrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsList`: AppSecrets
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Minimum secrets version to return. Returned when setting a new secret | 
 **showSecrets** | **bool** | Show the secret values. | 

### Return type

[**AppSecrets**](AppSecrets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

