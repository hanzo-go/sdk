# LimitsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**LimitsBlock**](LimitsBlock.md) | Limits is the plan&#39;s decision. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id the limits were resolved for, after the empty-means- world-free default. | [optional] 
**Unit** | Pointer to **string** | Unit names what the two rate numbers are counted in: requests/minute. | [optional] 

## Methods

### NewLimitsView

`func NewLimitsView() *LimitsView`

NewLimitsView instantiates a new LimitsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsViewWithDefaults

`func NewLimitsViewWithDefaults() *LimitsView`

NewLimitsViewWithDefaults instantiates a new LimitsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *LimitsView) GetLimits() LimitsBlock`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *LimitsView) GetLimitsOk() (*LimitsBlock, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *LimitsView) SetLimits(v LimitsBlock)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *LimitsView) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPlan

`func (o *LimitsView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LimitsView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LimitsView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LimitsView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUnit

`func (o *LimitsView) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LimitsView) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LimitsView) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LimitsView) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


