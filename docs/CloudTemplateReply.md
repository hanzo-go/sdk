# CloudTemplateReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Template** | Pointer to [**CloudLegalTemplate**](CloudLegalTemplate.md) | Template is the resolved template — the org&#39;s override if it has one, else the built-in. | [optional] 

## Methods

### NewCloudTemplateReply

`func NewCloudTemplateReply() *CloudTemplateReply`

NewCloudTemplateReply instantiates a new CloudTemplateReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTemplateReplyWithDefaults

`func NewCloudTemplateReplyWithDefaults() *CloudTemplateReply`

NewCloudTemplateReplyWithDefaults instantiates a new CloudTemplateReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *CloudTemplateReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *CloudTemplateReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *CloudTemplateReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *CloudTemplateReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetTemplate

`func (o *CloudTemplateReply) GetTemplate() CloudLegalTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CloudTemplateReply) GetTemplateOk() (*CloudLegalTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CloudTemplateReply) SetTemplate(v CloudLegalTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CloudTemplateReply) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


