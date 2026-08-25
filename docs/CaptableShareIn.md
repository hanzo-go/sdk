# CaptableShareIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardApprovalDate** | Pointer to **interface{}** |  | [optional] 
**CapitalContribution** | Pointer to **interface{}** |  | [optional] 
**CertificateId** | Pointer to **interface{}** |  | [optional] 
**CliffYears** | Pointer to **interface{}** |  | [optional] 
**CompanyLegends** | Pointer to **[]interface{}** | CompanyLegends are the restrictive legends printed on the certificate.  A LIST, and it must be sent as one: the bundle substitutes an EMPTY list for anything that is not an array, so a single string would be accepted and silently discarded — the certificate issued with no legends and nothing reporting it. | [optional] 
**DebtCancelled** | Pointer to **interface{}** |  | [optional] 
**IpContribution** | Pointer to **interface{}** |  | [optional] 
**IssueDate** | Pointer to **interface{}** |  | [optional] 
**OtherContributions** | Pointer to **interface{}** |  | [optional] 
**PricePerShare** | Pointer to **interface{}** |  | [optional] 
**Quantity** | Pointer to **interface{}** |  | [optional] 
**Rule144Date** | Pointer to **interface{}** |  | [optional] 
**ShareClassId** | Pointer to **interface{}** |  | [optional] 
**StakeholderId** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **interface{}** |  | [optional] 
**VestingStartDate** | Pointer to **interface{}** |  | [optional] 
**VestingYears** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCaptableShareIn

`func NewCaptableShareIn() *CaptableShareIn`

NewCaptableShareIn instantiates a new CaptableShareIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableShareInWithDefaults

`func NewCaptableShareInWithDefaults() *CaptableShareIn`

NewCaptableShareInWithDefaults instantiates a new CaptableShareIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardApprovalDate

`func (o *CaptableShareIn) GetBoardApprovalDate() interface{}`

GetBoardApprovalDate returns the BoardApprovalDate field if non-nil, zero value otherwise.

### GetBoardApprovalDateOk

`func (o *CaptableShareIn) GetBoardApprovalDateOk() (*interface{}, bool)`

GetBoardApprovalDateOk returns a tuple with the BoardApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardApprovalDate

`func (o *CaptableShareIn) SetBoardApprovalDate(v interface{})`

SetBoardApprovalDate sets BoardApprovalDate field to given value.

### HasBoardApprovalDate

`func (o *CaptableShareIn) HasBoardApprovalDate() bool`

HasBoardApprovalDate returns a boolean if a field has been set.

### SetBoardApprovalDateNil

`func (o *CaptableShareIn) SetBoardApprovalDateNil(b bool)`

 SetBoardApprovalDateNil sets the value for BoardApprovalDate to be an explicit nil

### UnsetBoardApprovalDate
`func (o *CaptableShareIn) UnsetBoardApprovalDate()`

UnsetBoardApprovalDate ensures that no value is present for BoardApprovalDate, not even an explicit nil
### GetCapitalContribution

`func (o *CaptableShareIn) GetCapitalContribution() interface{}`

GetCapitalContribution returns the CapitalContribution field if non-nil, zero value otherwise.

### GetCapitalContributionOk

`func (o *CaptableShareIn) GetCapitalContributionOk() (*interface{}, bool)`

GetCapitalContributionOk returns a tuple with the CapitalContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalContribution

`func (o *CaptableShareIn) SetCapitalContribution(v interface{})`

SetCapitalContribution sets CapitalContribution field to given value.

### HasCapitalContribution

`func (o *CaptableShareIn) HasCapitalContribution() bool`

HasCapitalContribution returns a boolean if a field has been set.

### SetCapitalContributionNil

`func (o *CaptableShareIn) SetCapitalContributionNil(b bool)`

 SetCapitalContributionNil sets the value for CapitalContribution to be an explicit nil

### UnsetCapitalContribution
`func (o *CaptableShareIn) UnsetCapitalContribution()`

UnsetCapitalContribution ensures that no value is present for CapitalContribution, not even an explicit nil
### GetCertificateId

