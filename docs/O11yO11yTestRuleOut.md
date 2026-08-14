# O11yO11yTestRuleOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGettableTestRule**](O11yGettableTestRule.md) | Data holds how many series would alert. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yTestRuleOut

`func NewO11yO11yTestRuleOut() *O11yO11yTestRuleOut`

NewO11yO11yTestRuleOut instantiates a new O11yO11yTestRuleOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTestRuleOutWithDefaults

`func NewO11yO11yTestRuleOutWithDefaults() *O11yO11yTestRuleOut`

NewO11yO11yTestRuleOutWithDefaults instantiates a new O11yO11yTestRuleOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yTestRuleOut) GetData() O11yGettableTestRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yTestRuleOut) GetDataOk() (*O11yGettableTestRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yTestRuleOut) SetData(v O11yGettableTestRule)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yTestRuleOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yTestRuleOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yTestRuleOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yTestRuleOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yTestRuleOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


