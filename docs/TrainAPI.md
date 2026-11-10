# \TrainAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CapabilitiesCapture**](TrainAPI.md#CapabilitiesCapture) | **Get** /v1/capabilities/capture | Answers whether the org&#39;s Kai decisions keep their states for its capabilities to learn from.
[**CapabilitiesCorrections**](TrainAPI.md#CapabilitiesCorrections) | **Get** /v1/capabilities/corrections | Answers every correction of the org&#39;s decisions, newest first, each captured (it can teach) or missing (its decision was made with capture off), and the questions they ask, each counted: what a capability is made from.
[**CapabilitiesCreate**](TrainAPI.md#CapabilitiesCreate) | **Post** /v1/capabilities | Makes a capability in draft from captured corrections of one question: that question (its key, type and labels) is the route it will be installed behind.
[**CapabilitiesDelete**](TrainAPI.md#CapabilitiesDelete) | **Delete** /v1/capabilities/{name} | Deletes a capability and its versions; their artifacts go with the org&#39;s retention.
[**CapabilitiesDisable**](TrainAPI.md#CapabilitiesDisable) | **Post** /v1/capabilities/{name}/disable | Takes a capability&#39;s installed version out of the org&#39;s published set: its questions are answered by base Kai, bit for bit, from the next decision.
[**CapabilitiesGet**](TrainAPI.md#CapabilitiesGet) | **Get** /v1/capabilities/{name} | Answers one capability: its versions with their evaluations, newest first, and what its installed version did live.
[**CapabilitiesList**](TrainAPI.md#CapabilitiesList) | **Get** /v1/capabilities | Answers the org&#39;s capabilities by name, each with its state, its latest and installed versions, and what the installed one did live.
[**CapabilitiesPublish**](TrainAPI.md#CapabilitiesPublish) | **Post** /v1/capabilities/{name}/publish | Installs a version in the org&#39;s published set, where ai attaches it to the org&#39;s decisions: the version&#39;s evaluation must have passed, and it must have been taught over the base Kai serves now.
[**CapabilitiesRollback**](TrainAPI.md#CapabilitiesRollback) | **Post** /v1/capabilities/{name}/rollback | Puts the version installed before the installed one back in the org&#39;s published set: exactly that version&#39;s artifact, nothing retrained.
[**CapabilitiesSetCapture**](TrainAPI.md#CapabilitiesSetCapture) | **Put** /v1/capabilities/capture | Turns capture on or off for the org&#39;s Kai decisions from the next one.
[**CapabilitiesTrain**](TrainAPI.md#CapabilitiesTrain) | **Post** /v1/capabilities/{name}/train | Teaches a capability a new version from its corrections, on the CPU over the states its org&#39;s decisions captured: no row is encoded again.
[**DeleteTrainArtifactsBySha256**](TrainAPI.md#DeleteTrainArtifactsBySha256) | **Delete** /v1/train/artifacts/{sha256} | Deletes one of the org&#39;s objects — an upload or a job&#39;s output — from the store: every artifact naming it then reads deleted, and its storage is no longer billed.
[**DeleteTrainClientsById**](TrainAPI.md#DeleteTrainClientsById) | **Delete** /v1/train/clients/{id} | Drops one of the org&#39;s clients and frees its memory on the engine.
[**GetTrainArtifacts**](TrainAPI.md#GetTrainArtifacts) | **Get** /v1/train/artifacts | Answers the org&#39;s objects — its uploads and its jobs&#39; outputs, one per sha256 — with what produced each, whether it is published and when retention deletes it.
[**GetTrainArtifactsBySha256**](TrainAPI.md#GetTrainArtifactsBySha256) | **Get** /v1/train/artifacts/{sha256} | Answers one of the org&#39;s objects, as the list does, with an address that fetches its bytes for ten minutes once it is stored: how an executor reads an &#x60;artifact:&lt;sha256&gt;&#x60; dataset its job names, and how the org downloads what it uploaded.
[**GetTrainClients**](TrainAPI.md#GetTrainClients) | **Get** /v1/train/clients | Answers the org&#39;s live clients as the engine reports them.
[**GetTrainClientsById**](TrainAPI.md#GetTrainClientsById) | **Get** /v1/train/clients/{id} | Answers one of the org&#39;s clients with its loss history.
[**GetTrainJobs**](TrainAPI.md#GetTrainJobs) | **Get** /v1/train/jobs | Answers the org&#39;s jobs, newest first, narrowed by ?status&#x3D;.
[**GetTrainJobsById**](TrainAPI.md#GetTrainJobsById) | **Get** /v1/train/jobs/{id} | Answers one job: what it asked for, its status and tasks, its artifacts, its result and verdict, and whether it was published.
[**GetTrainJobsByIdArtifacts**](TrainAPI.md#GetTrainJobsByIdArtifacts) | **Get** /v1/train/jobs/{id}/artifacts | Answers a job&#39;s artifacts — checkpoint, capability, lora, basis or merged, each named by the SHA-256 of its bytes — with a download address that lives ten minutes for each one stored.
[**GetTrainJobsByIdEvents**](TrainAPI.md#GetTrainJobsByIdEvents) | **Get** /v1/train/jobs/{id}/events | Answers a job&#39;s events after a cursor — status changes, the coordinator&#39;s address, metrics, log lines, artifacts and the result — waiting up to &#x60;wait&#x60; seconds for the first new one, so a caller following a run holds one request at a time.
[**GetTrainJobsByIdMetrics**](TrainAPI.md#GetTrainJobsByIdMetrics) | **Get** /v1/train/jobs/{id}/metrics | Answers a job&#39;s metric series — loss, throughput and each validation suite&#39;s numbers as the executor reported them — one per name, each point [step, value, unix seconds].
[**GetTrainModels**](TrainAPI.md#GetTrainModels) | **Get** /v1/train/models | Answers each base the caller can train today, in one shape: the base, the methods it trains by (adaptations), where it runs (runs: Hanzo&#39;s devices or the org&#39;s machines) with what one device-second costs there in nano-USD, then the resources that run it (nouns), its executor and devices, and the protections, objective terms, outputs and dataset schemes it takes.
[**PostTrainArtifacts**](TrainAPI.md#PostTrainArtifacts) | **Post** /v1/train/artifacts | Records rows the org is about to upload, named by the SHA-256 of their bytes, and answers the same grant an executor&#39;s artifact gets: a presigned PUT S3 accepts only with that checksum and length, encrypted under the org&#39;s key — a PUT per 64 MiB part over 64 MiB — or stored: true when the org already holds those bytes.
[**PostTrainClients**](TrainAPI.md#PostTrainClients) | **Post** /v1/train/clients | Loads base_model on the engine with a LoRA adapter and answers the client, loading; poll it until ready.
[**PostTrainClientsByIdForwardBackward**](TrainAPI.md#PostTrainClientsByIdForwardBackward) | **Post** /v1/train/clients/{id}/forward_backward | Accumulates the gradients of a batch on the client.
[**PostTrainClientsByIdOptimStep**](TrainAPI.md#PostTrainClientsByIdOptimStep) | **Post** /v1/train/clients/{id}/optim_step | Applies the optimizer to the accumulated gradients.
[**PostTrainClientsByIdSample**](TrainAPI.md#PostTrainClientsByIdSample) | **Post** /v1/train/clients/{id}/sample | Decodes from the client&#39;s current weights.
[**PostTrainClientsByIdSaveWeights**](TrainAPI.md#PostTrainClientsByIdSaveWeights) | **Post** /v1/train/clients/{id}/save_weights | Writes the client&#39;s adapter on the engine.
[**PostTrainJobs**](TrainAPI.md#PostTrainJobs) | **Post** /v1/train/jobs | Creates a job that teaches base_model one capability: the data, the objective, how the trained parameters are represented (adaptation), what they must not damage and by how much (protect), where it runs and how far (resources), the suites it is judged on (evaluation) and what it produces (output).
[**PostTrainJobsByIdArtifacts**](TrainAPI.md#PostTrainJobsByIdArtifacts) | **Post** /v1/train/jobs/{id}/artifacts | Records an output its executor is about to store, named by the SHA-256 of its bytes, and answers a presigned PUT for exactly that object under the org&#39;s prefix — S3 accepts it only with the checksum and encrypts it under the org&#39;s key — or, over 64 MiB, a presigned PUT per 64 MiB part, or stored: true when the org already holds those bytes.
[**PostTrainJobsByIdCancel**](TrainAPI.md#PostTrainJobsByIdCancel) | **Post** /v1/train/jobs/{id}/cancel | Stops a job that has not ended.
[**PostTrainJobsByIdEvents**](TrainAPI.md#PostTrainJobsByIdEvents) | **Post** /v1/train/jobs/{id}/events | ReportTrainTask is an executor&#39;s report for the task it holds: status changes, the lead&#39;s coordinator address, metrics, log lines, stored artifacts and, last, the lead&#39;s result.
[**PostTrainJobsByIdPublish**](TrainAPI.md#PostTrainJobsByIdPublish) | **Post** /v1/train/jobs/{id}/publish | Publishes a job&#39;s output under a name.
[**PostTrainJobsClaim**](TrainAPI.md#PostTrainJobsClaim) | **Post** /v1/train/jobs/claim | ClaimTrainTask is an executor&#39;s long poll for work: it names the machine it runs on, its devices and what its build runs, and is answered, within 25 seconds, one task it can run — a job&#39;s lead first; a join once its lead has said where it coordinates — or 204.



## CapabilitiesCapture

> TrainCapture CapabilitiesCapture(ctx).Execute()

Answers whether the org's Kai decisions keep their states for its capabilities to learn from.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesCapture(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesCapture`: TrainCapture
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesCapture`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesCaptureRequest struct via the builder pattern


### Return type

[**TrainCapture**](TrainCapture.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesCorrections

> TrainCorrectionsOut CapabilitiesCorrections(ctx).Question(question).Execute()

Answers every correction of the org's decisions, newest first, each captured (it can teach) or missing (its decision was made with capture off), and the questions they ask, each counted: what a capability is made from.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	question := "question_example" // string | Question keeps one question's corrections. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesCorrections(context.Background()).Question(question).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesCorrections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesCorrections`: TrainCorrectionsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesCorrections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesCorrectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | **string** | Question keeps one question&#39;s corrections. | 

### Return type

[**TrainCorrectionsOut**](TrainCorrectionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesCreate

> TrainCapability CapabilitiesCreate(ctx).TrainCreateIn(trainCreateIn).Execute()

Makes a capability in draft from captured corrections of one question: that question (its key, type and labels) is the route it will be installed behind.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainCreateIn := *openapiclient.NewTrainCreateIn() // TrainCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesCreate(context.Background()).TrainCreateIn(trainCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesCreate`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainCreateIn** | [**TrainCreateIn**](TrainCreateIn.md) |  | 

### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesDelete

> TrainUnmadeOut CapabilitiesDelete(ctx, name).Execute()

Deletes a capability and its versions; their artifacts go with the org's retention.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesDelete`: TrainUnmadeOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainUnmadeOut**](TrainUnmadeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesDisable

> TrainCapability CapabilitiesDisable(ctx, name).Execute()

Takes a capability's installed version out of the org's published set: its questions are answered by base Kai, bit for bit, from the next decision.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesDisable(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesDisable`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesGet

> TrainCapability CapabilitiesGet(ctx, name).Execute()

Answers one capability: its versions with their evaluations, newest first, and what its installed version did live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesGet`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesList

> TrainCapabilitiesOut CapabilitiesList(ctx).Execute()

Answers the org's capabilities by name, each with its state, its latest and installed versions, and what the installed one did live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesList`: TrainCapabilitiesOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesListRequest struct via the builder pattern


### Return type

[**TrainCapabilitiesOut**](TrainCapabilitiesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesPublish

> TrainCapability CapabilitiesPublish(ctx, name).TrainPublishCapIn(trainPublishCapIn).Execute()

Installs a version in the org's published set, where ai attaches it to the org's decisions: the version's evaluation must have passed, and it must have been taught over the base Kai serves now.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 
	trainPublishCapIn := *openapiclient.NewTrainPublishCapIn() // TrainPublishCapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesPublish(context.Background(), name).TrainPublishCapIn(trainPublishCapIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesPublish`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainPublishCapIn** | [**TrainPublishCapIn**](TrainPublishCapIn.md) |  | 

### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesRollback

> TrainCapability CapabilitiesRollback(ctx, name).Execute()

Puts the version installed before the installed one back in the org's published set: exactly that version's artifact, nothing retrained.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesRollback(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesRollback`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesSetCapture

> TrainCapture CapabilitiesSetCapture(ctx).TrainCaptureIn(trainCaptureIn).Execute()

Turns capture on or off for the org's Kai decisions from the next one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainCaptureIn := *openapiclient.NewTrainCaptureIn() // TrainCaptureIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesSetCapture(context.Background()).TrainCaptureIn(trainCaptureIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesSetCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesSetCapture`: TrainCapture
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesSetCapture`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesSetCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainCaptureIn** | [**TrainCaptureIn**](TrainCaptureIn.md) |  | 

### Return type

[**TrainCapture**](TrainCapture.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesTrain

> TrainCapability CapabilitiesTrain(ctx, name).TrainTrainIn(trainTrainIn).Execute()

Teaches a capability a new version from its corrections, on the CPU over the states its org's decisions captured: no row is encoded again.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 
	trainTrainIn := *openapiclient.NewTrainTrainIn() // TrainTrainIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.CapabilitiesTrain(context.Background(), name).TrainTrainIn(trainTrainIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.CapabilitiesTrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesTrain`: TrainCapability
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.CapabilitiesTrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesTrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainTrainIn** | [**TrainTrainIn**](TrainTrainIn.md) |  | 

### Return type

[**TrainCapability**](TrainCapability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrainArtifactsBySha256

> TrainDeletedOut DeleteTrainArtifactsBySha256(ctx, sha256).Execute()

Deletes one of the org's objects — an upload or a job's output — from the store: every artifact naming it then reads deleted, and its storage is no longer billed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sha256 := "sha256_example" // string | SHA256 names the object, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.DeleteTrainArtifactsBySha256(context.Background(), sha256).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.DeleteTrainArtifactsBySha256``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrainArtifactsBySha256`: TrainDeletedOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.DeleteTrainArtifactsBySha256`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sha256** | **string** | SHA256 names the object, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrainArtifactsBySha256Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainDeletedOut**](TrainDeletedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrainClientsById

> interface{} DeleteTrainClientsById(ctx, id).Execute()

Drops one of the org's clients and frees its memory on the engine.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.DeleteTrainClientsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.DeleteTrainClientsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrainClientsById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.DeleteTrainClientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrainClientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainArtifacts

> TrainObjectsOut GetTrainArtifacts(ctx).Execute()

Answers the org's objects — its uploads and its jobs' outputs, one per sha256 — with what produced each, whether it is published and when retention deletes it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainArtifacts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainArtifacts`: TrainObjectsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainArtifacts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainArtifactsRequest struct via the builder pattern


### Return type

[**TrainObjectsOut**](TrainObjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainArtifactsBySha256

> TrainObject GetTrainArtifactsBySha256(ctx, sha256).Execute()

Answers one of the org's objects, as the list does, with an address that fetches its bytes for ten minutes once it is stored: how an executor reads an `artifact:<sha256>` dataset its job names, and how the org downloads what it uploaded.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sha256 := "sha256_example" // string | SHA256 names the object, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainArtifactsBySha256(context.Background(), sha256).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainArtifactsBySha256``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainArtifactsBySha256`: TrainObject
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainArtifactsBySha256`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sha256** | **string** | SHA256 names the object, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainArtifactsBySha256Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainObject**](TrainObject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainClients

> TrainClientList GetTrainClients(ctx).Execute()

Answers the org's live clients as the engine reports them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainClients(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainClients`: TrainClientList
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainClientsRequest struct via the builder pattern


### Return type

[**TrainClientList**](TrainClientList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainClientsById

> interface{} GetTrainClientsById(ctx, id).Execute()

Answers one of the org's clients with its loss history.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainClientsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainClientsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainClientsById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainClientsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainClientsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainJobs

> TrainJobsOut GetTrainJobs(ctx).Status(status).Limit(limit).Execute()

Answers the org's jobs, newest first, narrowed by ?status=.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	status := "status_example" // string | Status narrows to one status. (optional)
	limit := int64(789) // int64 | Limit bounds the answer; default 50, at most 500. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainJobs(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainJobs`: TrainJobsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status narrows to one status. | 
 **limit** | **int64** | Limit bounds the answer; default 50, at most 500. | 

### Return type

[**TrainJobsOut**](TrainJobsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainJobsById

> TrainJob GetTrainJobsById(ctx, id).Execute()

Answers one job: what it asked for, its status and tasks, its artifacts, its result and verdict, and whether it was published.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainJobsById`: TrainJob
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainJobsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainJob**](TrainJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainJobsByIdArtifacts

> TrainArtifactsOut GetTrainJobsByIdArtifacts(ctx, id).Execute()

Answers a job's artifacts — checkpoint, capability, lora, basis or merged, each named by the SHA-256 of its bytes — with a download address that lives ten minutes for each one stored.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainJobsByIdArtifacts(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainJobsByIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainJobsByIdArtifacts`: TrainArtifactsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainJobsByIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainJobsByIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainArtifactsOut**](TrainArtifactsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainJobsByIdEvents

> TrainEventsOut GetTrainJobsByIdEvents(ctx, id).After(after).Wait(wait).Execute()

Answers a job's events after a cursor — status changes, the coordinator's address, metrics, log lines, artifacts and the result — waiting up to `wait` seconds for the first new one, so a caller following a run holds one request at a time.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the job.
	after := int64(789) // int64 | After is the last seq the caller holds; 0 reads from the start. (optional)
	wait := int64(789) // int64 | Wait is how many seconds to hold the request for a first new event, at most 25. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainJobsByIdEvents(context.Background(), id).After(after).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainJobsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainJobsByIdEvents`: TrainEventsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainJobsByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainJobsByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int64** | After is the last seq the caller holds; 0 reads from the start. | 
 **wait** | **int64** | Wait is how many seconds to hold the request for a first new event, at most 25. | 

### Return type

[**TrainEventsOut**](TrainEventsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainJobsByIdMetrics

> TrainMetricsOut GetTrainJobsByIdMetrics(ctx, id).Execute()

Answers a job's metric series — loss, throughput and each validation suite's numbers as the executor reported them — one per name, each point [step, value, unix seconds].



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainJobsByIdMetrics(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainJobsByIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainJobsByIdMetrics`: TrainMetricsOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainJobsByIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainJobsByIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainMetricsOut**](TrainMetricsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainModels

> TrainCatalog GetTrainModels(ctx).Execute()

Answers each base the caller can train today, in one shape: the base, the methods it trains by (adaptations), where it runs (runs: Hanzo's devices or the org's machines) with what one device-second costs there in nano-USD, then the resources that run it (nouns), its executor and devices, and the protections, objective terms, outputs and dataset schemes it takes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.GetTrainModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.GetTrainModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainModels`: TrainCatalog
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.GetTrainModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainModelsRequest struct via the builder pattern


### Return type

[**TrainCatalog**](TrainCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainArtifacts

> TrainRegisterOut PostTrainArtifacts(ctx).TrainUploadIn(trainUploadIn).Execute()

Records rows the org is about to upload, named by the SHA-256 of their bytes, and answers the same grant an executor's artifact gets: a presigned PUT S3 accepts only with that checksum and length, encrypted under the org's key — a PUT per 64 MiB part over 64 MiB — or stored: true when the org already holds those bytes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainUploadIn := *openapiclient.NewTrainUploadIn() // TrainUploadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainArtifacts(context.Background()).TrainUploadIn(trainUploadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainArtifacts`: TrainRegisterOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainUploadIn** | [**TrainUploadIn**](TrainUploadIn.md) |  | 

### Return type

[**TrainRegisterOut**](TrainRegisterOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainClients

> interface{} PostTrainClients(ctx).TrainCreateClientIn(trainCreateClientIn).Execute()

Loads base_model on the engine with a LoRA adapter and answers the client, loading; poll it until ready.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainCreateClientIn := *openapiclient.NewTrainCreateClientIn() // TrainCreateClientIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainClients(context.Background()).TrainCreateClientIn(trainCreateClientIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainClients`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainCreateClientIn** | [**TrainCreateClientIn**](TrainCreateClientIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainClientsByIdForwardBackward

> interface{} PostTrainClientsByIdForwardBackward(ctx, id).TrainForwardIn(trainForwardIn).Execute()

Accumulates the gradients of a batch on the client.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "tc-…" // string | ID names the client.
	trainForwardIn := *openapiclient.NewTrainForwardIn() // TrainForwardIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainClientsByIdForwardBackward(context.Background(), id).TrainForwardIn(trainForwardIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainClientsByIdForwardBackward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainClientsByIdForwardBackward`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainClientsByIdForwardBackward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainClientsByIdForwardBackwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainForwardIn** | [**TrainForwardIn**](TrainForwardIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainClientsByIdOptimStep

> interface{} PostTrainClientsByIdOptimStep(ctx, id).TrainOptimIn(trainOptimIn).Execute()

Applies the optimizer to the accumulated gradients.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the client.
	trainOptimIn := *openapiclient.NewTrainOptimIn() // TrainOptimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainClientsByIdOptimStep(context.Background(), id).TrainOptimIn(trainOptimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainClientsByIdOptimStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainClientsByIdOptimStep`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainClientsByIdOptimStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainClientsByIdOptimStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainOptimIn** | [**TrainOptimIn**](TrainOptimIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainClientsByIdSample

> interface{} PostTrainClientsByIdSample(ctx, id).TrainSampleIn(trainSampleIn).Execute()

Decodes from the client's current weights.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the client.
	trainSampleIn := *openapiclient.NewTrainSampleIn() // TrainSampleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainClientsByIdSample(context.Background(), id).TrainSampleIn(trainSampleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainClientsByIdSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainClientsByIdSample`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainClientsByIdSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainClientsByIdSampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainSampleIn** | [**TrainSampleIn**](TrainSampleIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainClientsByIdSaveWeights

> interface{} PostTrainClientsByIdSaveWeights(ctx, id).TrainSaveIn(trainSaveIn).Execute()

Writes the client's adapter on the engine.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the client.
	trainSaveIn := *openapiclient.NewTrainSaveIn() // TrainSaveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainClientsByIdSaveWeights(context.Background(), id).TrainSaveIn(trainSaveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainClientsByIdSaveWeights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainClientsByIdSaveWeights`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainClientsByIdSaveWeights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainClientsByIdSaveWeightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainSaveIn** | [**TrainSaveIn**](TrainSaveIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobs

> TrainJob PostTrainJobs(ctx).TrainSpec(trainSpec).Execute()

Creates a job that teaches base_model one capability: the data, the objective, how the trained parameters are represented (adaptation), what they must not damage and by how much (protect), where it runs and how far (resources), the suites it is judged on (evaluation) and what it produces (output).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainSpec := *openapiclient.NewTrainSpec() // TrainSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobs(context.Background()).TrainSpec(trainSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobs`: TrainJob
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainSpec** | [**TrainSpec**](TrainSpec.md) |  | 

### Return type

[**TrainJob**](TrainJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobsByIdArtifacts

> TrainRegisterOut PostTrainJobsByIdArtifacts(ctx, id).TrainRegisterIn(trainRegisterIn).Execute()

Records an output its executor is about to store, named by the SHA-256 of its bytes, and answers a presigned PUT for exactly that object under the org's prefix — S3 accepts it only with the checksum and encrypts it under the org's key — or, over 64 MiB, a presigned PUT per 64 MiB part, or stored: true when the org already holds those bytes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "job_0123456789abcdef" // string | ID names the job, from the path.
	trainRegisterIn := *openapiclient.NewTrainRegisterIn() // TrainRegisterIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobsByIdArtifacts(context.Background(), id).TrainRegisterIn(trainRegisterIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobsByIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobsByIdArtifacts`: TrainRegisterOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobsByIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsByIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainRegisterIn** | [**TrainRegisterIn**](TrainRegisterIn.md) |  | 

### Return type

[**TrainRegisterOut**](TrainRegisterOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobsByIdCancel

> TrainJob PostTrainJobsByIdCancel(ctx, id).Execute()

Stops a job that has not ended.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID names the job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobsByIdCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobsByIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobsByIdCancel`: TrainJob
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobsByIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsByIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrainJob**](TrainJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobsByIdEvents

> TrainReportOut PostTrainJobsByIdEvents(ctx, id).TrainReportIn(trainReportIn).Execute()

ReportTrainTask is an executor's report for the task it holds: status changes, the lead's coordinator address, metrics, log lines, stored artifacts and, last, the lead's result.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "job_0123456789abcdef" // string | ID names the job, from the path.
	trainReportIn := *openapiclient.NewTrainReportIn() // TrainReportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobsByIdEvents(context.Background(), id).TrainReportIn(trainReportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobsByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobsByIdEvents`: TrainReportOut
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobsByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainReportIn** | [**TrainReportIn**](TrainReportIn.md) |  | 

### Return type

[**TrainReportOut**](TrainReportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobsByIdPublish

> TrainJob PostTrainJobsByIdPublish(ctx, id).TrainPublishIn(trainPublishIn).Execute()

Publishes a job's output under a name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "job_0123456789abcdef" // string | ID names the job.
	trainPublishIn := *openapiclient.NewTrainPublishIn() // TrainPublishIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobsByIdPublish(context.Background(), id).TrainPublishIn(trainPublishIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobsByIdPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobsByIdPublish`: TrainJob
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobsByIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsByIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trainPublishIn** | [**TrainPublishIn**](TrainPublishIn.md) |  | 

### Return type

[**TrainJob**](TrainJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainJobsClaim

> TrainAssignment PostTrainJobsClaim(ctx).TrainClaimIn(trainClaimIn).Execute()

ClaimTrainTask is an executor's long poll for work: it names the machine it runs on, its devices and what its build runs, and is answered, within 25 seconds, one task it can run — a job's lead first; a join once its lead has said where it coordinates — or 204.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	trainClaimIn := *openapiclient.NewTrainClaimIn() // TrainClaimIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainAPI.PostTrainJobsClaim(context.Background()).TrainClaimIn(trainClaimIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainAPI.PostTrainJobsClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainJobsClaim`: TrainAssignment
	fmt.Fprintf(os.Stdout, "Response from `TrainAPI.PostTrainJobsClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainJobsClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainClaimIn** | [**TrainClaimIn**](TrainClaimIn.md) |  | 

### Return type

[**TrainAssignment**](TrainAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