`func (o *CaptableShareIn) GetCertificateId() interface{}`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CaptableShareIn) GetCertificateIdOk() (*interface{}, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CaptableShareIn) SetCertificateId(v interface{})`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CaptableShareIn) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *CaptableShareIn) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *CaptableShareIn) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetCliffYears

`func (o *CaptableShareIn) GetCliffYears() interface{}`

GetCliffYears returns the CliffYears field if non-nil, zero value otherwise.

### GetCliffYearsOk

`func (o *CaptableShareIn) GetCliffYearsOk() (*interface{}, bool)`

GetCliffYearsOk returns a tuple with the CliffYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYears

`func (o *CaptableShareIn) SetCliffYears(v interface{})`

SetCliffYears sets CliffYears field to given value.

### HasCliffYears

`func (o *CaptableShareIn) HasCliffYears() bool`

HasCliffYears returns a boolean if a field has been set.

### SetCliffYearsNil

`func (o *CaptableShareIn) SetCliffYearsNil(b bool)`

 SetCliffYearsNil sets the value for CliffYears to be an explicit nil

### UnsetCliffYears
`func (o *CaptableShareIn) UnsetCliffYears()`

UnsetCliffYears ensures that no value is present for CliffYears, not even an explicit nil
### GetCompanyLegends

`func (o *CaptableShareIn) GetCompanyLegends() []interface{}`

GetCompanyLegends returns the CompanyLegends field if non-nil, zero value otherwise.

### GetCompanyLegendsOk

`func (o *CaptableShareIn) GetCompanyLegendsOk() (*[]interface{}, bool)`

GetCompanyLegendsOk returns a tuple with the CompanyLegends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegends

`func (o *CaptableShareIn) SetCompanyLegends(v []interface{})`

SetCompanyLegends sets CompanyLegends field to given value.

### HasCompanyLegends

`func (o *CaptableShareIn) HasCompanyLegends() bool`

HasCompanyLegends returns a boolean if a field has been set.

### GetDebtCancelled

`func (o *CaptableShareIn) GetDebtCancelled() interface{}`

GetDebtCancelled returns the DebtCancelled field if non-nil, zero value otherwise.

### GetDebtCancelledOk

`func (o *CaptableShareIn) GetDebtCancelledOk() (*interface{}, bool)`

GetDebtCancelledOk returns a tuple with the DebtCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtCancelled

`func (o *CaptableShareIn) SetDebtCancelled(v interface{})`

SetDebtCancelled sets DebtCancelled field to given value.

### HasDebtCancelled

`func (o *CaptableShareIn) HasDebtCancelled() bool`

HasDebtCancelled returns a boolean if a field has been set.

### SetDebtCancelledNil

`func (o *CaptableShareIn) SetDebtCancelledNil(b bool)`

 SetDebtCancelledNil sets the value for DebtCancelled to be an explicit nil

### UnsetDebtCancelled
`func (o *CaptableShareIn) UnsetDebtCancelled()`

UnsetDebtCancelled ensures that no value is present for DebtCancelled, not even an explicit nil
### GetIpContribution

`func (o *CaptableShareIn) GetIpContribution() interface{}`

GetIpContribution returns the IpContribution field if non-nil, zero value otherwise.

### GetIpContributionOk

`func (o *CaptableShareIn) GetIpContributionOk() (*interface{}, bool)`

GetIpContributionOk returns a tuple with the IpContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpContribution

`func (o *CaptableShareIn) SetIpContribution(v interface{})`

SetIpContribution sets IpContribution field to given value.

### HasIpContribution

`func (o *CaptableShareIn) HasIpContribution() bool`

HasIpContribution returns a boolean if a field has been set.

### SetIpContributionNil

`func (o *CaptableShareIn) SetIpContributionNil(b bool)`

 SetIpContributionNil sets the value for IpContribution to be an explicit nil

### UnsetIpContribution
`func (o *CaptableShareIn) UnsetIpContribution()`

UnsetIpContribution ensures that no value is present for IpContribution, not even an explicit nil
### GetIssueDate

`func (o *CaptableShareIn) GetIssueDate() interface{}`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableShareIn) GetIssueDateOk() (*interface{}, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableShareIn) SetIssueDate(v interface{})`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableShareIn) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *CaptableShareIn) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *CaptableShareIn) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetOtherContributions

`func (o *CaptableShareIn) GetOtherContributions() interface{}`

GetOtherContributions returns the OtherContributions field if non-nil, zero value otherwise.

### GetOtherContributionsOk

`func (o *CaptableShareIn) GetOtherContributionsOk() (*interface{}, bool)`

GetOtherContributionsOk returns a tuple with the OtherContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherContributions

`func (o *CaptableShareIn) SetOtherContributions(v interface{})`

SetOtherContributions sets OtherContributions field to given value.

### HasOtherContributions

`func (o *CaptableShareIn) HasOtherContributions() bool`

HasOtherContributions returns a boolean if a field has been set.

### SetOtherContributionsNil

`func (o *CaptableShareIn) SetOtherContributionsNil(b bool)`

 SetOtherContributionsNil sets the value for OtherContributions to be an explicit nil

### UnsetOtherContributions
`func (o *CaptableShareIn) UnsetOtherContributions()`

UnsetOtherContributions ensures that no value is present for OtherContributions, not even an explicit nil
### GetPricePerShare

