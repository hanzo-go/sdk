# ConsoleCreateScoreConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DataType** | **string** |  | 
**Categories** | Pointer to [**[]ConsoleCreateScoreConfigRequestCategoriesInner**](ConsoleCreateScoreConfigRequestCategoriesInner.md) |  | [optional] 
**MinValue** | Pointer to **float32** |  | [optional] 
**MaxValue** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleCreateScoreConfigRequest

`func NewConsoleCreateScoreConfigRequest(name string, dataType string, ) *ConsoleCreateScoreConfigRequest`

NewConsoleCreateScoreConfigRequest instantiates a new ConsoleCreateScoreConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateScoreConfigRequestWithDefaults

`func NewConsoleCreateScoreConfigRequestWithDefaults() *ConsoleCreateScoreConfigRequest`

NewConsoleCreateScoreConfigRequestWithDefaults instantiates a new ConsoleCreateScoreConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsoleCreateScoreConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleCreateScoreConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleCreateScoreConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDataType

`func (o *ConsoleCreateScoreConfigRequest) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ConsoleCreateScoreConfigRequest) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ConsoleCreateScoreConfigRequest) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetCategories

`func (o *ConsoleCreateScoreConfigRequest) GetCategories() []ConsoleCreateScoreConfigRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConsoleCreateScoreConfigRequest) GetCategoriesOk() (*[]ConsoleCreateScoreConfigRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConsoleCreateScoreConfigRequest) SetCategories(v []ConsoleCreateScoreConfigRequestCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ConsoleCreateScoreConfigRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetMinValue

`func (o *ConsoleCreateScoreConfigRequest) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ConsoleCreateScoreConfigRequest) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ConsoleCreateScoreConfigRequest) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ConsoleCreateScoreConfigRequest) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ConsoleCreateScoreConfigRequest) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ConsoleCreateScoreConfigRequest) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ConsoleCreateScoreConfigRequest) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ConsoleCreateScoreConfigRequest) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleCreateScoreConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleCreateScoreConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleCreateScoreConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleCreateScoreConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


