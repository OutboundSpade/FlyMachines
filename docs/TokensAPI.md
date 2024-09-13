# \TokensAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokensRequestKms**](TokensAPI.md#TokensRequestKms) | **Post** /tokens/kms | Request a Petsem token for accessing KMS
[**TokensRequestOIDC**](TokensAPI.md#TokensRequestOIDC) | **Post** /tokens/oidc | Request an OIDC token



## TokensRequestKms

> string TokensRequestKms(ctx).Execute()

Request a Petsem token for accessing KMS



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
	resp, r, err := apiClient.TokensAPI.TokensRequestKms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.TokensRequestKms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensRequestKms`: string
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.TokensRequestKms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokensRequestKmsRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensRequestOIDC

> string TokensRequestOIDC(ctx).Request(request).Execute()

Request an OIDC token



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
	request := *openapiclient.NewCreateOIDCTokenRequest() // CreateOIDCTokenRequest | Optional request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.TokensRequestOIDC(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.TokensRequestOIDC``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensRequestOIDC`: string
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.TokensRequestOIDC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensRequestOIDCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateOIDCTokenRequest**](CreateOIDCTokenRequest.md) | Optional request body | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

