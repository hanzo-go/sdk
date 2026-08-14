# DocumentReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Document** | Pointer to [**DocumentView**](DocumentView.md) | Document is the document, rendered content included. | [optional] 

## Methods

### NewDocumentReply

`func NewDocumentReply() *DocumentReply`

NewDocumentReply instantiates a new DocumentReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentReplyWithDefaults

`func NewDocumentReplyWithDefaults() *DocumentReply`

NewDocumentReplyWithDefaults instantiates a new DocumentReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *DocumentReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *DocumentReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *DocumentReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *DocumentReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetDocument

`func (o *DocumentReply) GetDocument() DocumentView`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DocumentReply) GetDocumentOk() (*DocumentView, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DocumentReply) SetDocument(v DocumentView)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *DocumentReply) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


