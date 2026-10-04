# TaxonomyTaxonomy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]TaxonomyCategory**](TaxonomyCategory.md) | Categories are the groupings, in display order, each with its own taxa. | [optional] 

## Methods

### NewTaxonomyTaxonomy

`func NewTaxonomyTaxonomy() *TaxonomyTaxonomy`

NewTaxonomyTaxonomy instantiates a new TaxonomyTaxonomy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyTaxonomyWithDefaults

`func NewTaxonomyTaxonomyWithDefaults() *TaxonomyTaxonomy`

NewTaxonomyTaxonomyWithDefaults instantiates a new TaxonomyTaxonomy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *TaxonomyTaxonomy) GetCategories() []TaxonomyCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TaxonomyTaxonomy) GetCategoriesOk() (*[]TaxonomyCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TaxonomyTaxonomy) SetCategories(v []TaxonomyCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TaxonomyTaxonomy) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


