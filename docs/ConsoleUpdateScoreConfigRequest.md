# ConsoleUpdateScoreConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsArchived** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]ConsoleCreateScoreConfigRequestCategoriesInner**](ConsoleCreateScoreConfigRequestCategoriesInner.md) |  | [optional] 
**MinValue** | Pointer to **float32** |  | [optional] 
**MaxValue** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleUpdateScoreConfigRequest

`func NewConsoleUpdateScoreConfigRequest() *ConsoleUpdateScoreConfigRequest`

NewConsoleUpdateScoreConfigRequest instantiates a new ConsoleUpdateScoreConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleUpdateScoreConfigRequestWithDefaults

`func NewConsoleUpdateScoreConfigRequestWithDefaults() *ConsoleUpdateScoreConfigRequest`

NewConsoleUpdateScoreConfigRequestWithDefaults instantiates a new ConsoleUpdateScoreConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsArchived

`func (o *ConsoleUpdateScoreConfigRequest) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *ConsoleUpdateScoreConfigRequest) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *ConsoleUpdateScoreConfigRequest) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *ConsoleUpdateScoreConfigRequest) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetName

`func (o *ConsoleUpdateScoreConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleUpdateScoreConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleUpdateScoreConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleUpdateScoreConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategories

`func (o *ConsoleUpdateScoreConfigRequest) GetCategories() []ConsoleCreateScoreConfigRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConsoleUpdateScoreConfigRequest) GetCategoriesOk() (*[]ConsoleCreateScoreConfigRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConsoleUpdateScoreConfigRequest) SetCategories(v []ConsoleCreateScoreConfigRequestCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ConsoleUpdateScoreConfigRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetMinValue

`func (o *ConsoleUpdateScoreConfigRequest) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ConsoleUpdateScoreConfigRequest) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ConsoleUpdateScoreConfigRequest) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ConsoleUpdateScoreConfigRequest) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ConsoleUpdateScoreConfigRequest) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ConsoleUpdateScoreConfigRequest) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ConsoleUpdateScoreConfigRequest) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ConsoleUpdateScoreConfigRequest) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleUpdateScoreConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleUpdateScoreConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleUpdateScoreConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleUpdateScoreConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


