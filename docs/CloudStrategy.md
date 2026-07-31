# CloudStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Blog** | Pointer to [**CloudBlog**](CloudBlog.md) | long-form explainer (nil for un-blogged tactics) | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Era** | Pointer to **string** | modern | heritage | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Principle** | Pointer to **string** | the spine slug this tactic files under | [optional] 
**Source** | Pointer to **string** | provenance / attribution | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Workload** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudStrategy

`func NewCloudStrategy() *CloudStrategy`

NewCloudStrategy instantiates a new CloudStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStrategyWithDefaults

`func NewCloudStrategyWithDefaults() *CloudStrategy`

NewCloudStrategyWithDefaults instantiates a new CloudStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CloudStrategy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CloudStrategy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CloudStrategy) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CloudStrategy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlog

`func (o *CloudStrategy) GetBlog() CloudBlog`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *CloudStrategy) GetBlogOk() (*CloudBlog, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *CloudStrategy) SetBlog(v CloudBlog)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *CloudStrategy) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetCategory

`func (o *CloudStrategy) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudStrategy) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudStrategy) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudStrategy) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudStrategy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudStrategy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudStrategy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudStrategy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEra

`func (o *CloudStrategy) GetEra() string`

GetEra returns the Era field if non-nil, zero value otherwise.

### GetEraOk

`func (o *CloudStrategy) GetEraOk() (*string, bool)`

GetEraOk returns a tuple with the Era field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEra

`func (o *CloudStrategy) SetEra(v string)`

SetEra sets Era field to given value.

### HasEra

`func (o *CloudStrategy) HasEra() bool`

HasEra returns a boolean if a field has been set.

### GetId

`func (o *CloudStrategy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudStrategy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudStrategy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudStrategy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrinciple

`func (o *CloudStrategy) GetPrinciple() string`

GetPrinciple returns the Principle field if non-nil, zero value otherwise.

### GetPrincipleOk

`func (o *CloudStrategy) GetPrincipleOk() (*string, bool)`

GetPrincipleOk returns a tuple with the Principle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciple

`func (o *CloudStrategy) SetPrinciple(v string)`

SetPrinciple sets Principle field to given value.

### HasPrinciple

`func (o *CloudStrategy) HasPrinciple() bool`

HasPrinciple returns a boolean if a field has been set.

### GetSource

`func (o *CloudStrategy) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudStrategy) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudStrategy) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudStrategy) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *CloudStrategy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudStrategy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudStrategy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudStrategy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWorkload

`func (o *CloudStrategy) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *CloudStrategy) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *CloudStrategy) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *CloudStrategy) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


