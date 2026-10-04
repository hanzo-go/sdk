# GraphGraphSourceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when what the source says was so, RFC 3339. Required by ingest — which records — and read by nothing in extract, which records nothing. Required for the same reason /v1/graph requires it: it is part of the assertion&#39;s content address, so re-reading one source at one instant records one set of rows however many times it is delivered. A clock read here instead would append the whole document again on every re-read. | [optional] 
**Source** | **string** | Source names where the text came from — a URL, a document id, a page title. It is stamped on every assertion as its source, and with the section number as its evidence, so a claim can be traced back to the passage that made it. Required. | 
**Subject** | Pointer to **string** | Subject is the entity the text is about before any heading names one. A document that states relations above its first heading needs it; one whose every section is headed does not. | [optional] 
**Text** | **string** | Text is the document. Relations are read from it and from nothing else: a line written &#x60;relation:: value&#x60; states one, and prose states none. Required. | 

## Methods

### NewGraphGraphSourceIn

`func NewGraphGraphSourceIn(source string, text string, ) *GraphGraphSourceIn`

NewGraphGraphSourceIn instantiates a new GraphGraphSourceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphSourceInWithDefaults

`func NewGraphGraphSourceInWithDefaults() *GraphGraphSourceIn`

NewGraphGraphSourceInWithDefaults instantiates a new GraphGraphSourceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphSourceIn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphSourceIn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphSourceIn) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *GraphGraphSourceIn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetSource

`func (o *GraphGraphSourceIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphGraphSourceIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphGraphSourceIn) SetSource(v string)`

SetSource sets Source field to given value.


### GetSubject

`func (o *GraphGraphSourceIn) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphSourceIn) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphSourceIn) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GraphGraphSourceIn) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *GraphGraphSourceIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GraphGraphSourceIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GraphGraphSourceIn) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


