# MqRequestReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Target subject. | 
**Data** | **string** | Request payload (base64-encoded for binary data). | 
**Timeout** | Pointer to **string** | Maximum time to wait for a reply (e.g., \&quot;5s\&quot;, \&quot;30s\&quot;). Defaults to \&quot;10s\&quot;.  | [optional] [default to "10s"]
**Headers** | Pointer to **map[string][]string** | Optional message headers. | [optional] 

## Methods

### NewMqRequestReply

`func NewMqRequestReply(subject string, data string, ) *MqRequestReply`

NewMqRequestReply instantiates a new MqRequestReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqRequestReplyWithDefaults

`func NewMqRequestReplyWithDefaults() *MqRequestReply`

NewMqRequestReplyWithDefaults instantiates a new MqRequestReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MqRequestReply) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MqRequestReply) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MqRequestReply) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetData

`func (o *MqRequestReply) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MqRequestReply) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MqRequestReply) SetData(v string)`

SetData sets Data field to given value.


### GetTimeout

`func (o *MqRequestReply) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MqRequestReply) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MqRequestReply) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MqRequestReply) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetHeaders

`func (o *MqRequestReply) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MqRequestReply) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MqRequestReply) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MqRequestReply) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


