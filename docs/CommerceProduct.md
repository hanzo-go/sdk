# CommerceProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Slug** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Excerpt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**CommerceMedia**](CommerceMedia.md) |  | [optional] 
**Media** | Pointer to [**[]CommerceMedia**](CommerceMedia.md) |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Preorder** | Pointer to **bool** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**Variants** | Pointer to [**[]CommerceVariant**](CommerceVariant.md) |  | [optional] 
**Options** | Pointer to [**[]CommerceProductOption**](CommerceProductOption.md) |  | [optional] 
**Price** | Pointer to **int32** |  | [optional] 
**ListPrice** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceProduct

`func NewCommerceProduct() *CommerceProduct`

NewCommerceProduct instantiates a new CommerceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceProductWithDefaults

`func NewCommerceProductWithDefaults() *CommerceProduct`

NewCommerceProductWithDefaults instantiates a new CommerceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *CommerceProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CommerceProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CommerceProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CommerceProduct) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSku

`func (o *CommerceProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CommerceProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CommerceProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CommerceProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUpc

`func (o *CommerceProduct) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *CommerceProduct) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *CommerceProduct) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *CommerceProduct) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetName

`func (o *CommerceProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommerceProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommerceProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommerceProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeadline

`func (o *CommerceProduct) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CommerceProduct) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CommerceProduct) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *CommerceProduct) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetExcerpt

`func (o *CommerceProduct) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *CommerceProduct) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *CommerceProduct) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *CommerceProduct) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetDescription

`func (o *CommerceProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommerceProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommerceProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommerceProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *CommerceProduct) GetImage() CommerceMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CommerceProduct) GetImageOk() (*CommerceMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CommerceProduct) SetImage(v CommerceMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *CommerceProduct) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMedia

`func (o *CommerceProduct) GetMedia() []CommerceMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *CommerceProduct) GetMediaOk() (*[]CommerceMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *CommerceProduct) SetMedia(v []CommerceMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *CommerceProduct) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetAvailable

`func (o *CommerceProduct) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CommerceProduct) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CommerceProduct) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CommerceProduct) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetHidden

`func (o *CommerceProduct) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *CommerceProduct) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *CommerceProduct) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *CommerceProduct) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPreorder

`func (o *CommerceProduct) GetPreorder() bool`

GetPreorder returns the Preorder field if non-nil, zero value otherwise.

### GetPreorderOk

`func (o *CommerceProduct) GetPreorderOk() (*bool, bool)`

GetPreorderOk returns a tuple with the Preorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorder

`func (o *CommerceProduct) SetPreorder(v bool)`

SetPreorder sets Preorder field to given value.

### HasPreorder

`func (o *CommerceProduct) HasPreorder() bool`

HasPreorder returns a boolean if a field has been set.

### GetTaxable

`func (o *CommerceProduct) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *CommerceProduct) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *CommerceProduct) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *CommerceProduct) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetVariants

`func (o *CommerceProduct) GetVariants() []CommerceVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *CommerceProduct) GetVariantsOk() (*[]CommerceVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *CommerceProduct) SetVariants(v []CommerceVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *CommerceProduct) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetOptions

`func (o *CommerceProduct) GetOptions() []CommerceProductOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CommerceProduct) GetOptionsOk() (*[]CommerceProductOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CommerceProduct) SetOptions(v []CommerceProductOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CommerceProduct) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPrice

`func (o *CommerceProduct) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CommerceProduct) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CommerceProduct) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CommerceProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetListPrice

`func (o *CommerceProduct) GetListPrice() int32`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *CommerceProduct) GetListPriceOk() (*int32, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *CommerceProduct) SetListPrice(v int32)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *CommerceProduct) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceProduct) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceProduct) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceProduct) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceProduct) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceProduct) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceProduct) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceProduct) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceProduct) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceProduct) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceProduct) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


