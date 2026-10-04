# SecurityScanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Findings** | Pointer to [**[]SecurityFindingView**](SecurityFindingView.md) | Findings is every finding on that scan, so the detail view is one round-trip. | [optional] 
**Scan** | Pointer to [**SecurityScanView**](SecurityScanView.md) | Scan is the summary. | [optional] 

## Methods

### NewSecurityScanDetail

`func NewSecurityScanDetail() *SecurityScanDetail`

NewSecurityScanDetail instantiates a new SecurityScanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanDetailWithDefaults

`func NewSecurityScanDetailWithDefaults() *SecurityScanDetail`

NewSecurityScanDetailWithDefaults instantiates a new SecurityScanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindings

`func (o *SecurityScanDetail) GetFindings() []SecurityFindingView`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SecurityScanDetail) GetFindingsOk() (*[]SecurityFindingView, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SecurityScanDetail) SetFindings(v []SecurityFindingView)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *SecurityScanDetail) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetScan

`func (o *SecurityScanDetail) GetScan() SecurityScanView`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SecurityScanDetail) GetScanOk() (*SecurityScanView, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SecurityScanDetail) SetScan(v SecurityScanView)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SecurityScanDetail) HasScan() bool`

HasScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


