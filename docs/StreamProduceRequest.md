# StreamProduceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]StreamProduceRequestRecordsInner**](StreamProduceRequestRecordsInner.md) |  | 

## Methods

### NewStreamProduceRequest

`func NewStreamProduceRequest(records []StreamProduceRequestRecordsInner, ) *StreamProduceRequest`

NewStreamProduceRequest instantiates a new StreamProduceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamProduceRequestWithDefaults

`func NewStreamProduceRequestWithDefaults() *StreamProduceRequest`

NewStreamProduceRequestWithDefaults instantiates a new StreamProduceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *StreamProduceRequest) GetRecords() []StreamProduceRequestRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *StreamProduceRequest) GetRecordsOk() (*[]StreamProduceRequestRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *StreamProduceRequest) SetRecords(v []StreamProduceRequestRecordsInner)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


