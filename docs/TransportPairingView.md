# TransportPairingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the CAPABILITY that authorises the approval: eight characters from a 32-symbol uppercase alphabet (A-Z0-9 minus the confusables 0, O, 1 and I), minted with crypto/rand and also sent to the requester in chat. An org admin passes it with the transport to POST /v1/transport/pairing/approve, which CONSUMES it — the request row is deleted, so a code approves once — and which takes org admin as well as the code. It lives ONE HOUR from CreatedAt; expired requests are not listed here, and approving one is a 404. It is shown on this admin surface and NEVER logged. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is Unix SECONDS of FIRST contact: when the request was minted and the code sent. Expiry is measured from here and from nowhere else. | [optional] 
**LastSeen** | Pointer to **int64** | LastSeen is Unix SECONDS of the MOST RECENT message from this sender while the request has been pending. It moves as they keep writing, which is how an admin tells a live request from an abandoned one — but it does not extend the hour and does not re-send the code, since one request sends exactly one chat reply. | [optional] 
**Sender** | Pointer to **string** | Sender is the transport-native user id waiting for access — the same identity inbox messages carry. Approving mints a DM allow entry for exactly this value and nothing wider: pairing never grants group access. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport the request arrived on — discord, slack, teams or telegram — and half of what approval names. The cap of three unapproved requests applies per (org, transport); while it is full no further code is minted until one is approved or expires. | [optional] 

## Methods

### NewTransportPairingView

`func NewTransportPairingView() *TransportPairingView`

NewTransportPairingView instantiates a new TransportPairingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPairingViewWithDefaults

`func NewTransportPairingViewWithDefaults() *TransportPairingView`

NewTransportPairingViewWithDefaults instantiates a new TransportPairingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransportPairingView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransportPairingView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransportPairingView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransportPairingView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransportPairingView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransportPairingView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransportPairingView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransportPairingView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSeen

`func (o *TransportPairingView) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *TransportPairingView) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *TransportPairingView) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *TransportPairingView) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetSender

`func (o *TransportPairingView) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TransportPairingView) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TransportPairingView) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TransportPairingView) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetTransport

`func (o *TransportPairingView) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportPairingView) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportPairingView) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportPairingView) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


