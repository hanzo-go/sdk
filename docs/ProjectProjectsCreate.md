# ProjectProjectsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analytics** | Pointer to **bool** | Analytics is the opt-OUT for the wired-by-default analytics beacon: absent (nil) ⇒ ON (the default); explicit false ⇒ off. A pointer so \&quot;unset\&quot; is distinguishable from \&quot;false\&quot; — the only way to turn the default off. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary, copied onto anything forked from this project. | [optional] 
**Framework** | Pointer to **string** | Framework is a BUILD HINT from a closed set, defaulting to static. It tells CI how to build a linked repo and never gates a deploy. | [optional] 
**License** | Pointer to **string** | License is the terms that upstream work carries. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s display name and the only REQUIRED field. When slug is omitted it is also what the slug is derived from. | [optional] 
**Repo** | Pointer to [**ProjectProjectsCreateRepo**](ProjectProjectsCreateRepo.md) |  | [optional] 
**Slug** | Pointer to **string** | Slug is the handle everything else addresses this project by: the public host &#x60;&lt;slug&gt;.hanzo.app&#x60;, the object-store key segment, and the path parameter of every later call. Derived from the name when omitted. It is a hostname label, so it is constrained and reserved labels such as &#x60;api&#x60; or &#x60;admin&#x60; are refused. | [optional] 
**Upstream** | Pointer to **string** | Upstream credits the third-party work this project was published from. It is accepted from any caller: giving away credit can only cost the publisher, so it needs no gate. | [optional] 
**Visibility** | Pointer to **string** | Visibility is \&quot;public\&quot; (the default when absent) or \&quot;private\&quot;. Publishing publicly is ungated — that is the point of a community. Going PRIVATE is the paid feature, so an unfunded org asking for it is refused rather than silently downgraded (see resolve). | [optional] 

## Methods

### NewProjectProjectsCreate

`func NewProjectProjectsCreate() *ProjectProjectsCreate`

NewProjectProjectsCreate instantiates a new ProjectProjectsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsCreateWithDefaults

`func NewProjectProjectsCreateWithDefaults() *ProjectProjectsCreate`

NewProjectProjectsCreateWithDefaults instantiates a new ProjectProjectsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalytics

`func (o *ProjectProjectsCreate) GetAnalytics() bool`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *ProjectProjectsCreate) GetAnalyticsOk() (*bool, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *ProjectProjectsCreate) SetAnalytics(v bool)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *ProjectProjectsCreate) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectProjectsCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectProjectsCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectProjectsCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectProjectsCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectProjectsCreate) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectProjectsCreate) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectProjectsCreate) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectProjectsCreate) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetLicense

`func (o *ProjectProjectsCreate) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ProjectProjectsCreate) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ProjectProjectsCreate) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ProjectProjectsCreate) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectsCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectsCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectsCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectsCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectProjectsCreate) GetRepo() ProjectProjectsCreateRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectProjectsCreate) GetRepoOk() (*ProjectProjectsCreateRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectProjectsCreate) SetRepo(v ProjectProjectsCreateRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectProjectsCreate) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsCreate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUpstream

`func (o *ProjectProjectsCreate) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ProjectProjectsCreate) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ProjectProjectsCreate) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ProjectProjectsCreate) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetVisibility

`func (o *ProjectProjectsCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ProjectProjectsCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ProjectProjectsCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ProjectProjectsCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


