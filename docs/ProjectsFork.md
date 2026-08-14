# ProjectsFork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | target project name (optional; defaults to the parent&#39;s title) | [optional] 
**Slug** | Pointer to **string** | parent slug to fork — catalog template or published project (required) | [optional] 
**Target** | Pointer to **string** | Target overrides the derived project slug (optional; defaults to the parent slug). Kept distinct from Slug so callers can rename on fork. | [optional] 
**Variant** | Pointer to **string** | Variant picks a template&#39;s format/page/theme (optional; defaults to the template&#39;s first shape). This is the axis the catalog used to spend sibling slugs on, so it is expressed here, where the user&#39;s preference is. | [optional] 

## Methods

### NewProjectsFork

`func NewProjectsFork() *ProjectsFork`

NewProjectsFork instantiates a new ProjectsFork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsForkWithDefaults

`func NewProjectsForkWithDefaults() *ProjectsFork`

NewProjectsForkWithDefaults instantiates a new ProjectsFork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsFork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsFork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsFork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsFork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsFork) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsFork) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsFork) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsFork) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTarget

`func (o *ProjectsFork) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ProjectsFork) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ProjectsFork) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ProjectsFork) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetVariant

`func (o *ProjectsFork) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ProjectsFork) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ProjectsFork) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ProjectsFork) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


