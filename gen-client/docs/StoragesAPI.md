# \StoragesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoragesGet**](StoragesAPI.md#StoragesGet) | **Get** /storages | List all storages
[**StoragesIdAttachVmidPost**](StoragesAPI.md#StoragesIdAttachVmidPost) | **Post** /storages/{id}/attach/{vmid} | Attach storage to VM
[**StoragesIdDelete**](StoragesAPI.md#StoragesIdDelete) | **Delete** /storages/{id} | Delete storage by ID
[**StoragesIdDetachVmidPost**](StoragesAPI.md#StoragesIdDetachVmidPost) | **Post** /storages/{id}/detach/{vmid} | Detach storage from VM
[**StoragesIdGet**](StoragesAPI.md#StoragesIdGet) | **Get** /storages/{id} | Get storage by ID
[**StoragesIdPut**](StoragesAPI.md#StoragesIdPut) | **Put** /storages/{id} | Update storage size
[**StoragesPost**](StoragesAPI.md#StoragesPost) | **Post** /storages | Create a new storage volume



## StoragesGet

> map[string]interface{} StoragesGet(ctx).Execute()

List all storages



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesIdAttachVmidPost

> map[string]interface{} StoragesIdAttachVmidPost(ctx, id, vmid).Execute()

Attach storage to VM



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
	id := int32(56) // int32 | Storage ID
	vmid := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesIdAttachVmidPost(context.Background(), id, vmid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesIdAttachVmidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesIdAttachVmidPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesIdAttachVmidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Storage ID | 
**vmid** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesIdAttachVmidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesIdDelete

> map[string]string StoragesIdDelete(ctx, id).Execute()

Delete storage by ID



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
	id := int32(56) // int32 | Storage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesIdDelete`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesIdDetachVmidPost

> map[string]interface{} StoragesIdDetachVmidPost(ctx, id, vmid).Execute()

Detach storage from VM



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
	id := int32(56) // int32 | Storage ID
	vmid := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesIdDetachVmidPost(context.Background(), id, vmid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesIdDetachVmidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesIdDetachVmidPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesIdDetachVmidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Storage ID | 
**vmid** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesIdDetachVmidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesIdGet

> map[string]interface{} StoragesIdGet(ctx, id).Execute()

Get storage by ID



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
	id := int32(56) // int32 | Storage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesIdPut

> map[string]interface{} StoragesIdPut(ctx, id).Storage(storage).Execute()

Update storage size



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
	id := int32(56) // int32 | Storage ID
	storage := *openapiclient.NewModelsStorage() // ModelsStorage | Updated storage data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesIdPut(context.Background(), id).Storage(storage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesIdPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | [**ModelsStorage**](ModelsStorage.md) | Updated storage data | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragesPost

> map[string]interface{} StoragesPost(ctx).Storage(storage).Execute()

Create a new storage volume



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
	storage := *openapiclient.NewModelsStorage() // ModelsStorage | Storage object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoragesAPI.StoragesPost(context.Background()).Storage(storage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoragesAPI.StoragesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragesPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StoragesAPI.StoragesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoragesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storage** | [**ModelsStorage**](ModelsStorage.md) | Storage object | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

