# CaptableOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliffYears** | Pointer to **int64** | CliffYears is how many years before any of the grant vests. | [optional] 
**EquityPlanId** | Pointer to **string** | EquityPlanID is the plan the grant draws from. | [optional] 
**EquityPlanName** | Pointer to **string** | EquityPlanName is that plan&#39;s name. | [optional] 
**ExercisePrice** | Pointer to **float64** | ExercisePrice is the strike price per share. | [optional] 
**ExpirationDate** | Pointer to **string** | ExpirationDate is the ISO date the grant expires. | [optional] 
**GrantId** | Pointer to **string** | GrantID is the grant number, unique within the company. | [optional] 
**Id** | Pointer to **string** | ID is the option id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the grant was issued. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is how many shares the grant covers. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the grantee. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that grantee&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the grant&#39;s state, e.g. DRAFT, ACTIVE, EXERCISED, EXPIRED or CANCELLED. Only non-terminal grants dilute the cap table. | [optional] 
**Type** | Pointer to **string** | Type is the grant kind, ISO or NSO. | [optional] 
**VestingYears** | Pointer to **int64** | VestingYears is the total vesting period in years. | [optional] 

## Methods

### NewCaptableOption

`func NewCaptableOption() *CaptableOption`

NewCaptableOption instantiates a new CaptableOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableOptionWithDefaults

`func NewCaptableOptionWithDefaults() *CaptableOption`

NewCaptableOptionWithDefaults instantiates a new CaptableOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliffYears

`func (o *CaptableOption) GetCliffYears() int64`

GetCliffYears returns the CliffYears field if non-nil, zero value otherwise.

### GetCliffYearsOk

`func (o *CaptableOption) GetCliffYearsOk() (*int64, bool)`

GetCliffYearsOk returns a tuple with the CliffYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYears

`func (o *CaptableOption) SetCliffYears(v int64)`

SetCliffYears sets CliffYears field to given value.

### HasCliffYears

`func (o *CaptableOption) HasCliffYears() bool`

HasCliffYears returns a boolean if a field has been set.

### GetEquityPlanId

`func (o *CaptableOption) GetEquityPlanId() string`

GetEquityPlanId returns the EquityPlanId field if non-nil, zero value otherwise.

### GetEquityPlanIdOk

`func (o *CaptableOption) GetEquityPlanIdOk() (*string, bool)`

GetEquityPlanIdOk returns a tuple with the EquityPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanId

`func (o *CaptableOption) SetEquityPlanId(v string)`

SetEquityPlanId sets EquityPlanId field to given value.

### HasEquityPlanId

`func (o *CaptableOption) HasEquityPlanId() bool`

HasEquityPlanId returns a boolean if a field has been set.

### GetEquityPlanName

`func (o *CaptableOption) GetEquityPlanName() string`

GetEquityPlanName returns the EquityPlanName field if non-nil, zero value otherwise.

### GetEquityPlanNameOk

`func (o *CaptableOption) GetEquityPlanNameOk() (*string, bool)`

GetEquityPlanNameOk returns a tuple with the EquityPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanName

`func (o *CaptableOption) SetEquityPlanName(v string)`

SetEquityPlanName sets EquityPlanName field to given value.

### HasEquityPlanName

`func (o *CaptableOption) HasEquityPlanName() bool`

HasEquityPlanName returns a boolean if a field has been set.

### GetExercisePrice

`func (o *CaptableOption) GetExercisePrice() float64`

GetExercisePrice returns the ExercisePrice field if non-nil, zero value otherwise.

### GetExercisePriceOk

`func (o *CaptableOption) GetExercisePriceOk() (*float64, bool)`

GetExercisePriceOk returns a tuple with the ExercisePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercisePrice

`func (o *CaptableOption) SetExercisePrice(v float64)`

SetExercisePrice sets ExercisePrice field to given value.

### HasExercisePrice

`func (o *CaptableOption) HasExercisePrice() bool`

HasExercisePrice returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CaptableOption) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CaptableOption) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CaptableOption) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CaptableOption) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGrantId

`func (o *CaptableOption) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *CaptableOption) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *CaptableOption) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *CaptableOption) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetId

`func (o *CaptableOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableOption) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableOption) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableOption) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableOption) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetQuantity

`func (o *CaptableOption) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptableOption) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptableOption) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptableOption) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableOption) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableOption) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableOption) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableOption) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableOption) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableOption) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableOption) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableOption) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableOption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableOption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableOption) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableOption) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CaptableOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptableOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptableOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptableOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVestingYears

`func (o *CaptableOption) GetVestingYears() int64`

GetVestingYears returns the VestingYears field if non-nil, zero value otherwise.

### GetVestingYearsOk

`func (o *CaptableOption) GetVestingYearsOk() (*int64, bool)`

GetVestingYearsOk returns a tuple with the VestingYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingYears

`func (o *CaptableOption) SetVestingYears(v int64)`

SetVestingYears sets VestingYears field to given value.

### HasVestingYears

`func (o *CaptableOption) HasVestingYears() bool`

HasVestingYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


