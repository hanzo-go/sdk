# ComplianceCheckList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComplianceCheckView**](ComplianceCheckView.md) | Data is the org&#39;s verifications, newest first, without subject PII. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 

## Methods

### NewComplianceCheckList

`func NewComplianceCheckList() *ComplianceCheckList`

NewComplianceCheckList instantiates a new ComplianceCheckList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceCheckListWithDefaults

`func NewComplianceCheckListWithDefaults() *ComplianceCheckList`

NewComplianceCheckListWithDefaults instantiates a new ComplianceCheckList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComplianceCheckList) GetData() []ComplianceCheckView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComplianceCheckList) GetDataOk() (*[]ComplianceCheckView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComplianceCheckList) SetData(v []ComplianceCheckView)`

SetData sets Data field to given value.

### HasData

`func (o *ComplianceCheckList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *ComplianceCheckList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *ComplianceCheckList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *ComplianceCheckList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *ComplianceCheckList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


