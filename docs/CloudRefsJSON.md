# CloudRefsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]CloudRefJSON**](CloudRefJSON.md) | Branches are the repo&#39;s heads; empty on a repo with no commits. | [optional] 
**Default** | Pointer to **string** | Default is the branch name a caller gets when it asks for no ref. | [optional] 
**Tags** | Pointer to [**[]CloudRefJSON**](CloudRefJSON.md) | Tags are the repo&#39;s tags; empty when there are none. | [optional] 

## Methods

### NewCloudRefsJSON

`func NewCloudRefsJSON() *CloudRefsJSON`

NewCloudRefsJSON instantiates a new CloudRefsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRefsJSONWithDefaults

`func NewCloudRefsJSONWithDefaults() *CloudRefsJSON`

NewCloudRefsJSONWithDefaults instantiates a new CloudRefsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *CloudRefsJSON) GetBranches() []CloudRefJSON`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *CloudRefsJSON) GetBranchesOk() (*[]CloudRefJSON, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *CloudRefsJSON) SetBranches(v []CloudRefJSON)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *CloudRefsJSON) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetDefault

`func (o *CloudRefsJSON) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CloudRefsJSON) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CloudRefsJSON) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CloudRefsJSON) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTags

`func (o *CloudRefsJSON) GetTags() []CloudRefJSON`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudRefsJSON) GetTagsOk() (*[]CloudRefJSON, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudRefsJSON) SetTags(v []CloudRefJSON)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudRefsJSON) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


