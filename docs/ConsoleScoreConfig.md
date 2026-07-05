# ConsoleScoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]ConsoleCreateScoreConfigRequestCategoriesInner**](ConsoleCreateScoreConfigRequestCategoriesInner.md) |  | [optional] 
**MinValue** | Pointer to **float32** |  | [optional] 
**MaxValue** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConsoleScoreConfig

`func NewConsoleScoreConfig() *ConsoleScoreConfig`

NewConsoleScoreConfig instantiates a new ConsoleScoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleScoreConfigWithDefaults

`func NewConsoleScoreConfigWithDefaults() *ConsoleScoreConfig`

NewConsoleScoreConfigWithDefaults instantiates a new ConsoleScoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleScoreConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleScoreConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleScoreConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleScoreConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConsoleScoreConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleScoreConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleScoreConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleScoreConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataType

`func (o *ConsoleScoreConfig) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ConsoleScoreConfig) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ConsoleScoreConfig) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ConsoleScoreConfig) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetCategories

`func (o *ConsoleScoreConfig) GetCategories() []ConsoleCreateScoreConfigRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConsoleScoreConfig) GetCategoriesOk() (*[]ConsoleCreateScoreConfigRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConsoleScoreConfig) SetCategories(v []ConsoleCreateScoreConfigRequestCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ConsoleScoreConfig) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetMinValue

`func (o *ConsoleScoreConfig) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ConsoleScoreConfig) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ConsoleScoreConfig) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ConsoleScoreConfig) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ConsoleScoreConfig) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ConsoleScoreConfig) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ConsoleScoreConfig) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ConsoleScoreConfig) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleScoreConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleScoreConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleScoreConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleScoreConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsArchived

`func (o *ConsoleScoreConfig) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *ConsoleScoreConfig) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *ConsoleScoreConfig) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *ConsoleScoreConfig) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConsoleScoreConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsoleScoreConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsoleScoreConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsoleScoreConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConsoleScoreConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConsoleScoreConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConsoleScoreConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConsoleScoreConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


