# ComplianceStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired verification provider&#39;s name. | [optional] 
**Verifications** | Pointer to [**ComplianceVerificationTally**](ComplianceVerificationTally.md) | Verifications tallies the org&#39;s verifications by provider-reported status. | [optional] 

## Methods

### NewComplianceStatusView

`func NewComplianceStatusView() *ComplianceStatusView`

NewComplianceStatusView instantiates a new ComplianceStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceStatusViewWithDefaults

`func NewComplianceStatusViewWithDefaults() *ComplianceStatusView`

NewComplianceStatusViewWithDefaults instantiates a new ComplianceStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *ComplianceStatusView) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *ComplianceStatusView) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *ComplianceStatusView) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *ComplianceStatusView) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetProvider

`func (o *ComplianceStatusView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComplianceStatusView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComplianceStatusView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComplianceStatusView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVerifications

`func (o *ComplianceStatusView) GetVerifications() ComplianceVerificationTally`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *ComplianceStatusView) GetVerificationsOk() (*ComplianceVerificationTally, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *ComplianceStatusView) SetVerifications(v ComplianceVerificationTally)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *ComplianceStatusView) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


