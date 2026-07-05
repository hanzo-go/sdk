# CrmOpportunityInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amount** | Pointer to **int64** | Deal size in minor units (cents) of currency | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Stage** | Pointer to [**CrmStage**](CrmStage.md) |  | [optional] 
**CloseDate** | Pointer to **int64** | Unix seconds (0 &#x3D; unset) | [optional] 
**CompanyId** | Pointer to **string** | Optional in-org relation to a Company | [optional] 
**PointOfContactId** | Pointer to **string** | Optional in-org relation to a Contact | [optional] 

## Methods

### NewCrmOpportunityInput

`func NewCrmOpportunityInput(name string, ) *CrmOpportunityInput`

NewCrmOpportunityInput instantiates a new CrmOpportunityInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmOpportunityInputWithDefaults

`func NewCrmOpportunityInputWithDefaults() *CrmOpportunityInput`

NewCrmOpportunityInputWithDefaults instantiates a new CrmOpportunityInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrmOpportunityInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmOpportunityInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmOpportunityInput) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *CrmOpportunityInput) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CrmOpportunityInput) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CrmOpportunityInput) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CrmOpportunityInput) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CrmOpportunityInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CrmOpportunityInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CrmOpportunityInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CrmOpportunityInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStage

`func (o *CrmOpportunityInput) GetStage() CrmStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CrmOpportunityInput) GetStageOk() (*CrmStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CrmOpportunityInput) SetStage(v CrmStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CrmOpportunityInput) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetCloseDate

`func (o *CrmOpportunityInput) GetCloseDate() int64`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CrmOpportunityInput) GetCloseDateOk() (*int64, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CrmOpportunityInput) SetCloseDate(v int64)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CrmOpportunityInput) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCompanyId

`func (o *CrmOpportunityInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CrmOpportunityInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CrmOpportunityInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CrmOpportunityInput) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetPointOfContactId

`func (o *CrmOpportunityInput) GetPointOfContactId() string`

GetPointOfContactId returns the PointOfContactId field if non-nil, zero value otherwise.

### GetPointOfContactIdOk

`func (o *CrmOpportunityInput) GetPointOfContactIdOk() (*string, bool)`

GetPointOfContactIdOk returns a tuple with the PointOfContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContactId

`func (o *CrmOpportunityInput) SetPointOfContactId(v string)`

SetPointOfContactId sets PointOfContactId field to given value.

### HasPointOfContactId

`func (o *CrmOpportunityInput) HasPointOfContactId() bool`

HasPointOfContactId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


