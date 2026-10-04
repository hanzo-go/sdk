# GitSubscriptionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the Slack channel id or name the notifier posts to. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Events** | Pointer to **[]string** | Events is the kind filter; absent means every deliverable kind. | [optional] 
**Id** | Pointer to **string** | ID is the subscription&#39;s identifier (\&quot;sub_…\&quot;), the handle to delete it by. | [optional] 
**Repo** | Pointer to **string** | Repo is the repo whose lifecycle events are delivered. | [optional] 

## Methods

### NewGitSubscriptionView

`func NewGitSubscriptionView() *GitSubscriptionView`

NewGitSubscriptionView instantiates a new GitSubscriptionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitSubscriptionViewWithDefaults

`func NewGitSubscriptionViewWithDefaults() *GitSubscriptionView`

NewGitSubscriptionViewWithDefaults instantiates a new GitSubscriptionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *GitSubscriptionView) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GitSubscriptionView) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GitSubscriptionView) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GitSubscriptionView) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GitSubscriptionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitSubscriptionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitSubscriptionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitSubscriptionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *GitSubscriptionView) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GitSubscriptionView) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GitSubscriptionView) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GitSubscriptionView) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *GitSubscriptionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitSubscriptionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitSubscriptionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GitSubscriptionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepo

`func (o *GitSubscriptionView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitSubscriptionView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitSubscriptionView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitSubscriptionView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


