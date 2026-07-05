# StreamConsumeMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]StreamConsumeRecord**](StreamConsumeRecord.md) |  | [optional] 
**NextOffset** | Pointer to **int32** |  | [optional] 

## Methods

### NewStreamConsumeMessages200Response

`func NewStreamConsumeMessages200Response() *StreamConsumeMessages200Response`

NewStreamConsumeMessages200Response instantiates a new StreamConsumeMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConsumeMessages200ResponseWithDefaults

`func NewStreamConsumeMessages200ResponseWithDefaults() *StreamConsumeMessages200Response`

NewStreamConsumeMessages200ResponseWithDefaults instantiates a new StreamConsumeMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *StreamConsumeMessages200Response) GetRecords() []StreamConsumeRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *StreamConsumeMessages200Response) GetRecordsOk() (*[]StreamConsumeRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *StreamConsumeMessages200Response) SetRecords(v []StreamConsumeRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *StreamConsumeMessages200Response) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetNextOffset

`func (o *StreamConsumeMessages200Response) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *StreamConsumeMessages200Response) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *StreamConsumeMessages200Response) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *StreamConsumeMessages200Response) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


