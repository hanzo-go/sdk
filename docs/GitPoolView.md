# GitPoolView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the pool was declared, in unix seconds. | [optional] 
**Labels** | Pointer to **[]string** | Labels are what a workflow&#39;s &#x60;runs-on:&#x60; selects this pool by. | [optional] 
**Name** | Pointer to **string** | Name is the pool&#39;s handle in this org. | [optional] 
**Runners** | Pointer to **int64** | Runners is how many daemons have entered the pool. Zero means the capacity is declared and nothing has turned up to provide it. | [optional] 

## Methods

### NewGitPoolView

`func NewGitPoolView() *GitPoolView`

NewGitPoolView instantiates a new GitPoolView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPoolViewWithDefaults

`func NewGitPoolViewWithDefaults() *GitPoolView`

NewGitPoolViewWithDefaults instantiates a new GitPoolView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GitPoolView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitPoolView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitPoolView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitPoolView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *GitPoolView) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GitPoolView) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GitPoolView) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GitPoolView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *GitPoolView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitPoolView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitPoolView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitPoolView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunners

`func (o *GitPoolView) GetRunners() int64`

GetRunners returns the Runners field if non-nil, zero value otherwise.

### GetRunnersOk

`func (o *GitPoolView) GetRunnersOk() (*int64, bool)`

GetRunnersOk returns a tuple with the Runners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunners

`func (o *GitPoolView) SetRunners(v int64)`

SetRunners sets Runners field to given value.

### HasRunners

`func (o *GitPoolView) HasRunners() bool`

HasRunners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


