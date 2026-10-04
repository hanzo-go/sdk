# BooksVoucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the human line for the event, e.g. the vendor a bill came from. | [optional] 
**Legs** | Pointer to [**[]BooksLeg**](BooksLeg.md) | Legs are the sides of the posting. They must balance: Σdebit &#x3D;&#x3D; Σcredit, give or take the 2¢ round-off allowance. | [optional] 
**PostingAt** | Pointer to **string** | PostingAt is the RFC3339 instant the event posts at — the time every statement window filters on. | [optional] 
**SourceId** | Pointer to **string** | SourceID is the source event&#39;s own id within that namespace. Together with SourceKind it is the key that makes a repeat posting a no-op. | [optional] 
**SourceKind** | Pointer to **string** | SourceKind is the idempotency namespace naming what booked this, e.g. \&quot;scan\&quot;. | [optional] 

## Methods

### NewBooksVoucher

`func NewBooksVoucher() *BooksVoucher`

NewBooksVoucher instantiates a new BooksVoucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksVoucherWithDefaults

`func NewBooksVoucherWithDefaults() *BooksVoucher`

NewBooksVoucherWithDefaults instantiates a new BooksVoucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BooksVoucher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooksVoucher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooksVoucher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooksVoucher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegs

`func (o *BooksVoucher) GetLegs() []BooksLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *BooksVoucher) GetLegsOk() (*[]BooksLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *BooksVoucher) SetLegs(v []BooksLeg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *BooksVoucher) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPostingAt

`func (o *BooksVoucher) GetPostingAt() string`

GetPostingAt returns the PostingAt field if non-nil, zero value otherwise.

### GetPostingAtOk

`func (o *BooksVoucher) GetPostingAtOk() (*string, bool)`

GetPostingAtOk returns a tuple with the PostingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAt

`func (o *BooksVoucher) SetPostingAt(v string)`

SetPostingAt sets PostingAt field to given value.

### HasPostingAt

`func (o *BooksVoucher) HasPostingAt() bool`

HasPostingAt returns a boolean if a field has been set.

### GetSourceId

`func (o *BooksVoucher) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BooksVoucher) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BooksVoucher) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BooksVoucher) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceKind

`func (o *BooksVoucher) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *BooksVoucher) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *BooksVoucher) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *BooksVoucher) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


