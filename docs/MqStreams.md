# MqStreams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | Pointer to [**[]MqStream**](MqStream.md) | Streams is the page, ordered by name. | [optional] 
**Total** | Pointer to **int64** | Total is the org&#39;s stream count before paging. | [optional] 

## Methods

### NewMqStreams

`func NewMqStreams() *MqStreams`

NewMqStreams instantiates a new MqStreams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStreamsWithDefaults

`func NewMqStreamsWithDefaults() *MqStreams`

NewMqStreamsWithDefaults instantiates a new MqStreams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *MqStreams) GetStreams() []MqStream`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *MqStreams) GetStreamsOk() (*[]MqStream, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *MqStreams) SetStreams(v []MqStream)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *MqStreams) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTotal

`func (o *MqStreams) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqStreams) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqStreams) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqStreams) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


