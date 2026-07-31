# CloudBusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the request payload, carried verbatim as UTF-8 text. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers are optional request headers, one value per name. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject a responder listens on, in the org&#39;s namespace. | [optional] 
**TimeoutMs** | Pointer to **int32** | TimeoutMs bounds the wait for a reply. 0 or less means the default of 5000; anything above 30000 is clamped to 30000. | [optional] 

## Methods

### NewCloudBusRequest

`func NewCloudBusRequest() *CloudBusRequest`

NewCloudBusRequest instantiates a new CloudBusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBusRequestWithDefaults

`func NewCloudBusRequestWithDefaults() *CloudBusRequest`

NewCloudBusRequestWithDefaults instantiates a new CloudBusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudBusRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudBusRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudBusRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CloudBusRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *CloudBusRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CloudBusRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CloudBusRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CloudBusRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubject

`func (o *CloudBusRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudBusRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudBusRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudBusRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *CloudBusRequest) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *CloudBusRequest) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *CloudBusRequest) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *CloudBusRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


