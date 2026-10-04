# BooksVendorRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases are the other spellings a receipt may print the vendor under; a scan matching any of them resolves to this vendor. | [optional] 
**Canonical** | Pointer to **string** | Canonical is the vendor&#39;s one true name, and the key an upsert writes by. | [optional] 
**DefaultCategory** | Pointer to **string** | DefaultCategory is the COA expense account new bills from this vendor book to. An upsert normalizes a slug (\&quot;software\&quot;) to its account number. | [optional] 

## Methods

### NewBooksVendorRow

`func NewBooksVendorRow() *BooksVendorRow`

NewBooksVendorRow instantiates a new BooksVendorRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksVendorRowWithDefaults

`func NewBooksVendorRowWithDefaults() *BooksVendorRow`

NewBooksVendorRowWithDefaults instantiates a new BooksVendorRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *BooksVendorRow) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *BooksVendorRow) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *BooksVendorRow) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *BooksVendorRow) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCanonical

`func (o *BooksVendorRow) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *BooksVendorRow) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *BooksVendorRow) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *BooksVendorRow) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetDefaultCategory

`func (o *BooksVendorRow) GetDefaultCategory() string`

GetDefaultCategory returns the DefaultCategory field if non-nil, zero value otherwise.

### GetDefaultCategoryOk

`func (o *BooksVendorRow) GetDefaultCategoryOk() (*string, bool)`

GetDefaultCategoryOk returns a tuple with the DefaultCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCategory

`func (o *BooksVendorRow) SetDefaultCategory(v string)`

SetDefaultCategory sets DefaultCategory field to given value.

### HasDefaultCategory

`func (o *BooksVendorRow) HasDefaultCategory() bool`

HasDefaultCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


