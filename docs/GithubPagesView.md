# GithubPagesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** | BuildType is the builder in use: \&quot;legacy\&quot; (branch source) or \&quot;workflow\&quot;. | [optional] 
**Cname** | Pointer to **string** | CNAME is the custom domain, absent when none is set. | [optional] 
**Custom404** | Pointer to **bool** | Custom404 is whether the repo ships its own 404 page. | [optional] 
**HttpsEnforced** | Pointer to **bool** | HTTPSEnforced is GitHub&#39;s enforce-HTTPS bit. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the site belongs to. | [optional] 
**Source** | Pointer to [**GithubPagesSource**](GithubPagesSource.md) | Source is the branch + path the site builds from. Absent under \&quot;workflow\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is GitHub&#39;s build state: \&quot;built\&quot;, \&quot;building\&quot; or \&quot;errored\&quot;. Absent before the first build. | [optional] 
**Url** | Pointer to **string** | URL is the live site (GitHub&#39;s html_url). | [optional] 

## Methods

### NewGithubPagesView

`func NewGithubPagesView() *GithubPagesView`

NewGithubPagesView instantiates a new GithubPagesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubPagesViewWithDefaults

`func NewGithubPagesViewWithDefaults() *GithubPagesView`

NewGithubPagesViewWithDefaults instantiates a new GithubPagesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *GithubPagesView) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *GithubPagesView) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *GithubPagesView) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *GithubPagesView) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCname

`func (o *GithubPagesView) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *GithubPagesView) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *GithubPagesView) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *GithubPagesView) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetCustom404

`func (o *GithubPagesView) GetCustom404() bool`

GetCustom404 returns the Custom404 field if non-nil, zero value otherwise.

### GetCustom404Ok

`func (o *GithubPagesView) GetCustom404Ok() (*bool, bool)`

GetCustom404Ok returns a tuple with the Custom404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom404

`func (o *GithubPagesView) SetCustom404(v bool)`

SetCustom404 sets Custom404 field to given value.

### HasCustom404

`func (o *GithubPagesView) HasCustom404() bool`

HasCustom404 returns a boolean if a field has been set.

### GetHttpsEnforced

`func (o *GithubPagesView) GetHttpsEnforced() bool`

GetHttpsEnforced returns the HttpsEnforced field if non-nil, zero value otherwise.

### GetHttpsEnforcedOk

`func (o *GithubPagesView) GetHttpsEnforcedOk() (*bool, bool)`

GetHttpsEnforcedOk returns a tuple with the HttpsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnforced

`func (o *GithubPagesView) SetHttpsEnforced(v bool)`

SetHttpsEnforced sets HttpsEnforced field to given value.

### HasHttpsEnforced

`func (o *GithubPagesView) HasHttpsEnforced() bool`

HasHttpsEnforced returns a boolean if a field has been set.

### GetRepo

`func (o *GithubPagesView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GithubPagesView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GithubPagesView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GithubPagesView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSource

`func (o *GithubPagesView) GetSource() GithubPagesSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GithubPagesView) GetSourceOk() (*GithubPagesSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GithubPagesView) SetSource(v GithubPagesSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GithubPagesView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *GithubPagesView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubPagesView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubPagesView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubPagesView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GithubPagesView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GithubPagesView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GithubPagesView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GithubPagesView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


