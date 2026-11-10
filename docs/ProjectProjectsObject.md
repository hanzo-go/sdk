# ProjectProjectsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** | ETag is the store&#39;s content identity with its quotes removed. For an object written in one part it is the lowercase hex MD5 of the bytes, so a publisher compares it with its own file&#39;s digest; a multipart upload&#39;s ends in &#x60;-&lt;parts&gt;&#x60; and equals no single MD5, so it reads as different. | [optional] 
**Key** | Pointer to **string** | Key is the object&#39;s path RELATIVE to the deployment prefix — the form an upload names it by and a completion&#39;s &#x60;keys&#x60; lists it in. | [optional] 
**Size** | Pointer to **int64** | Size is the object&#39;s length in bytes. | [optional] 

## Methods

### NewProjectProjectsObject

`func NewProjectProjectsObject() *ProjectProjectsObject`

NewProjectProjectsObject instantiates a new ProjectProjectsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsObjectWithDefaults

`func NewProjectProjectsObjectWithDefaults() *ProjectProjectsObject`

NewProjectProjectsObjectWithDefaults instantiates a new ProjectProjectsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *ProjectProjectsObject) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ProjectProjectsObject) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ProjectProjectsObject) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ProjectProjectsObject) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetKey

`func (o *ProjectProjectsObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectProjectsObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectProjectsObject) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectProjectsObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSize

`func (o *ProjectProjectsObject) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProjectProjectsObject) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProjectProjectsObject) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ProjectProjectsObject) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


