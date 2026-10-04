# LegalTemplateReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Template** | Pointer to [**LegalLegalTemplate**](LegalLegalTemplate.md) | Template is the resolved template — the org&#39;s override if it has one, else the built-in. | [optional] 

## Methods

### NewLegalTemplateReply

`func NewLegalTemplateReply() *LegalTemplateReply`

NewLegalTemplateReply instantiates a new LegalTemplateReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalTemplateReplyWithDefaults

`func NewLegalTemplateReplyWithDefaults() *LegalTemplateReply`

NewLegalTemplateReplyWithDefaults instantiates a new LegalTemplateReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *LegalTemplateReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalTemplateReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalTemplateReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalTemplateReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetTemplate

`func (o *LegalTemplateReply) GetTemplate() LegalLegalTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *LegalTemplateReply) GetTemplateOk() (*LegalLegalTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *LegalTemplateReply) SetTemplate(v LegalLegalTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *LegalTemplateReply) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


