# GraphGraphAppleScorecard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** | Entity is the person&#39;s identifier in the graph. Required. | 
**Features** | Pointer to **[]float64** | Features is the optional 144-element numeric input vector in metadata.json order (0..19 identity, 20..59 communication, 60..143 photos). | [optional] 
**Scores** | **map[string]float64** | Scores maps each relationship column (mother, father, parent, sister, brother, sibling, daughter, son, child, partner, coworker, friend, family, familyAndFriends, alumni, unknown) to its probability in [0,1]. | 

## Methods

### NewGraphGraphAppleScorecard

`func NewGraphGraphAppleScorecard(entity string, scores map[string]float64, ) *GraphGraphAppleScorecard`

NewGraphGraphAppleScorecard instantiates a new GraphGraphAppleScorecard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAppleScorecardWithDefaults

`func NewGraphGraphAppleScorecardWithDefaults() *GraphGraphAppleScorecard`

NewGraphGraphAppleScorecardWithDefaults instantiates a new GraphGraphAppleScorecard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *GraphGraphAppleScorecard) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphAppleScorecard) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphAppleScorecard) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetFeatures

`func (o *GraphGraphAppleScorecard) GetFeatures() []float64`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GraphGraphAppleScorecard) GetFeaturesOk() (*[]float64, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GraphGraphAppleScorecard) SetFeatures(v []float64)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GraphGraphAppleScorecard) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetScores

`func (o *GraphGraphAppleScorecard) GetScores() map[string]float64`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *GraphGraphAppleScorecard) GetScoresOk() (*map[string]float64, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *GraphGraphAppleScorecard) SetScores(v map[string]float64)`

SetScores sets Scores field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


