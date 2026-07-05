# StreamTopicMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Partitions** | Pointer to [**[]StreamTopicMetadataPartitionsInner**](StreamTopicMetadataPartitionsInner.md) |  | [optional] 

## Methods

### NewStreamTopicMetadata

`func NewStreamTopicMetadata() *StreamTopicMetadata`

NewStreamTopicMetadata instantiates a new StreamTopicMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamTopicMetadataWithDefaults

`func NewStreamTopicMetadataWithDefaults() *StreamTopicMetadata`

NewStreamTopicMetadataWithDefaults instantiates a new StreamTopicMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamTopicMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamTopicMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamTopicMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamTopicMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitions

`func (o *StreamTopicMetadata) GetPartitions() []StreamTopicMetadataPartitionsInner`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *StreamTopicMetadata) GetPartitionsOk() (*[]StreamTopicMetadataPartitionsInner, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *StreamTopicMetadata) SetPartitions(v []StreamTopicMetadataPartitionsInner)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *StreamTopicMetadata) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


