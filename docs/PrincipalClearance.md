# PrincipalClearance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to **string** | Amount is the gross amount, U.S. dollars. | [optional] 
**Category** | Pointer to **string** | Category is what is bought: services, attorney, rents, royalties, other or merchandise. | [optional] 
**DecidedAt** | Pointer to **int64** | DecidedAt is when, unix seconds. | [optional] 
**FactsRequired** | Pointer to [**[]PrincipalFact**](PrincipalFact.md) |  | [optional] 
**Id** | Pointer to **string** | ID is the clearance, \&quot;clr_\&quot;-prefixed. | [optional] 
**Notice** | Pointer to **string** | Notice is what a clearance is and is not. | [optional] 
**Payee** | Pointer to **string** | Payee is the org to be paid. | [optional] 
**Payer** | Pointer to **string** | Payer is the caller&#39;s org. | [optional] 
**Performed** | Pointer to **string** | Performed is where the service is performed, as stated. | [optional] 
**Platform** | Pointer to **bool** | Platform is whether the caller operates the payee&#39;s platform, as stated. | [optional] 
**Rail** | Pointer to **string** | Rail is how it moves. | [optional] 
**ReportingObligations** | Pointer to [**[]PrincipalObligation**](PrincipalObligation.md) |  | [optional] 
**RequiredBeforePayment** | Pointer to [**[]PrincipalStep**](PrincipalStep.md) |  | [optional] 
**Rules** | Pointer to [**[]PrincipalDecided**](PrincipalDecided.md) |  | [optional] 
**SettlementMethods** | Pointer to [**[]PrincipalMethod**](PrincipalMethod.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Withholding** | Pointer to [**PrincipalWithholding**](PrincipalWithholding.md) |  | [optional] 

## Methods

### NewPrincipalClearance

`func NewPrincipalClearance() *PrincipalClearance`

NewPrincipalClearance instantiates a new PrincipalClearance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalClearanceWithDefaults

`func NewPrincipalClearanceWithDefaults() *PrincipalClearance`

NewPrincipalClearanceWithDefaults instantiates a new PrincipalClearance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *PrincipalClearance) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *PrincipalClearance) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *PrincipalClearance) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *PrincipalClearance) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAmount

`func (o *PrincipalClearance) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrincipalClearance) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrincipalClearance) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PrincipalClearance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategory

`func (o *PrincipalClearance) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PrincipalClearance) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PrincipalClearance) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PrincipalClearance) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDecidedAt

`func (o *PrincipalClearance) GetDecidedAt() int64`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *PrincipalClearance) GetDecidedAtOk() (*int64, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *PrincipalClearance) SetDecidedAt(v int64)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *PrincipalClearance) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetFactsRequired

`func (o *PrincipalClearance) GetFactsRequired() []PrincipalFact`

GetFactsRequired returns the FactsRequired field if non-nil, zero value otherwise.

### GetFactsRequiredOk

`func (o *PrincipalClearance) GetFactsRequiredOk() (*[]PrincipalFact, bool)`

GetFactsRequiredOk returns a tuple with the FactsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactsRequired

`func (o *PrincipalClearance) SetFactsRequired(v []PrincipalFact)`

SetFactsRequired sets FactsRequired field to given value.

### HasFactsRequired

`func (o *PrincipalClearance) HasFactsRequired() bool`

HasFactsRequired returns a boolean if a field has been set.

### GetId

`func (o *PrincipalClearance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalClearance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalClearance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalClearance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotice

`func (o *PrincipalClearance) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *PrincipalClearance) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *PrincipalClearance) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *PrincipalClearance) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### GetPayee

`func (o *PrincipalClearance) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *PrincipalClearance) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *PrincipalClearance) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *PrincipalClearance) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetPayer

`func (o *PrincipalClearance) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *PrincipalClearance) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *PrincipalClearance) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *PrincipalClearance) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetPerformed

`func (o *PrincipalClearance) GetPerformed() string`

GetPerformed returns the Performed field if non-nil, zero value otherwise.

### GetPerformedOk

`func (o *PrincipalClearance) GetPerformedOk() (*string, bool)`

GetPerformedOk returns a tuple with the Performed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformed

`func (o *PrincipalClearance) SetPerformed(v string)`

SetPerformed sets Performed field to given value.

### HasPerformed

`func (o *PrincipalClearance) HasPerformed() bool`

HasPerformed returns a boolean if a field has been set.

### GetPlatform

`func (o *PrincipalClearance) GetPlatform() bool`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PrincipalClearance) GetPlatformOk() (*bool, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PrincipalClearance) SetPlatform(v bool)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PrincipalClearance) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRail

`func (o *PrincipalClearance) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *PrincipalClearance) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *PrincipalClearance) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *PrincipalClearance) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetReportingObligations

`func (o *PrincipalClearance) GetReportingObligations() []PrincipalObligation`

GetReportingObligations returns the ReportingObligations field if non-nil, zero value otherwise.

### GetReportingObligationsOk

`func (o *PrincipalClearance) GetReportingObligationsOk() (*[]PrincipalObligation, bool)`

GetReportingObligationsOk returns a tuple with the ReportingObligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingObligations

`func (o *PrincipalClearance) SetReportingObligations(v []PrincipalObligation)`

SetReportingObligations sets ReportingObligations field to given value.

### HasReportingObligations

`func (o *PrincipalClearance) HasReportingObligations() bool`

HasReportingObligations returns a boolean if a field has been set.

### GetRequiredBeforePayment

`func (o *PrincipalClearance) GetRequiredBeforePayment() []PrincipalStep`

GetRequiredBeforePayment returns the RequiredBeforePayment field if non-nil, zero value otherwise.

### GetRequiredBeforePaymentOk

`func (o *PrincipalClearance) GetRequiredBeforePaymentOk() (*[]PrincipalStep, bool)`

GetRequiredBeforePaymentOk returns a tuple with the RequiredBeforePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeforePayment

`func (o *PrincipalClearance) SetRequiredBeforePayment(v []PrincipalStep)`

SetRequiredBeforePayment sets RequiredBeforePayment field to given value.

### HasRequiredBeforePayment

`func (o *PrincipalClearance) HasRequiredBeforePayment() bool`

HasRequiredBeforePayment returns a boolean if a field has been set.

### GetRules

`func (o *PrincipalClearance) GetRules() []PrincipalDecided`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PrincipalClearance) GetRulesOk() (*[]PrincipalDecided, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PrincipalClearance) SetRules(v []PrincipalDecided)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PrincipalClearance) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSettlementMethods

`func (o *PrincipalClearance) GetSettlementMethods() []PrincipalMethod`

GetSettlementMethods returns the SettlementMethods field if non-nil, zero value otherwise.

### GetSettlementMethodsOk

`func (o *PrincipalClearance) GetSettlementMethodsOk() (*[]PrincipalMethod, bool)`

GetSettlementMethodsOk returns a tuple with the SettlementMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementMethods

`func (o *PrincipalClearance) SetSettlementMethods(v []PrincipalMethod)`

SetSettlementMethods sets SettlementMethods field to given value.

### HasSettlementMethods

`func (o *PrincipalClearance) HasSettlementMethods() bool`

HasSettlementMethods returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalClearance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalClearance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalClearance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalClearance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWithholding

`func (o *PrincipalClearance) GetWithholding() PrincipalWithholding`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *PrincipalClearance) GetWithholdingOk() (*PrincipalWithholding, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *PrincipalClearance) SetWithholding(v PrincipalWithholding)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *PrincipalClearance) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


