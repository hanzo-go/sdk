# CloudSignReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**CloudDocumentSummary**](CloudDocumentSummary.md) | Document is the document, now out for signature. Its rendered body is not repeated here. | [optional] 
**EsignRef** | Pointer to **string** | EsignRef is the provider&#39;s own reference for the request — what a webhook or a status poll quotes. | [optional] 
**Provider** | Pointer to **string** | Provider names the e-signature provider that took the request. \&quot;manual\&quot; means no provider is wired on this deployment and the org fulfils it out of band. | [optional] 

## Methods

### NewCloudSignReply

`func NewCloudSignReply() *CloudSignReply`

NewCloudSignReply instantiates a new CloudSignReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSignReplyWithDefaults

`func NewCloudSignReplyWithDefaults() *CloudSignReply`

NewCloudSignReplyWithDefaults instantiates a new CloudSignReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *CloudSignReply) GetDocument() CloudDocumentSummary`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CloudSignReply) GetDocumentOk() (*CloudDocumentSummary, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CloudSignReply) SetDocument(v CloudDocumentSummary)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *CloudSignReply) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetEsignRef

`func (o *CloudSignReply) GetEsignRef() string`

GetEsignRef returns the EsignRef field if non-nil, zero value otherwise.

### GetEsignRefOk

`func (o *CloudSignReply) GetEsignRefOk() (*string, bool)`

GetEsignRefOk returns a tuple with the EsignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignRef

`func (o *CloudSignReply) SetEsignRef(v string)`

SetEsignRef sets EsignRef field to given value.

### HasEsignRef

`func (o *CloudSignReply) HasEsignRef() bool`

HasEsignRef returns a boolean if a field has been set.

### GetProvider

`func (o *CloudSignReply) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudSignReply) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudSignReply) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudSignReply) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


