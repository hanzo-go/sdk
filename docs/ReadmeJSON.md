# ReadmeJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s text, verbatim and unrendered. | [optional] 
**Encoding** | Pointer to **string** | Encoding is always \&quot;utf8\&quot; — a README is text by definition. | [optional] 
**Path** | Pointer to **string** | Path is the file the README was found at (README.md, README, …). | [optional] 

## Methods

### NewReadmeJSON

`func NewReadmeJSON() *ReadmeJSON`

NewReadmeJSON instantiates a new ReadmeJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadmeJSONWithDefaults

`func NewReadmeJSONWithDefaults() *ReadmeJSON`

NewReadmeJSONWithDefaults instantiates a new ReadmeJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ReadmeJSON) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReadmeJSON) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReadmeJSON) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ReadmeJSON) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *ReadmeJSON) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ReadmeJSON) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ReadmeJSON) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ReadmeJSON) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *ReadmeJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReadmeJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReadmeJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReadmeJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


