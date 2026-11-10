# GraphGraphOntologyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | Pointer to [**[]GraphGraphOntologyClass**](GraphGraphOntologyClass.md) | Classes is the matching entity classes (181 in the unfiltered catalog). | [optional] 
**Features** | Pointer to [**[]GraphGraphOntologyFeature**](GraphGraphOntologyFeature.md) | Features is the matching EntityTagging classifier features (144 in the unfiltered catalog). | [optional] 
**Predicates** | Pointer to [**[]GraphGraphOntologyPredicate**](GraphGraphOntologyPredicate.md) | Predicates is the matching predicates (278 in the unfiltered catalog). | [optional] 
**ScoreColumns** | Pointer to **[]string** | ScoreColumns is the 17 relationship columns of entityTagging.db&#39;s kinship scorecard. | [optional] 

## Methods

### NewGraphGraphOntologyOut

`func NewGraphGraphOntologyOut() *GraphGraphOntologyOut`

NewGraphGraphOntologyOut instantiates a new GraphGraphOntologyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphOntologyOutWithDefaults

`func NewGraphGraphOntologyOutWithDefaults() *GraphGraphOntologyOut`

NewGraphGraphOntologyOutWithDefaults instantiates a new GraphGraphOntologyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *GraphGraphOntologyOut) GetClasses() []GraphGraphOntologyClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *GraphGraphOntologyOut) GetClassesOk() (*[]GraphGraphOntologyClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *GraphGraphOntologyOut) SetClasses(v []GraphGraphOntologyClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *GraphGraphOntologyOut) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetFeatures

`func (o *GraphGraphOntologyOut) GetFeatures() []GraphGraphOntologyFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GraphGraphOntologyOut) GetFeaturesOk() (*[]GraphGraphOntologyFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GraphGraphOntologyOut) SetFeatures(v []GraphGraphOntologyFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GraphGraphOntologyOut) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPredicates

`func (o *GraphGraphOntologyOut) GetPredicates() []GraphGraphOntologyPredicate`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *GraphGraphOntologyOut) GetPredicatesOk() (*[]GraphGraphOntologyPredicate, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *GraphGraphOntologyOut) SetPredicates(v []GraphGraphOntologyPredicate)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *GraphGraphOntologyOut) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetScoreColumns

`func (o *GraphGraphOntologyOut) GetScoreColumns() []string`

GetScoreColumns returns the ScoreColumns field if non-nil, zero value otherwise.

### GetScoreColumnsOk

`func (o *GraphGraphOntologyOut) GetScoreColumnsOk() (*[]string, bool)`

GetScoreColumnsOk returns a tuple with the ScoreColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreColumns

`func (o *GraphGraphOntologyOut) SetScoreColumns(v []string)`

SetScoreColumns sets ScoreColumns field to given value.

### HasScoreColumns

`func (o *GraphGraphOntologyOut) HasScoreColumns() bool`

HasScoreColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


