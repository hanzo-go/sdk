# ConsoleCreateDatasetItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | **string** |  | 
**Input** | Pointer to **interface{}** |  | [optional] 
**ExpectedOutput** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**SourceTraceId** | Pointer to **string** |  | [optional] 
**SourceObservationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Upsert key. Must be unique within the project. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleCreateDatasetItemRequest

`func NewConsoleCreateDatasetItemRequest(datasetName string, ) *ConsoleCreateDatasetItemRequest`

NewConsoleCreateDatasetItemRequest instantiates a new ConsoleCreateDatasetItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateDatasetItemRequestWithDefaults

`func NewConsoleCreateDatasetItemRequestWithDefaults() *ConsoleCreateDatasetItemRequest`

NewConsoleCreateDatasetItemRequestWithDefaults instantiates a new ConsoleCreateDatasetItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *ConsoleCreateDatasetItemRequest) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ConsoleCreateDatasetItemRequest) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ConsoleCreateDatasetItemRequest) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.


### GetInput

`func (o *ConsoleCreateDatasetItemRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConsoleCreateDatasetItemRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConsoleCreateDatasetItemRequest) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConsoleCreateDatasetItemRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConsoleCreateDatasetItemRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConsoleCreateDatasetItemRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetExpectedOutput

`func (o *ConsoleCreateDatasetItemRequest) GetExpectedOutput() interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *ConsoleCreateDatasetItemRequest) GetExpectedOutputOk() (*interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *ConsoleCreateDatasetItemRequest) SetExpectedOutput(v interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *ConsoleCreateDatasetItemRequest) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### SetExpectedOutputNil

`func (o *ConsoleCreateDatasetItemRequest) SetExpectedOutputNil(b bool)`

 SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil

### UnsetExpectedOutput
`func (o *ConsoleCreateDatasetItemRequest) UnsetExpectedOutput()`

UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil
### GetMetadata

`func (o *ConsoleCreateDatasetItemRequest) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleCreateDatasetItemRequest) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleCreateDatasetItemRequest) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleCreateDatasetItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ConsoleCreateDatasetItemRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ConsoleCreateDatasetItemRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSourceTraceId

`func (o *ConsoleCreateDatasetItemRequest) GetSourceTraceId() string`

GetSourceTraceId returns the SourceTraceId field if non-nil, zero value otherwise.

### GetSourceTraceIdOk

`func (o *ConsoleCreateDatasetItemRequest) GetSourceTraceIdOk() (*string, bool)`

GetSourceTraceIdOk returns a tuple with the SourceTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTraceId

`func (o *ConsoleCreateDatasetItemRequest) SetSourceTraceId(v string)`

SetSourceTraceId sets SourceTraceId field to given value.

### HasSourceTraceId

`func (o *ConsoleCreateDatasetItemRequest) HasSourceTraceId() bool`

HasSourceTraceId returns a boolean if a field has been set.

### GetSourceObservationId

`func (o *ConsoleCreateDatasetItemRequest) GetSourceObservationId() string`

GetSourceObservationId returns the SourceObservationId field if non-nil, zero value otherwise.

### GetSourceObservationIdOk

`func (o *ConsoleCreateDatasetItemRequest) GetSourceObservationIdOk() (*string, bool)`

GetSourceObservationIdOk returns a tuple with the SourceObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObservationId

`func (o *ConsoleCreateDatasetItemRequest) SetSourceObservationId(v string)`

SetSourceObservationId sets SourceObservationId field to given value.

### HasSourceObservationId

`func (o *ConsoleCreateDatasetItemRequest) HasSourceObservationId() bool`

HasSourceObservationId returns a boolean if a field has been set.

### GetId

`func (o *ConsoleCreateDatasetItemRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleCreateDatasetItemRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleCreateDatasetItemRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleCreateDatasetItemRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ConsoleCreateDatasetItemRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsoleCreateDatasetItemRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsoleCreateDatasetItemRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConsoleCreateDatasetItemRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


