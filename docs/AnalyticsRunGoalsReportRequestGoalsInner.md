# AnalyticsRunGoalsReportRequestGoalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** |  | 
**Goal** | **float32** |  | 
**Operator** | Pointer to **string** | Required when type is event-data | [optional] 
**Property** | Pointer to **string** | Required when type is event-data | [optional] 

## Methods

### NewAnalyticsRunGoalsReportRequestGoalsInner

`func NewAnalyticsRunGoalsReportRequestGoalsInner(type_ string, value string, goal float32, ) *AnalyticsRunGoalsReportRequestGoalsInner`

NewAnalyticsRunGoalsReportRequestGoalsInner instantiates a new AnalyticsRunGoalsReportRequestGoalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRunGoalsReportRequestGoalsInnerWithDefaults

`func NewAnalyticsRunGoalsReportRequestGoalsInnerWithDefaults() *AnalyticsRunGoalsReportRequestGoalsInner`

NewAnalyticsRunGoalsReportRequestGoalsInnerWithDefaults instantiates a new AnalyticsRunGoalsReportRequestGoalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetGoal

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetGoal() float32`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetGoalOk() (*float32, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) SetGoal(v float32)`

SetGoal sets Goal field to given value.


### GetOperator

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetProperty

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AnalyticsRunGoalsReportRequestGoalsInner) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


