# CloudAnnQueueDetailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedCount** | Pointer to **int32** | CompletedCount is how many have been reviewed. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when it was created, RFC3339 in UTC. | [optional] 
**Description** | Pointer to **string** | Description is its free text, omitted when empty. | [optional] 
**Id** | Pointer to **string** | ID is the queue&#39;s id. | [optional] 
**Items** | Pointer to [**[]CloudAnnItemView**](CloudAnnItemView.md) | Items is the queue&#39;s first page of items (up to 100). | [optional] 
**Name** | Pointer to **string** | Name is its display handle. | [optional] 
**PendingCount** | Pointer to **int32** | PendingCount is how many of its items are still awaiting review. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs are the eval score-configs reviewers grade against. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed, RFC3339 in UTC. | [optional] 

## Methods

### NewCloudAnnQueueDetailView

`func NewCloudAnnQueueDetailView() *CloudAnnQueueDetailView`

NewCloudAnnQueueDetailView instantiates a new CloudAnnQueueDetailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnnQueueDetailViewWithDefaults

`func NewCloudAnnQueueDetailViewWithDefaults() *CloudAnnQueueDetailView`

NewCloudAnnQueueDetailViewWithDefaults instantiates a new CloudAnnQueueDetailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedCount

`func (o *CloudAnnQueueDetailView) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *CloudAnnQueueDetailView) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *CloudAnnQueueDetailView) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *CloudAnnQueueDetailView) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAnnQueueDetailView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAnnQueueDetailView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAnnQueueDetailView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAnnQueueDetailView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAnnQueueDetailView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAnnQueueDetailView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAnnQueueDetailView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAnnQueueDetailView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CloudAnnQueueDetailView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAnnQueueDetailView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAnnQueueDetailView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAnnQueueDetailView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *CloudAnnQueueDetailView) GetItems() []CloudAnnItemView`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudAnnQueueDetailView) GetItemsOk() (*[]CloudAnnItemView, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudAnnQueueDetailView) SetItems(v []CloudAnnItemView)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudAnnQueueDetailView) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *CloudAnnQueueDetailView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAnnQueueDetailView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAnnQueueDetailView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAnnQueueDetailView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingCount

`func (o *CloudAnnQueueDetailView) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *CloudAnnQueueDetailView) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *CloudAnnQueueDetailView) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *CloudAnnQueueDetailView) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *CloudAnnQueueDetailView) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *CloudAnnQueueDetailView) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *CloudAnnQueueDetailView) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *CloudAnnQueueDetailView) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAnnQueueDetailView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAnnQueueDetailView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAnnQueueDetailView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAnnQueueDetailView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


