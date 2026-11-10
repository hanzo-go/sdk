# FormWebhookIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled is whether responses are posted. Default true. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s id or slug, from the path. | [optional] 
**Rotate** | Pointer to **bool** | Rotate mints a new secret on a replace, and answers it. | [optional] 
**Secret** | Pointer to **string** | Secret sets the signing secret: 16 to 256 characters. Absent on a make, one is minted; absent on a replace, the secret held is kept. | [optional] 
**Tag** | Pointer to **string** | Tag names the webhook in its form, at most 64 characters. | [optional] 
**Url** | Pointer to **string** | URL is the https address posted to. Required. | [optional] 
**Webhook** | Pointer to **string** | Webhook is the webhook&#39;s id, from the path of a replace. | [optional] 

## Methods

### NewFormWebhookIn

`func NewFormWebhookIn() *FormWebhookIn`

NewFormWebhookIn instantiates a new FormWebhookIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWebhookInWithDefaults

`func NewFormWebhookInWithDefaults() *FormWebhookIn`

NewFormWebhookInWithDefaults instantiates a new FormWebhookIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FormWebhookIn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FormWebhookIn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FormWebhookIn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FormWebhookIn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *FormWebhookIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormWebhookIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormWebhookIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormWebhookIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRotate

`func (o *FormWebhookIn) GetRotate() bool`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *FormWebhookIn) GetRotateOk() (*bool, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *FormWebhookIn) SetRotate(v bool)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *FormWebhookIn) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetSecret

`func (o *FormWebhookIn) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *FormWebhookIn) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *FormWebhookIn) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *FormWebhookIn) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTag

`func (o *FormWebhookIn) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FormWebhookIn) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FormWebhookIn) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *FormWebhookIn) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUrl

`func (o *FormWebhookIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormWebhookIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormWebhookIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FormWebhookIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebhook

`func (o *FormWebhookIn) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *FormWebhookIn) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *FormWebhookIn) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *FormWebhookIn) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


