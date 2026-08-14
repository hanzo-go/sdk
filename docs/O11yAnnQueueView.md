# O11yAnnQueueView

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

### NewO11yAnnQueueView

`func NewO11yAnnQueueView() *O11yAnnQueueView`

NewO11yAnnQueueView instantiates a new O11yAnnQueueView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAnnQueueViewWithDefaults

`func NewO11yAnnQueueViewWithDefaults() *O11yAnnQueueView`

NewO11yAnnQueueViewWithDefaults instantiates a new O11yAnnQueueView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yAnnQueueView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yAnnQueueView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yAnnQueueView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yAnnQueueView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *O11yAnnQueueView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yAnnQueueView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yAnnQueueView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yAnnQueueView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *O11yAnnQueueView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAnnQueueView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAnnQueueView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yAnnQueueView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yAnnQueueView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yAnnQueueView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yAnnQueueView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yAnnQueueView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *O11yAnnQueueView) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *O11yAnnQueueView) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *O11yAnnQueueView) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *O11yAnnQueueView) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yAnnQueueView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yAnnQueueView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yAnnQueueView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yAnnQueueView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


