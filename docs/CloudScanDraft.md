# CloudScanDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balanced** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**Extracted** | Pointer to [**CloudExtracted**](CloudExtracted.md) |  | [optional] 
**Questions** | Pointer to [**[]CloudQuestion**](CloudQuestion.md) |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Voucher** | Pointer to [**CloudVoucher**](CloudVoucher.md) |  | [optional] 

## Methods

### NewCloudScanDraft

`func NewCloudScanDraft() *CloudScanDraft`

NewCloudScanDraft instantiates a new CloudScanDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudScanDraftWithDefaults

`func NewCloudScanDraftWithDefaults() *CloudScanDraft`

NewCloudScanDraftWithDefaults instantiates a new CloudScanDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanced

`func (o *CloudScanDraft) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *CloudScanDraft) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *CloudScanDraft) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *CloudScanDraft) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetCategory

`func (o *CloudScanDraft) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudScanDraft) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudScanDraft) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudScanDraft) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfidence

`func (o *CloudScanDraft) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CloudScanDraft) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CloudScanDraft) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CloudScanDraft) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetExtracted

`func (o *CloudScanDraft) GetExtracted() CloudExtracted`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *CloudScanDraft) GetExtractedOk() (*CloudExtracted, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *CloudScanDraft) SetExtracted(v CloudExtracted)`

SetExtracted sets Extracted field to given value.

### HasExtracted

`func (o *CloudScanDraft) HasExtracted() bool`

HasExtracted returns a boolean if a field has been set.

### GetQuestions

`func (o *CloudScanDraft) GetQuestions() []CloudQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CloudScanDraft) GetQuestionsOk() (*[]CloudQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CloudScanDraft) SetQuestions(v []CloudQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CloudScanDraft) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetScanId

`func (o *CloudScanDraft) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *CloudScanDraft) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *CloudScanDraft) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *CloudScanDraft) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetVendor

`func (o *CloudScanDraft) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CloudScanDraft) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CloudScanDraft) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CloudScanDraft) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVoucher

`func (o *CloudScanDraft) GetVoucher() CloudVoucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *CloudScanDraft) GetVoucherOk() (*CloudVoucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *CloudScanDraft) SetVoucher(v CloudVoucher)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *CloudScanDraft) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


