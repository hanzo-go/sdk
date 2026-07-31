# CloudPushFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s bytes, carried per Encoding. | [optional] 
**Encoding** | Pointer to **string** | Encoding is \&quot;base64\&quot;, or \&quot;utf-8\&quot; (the default, also \&quot;utf8\&quot; / \&quot;text\&quot;). | [optional] 
**Path** | Pointer to **string** | Path is repo-relative. Absolute or traversing paths are refused. | [optional] 

## Methods

### NewCloudPushFile

`func NewCloudPushFile() *CloudPushFile`

NewCloudPushFile instantiates a new CloudPushFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPushFileWithDefaults

`func NewCloudPushFileWithDefaults() *CloudPushFile`

NewCloudPushFileWithDefaults instantiates a new CloudPushFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudPushFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudPushFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudPushFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudPushFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *CloudPushFile) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *CloudPushFile) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *CloudPushFile) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *CloudPushFile) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *CloudPushFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudPushFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudPushFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudPushFile) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


