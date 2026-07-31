# CloudCaptableOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliffYears** | Pointer to **int32** | CliffYears is how many years before any of the grant vests. | [optional] 
**EquityPlanId** | Pointer to **string** | EquityPlanID is the plan the grant draws from. | [optional] 
**EquityPlanName** | Pointer to **string** | EquityPlanName is that plan&#39;s name. | [optional] 
**ExercisePrice** | Pointer to **float32** | ExercisePrice is the strike price per share. | [optional] 
**ExpirationDate** | Pointer to **string** | ExpirationDate is the ISO date the grant expires. | [optional] 
**GrantId** | Pointer to **string** | GrantID is the grant number, unique within the company. | [optional] 
**Id** | Pointer to **string** | ID is the option id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the grant was issued. | [optional] 
**Quantity** | Pointer to **int32** | Quantity is how many shares the grant covers. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the grantee. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that grantee&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the grant&#39;s state, e.g. DRAFT, ACTIVE, EXERCISED, EXPIRED or CANCELLED. Only non-terminal grants dilute the cap table. | [optional] 
**Type** | Pointer to **string** | Type is the grant kind, ISO or NSO. | [optional] 
**VestingYears** | Pointer to **int32** | VestingYears is the total vesting period in years. | [optional] 

## Methods

### NewCloudCaptableOption

`func NewCloudCaptableOption() *CloudCaptableOption`

NewCloudCaptableOption instantiates a new CloudCaptableOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableOptionWithDefaults

`func NewCloudCaptableOptionWithDefaults() *CloudCaptableOption`

NewCloudCaptableOptionWithDefaults instantiates a new CloudCaptableOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliffYears

`func (o *CloudCaptableOption) GetCliffYears() int32`

GetCliffYears returns the CliffYears field if non-nil, zero value otherwise.

### GetCliffYearsOk

`func (o *CloudCaptableOption) GetCliffYearsOk() (*int32, bool)`

GetCliffYearsOk returns a tuple with the CliffYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYears

`func (o *CloudCaptableOption) SetCliffYears(v int32)`

SetCliffYears sets CliffYears field to given value.

### HasCliffYears

`func (o *CloudCaptableOption) HasCliffYears() bool`

HasCliffYears returns a boolean if a field has been set.

### GetEquityPlanId

`func (o *CloudCaptableOption) GetEquityPlanId() string`

GetEquityPlanId returns the EquityPlanId field if non-nil, zero value otherwise.

### GetEquityPlanIdOk

`func (o *CloudCaptableOption) GetEquityPlanIdOk() (*string, bool)`

GetEquityPlanIdOk returns a tuple with the EquityPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanId

`func (o *CloudCaptableOption) SetEquityPlanId(v string)`

SetEquityPlanId sets EquityPlanId field to given value.

### HasEquityPlanId

`func (o *CloudCaptableOption) HasEquityPlanId() bool`

HasEquityPlanId returns a boolean if a field has been set.

### GetEquityPlanName

`func (o *CloudCaptableOption) GetEquityPlanName() string`

GetEquityPlanName returns the EquityPlanName field if non-nil, zero value otherwise.

### GetEquityPlanNameOk

`func (o *CloudCaptableOption) GetEquityPlanNameOk() (*string, bool)`

GetEquityPlanNameOk returns a tuple with the EquityPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPlanName

`func (o *CloudCaptableOption) SetEquityPlanName(v string)`

SetEquityPlanName sets EquityPlanName field to given value.

### HasEquityPlanName

`func (o *CloudCaptableOption) HasEquityPlanName() bool`

HasEquityPlanName returns a boolean if a field has been set.

### GetExercisePrice

`func (o *CloudCaptableOption) GetExercisePrice() float32`

GetExercisePrice returns the ExercisePrice field if non-nil, zero value otherwise.

### GetExercisePriceOk

`func (o *CloudCaptableOption) GetExercisePriceOk() (*float32, bool)`

GetExercisePriceOk returns a tuple with the ExercisePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercisePrice

`func (o *CloudCaptableOption) SetExercisePrice(v float32)`

SetExercisePrice sets ExercisePrice field to given value.

### HasExercisePrice

`func (o *CloudCaptableOption) HasExercisePrice() bool`

HasExercisePrice returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CloudCaptableOption) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CloudCaptableOption) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CloudCaptableOption) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CloudCaptableOption) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGrantId

`func (o *CloudCaptableOption) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *CloudCaptableOption) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *CloudCaptableOption) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *CloudCaptableOption) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CloudCaptableOption) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CloudCaptableOption) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CloudCaptableOption) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CloudCaptableOption) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetQuantity

`func (o *CloudCaptableOption) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CloudCaptableOption) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CloudCaptableOption) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CloudCaptableOption) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableOption) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableOption) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableOption) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableOption) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CloudCaptableOption) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CloudCaptableOption) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CloudCaptableOption) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CloudCaptableOption) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCaptableOption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCaptableOption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCaptableOption) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCaptableOption) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CloudCaptableOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCaptableOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCaptableOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCaptableOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVestingYears

`func (o *CloudCaptableOption) GetVestingYears() int32`

GetVestingYears returns the VestingYears field if non-nil, zero value otherwise.

### GetVestingYearsOk

`func (o *CloudCaptableOption) GetVestingYearsOk() (*int32, bool)`

GetVestingYearsOk returns a tuple with the VestingYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingYears

`func (o *CloudCaptableOption) SetVestingYears(v int32)`

SetVestingYears sets VestingYears field to given value.

### HasVestingYears

`func (o *CloudCaptableOption) HasVestingYears() bool`

HasVestingYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


