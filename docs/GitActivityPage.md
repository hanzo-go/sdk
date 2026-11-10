# GitActivityPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GitActivityRepo**](GitActivityRepo.md) | Data is the page. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page size asked for. | [optional] 
**Page** | Pointer to **int64** | Page is this page&#39;s number, from 1. | [optional] 
**Total** | Pointer to **int64** | Total is how many public repositories match, across every page. | [optional] 

## Methods

### NewGitActivityPage

`func NewGitActivityPage() *GitActivityPage`

NewGitActivityPage instantiates a new GitActivityPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityPageWithDefaults

`func NewGitActivityPageWithDefaults() *GitActivityPage`

NewGitActivityPageWithDefaults instantiates a new GitActivityPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GitActivityPage) GetData() []GitActivityRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GitActivityPage) GetDataOk() (*[]GitActivityRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GitActivityPage) SetData(v []GitActivityRepo)`

SetData sets Data field to given value.

### HasData

`func (o *GitActivityPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *GitActivityPage) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GitActivityPage) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GitActivityPage) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GitActivityPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *GitActivityPage) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GitActivityPage) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GitActivityPage) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *GitActivityPage) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotal

`func (o *GitActivityPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GitActivityPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GitActivityPage) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GitActivityPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


