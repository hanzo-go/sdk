# IamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationSchemes** | Pointer to [**[]IamScheme**](IamScheme.md) |  | [optional] 
**Bulk** | Pointer to [**IamBulk**](IamBulk.md) |  | [optional] 
**ChangePassword** | Pointer to [**IamToggle**](IamToggle.md) |  | [optional] 
**DocumentationUri** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to [**IamToggle**](IamToggle.md) |  | [optional] 
**Filter** | Pointer to [**IamFilter**](IamFilter.md) |  | [optional] 
**Patch** | Pointer to [**IamToggle**](IamToggle.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**IamToggle**](IamToggle.md) |  | [optional] 

## Methods

### NewIamConfig

`func NewIamConfig() *IamConfig`

NewIamConfig instantiates a new IamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamConfigWithDefaults

`func NewIamConfigWithDefaults() *IamConfig`

NewIamConfigWithDefaults instantiates a new IamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationSchemes

`func (o *IamConfig) GetAuthenticationSchemes() []IamScheme`

GetAuthenticationSchemes returns the AuthenticationSchemes field if non-nil, zero value otherwise.

### GetAuthenticationSchemesOk

`func (o *IamConfig) GetAuthenticationSchemesOk() (*[]IamScheme, bool)`

GetAuthenticationSchemesOk returns a tuple with the AuthenticationSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSchemes

`func (o *IamConfig) SetAuthenticationSchemes(v []IamScheme)`

SetAuthenticationSchemes sets AuthenticationSchemes field to given value.

### HasAuthenticationSchemes

`func (o *IamConfig) HasAuthenticationSchemes() bool`

HasAuthenticationSchemes returns a boolean if a field has been set.

### GetBulk

`func (o *IamConfig) GetBulk() IamBulk`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *IamConfig) GetBulkOk() (*IamBulk, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *IamConfig) SetBulk(v IamBulk)`

SetBulk sets Bulk field to given value.

### HasBulk

`func (o *IamConfig) HasBulk() bool`

HasBulk returns a boolean if a field has been set.

### GetChangePassword

`func (o *IamConfig) GetChangePassword() IamToggle`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *IamConfig) GetChangePasswordOk() (*IamToggle, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *IamConfig) SetChangePassword(v IamToggle)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *IamConfig) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### GetDocumentationUri

`func (o *IamConfig) GetDocumentationUri() string`

GetDocumentationUri returns the DocumentationUri field if non-nil, zero value otherwise.

### GetDocumentationUriOk

`func (o *IamConfig) GetDocumentationUriOk() (*string, bool)`

GetDocumentationUriOk returns a tuple with the DocumentationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUri

`func (o *IamConfig) SetDocumentationUri(v string)`

SetDocumentationUri sets DocumentationUri field to given value.

### HasDocumentationUri

`func (o *IamConfig) HasDocumentationUri() bool`

HasDocumentationUri returns a boolean if a field has been set.

### GetEtag

`func (o *IamConfig) GetEtag() IamToggle`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *IamConfig) GetEtagOk() (*IamToggle, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *IamConfig) SetEtag(v IamToggle)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *IamConfig) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetFilter

`func (o *IamConfig) GetFilter() IamFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamConfig) GetFilterOk() (*IamFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamConfig) SetFilter(v IamFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IamConfig) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPatch

`func (o *IamConfig) GetPatch() IamToggle`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *IamConfig) GetPatchOk() (*IamToggle, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *IamConfig) SetPatch(v IamToggle)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *IamConfig) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetSchemas

`func (o *IamConfig) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IamConfig) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IamConfig) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *IamConfig) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetSort

`func (o *IamConfig) GetSort() IamToggle`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IamConfig) GetSortOk() (*IamToggle, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IamConfig) SetSort(v IamToggle)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IamConfig) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


