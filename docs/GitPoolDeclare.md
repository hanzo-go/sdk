# GitPoolDeclare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels are what a workflow&#39;s &#x60;runs-on:&#x60; selects this pool by. A job is assigned to a pool that carries EVERY label it asked for. | [optional] 
**Name** | Pointer to **string** | Name is the pool&#39;s handle in this org, e.g. \&quot;evo\&quot; or \&quot;spark\&quot;. | [optional] 

## Methods

### NewGitPoolDeclare

`func NewGitPoolDeclare() *GitPoolDeclare`

NewGitPoolDeclare instantiates a new GitPoolDeclare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPoolDeclareWithDefaults

`func NewGitPoolDeclareWithDefaults() *GitPoolDeclare`

NewGitPoolDeclareWithDefaults instantiates a new GitPoolDeclare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *GitPoolDeclare) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GitPoolDeclare) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GitPoolDeclare) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GitPoolDeclare) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *GitPoolDeclare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitPoolDeclare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitPoolDeclare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitPoolDeclare) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


