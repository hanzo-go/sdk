# Strategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Blog** | Pointer to [**Blog**](Blog.md) | long-form explainer (nil for un-blogged tactics) | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Era** | Pointer to **string** | modern | heritage | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Principle** | Pointer to **string** | the spine slug this tactic files under | [optional] 
**Source** | Pointer to **string** | provenance / attribution | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Workload** | Pointer to **string** |  | [optional] 

## Methods

### NewStrategy

`func NewStrategy() *Strategy`

NewStrategy instantiates a new Strategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyWithDefaults

`func NewStrategyWithDefaults() *Strategy`

NewStrategyWithDefaults instantiates a new Strategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Strategy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Strategy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Strategy) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Strategy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlog

`func (o *Strategy) GetBlog() Blog`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *Strategy) GetBlogOk() (*Blog, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *Strategy) SetBlog(v Blog)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *Strategy) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetCategory

`func (o *Strategy) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Strategy) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Strategy) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Strategy) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *Strategy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Strategy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Strategy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Strategy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEra

`func (o *Strategy) GetEra() string`

GetEra returns the Era field if non-nil, zero value otherwise.

### GetEraOk

`func (o *Strategy) GetEraOk() (*string, bool)`

GetEraOk returns a tuple with the Era field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEra

`func (o *Strategy) SetEra(v string)`

SetEra sets Era field to given value.

### HasEra

`func (o *Strategy) HasEra() bool`

HasEra returns a boolean if a field has been set.

### GetId

`func (o *Strategy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Strategy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Strategy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Strategy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrinciple

`func (o *Strategy) GetPrinciple() string`

GetPrinciple returns the Principle field if non-nil, zero value otherwise.

### GetPrincipleOk

`func (o *Strategy) GetPrincipleOk() (*string, bool)`

GetPrincipleOk returns a tuple with the Principle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciple

`func (o *Strategy) SetPrinciple(v string)`

SetPrinciple sets Principle field to given value.

### HasPrinciple

`func (o *Strategy) HasPrinciple() bool`

HasPrinciple returns a boolean if a field has been set.

### GetSource

`func (o *Strategy) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Strategy) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Strategy) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Strategy) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *Strategy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Strategy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Strategy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Strategy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWorkload

`func (o *Strategy) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *Strategy) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *Strategy) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *Strategy) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


