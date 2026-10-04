# CompanyWaiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the founder&#39;s email — the key a decision is posted against. | [optional] 
**Founder** | Pointer to **string** | Founder is the founder&#39;s name. | [optional] 
**KycRef** | Pointer to **string** | KYCRef is the identity-verification session reference, when one was opened. | [optional] 
**KycStatus** | Pointer to **string** | KYCStatus is the founder&#39;s unsettled status. | [optional] 
**Name** | Pointer to **string** | Name is the proposed company name. | [optional] 
**Org** | Pointer to **string** | Org is the tenant whose formation the founder belongs to. | [optional] 
**Since** | Pointer to **int64** | Since is when the formation was last touched, as a unix second. | [optional] 

## Methods

### NewCompanyWaiting

`func NewCompanyWaiting() *CompanyWaiting`

NewCompanyWaiting instantiates a new CompanyWaiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWaitingWithDefaults

`func NewCompanyWaitingWithDefaults() *CompanyWaiting`

NewCompanyWaitingWithDefaults instantiates a new CompanyWaiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CompanyWaiting) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyWaiting) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyWaiting) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyWaiting) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFounder

`func (o *CompanyWaiting) GetFounder() string`

GetFounder returns the Founder field if non-nil, zero value otherwise.

### GetFounderOk

`func (o *CompanyWaiting) GetFounderOk() (*string, bool)`

GetFounderOk returns a tuple with the Founder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounder

`func (o *CompanyWaiting) SetFounder(v string)`

SetFounder sets Founder field to given value.

### HasFounder

`func (o *CompanyWaiting) HasFounder() bool`

HasFounder returns a boolean if a field has been set.

### GetKycRef

`func (o *CompanyWaiting) GetKycRef() string`

GetKycRef returns the KycRef field if non-nil, zero value otherwise.

### GetKycRefOk

`func (o *CompanyWaiting) GetKycRefOk() (*string, bool)`

GetKycRefOk returns a tuple with the KycRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycRef

`func (o *CompanyWaiting) SetKycRef(v string)`

SetKycRef sets KycRef field to given value.

### HasKycRef

`func (o *CompanyWaiting) HasKycRef() bool`

HasKycRef returns a boolean if a field has been set.

### GetKycStatus

`func (o *CompanyWaiting) GetKycStatus() string`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CompanyWaiting) GetKycStatusOk() (*string, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CompanyWaiting) SetKycStatus(v string)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *CompanyWaiting) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetName

`func (o *CompanyWaiting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyWaiting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyWaiting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyWaiting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CompanyWaiting) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CompanyWaiting) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CompanyWaiting) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CompanyWaiting) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSince

`func (o *CompanyWaiting) GetSince() int64`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *CompanyWaiting) GetSinceOk() (*int64, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *CompanyWaiting) SetSince(v int64)`

SetSince sets Since field to given value.

### HasSince

`func (o *CompanyWaiting) HasSince() bool`

HasSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


