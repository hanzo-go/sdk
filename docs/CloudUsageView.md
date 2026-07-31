# CloudUsageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org the rollup is for. | [optional] 
**Repos** | Pointer to [**[]CloudUsageRepo**](CloudUsageRepo.md) | Repos is every repo the org owns, across every project sub-scope. | [optional] 
**TotalBytes** | Pointer to **int32** | TotalBytes is the sum over Repos — the org&#39;s whole git footprint. | [optional] 

## Methods

### NewCloudUsageView

`func NewCloudUsageView() *CloudUsageView`

NewCloudUsageView instantiates a new CloudUsageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageViewWithDefaults

`func NewCloudUsageViewWithDefaults() *CloudUsageView`

NewCloudUsageViewWithDefaults instantiates a new CloudUsageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *CloudUsageView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudUsageView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudUsageView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudUsageView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepos

`func (o *CloudUsageView) GetRepos() []CloudUsageRepo`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CloudUsageView) GetReposOk() (*[]CloudUsageRepo, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CloudUsageView) SetRepos(v []CloudUsageRepo)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CloudUsageView) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetTotalBytes

`func (o *CloudUsageView) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *CloudUsageView) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *CloudUsageView) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *CloudUsageView) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


