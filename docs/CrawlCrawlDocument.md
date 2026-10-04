# CrawlCrawlDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markdown** | Pointer to **string** | Markdown is the page&#39;s content, extracted and rendered to markdown. This is the field to read. | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s title, when it carried one. | [optional] 
**Url** | Pointer to **string** | URL is the address actually read, after redirects. | [optional] 

## Methods

### NewCrawlCrawlDocument

`func NewCrawlCrawlDocument() *CrawlCrawlDocument`

NewCrawlCrawlDocument instantiates a new CrawlCrawlDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlCrawlDocumentWithDefaults

`func NewCrawlCrawlDocumentWithDefaults() *CrawlCrawlDocument`

NewCrawlCrawlDocumentWithDefaults instantiates a new CrawlCrawlDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkdown

`func (o *CrawlCrawlDocument) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *CrawlCrawlDocument) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *CrawlCrawlDocument) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *CrawlCrawlDocument) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.

### GetMetadata

`func (o *CrawlCrawlDocument) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CrawlCrawlDocument) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CrawlCrawlDocument) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CrawlCrawlDocument) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CrawlCrawlDocument) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CrawlCrawlDocument) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTitle

`func (o *CrawlCrawlDocument) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CrawlCrawlDocument) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CrawlCrawlDocument) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CrawlCrawlDocument) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *CrawlCrawlDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CrawlCrawlDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CrawlCrawlDocument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CrawlCrawlDocument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


