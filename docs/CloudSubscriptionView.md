# CloudSubscriptionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the Slack channel id or name the notifier posts to. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Events** | Pointer to **[]string** | Events is the kind filter; absent means every deliverable kind. | [optional] 
**Id** | Pointer to **string** | ID is the subscription&#39;s identifier (\&quot;sub_…\&quot;), the handle to delete it by. | [optional] 
**Repo** | Pointer to **string** | Repo is the repo whose lifecycle events are delivered. | [optional] 

## Methods

### NewCloudSubscriptionView

`func NewCloudSubscriptionView() *CloudSubscriptionView`

NewCloudSubscriptionView instantiates a new CloudSubscriptionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubscriptionViewWithDefaults

`func NewCloudSubscriptionViewWithDefaults() *CloudSubscriptionView`

NewCloudSubscriptionViewWithDefaults instantiates a new CloudSubscriptionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CloudSubscriptionView) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudSubscriptionView) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudSubscriptionView) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudSubscriptionView) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSubscriptionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSubscriptionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSubscriptionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSubscriptionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *CloudSubscriptionView) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudSubscriptionView) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudSubscriptionView) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudSubscriptionView) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *CloudSubscriptionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSubscriptionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSubscriptionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSubscriptionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepo

`func (o *CloudSubscriptionView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudSubscriptionView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudSubscriptionView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudSubscriptionView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


