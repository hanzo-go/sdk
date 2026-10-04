# Voucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Legs** | Pointer to [**[]Leg**](Leg.md) |  | [optional] 
**PostingAt** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceKind** | Pointer to **string** |  | [optional] 

## Methods

### NewVoucher

`func NewVoucher() *Voucher`

NewVoucher instantiates a new Voucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherWithDefaults

`func NewVoucherWithDefaults() *Voucher`

NewVoucherWithDefaults instantiates a new Voucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Voucher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Voucher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Voucher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Voucher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegs

`func (o *Voucher) GetLegs() []Leg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *Voucher) GetLegsOk() (*[]Leg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *Voucher) SetLegs(v []Leg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *Voucher) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPostingAt

`func (o *Voucher) GetPostingAt() string`

GetPostingAt returns the PostingAt field if non-nil, zero value otherwise.

### GetPostingAtOk

`func (o *Voucher) GetPostingAtOk() (*string, bool)`

GetPostingAtOk returns a tuple with the PostingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAt

`func (o *Voucher) SetPostingAt(v string)`

SetPostingAt sets PostingAt field to given value.

### HasPostingAt

`func (o *Voucher) HasPostingAt() bool`

HasPostingAt returns a boolean if a field has been set.

### GetSourceId

`func (o *Voucher) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Voucher) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Voucher) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Voucher) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceKind

`func (o *Voucher) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *Voucher) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *Voucher) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *Voucher) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


