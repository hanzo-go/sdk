# CaptableCaptableShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapitalContribution** | Pointer to **float64** | CapitalContribution is the cash paid for the certificate, if recorded. | [optional] 
**CertificateId** | Pointer to **string** | CertificateID is the certificate number, unique within the company. | [optional] 
**CompanyLegends** | Pointer to **[]string** | CompanyLegends are the restrictive legends printed on the certificate. | [optional] 
**Id** | Pointer to **string** | ID is the share id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the certificate was issued. | [optional] 
**PricePerShare** | Pointer to **float64** | PricePerShare is the price paid per share, if recorded. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is how many shares the certificate covers. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class the shares belong to. | [optional] 
**ShareClassName** | Pointer to **string** | ShareClassName is that class&#39;s name. | [optional] 
**ShareClassType** | Pointer to **string** | ShareClassType is that class&#39;s type, COMMON or PREFERRED. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the holder of the certificate. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that holder&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE or DRAFT. | [optional] 

## Methods

### NewCaptableCaptableShare

`func NewCaptableCaptableShare() *CaptableCaptableShare`

NewCaptableCaptableShare instantiates a new CaptableCaptableShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableShareWithDefaults

`func NewCaptableCaptableShareWithDefaults() *CaptableCaptableShare`

NewCaptableCaptableShareWithDefaults instantiates a new CaptableCaptableShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapitalContribution

`func (o *CaptableCaptableShare) GetCapitalContribution() float64`

GetCapitalContribution returns the CapitalContribution field if non-nil, zero value otherwise.

### GetCapitalContributionOk

`func (o *CaptableCaptableShare) GetCapitalContributionOk() (*float64, bool)`

GetCapitalContributionOk returns a tuple with the CapitalContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalContribution

`func (o *CaptableCaptableShare) SetCapitalContribution(v float64)`

SetCapitalContribution sets CapitalContribution field to given value.

### HasCapitalContribution

`func (o *CaptableCaptableShare) HasCapitalContribution() bool`

HasCapitalContribution returns a boolean if a field has been set.

### GetCertificateId

`func (o *CaptableCaptableShare) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CaptableCaptableShare) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CaptableCaptableShare) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CaptableCaptableShare) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetCompanyLegends

`func (o *CaptableCaptableShare) GetCompanyLegends() []string`

GetCompanyLegends returns the CompanyLegends field if non-nil, zero value otherwise.

### GetCompanyLegendsOk

`func (o *CaptableCaptableShare) GetCompanyLegendsOk() (*[]string, bool)`

GetCompanyLegendsOk returns a tuple with the CompanyLegends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegends

`func (o *CaptableCaptableShare) SetCompanyLegends(v []string)`

SetCompanyLegends sets CompanyLegends field to given value.

### HasCompanyLegends

`func (o *CaptableCaptableShare) HasCompanyLegends() bool`

HasCompanyLegends returns a boolean if a field has been set.

### GetId

`func (o *CaptableCaptableShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableCaptableShare) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableCaptableShare) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableCaptableShare) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableCaptableShare) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CaptableCaptableShare) GetPricePerShare() float64`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableCaptableShare) GetPricePerShareOk() (*float64, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableCaptableShare) SetPricePerShare(v float64)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableCaptableShare) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetQuantity

`func (o *CaptableCaptableShare) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptableCaptableShare) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptableCaptableShare) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptableCaptableShare) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableCaptableShare) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableCaptableShare) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableCaptableShare) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableCaptableShare) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetShareClassName

`func (o *CaptableCaptableShare) GetShareClassName() string`

GetShareClassName returns the ShareClassName field if non-nil, zero value otherwise.

### GetShareClassNameOk

`func (o *CaptableCaptableShare) GetShareClassNameOk() (*string, bool)`

GetShareClassNameOk returns a tuple with the ShareClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassName

`func (o *CaptableCaptableShare) SetShareClassName(v string)`

SetShareClassName sets ShareClassName field to given value.

### HasShareClassName

`func (o *CaptableCaptableShare) HasShareClassName() bool`

HasShareClassName returns a boolean if a field has been set.

### GetShareClassType

`func (o *CaptableCaptableShare) GetShareClassType() string`

GetShareClassType returns the ShareClassType field if non-nil, zero value otherwise.

### GetShareClassTypeOk

`func (o *CaptableCaptableShare) GetShareClassTypeOk() (*string, bool)`

GetShareClassTypeOk returns a tuple with the ShareClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassType

`func (o *CaptableCaptableShare) SetShareClassType(v string)`

SetShareClassType sets ShareClassType field to given value.

### HasShareClassType

`func (o *CaptableCaptableShare) HasShareClassType() bool`

HasShareClassType returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableCaptableShare) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableCaptableShare) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableCaptableShare) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableCaptableShare) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableCaptableShare) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableCaptableShare) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableCaptableShare) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableCaptableShare) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableCaptableShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableCaptableShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableCaptableShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableCaptableShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


