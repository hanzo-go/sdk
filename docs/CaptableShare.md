# CaptableShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapitalContribution** | Pointer to **float32** | CapitalContribution is the cash paid for the certificate, if recorded. | [optional] 
**CertificateId** | Pointer to **string** | CertificateID is the certificate number, unique within the company. | [optional] 
**CompanyLegends** | Pointer to **[]string** | CompanyLegends are the restrictive legends printed on the certificate. | [optional] 
**Id** | Pointer to **string** | ID is the share id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the certificate was issued. | [optional] 
**PricePerShare** | Pointer to **float32** | PricePerShare is the price paid per share, if recorded. | [optional] 
**Quantity** | Pointer to **int32** | Quantity is how many shares the certificate covers. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class the shares belong to. | [optional] 
**ShareClassName** | Pointer to **string** | ShareClassName is that class&#39;s name. | [optional] 
**ShareClassType** | Pointer to **string** | ShareClassType is that class&#39;s type, COMMON or PREFERRED. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the holder of the certificate. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that holder&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is ACTIVE or DRAFT. | [optional] 

## Methods

### NewCaptableShare

`func NewCaptableShare() *CaptableShare`

NewCaptableShare instantiates a new CaptableShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableShareWithDefaults

`func NewCaptableShareWithDefaults() *CaptableShare`

NewCaptableShareWithDefaults instantiates a new CaptableShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapitalContribution

`func (o *CaptableShare) GetCapitalContribution() float32`

GetCapitalContribution returns the CapitalContribution field if non-nil, zero value otherwise.

### GetCapitalContributionOk

`func (o *CaptableShare) GetCapitalContributionOk() (*float32, bool)`

GetCapitalContributionOk returns a tuple with the CapitalContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalContribution

`func (o *CaptableShare) SetCapitalContribution(v float32)`

SetCapitalContribution sets CapitalContribution field to given value.

### HasCapitalContribution

`func (o *CaptableShare) HasCapitalContribution() bool`

HasCapitalContribution returns a boolean if a field has been set.

### GetCertificateId

`func (o *CaptableShare) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CaptableShare) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CaptableShare) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CaptableShare) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetCompanyLegends

`func (o *CaptableShare) GetCompanyLegends() []string`

GetCompanyLegends returns the CompanyLegends field if non-nil, zero value otherwise.

### GetCompanyLegendsOk

`func (o *CaptableShare) GetCompanyLegendsOk() (*[]string, bool)`

GetCompanyLegendsOk returns a tuple with the CompanyLegends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegends

`func (o *CaptableShare) SetCompanyLegends(v []string)`

SetCompanyLegends sets CompanyLegends field to given value.

### HasCompanyLegends

`func (o *CaptableShare) HasCompanyLegends() bool`

HasCompanyLegends returns a boolean if a field has been set.

### GetId

`func (o *CaptableShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableShare) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableShare) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableShare) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableShare) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CaptableShare) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableShare) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableShare) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableShare) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetQuantity

`func (o *CaptableShare) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptableShare) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptableShare) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptableShare) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableShare) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableShare) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableShare) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableShare) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetShareClassName

`func (o *CaptableShare) GetShareClassName() string`

GetShareClassName returns the ShareClassName field if non-nil, zero value otherwise.

### GetShareClassNameOk

`func (o *CaptableShare) GetShareClassNameOk() (*string, bool)`

GetShareClassNameOk returns a tuple with the ShareClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassName

`func (o *CaptableShare) SetShareClassName(v string)`

SetShareClassName sets ShareClassName field to given value.

### HasShareClassName

`func (o *CaptableShare) HasShareClassName() bool`

HasShareClassName returns a boolean if a field has been set.

### GetShareClassType

`func (o *CaptableShare) GetShareClassType() string`

GetShareClassType returns the ShareClassType field if non-nil, zero value otherwise.

### GetShareClassTypeOk

`func (o *CaptableShare) GetShareClassTypeOk() (*string, bool)`

GetShareClassTypeOk returns a tuple with the ShareClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassType

`func (o *CaptableShare) SetShareClassType(v string)`

SetShareClassType sets ShareClassType field to given value.

### HasShareClassType

`func (o *CaptableShare) HasShareClassType() bool`

HasShareClassType returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableShare) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableShare) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableShare) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableShare) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableShare) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableShare) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableShare) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableShare) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


