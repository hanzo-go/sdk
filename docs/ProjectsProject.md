# ProjectsProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analytics** | Pointer to **bool** | Analytics is whether the web-analytics beacon is injected into this site&#39;s pages. It is ON by default — a project has to opt out — and it is what the static builder reads to decide whether to inject at all. | [optional] 
**Bucket** | Pointer to **string** | Bucket is the object-store bucket the site&#39;s files are served out of. | [optional] 
**CacheControl** | Pointer to **string** | CacheControl is the Cache-Control policy the edge serves this site&#39;s HTML under — how long a reader may hold a stale page before asking again. Assets are content-addressed and are not governed by it. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the project was created, as Unix seconds. | [optional] 
**CurrentDeploymentId** | Pointer to **string** | CurrentDeploymentID names the deployment currently serving, so a caller can ask what is live without scanning the history. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary, which is copied onto forks of this project and shown on a gallery card. | [optional] 
**ForkedFrom** | Pointer to **string** | ForkedFrom is the parent this project was forked from (\&quot;&lt;org&gt;/&lt;slug&gt;\&quot; of a published project, or a catalog template slug) — the attribution edge a gallery credits. | [optional] 
**Framework** | Pointer to **string** | Framework is a BUILD HINT from a closed set, defaulting to static. It tells CI how to build a linked repo and never gates a deploy, so a wrong value costs a build rather than access. | [optional] 
**Hidden** | Pointer to **bool** | Hidden is PLATFORM MODERATION, and it is a different axis from visibility: it pulls a public project out of the catalogue without editing the publisher&#39;s own choice, so un-hiding restores exactly what they asked for. A project is listed only when it is public AND not hidden. Always present, never omitted, for the same reason as visibility. | [optional] 
**HiddenReason** | Pointer to **string** | HiddenReason is why moderation hid it. Absent when it is not hidden. | [optional] 
**Id** | Pointer to **string** | ID is the project&#39;s internal identifier. It is stable across a rename, but it is not what the API addresses this project by — &#x60;slug&#x60; is. | [optional] 
**Key** | Pointer to **string** | Key is the project&#39;s publishable ingest key, minted at create. It is the value the injected beacon carries and the ONE thing that attributes this site&#39;s events; the static-builder reads it beside analytics.  Publishable means it belongs in a page&#39;s source: it names a write scope and mints no principal, so it is returned in full rather than masked. Masking it would only mean every caller needed a second endpoint to get the thing the page already ships. | [optional] 
**LastPurgeAt** | Pointer to **int32** | LastPurgeAt is when the edge cache was last cleared, as Unix seconds, so a console can say how fresh what readers see actually is. Absent means never. | [optional] 
**License** | Pointer to **string** | License is the terms that upstream work carries. Absent has the same reading: undeclared, not unencumbered. | [optional] 
**LiveUrl** | Pointer to **string** | LiveURL is where the site answers today. Absent until something has been deployed. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s display name, free text a person chose. | [optional] 
**Org** | Pointer to **string** | Org is the organisation that owns the project, and therefore who pays for it and who may change it. It is also the AUTHORSHIP line a gallery credits; there is no separate author field. | [optional] 
**Repo** | Pointer to [**ProjectsRepo**](ProjectsRepo.md) | Repo is the git source this project builds from, empty when it is deployed by uploading an artifact instead. | [optional] 
**Slug** | Pointer to **string** | Slug is the identifier that MATTERS: the handle every later call addresses, the S3 key segment the site&#39;s objects live under, and the label of the public host &#x60;&lt;slug&gt;.hanzo.app&#x60;. Because it is a hostname it is constrained and reserved labels such as &#x60;api&#x60; are refused. | [optional] 
**Space** | Pointer to **string** | Space is the project&#39;s Base data space, which is where a deployed site&#39;s form, forum and data submissions land. Absent means the site stores nothing. | [optional] 
**Starred** | Pointer to **bool** | Starred is THIS CALLER&#39;s star, not a property of the project — two people in the same org see different values for the same row, which is the whole point of it. Always present so a client can tell \&quot;not starred\&quot; from \&quot;this API is too old to say\&quot;, the same reason visibility and hidden are. | [optional] 
**Status** | Pointer to **string** | Status is where the project stands — whether a build has ever succeeded and whether anything is serving right now. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags is the site&#39;s browser tag config: platform slug → non-secret pixel id (GA measurement, Meta pixel, …) — what track.js injects and the server CAPI reads, per site. Omitted when none are set. The API SECRET is never here (KMS). | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the project&#39;s own record last changed, as Unix seconds. A deploy is not an edit of the project, so this does not move on every publish. | [optional] 
**Upstream** | Pointer to **string** | Upstream credits the third-party work this project was published from — a free-text line, because the honest answer is a name and a title that no enum could hold. Absent means NOBODY HAS SAID, not that there is nothing to say. | [optional] 
**Visibility** | Pointer to **string** | Visibility is \&quot;public\&quot; or \&quot;private\&quot;, and Hidden reports platform moderation. Both are always present (never omitempty) so a consumer can tell a real answer from \&quot;this API is too old to say\&quot; — and so a console never renders a project as public because a field was missing.  Authorship is deliberately absent: it is Org, above. | [optional] 

