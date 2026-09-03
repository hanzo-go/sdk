# TrustAskView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the ask arrived, in unix milliseconds. | [optional] 
**DecidedAt** | Pointer to **int64** | DecidedAt is when it was answered, in unix milliseconds; 0 while open. | [optional] 
**DecidedBy** | Pointer to **string** | DecidedBy is who answered it. | [optional] 
**Email** | Pointer to **string** | Email is the address that asked, as stated and UNVERIFIED — it names a party and proves nothing, which is why the grant is addressed to it rather than trusting it. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when a granted ask closes, in unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the request&#39;s id. | [optional] 
**Item** | Pointer to **string** | Item is the item asked for, empty when the whole released tier was asked for. | [optional] 
**Link** | Pointer to **string** | Link is the share link a granted ask became. | [optional] 
**Nda** | Pointer to **string** | Nda is the text this party accepted, verbatim as it stood when they accepted. | [optional] 
**Note** | Pointer to **string** | Note is what the decider wrote when refusing. | [optional] 
**Party** | Pointer to **string** | Party is the company the asker stated. | [optional] 
**Reason** | Pointer to **string** | Reason is why they said they want it. | [optional] 
**State** | Pointer to **string** | State is open, granted or refused. | [optional] 

## Methods

### NewTrustAskView

`func NewTrustAskView() *TrustAskView`

NewTrustAskView instantiates a new TrustAskView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustAskViewWithDefaults

`func NewTrustAskViewWithDefaults() *TrustAskView`

NewTrustAskViewWithDefaults instantiates a new TrustAskView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TrustAskView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrustAskView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrustAskView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrustAskView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDecidedAt

`func (o *TrustAskView) GetDecidedAt() int64`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *TrustAskView) GetDecidedAtOk() (*int64, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *TrustAskView) SetDecidedAt(v int64)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *TrustAskView) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *TrustAskView) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *TrustAskView) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *TrustAskView) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *TrustAskView) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetEmail

`func (o *TrustAskView) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrustAskView) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrustAskView) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrustAskView) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TrustAskView) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TrustAskView) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TrustAskView) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TrustAskView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *TrustAskView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustAskView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustAskView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustAskView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItem

`func (o *TrustAskView) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TrustAskView) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TrustAskView) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *TrustAskView) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLink

`func (o *TrustAskView) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *TrustAskView) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *TrustAskView) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *TrustAskView) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetNda

`func (o *TrustAskView) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *TrustAskView) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *TrustAskView) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *TrustAskView) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetNote

`func (o *TrustAskView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TrustAskView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TrustAskView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TrustAskView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetParty

`func (o *TrustAskView) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *TrustAskView) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *TrustAskView) SetParty(v string)`

SetParty sets Party field to given value.

### HasParty

`func (o *TrustAskView) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetReason

`func (o *TrustAskView) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TrustAskView) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TrustAskView) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TrustAskView) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *TrustAskView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrustAskView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrustAskView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrustAskView) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


