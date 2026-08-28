# Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **bool** | Actions is whether the transport renders an INTERACTIVE control natively, and it is the flag to read before composing one. The vocabulary is a closed kind-tagged union (envelope.go), exactly four kinds, each carrying only its own field plus an optional label:   command  — a bot command to run (&#x60;command&#x60;), rendered as a button that             invokes it.  url      — an external link (&#x60;url&#x60;), rendered as a link button.  select   — a menu (&#x60;options&#x60;, each a label and the value choosing it             returns), rendered as a picker.  approval — a reference to an approval request (&#x60;approval.id&#x60;), rendered as             approve/deny controls bound to that id.  False on every transport, and nothing refuses a send for it: actions are accepted, validated per kind, and flattened by renderText to one line each after the text — &#x60;[label] command&#x60;, &#x60;[label] url&#x60;, &#x60;[label] opt | opt&#x60;, &#x60;[label] approval requested: &lt;id&gt;&#x60;. So a caller that needs a real control must read this flag and degrade itself; a caller that only needs the choice communicated can send actions and take the text form. | [optional] 
**Dm** | Pointer to **bool** | DM is whether the transport carries a DIRECT message at all. True for slack, teams, telegram and whatsapp — whatsapp is nothing else, since the Cloud API addresses a person&#39;s number and there is no room a third party joins. False for discord, honestly: that ingress is guild-scoped slash commands — an interaction without a guild id is refused at the endpoint — so nothing ever arrives classified as a DM, no reply route is ever learned for one, and a send addressed at a Discord DM is refused 409. | [optional] 
**Group** | Pointer to **bool** | Group is whether the transport carries multi-person rooms — a Discord guild channel, a Slack channel, a Teams channel or group chat, a Telegram group or supergroup. False on whatsapp alone, which has no such room to carry. | [optional] 
**Media** | Pointer to **bool** | Media is whether the transport renders an ATTACHMENT natively. False everywhere, and a send is not refused for it: renderText flattens each attachment to one &#x60;kind: url (mime)&#x60; line after the text rather than dropping it. A transport whose egress hands its door the raw text would drop the attachment instead, and an attachment-only send would reach the platform with nothing to say — which is why the flag and the flattening are pinned together. | [optional] 
**Thread** | Pointer to **bool** | Thread is whether a reply can be threaded UNDER a specific message. True for slack alone: it is the only transport whose ingress reports a thread (thread_ts, published as the envelope&#39;s replyTo) and whose send posts back into it. Discord&#39;s replyTo makes an inline reply rather than a thread, Telegram&#39;s and WhatsApp&#39;s each quote one message, and Teams carries no reply target at all — a replyTo sent to it is ignored. | [optional] 

## Methods

### NewCapabilities

`func NewCapabilities() *Capabilities`

NewCapabilities instantiates a new Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesWithDefaults

`func NewCapabilitiesWithDefaults() *Capabilities`

NewCapabilitiesWithDefaults instantiates a new Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *Capabilities) GetActions() bool`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Capabilities) GetActionsOk() (*bool, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Capabilities) SetActions(v bool)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Capabilities) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDm

`func (o *Capabilities) GetDm() bool`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *Capabilities) GetDmOk() (*bool, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *Capabilities) SetDm(v bool)`

SetDm sets Dm field to given value.

### HasDm

`func (o *Capabilities) HasDm() bool`

HasDm returns a boolean if a field has been set.

### GetGroup

`func (o *Capabilities) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Capabilities) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Capabilities) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Capabilities) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMedia

`func (o *Capabilities) GetMedia() bool`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Capabilities) GetMediaOk() (*bool, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Capabilities) SetMedia(v bool)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Capabilities) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetThread

`func (o *Capabilities) GetThread() bool`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *Capabilities) GetThreadOk() (*bool, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *Capabilities) SetThread(v bool)`

SetThread sets Thread field to given value.

### HasThread

`func (o *Capabilities) HasThread() bool`

HasThread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


