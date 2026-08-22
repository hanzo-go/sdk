# StorefrontResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrl** | Pointer to **string** | ImageURL is the absolute URL the listing&#39;s headerImage now points at. The image is REFERENCED, not copied — it stays in the org&#39;s studio output bucket, so removing it there empties the storefront tile. | [optional] 
**Slug** | Pointer to **string** | Slug is the product handle the image was attached to. It IS the asset&#39;s &#x60;design&#x60; field — that equality is the whole join between the studio and the catalog, which is why an asset with no design produces no storefront result at all. | [optional] 
**Status** | Pointer to **string** | Status is one of \&quot;published\&quot; (the product image was set), \&quot;not_configured\&quot; (no commerce edge, no store provisioned for the org, or a token that is not admin on the store — a fail-closed no-op) or \&quot;failed\&quot; (commerce answered and errored). None of the three fails the transition that produced it. | [optional] 
**Store** | Pointer to **string** | Store is the commerce store id the image landed in, resolved for the org mid-call. Present only on \&quot;published\&quot;: a result that never got that far carries none. | [optional] 

## Methods

### NewStorefrontResult

`func NewStorefrontResult() *StorefrontResult`

NewStorefrontResult instantiates a new StorefrontResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontResultWithDefaults

`func NewStorefrontResultWithDefaults() *StorefrontResult`

NewStorefrontResultWithDefaults instantiates a new StorefrontResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrl

`func (o *StorefrontResult) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *StorefrontResult) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *StorefrontResult) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *StorefrontResult) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetSlug

`func (o *StorefrontResult) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StorefrontResult) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StorefrontResult) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StorefrontResult) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *StorefrontResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorefrontResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorefrontResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorefrontResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *StorefrontResult) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *StorefrontResult) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *StorefrontResult) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *StorefrontResult) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


