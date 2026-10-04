# TeamTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is false for an agent the org has retired: it stays in the roster so its past messages keep an author, and it no longer answers. | [optional] 
**Avatar** | Pointer to **string** | Avatar is the blob id of an uploaded picture, readable at GET /v1/team/files/{space}/avatar?file&#x3D;{avatar}. Empty when the person has none and the client draws initials. | [optional] 
**Bot** | Pointer to **bool** | Bot reports that this member is one of the org&#39;s agents rather than a person. An agent answers when it is @-mentioned or messaged directly. | [optional] 
**Id** | Pointer to **string** | ID is the account uuid — the value every message&#39;s author, every room&#39;s members list and every &#x60;&lt;@id&gt;&#x60; mention token names. | [optional] 
**Name** | Pointer to **string** | Name is what to call them: the name IAM holds for a person, the agent&#39;s own name for an agent. Falls back to the id when neither has one. | [optional] 
**PersonRef** | Pointer to **string** | PersonRef is the member&#39;s Person document id, the reference a mention stores. | [optional] 
**Presence** | Pointer to **string** | Presence is \&quot;online\&quot; while the member holds a live connection to this space (a Team client or an event stream), \&quot;offline\&quot; otherwise. | [optional] 
**Role** | Pointer to **string** | Role is the member&#39;s role in the space — owner, admin, member or guest. Agents read \&quot;member\&quot;. | [optional] 

## Methods

### NewTeamTeamMember

`func NewTeamTeamMember() *TeamTeamMember`

NewTeamTeamMember instantiates a new TeamTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamMemberWithDefaults

`func NewTeamTeamMemberWithDefaults() *TeamTeamMember`

NewTeamTeamMemberWithDefaults instantiates a new TeamTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TeamTeamMember) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TeamTeamMember) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TeamTeamMember) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TeamTeamMember) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvatar

`func (o *TeamTeamMember) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *TeamTeamMember) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *TeamTeamMember) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *TeamTeamMember) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBot

`func (o *TeamTeamMember) GetBot() bool`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *TeamTeamMember) GetBotOk() (*bool, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *TeamTeamMember) SetBot(v bool)`

SetBot sets Bot field to given value.

### HasBot

`func (o *TeamTeamMember) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonRef

`func (o *TeamTeamMember) GetPersonRef() string`

GetPersonRef returns the PersonRef field if non-nil, zero value otherwise.

### GetPersonRefOk

`func (o *TeamTeamMember) GetPersonRefOk() (*string, bool)`

GetPersonRefOk returns a tuple with the PersonRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonRef

`func (o *TeamTeamMember) SetPersonRef(v string)`

SetPersonRef sets PersonRef field to given value.

### HasPersonRef

`func (o *TeamTeamMember) HasPersonRef() bool`

HasPersonRef returns a boolean if a field has been set.

### GetPresence

`func (o *TeamTeamMember) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *TeamTeamMember) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *TeamTeamMember) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *TeamTeamMember) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRole

`func (o *TeamTeamMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamTeamMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamTeamMember) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TeamTeamMember) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


