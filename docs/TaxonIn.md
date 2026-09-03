# TaxonIn

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

### NewTaxonIn

`func NewTaxonIn() *TaxonIn`

NewTaxonIn instantiates a new TaxonIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonInWithDefaults

`func NewTaxonInWithDefaults() *TaxonIn`

NewTaxonInWithDefaults instantiates a new TaxonIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *TaxonIn) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *TaxonIn) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *TaxonIn) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *TaxonIn) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCategory

`func (o *TaxonIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaxonIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaxonIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaxonIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *TaxonIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxonIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxonIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxonIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *TaxonIn) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TaxonIn) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TaxonIn) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TaxonIn) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIcon

`func (o *TaxonIn) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TaxonIn) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TaxonIn) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TaxonIn) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *TaxonIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxonIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxonIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxonIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxonIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxonIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxonIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxonIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *TaxonIn) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TaxonIn) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TaxonIn) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TaxonIn) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPublished

`func (o *TaxonIn) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TaxonIn) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TaxonIn) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TaxonIn) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRoute

`func (o *TaxonIn) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *TaxonIn) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *TaxonIn) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *TaxonIn) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetTags

`func (o *TaxonIn) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TaxonIn) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TaxonIn) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TaxonIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


