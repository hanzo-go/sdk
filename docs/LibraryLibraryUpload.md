# LibraryLibraryUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]LibraryField**](LibraryField.md) | Headers are the ones the PUT must carry. | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewLibraryLibraryUpload

`func NewLibraryLibraryUpload() *LibraryLibraryUpload`

NewLibraryLibraryUpload instantiates a new LibraryLibraryUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryLibraryUploadWithDefaults

`func NewLibraryLibraryUploadWithDefaults() *LibraryLibraryUpload`

NewLibraryLibraryUploadWithDefaults instantiates a new LibraryLibraryUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *LibraryLibraryUpload) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *LibraryLibraryUpload) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *LibraryLibraryUpload) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *LibraryLibraryUpload) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetHeaders

`func (o *LibraryLibraryUpload) GetHeaders() []LibraryField`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *LibraryLibraryUpload) GetHeadersOk() (*[]LibraryField, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *LibraryLibraryUpload) SetHeaders(v []LibraryField)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *LibraryLibraryUpload) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *LibraryLibraryUpload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LibraryLibraryUpload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LibraryLibraryUpload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LibraryLibraryUpload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


