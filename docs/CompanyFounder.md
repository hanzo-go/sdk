# CompanyFounder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedBy** | Pointer to **string** | DecidedBy is who settled a terminal KYC status: the provider name, or a reviewer&#39;s user id. | [optional] 
**Email** | Pointer to **string** | Email is the founder&#39;s email, and the key a KYC decision addresses a founder by — POST /v1/company/kyc/decision matches on it. | [optional] 
**EquityBps** | Pointer to **int64** | EquityBps is the founder&#39;s ownership in basis points, 0–10000 (1% &#x3D;&#x3D; 100 bps, so 10000 is the whole company). The founders&#39; shares seed the cap-table genesis. | [optional] 
**KycRef** | Pointer to **string** | KYCRef is the idv provider&#39;s session reference for this founder. | [optional] 
**KycStatus** | Pointer to **string** | KYCStatus is the founder&#39;s identity-verification state: pending, verified (a real idv provider reported a pass), reviewer_confirmed (a privileged reviewer confirmed out-of-band) or failed. The payment stage is unreachable until every founder passes. | [optional] 
**Name** | Pointer to **string** | Name is the founder&#39;s full legal name, as it appears on the formation documents. | [optional] 

## Methods

### NewCompanyFounder

`func NewCompanyFounder() *CompanyFounder`

NewCompanyFounder instantiates a new CompanyFounder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFounderWithDefaults

`func NewCompanyFounderWithDefaults() *CompanyFounder`

NewCompanyFounderWithDefaults instantiates a new CompanyFounder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedBy

`func (o *CompanyFounder) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *CompanyFounder) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *CompanyFounder) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *CompanyFounder) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetEmail

`func (o *CompanyFounder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyFounder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyFounder) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyFounder) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEquityBps

`func (o *CompanyFounder) GetEquityBps() int64`

GetEquityBps returns the EquityBps field if non-nil, zero value otherwise.

### GetEquityBpsOk

`func (o *CompanyFounder) GetEquityBpsOk() (*int64, bool)`

GetEquityBpsOk returns a tuple with the EquityBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityBps

`func (o *CompanyFounder) SetEquityBps(v int64)`

SetEquityBps sets EquityBps field to given value.

### HasEquityBps

`func (o *CompanyFounder) HasEquityBps() bool`

HasEquityBps returns a boolean if a field has been set.

### GetKycRef

`func (o *CompanyFounder) GetKycRef() string`

GetKycRef returns the KycRef field if non-nil, zero value otherwise.

### GetKycRefOk

`func (o *CompanyFounder) GetKycRefOk() (*string, bool)`

GetKycRefOk returns a tuple with the KycRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycRef

`func (o *CompanyFounder) SetKycRef(v string)`

SetKycRef sets KycRef field to given value.

### HasKycRef

`func (o *CompanyFounder) HasKycRef() bool`

HasKycRef returns a boolean if a field has been set.

### GetKycStatus

`func (o *CompanyFounder) GetKycStatus() string`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *CompanyFounder) GetKycStatusOk() (*string, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *CompanyFounder) SetKycStatus(v string)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *CompanyFounder) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetName

`func (o *CompanyFounder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyFounder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyFounder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyFounder) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


