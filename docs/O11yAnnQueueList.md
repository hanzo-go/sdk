# O11yAnnQueueList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yAnnQueueView**](O11yAnnQueueView.md) | Data is the page of queues. | [optional] 
**Meta** | Pointer to [**O11yListMeta**](O11yListMeta.md) | Meta is the paging that produced it. | [optional] 

## Methods

### NewO11yAnnQueueList

`func NewO11yAnnQueueList() *O11yAnnQueueList`

NewO11yAnnQueueList instantiates a new O11yAnnQueueList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAnnQueueListWithDefaults

`func NewO11yAnnQueueListWithDefaults() *O11yAnnQueueList`

NewO11yAnnQueueListWithDefaults instantiates a new O11yAnnQueueList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yAnnQueueList) GetData() []O11yAnnQueueView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yAnnQueueList) GetDataOk() (*[]O11yAnnQueueView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yAnnQueueList) SetData(v []O11yAnnQueueView)`

SetData sets Data field to given value.

### HasData

`func (o *O11yAnnQueueList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *O11yAnnQueueList) GetMeta() O11yListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yAnnQueueList) GetMetaOk() (*O11yListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yAnnQueueList) SetMeta(v O11yListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yAnnQueueList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


