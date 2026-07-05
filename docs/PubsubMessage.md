# PubsubMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Message subject (e.g. orders.created) | [optional] 
**Data** | Pointer to **string** | Base64-encoded message payload | [optional] 
**Headers** | Pointer to **map[string][]string** | Optional message headers | [optional] 

## Methods

### NewPubsubMessage

`func NewPubsubMessage() *PubsubMessage`

NewPubsubMessage instantiates a new PubsubMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubMessageWithDefaults

`func NewPubsubMessageWithDefaults() *PubsubMessage`

NewPubsubMessageWithDefaults instantiates a new PubsubMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PubsubMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PubsubMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PubsubMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PubsubMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetData

`func (o *PubsubMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PubsubMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PubsubMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PubsubMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *PubsubMessage) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PubsubMessage) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PubsubMessage) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PubsubMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


