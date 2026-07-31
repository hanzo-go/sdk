# CloudAdCampaign

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

### NewCloudAdCampaign

`func NewCloudAdCampaign() *CloudAdCampaign`

NewCloudAdCampaign instantiates a new CloudAdCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdCampaignWithDefaults

`func NewCloudAdCampaignWithDefaults() *CloudAdCampaign`

NewCloudAdCampaignWithDefaults instantiates a new CloudAdCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudAdCampaign) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudAdCampaign) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudAdCampaign) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudAdCampaign) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *CloudAdCampaign) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudAdCampaign) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudAdCampaign) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudAdCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAdCampaign) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAdCampaign) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAdCampaign) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAdCampaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudAdCampaign) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudAdCampaign) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudAdCampaign) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudAdCampaign) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *CloudAdCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAdCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAdCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAdCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAdCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAdCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAdCampaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAdCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *CloudAdCampaign) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CloudAdCampaign) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CloudAdCampaign) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *CloudAdCampaign) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudAdCampaign) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudAdCampaign) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudAdCampaign) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudAdCampaign) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *CloudAdCampaign) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *CloudAdCampaign) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *CloudAdCampaign) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *CloudAdCampaign) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAdCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAdCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAdCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAdCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAdCampaign) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAdCampaign) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAdCampaign) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAdCampaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


