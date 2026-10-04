# ComplianceRecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accreditation** | Pointer to [**[]ComplianceAccView**](ComplianceAccView.md) | Accreditation is the org&#39;s tracked accreditation-state records. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported or tracked, never a platform assertion of legal or regulatory compliance. | [optional] 
**Verifications** | Pointer to [**[]ComplianceCheckView**](ComplianceCheckView.md) | Verifications is the org&#39;s KYC/KYB checks, provider-reported statuses only. | [optional] 

## Methods

### NewComplianceRecordList

`func NewComplianceRecordList() *ComplianceRecordList`

NewComplianceRecordList instantiates a new ComplianceRecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceRecordListWithDefaults

`func NewComplianceRecordListWithDefaults() *ComplianceRecordList`

NewComplianceRecordListWithDefaults instantiates a new ComplianceRecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditation

`func (o *ComplianceRecordList) GetAccreditation() []ComplianceAccView`

GetAccreditation returns the Accreditation field if non-nil, zero value otherwise.

### GetAccreditationOk

`func (o *ComplianceRecordList) GetAccreditationOk() (*[]ComplianceAccView, bool)`

GetAccreditationOk returns a tuple with the Accreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditation

`func (o *ComplianceRecordList) SetAccreditation(v []ComplianceAccView)`

SetAccreditation sets Accreditation field to given value.

### HasAccreditation

`func (o *ComplianceRecordList) HasAccreditation() bool`

HasAccreditation returns a boolean if a field has been set.

### GetDisclaimer

`func (o *ComplianceRecordList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *ComplianceRecordList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *ComplianceRecordList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *ComplianceRecordList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetVerifications

`func (o *ComplianceRecordList) GetVerifications() []ComplianceCheckView`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *ComplianceRecordList) GetVerificationsOk() (*[]ComplianceCheckView, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *ComplianceRecordList) SetVerifications(v []ComplianceCheckView)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *ComplianceRecordList) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


