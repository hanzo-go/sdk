# Listed

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

### NewListed

`func NewListed() *Listed`

NewListed instantiates a new Listed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedWithDefaults

`func NewListedWithDefaults() *Listed`

NewListedWithDefaults instantiates a new Listed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *Listed) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Listed) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Listed) SetMembers(v int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Listed) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *Listed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Listed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Listed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Listed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *Listed) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Listed) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Listed) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Listed) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRoom

`func (o *Listed) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *Listed) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *Listed) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *Listed) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetSpace

`func (o *Listed) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *Listed) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *Listed) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *Listed) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTopic

`func (o *Listed) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Listed) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Listed) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Listed) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUpdated

`func (o *Listed) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Listed) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Listed) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Listed) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


