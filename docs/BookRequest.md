# BookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Override** | Pointer to **bool** | Override books this bill even when one of the SAME economic identity (vendor, total, issue date) already posted — the explicit human confirmation that a same-looking bill is a genuine second spend, not the same receipt re-scanned. | [optional] 
**ScanId** | Pointer to **string** | ScanID is the scanned document&#39;s file hash, as GET /v1/books/inbox and the scan draft report it. It is the idempotency key: re-booking the same scan writes nothing. | [optional] 
**Voucher** | Pointer to [**Voucher**](Voucher.md) | Voucher is the reviewed voucher to post. Its source is FORCED to (scan, scanId) server-side, so it can never be booked under another source&#39;s key. | [optional] 

## Methods

### NewBookRequest

`func NewBookRequest() *BookRequest`

NewBookRequest instantiates a new BookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookRequestWithDefaults

`func NewBookRequestWithDefaults() *BookRequest`

NewBookRequestWithDefaults instantiates a new BookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverride

`func (o *BookRequest) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *BookRequest) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *BookRequest) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *BookRequest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetScanId

`func (o *BookRequest) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *BookRequest) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *BookRequest) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *BookRequest) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetVoucher

`func (o *BookRequest) GetVoucher() Voucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *BookRequest) GetVoucherOk() (*Voucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *BookRequest) SetVoucher(v Voucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *BookRequest) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


