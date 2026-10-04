# PubsubBusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the request payload, carried verbatim as UTF-8 text. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers are optional request headers, one value per name. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject a responder listens on, in the org&#39;s namespace. | [optional] 
**TimeoutMs** | Pointer to **int64** | TimeoutMs bounds the wait for a reply. 0 or less means the default of 5000; anything above 30000 is clamped to 30000. | [optional] 

## Methods

### NewPubsubBusRequest

`func NewPubsubBusRequest() *PubsubBusRequest`

NewPubsubBusRequest instantiates a new PubsubBusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubBusRequestWithDefaults

`func NewPubsubBusRequestWithDefaults() *PubsubBusRequest`

NewPubsubBusRequestWithDefaults instantiates a new PubsubBusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PubsubBusRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PubsubBusRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PubsubBusRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PubsubBusRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *PubsubBusRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PubsubBusRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PubsubBusRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PubsubBusRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubject

`func (o *PubsubBusRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PubsubBusRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PubsubBusRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PubsubBusRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *PubsubBusRequest) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *PubsubBusRequest) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *PubsubBusRequest) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *PubsubBusRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


