# GatewayKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The API key (only shown once on creation) | [optional] 
**KeyName** | Pointer to **string** |  | [optional] 
**KeyAlias** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Spend** | Pointer to **float32** |  | [optional] 
**MaxBudget** | Pointer to **float32** |  | [optional] 
**BudgetDuration** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGatewayKey

`func NewGatewayKey() *GatewayKey`

NewGatewayKey instantiates a new GatewayKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayKeyWithDefaults

`func NewGatewayKeyWithDefaults() *GatewayKey`

NewGatewayKeyWithDefaults instantiates a new GatewayKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GatewayKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GatewayKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GatewayKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GatewayKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyName

`func (o *GatewayKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *GatewayKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *GatewayKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *GatewayKey) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetKeyAlias

`func (o *GatewayKey) GetKeyAlias() string`

GetKeyAlias returns the KeyAlias field if non-nil, zero value otherwise.

### GetKeyAliasOk

`func (o *GatewayKey) GetKeyAliasOk() (*string, bool)`

GetKeyAliasOk returns a tuple with the KeyAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlias

`func (o *GatewayKey) SetKeyAlias(v string)`

SetKeyAlias sets KeyAlias field to given value.

### HasKeyAlias

`func (o *GatewayKey) HasKeyAlias() bool`

HasKeyAlias returns a boolean if a field has been set.

### GetUserId

`func (o *GatewayKey) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GatewayKey) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GatewayKey) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GatewayKey) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTeamId

`func (o *GatewayKey) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GatewayKey) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GatewayKey) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GatewayKey) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetOrgId

`func (o *GatewayKey) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GatewayKey) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GatewayKey) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GatewayKey) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModels

`func (o *GatewayKey) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GatewayKey) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GatewayKey) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *GatewayKey) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetSpend

`func (o *GatewayKey) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GatewayKey) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GatewayKey) SetSpend(v float32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *GatewayKey) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetMaxBudget

`func (o *GatewayKey) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *GatewayKey) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *GatewayKey) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *GatewayKey) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### GetBudgetDuration

`func (o *GatewayKey) GetBudgetDuration() string`

GetBudgetDuration returns the BudgetDuration field if non-nil, zero value otherwise.

### GetBudgetDurationOk

`func (o *GatewayKey) GetBudgetDurationOk() (*string, bool)`

GetBudgetDurationOk returns a tuple with the BudgetDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDuration

`func (o *GatewayKey) SetBudgetDuration(v string)`

SetBudgetDuration sets BudgetDuration field to given value.

### HasBudgetDuration

`func (o *GatewayKey) HasBudgetDuration() bool`

HasBudgetDuration returns a boolean if a field has been set.

### GetExpires

`func (o *GatewayKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *GatewayKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *GatewayKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *GatewayKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetMetadata

`func (o *GatewayKey) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GatewayKey) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GatewayKey) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GatewayKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


