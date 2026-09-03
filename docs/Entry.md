# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archetype** | Pointer to **string** | Archetype is WHAT KIND OF THING this is, from a closed and ordered list — model | contract | chain | sdk | template | infra | site | app — derived from the repository&#39;s own topics, name and description, first match winning, and always &#x60;site&#x60; for a deployed site. It is DERIVED, never guessed by a model, because a wrong archetype hides a row from the browse rail more thoroughly than a missing one does. Empty when no topic matched: unclassified, not uncategorisable. | [optional] 
**Description** | Pointer to **string** | Description is the repository&#39;s own one-line GitHub description, carried verbatim. It comes from the SOURCE half of a row, so a site that was never matched to a repository has none, and nothing here is written by us. | [optional] 
**Forkable** | Pointer to **bool** | Forkable is NOT omitempty: false is an answer here, not a missing field. Omitted, a client could not tell \&quot;you cannot fork this\&quot; from \&quot;nobody said\&quot;. | [optional] 
**Id** | Pointer to **string** | ID is \&quot;&lt;org&gt;/&lt;name&gt;\&quot; and is the corpus&#39;s primary key: a re-published entry updates in place under it rather than accumulating duplicates, so it is the one handle stable enough to link to or to name in a &#x60;template&#x60; filter. Two orgs can spell the same id, and &#x60;canonical&#x60; picks which one keeps it. | [optional] 
**Kind** | Pointer to **string** | repo | site | [optional] 
**Language** | Pointer to **string** | Language is the repository&#39;s primary implementation language as GitHub computes it (\&quot;Go\&quot;, \&quot;TypeScript\&quot;), and the case is GitHub&#39;s. Empty for a site with no source half and for a repository GitHub could not classify. | [optional] 
**License** | Pointer to **string** | License is the terms that upstream work carries, in whichever form the half that credited it had: an SPDX id (\&quot;MIT\&quot;, \&quot;Apache-2.0\&quot;) on a GitHub fork, free text on a site whose publisher declared it. GitHub&#39;s NOASSERTION — \&quot;we could not identify it\&quot; — reads as none rather than as a licence by that name. So empty means UNDECLARED and never unencumbered, and Upstream is what says whether the question applies at all. | [optional] 
**Name** | Pointer to **string** | Name is the short identifier inside the org — the repository&#39;s name, or the site&#39;s slug — and is the half of ID after the slash. Not a display name; Title is. | [optional] 
**Note** | Pointer to **string** | Note is why a row is NOT in the published catalog, set by the admission gate (gate.go) on the sites it holds back. It is the difference between a demo that silently vanished from the public lens and one whose owner can read the reason and fix it. A published row never carries one. | [optional] 
**Org** | Pointer to **string** | hanzo | lux | zoo | [optional] 
**Origin** | Pointer to **string** | Origin is WHAT THIS IS TO YOU: template | community | third-party | product (origin.go owns the four nouns and derives them). Not omitempty, for the same reason Forkable is not: every row has an answer, and a missing one is exactly the flattening this field exists to end. | [optional] 
**Repo** | Pointer to **string** | source | [optional] 
**Scope** | Pointer to **string** | Scope is provenance, not storage: \&quot;public\&quot; for a row from the published corpus, \&quot;org\&quot; for one only this caller can see. A UI that cannot tell them apart cannot warn before sharing a link. | [optional] 
**Stars** | Pointer to **int64** | Stars is GitHub&#39;s stargazer count for the source repository, read at the last sync and never accumulated here. It is not a ranking — the page sorts on Updated — but it is the tiebreak when two orgs claim one ID. Absent for a site with no repository behind it, and for a repository nobody has starred. | [optional] 
**Template** | Pointer to **string** | lineage, if forked from one | [optional] 
**Title** | Pointer to **string** | Title is what to SHOW. A site&#39;s human name wins where it has one; a repo row falls back to the repository name, so on a repo this usually just repeats Name. Absent only for a site whose project was never named — render Name. | [optional] 
**Updated** | Pointer to **string** | Updated is when the thing last MOVED, as RFC 3339 in UTC: a repository&#39;s last push, or a site&#39;s last deploy. The page is ordered on it, most recent first, by comparing these strings — so the format is load-bearing and not cosmetic. Absent means the source reported no timestamp, and such a row sorts last. | [optional] 
**Upstream** | Pointer to **string** | Upstream/License credit the third-party work an entry was published from: the difference between \&quot;this org built it\&quot; and \&quot;somebody else built it and we are showing it to you\&quot;.  WHO built it is Org, above — the account that paid for the project. There was once a separate admin-gated &#x60;official&#x60; boolean here claiming the same thing, and because it was gated it disagreed: apps Hanzo wrote and hosts were published by a script holding an ordinary org token, so it stayed false on all of them and this directory filed our own work as somebody else&#39;s. A field that restates an unforgeable fact can only ever be the wrong copy of it. | [optional] 
**Url** | Pointer to **string** | live, if it is deployed | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchetype

`func (o *Entry) GetArchetype() string`

GetArchetype returns the Archetype field if non-nil, zero value otherwise.

### GetArchetypeOk

`func (o *Entry) GetArchetypeOk() (*string, bool)`

GetArchetypeOk returns a tuple with the Archetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchetype

`func (o *Entry) SetArchetype(v string)`

SetArchetype sets Archetype field to given value.

### HasArchetype

`func (o *Entry) HasArchetype() bool`

HasArchetype returns a boolean if a field has been set.

### GetDescription

`func (o *Entry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForkable

`func (o *Entry) GetForkable() bool`

GetForkable returns the Forkable field if non-nil, zero value otherwise.

### GetForkableOk

`func (o *Entry) GetForkableOk() (*bool, bool)`

GetForkableOk returns a tuple with the Forkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkable

`func (o *Entry) SetForkable(v bool)`

SetForkable sets Forkable field to given value.

### HasForkable

`func (o *Entry) HasForkable() bool`

HasForkable returns a boolean if a field has been set.

### GetId

`func (o *Entry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Entry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Entry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Entry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Entry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguage

`func (o *Entry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Entry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Entry) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Entry) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLicense

`func (o *Entry) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Entry) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Entry) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Entry) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *Entry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *Entry) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Entry) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Entry) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Entry) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrg

`func (o *Entry) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Entry) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Entry) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Entry) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOrigin

`func (o *Entry) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Entry) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Entry) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Entry) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetRepo

`func (o *Entry) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Entry) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Entry) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Entry) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetScope

`func (o *Entry) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Entry) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Entry) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Entry) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStars

`func (o *Entry) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *Entry) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *Entry) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *Entry) HasStars() bool`

HasStars returns a boolean if a field has been set.

### GetTemplate

`func (o *Entry) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Entry) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Entry) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Entry) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *Entry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Entry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Entry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Entry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *Entry) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Entry) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Entry) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Entry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpstream

`func (o *Entry) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *Entry) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *Entry) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *Entry) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetUrl

`func (o *Entry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Entry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Entry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Entry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


