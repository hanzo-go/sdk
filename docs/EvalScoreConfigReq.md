# EvalScoreConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Categories is the closed set of labels a CATEGORICAL score may carry. A CATEGORICAL rubric with none is refused. | [optional] 
**DataType** | Pointer to **string** | DataType is NUMERIC (the default), CATEGORICAL or BOOLEAN. | [optional] 
**MaxValue** | Pointer to **float64** | MaxValue is the inclusive ceiling a NUMERIC score must stay under, finite. | [optional] 
**MinValue** | Pointer to **float64** | MinValue is the inclusive floor a NUMERIC score must clear. It must be finite and must not exceed MaxValue. | [optional] 
**Name** | **string** | Name is the score name this rubric governs, matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. The name is the key, so re-posting one replaces its rules. | 

## Methods

### NewEvalScoreConfigReq

`func NewEvalScoreConfigReq(name string, ) *EvalScoreConfigReq`

NewEvalScoreConfigReq instantiates a new EvalScoreConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalScoreConfigReqWithDefaults

`func NewEvalScoreConfigReqWithDefaults() *EvalScoreConfigReq`

NewEvalScoreConfigReqWithDefaults instantiates a new EvalScoreConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *EvalScoreConfigReq) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EvalScoreConfigReq) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EvalScoreConfigReq) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *EvalScoreConfigReq) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDataType

`func (o *EvalScoreConfigReq) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EvalScoreConfigReq) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EvalScoreConfigReq) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EvalScoreConfigReq) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetMaxValue

`func (o *EvalScoreConfigReq) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *EvalScoreConfigReq) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *EvalScoreConfigReq) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *EvalScoreConfigReq) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *EvalScoreConfigReq) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *EvalScoreConfigReq) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *EvalScoreConfigReq) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *EvalScoreConfigReq) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetName

`func (o *EvalScoreConfigReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalScoreConfigReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalScoreConfigReq) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


