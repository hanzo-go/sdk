# TaxonomyCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** | Brands are the brands whose console shows this category. Absent means every brand. | [optional] 
**Id** | Pointer to **string** | ID is the stable slug this category is addressed by, e.g. \&quot;observe\&quot;. | [optional] 
**Label** | Pointer to **string** | Label is the display name, e.g. \&quot;Observe\&quot;. | [optional] 
**Order** | Pointer to **int64** | Order is where the category sits among its siblings, ascending. | [optional] 
**Owner** | Pointer to **string** | Owner is the org this category belongs to: the platform&#39;s own org for a category every tenant sees, or your org for one you added. It tells a console which rows it may offer to edit. | [optional] 
**Summary** | Pointer to **string** | Summary is the one line describing what the category groups, shown as the header copy on its landing page. | [optional] 
**Taxa** | Pointer to [**[]TaxonomyTaxon**](TaxonomyTaxon.md) | Taxa are the products filed under this category, in display order. | [optional] 

## Methods

### NewTaxonomyCategory

`func NewTaxonomyCategory() *TaxonomyCategory`

NewTaxonomyCategory instantiates a new TaxonomyCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyCategoryWithDefaults

`func NewTaxonomyCategoryWithDefaults() *TaxonomyCategory`

NewTaxonomyCategoryWithDefaults instantiates a new TaxonomyCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *TaxonomyCategory) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *TaxonomyCategory) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *TaxonomyCategory) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *TaxonomyCategory) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetId

`func (o *TaxonomyCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxonomyCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxonomyCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxonomyCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TaxonomyCategory) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TaxonomyCategory) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TaxonomyCategory) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TaxonomyCategory) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOrder

`func (o *TaxonomyCategory) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TaxonomyCategory) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TaxonomyCategory) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TaxonomyCategory) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOwner

`func (o *TaxonomyCategory) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TaxonomyCategory) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TaxonomyCategory) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TaxonomyCategory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSummary

`func (o *TaxonomyCategory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TaxonomyCategory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TaxonomyCategory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TaxonomyCategory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTaxa

`func (o *TaxonomyCategory) GetTaxa() []TaxonomyTaxon`

GetTaxa returns the Taxa field if non-nil, zero value otherwise.

### GetTaxaOk

`func (o *TaxonomyCategory) GetTaxaOk() (*[]TaxonomyTaxon, bool)`

GetTaxaOk returns a tuple with the Taxa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxa

`func (o *TaxonomyCategory) SetTaxa(v []TaxonomyTaxon)`

SetTaxa sets Taxa field to given value.

### HasTaxa

`func (o *TaxonomyCategory) HasTaxa() bool`

HasTaxa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


