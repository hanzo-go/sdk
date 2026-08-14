# ScoreConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Categories is the closed set of labels a CATEGORICAL score may carry. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the rubric was first declared. | [optional] 
**DataType** | Pointer to **string** | DataType is NUMERIC, CATEGORICAL or BOOLEAN, and is authoritative — a score recorded under this name cannot claim a different one. | [optional] 
**MaxValue** | Pointer to **float32** | MaxValue is the inclusive ceiling a NUMERIC score must stay under, absent when unbounded. | [optional] 
**MinValue** | Pointer to **float32** | MinValue is the inclusive floor a NUMERIC score must clear, absent when unbounded. | [optional] 
**Name** | Pointer to **string** | Name is the score name this rubric governs. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewScoreConfigView

`func NewScoreConfigView() *ScoreConfigView`

NewScoreConfigView instantiates a new ScoreConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreConfigViewWithDefaults

`func NewScoreConfigViewWithDefaults() *ScoreConfigView`

NewScoreConfigViewWithDefaults instantiates a new ScoreConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ScoreConfigView) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ScoreConfigView) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ScoreConfigView) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ScoreConfigView) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScoreConfigView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScoreConfigView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScoreConfigView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScoreConfigView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataType

`func (o *ScoreConfigView) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ScoreConfigView) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ScoreConfigView) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ScoreConfigView) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetMaxValue

`func (o *ScoreConfigView) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ScoreConfigView) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ScoreConfigView) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ScoreConfigView) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *ScoreConfigView) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ScoreConfigView) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ScoreConfigView) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ScoreConfigView) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetName

`func (o *ScoreConfigView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScoreConfigView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScoreConfigView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScoreConfigView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScoreConfigView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScoreConfigView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScoreConfigView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ScoreConfigView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


