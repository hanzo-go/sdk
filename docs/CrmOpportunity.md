# CrmOpportunity

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
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix seconds | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix seconds | [optional] 

## Methods

### NewCrmOpportunity

`func NewCrmOpportunity(name string, ) *CrmOpportunity`

NewCrmOpportunity instantiates a new CrmOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmOpportunityWithDefaults

`func NewCrmOpportunityWithDefaults() *CrmOpportunity`

NewCrmOpportunityWithDefaults instantiates a new CrmOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrmOpportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmOpportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmOpportunity) SetName(v string)`

SetName sets Name field to given value.


### GetAmount

`func (o *CrmOpportunity) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CrmOpportunity) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CrmOpportunity) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CrmOpportunity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CrmOpportunity) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CrmOpportunity) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CrmOpportunity) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CrmOpportunity) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStage

`func (o *CrmOpportunity) GetStage() CrmStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CrmOpportunity) GetStageOk() (*CrmStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CrmOpportunity) SetStage(v CrmStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CrmOpportunity) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetCloseDate

`func (o *CrmOpportunity) GetCloseDate() int64`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CrmOpportunity) GetCloseDateOk() (*int64, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CrmOpportunity) SetCloseDate(v int64)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CrmOpportunity) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCompanyId

`func (o *CrmOpportunity) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CrmOpportunity) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CrmOpportunity) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CrmOpportunity) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetPointOfContactId

`func (o *CrmOpportunity) GetPointOfContactId() string`

GetPointOfContactId returns the PointOfContactId field if non-nil, zero value otherwise.

### GetPointOfContactIdOk

`func (o *CrmOpportunity) GetPointOfContactIdOk() (*string, bool)`

GetPointOfContactIdOk returns a tuple with the PointOfContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContactId

`func (o *CrmOpportunity) SetPointOfContactId(v string)`

SetPointOfContactId sets PointOfContactId field to given value.

### HasPointOfContactId

`func (o *CrmOpportunity) HasPointOfContactId() bool`

HasPointOfContactId returns a boolean if a field has been set.

### GetId

`func (o *CrmOpportunity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrmOpportunity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrmOpportunity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CrmOpportunity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CrmOpportunity) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmOpportunity) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmOpportunity) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CrmOpportunity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CrmOpportunity) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmOpportunity) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmOpportunity) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CrmOpportunity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


