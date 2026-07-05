# PubsubPublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Subject to publish to (e.g. orders.created) | 
**Data** | **string** | Message payload (string or base64-encoded binary) | 
**Headers** | Pointer to **map[string]string** | Optional message headers | [optional] 
**Reply** | Pointer to **string** | Reply-to subject for request/reply pattern | [optional] 

## Methods

### NewPubsubPublishRequest

`func NewPubsubPublishRequest(subject string, data string, ) *PubsubPublishRequest`

NewPubsubPublishRequest instantiates a new PubsubPublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubPublishRequestWithDefaults

`func NewPubsubPublishRequestWithDefaults() *PubsubPublishRequest`

NewPubsubPublishRequestWithDefaults instantiates a new PubsubPublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PubsubPublishRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PubsubPublishRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PubsubPublishRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetData

`func (o *PubsubPublishRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PubsubPublishRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PubsubPublishRequest) SetData(v string)`

SetData sets Data field to given value.


### GetHeaders

`func (o *PubsubPublishRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PubsubPublishRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PubsubPublishRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PubsubPublishRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReply

`func (o *PubsubPublishRequest) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *PubsubPublishRequest) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *PubsubPublishRequest) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *PubsubPublishRequest) HasReply() bool`

HasReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


