# FormWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Created is when the webhook was made, RFC 3339. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether responses are posted. | [optional] 
**Form** | Pointer to **string** | Form is the form&#39;s id. | [optional] 
**Id** | Pointer to **string** | ID is the webhook&#39;s id. | [optional] 
**Secret** | Pointer to **string** | Secret is the signing secret, answered once: when the webhook is made or its secret is set or rotated. Every delivery carries Hanzo-Signature: t&#x3D;&lt;unix seconds&gt;,sha256&#x3D;&lt;base64 of HMAC-SHA256(secret, \&quot;&lt;seconds&gt;.\&quot; + body)&gt;; a receiver refuses one older than five minutes. It is kept in KMS and never read back. | [optional] 
**Tag** | Pointer to **string** | Tag names the webhook in its form. At most 64 characters. | [optional] 
**Updated** | Pointer to **string** | Updated is when it last changed, RFC 3339. | [optional] 
**Url** | Pointer to **string** | URL is the https address posted to. | [optional] 

## Methods

### NewFormWebhook

`func NewFormWebhook() *FormWebhook`

NewFormWebhook instantiates a new FormWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWebhookWithDefaults

`func NewFormWebhookWithDefaults() *FormWebhook`

NewFormWebhookWithDefaults instantiates a new FormWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FormWebhook) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormWebhook) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormWebhook) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormWebhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *FormWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FormWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FormWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FormWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForm

`func (o *FormWebhook) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormWebhook) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormWebhook) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormWebhook) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetId

`func (o *FormWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecret

`func (o *FormWebhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *FormWebhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *FormWebhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *FormWebhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTag

`func (o *FormWebhook) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FormWebhook) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FormWebhook) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *FormWebhook) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUpdated

`func (o *FormWebhook) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FormWebhook) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FormWebhook) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FormWebhook) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *FormWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FormWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


