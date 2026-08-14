# ScanDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balanced** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**Extracted** | Pointer to [**Extracted**](Extracted.md) |  | [optional] 
**Questions** | Pointer to [**[]Question**](Question.md) |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Voucher** | Pointer to [**Voucher**](Voucher.md) |  | [optional] 

## Methods

### NewScanDraft

`func NewScanDraft() *ScanDraft`

NewScanDraft instantiates a new ScanDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanDraftWithDefaults

`func NewScanDraftWithDefaults() *ScanDraft`

NewScanDraftWithDefaults instantiates a new ScanDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanced

`func (o *ScanDraft) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *ScanDraft) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *ScanDraft) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *ScanDraft) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetCategory

`func (o *ScanDraft) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ScanDraft) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ScanDraft) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ScanDraft) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfidence

`func (o *ScanDraft) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ScanDraft) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ScanDraft) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ScanDraft) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetExtracted

`func (o *ScanDraft) GetExtracted() Extracted`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *ScanDraft) GetExtractedOk() (*Extracted, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *ScanDraft) SetExtracted(v Extracted)`

SetExtracted sets Extracted field to given value.

### HasExtracted

`func (o *ScanDraft) HasExtracted() bool`

HasExtracted returns a boolean if a field has been set.

### GetQuestions

`func (o *ScanDraft) GetQuestions() []Question`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *ScanDraft) GetQuestionsOk() (*[]Question, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *ScanDraft) SetQuestions(v []Question)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *ScanDraft) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetScanId

`func (o *ScanDraft) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ScanDraft) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ScanDraft) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ScanDraft) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetVendor

`func (o *ScanDraft) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ScanDraft) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ScanDraft) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ScanDraft) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVoucher

`func (o *ScanDraft) GetVoucher() Voucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *ScanDraft) GetVoucherOk() (*Voucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *ScanDraft) SetVoucher(v Voucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *ScanDraft) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


