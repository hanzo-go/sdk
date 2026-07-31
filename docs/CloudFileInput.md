# CloudFileInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s full text. Max 1 MiB per file; binary files should simply be omitted rather than sent. | [optional] 
**Path** | Pointer to **string** | Path is the file&#39;s repo-relative path, e.g. \&quot;internal/store/db.go\&quot;. | [optional] 

## Methods

### NewCloudFileInput

`func NewCloudFileInput() *CloudFileInput`

NewCloudFileInput instantiates a new CloudFileInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFileInputWithDefaults

`func NewCloudFileInputWithDefaults() *CloudFileInput`

NewCloudFileInputWithDefaults instantiates a new CloudFileInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudFileInput) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudFileInput) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudFileInput) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudFileInput) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPath

`func (o *CloudFileInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudFileInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudFileInput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudFileInput) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


