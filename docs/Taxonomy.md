# Taxonomy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]Category**](Category.md) | Categories are the groupings, in display order, each with its own taxa. | [optional] 

## Methods

### NewTaxonomy

`func NewTaxonomy() *Taxonomy`

NewTaxonomy instantiates a new Taxonomy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyWithDefaults

`func NewTaxonomyWithDefaults() *Taxonomy`

NewTaxonomyWithDefaults instantiates a new Taxonomy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *Taxonomy) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Taxonomy) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Taxonomy) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Taxonomy) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


