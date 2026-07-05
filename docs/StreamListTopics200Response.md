# StreamListTopics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topics** | Pointer to [**[]StreamTopic**](StreamTopic.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewStreamListTopics200Response

`func NewStreamListTopics200Response() *StreamListTopics200Response`

NewStreamListTopics200Response instantiates a new StreamListTopics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamListTopics200ResponseWithDefaults

`func NewStreamListTopics200ResponseWithDefaults() *StreamListTopics200Response`

NewStreamListTopics200ResponseWithDefaults instantiates a new StreamListTopics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopics

`func (o *StreamListTopics200Response) GetTopics() []StreamTopic`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *StreamListTopics200Response) GetTopicsOk() (*[]StreamTopic, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *StreamListTopics200Response) SetTopics(v []StreamTopic)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *StreamListTopics200Response) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetTotal

`func (o *StreamListTopics200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StreamListTopics200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StreamListTopics200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StreamListTopics200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


