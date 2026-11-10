# TeamTeamChannelNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **[]string** | Bindings are what the channel is about, each \&quot;&lt;kind&gt;:&lt;ref&gt;\&quot;. | [optional] 
**Life** | Pointer to **string** | Life is the lifecycle intent, \&quot;standing\&quot; or \&quot;bound\&quot;; empty reads standing. | [optional] 
**Members** | Pointer to **[]string** | Members are the account uuids in the channel. A public channel may open empty — anyone in the org can find it — and a private one that names nobody is refused rather than created unreachable. | [optional] 
**Name** | Pointer to **string** | Name is what a person sees in a sidebar — \&quot;bugfix-1010\&quot;, not \&quot;#bugfix-1010\&quot;. The sigil is how a client DRAWS a channel, and storing it would put it in the name twice the first time a client added its own. | [optional] 
**Private** | Pointer to **bool** | Private restricts the channel to its members. Public is the default because a channel nobody can find is the more surprising of the two. | [optional] 
**Space** | Pointer to **string** | Space is where the channel is opened. Optional: an org with one space has no choice to make, so it does not have to state one. An org with several must, because picking for it would make the channel&#39;s home depend on iteration order. | [optional] 
**Topic** | Pointer to **string** | Topic is the channel&#39;s one-line subject. | [optional] 

## Methods

### NewTeamTeamChannelNew

`func NewTeamTeamChannelNew() *TeamTeamChannelNew`

NewTeamTeamChannelNew instantiates a new TeamTeamChannelNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelNewWithDefaults

`func NewTeamTeamChannelNewWithDefaults() *TeamTeamChannelNew`

NewTeamTeamChannelNewWithDefaults instantiates a new TeamTeamChannelNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *TeamTeamChannelNew) GetBindings() []string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *TeamTeamChannelNew) GetBindingsOk() (*[]string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *TeamTeamChannelNew) SetBindings(v []string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *TeamTeamChannelNew) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetLife

`func (o *TeamTeamChannelNew) GetLife() string`

GetLife returns the Life field if non-nil, zero value otherwise.

### GetLifeOk

`func (o *TeamTeamChannelNew) GetLifeOk() (*string, bool)`

GetLifeOk returns a tuple with the Life field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLife

`func (o *TeamTeamChannelNew) SetLife(v string)`

SetLife sets Life field to given value.

### HasLife

`func (o *TeamTeamChannelNew) HasLife() bool`

HasLife returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamChannelNew) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamChannelNew) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamChannelNew) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamChannelNew) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamChannelNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamChannelNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamChannelNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamChannelNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *TeamTeamChannelNew) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamTeamChannelNew) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamTeamChannelNew) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamTeamChannelNew) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamChannelNew) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamChannelNew) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamChannelNew) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamChannelNew) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamTeamChannelNew) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamTeamChannelNew) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamTeamChannelNew) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamTeamChannelNew) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


