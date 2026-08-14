# FetchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **int32** | Batch is the most messages to return. 0 or less means 1; anything above 100 is clamped to 100. | [optional] 
**Name** | Pointer to **string** | Name is the consumer, from the path. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream, from the path. | [optional] 
**WaitMs** | Pointer to **int32** | WaitMs is how long to wait for the batch to fill before answering with what arrived. 0 or less means the default of 5000; clamped to 30000. | [optional] 

## Methods

### NewFetchQuery

`func NewFetchQuery() *FetchQuery`

NewFetchQuery instantiates a new FetchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchQueryWithDefaults

`func NewFetchQueryWithDefaults() *FetchQuery`

NewFetchQueryWithDefaults instantiates a new FetchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *FetchQuery) GetBatch() int32`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *FetchQuery) GetBatchOk() (*int32, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *FetchQuery) SetBatch(v int32)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *FetchQuery) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetName

`func (o *FetchQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FetchQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FetchQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FetchQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStream

`func (o *FetchQuery) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *FetchQuery) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *FetchQuery) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *FetchQuery) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetWaitMs

`func (o *FetchQuery) GetWaitMs() int32`

GetWaitMs returns the WaitMs field if non-nil, zero value otherwise.

### GetWaitMsOk

`func (o *FetchQuery) GetWaitMsOk() (*int32, bool)`

GetWaitMsOk returns a tuple with the WaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMs

`func (o *FetchQuery) SetWaitMs(v int32)`

SetWaitMs sets WaitMs field to given value.

### HasWaitMs

`func (o *FetchQuery) HasWaitMs() bool`

HasWaitMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


