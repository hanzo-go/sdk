# Strategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the tactic itself: the thing to go and do, stated imperatively. | [optional] 
**Blog** | Pointer to [**Blog**](Blog.md) | Blog is the tactic&#39;s long-form explainer; absent for tactics that have none. | [optional] 
**Category** | Pointer to **string** | Category is the growth discipline the tactic belongs to — the axis &#x60;?category&#x3D;&#x60; narrows the corpus on, and one of the facets a caller browses by. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin lever. Absent reads as ON; an explicit false drops the tactic from every org-facing corpus read while leaving it in the document. | [optional] 
**Era** | Pointer to **string** | Era separates an AI-era tactic (&#x60;modern&#x60;) from a classical one (&#x60;heritage&#x60;). | [optional] 
**Id** | Pointer to **string** | ID is the tactic&#39;s stable slug, unique across the corpus. | [optional] 
**Principle** | Pointer to **string** | Principle is the spine slug this tactic files under (a Principle.Slug). | [optional] 
**Source** | Pointer to **string** | Source is where the tactic came from — the attribution a reader is owed. | [optional] 
**Tags** | Pointer to **[]string** | Tags are PRECONDITIONS, not labels — every one must be satisfied by the org&#39;s observed profile before the tactic surfaces, so an untagged tactic is universally applicable. Two vocabularies: &#x60;stage:&lt;research|formed|launched| activated|scaling&gt;&#x60; reads the org&#39;s growth stage, &#x60;has:&lt;capability&gt;&#x60; reads an observed signal. | [optional] 
**Workload** | Pointer to **string** | Workload is how much effort running the tactic costs, so a corpus can be narrowed to what an org has the hands for right now. | [optional] 

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


