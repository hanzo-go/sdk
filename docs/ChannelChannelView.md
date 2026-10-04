# ChannelChannelView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the id-shaped fact about that connection: the lowercased external id integrations custodies for it — a Discord guild id, a Slack team (workspace) id, a Teams AAD tenant id, or the Telegram chat the org bound. Empty when not connected. Informational: the access policy keys on (org, channel), so exactly one account is representable per pair. | [optional] 
**AccountLabel** | Pointer to **string** | AccountLabel is the human label of that same account — the Discord guild name, the Slack team name, the Teams tenant name (falling back to the tenant id), the Telegram chat title. DISPLAY ONLY: never a key, and never swapped with Account, on any surface. | [optional] 
**Capabilities** | Pointer to [**ChannelCapabilities**](ChannelCapabilities.md) | Capabilities is what this transport renders natively — read it before composing a message that needs threading, media or interactive actions. | [optional] 
**Connected** | Pointer to **bool** | Connected is whether integrations holds a connection for (this org, this transport) — whether someone finished its connect flow. False leaves Account and AccountLabel empty, and a send is then refused downstream rather than here: by the transport&#39;s own binding check (403 for a Telegram chat this org has not bound, 409 for a Discord or Teams room with no inbound-learned route), or on Slack by the absent per-org bot token, which surfaces as 502. | [optional] 
**DmPolicy** | Pointer to **string** | DMPolicy is how this org admits direct messages here: \&quot;pairing\&quot;, \&quot;allowlist\&quot; or \&quot;open\&quot;, defaulting to \&quot;pairing\&quot; when the org has never set one. | [optional] 
**GroupPolicy** | Pointer to **string** | GroupPolicy is how this org admits group and thread rooms here: \&quot;open\&quot;, \&quot;allowlist\&quot; or \&quot;disabled\&quot;, defaulting to \&quot;open\&quot;. Both policy fields come back EMPTY — rather than the listing failing — when the policy cannot be read; GET /v1/channel/allowlist carries the same two with the entries they consult. | [optional] 
**Id** | Pointer to **string** | ID is the fixed transport identifier — discord, slack, teams, telegram or whatsapp — and the value every route on this surface names a channel by, including the &#x60;:channel&#x60; segment of the send path. The listing is always in that order. | [optional] 
**PendingPairing** | Pointer to **int64** | PendingPairing counts the org&#39;s UNEXPIRED pairing requests on this channel: exactly the rows GET /v1/channel/pairing returns for it, one per person waiting on an admin. It never exceeds three — the pending cap per (org, channel) — and expired requests are not counted. | [optional] 

## Methods

### NewChannelChannelView

`func NewChannelChannelView() *ChannelChannelView`

NewChannelChannelView instantiates a new ChannelChannelView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelChannelViewWithDefaults

`func NewChannelChannelViewWithDefaults() *ChannelChannelView`

NewChannelChannelViewWithDefaults instantiates a new ChannelChannelView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ChannelChannelView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ChannelChannelView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ChannelChannelView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ChannelChannelView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountLabel

`func (o *ChannelChannelView) GetAccountLabel() string`

GetAccountLabel returns the AccountLabel field if non-nil, zero value otherwise.

### GetAccountLabelOk

`func (o *ChannelChannelView) GetAccountLabelOk() (*string, bool)`

GetAccountLabelOk returns a tuple with the AccountLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLabel

`func (o *ChannelChannelView) SetAccountLabel(v string)`

SetAccountLabel sets AccountLabel field to given value.

### HasAccountLabel

`func (o *ChannelChannelView) HasAccountLabel() bool`

HasAccountLabel returns a boolean if a field has been set.

### GetCapabilities

`func (o *ChannelChannelView) GetCapabilities() ChannelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ChannelChannelView) GetCapabilitiesOk() (*ChannelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ChannelChannelView) SetCapabilities(v ChannelCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ChannelChannelView) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetConnected

`func (o *ChannelChannelView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ChannelChannelView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ChannelChannelView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ChannelChannelView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDmPolicy

`func (o *ChannelChannelView) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *ChannelChannelView) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *ChannelChannelView) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *ChannelChannelView) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *ChannelChannelView) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *ChannelChannelView) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *ChannelChannelView) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *ChannelChannelView) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.

### GetId

`func (o *ChannelChannelView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelChannelView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelChannelView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelChannelView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPendingPairing

`func (o *ChannelChannelView) GetPendingPairing() int64`

GetPendingPairing returns the PendingPairing field if non-nil, zero value otherwise.

### GetPendingPairingOk

`func (o *ChannelChannelView) GetPendingPairingOk() (*int64, bool)`

GetPendingPairingOk returns a tuple with the PendingPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPairing

`func (o *ChannelChannelView) SetPendingPairing(v int64)`

SetPendingPairing sets PendingPairing field to given value.

### HasPendingPairing

`func (o *ChannelChannelView) HasPendingPairing() bool`

HasPendingPairing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


