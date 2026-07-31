# CloudAnnQueueView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when it was created, RFC3339 in UTC. | [optional] 
**Description** | Pointer to **string** | Description is its free text, omitted when empty. | [optional] 
**Id** | Pointer to **string** | ID is the queue&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is its display handle. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs are the eval score-configs reviewers grade against. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed, RFC3339 in UTC. | [optional] 

## Methods

### NewCloudAnnQueueView

`func NewCloudAnnQueueView() *CloudAnnQueueView`

NewCloudAnnQueueView instantiates a new CloudAnnQueueView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnnQueueViewWithDefaults

`func NewCloudAnnQueueViewWithDefaults() *CloudAnnQueueView`

NewCloudAnnQueueViewWithDefaults instantiates a new CloudAnnQueueView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudAnnQueueView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAnnQueueView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAnnQueueView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAnnQueueView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAnnQueueView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAnnQueueView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAnnQueueView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAnnQueueView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CloudAnnQueueView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAnnQueueView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAnnQueueView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAnnQueueView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAnnQueueView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAnnQueueView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAnnQueueView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAnnQueueView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *CloudAnnQueueView) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *CloudAnnQueueView) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *CloudAnnQueueView) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *CloudAnnQueueView) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAnnQueueView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAnnQueueView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAnnQueueView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAnnQueueView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


