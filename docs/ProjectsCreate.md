# ProjectsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analytics** | Pointer to **bool** | Analytics is the opt-OUT for the wired-by-default analytics beacon: absent (nil) ⇒ ON (the default); explicit false ⇒ off. A pointer so \&quot;unset\&quot; is distinguishable from \&quot;false\&quot; — the only way to turn the default off. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to [**ProjectsCreateRepo**](ProjectsCreateRepo.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Upstream** | Pointer to **string** | Upstream/License credit the third-party work this project was published from. Taken from any caller: disclaiming authorship can only cost the publisher credit, so it needs no gate (see Project.Upstream). | [optional] 
**Visibility** | Pointer to **string** | Visibility is \&quot;public\&quot; (the default when absent) or \&quot;private\&quot;. Publishing publicly is ungated — that is the point of a community. Going PRIVATE is the paid feature, so an unfunded org asking for it is refused rather than silently downgraded (see resolve). | [optional] 

## Methods

### NewProjectsCreate

`func NewProjectsCreate() *ProjectsCreate`

NewProjectsCreate instantiates a new ProjectsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCreateWithDefaults

`func NewProjectsCreateWithDefaults() *ProjectsCreate`

NewProjectsCreateWithDefaults instantiates a new ProjectsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalytics

`func (o *ProjectsCreate) GetAnalytics() bool`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *ProjectsCreate) GetAnalyticsOk() (*bool, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *ProjectsCreate) SetAnalytics(v bool)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *ProjectsCreate) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectsCreate) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectsCreate) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectsCreate) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectsCreate) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetLicense

`func (o *ProjectsCreate) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ProjectsCreate) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ProjectsCreate) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ProjectsCreate) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *ProjectsCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectsCreate) GetRepo() ProjectsCreateRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectsCreate) GetRepoOk() (*ProjectsCreateRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectsCreate) SetRepo(v ProjectsCreateRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectsCreate) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsCreate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUpstream

`func (o *ProjectsCreate) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ProjectsCreate) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ProjectsCreate) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ProjectsCreate) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetVisibility

`func (o *ProjectsCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ProjectsCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ProjectsCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ProjectsCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


