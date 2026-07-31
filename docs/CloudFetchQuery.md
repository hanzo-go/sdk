# CloudFetchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **int32** | Batch is the most messages to return. 0 or less means 1; anything above 100 is clamped to 100. | [optional] 
**Name** | Pointer to **string** | Name is the consumer, from the path. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream, from the path. | [optional] 
**WaitMs** | Pointer to **int32** | WaitMs is how long to wait for the batch to fill before answering with what arrived. 0 or less means the default of 5000; clamped to 30000. | [optional] 

## Methods

### NewCloudFetchQuery

`func NewCloudFetchQuery() *CloudFetchQuery`

NewCloudFetchQuery instantiates a new CloudFetchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFetchQueryWithDefaults

`func NewCloudFetchQueryWithDefaults() *CloudFetchQuery`

NewCloudFetchQueryWithDefaults instantiates a new CloudFetchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *CloudFetchQuery) GetBatch() int32`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *CloudFetchQuery) GetBatchOk() (*int32, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *CloudFetchQuery) SetBatch(v int32)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *CloudFetchQuery) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetName

`func (o *CloudFetchQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudFetchQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudFetchQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudFetchQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStream

`func (o *CloudFetchQuery) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudFetchQuery) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudFetchQuery) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudFetchQuery) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetWaitMs

`func (o *CloudFetchQuery) GetWaitMs() int32`

GetWaitMs returns the WaitMs field if non-nil, zero value otherwise.

### GetWaitMsOk

`func (o *CloudFetchQuery) GetWaitMsOk() (*int32, bool)`

GetWaitMsOk returns a tuple with the WaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMs

`func (o *CloudFetchQuery) SetWaitMs(v int32)`

SetWaitMs sets WaitMs field to given value.

### HasWaitMs

`func (o *CloudFetchQuery) HasWaitMs() bool`

HasWaitMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


