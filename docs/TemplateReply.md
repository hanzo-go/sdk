# TemplateReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Template** | Pointer to [**LegalTemplate**](LegalTemplate.md) | Template is the resolved template — the org&#39;s override if it has one, else the built-in. | [optional] 

## Methods

### NewTemplateReply

`func NewTemplateReply() *TemplateReply`

NewTemplateReply instantiates a new TemplateReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateReplyWithDefaults

`func NewTemplateReplyWithDefaults() *TemplateReply`

NewTemplateReplyWithDefaults instantiates a new TemplateReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *TemplateReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *TemplateReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *TemplateReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *TemplateReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetTemplate

`func (o *TemplateReply) GetTemplate() LegalTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *TemplateReply) GetTemplateOk() (*LegalTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *TemplateReply) SetTemplate(v LegalTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *TemplateReply) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


