# ProviderGithubBranchesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]ProviderGithubBranch**](ProviderGithubBranch.md) | Branches is this page: the default branch first, then the rest by name. Never null. | [optional] 
**Next** | Pointer to **string** | Next is the cursor for the page after this one; empty on the last page. | [optional] 
**Total** | Pointer to **int64** | Total is how many branches match q, across every page. | [optional] 

## Methods

### NewProviderGithubBranchesOut

`func NewProviderGithubBranchesOut() *ProviderGithubBranchesOut`

NewProviderGithubBranchesOut instantiates a new ProviderGithubBranchesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubBranchesOutWithDefaults

`func NewProviderGithubBranchesOutWithDefaults() *ProviderGithubBranchesOut`

NewProviderGithubBranchesOutWithDefaults instantiates a new ProviderGithubBranchesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *ProviderGithubBranchesOut) GetBranches() []ProviderGithubBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *ProviderGithubBranchesOut) GetBranchesOk() (*[]ProviderGithubBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *ProviderGithubBranchesOut) SetBranches(v []ProviderGithubBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *ProviderGithubBranchesOut) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetNext

`func (o *ProviderGithubBranchesOut) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ProviderGithubBranchesOut) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ProviderGithubBranchesOut) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ProviderGithubBranchesOut) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotal

`func (o *ProviderGithubBranchesOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderGithubBranchesOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderGithubBranchesOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProviderGithubBranchesOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


