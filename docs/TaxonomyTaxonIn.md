# TaxonomyTaxonIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** | Brands are the brands whose console shows it. Omit for every brand its category admits. | [optional] 
**Category** | Pointer to **string** | Category is the id of an EXISTING category to file it under. Required. | [optional] 
**Description** | Pointer to **string** | Description is the one line shown beneath the name. | [optional] 
**Href** | Pointer to **string** | Href is the absolute URL an external product launches. Give this or route, never both. | [optional] 
**Icon** | Pointer to **string** | Icon names the icon the surface renders, e.g. \&quot;Database\&quot;. | [optional] 
**Id** | Pointer to **string** | ID is the taxon slug to write, from the path. | [optional] 
**Name** | Pointer to **string** | Name is the display name. Required. | [optional] 
**Order** | Pointer to **int64** | Order is where it sits within its category, ascending. | [optional] 
**Published** | Pointer to **bool** | Published is whether it is shown. Omitted means published — a taxon someone took the trouble to write is meant to be seen, and hiding one is the deliberate act. | [optional] 
**Route** | Pointer to **string** | Route is the in-console path it opens, e.g. \&quot;/vector\&quot;. Give this or href, never both. | [optional] 
**Tags** | Pointer to **[]string** | Tags are free-form labels for search and grouping across categories. | [optional] 

## Methods

### NewTaxonomyTaxonIn

`func NewTaxonomyTaxonIn() *TaxonomyTaxonIn`

NewTaxonomyTaxonIn instantiates a new TaxonomyTaxonIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyTaxonInWithDefaults

`func NewTaxonomyTaxonInWithDefaults() *TaxonomyTaxonIn`

NewTaxonomyTaxonInWithDefaults instantiates a new TaxonomyTaxonIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *TaxonomyTaxonIn) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *TaxonomyTaxonIn) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *TaxonomyTaxonIn) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *TaxonomyTaxonIn) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCategory

`func (o *TaxonomyTaxonIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaxonomyTaxonIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaxonomyTaxonIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaxonomyTaxonIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *TaxonomyTaxonIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxonomyTaxonIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxonomyTaxonIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxonomyTaxonIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *TaxonomyTaxonIn) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TaxonomyTaxonIn) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TaxonomyTaxonIn) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TaxonomyTaxonIn) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIcon

`func (o *TaxonomyTaxonIn) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TaxonomyTaxonIn) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TaxonomyTaxonIn) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TaxonomyTaxonIn) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *TaxonomyTaxonIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxonomyTaxonIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxonomyTaxonIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxonomyTaxonIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxonomyTaxonIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxonomyTaxonIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxonomyTaxonIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxonomyTaxonIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *TaxonomyTaxonIn) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TaxonomyTaxonIn) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TaxonomyTaxonIn) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TaxonomyTaxonIn) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPublished

`func (o *TaxonomyTaxonIn) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TaxonomyTaxonIn) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TaxonomyTaxonIn) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TaxonomyTaxonIn) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRoute

`func (o *TaxonomyTaxonIn) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *TaxonomyTaxonIn) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *TaxonomyTaxonIn) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *TaxonomyTaxonIn) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetTags

`func (o *TaxonomyTaxonIn) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TaxonomyTaxonIn) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TaxonomyTaxonIn) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TaxonomyTaxonIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


