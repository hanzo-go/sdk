# TeamTeamChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Archived reports that the channel has been closed. It is the platform&#39;s own Space attribute — the same one the Team client writes — and NOT a field of the work facet, so there is exactly one answer to \&quot;is this channel open\&quot;. | [optional] 
**Bindings** | Pointer to **[]string** | Bindings are what this channel is ABOUT, each a \&quot;&lt;kind&gt;:&lt;ref&gt;\&quot; string — \&quot;project:acme/web\&quot;, \&quot;repo:hanzoai/cloud\&quot;, \&quot;issue:1010\&quot;. One list rather than one field per kind, because the next thing a channel can be about should not be a schema change; and a bound value is opaque here on purpose, since the app that owns a project is the app that can resolve one. HIP-0523 §2: a binding is a REFERENCE, never a copy — a channel holding an issue&#39;s title or status would be the parallel work-item store HIP-1160 §1 forbids. | [optional] 
**Direct** | Pointer to **bool** | Direct reports that this is a channel between people rather than a named channel. It is derived from the document&#39;s class, so it cannot disagree with what the client will render. | [optional] 
**Id** | Pointer to **string** | ID is the channel document&#39;s own id, and the value the bind op addresses. It is unique within a space, not across the org. | [optional] 
**Kind** | Pointer to **string** | Kind is what sort of channel this is: \&quot;public\&quot; (open to everyone in the space), \&quot;private\&quot; (open to its members) or \&quot;dm\&quot; (a conversation between the people in it). It is derived from the class and the private flag, so it cannot disagree with Direct and Private beside it. | [optional] 
**Life** | Pointer to **string** | Life is the channel&#39;s lifecycle INTENT — \&quot;standing\&quot; or \&quot;bound\&quot; (HIP-0523 §2). Absent on the document it reads \&quot;standing\&quot;: a channel nobody classified is one that persists. | [optional] 
**Members** | Pointer to **[]string** | Members are the account uuids in the channel, agents included: an agent projects as a space member under a uuid derived from its id, so a caller comparing this against GET /v1/bot/members learns which channels an agent is in. | [optional] 
**Name** | Pointer to **string** | Name is what a person sees in a sidebar. A direct message carries none, so this is empty for one — the members are its name. | [optional] 
**Private** | Pointer to **bool** | Private reports that the channel is restricted to its members. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding this channel. It is part of the channel&#39;s address: two spaces of one org may each hold a channel with the same name, and only the pair identifies one. | [optional] 
**Topic** | Pointer to **string** | Topic is the channel&#39;s own one-line subject, as the Team client sets it. | [optional] 

## Methods

### NewTeamTeamChannel

`func NewTeamTeamChannel() *TeamTeamChannel`

NewTeamTeamChannel instantiates a new TeamTeamChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelWithDefaults

`func NewTeamTeamChannelWithDefaults() *TeamTeamChannel`

NewTeamTeamChannelWithDefaults instantiates a new TeamTeamChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TeamTeamChannel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamTeamChannel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamTeamChannel) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TeamTeamChannel) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetBindings

`func (o *TeamTeamChannel) GetBindings() []string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *TeamTeamChannel) GetBindingsOk() (*[]string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *TeamTeamChannel) SetBindings(v []string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *TeamTeamChannel) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetDirect

`func (o *TeamTeamChannel) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *TeamTeamChannel) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *TeamTeamChannel) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *TeamTeamChannel) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TeamTeamChannel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TeamTeamChannel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TeamTeamChannel) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TeamTeamChannel) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLife

`func (o *TeamTeamChannel) GetLife() string`

GetLife returns the Life field if non-nil, zero value otherwise.

### GetLifeOk

`func (o *TeamTeamChannel) GetLifeOk() (*string, bool)`

GetLifeOk returns a tuple with the Life field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLife

`func (o *TeamTeamChannel) SetLife(v string)`

SetLife sets Life field to given value.

### HasLife

`func (o *TeamTeamChannel) HasLife() bool`

HasLife returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamChannel) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamChannel) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamChannel) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamChannel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *TeamTeamChannel) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamTeamChannel) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamTeamChannel) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamTeamChannel) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamChannel) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamChannel) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamChannel) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamChannel) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamTeamChannel) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamTeamChannel) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamTeamChannel) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamTeamChannel) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


