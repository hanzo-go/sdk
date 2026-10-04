# TeamTeamMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **[]string** | Degraded names a roster source that did not answer — \&quot;agents\&quot; — so an empty agent list reads as an outage rather than as an org with no agents. Absent when every source answered. | [optional] 
**Members** | Pointer to [**[]TeamTeamMember**](TeamTeamMember.md) | Members are the space&#39;s people and the org&#39;s agents, people first, each group by name. | [optional] 

## Methods

### NewTeamTeamMembers

`func NewTeamTeamMembers() *TeamTeamMembers`

NewTeamTeamMembers instantiates a new TeamTeamMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamMembersWithDefaults

`func NewTeamTeamMembersWithDefaults() *TeamTeamMembers`

NewTeamTeamMembersWithDefaults instantiates a new TeamTeamMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *TeamTeamMembers) GetDegraded() []string`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *TeamTeamMembers) GetDegradedOk() (*[]string, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *TeamTeamMembers) SetDegraded(v []string)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *TeamTeamMembers) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamMembers) GetMembers() []TeamTeamMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamMembers) GetMembersOk() (*[]TeamTeamMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamMembers) SetMembers(v []TeamTeamMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


