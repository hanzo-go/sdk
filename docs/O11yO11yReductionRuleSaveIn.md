# O11yO11yReductionRuleSaveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID is the rule&#39;s id. | 
**Labels** | **[]string** | Labels are the label names the rule matches. Required, at least one. | 
**MatchType** | **string** | MatchType is drop or keep. Required. | 

## Methods

### NewO11yO11yReductionRuleSaveIn

`func NewO11yO11yReductionRuleSaveIn(id string, labels []string, matchType string, ) *O11yO11yReductionRuleSaveIn`

NewO11yO11yReductionRuleSaveIn instantiates a new O11yO11yReductionRuleSaveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRuleSaveInWithDefaults

`func NewO11yO11yReductionRuleSaveInWithDefaults() *O11yO11yReductionRuleSaveIn`

NewO11yO11yReductionRuleSaveInWithDefaults instantiates a new O11yO11yReductionRuleSaveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yO11yReductionRuleSaveIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yReductionRuleSaveIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yReductionRuleSaveIn) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *O11yO11yReductionRuleSaveIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yO11yReductionRuleSaveIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yO11yReductionRuleSaveIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMatchType

`func (o *O11yO11yReductionRuleSaveIn) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *O11yO11yReductionRuleSaveIn) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *O11yO11yReductionRuleSaveIn) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


