# \DynamicAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelFilterPost**](DynamicAPI.md#ModelFilterPost) | **Post** /{model}/filter | Filter entities
[**ModelGet**](DynamicAPI.md#ModelGet) | **Get** /{model} | List and filter entities
[**ModelIdDelete**](DynamicAPI.md#ModelIdDelete) | **Delete** /{model}/{id} | Delete an entity
[**ModelIdGet**](DynamicAPI.md#ModelIdGet) | **Get** /{model}/{id} | Get an entity by ID
[**ModelIdPut**](DynamicAPI.md#ModelIdPut) | **Put** /{model}/{id} | Update an entity
[**ModelPost**](DynamicAPI.md#ModelPost) | **Post** /{model} | Create a new entity



## ModelFilterPost

> QueryFilterResponse ModelFilterPost(ctx, model).Filter(filter).Page(page).PageSize(pageSize).Sort(sort).Execute()

Filter entities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model name
	filter := *openapiclient.NewQueryQueryFilter() // QueryQueryFilter | Filter conditions
	page := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page (optional) (default to 10)
	sort := "sort_example" // string | Sort field and direction (e.g., name:asc,age:desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelFilterPost(context.Background(), model).Filter(filter).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelFilterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelFilterPost`: QueryFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelFilterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**QueryQueryFilter**](QueryQueryFilter.md) | Filter conditions | 
 **page** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Items per page | [default to 10]
 **sort** | **string** | Sort field and direction (e.g., name:asc,age:desc) | 

### Return type

[**QueryFilterResponse**](QueryFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelGet

> QueryFilterResponse ModelGet(ctx, model).Page(page).PageSize(pageSize).Sort(sort).Execute()

List and filter entities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	page := int32(56) // int32 | Page number (optional)
	pageSize := int32(56) // int32 | Items per page (optional)
	sort := "sort_example" // string | Sort field and direction (e.g., name:asc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelGet(context.Background(), model).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelGet`: QueryFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | 
 **pageSize** | **int32** | Items per page | 
 **sort** | **string** | Sort field and direction (e.g., name:asc) | 

### Return type

[**QueryFilterResponse**](QueryFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdDelete

> ApiErrorResponse ModelIdDelete(ctx, model, id).Execute()

Delete an entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdDelete(context.Background(), model, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdDelete`: ApiErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiErrorResponse**](ApiErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdGet

> QueryEntityWithRelations ModelIdGet(ctx, model, id).Execute()

Get an entity by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdGet(context.Background(), model, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdGet`: QueryEntityWithRelations
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueryEntityWithRelations**](QueryEntityWithRelations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdPut

> QueryEntityWithRelations ModelIdPut(ctx, model, id).Entity(entity).Execute()

Update an entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID
	entity := *openapiclient.NewQueryEntityWithRelations() // QueryEntityWithRelations | Entity Data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdPut(context.Background(), model, id).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdPut`: QueryEntityWithRelations
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entity** | [**QueryEntityWithRelations**](QueryEntityWithRelations.md) | Entity Data | 

### Return type

[**QueryEntityWithRelations**](QueryEntityWithRelations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelPost

> QueryEntityWithRelations ModelPost(ctx, model).Entity(entity).Execute()

Create a new entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	entity := *openapiclient.NewQueryEntityWithRelations() // QueryEntityWithRelations | Entity Data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelPost(context.Background(), model).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelPost`: QueryEntityWithRelations
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entity** | [**QueryEntityWithRelations**](QueryEntityWithRelations.md) | Entity Data | 

### Return type

[**QueryEntityWithRelations**](QueryEntityWithRelations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

