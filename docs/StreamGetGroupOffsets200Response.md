# StreamGetGroupOffsets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offsets** | Pointer to [**[]StreamOffsetCommit**](StreamOffsetCommit.md) |  | [optional] 

## Methods

### NewStreamGetGroupOffsets200Response

`func NewStreamGetGroupOffsets200Response() *StreamGetGroupOffsets200Response`

NewStreamGetGroupOffsets200Response instantiates a new StreamGetGroupOffsets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGetGroupOffsets200ResponseWithDefaults

`func NewStreamGetGroupOffsets200ResponseWithDefaults() *StreamGetGroupOffsets200Response`

NewStreamGetGroupOffsets200ResponseWithDefaults instantiates a new StreamGetGroupOffsets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsets

`func (o *StreamGetGroupOffsets200Response) GetOffsets() []StreamOffsetCommit`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *StreamGetGroupOffsets200Response) GetOffsetsOk() (*[]StreamOffsetCommit, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *StreamGetGroupOffsets200Response) SetOffsets(v []StreamOffsetCommit)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *StreamGetGroupOffsets200Response) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


