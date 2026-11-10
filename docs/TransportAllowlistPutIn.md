# TransportAllowlistPutIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroups** | Pointer to [**map[string]map[string][]string**](map.md) | AccessGroups REPLACES the org&#39;s named access groups, as group name -&gt; transport -&gt; entries. Absent or null leaves them alone. | [optional] 
**Dm** | Pointer to **[]string** | DM REPLACES the config-managed DM allow entries. Absent or null leaves them alone; an empty list clears them. It never touches senders approved through pairing — a policy edit cannot revoke an approved pairing. | [optional] 
**DmPolicy** | Pointer to **string** | DMPolicy sets how direct messages are admitted: \&quot;pairing\&quot; (a person must be approved first), \&quot;allowlist\&quot; (only listed senders) or \&quot;open\&quot;. Empty leaves it unchanged. | [optional] 
**Group** | Pointer to **[]string** | Group REPLACES the config-managed group allow entries. Absent or null leaves them alone; an empty list clears them. | [optional] 
**GroupPolicy** | Pointer to **string** | GroupPolicy sets how group and thread channels are admitted: \&quot;open\&quot;, \&quot;allowlist\&quot; or \&quot;disabled\&quot;. Empty leaves it unchanged. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport to edit: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | [optional] 

## Methods

### NewTransportAllowlistPutIn

`func NewTransportAllowlistPutIn() *TransportAllowlistPutIn`

NewTransportAllowlistPutIn instantiates a new TransportAllowlistPutIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportAllowlistPutInWithDefaults

`func NewTransportAllowlistPutInWithDefaults() *TransportAllowlistPutIn`

NewTransportAllowlistPutInWithDefaults instantiates a new TransportAllowlistPutIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroups

`func (o *TransportAllowlistPutIn) GetAccessGroups() map[string]map[string][]string`

GetAccessGroups returns the AccessGroups field if non-nil, zero value otherwise.

### GetAccessGroupsOk

`func (o *TransportAllowlistPutIn) GetAccessGroupsOk() (*map[string]map[string][]string, bool)`

GetAccessGroupsOk returns a tuple with the AccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroups

`func (o *TransportAllowlistPutIn) SetAccessGroups(v map[string]map[string][]string)`

SetAccessGroups sets AccessGroups field to given value.

### HasAccessGroups

`func (o *TransportAllowlistPutIn) HasAccessGroups() bool`

HasAccessGroups returns a boolean if a field has been set.

### GetDm

`func (o *TransportAllowlistPutIn) GetDm() []string`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *TransportAllowlistPutIn) GetDmOk() (*[]string, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *TransportAllowlistPutIn) SetDm(v []string)`

SetDm sets Dm field to given value.

### HasDm

`func (o *TransportAllowlistPutIn) HasDm() bool`

HasDm returns a boolean if a field has been set.

### GetDmPolicy

`func (o *TransportAllowlistPutIn) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *TransportAllowlistPutIn) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *TransportAllowlistPutIn) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *TransportAllowlistPutIn) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroup

`func (o *TransportAllowlistPutIn) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TransportAllowlistPutIn) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TransportAllowlistPutIn) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TransportAllowlistPutIn) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *TransportAllowlistPutIn) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *TransportAllowlistPutIn) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *TransportAllowlistPutIn) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *TransportAllowlistPutIn) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.

### GetTransport

`func (o *TransportAllowlistPutIn) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportAllowlistPutIn) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportAllowlistPutIn) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportAllowlistPutIn) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


