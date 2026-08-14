# AdminReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | Pointer to [**AnchorStatus**](AnchorStatus.md) | Anchor is the Hanzo L1 anchoring status of the ledger root. | [optional] 
**Journal** | Pointer to [**[]JournalEntry**](JournalEntry.md) | Journal is the recent double-entry entries, newest first. | [optional] 
**Report** | Pointer to [**TreasuryReport**](TreasuryReport.md) | Report is the reserve-fund snapshot — available, accrued, paid, per-program, policy. | [optional] 

## Methods

### NewAdminReportData

`func NewAdminReportData() *AdminReportData`

NewAdminReportData instantiates a new AdminReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminReportDataWithDefaults

`func NewAdminReportDataWithDefaults() *AdminReportData`

NewAdminReportDataWithDefaults instantiates a new AdminReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *AdminReportData) GetAnchor() AnchorStatus`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *AdminReportData) GetAnchorOk() (*AnchorStatus, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *AdminReportData) SetAnchor(v AnchorStatus)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *AdminReportData) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetJournal

`func (o *AdminReportData) GetJournal() []JournalEntry`

GetJournal returns the Journal field if non-nil, zero value otherwise.

### GetJournalOk

`func (o *AdminReportData) GetJournalOk() (*[]JournalEntry, bool)`

GetJournalOk returns a tuple with the Journal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournal

`func (o *AdminReportData) SetJournal(v []JournalEntry)`

SetJournal sets Journal field to given value.

### HasJournal

`func (o *AdminReportData) HasJournal() bool`

HasJournal returns a boolean if a field has been set.

### GetReport

`func (o *AdminReportData) GetReport() TreasuryReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *AdminReportData) GetReportOk() (*TreasuryReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *AdminReportData) SetReport(v TreasuryReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *AdminReportData) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


