# BusPublish

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the payload, carried verbatim as UTF-8 text (typically JSON). Binary payloads belong on the NATS port. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers are optional message headers, one value per name. A Nats-Msg-Id header is JetStream&#39;s deduplication key: a repeat within the stream&#39;s dedup window is acknowledged as duplicate rather than stored twice. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject to publish to, in the org&#39;s own namespace — e.g. orders.created. No wildcards. | [optional] 

## Methods

### NewBusPublish

`func NewBusPublish() *BusPublish`

NewBusPublish instantiates a new BusPublish object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusPublishWithDefaults

`func NewBusPublishWithDefaults() *BusPublish`

NewBusPublishWithDefaults instantiates a new BusPublish object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BusPublish) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BusPublish) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BusPublish) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BusPublish) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *BusPublish) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BusPublish) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BusPublish) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BusPublish) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubject

`func (o *BusPublish) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BusPublish) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BusPublish) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *BusPublish) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


