# Taxon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** | Brands are the brands whose console shows this taxon. Absent means every brand its category admits. | [optional] 
**Category** | Pointer to **string** | Category is the id of the category this taxon is filed under. | [optional] 
**Description** | Pointer to **string** | Description is the one line shown beneath the name in the catalogue and nav. | [optional] 
**Href** | Pointer to **string** | Href is the absolute URL an external product launches, for the taxa that genuinely live at their own domain. Empty for an in-console product. | [optional] 
**Icon** | Pointer to **string** | Icon names the icon the surface renders, e.g. \&quot;Database\&quot;. It is a NAME, not an image: which icon set draws it is the rendering surface&#39;s business. | [optional] 
**Id** | Pointer to **string** | ID is the stable slug this taxon is addressed by, e.g. \&quot;vector\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the display name, e.g. \&quot;Vector\&quot;. | [optional] 
**Order** | Pointer to **int32** | Order is where the taxon sits within its category, ascending. | [optional] 
**Owner** | Pointer to **string** | Owner is the org this product belongs to: the platform&#39;s own org for one every tenant sees, or your org for one you added. Where two rows share an id, yours is the one served. | [optional] 
**Published** | Pointer to **bool** | Published is whether the taxon is shown. An unpublished taxon is served only to an editor, so a product can be staged before anyone sees it. | [optional] 
**Route** | Pointer to **string** | Route is the in-console path this taxon opens, e.g. \&quot;/vector\&quot;. Set for a product the console renders itself; empty for an external one. | [optional] 
**Tags** | Pointer to **[]string** | Tags are free-form labels for search and grouping across categories. | [optional] 

## Methods

### NewTaxon

`func NewTaxon() *Taxon`

NewTaxon instantiates a new Taxon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonWithDefaults

`func NewTaxonWithDefaults() *Taxon`

NewTaxonWithDefaults instantiates a new Taxon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *Taxon) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *Taxon) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *Taxon) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *Taxon) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCategory

`func (o *Taxon) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Taxon) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Taxon) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Taxon) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Taxon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Taxon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Taxon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Taxon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *Taxon) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Taxon) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Taxon) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Taxon) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIcon

`func (o *Taxon) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Taxon) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Taxon) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Taxon) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *Taxon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Taxon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Taxon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Taxon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Taxon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Taxon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Taxon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Taxon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *Taxon) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Taxon) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Taxon) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Taxon) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOwner

`func (o *Taxon) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Taxon) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Taxon) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Taxon) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublished

`func (o *Taxon) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Taxon) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Taxon) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Taxon) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRoute

`func (o *Taxon) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *Taxon) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *Taxon) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *Taxon) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetTags

`func (o *Taxon) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Taxon) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Taxon) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Taxon) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


