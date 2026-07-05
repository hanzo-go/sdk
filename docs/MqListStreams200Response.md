# MqListStreams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | Pointer to [**[]MqStream**](MqStream.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListStreams200Response

`func NewMqListStreams200Response() *MqListStreams200Response`

NewMqListStreams200Response instantiates a new MqListStreams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListStreams200ResponseWithDefaults

`func NewMqListStreams200ResponseWithDefaults() *MqListStreams200Response`

NewMqListStreams200ResponseWithDefaults instantiates a new MqListStreams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *MqListStreams200Response) GetStreams() []MqStream`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *MqListStreams200Response) GetStreamsOk() (*[]MqStream, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *MqListStreams200Response) SetStreams(v []MqStream)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *MqListStreams200Response) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTotal

`func (o *MqListStreams200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListStreams200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListStreams200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListStreams200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


