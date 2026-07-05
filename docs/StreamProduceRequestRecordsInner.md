# StreamProduceRequestRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Optional partition key | [optional] 
**Value** | Pointer to **string** | Message value | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Partition** | Pointer to **int32** | Target partition (optional, auto-assigned if omitted) | [optional] 

## Methods

### NewStreamProduceRequestRecordsInner

`func NewStreamProduceRequestRecordsInner() *StreamProduceRequestRecordsInner`

NewStreamProduceRequestRecordsInner instantiates a new StreamProduceRequestRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamProduceRequestRecordsInnerWithDefaults

`func NewStreamProduceRequestRecordsInnerWithDefaults() *StreamProduceRequestRecordsInner`

NewStreamProduceRequestRecordsInnerWithDefaults instantiates a new StreamProduceRequestRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *StreamProduceRequestRecordsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StreamProduceRequestRecordsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StreamProduceRequestRecordsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StreamProduceRequestRecordsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *StreamProduceRequestRecordsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StreamProduceRequestRecordsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StreamProduceRequestRecordsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StreamProduceRequestRecordsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHeaders

`func (o *StreamProduceRequestRecordsInner) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *StreamProduceRequestRecordsInner) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *StreamProduceRequestRecordsInner) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *StreamProduceRequestRecordsInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPartition

`func (o *StreamProduceRequestRecordsInner) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *StreamProduceRequestRecordsInner) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *StreamProduceRequestRecordsInner) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *StreamProduceRequestRecordsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


