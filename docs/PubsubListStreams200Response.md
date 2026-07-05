# PubsubListStreams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | Pointer to [**[]PubsubStreamInfo**](PubsubStreamInfo.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubListStreams200Response

`func NewPubsubListStreams200Response() *PubsubListStreams200Response`

NewPubsubListStreams200Response instantiates a new PubsubListStreams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubListStreams200ResponseWithDefaults

`func NewPubsubListStreams200ResponseWithDefaults() *PubsubListStreams200Response`

NewPubsubListStreams200ResponseWithDefaults instantiates a new PubsubListStreams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *PubsubListStreams200Response) GetStreams() []PubsubStreamInfo`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *PubsubListStreams200Response) GetStreamsOk() (*[]PubsubStreamInfo, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *PubsubListStreams200Response) SetStreams(v []PubsubStreamInfo)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *PubsubListStreams200Response) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTotal

`func (o *PubsubListStreams200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PubsubListStreams200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PubsubListStreams200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PubsubListStreams200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


