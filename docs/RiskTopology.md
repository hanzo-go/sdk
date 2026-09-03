# RiskTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blend** | Pointer to **float64** | Blend is how much of a closing window folds into the reference: 1 replaces it outright, less makes the reference expensive to move. | [optional] 
**Depth** | Pointer to **int64** | Depth is how deep each tree is. With Trees it sets how finely the space is partitioned, and therefore how much history it takes to fill. | [optional] 
**Family** | Pointer to **string** | Family is the KIND of model this candidate is: &#x60;halfspace&#x60; is an ensemble of half-space trees whose masses are counters, and it is the family this search grid ranks. The parameters below are that family&#39;s own — a family that does not partition space with trees has different ones — so read them against this. | [optional] 
**Review** | Pointer to **float64** | Review is the appetite this shape was tried at. | [optional] 
**Trees** | Pointer to **int64** | Trees is how many half-space trees the ensemble holds. | [optional] 
**Window** | Pointer to **int64** | Window is how many events make one reference window. | [optional] 

## Methods

### NewRiskTopology

`func NewRiskTopology() *RiskTopology`

NewRiskTopology instantiates a new RiskTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskTopologyWithDefaults

`func NewRiskTopologyWithDefaults() *RiskTopology`

NewRiskTopologyWithDefaults instantiates a new RiskTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlend

`func (o *RiskTopology) GetBlend() float64`

GetBlend returns the Blend field if non-nil, zero value otherwise.

### GetBlendOk

`func (o *RiskTopology) GetBlendOk() (*float64, bool)`

GetBlendOk returns a tuple with the Blend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlend

`func (o *RiskTopology) SetBlend(v float64)`

SetBlend sets Blend field to given value.

### HasBlend

`func (o *RiskTopology) HasBlend() bool`

HasBlend returns a boolean if a field has been set.

### GetDepth

`func (o *RiskTopology) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *RiskTopology) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *RiskTopology) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *RiskTopology) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetFamily

`func (o *RiskTopology) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *RiskTopology) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *RiskTopology) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *RiskTopology) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetReview

`func (o *RiskTopology) GetReview() float64`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *RiskTopology) GetReviewOk() (*float64, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *RiskTopology) SetReview(v float64)`

SetReview sets Review field to given value.

### HasReview

`func (o *RiskTopology) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetTrees

`func (o *RiskTopology) GetTrees() int64`

GetTrees returns the Trees field if non-nil, zero value otherwise.

### GetTreesOk

`func (o *RiskTopology) GetTreesOk() (*int64, bool)`

GetTreesOk returns a tuple with the Trees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrees

`func (o *RiskTopology) SetTrees(v int64)`

SetTrees sets Trees field to given value.

### HasTrees

`func (o *RiskTopology) HasTrees() bool`

HasTrees returns a boolean if a field has been set.

### GetWindow

`func (o *RiskTopology) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RiskTopology) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RiskTopology) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RiskTopology) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


