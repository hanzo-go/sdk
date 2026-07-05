# RegistryScanReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** |  | [optional] 
**ScanStatus** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]RegistryScanReportVulnerabilitiesInner**](RegistryScanReportVulnerabilitiesInner.md) |  | [optional] 

## Methods

### NewRegistryScanReport

`func NewRegistryScanReport() *RegistryScanReport`

NewRegistryScanReport instantiates a new RegistryScanReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryScanReportWithDefaults

`func NewRegistryScanReportWithDefaults() *RegistryScanReport`

NewRegistryScanReportWithDefaults instantiates a new RegistryScanReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *RegistryScanReport) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *RegistryScanReport) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *RegistryScanReport) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *RegistryScanReport) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetScanStatus

`func (o *RegistryScanReport) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *RegistryScanReport) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *RegistryScanReport) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *RegistryScanReport) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *RegistryScanReport) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RegistryScanReport) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RegistryScanReport) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RegistryScanReport) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *RegistryScanReport) GetVulnerabilities() []RegistryScanReportVulnerabilitiesInner`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *RegistryScanReport) GetVulnerabilitiesOk() (*[]RegistryScanReportVulnerabilitiesInner, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *RegistryScanReport) SetVulnerabilities(v []RegistryScanReportVulnerabilitiesInner)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *RegistryScanReport) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


