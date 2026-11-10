# BillingInvoiceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand is the brand that bills (hanzo, lux, zoo, pars), whose mark and address a rendered invoice carries. The edge sets it from the caller&#39;s issuer; empty is Hanzo. | [optional] 
**Id** | Pointer to **string** | ID is the invoice id. | [optional] 

## Methods

### NewBillingInvoiceRef

`func NewBillingInvoiceRef() *BillingInvoiceRef`

NewBillingInvoiceRef instantiates a new BillingInvoiceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceRefWithDefaults

`func NewBillingInvoiceRefWithDefaults() *BillingInvoiceRef`

NewBillingInvoiceRefWithDefaults instantiates a new BillingInvoiceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *BillingInvoiceRef) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *BillingInvoiceRef) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *BillingInvoiceRef) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *BillingInvoiceRef) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetId

`func (o *BillingInvoiceRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoiceRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoiceRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoiceRef) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


