# RecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accreditation** | Pointer to [**[]AccView**](AccView.md) | Accreditation is the org&#39;s tracked accreditation-state records. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are provider-reported or tracked, never a platform assertion of legal or regulatory compliance. | [optional] 
**Verifications** | Pointer to [**[]CheckView**](CheckView.md) | Verifications is the org&#39;s KYC/KYB checks, provider-reported statuses only. | [optional] 

## Methods

### NewRecordList

`func NewRecordList() *RecordList`

NewRecordList instantiates a new RecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordListWithDefaults

`func NewRecordListWithDefaults() *RecordList`

NewRecordListWithDefaults instantiates a new RecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccreditation

`func (o *RecordList) GetAccreditation() []AccView`

GetAccreditation returns the Accreditation field if non-nil, zero value otherwise.

### GetAccreditationOk

`func (o *RecordList) GetAccreditationOk() (*[]AccView, bool)`

GetAccreditationOk returns a tuple with the Accreditation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccreditation

`func (o *RecordList) SetAccreditation(v []AccView)`

SetAccreditation sets Accreditation field to given value.

### HasAccreditation

`func (o *RecordList) HasAccreditation() bool`

HasAccreditation returns a boolean if a field has been set.

### GetDisclaimer

`func (o *RecordList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *RecordList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *RecordList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *RecordList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetVerifications

`func (o *RecordList) GetVerifications() []CheckView`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *RecordList) GetVerificationsOk() (*[]CheckView, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *RecordList) SetVerifications(v []CheckView)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *RecordList) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


