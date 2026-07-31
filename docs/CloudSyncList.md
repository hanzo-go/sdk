# CloudSyncList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudSyncView**](CloudSyncView.md) | Data is the org&#39;s syncs, each with its endpoints, policy and last-synced time. | [optional] 

## Methods

### NewCloudSyncList

`func NewCloudSyncList() *CloudSyncList`

NewCloudSyncList instantiates a new CloudSyncList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSyncListWithDefaults

`func NewCloudSyncListWithDefaults() *CloudSyncList`

NewCloudSyncListWithDefaults instantiates a new CloudSyncList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSyncList) GetData() []CloudSyncView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSyncList) GetDataOk() (*[]CloudSyncView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSyncList) SetData(v []CloudSyncView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSyncList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


