# EsignRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is where this recipient&#39;s signing link is meant to go, lower-cased. | [optional] 
**Id** | Pointer to **string** | ID is the recipient id, which is what a field is placed against. | [optional] 
**Name** | Pointer to **string** | Name is the recipient&#39;s display name, empty when none was given. | [optional] 
**ReadStatus** | Pointer to **string** | ReadStatus is NOT_OPENED until they first open their link, then OPENED. | [optional] 
**RejectionReason** | Pointer to **string** | RejectionReason is why they declined, null unless they did. | [optional] 
**Role** | Pointer to **string** | Role is SIGNER, CC, VIEWER, APPROVER or ASSISTANT. A document waits only for its SIGNERs and APPROVERs before it can seal. | [optional] 
**SendStatus** | Pointer to **string** | SendStatus is NOT_SENT until the document goes out, then SENT. A CC recipient is SENT from the moment they are added. | [optional] 
**SignedAt** | Pointer to **int32** | SignedAt is when they finished or declined, in unix milliseconds; null while neither has happened. | [optional] 
**SigningOrder** | Pointer to **float32** | SigningOrder is their position in a SEQUENTIAL document, null when they were added without one. A PARALLEL document ignores it. | [optional] 
**SigningStatus** | Pointer to **string** | SigningStatus is NOT_SIGNED, SIGNED or REJECTED. A CC recipient is SIGNED from the moment they are added, because they are never asked. | [optional] 

## Methods

### NewEsignRecipient

`func NewEsignRecipient() *EsignRecipient`

NewEsignRecipient instantiates a new EsignRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignRecipientWithDefaults

`func NewEsignRecipientWithDefaults() *EsignRecipient`

NewEsignRecipientWithDefaults instantiates a new EsignRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EsignRecipient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignRecipient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignRecipient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignRecipient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EsignRecipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsignRecipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsignRecipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsignRecipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadStatus

`func (o *EsignRecipient) GetReadStatus() string`

GetReadStatus returns the ReadStatus field if non-nil, zero value otherwise.

### GetReadStatusOk

`func (o *EsignRecipient) GetReadStatusOk() (*string, bool)`

GetReadStatusOk returns a tuple with the ReadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStatus

`func (o *EsignRecipient) SetReadStatus(v string)`

SetReadStatus sets ReadStatus field to given value.

### HasReadStatus

`func (o *EsignRecipient) HasReadStatus() bool`

HasReadStatus returns a boolean if a field has been set.

### GetRejectionReason

`func (o *EsignRecipient) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *EsignRecipient) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *EsignRecipient) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *EsignRecipient) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetRole

`func (o *EsignRecipient) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignRecipient) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignRecipient) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignRecipient) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSendStatus

`func (o *EsignRecipient) GetSendStatus() string`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *EsignRecipient) GetSendStatusOk() (*string, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *EsignRecipient) SetSendStatus(v string)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *EsignRecipient) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetSignedAt

`func (o *EsignRecipient) GetSignedAt() int32`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *EsignRecipient) GetSignedAtOk() (*int32, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *EsignRecipient) SetSignedAt(v int32)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *EsignRecipient) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### GetSigningOrder

`func (o *EsignRecipient) GetSigningOrder() float32`

GetSigningOrder returns the SigningOrder field if non-nil, zero value otherwise.

### GetSigningOrderOk

`func (o *EsignRecipient) GetSigningOrderOk() (*float32, bool)`

GetSigningOrderOk returns a tuple with the SigningOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOrder

`func (o *EsignRecipient) SetSigningOrder(v float32)`

SetSigningOrder sets SigningOrder field to given value.

### HasSigningOrder

`func (o *EsignRecipient) HasSigningOrder() bool`

HasSigningOrder returns a boolean if a field has been set.

### GetSigningStatus

`func (o *EsignRecipient) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *EsignRecipient) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *EsignRecipient) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *EsignRecipient) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


