# UsageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org the rollup is for. | [optional] 
**Repos** | Pointer to [**[]UsageRepo**](UsageRepo.md) | Repos is every repo the org owns, across every project sub-scope. | [optional] 
**TotalBytes** | Pointer to **int32** | TotalBytes is the sum over Repos — the org&#39;s whole git footprint. | [optional] 

## Methods

### NewUsageView

`func NewUsageView() *UsageView`

NewUsageView instantiates a new UsageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageViewWithDefaults

`func NewUsageViewWithDefaults() *UsageView`

NewUsageViewWithDefaults instantiates a new UsageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *UsageView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *UsageView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *UsageView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *UsageView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepos

`func (o *UsageView) GetRepos() []UsageRepo`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *UsageView) GetReposOk() (*[]UsageRepo, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *UsageView) SetRepos(v []UsageRepo)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *UsageView) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetTotalBytes

`func (o *UsageView) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *UsageView) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *UsageView) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *UsageView) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


