# ProjectProjectsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheControl** | Pointer to **string** | CacheControl replaces the Cache-Control policy the edge serves this site&#39;s HTML under. Absent leaves it. | [optional] 
**Description** | Pointer to **string** | Description replaces the one-line summary. Absent leaves it. | [optional] 
**DropLocal** | Pointer to **bool** | DropLocal turns the drop of localhost events at ingest on or off. Absent leaves it. | [optional] 
**Framework** | Pointer to **string** | Framework replaces the build hint. It affects the NEXT build only — nothing already deployed is rebuilt. | [optional] 
**Hidden** | Pointer to **bool** | Hidden is MODERATION, and the only admin-gated field on this body: it pulls a public project out of the catalogue from admin.hanzo.ai without editing the publisher&#39;s own visibility choice, so un-hiding restores exactly what they asked for. A tenant sending it is ignored. | [optional] 
**HiddenReason** | Pointer to **string** | HiddenReason records WHY moderation hid it, so the action can be explained and reviewed later. Admin-gated like hidden itself. | [optional] 
**InternalNetworks** | Pointer to **[]string** | InternalNetworks replaces the site&#39;s own addresses: CIDRs or bare addresses, at most 64. Absent leaves them; send [] to clear. | [optional] 
**License** | Pointer to **string** | License is the terms that upstream work carries, with the same clear-versus- leave rule. | [optional] 
**Name** | Pointer to **string** | Name replaces the display name. Absent leaves it; the slug never moves with it. | [optional] 
**Repo** | Pointer to [**ProjectProjectsUpdateRepo**](ProjectProjectsUpdateRepo.md) |  | [optional] 
**Slug** | Pointer to **string** | Slug is the project to update, from the path. The URL is the addressing authority — a &#x60;slug&#x60; in the body cannot move the write to another project. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags sets the site&#39;s browser tag config: platform slug → non-secret pixel id (e.g. {\&quot;ga4\&quot;:\&quot;G-…\&quot;,\&quot;meta\&quot;:\&quot;…\&quot;}). track.js injects these first-party and the server CAPI reads them, per site. Absent LEAVES them; a present object REPLACES the set (send {} to clear). The ids are public — they ship in the page — so this is not the SECRET path (a CAPI token is sealed via POST /v1/destination). | [optional] 
**Upstream** | Pointer to **string** | Upstream credits the third-party work this project was published from, and is settable after the fact because the live demos are the ones that most need crediting. An explicit empty string CLEARS the credit; absent leaves it. | [optional] 
**Visibility** | Pointer to **string** | Visibility flips an existing project between \&quot;public\&quot; and \&quot;private\&quot;. Same ONE rule as at create: public is free, private needs a paid plan. | [optional] 

## Methods

### NewProjectProjectsUpdate

`func NewProjectProjectsUpdate() *ProjectProjectsUpdate`

NewProjectProjectsUpdate instantiates a new ProjectProjectsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsUpdateWithDefaults

`func NewProjectProjectsUpdateWithDefaults() *ProjectProjectsUpdate`

NewProjectProjectsUpdateWithDefaults instantiates a new ProjectProjectsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheControl

`func (o *ProjectProjectsUpdate) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *ProjectProjectsUpdate) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *ProjectProjectsUpdate) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *ProjectProjectsUpdate) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectProjectsUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectProjectsUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectProjectsUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectProjectsUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDropLocal

`func (o *ProjectProjectsUpdate) GetDropLocal() bool`

GetDropLocal returns the DropLocal field if non-nil, zero value otherwise.

### GetDropLocalOk

`func (o *ProjectProjectsUpdate) GetDropLocalOk() (*bool, bool)`

GetDropLocalOk returns a tuple with the DropLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropLocal

`func (o *ProjectProjectsUpdate) SetDropLocal(v bool)`

SetDropLocal sets DropLocal field to given value.

### HasDropLocal

`func (o *ProjectProjectsUpdate) HasDropLocal() bool`

HasDropLocal returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectProjectsUpdate) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectProjectsUpdate) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectProjectsUpdate) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectProjectsUpdate) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetHidden

`func (o *ProjectProjectsUpdate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ProjectProjectsUpdate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ProjectProjectsUpdate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ProjectProjectsUpdate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHiddenReason

`func (o *ProjectProjectsUpdate) GetHiddenReason() string`

GetHiddenReason returns the HiddenReason field if non-nil, zero value otherwise.

### GetHiddenReasonOk

`func (o *ProjectProjectsUpdate) GetHiddenReasonOk() (*string, bool)`

GetHiddenReasonOk returns a tuple with the HiddenReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenReason

`func (o *ProjectProjectsUpdate) SetHiddenReason(v string)`

SetHiddenReason sets HiddenReason field to given value.

### HasHiddenReason

`func (o *ProjectProjectsUpdate) HasHiddenReason() bool`

HasHiddenReason returns a boolean if a field has been set.

### GetInternalNetworks

`func (o *ProjectProjectsUpdate) GetInternalNetworks() []string`

GetInternalNetworks returns the InternalNetworks field if non-nil, zero value otherwise.

### GetInternalNetworksOk

`func (o *ProjectProjectsUpdate) GetInternalNetworksOk() (*[]string, bool)`

GetInternalNetworksOk returns a tuple with the InternalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNetworks

`func (o *ProjectProjectsUpdate) SetInternalNetworks(v []string)`

SetInternalNetworks sets InternalNetworks field to given value.

### HasInternalNetworks

`func (o *ProjectProjectsUpdate) HasInternalNetworks() bool`

HasInternalNetworks returns a boolean if a field has been set.

### GetLicense

`func (o *ProjectProjectsUpdate) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ProjectProjectsUpdate) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ProjectProjectsUpdate) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ProjectProjectsUpdate) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectsUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectsUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectProjectsUpdate) GetRepo() ProjectProjectsUpdateRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectProjectsUpdate) GetRepoOk() (*ProjectProjectsUpdateRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectProjectsUpdate) SetRepo(v ProjectProjectsUpdateRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectProjectsUpdate) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTags

`func (o *ProjectProjectsUpdate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectProjectsUpdate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectProjectsUpdate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectProjectsUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpstream

`func (o *ProjectProjectsUpdate) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ProjectProjectsUpdate) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ProjectProjectsUpdate) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ProjectProjectsUpdate) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetVisibility

`func (o *ProjectProjectsUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ProjectProjectsUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ProjectProjectsUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ProjectProjectsUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


