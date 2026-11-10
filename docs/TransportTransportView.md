# TransportTransportView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the id-shaped fact about that connection: the lowercased external id integrations custodies for it — a Discord guild id, a Slack team (workspace) id, a Teams AAD tenant id, or the Telegram chat the org bound. Empty when not connected. Informational: the access policy keys on (org, transport), so exactly one account is representable per pair. | [optional] 
**AccountLabel** | Pointer to **string** | AccountLabel is the human label of that same account — the Discord guild name, the Slack team name, the Teams tenant name (falling back to the tenant id), the Telegram chat title. DISPLAY ONLY: never a key, and never swapped with Account, on any surface. | [optional] 
**Capabilities** | Pointer to [**TransportCapabilities**](TransportCapabilities.md) | Capabilities is what this transport renders natively — read it before composing a message that needs threading, media or interactive actions. | [optional] 
**Connected** | Pointer to **bool** | Connected is whether integrations holds a connection for (this org, this transport) — whether someone finished its connect flow. False leaves Account and AccountLabel empty, and a send is then refused downstream rather than here: by the transport&#39;s own binding check (403 for a Telegram chat this org has not bound, 409 for a Discord or Teams channel with no inbound-learned route), or on Slack by the absent per-org bot token, which surfaces as 502. | [optional] 
**DmPolicy** | Pointer to **string** | DMPolicy is how this org admits direct messages here: \&quot;pairing\&quot;, \&quot;allowlist\&quot; or \&quot;open\&quot;, defaulting to \&quot;pairing\&quot; when the org has never set one. | [optional] 
**GroupPolicy** | Pointer to **string** | GroupPolicy is how this org admits group and thread channels here: \&quot;open\&quot;, \&quot;allowlist\&quot; or \&quot;disabled\&quot;, defaulting to \&quot;open\&quot;. Both policy fields come back EMPTY — rather than the listing failing — when the policy cannot be read; GET /v1/transport/allowlist carries the same two with the entries they consult. | [optional] 
**Id** | Pointer to **string** | ID is the fixed transport identifier — discord, slack, teams, telegram or whatsapp — and the value every route on this surface names a transport by, including the &#x60;:transport&#x60; segment of the send path. The listing is always in that order. | [optional] 
**PendingPairing** | Pointer to **int64** | PendingPairing counts the org&#39;s UNEXPIRED pairing requests on this transport: exactly the rows GET /v1/transport/pairing returns for it, one per person waiting on an admin. It never exceeds three — the pending cap per (org, transport) — and expired requests are not counted. | [optional] 

## Methods

### NewTransportTransportView

`func NewTransportTransportView() *TransportTransportView`

NewTransportTransportView instantiates a new TransportTransportView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportTransportViewWithDefaults

`func NewTransportTransportViewWithDefaults() *TransportTransportView`

NewTransportTransportViewWithDefaults instantiates a new TransportTransportView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransportTransportView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransportTransportView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransportTransportView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TransportTransportView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountLabel

`func (o *TransportTransportView) GetAccountLabel() string`

GetAccountLabel returns the AccountLabel field if non-nil, zero value otherwise.

### GetAccountLabelOk

`func (o *TransportTransportView) GetAccountLabelOk() (*string, bool)`

GetAccountLabelOk returns a tuple with the AccountLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLabel

`func (o *TransportTransportView) SetAccountLabel(v string)`

SetAccountLabel sets AccountLabel field to given value.

### HasAccountLabel

`func (o *TransportTransportView) HasAccountLabel() bool`

HasAccountLabel returns a boolean if a field has been set.

### GetCapabilities

`func (o *TransportTransportView) GetCapabilities() TransportCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *TransportTransportView) GetCapabilitiesOk() (*TransportCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *TransportTransportView) SetCapabilities(v TransportCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *TransportTransportView) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetConnected

`func (o *TransportTransportView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *TransportTransportView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *TransportTransportView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *TransportTransportView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDmPolicy

`func (o *TransportTransportView) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *TransportTransportView) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *TransportTransportView) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *TransportTransportView) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *TransportTransportView) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *TransportTransportView) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *TransportTransportView) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *TransportTransportView) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.

### GetId

`func (o *TransportTransportView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransportTransportView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransportTransportView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransportTransportView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPendingPairing

`func (o *TransportTransportView) GetPendingPairing() int64`

GetPendingPairing returns the PendingPairing field if non-nil, zero value otherwise.

### GetPendingPairingOk

`func (o *TransportTransportView) GetPendingPairingOk() (*int64, bool)`

GetPendingPairingOk returns a tuple with the PendingPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPairing

`func (o *TransportTransportView) SetPendingPairing(v int64)`

SetPendingPairing sets PendingPairing field to given value.

### HasPendingPairing

`func (o *TransportTransportView) HasPendingPairing() bool`

HasPendingPairing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


