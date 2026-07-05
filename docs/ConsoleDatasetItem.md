# ConsoleDatasetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DatasetId** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**ExpectedOutput** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**SourceTraceId** | Pointer to **string** |  | [optional] 
**SourceObservationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConsoleDatasetItem

`func NewConsoleDatasetItem() *ConsoleDatasetItem`

NewConsoleDatasetItem instantiates a new ConsoleDatasetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleDatasetItemWithDefaults

`func NewConsoleDatasetItemWithDefaults() *ConsoleDatasetItem`

NewConsoleDatasetItemWithDefaults instantiates a new ConsoleDatasetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleDatasetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleDatasetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleDatasetItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleDatasetItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatasetId

`func (o *ConsoleDatasetItem) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *ConsoleDatasetItem) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *ConsoleDatasetItem) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *ConsoleDatasetItem) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetInput

`func (o *ConsoleDatasetItem) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConsoleDatasetItem) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConsoleDatasetItem) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConsoleDatasetItem) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConsoleDatasetItem) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConsoleDatasetItem) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetExpectedOutput

`func (o *ConsoleDatasetItem) GetExpectedOutput() interface{}`

GetExpectedOutput returns the ExpectedOutput field if non-nil, zero value otherwise.

### GetExpectedOutputOk

`func (o *ConsoleDatasetItem) GetExpectedOutputOk() (*interface{}, bool)`

GetExpectedOutputOk returns a tuple with the ExpectedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutput

`func (o *ConsoleDatasetItem) SetExpectedOutput(v interface{})`

SetExpectedOutput sets ExpectedOutput field to given value.

### HasExpectedOutput

`func (o *ConsoleDatasetItem) HasExpectedOutput() bool`

HasExpectedOutput returns a boolean if a field has been set.

### SetExpectedOutputNil

`func (o *ConsoleDatasetItem) SetExpectedOutputNil(b bool)`

 SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil

### UnsetExpectedOutput
`func (o *ConsoleDatasetItem) UnsetExpectedOutput()`

UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil
### GetMetadata

`func (o *ConsoleDatasetItem) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleDatasetItem) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleDatasetItem) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleDatasetItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ConsoleDatasetItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ConsoleDatasetItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSourceTraceId

`func (o *ConsoleDatasetItem) GetSourceTraceId() string`

GetSourceTraceId returns the SourceTraceId field if non-nil, zero value otherwise.

### GetSourceTraceIdOk

`func (o *ConsoleDatasetItem) GetSourceTraceIdOk() (*string, bool)`

GetSourceTraceIdOk returns a tuple with the SourceTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTraceId

`func (o *ConsoleDatasetItem) SetSourceTraceId(v string)`

SetSourceTraceId sets SourceTraceId field to given value.

### HasSourceTraceId

`func (o *ConsoleDatasetItem) HasSourceTraceId() bool`

HasSourceTraceId returns a boolean if a field has been set.

### GetSourceObservationId

`func (o *ConsoleDatasetItem) GetSourceObservationId() string`

GetSourceObservationId returns the SourceObservationId field if non-nil, zero value otherwise.

### GetSourceObservationIdOk

`func (o *ConsoleDatasetItem) GetSourceObservationIdOk() (*string, bool)`

GetSourceObservationIdOk returns a tuple with the SourceObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObservationId

`func (o *ConsoleDatasetItem) SetSourceObservationId(v string)`

SetSourceObservationId sets SourceObservationId field to given value.

### HasSourceObservationId

`func (o *ConsoleDatasetItem) HasSourceObservationId() bool`

HasSourceObservationId returns a boolean if a field has been set.

### GetStatus

`func (o *ConsoleDatasetItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsoleDatasetItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsoleDatasetItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConsoleDatasetItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConsoleDatasetItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsoleDatasetItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsoleDatasetItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsoleDatasetItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConsoleDatasetItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConsoleDatasetItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConsoleDatasetItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConsoleDatasetItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


