# ProjectsSiteFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Relative path within the site (no leading \&quot;/\&quot;, no \&quot;..\&quot;); index.html must be present at the root. | 
**Content** | **string** | Full file contents. HTML files get a mobile viewport tag injected when absent. | 

## Methods

### NewProjectsSiteFile

`func NewProjectsSiteFile(path string, content string, ) *ProjectsSiteFile`

NewProjectsSiteFile instantiates a new ProjectsSiteFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSiteFileWithDefaults

`func NewProjectsSiteFileWithDefaults() *ProjectsSiteFile`

NewProjectsSiteFileWithDefaults instantiates a new ProjectsSiteFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ProjectsSiteFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ProjectsSiteFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ProjectsSiteFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *ProjectsSiteFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProjectsSiteFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProjectsSiteFile) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


