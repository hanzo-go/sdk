# CloudBookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Override** | Pointer to **bool** | Override books this bill even when one of the SAME economic identity (vendor, total, issue date) already posted — the explicit human confirmation that a same-looking bill is a genuine second spend, not the same receipt re-scanned. | [optional] 
**ScanId** | Pointer to **string** | ScanID is the scanned document&#39;s file hash, as GET /v1/books/inbox and the scan draft report it. It is the idempotency key: re-booking the same scan writes nothing. | [optional] 
**Voucher** | Pointer to [**CloudVoucher**](CloudVoucher.md) | Voucher is the reviewed voucher to post. Its source is FORCED to (scan, scanId) server-side, so it can never be booked under another source&#39;s key. | [optional] 

## Methods

### NewCloudBookRequest

`func NewCloudBookRequest() *CloudBookRequest`

NewCloudBookRequest instantiates a new CloudBookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBookRequestWithDefaults

`func NewCloudBookRequestWithDefaults() *CloudBookRequest`

NewCloudBookRequestWithDefaults instantiates a new CloudBookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverride

`func (o *CloudBookRequest) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *CloudBookRequest) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *CloudBookRequest) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *CloudBookRequest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetScanId

`func (o *CloudBookRequest) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *CloudBookRequest) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *CloudBookRequest) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *CloudBookRequest) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetVoucher

`func (o *CloudBookRequest) GetVoucher() CloudVoucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *CloudBookRequest) GetVoucherOk() (*CloudVoucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *CloudBookRequest) SetVoucher(v CloudVoucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *CloudBookRequest) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


