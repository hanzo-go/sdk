# EmailSendMessageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is an alias for Text for backward compatibility. | [optional] 
**Campaign** | Pointer to **string** | Campaign tags the messages: 1 to 128 letters, digits, &#39;-&#39; or &#39;_&#39;, which is what an SES message tag may hold. Every link carries it as utm_campaign, the provider carries it back on every event, and the stats count by it. | [optional] 
**From** | Pointer to **string** | From is the sender address. If empty, the provider credential&#39;s from-email is used. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers provides optional custom MIME headers. | [optional] 
**Html** | Pointer to **string** | HTML is the HTML body. | [optional] 
**IdempotencyKey** | Pointer to **string** | IdempotencyKey names the send, so a retried request sends each recipient&#39;s message once: the same key and recipient name the same message. | [optional] 
**Provider** | Pointer to **string** | Provider pins the rail, ses or sendgrid. Empty is the org&#39;s default (PUT /v1/email/provider), else ses, else sendgrid, whichever the org holds credentials for. | [optional] 
**Subject** | Pointer to **string** | Subject is the message subject. | [optional] 
**Sync** | Pointer to **string** | Sync is accepted for backwards compatibility with notify sends. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID optionally references a built-in template. | [optional] 
**TemplateVars** | Pointer to **interface{}** |  | [optional] 
**Text** | Pointer to **string** | Text is the plain text body. | [optional] 
**To** | Pointer to **[]string** | To lists destination recipient addresses. Each gets a message of its own. | [optional] 

## Methods

### NewEmailSendMessageInput

`func NewEmailSendMessageInput() *EmailSendMessageInput`

NewEmailSendMessageInput instantiates a new EmailSendMessageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendMessageInputWithDefaults

`func NewEmailSendMessageInputWithDefaults() *EmailSendMessageInput`

NewEmailSendMessageInputWithDefaults instantiates a new EmailSendMessageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailSendMessageInput) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailSendMessageInput) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailSendMessageInput) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailSendMessageInput) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCampaign

`func (o *EmailSendMessageInput) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *EmailSendMessageInput) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *EmailSendMessageInput) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *EmailSendMessageInput) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetFrom

`func (o *EmailSendMessageInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailSendMessageInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailSendMessageInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailSendMessageInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailSendMessageInput) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailSendMessageInput) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailSendMessageInput) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailSendMessageInput) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHtml

`func (o *EmailSendMessageInput) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EmailSendMessageInput) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EmailSendMessageInput) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *EmailSendMessageInput) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *EmailSendMessageInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *EmailSendMessageInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *EmailSendMessageInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *EmailSendMessageInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetProvider

`func (o *EmailSendMessageInput) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmailSendMessageInput) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmailSendMessageInput) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmailSendMessageInput) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSubject

`func (o *EmailSendMessageInput) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailSendMessageInput) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailSendMessageInput) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailSendMessageInput) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSync

`func (o *EmailSendMessageInput) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *EmailSendMessageInput) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *EmailSendMessageInput) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *EmailSendMessageInput) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetTemplateId

`func (o *EmailSendMessageInput) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *EmailSendMessageInput) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *EmailSendMessageInput) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *EmailSendMessageInput) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVars

`func (o *EmailSendMessageInput) GetTemplateVars() interface{}`

GetTemplateVars returns the TemplateVars field if non-nil, zero value otherwise.

### GetTemplateVarsOk

`func (o *EmailSendMessageInput) GetTemplateVarsOk() (*interface{}, bool)`

GetTemplateVarsOk returns a tuple with the TemplateVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVars

`func (o *EmailSendMessageInput) SetTemplateVars(v interface{})`

SetTemplateVars sets TemplateVars field to given value.

### HasTemplateVars

`func (o *EmailSendMessageInput) HasTemplateVars() bool`

HasTemplateVars returns a boolean if a field has been set.

### SetTemplateVarsNil

`func (o *EmailSendMessageInput) SetTemplateVarsNil(b bool)`

 SetTemplateVarsNil sets the value for TemplateVars to be an explicit nil

### UnsetTemplateVars
`func (o *EmailSendMessageInput) UnsetTemplateVars()`

UnsetTemplateVars ensures that no value is present for TemplateVars, not even an explicit nil
### GetText

`func (o *EmailSendMessageInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EmailSendMessageInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EmailSendMessageInput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EmailSendMessageInput) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *EmailSendMessageInput) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailSendMessageInput) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailSendMessageInput) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailSendMessageInput) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