## Methods

### NewProjectsProject

`func NewProjectsProject() *ProjectsProject`

NewProjectsProject instantiates a new ProjectsProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsProjectWithDefaults

`func NewProjectsProjectWithDefaults() *ProjectsProject`

NewProjectsProjectWithDefaults instantiates a new ProjectsProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalytics

`func (o *ProjectsProject) GetAnalytics() bool`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *ProjectsProject) GetAnalyticsOk() (*bool, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *ProjectsProject) SetAnalytics(v bool)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *ProjectsProject) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetBucket

`func (o *ProjectsProject) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ProjectsProject) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ProjectsProject) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ProjectsProject) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCacheControl

`func (o *ProjectsProject) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *ProjectsProject) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *ProjectsProject) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *ProjectsProject) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectsProject) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsProject) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsProject) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectsProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *ProjectsProject) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *ProjectsProject) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *ProjectsProject) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *ProjectsProject) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForkedFrom

`func (o *ProjectsProject) GetForkedFrom() string`

GetForkedFrom returns the ForkedFrom field if non-nil, zero value otherwise.

### GetForkedFromOk

`func (o *ProjectsProject) GetForkedFromOk() (*string, bool)`

GetForkedFromOk returns a tuple with the ForkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkedFrom

`func (o *ProjectsProject) SetForkedFrom(v string)`

SetForkedFrom sets ForkedFrom field to given value.

### HasForkedFrom

`func (o *ProjectsProject) HasForkedFrom() bool`

HasForkedFrom returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectsProject) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectsProject) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectsProject) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectsProject) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetHidden

`func (o *ProjectsProject) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ProjectsProject) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ProjectsProject) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ProjectsProject) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHiddenReason

`func (o *ProjectsProject) GetHiddenReason() string`

GetHiddenReason returns the HiddenReason field if non-nil, zero value otherwise.

### GetHiddenReasonOk

`func (o *ProjectsProject) GetHiddenReasonOk() (*string, bool)`

GetHiddenReasonOk returns a tuple with the HiddenReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenReason

`func (o *ProjectsProject) SetHiddenReason(v string)`

SetHiddenReason sets HiddenReason field to given value.

### HasHiddenReason

`func (o *ProjectsProject) HasHiddenReason() bool`

HasHiddenReason returns a boolean if a field has been set.

### GetId

`func (o *ProjectsProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectsProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectsProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectsProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ProjectsProject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectsProject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectsProject) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectsProject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastPurgeAt

`func (o *ProjectsProject) GetLastPurgeAt() int32`

GetLastPurgeAt returns the LastPurgeAt field if non-nil, zero value otherwise.

### GetLastPurgeAtOk

`func (o *ProjectsProject) GetLastPurgeAtOk() (*int32, bool)`

GetLastPurgeAtOk returns a tuple with the LastPurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPurgeAt

`func (o *ProjectsProject) SetLastPurgeAt(v int32)`

SetLastPurgeAt sets LastPurgeAt field to given value.

### HasLastPurgeAt

`func (o *ProjectsProject) HasLastPurgeAt() bool`

HasLastPurgeAt returns a boolean if a field has been set.

### GetLicense

`func (o *ProjectsProject) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ProjectsProject) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ProjectsProject) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ProjectsProject) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectsProject) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectsProject) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectsProject) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectsProject) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetName

`func (o *ProjectsProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectsProject) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectsProject) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectsProject) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectsProject) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectsProject) GetRepo() ProjectsRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectsProject) GetRepoOk() (*ProjectsRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectsProject) SetRepo(v ProjectsRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectsProject) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsProject) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsProject) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpace

`func (o *ProjectsProject) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *ProjectsProject) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *ProjectsProject) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *ProjectsProject) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetStarred

`func (o *ProjectsProject) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *ProjectsProject) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *ProjectsProject) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *ProjectsProject) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectsProject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsProject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsProject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectsProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ProjectsProject) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectsProject) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectsProject) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectsProject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectsProject) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectsProject) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectsProject) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectsProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstream

`func (o *ProjectsProject) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ProjectsProject) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ProjectsProject) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ProjectsProject) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetVisibility

`func (o *ProjectsProject) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ProjectsProject) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ProjectsProject) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ProjectsProject) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


