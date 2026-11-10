# TeamTeamChannelBind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **[]string** | Bindings REPLACES what the channel is about, wholly. It is a replace and not a merge because a caller that cannot remove a binding would have no way to correct a wrong one, and an empty list sent explicitly is how a channel is unbound. Absent (null) leaves the existing list alone. | [optional] 
**Id** | Pointer to **string** | ID is the channel to bind, from the path. The URL is the authority; a body carrying another id cannot redirect the write. | [optional] 
**Life** | Pointer to **string** | Life sets the lifecycle intent: \&quot;standing\&quot; or \&quot;bound\&quot;. Any other value is refused rather than stored, so a reader never has to interpret a third one. Empty leaves the current intent unchanged. | [optional] 
**Space** | Pointer to **string** | Space names the space holding the channel. It is required, because a channel id is unique only within one and searching every space for a matching id would make the write&#39;s target depend on iteration order. | [optional] 

## Methods

### NewTeamTeamChannelBind

`func NewTeamTeamChannelBind() *TeamTeamChannelBind`

NewTeamTeamChannelBind instantiates a new TeamTeamChannelBind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelBindWithDefaults

`func NewTeamTeamChannelBindWithDefaults() *TeamTeamChannelBind`

NewTeamTeamChannelBindWithDefaults instantiates a new TeamTeamChannelBind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *TeamTeamChannelBind) GetBindings() []string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *TeamTeamChannelBind) GetBindingsOk() (*[]string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *TeamTeamChannelBind) SetBindings(v []string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *TeamTeamChannelBind) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetId

`func (o *TeamTeamChannelBind) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamChannelBind) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamChannelBind) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamChannelBind) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLife

`func (o *TeamTeamChannelBind) GetLife() string`

GetLife returns the Life field if non-nil, zero value otherwise.

### GetLifeOk

`func (o *TeamTeamChannelBind) GetLifeOk() (*string, bool)`

GetLifeOk returns a tuple with the Life field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLife

`func (o *TeamTeamChannelBind) SetLife(v string)`

SetLife sets Life field to given value.

### HasLife

`func (o *TeamTeamChannelBind) HasLife() bool`

HasLife returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamChannelBind) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamChannelBind) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamChannelBind) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamChannelBind) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


