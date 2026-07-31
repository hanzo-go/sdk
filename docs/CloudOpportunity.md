# CloudOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount is the deal value in minor units (cents) of Currency. | [optional] 
**CloseDate** | Pointer to **int32** | CloseDate is the expected close as a unix second (0 &#x3D; unset). | [optional] 
**CompanyId** | Pointer to **string** | CompanyID links the deal to one of the org&#39;s companies; empty when unlinked, and cleared when that company is deleted. A write naming a company the org does not own is refused with 422. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the opportunity was created. Server-owned. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code Amount is denominated in; a write that names none stores USD. | [optional] 
**Id** | Pointer to **string** | ID is the server-minted opportunity id (\&quot;oppo_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the deal name. | [optional] 
**PointOfContactId** | Pointer to **string** | PointOfContact links the deal to one of the org&#39;s contacts; empty when unlinked, and cleared when that contact is deleted. A write naming a contact the org does not own is refused with 422. | [optional] 
**Stage** | Pointer to **string** | Stage is the pipeline stage, always one of NEW, SCREENING, MEETING, PROPOSAL or CUSTOMER — stored upper-case whatever case the write used. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the last write. Server-owned. | [optional] 

## Methods

### NewCloudOpportunity

`func NewCloudOpportunity() *CloudOpportunity`

NewCloudOpportunity instantiates a new CloudOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOpportunityWithDefaults

`func NewCloudOpportunityWithDefaults() *CloudOpportunity`

NewCloudOpportunityWithDefaults instantiates a new CloudOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CloudOpportunity) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudOpportunity) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudOpportunity) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudOpportunity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCloseDate

`func (o *CloudOpportunity) GetCloseDate() int32`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CloudOpportunity) GetCloseDateOk() (*int32, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CloudOpportunity) SetCloseDate(v int32)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CloudOpportunity) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCompanyId

`func (o *CloudOpportunity) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CloudOpportunity) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CloudOpportunity) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CloudOpportunity) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudOpportunity) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudOpportunity) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudOpportunity) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudOpportunity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudOpportunity) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudOpportunity) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudOpportunity) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudOpportunity) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *CloudOpportunity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudOpportunity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudOpportunity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudOpportunity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudOpportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudOpportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudOpportunity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudOpportunity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPointOfContactId

`func (o *CloudOpportunity) GetPointOfContactId() string`

GetPointOfContactId returns the PointOfContactId field if non-nil, zero value otherwise.

### GetPointOfContactIdOk

`func (o *CloudOpportunity) GetPointOfContactIdOk() (*string, bool)`

GetPointOfContactIdOk returns a tuple with the PointOfContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContactId

`func (o *CloudOpportunity) SetPointOfContactId(v string)`

SetPointOfContactId sets PointOfContactId field to given value.

### HasPointOfContactId

`func (o *CloudOpportunity) HasPointOfContactId() bool`

HasPointOfContactId returns a boolean if a field has been set.

### GetStage

`func (o *CloudOpportunity) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CloudOpportunity) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CloudOpportunity) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CloudOpportunity) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudOpportunity) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudOpportunity) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudOpportunity) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudOpportunity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


