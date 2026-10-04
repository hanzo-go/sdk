# TeamTeamRoomNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **[]string** | Bindings are what the room is about, each \&quot;&lt;kind&gt;:&lt;ref&gt;\&quot;. | [optional] 
**Life** | Pointer to **string** | Life is the lifecycle intent, \&quot;standing\&quot; or \&quot;bound\&quot;; empty reads standing. | [optional] 
**Members** | Pointer to **[]string** | Members are the account uuids in the room. A public room may open empty — anyone in the org can find it — and a private one that names nobody is refused rather than created unreachable. | [optional] 
**Name** | Pointer to **string** | Name is what a person sees in a sidebar — \&quot;bugfix-1010\&quot;, not \&quot;#bugfix-1010\&quot;. The sigil is how a client DRAWS a room, and storing it would put it in the name twice the first time a client added its own. | [optional] 
**Private** | Pointer to **bool** | Private restricts the room to its members. Public is the default because a room nobody can find is the more surprising of the two. | [optional] 
**Space** | Pointer to **string** | Space is where the room is opened. Optional: an org with one space has no choice to make, so it does not have to state one. An org with several must, because picking for it would make the room&#39;s home depend on iteration order. | [optional] 
**Topic** | Pointer to **string** | Topic is the room&#39;s one-line subject. | [optional] 

## Methods

### NewTeamTeamRoomNew

`func NewTeamTeamRoomNew() *TeamTeamRoomNew`

NewTeamTeamRoomNew instantiates a new TeamTeamRoomNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamRoomNewWithDefaults

`func NewTeamTeamRoomNewWithDefaults() *TeamTeamRoomNew`

NewTeamTeamRoomNewWithDefaults instantiates a new TeamTeamRoomNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *TeamTeamRoomNew) GetBindings() []string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *TeamTeamRoomNew) GetBindingsOk() (*[]string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *TeamTeamRoomNew) SetBindings(v []string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *TeamTeamRoomNew) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetLife

`func (o *TeamTeamRoomNew) GetLife() string`

GetLife returns the Life field if non-nil, zero value otherwise.

### GetLifeOk

`func (o *TeamTeamRoomNew) GetLifeOk() (*string, bool)`

GetLifeOk returns a tuple with the Life field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLife

`func (o *TeamTeamRoomNew) SetLife(v string)`

SetLife sets Life field to given value.

### HasLife

`func (o *TeamTeamRoomNew) HasLife() bool`

HasLife returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamRoomNew) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamRoomNew) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamRoomNew) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamRoomNew) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamRoomNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamRoomNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamRoomNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamRoomNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *TeamTeamRoomNew) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamTeamRoomNew) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamTeamRoomNew) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TeamTeamRoomNew) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamRoomNew) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamRoomNew) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamRoomNew) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamRoomNew) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamTeamRoomNew) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamTeamRoomNew) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamTeamRoomNew) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamTeamRoomNew) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


