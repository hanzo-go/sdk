# CloudRecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accreditation** | Pointer to [**[]CloudAccView**](CloudAccView.md) | Accreditation is the org&#39;s tracked accreditation-state records. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported or tracked, never a platform assertion of legal or regulatory compliance. | [optional] 
**Verifications** | Pointer to [**[]CloudCheckView**](CloudCheckView.md) | Verifications is the org&#39;s KYC/KYB checks, provider-reported statuses only. | [optional] 

## Methods

### NewCloudRecordList

`func NewCloudRecordList() *CloudRecordList`

NewCloudRecordList instantiates a new CloudRecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRecordListWithDefaults

`func NewCloudRecordListWithDefaults() *CloudRecordList`

NewCloudRecordListWithDefaults instantiates a new CloudRecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditation

`func (o *CloudRecordList) GetAccreditation() []CloudAccView`

GetAccreditation returns the Accreditation field if non-nil, zero value otherwise.

### GetAccreditationOk

`func (o *CloudRecordList) GetAccreditationOk() (*[]CloudAccView, bool)`

GetAccreditationOk returns a tuple with the Accreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditation

`func (o *CloudRecordList) SetAccreditation(v []CloudAccView)`

SetAccreditation sets Accreditation field to given value.

### HasAccreditation

`func (o *CloudRecordList) HasAccreditation() bool`

HasAccreditation returns a boolean if a field has been set.

### GetDisclaimer

`func (o *CloudRecordList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *CloudRecordList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *CloudRecordList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *CloudRecordList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetVerifications

`func (o *CloudRecordList) GetVerifications() []CloudCheckView`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CloudRecordList) GetVerificationsOk() (*[]CloudCheckView, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CloudRecordList) SetVerifications(v []CloudCheckView)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *CloudRecordList) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


