# MarketingBetaActIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is star or follow. | [optional] 
**Github** | Pointer to **string** | GitHub is the person&#39;s GitHub login, for a star. Once a star is counted it is the record&#39;s, and another login is refused. | [optional] 
**Target** | Pointer to **string** | Target is the repository starred (hanzoai/&lt;name&gt;) or the founder&#39;s link followed (antje or zach, then github, linkedin or x: zach/github). | [optional] 

## Methods

### NewMarketingBetaActIn

`func NewMarketingBetaActIn() *MarketingBetaActIn`

NewMarketingBetaActIn instantiates a new MarketingBetaActIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaActInWithDefaults

`func NewMarketingBetaActInWithDefaults() *MarketingBetaActIn`

NewMarketingBetaActInWithDefaults instantiates a new MarketingBetaActIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MarketingBetaActIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MarketingBetaActIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MarketingBetaActIn) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MarketingBetaActIn) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetGithub

`func (o *MarketingBetaActIn) GetGithub() string`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *MarketingBetaActIn) GetGithubOk() (*string, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *MarketingBetaActIn) SetGithub(v string)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *MarketingBetaActIn) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetTarget

`func (o *MarketingBetaActIn) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MarketingBetaActIn) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MarketingBetaActIn) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MarketingBetaActIn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


