# CreditGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AmountCents** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EffectiveAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RemainingCents** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Voided** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreditGrant

`func NewCreditGrant() *CreditGrant`

NewCreditGrant instantiates a new CreditGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditGrantWithDefaults

`func NewCreditGrantWithDefaults() *CreditGrant`

NewCreditGrantWithDefaults instantiates a new CreditGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreditGrant) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreditGrant) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreditGrant) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreditGrant) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountCents

`func (o *CreditGrant) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditGrant) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditGrant) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CreditGrant) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreditGrant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditGrant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditGrant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreditGrant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditGrant) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditGrant) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditGrant) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreditGrant) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *CreditGrant) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *CreditGrant) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *CreditGrant) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *CreditGrant) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreditGrant) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreditGrant) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreditGrant) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreditGrant) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *CreditGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreditGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreditGrant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditGrant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditGrant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreditGrant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *CreditGrant) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreditGrant) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreditGrant) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreditGrant) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRemainingCents

`func (o *CreditGrant) GetRemainingCents() int32`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *CreditGrant) GetRemainingCentsOk() (*int32, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *CreditGrant) SetRemainingCents(v int32)`

SetRemainingCents sets RemainingCents field to given value.

### HasRemainingCents

`func (o *CreditGrant) HasRemainingCents() bool`

HasRemainingCents returns a boolean if a field has been set.

### GetTags

`func (o *CreditGrant) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreditGrant) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreditGrant) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreditGrant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserId

`func (o *CreditGrant) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreditGrant) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreditGrant) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreditGrant) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVoided

`func (o *CreditGrant) GetVoided() bool`

GetVoided returns the Voided field if non-nil, zero value otherwise.

### GetVoidedOk

`func (o *CreditGrant) GetVoidedOk() (*bool, bool)`

GetVoidedOk returns a tuple with the Voided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoided

`func (o *CreditGrant) SetVoided(v bool)`

SetVoided sets Voided field to given value.

### HasVoided

`func (o *CreditGrant) HasVoided() bool`

HasVoided returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


