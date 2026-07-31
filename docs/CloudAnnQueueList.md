# CloudAnnQueueList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudAnnQueueView**](CloudAnnQueueView.md) | Data is the page of queues. | [optional] 
**Meta** | Pointer to [**CloudListMeta**](CloudListMeta.md) | Meta is the paging that produced it. | [optional] 

## Methods

### NewCloudAnnQueueList

`func NewCloudAnnQueueList() *CloudAnnQueueList`

NewCloudAnnQueueList instantiates a new CloudAnnQueueList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnnQueueListWithDefaults

`func NewCloudAnnQueueListWithDefaults() *CloudAnnQueueList`

NewCloudAnnQueueListWithDefaults instantiates a new CloudAnnQueueList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAnnQueueList) GetData() []CloudAnnQueueView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAnnQueueList) GetDataOk() (*[]CloudAnnQueueView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAnnQueueList) SetData(v []CloudAnnQueueView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAnnQueueList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *CloudAnnQueueList) GetMeta() CloudListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CloudAnnQueueList) GetMetaOk() (*CloudListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CloudAnnQueueList) SetMeta(v CloudListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CloudAnnQueueList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


