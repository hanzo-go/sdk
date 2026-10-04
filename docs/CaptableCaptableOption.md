# CaptableCaptableOption

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

### NewCaptableCaptableOption

`func NewCaptableCaptableOption() *CaptableCaptableOption`

NewCaptableCaptableOption instantiates a new CaptableCaptableOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableOptionWithDefaults

`func NewCaptableCaptableOptionWithDefaults() *CaptableCaptableOption`

NewCaptableCaptableOptionWithDefaults instantiates a new CaptableCaptableOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliffYears

`func (o *CaptableCaptableOption) GetCliffYears() int64`

GetCliffYears returns the CliffYears field if non-nil, zero value otherwise.

### GetCliffYearsOk

`func (o *CaptableCaptableOption) GetCliffYearsOk() (*int64, bool)`

GetCliffYearsOk returns a tuple with the CliffYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYears

`func (o *CaptableCaptableOption) SetCliffYears(v int64)`

SetCliffYears sets CliffYears field to given value.

### HasCliffYears

`func (o *CaptableCaptableOption) HasCliffYears() bool`

HasCliffYears returns a boolean if a field has been set.

### GetEquityPlanId

`func (o *CaptableCaptableOption) GetEquityPlanId() string`

GetEquityPlanId returns the EquityPlanId field if non-nil, zero value otherwise.

### GetEquityPlanIdOk

`func (o *CaptableCaptableOption) GetEquityPlanIdOk() (*string, bool)`

GetEquityPlanIdOk returns a tuple with the EquityPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanId

`func (o *CaptableCaptableOption) SetEquityPlanId(v string)`

SetEquityPlanId sets EquityPlanId field to given value.

### HasEquityPlanId

`func (o *CaptableCaptableOption) HasEquityPlanId() bool`

HasEquityPlanId returns a boolean if a field has been set.

### GetEquityPlanName

`func (o *CaptableCaptableOption) GetEquityPlanName() string`

GetEquityPlanName returns the EquityPlanName field if non-nil, zero value otherwise.

### GetEquityPlanNameOk

`func (o *CaptableCaptableOption) GetEquityPlanNameOk() (*string, bool)`

GetEquityPlanNameOk returns a tuple with the EquityPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanName

`func (o *CaptableCaptableOption) SetEquityPlanName(v string)`

SetEquityPlanName sets EquityPlanName field to given value.

### HasEquityPlanName

`func (o *CaptableCaptableOption) HasEquityPlanName() bool`

HasEquityPlanName returns a boolean if a field has been set.

### GetExercisePrice

`func (o *CaptableCaptableOption) GetExercisePrice() float64`

GetExercisePrice returns the ExercisePrice field if non-nil, zero value otherwise.

### GetExercisePriceOk

`func (o *CaptableCaptableOption) GetExercisePriceOk() (*float64, bool)`

GetExercisePriceOk returns a tuple with the ExercisePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercisePrice

`func (o *CaptableCaptableOption) SetExercisePrice(v float64)`

SetExercisePrice sets ExercisePrice field to given value.

### HasExercisePrice

`func (o *CaptableCaptableOption) HasExercisePrice() bool`

HasExercisePrice returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CaptableCaptableOption) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CaptableCaptableOption) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CaptableCaptableOption) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CaptableCaptableOption) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGrantId

`func (o *CaptableCaptableOption) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *CaptableCaptableOption) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *CaptableCaptableOption) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *CaptableCaptableOption) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetId

`func (o *CaptableCaptableOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableCaptableOption) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableCaptableOption) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableCaptableOption) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableCaptableOption) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetQuantity

`func (o *CaptableCaptableOption) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptableCaptableOption) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptableCaptableOption) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptableCaptableOption) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableCaptableOption) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableCaptableOption) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableCaptableOption) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableCaptableOption) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableCaptableOption) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableCaptableOption) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableCaptableOption) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableCaptableOption) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableCaptableOption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableCaptableOption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableCaptableOption) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableCaptableOption) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CaptableCaptableOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptableCaptableOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptableCaptableOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptableCaptableOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVestingYears

`func (o *CaptableCaptableOption) GetVestingYears() int64`

GetVestingYears returns the VestingYears field if non-nil, zero value otherwise.

### GetVestingYearsOk

`func (o *CaptableCaptableOption) GetVestingYearsOk() (*int64, bool)`

GetVestingYearsOk returns a tuple with the VestingYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingYears

`func (o *CaptableCaptableOption) SetVestingYears(v int64)`

SetVestingYears sets VestingYears field to given value.

### HasVestingYears

`func (o *CaptableCaptableOption) HasVestingYears() bool`

HasVestingYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


