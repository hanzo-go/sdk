# MarketingBetaIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]MarketingBetaAnswer**](MarketingBetaAnswer.md) | Answers replace the record&#39;s answers. | [optional] 
**Betas** | Pointer to **[]string** | Betas replace the ids of the unreleased products the account asked for. | [optional] 
**Lead** | Pointer to **string** | Lead is the reference the application was filed under as a lead (POST /v1/marketing/leads), so the record and the lead are one person&#39;s. | [optional] 
**Program** | Pointer to **string** | Program is the program applied to: beta, research or partner. | [optional] 

## Methods

### NewMarketingBetaIn

`func NewMarketingBetaIn() *MarketingBetaIn`

NewMarketingBetaIn instantiates a new MarketingBetaIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInWithDefaults

`func NewMarketingBetaInWithDefaults() *MarketingBetaIn`

NewMarketingBetaInWithDefaults instantiates a new MarketingBetaIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *MarketingBetaIn) GetAnswers() []MarketingBetaAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *MarketingBetaIn) GetAnswersOk() (*[]MarketingBetaAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *MarketingBetaIn) SetAnswers(v []MarketingBetaAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *MarketingBetaIn) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetBetas

`func (o *MarketingBetaIn) GetBetas() []string`

GetBetas returns the Betas field if non-nil, zero value otherwise.

### GetBetasOk

`func (o *MarketingBetaIn) GetBetasOk() (*[]string, bool)`

GetBetasOk returns a tuple with the Betas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetas

`func (o *MarketingBetaIn) SetBetas(v []string)`

SetBetas sets Betas field to given value.

### HasBetas

`func (o *MarketingBetaIn) HasBetas() bool`

HasBetas returns a boolean if a field has been set.

### GetLead

`func (o *MarketingBetaIn) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *MarketingBetaIn) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *MarketingBetaIn) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *MarketingBetaIn) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetProgram

`func (o *MarketingBetaIn) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *MarketingBetaIn) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *MarketingBetaIn) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *MarketingBetaIn) HasProgram() bool`

HasProgram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


