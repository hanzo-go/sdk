# AllowlistPutIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroups** | Pointer to [**map[string]map[string][]string**](map.md) | AccessGroups REPLACES the org&#39;s named access groups, as group name -&gt; channel -&gt; entries. Absent or null leaves them alone. | [optional] 
**Channel** | Pointer to **string** | Channel is the transport to edit: discord, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | [optional] 
**Dm** | Pointer to **[]string** | DM REPLACES the config-managed DM allow entries. Absent or null leaves them alone; an empty list clears them. It never touches senders approved through pairing — a policy edit cannot revoke an approved pairing. | [optional] 
**DmPolicy** | Pointer to **string** | DMPolicy sets how direct messages are admitted: \&quot;pairing\&quot; (a person must be approved first), \&quot;allowlist\&quot; (only listed senders) or \&quot;open\&quot;. Empty leaves it unchanged. | [optional] 
**Group** | Pointer to **[]string** | Group REPLACES the config-managed group allow entries. Absent or null leaves them alone; an empty list clears them. | [optional] 
**GroupPolicy** | Pointer to **string** | GroupPolicy sets how group and thread rooms are admitted: \&quot;open\&quot;, \&quot;allowlist\&quot; or \&quot;disabled\&quot;. Empty leaves it unchanged. | [optional] 

## Methods

### NewAllowlistPutIn

`func NewAllowlistPutIn() *AllowlistPutIn`

NewAllowlistPutIn instantiates a new AllowlistPutIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistPutInWithDefaults

`func NewAllowlistPutInWithDefaults() *AllowlistPutIn`

NewAllowlistPutInWithDefaults instantiates a new AllowlistPutIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroups

`func (o *AllowlistPutIn) GetAccessGroups() map[string]map[string][]string`

GetAccessGroups returns the AccessGroups field if non-nil, zero value otherwise.

### GetAccessGroupsOk

`func (o *AllowlistPutIn) GetAccessGroupsOk() (*map[string]map[string][]string, bool)`

GetAccessGroupsOk returns a tuple with the AccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroups

`func (o *AllowlistPutIn) SetAccessGroups(v map[string]map[string][]string)`

SetAccessGroups sets AccessGroups field to given value.

### HasAccessGroups

`func (o *AllowlistPutIn) HasAccessGroups() bool`

HasAccessGroups returns a boolean if a field has been set.

### GetChannel

`func (o *AllowlistPutIn) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AllowlistPutIn) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AllowlistPutIn) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AllowlistPutIn) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDm

`func (o *AllowlistPutIn) GetDm() []string`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *AllowlistPutIn) GetDmOk() (*[]string, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *AllowlistPutIn) SetDm(v []string)`

SetDm sets Dm field to given value.

### HasDm

`func (o *AllowlistPutIn) HasDm() bool`

HasDm returns a boolean if a field has been set.

### GetDmPolicy

`func (o *AllowlistPutIn) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *AllowlistPutIn) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *AllowlistPutIn) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *AllowlistPutIn) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroup

`func (o *AllowlistPutIn) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AllowlistPutIn) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AllowlistPutIn) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AllowlistPutIn) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *AllowlistPutIn) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *AllowlistPutIn) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *AllowlistPutIn) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *AllowlistPutIn) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


