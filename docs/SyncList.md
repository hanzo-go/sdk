# SyncList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SyncView**](SyncView.md) | Data is the org&#39;s syncs, each with its endpoints, policy and last-synced time. | [optional] 

## Methods

### NewSyncList

`func NewSyncList() *SyncList`

NewSyncList instantiates a new SyncList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncListWithDefaults

`func NewSyncListWithDefaults() *SyncList`

NewSyncListWithDefaults instantiates a new SyncList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SyncList) GetData() []SyncView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncList) GetDataOk() (*[]SyncView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncList) SetData(v []SyncView)`

SetData sets Data field to given value.

### HasData

`func (o *SyncList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


