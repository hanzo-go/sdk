# WorldLimitsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**WorldLimitsBlock**](WorldLimitsBlock.md) | Limits is the plan&#39;s decision. | [optional] 
**Plan** | Pointer to **string** | Plan echoes the plan id the limits were resolved for, after the empty-means- world-free default. | [optional] 
**Unit** | Pointer to **string** | Unit names what the two rate numbers are counted in: requests/minute. | [optional] 

## Methods

### NewWorldLimitsView

`func NewWorldLimitsView() *WorldLimitsView`

NewWorldLimitsView instantiates a new WorldLimitsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldLimitsViewWithDefaults

`func NewWorldLimitsViewWithDefaults() *WorldLimitsView`

NewWorldLimitsViewWithDefaults instantiates a new WorldLimitsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *WorldLimitsView) GetLimits() WorldLimitsBlock`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *WorldLimitsView) GetLimitsOk() (*WorldLimitsBlock, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *WorldLimitsView) SetLimits(v WorldLimitsBlock)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *WorldLimitsView) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPlan

`func (o *WorldLimitsView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *WorldLimitsView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *WorldLimitsView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *WorldLimitsView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUnit

`func (o *WorldLimitsView) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *WorldLimitsView) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *WorldLimitsView) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *WorldLimitsView) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


