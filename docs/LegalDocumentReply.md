# LegalDocumentReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Document** | Pointer to [**LegalDocumentView**](LegalDocumentView.md) | Document is the document, rendered content included. | [optional] 

## Methods

### NewLegalDocumentReply

`func NewLegalDocumentReply() *LegalDocumentReply`

NewLegalDocumentReply instantiates a new LegalDocumentReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalDocumentReplyWithDefaults

`func NewLegalDocumentReplyWithDefaults() *LegalDocumentReply`

NewLegalDocumentReplyWithDefaults instantiates a new LegalDocumentReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *LegalDocumentReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalDocumentReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalDocumentReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalDocumentReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetDocument

`func (o *LegalDocumentReply) GetDocument() LegalDocumentView`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *LegalDocumentReply) GetDocumentOk() (*LegalDocumentView, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *LegalDocumentReply) SetDocument(v LegalDocumentView)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *LegalDocumentReply) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


