# O11yAnnQueueDetailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedCount** | Pointer to **int32** | CompletedCount is how many have been reviewed. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when it was created, RFC3339 in UTC. | [optional] 
**Description** | Pointer to **string** | Description is its free text, omitted when empty. | [optional] 
**Id** | Pointer to **string** | ID is the queue&#39;s id. | [optional] 
**Items** | Pointer to [**[]O11yAnnItemView**](O11yAnnItemView.md) | Items is the queue&#39;s first page of items (up to 100). | [optional] 
**Name** | Pointer to **string** | Name is its display handle. | [optional] 
**PendingCount** | Pointer to **int32** | PendingCount is how many of its items are still awaiting review. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs are the eval score-configs reviewers grade against. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed, RFC3339 in UTC. | [optional] 

## Methods

### NewO11yAnnQueueDetailView

`func NewO11yAnnQueueDetailView() *O11yAnnQueueDetailView`

NewO11yAnnQueueDetailView instantiates a new O11yAnnQueueDetailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAnnQueueDetailViewWithDefaults

`func NewO11yAnnQueueDetailViewWithDefaults() *O11yAnnQueueDetailView`

NewO11yAnnQueueDetailViewWithDefaults instantiates a new O11yAnnQueueDetailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedCount

`func (o *O11yAnnQueueDetailView) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *O11yAnnQueueDetailView) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *O11yAnnQueueDetailView) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *O11yAnnQueueDetailView) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yAnnQueueDetailView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yAnnQueueDetailView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yAnnQueueDetailView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yAnnQueueDetailView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *O11yAnnQueueDetailView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yAnnQueueDetailView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yAnnQueueDetailView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yAnnQueueDetailView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *O11yAnnQueueDetailView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAnnQueueDetailView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAnnQueueDetailView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yAnnQueueDetailView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *O11yAnnQueueDetailView) GetItems() []O11yAnnItemView`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yAnnQueueDetailView) GetItemsOk() (*[]O11yAnnItemView, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yAnnQueueDetailView) SetItems(v []O11yAnnItemView)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yAnnQueueDetailView) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *O11yAnnQueueDetailView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yAnnQueueDetailView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yAnnQueueDetailView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yAnnQueueDetailView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingCount

`func (o *O11yAnnQueueDetailView) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *O11yAnnQueueDetailView) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *O11yAnnQueueDetailView) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *O11yAnnQueueDetailView) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *O11yAnnQueueDetailView) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *O11yAnnQueueDetailView) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *O11yAnnQueueDetailView) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *O11yAnnQueueDetailView) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yAnnQueueDetailView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yAnnQueueDetailView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yAnnQueueDetailView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yAnnQueueDetailView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


