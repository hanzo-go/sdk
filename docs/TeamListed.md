# TeamListed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **int64** | Members counts the room, and never names anybody in it. | [optional] 
**Name** | Pointer to **string** | Name is what a person sees, without the sigil a client draws. | [optional] 
**Org** | Pointer to **string** | Org owns the room. It is also what a caller filters by to browse one org. | [optional] 
**Room** | Pointer to **string** | Room addresses it in the owning store — what a join is called with. | [optional] 
**Space** | Pointer to **string** | Space is where the room lives inside that org. | [optional] 
**Topic** | Pointer to **string** | Topic is the room&#39;s one-line subject, empty when it has none. | [optional] 
**Updated** | Pointer to **int64** | Updated is when this row was last written, unix seconds. | [optional] 

## Methods

### NewTeamListed

`func NewTeamListed() *TeamListed`

NewTeamListed instantiates a new TeamListed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamListedWithDefaults

`func NewTeamListedWithDefaults() *TeamListed`

NewTeamListedWithDefaults instantiates a new TeamListed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *TeamListed) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamListed) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamListed) SetMembers(v int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamListed) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *TeamListed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamListed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamListed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamListed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *TeamListed) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TeamListed) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TeamListed) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TeamListed) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRoom

`func (o *TeamListed) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TeamListed) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TeamListed) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TeamListed) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetSpace

`func (o *TeamListed) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamListed) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamListed) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamListed) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *TeamListed) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TeamListed) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TeamListed) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TeamListed) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUpdated

`func (o *TeamListed) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TeamListed) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TeamListed) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TeamListed) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


