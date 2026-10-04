# ComplianceAuditList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComplianceWire**](ComplianceWire.md) | Data is the org&#39;s compliance.* audit rows, newest first. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported or tracked, never a platform assertion of legal or regulatory compliance. | [optional] 

## Methods

### NewComplianceAuditList

`func NewComplianceAuditList() *ComplianceAuditList`

NewComplianceAuditList instantiates a new ComplianceAuditList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceAuditListWithDefaults

`func NewComplianceAuditListWithDefaults() *ComplianceAuditList`

NewComplianceAuditListWithDefaults instantiates a new ComplianceAuditList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComplianceAuditList) GetData() []ComplianceWire`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComplianceAuditList) GetDataOk() (*[]ComplianceWire, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComplianceAuditList) SetData(v []ComplianceWire)`

SetData sets Data field to given value.

### HasData

`func (o *ComplianceAuditList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *ComplianceAuditList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *ComplianceAuditList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *ComplianceAuditList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *ComplianceAuditList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


