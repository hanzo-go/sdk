# PairingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport the request arrived on — discord, slack, teams or telegram — and half of what approval names. The cap of three unapproved requests applies per (org, channel); while it is full no further code is minted until one is approved or expires. | [optional] 
**Code** | Pointer to **string** | Code is the CAPABILITY that authorises the approval: eight characters from a 32-symbol uppercase alphabet (A-Z0-9 minus the confusables 0, O, 1 and I), minted with crypto/rand and also sent to the requester in chat. An org admin passes it with the channel to POST /v1/channels/pairing/approve, which CONSUMES it — the request row is deleted, so a code approves once — and which takes org admin as well as the code. It lives ONE HOUR from CreatedAt; expired requests are not listed here, and approving one is a 404. It is shown on this admin surface and NEVER logged. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is Unix SECONDS of FIRST contact: when the request was minted and the code sent. Expiry is measured from here and from nowhere else. | [optional] 
**LastSeen** | Pointer to **int64** | LastSeen is Unix SECONDS of the MOST RECENT message from this sender while the request has been pending. It moves as they keep writing, which is how an admin tells a live request from an abandoned one — but it does not extend the hour and does not re-send the code, since one request sends exactly one chat reply. | [optional] 
**Sender** | Pointer to **string** | Sender is the transport-native user id waiting for access — the same identity inbox messages carry. Approving mints a DM allow entry for exactly this value and nothing wider: pairing never grants group access. | [optional] 

## Methods

### NewPairingView

`func NewPairingView() *PairingView`

NewPairingView instantiates a new PairingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairingViewWithDefaults

`func NewPairingViewWithDefaults() *PairingView`

NewPairingViewWithDefaults instantiates a new PairingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PairingView) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PairingView) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PairingView) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PairingView) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCode

`func (o *PairingView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PairingView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PairingView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PairingView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PairingView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PairingView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PairingView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PairingView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSeen

`func (o *PairingView) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *PairingView) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *PairingView) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *PairingView) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetSender

`func (o *PairingView) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PairingView) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PairingView) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *PairingView) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


