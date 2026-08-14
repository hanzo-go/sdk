# AllowlistView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroups** | Pointer to [**map[string]map[string][]string**](map.md) | AccessGroups is the org&#39;s named sender sets, as group name -&gt; channel -&gt; member entries, held once for the whole org. A DM or Group entry written &#x60;accessGroup:&lt;name&gt;&#x60; admits any sender listed under that name for THIS channel, or under the channel &#x60;*&#x60;, which is how one set covers all four transports. Replaced wholesale by the PUT. | [optional] 
**Dm** | Pointer to **[]string** | DM is the CONFIG-managed DM allow entries — the list PUT /v1/channels/allowlist owns and replaces wholesale. An entry matches a sender either EXACTLY, as the transport-native id inbox messages carry, or as &#x60;accessGroup:&lt;name&gt;&#x60; resolved through AccessGroups. A bare &#x60;*&#x60; admits everyone, but only while DMPolicy is \&quot;open\&quot;: it is gate syntax, not an identity, so under \&quot;allowlist\&quot; it matches nobody. | [optional] 
**DmPolicy** | Pointer to **string** | DMPolicy decides every inbound DIRECT message, defaulting to \&quot;pairing\&quot; when the org has never set one. \&quot;pairing\&quot;: a sender with no entry is sent a pairing code and the message is DROPPED — it never reaches the inbox — and they are admitted only once an admin approves. \&quot;allowlist\&quot;: only DM admits, and Paired senders are suspended, since a pairing grant counts under \&quot;pairing\&quot; alone. \&quot;open\&quot; is not unconditional either — it still requires &#x60;*&#x60; or a matching entry in DM. | [optional] 
**Group** | Pointer to **[]string** | Group is the CONFIG-managed group allow entries, consulted only while GroupPolicy is \&quot;allowlist\&quot;. Entries match the same two ways as DM, and here a bare &#x60;*&#x60; admits every sender in the room. | [optional] 
**GroupPolicy** | Pointer to **string** | GroupPolicy decides every inbound GROUP or THREAD message — a thread is a group surface — defaulting to \&quot;open\&quot;. \&quot;open\&quot; admits every sender in the room. \&quot;allowlist\&quot; admits only what Group lists, so an EMPTY Group blocks the channel&#39;s group rooms outright. \&quot;disabled\&quot; drops all of them. | [optional] 
**Paired** | Pointer to **[]string** | Paired is the senders admitted by PAIRING — the entries POST /v1/channels/pairing/approve minted, DM scope only. READ-ONLY on this endpoint: the PUT writes config entries and can never revoke one of these (listing a paired sender under DM instead promotes that entry to config, which the admin then owns). They admit only while DMPolicy is \&quot;pairing\&quot;. | [optional] 

## Methods

### NewAllowlistView

`func NewAllowlistView() *AllowlistView`

NewAllowlistView instantiates a new AllowlistView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistViewWithDefaults

`func NewAllowlistViewWithDefaults() *AllowlistView`

NewAllowlistViewWithDefaults instantiates a new AllowlistView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroups

`func (o *AllowlistView) GetAccessGroups() map[string]map[string][]string`

GetAccessGroups returns the AccessGroups field if non-nil, zero value otherwise.

### GetAccessGroupsOk

`func (o *AllowlistView) GetAccessGroupsOk() (*map[string]map[string][]string, bool)`

GetAccessGroupsOk returns a tuple with the AccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroups

`func (o *AllowlistView) SetAccessGroups(v map[string]map[string][]string)`

SetAccessGroups sets AccessGroups field to given value.

### HasAccessGroups

`func (o *AllowlistView) HasAccessGroups() bool`

HasAccessGroups returns a boolean if a field has been set.

### GetDm

`func (o *AllowlistView) GetDm() []string`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *AllowlistView) GetDmOk() (*[]string, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *AllowlistView) SetDm(v []string)`

SetDm sets Dm field to given value.

### HasDm

`func (o *AllowlistView) HasDm() bool`

HasDm returns a boolean if a field has been set.

### GetDmPolicy

`func (o *AllowlistView) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *AllowlistView) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *AllowlistView) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *AllowlistView) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroup

`func (o *AllowlistView) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AllowlistView) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AllowlistView) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AllowlistView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *AllowlistView) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *AllowlistView) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *AllowlistView) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *AllowlistView) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.

### GetPaired

`func (o *AllowlistView) GetPaired() []string`

GetPaired returns the Paired field if non-nil, zero value otherwise.

### GetPairedOk

`func (o *AllowlistView) GetPairedOk() (*[]string, bool)`

GetPairedOk returns a tuple with the Paired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaired

`func (o *AllowlistView) SetPaired(v []string)`

SetPaired sets Paired field to given value.

### HasPaired

`func (o *AllowlistView) HasPaired() bool`

HasPaired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


