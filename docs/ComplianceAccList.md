# ComplianceAccList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComplianceAccView**](ComplianceAccView.md) | Data is the org&#39;s tracked accreditation records, newest first. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are tracked or provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 

## Methods

### NewComplianceAccList

`func NewComplianceAccList() *ComplianceAccList`

NewComplianceAccList instantiates a new ComplianceAccList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceAccListWithDefaults

`func NewComplianceAccListWithDefaults() *ComplianceAccList`

NewComplianceAccListWithDefaults instantiates a new ComplianceAccList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComplianceAccList) GetData() []ComplianceAccView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComplianceAccList) GetDataOk() (*[]ComplianceAccView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComplianceAccList) SetData(v []ComplianceAccView)`

SetData sets Data field to given value.

### HasData

`func (o *ComplianceAccList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *ComplianceAccList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *ComplianceAccList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *ComplianceAccList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *ComplianceAccList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