`func (o *CaptableShareIn) GetPricePerShare() interface{}`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableShareIn) GetPricePerShareOk() (*interface{}, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableShareIn) SetPricePerShare(v interface{})`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableShareIn) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### SetPricePerShareNil

`func (o *CaptableShareIn) SetPricePerShareNil(b bool)`

 SetPricePerShareNil sets the value for PricePerShare to be an explicit nil

### UnsetPricePerShare
`func (o *CaptableShareIn) UnsetPricePerShare()`

UnsetPricePerShare ensures that no value is present for PricePerShare, not even an explicit nil
### GetQuantity

`func (o *CaptableShareIn) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptableShareIn) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptableShareIn) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptableShareIn) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *CaptableShareIn) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CaptableShareIn) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetRule144Date

`func (o *CaptableShareIn) GetRule144Date() interface{}`

GetRule144Date returns the Rule144Date field if non-nil, zero value otherwise.

### GetRule144DateOk

`func (o *CaptableShareIn) GetRule144DateOk() (*interface{}, bool)`

GetRule144DateOk returns a tuple with the Rule144Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule144Date

`func (o *CaptableShareIn) SetRule144Date(v interface{})`

SetRule144Date sets Rule144Date field to given value.

### HasRule144Date

`func (o *CaptableShareIn) HasRule144Date() bool`

HasRule144Date returns a boolean if a field has been set.

### SetRule144DateNil

`func (o *CaptableShareIn) SetRule144DateNil(b bool)`

 SetRule144DateNil sets the value for Rule144Date to be an explicit nil

### UnsetRule144Date
`func (o *CaptableShareIn) UnsetRule144Date()`

UnsetRule144Date ensures that no value is present for Rule144Date, not even an explicit nil
### GetShareClassId

`func (o *CaptableShareIn) GetShareClassId() interface{}`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableShareIn) GetShareClassIdOk() (*interface{}, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableShareIn) SetShareClassId(v interface{})`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableShareIn) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### SetShareClassIdNil

`func (o *CaptableShareIn) SetShareClassIdNil(b bool)`

 SetShareClassIdNil sets the value for ShareClassId to be an explicit nil

### UnsetShareClassId
`func (o *CaptableShareIn) UnsetShareClassId()`

UnsetShareClassId ensures that no value is present for ShareClassId, not even an explicit nil
### GetStakeholderId

`func (o *CaptableShareIn) GetStakeholderId() interface{}`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableShareIn) GetStakeholderIdOk() (*interface{}, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableShareIn) SetStakeholderId(v interface{})`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableShareIn) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### SetStakeholderIdNil

`func (o *CaptableShareIn) SetStakeholderIdNil(b bool)`

 SetStakeholderIdNil sets the value for StakeholderId to be an explicit nil

### UnsetStakeholderId
`func (o *CaptableShareIn) UnsetStakeholderId()`

UnsetStakeholderId ensures that no value is present for StakeholderId, not even an explicit nil
### GetStatus

`func (o *CaptableShareIn) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableShareIn) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableShareIn) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableShareIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CaptableShareIn) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CaptableShareIn) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVestingStartDate

`func (o *CaptableShareIn) GetVestingStartDate() interface{}`

GetVestingStartDate returns the VestingStartDate field if non-nil, zero value otherwise.

### GetVestingStartDateOk

`func (o *CaptableShareIn) GetVestingStartDateOk() (*interface{}, bool)`

GetVestingStartDateOk returns a tuple with the VestingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingStartDate

`func (o *CaptableShareIn) SetVestingStartDate(v interface{})`

SetVestingStartDate sets VestingStartDate field to given value.

### HasVestingStartDate

`func (o *CaptableShareIn) HasVestingStartDate() bool`

HasVestingStartDate returns a boolean if a field has been set.

### SetVestingStartDateNil

`func (o *CaptableShareIn) SetVestingStartDateNil(b bool)`

 SetVestingStartDateNil sets the value for VestingStartDate to be an explicit nil

### UnsetVestingStartDate
`func (o *CaptableShareIn) UnsetVestingStartDate()`

UnsetVestingStartDate ensures that no value is present for VestingStartDate, not even an explicit nil
### GetVestingYears

`func (o *CaptableShareIn) GetVestingYears() interface{}`

GetVestingYears returns the VestingYears field if non-nil, zero value otherwise.

### GetVestingYearsOk

`func (o *CaptableShareIn) GetVestingYearsOk() (*interface{}, bool)`

GetVestingYearsOk returns a tuple with the VestingYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingYears

`func (o *CaptableShareIn) SetVestingYears(v interface{})`

SetVestingYears sets VestingYears field to given value.

### HasVestingYears

`func (o *CaptableShareIn) HasVestingYears() bool`

HasVestingYears returns a boolean if a field has been set.

### SetVestingYearsNil

`func (o *CaptableShareIn) SetVestingYearsNil(b bool)`

 SetVestingYearsNil sets the value for VestingYears to be an explicit nil

### UnsetVestingYears
`func (o *CaptableShareIn) UnsetVestingYears()`

UnsetVestingYears ensures that no value is present for VestingYears, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


