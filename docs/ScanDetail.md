# ScanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Findings** | Pointer to [**[]FindingView**](FindingView.md) | Findings is every finding on that scan, so the detail view is one round-trip. | [optional] 
**Scan** | Pointer to [**ScanView**](ScanView.md) | Scan is the summary. | [optional] 

## Methods

### NewScanDetail

`func NewScanDetail() *ScanDetail`

NewScanDetail instantiates a new ScanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanDetailWithDefaults

`func NewScanDetailWithDefaults() *ScanDetail`

NewScanDetailWithDefaults instantiates a new ScanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindings

`func (o *ScanDetail) GetFindings() []FindingView`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ScanDetail) GetFindingsOk() (*[]FindingView, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ScanDetail) SetFindings(v []FindingView)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *ScanDetail) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetScan

`func (o *ScanDetail) GetScan() ScanView`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *ScanDetail) GetScanOk() (*ScanView, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *ScanDetail) SetScan(v ScanView)`

SetScan sets Scan field to given value.

### HasScan

`func (o *ScanDetail) HasScan() bool`

HasScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


