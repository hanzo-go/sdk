# CloudBusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the payload as UTF-8 text. | [optional] 
**Headers** | Pointer to **map[string][]string** | Headers are the message&#39;s headers, when it carries any. | [optional] 
**Seq** | Pointer to **int32** | Seq is the message&#39;s stream sequence — fetched messages only. | [optional] 
**Subject** | Pointer to **string** | Subject is the message&#39;s subject in the org&#39;s own namespace. | [optional] 
**Time** | Pointer to **string** | Time is when the stream stored the message, RFC3339 — fetched messages only. | [optional] 

## Methods

### NewCloudBusMessage

`func NewCloudBusMessage() *CloudBusMessage`

NewCloudBusMessage instantiates a new CloudBusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBusMessageWithDefaults

`func NewCloudBusMessageWithDefaults() *CloudBusMessage`

NewCloudBusMessageWithDefaults instantiates a new CloudBusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudBusMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudBusMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudBusMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CloudBusMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *CloudBusMessage) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CloudBusMessage) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CloudBusMessage) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CloudBusMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSeq

`func (o *CloudBusMessage) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *CloudBusMessage) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *CloudBusMessage) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *CloudBusMessage) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSubject

`func (o *CloudBusMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudBusMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudBusMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudBusMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTime

`func (o *CloudBusMessage) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CloudBusMessage) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CloudBusMessage) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CloudBusMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


