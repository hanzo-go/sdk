# MqPublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Target subject. | 
**Data** | **string** | Message payload (base64-encoded for binary data). | 
**Headers** | Pointer to **map[string][]string** | Optional message headers. | [optional] 
**Reply** | Pointer to **string** | Optional reply-to subject. | [optional] 

## Methods

### NewMqPublishRequest

`func NewMqPublishRequest(subject string, data string, ) *MqPublishRequest`

NewMqPublishRequest instantiates a new MqPublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqPublishRequestWithDefaults

`func NewMqPublishRequestWithDefaults() *MqPublishRequest`

NewMqPublishRequestWithDefaults instantiates a new MqPublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MqPublishRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MqPublishRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MqPublishRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetData

`func (o *MqPublishRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MqPublishRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MqPublishRequest) SetData(v string)`

SetData sets Data field to given value.


### GetHeaders

`func (o *MqPublishRequest) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MqPublishRequest) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MqPublishRequest) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MqPublishRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReply

`func (o *MqPublishRequest) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *MqPublishRequest) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *MqPublishRequest) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *MqPublishRequest) HasReply() bool`

HasReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


