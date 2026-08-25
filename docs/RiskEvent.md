# RiskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it happened, RFC 3339. Empty means now. It must sit inside the thirty-day window the aggregates keep and no more than two minutes ahead of this plane&#39;s clock; anything outside that is REFUSED rather than quietly accepted, because a future timestamp moves the aggregates&#39; leading edge and leaves every later event for that subject reading as though it never happened. History older than the window is folded in from your own event surface, not through this endpoint. | [optional] 
**Device** | Pointer to **string** | Device is the device fingerprint, if any. It is the axis that surfaces several nominally unrelated subjects acting as one. | [optional] 
**Id** | Pointer to **string** | ID is the caller&#39;s own stable identifier for the event. It selects the below-the-line review sample by hash, so a counter would make the sample steerable — use the id the event already has. | [optional] 
**Kind** | Pointer to **string** | Kind is whose behaviour this is: person, session or account. It namespaces the subject, so a person and an account that share an identifier stay two subjects. | [optional] 
**Nano** | Pointer to **int32** | Nano is the value moved, in nano-USD. Omit it for an event that moves no money: the value features then read BLIND rather than being told the amount was zero, and the difference is reported on the model state. | [optional] 
**Peer** | Pointer to **string** | Peer is the counterparty, if any. It is an aggregation axis of its own — \&quot;unfamiliar\&quot; is a fact about a relationship and not about either party. | [optional] 
**Subject** | Pointer to **string** | Subject is the identifier on that kind. | [optional] 

## Methods

### NewRiskEvent

`func NewRiskEvent() *RiskEvent`

NewRiskEvent instantiates a new RiskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEventWithDefaults

`func NewRiskEventWithDefaults() *RiskEvent`

NewRiskEventWithDefaults instantiates a new RiskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskEvent) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskEvent) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskEvent) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDevice

`func (o *RiskEvent) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RiskEvent) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RiskEvent) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *RiskEvent) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetId

`func (o *RiskEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RiskEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNano

`func (o *RiskEvent) GetNano() int32`

GetNano returns the Nano field if non-nil, zero value otherwise.

### GetNanoOk

`func (o *RiskEvent) GetNanoOk() (*int32, bool)`

GetNanoOk returns a tuple with the Nano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNano

`func (o *RiskEvent) SetNano(v int32)`

SetNano sets Nano field to given value.

### HasNano

`func (o *RiskEvent) HasNano() bool`

HasNano returns a boolean if a field has been set.

### GetPeer

`func (o *RiskEvent) GetPeer() string`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *RiskEvent) GetPeerOk() (*string, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *RiskEvent) SetPeer(v string)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *RiskEvent) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetSubject

`func (o *RiskEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RiskEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


