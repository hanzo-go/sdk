# MarketingBetaAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is what was done: star, verified on GitHub, or follow, counted on trust. | [optional] 
**At** | Pointer to **string** | At is when it was counted, RFC 3339 in UTC. | [optional] 
**Points** | Pointer to **int64** | Points is what the action earned. | [optional] 
**Target** | Pointer to **string** | Target is what it was done to: a repository (hanzoai/ai) for a star, a founder&#39;s link (zach/github) for a follow. | [optional] 

## Methods

### NewMarketingBetaAction

`func NewMarketingBetaAction() *MarketingBetaAction`

NewMarketingBetaAction instantiates a new MarketingBetaAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaActionWithDefaults

`func NewMarketingBetaActionWithDefaults() *MarketingBetaAction`

NewMarketingBetaActionWithDefaults instantiates a new MarketingBetaAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MarketingBetaAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MarketingBetaAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MarketingBetaAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MarketingBetaAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAt

`func (o *MarketingBetaAction) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketingBetaAction) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketingBetaAction) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketingBetaAction) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetPoints

`func (o *MarketingBetaAction) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *MarketingBetaAction) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *MarketingBetaAction) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *MarketingBetaAction) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTarget

`func (o *MarketingBetaAction) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MarketingBetaAction) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MarketingBetaAction) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MarketingBetaAction) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


