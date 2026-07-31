# CloudStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired verification provider&#39;s name. | [optional] 
**Verifications** | Pointer to [**CloudVerificationTally**](CloudVerificationTally.md) | Verifications tallies the org&#39;s verifications by provider-reported status. | [optional] 

## Methods

### NewCloudStatusView

`func NewCloudStatusView() *CloudStatusView`

NewCloudStatusView instantiates a new CloudStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStatusViewWithDefaults

`func NewCloudStatusViewWithDefaults() *CloudStatusView`

NewCloudStatusViewWithDefaults instantiates a new CloudStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *CloudStatusView) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *CloudStatusView) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *CloudStatusView) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *CloudStatusView) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetProvider

`func (o *CloudStatusView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudStatusView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudStatusView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudStatusView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVerifications

`func (o *CloudStatusView) GetVerifications() CloudVerificationTally`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CloudStatusView) GetVerificationsOk() (*CloudVerificationTally, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CloudStatusView) SetVerifications(v CloudVerificationTally)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *CloudStatusView) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


