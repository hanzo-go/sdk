# ConsoleCreateDatasetRunItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetItemId** | **string** |  | 
**TraceId** | **string** |  | 
**ObservationId** | Pointer to **string** |  | [optional] 
**RunName** | **string** |  | 
**RunDescription** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConsoleCreateDatasetRunItemRequest

`func NewConsoleCreateDatasetRunItemRequest(datasetItemId string, traceId string, runName string, ) *ConsoleCreateDatasetRunItemRequest`

NewConsoleCreateDatasetRunItemRequest instantiates a new ConsoleCreateDatasetRunItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateDatasetRunItemRequestWithDefaults

`func NewConsoleCreateDatasetRunItemRequestWithDefaults() *ConsoleCreateDatasetRunItemRequest`

NewConsoleCreateDatasetRunItemRequestWithDefaults instantiates a new ConsoleCreateDatasetRunItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetItemId

`func (o *ConsoleCreateDatasetRunItemRequest) GetDatasetItemId() string`

GetDatasetItemId returns the DatasetItemId field if non-nil, zero value otherwise.

### GetDatasetItemIdOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetDatasetItemIdOk() (*string, bool)`

GetDatasetItemIdOk returns a tuple with the DatasetItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetItemId

`func (o *ConsoleCreateDatasetRunItemRequest) SetDatasetItemId(v string)`

SetDatasetItemId sets DatasetItemId field to given value.


### GetTraceId

`func (o *ConsoleCreateDatasetRunItemRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ConsoleCreateDatasetRunItemRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetObservationId

`func (o *ConsoleCreateDatasetRunItemRequest) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *ConsoleCreateDatasetRunItemRequest) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *ConsoleCreateDatasetRunItemRequest) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetRunName

`func (o *ConsoleCreateDatasetRunItemRequest) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *ConsoleCreateDatasetRunItemRequest) SetRunName(v string)`

SetRunName sets RunName field to given value.


### GetRunDescription

`func (o *ConsoleCreateDatasetRunItemRequest) GetRunDescription() string`

GetRunDescription returns the RunDescription field if non-nil, zero value otherwise.

### GetRunDescriptionOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetRunDescriptionOk() (*string, bool)`

GetRunDescriptionOk returns a tuple with the RunDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDescription

`func (o *ConsoleCreateDatasetRunItemRequest) SetRunDescription(v string)`

SetRunDescription sets RunDescription field to given value.

### HasRunDescription

`func (o *ConsoleCreateDatasetRunItemRequest) HasRunDescription() bool`

HasRunDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleCreateDatasetRunItemRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleCreateDatasetRunItemRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleCreateDatasetRunItemRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleCreateDatasetRunItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


