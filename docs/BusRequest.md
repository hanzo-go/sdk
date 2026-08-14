# BusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the request payload, carried verbatim as UTF-8 text. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers are optional request headers, one value per name. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject a responder listens on, in the org&#39;s namespace. | [optional] 
**TimeoutMs** | Pointer to **int32** | TimeoutMs bounds the wait for a reply. 0 or less means the default of 5000; anything above 30000 is clamped to 30000. | [optional] 

## Methods

### NewBusRequest

`func NewBusRequest() *BusRequest`

NewBusRequest instantiates a new BusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusRequestWithDefaults

`func NewBusRequestWithDefaults() *BusRequest`

NewBusRequestWithDefaults instantiates a new BusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BusRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BusRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BusRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BusRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *BusRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BusRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BusRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BusRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubject

`func (o *BusRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BusRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BusRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *BusRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *BusRequest) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *BusRequest) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *BusRequest) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *BusRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


