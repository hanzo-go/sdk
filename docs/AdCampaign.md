# AdCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | provider ad-account ref (Meta act_&lt;id&gt;) | [optional] 
**Budget** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**ExternalId** | Pointer to **string** | provider campaign id after a launch | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Objective** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Spend** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdCampaign

`func NewAdCampaign() *AdCampaign`

NewAdCampaign instantiates a new AdCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCampaignWithDefaults

`func NewAdCampaignWithDefaults() *AdCampaign`

NewAdCampaignWithDefaults instantiates a new AdCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AdCampaign) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AdCampaign) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AdCampaign) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AdCampaign) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *AdCampaign) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdCampaign) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdCampaign) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdCampaign) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdCampaign) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdCampaign) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdCampaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *AdCampaign) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AdCampaign) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AdCampaign) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AdCampaign) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *AdCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdCampaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *AdCampaign) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *AdCampaign) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *AdCampaign) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *AdCampaign) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *AdCampaign) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdCampaign) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdCampaign) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdCampaign) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *AdCampaign) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdCampaign) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdCampaign) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AdCampaign) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *AdCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdCampaign) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdCampaign) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdCampaign) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdCampaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


