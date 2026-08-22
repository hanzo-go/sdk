# StrategyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the tactic itself: the thing to go and do. | [optional] 
**Category** | Pointer to **string** | Category is the growth discipline the tactic belongs to — the axis &#x60;?category&#x3D;&#x60; narrows on. | [optional] 
**Id** | Pointer to **string** | ID is the tactic&#39;s stable slug in the corpus. | [optional] 
**Tags** | Pointer to **[]string** | Tags are the PRECONDITIONS this tactic already satisfied to appear in the answer — &#x60;stage:&lt;name&gt;&#x60; and &#x60;has:&lt;capability&gt;&#x60; predicates over the org&#39;s observed profile. They are carried back so a caller can show why a tactic surfaced, not so it can filter again. | [optional] 
**Workload** | Pointer to **string** | Workload is how much effort running the tactic costs, so a corpus can be cut to what the org has the hands for. | [optional] 

## Methods

### NewStrategyView

`func NewStrategyView() *StrategyView`

NewStrategyView instantiates a new StrategyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyViewWithDefaults

`func NewStrategyViewWithDefaults() *StrategyView`

NewStrategyViewWithDefaults instantiates a new StrategyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *StrategyView) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StrategyView) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StrategyView) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *StrategyView) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCategory

`func (o *StrategyView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StrategyView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StrategyView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StrategyView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *StrategyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrategyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrategyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StrategyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTags

`func (o *StrategyView) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StrategyView) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StrategyView) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StrategyView) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWorkload

`func (o *StrategyView) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *StrategyView) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *StrategyView) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *StrategyView) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


