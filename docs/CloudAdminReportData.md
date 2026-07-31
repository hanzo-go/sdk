# CloudAdminReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | Pointer to [**CloudAnchorStatus**](CloudAnchorStatus.md) | Anchor is the Hanzo L1 anchoring status of the ledger root. | [optional] 
**Journal** | Pointer to [**[]CloudJournalEntry**](CloudJournalEntry.md) | Journal is the recent double-entry entries, newest first. | [optional] 
**Report** | Pointer to [**CloudTreasuryReport**](CloudTreasuryReport.md) | Report is the reserve-fund snapshot — available, accrued, paid, per-program, policy. | [optional] 

## Methods

### NewCloudAdminReportData

`func NewCloudAdminReportData() *CloudAdminReportData`

NewCloudAdminReportData instantiates a new CloudAdminReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminReportDataWithDefaults

`func NewCloudAdminReportDataWithDefaults() *CloudAdminReportData`

NewCloudAdminReportDataWithDefaults instantiates a new CloudAdminReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *CloudAdminReportData) GetAnchor() CloudAnchorStatus`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *CloudAdminReportData) GetAnchorOk() (*CloudAnchorStatus, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *CloudAdminReportData) SetAnchor(v CloudAnchorStatus)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *CloudAdminReportData) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetJournal

`func (o *CloudAdminReportData) GetJournal() []CloudJournalEntry`

GetJournal returns the Journal field if non-nil, zero value otherwise.

### GetJournalOk

`func (o *CloudAdminReportData) GetJournalOk() (*[]CloudJournalEntry, bool)`

GetJournalOk returns a tuple with the Journal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournal

`func (o *CloudAdminReportData) SetJournal(v []CloudJournalEntry)`

SetJournal sets Journal field to given value.

### HasJournal

`func (o *CloudAdminReportData) HasJournal() bool`

HasJournal returns a boolean if a field has been set.

### GetReport

`func (o *CloudAdminReportData) GetReport() CloudTreasuryReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *CloudAdminReportData) GetReportOk() (*CloudTreasuryReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *CloudAdminReportData) SetReport(v CloudTreasuryReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *CloudAdminReportData) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


