# ProjectsBuildSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brief** | **string** | Natural-language description of the site to generate (required; capped at 8 KiB). | 
**Slug** | Pointer to **string** | Target project slug; derived from name (or minted \&quot;site-&lt;token&gt;\&quot;) when omitted. | [optional] 
**Name** | Pointer to **string** | Display name; defaults to the generated site name, else \&quot;Site\&quot;. | [optional] 
**Model** | Pointer to **string** | Inference model override; empty selects the gateway default. | [optional] 

## Methods

### NewProjectsBuildSiteRequest

`func NewProjectsBuildSiteRequest(brief string, ) *ProjectsBuildSiteRequest`

NewProjectsBuildSiteRequest instantiates a new ProjectsBuildSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsBuildSiteRequestWithDefaults

`func NewProjectsBuildSiteRequestWithDefaults() *ProjectsBuildSiteRequest`

NewProjectsBuildSiteRequestWithDefaults instantiates a new ProjectsBuildSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrief

`func (o *ProjectsBuildSiteRequest) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *ProjectsBuildSiteRequest) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *ProjectsBuildSiteRequest) SetBrief(v string)`

SetBrief sets Brief field to given value.


### GetSlug

`func (o *ProjectsBuildSiteRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsBuildSiteRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsBuildSiteRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsBuildSiteRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *ProjectsBuildSiteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsBuildSiteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsBuildSiteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsBuildSiteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *ProjectsBuildSiteRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProjectsBuildSiteRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProjectsBuildSiteRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProjectsBuildSiteRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


