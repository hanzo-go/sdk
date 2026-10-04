# DataroomTrustAsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | Pointer to **bool** | Accept must be true when the centre states an NDA. The text accepted is recorded verbatim on the request, so a later edit to the NDA cannot rewrite what this party agreed to. | [optional] 
**Email** | Pointer to **string** | Email is where the grant will be sent, and the ONLY address the resulting link admits. Required. | [optional] 
**Item** | Pointer to **string** | Item names one published item to ask for. Optional; omitting it asks for everything released on request. | [optional] 
**Party** | Pointer to **string** | Party is the company the asker is from. Optional, and recorded as stated. | [optional] 
**Reason** | Pointer to **string** | Reason is why they want it. Optional, and recorded as stated — it is what the person deciding reads. | [optional] 
**Slug** | Pointer to **string** | Slug is the centre&#39;s public address, taken from the path. | [optional] 

## Methods

### NewDataroomTrustAsk

`func NewDataroomTrustAsk() *DataroomTrustAsk`

NewDataroomTrustAsk instantiates a new DataroomTrustAsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustAskWithDefaults

`func NewDataroomTrustAskWithDefaults() *DataroomTrustAsk`

NewDataroomTrustAskWithDefaults instantiates a new DataroomTrustAsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccept

`func (o *DataroomTrustAsk) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *DataroomTrustAsk) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *DataroomTrustAsk) SetAccept(v bool)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *DataroomTrustAsk) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetEmail

`func (o *DataroomTrustAsk) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DataroomTrustAsk) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DataroomTrustAsk) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DataroomTrustAsk) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetItem

`func (o *DataroomTrustAsk) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *DataroomTrustAsk) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *DataroomTrustAsk) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *DataroomTrustAsk) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetParty

`func (o *DataroomTrustAsk) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *DataroomTrustAsk) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *DataroomTrustAsk) SetParty(v string)`

SetParty sets Party field to given value.

### HasParty

`func (o *DataroomTrustAsk) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetReason

`func (o *DataroomTrustAsk) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DataroomTrustAsk) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DataroomTrustAsk) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DataroomTrustAsk) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSlug

`func (o *DataroomTrustAsk) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DataroomTrustAsk) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DataroomTrustAsk) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DataroomTrustAsk) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


