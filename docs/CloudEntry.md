# CloudEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archetype** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Forkable** | Pointer to **bool** | Forkable is NOT omitempty: false is an answer here, not a missing field. Omitted, a client could not tell \&quot;you cannot fork this\&quot; from \&quot;nobody said\&quot;. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | repo | site | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** | Note is why a row is NOT in the published catalog, set by the admission gate (gate.go) on the sites it holds back. It is the difference between a demo that silently vanished from the public lens and one whose owner can read the reason and fix it. A published row never carries one. | [optional] 
**Org** | Pointer to **string** | hanzo | lux | zoo | [optional] 
**Origin** | Pointer to **string** | Origin is WHAT THIS IS TO YOU: template | community | third-party | product (origin.go owns the four nouns and derives them). Not omitempty, for the same reason Forkable is not: every row has an answer, and a missing one is exactly the flattening this field exists to end. | [optional] 
**Repo** | Pointer to **string** | source | [optional] 
**Scope** | Pointer to **string** | Scope is provenance, not storage: \&quot;public\&quot; for a row from the published corpus, \&quot;org\&quot; for one only this caller can see. A UI that cannot tell them apart cannot warn before sharing a link. | [optional] 
**Stars** | Pointer to **int32** |  | [optional] 
**Template** | Pointer to **string** | lineage, if forked from one | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Upstream** | Pointer to **string** | Upstream/License credit the third-party work an entry was published from: the difference between \&quot;this org built it\&quot; and \&quot;somebody else built it and we are showing it to you\&quot;.  WHO built it is Org, above — the account that paid for the project. There was once a separate admin-gated &#x60;official&#x60; boolean here claiming the same thing, and because it was gated it disagreed: apps Hanzo wrote and hosts were published by a script holding an ordinary org token, so it stayed false on all of them and this directory filed our own work as somebody else&#39;s. A field that restates an unforgeable fact can only ever be the wrong copy of it. | [optional] 
**Url** | Pointer to **string** | live, if it is deployed | [optional] 

## Methods

### NewCloudEntry

`func NewCloudEntry() *CloudEntry`

NewCloudEntry instantiates a new CloudEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEntryWithDefaults

`func NewCloudEntryWithDefaults() *CloudEntry`

NewCloudEntryWithDefaults instantiates a new CloudEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchetype

`func (o *CloudEntry) GetArchetype() string`

GetArchetype returns the Archetype field if non-nil, zero value otherwise.

### GetArchetypeOk

`func (o *CloudEntry) GetArchetypeOk() (*string, bool)`

GetArchetypeOk returns a tuple with the Archetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchetype

`func (o *CloudEntry) SetArchetype(v string)`

SetArchetype sets Archetype field to given value.

### HasArchetype

`func (o *CloudEntry) HasArchetype() bool`

HasArchetype returns a boolean if a field has been set.

### GetDescription

`func (o *CloudEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForkable

`func (o *CloudEntry) GetForkable() bool`

GetForkable returns the Forkable field if non-nil, zero value otherwise.

### GetForkableOk

`func (o *CloudEntry) GetForkableOk() (*bool, bool)`

GetForkableOk returns a tuple with the Forkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkable

`func (o *CloudEntry) SetForkable(v bool)`

SetForkable sets Forkable field to given value.

### HasForkable

`func (o *CloudEntry) HasForkable() bool`

HasForkable returns a boolean if a field has been set.

### GetId

`func (o *CloudEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguage

`func (o *CloudEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CloudEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CloudEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CloudEntry) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLicense

`func (o *CloudEntry) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CloudEntry) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CloudEntry) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CloudEntry) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *CloudEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *CloudEntry) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CloudEntry) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CloudEntry) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CloudEntry) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrg

`func (o *CloudEntry) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudEntry) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudEntry) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudEntry) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOrigin

`func (o *CloudEntry) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CloudEntry) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CloudEntry) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CloudEntry) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetRepo

`func (o *CloudEntry) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudEntry) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudEntry) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudEntry) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetScope

`func (o *CloudEntry) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudEntry) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudEntry) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudEntry) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStars

`func (o *CloudEntry) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *CloudEntry) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *CloudEntry) SetStars(v int32)`

SetStars sets Stars field to given value.

### HasStars

`func (o *CloudEntry) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetTemplate

`func (o *CloudEntry) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CloudEntry) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CloudEntry) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CloudEntry) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *CloudEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudEntry) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudEntry) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudEntry) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudEntry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpstream

`func (o *CloudEntry) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *CloudEntry) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *CloudEntry) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *CloudEntry) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetUrl

`func (o *CloudEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudEntry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


