# SettingsSettingsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Config is the product&#39;s non-secret configuration, stored verbatim. Bounded at 64 KiB once serialized. Omit it to store an empty object. | [optional] 
**Product** | Pointer to **string** | Product is the catalog slug, from the PATH. zip binds the path last, so the URL names the product being written whatever a body field claims. | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets are the secret fields, by name. Each VALUE is sealed into KMS and never reaches this deployment&#39;s database; a value that is empty or equal to the mask the read path returns means \&quot;unchanged\&quot; and is skipped, so a console round-trip cannot blank a stored secret. A key must match ^[a-z0-9][a-z0-9._-]{0,62}$, a value is bounded at 8 KiB, and an org may hold at most 64 secret fields per product. | [optional] 

## Methods

### NewSettingsSettingsReq

`func NewSettingsSettingsReq() *SettingsSettingsReq`

NewSettingsSettingsReq instantiates a new SettingsSettingsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsSettingsReqWithDefaults

`func NewSettingsSettingsReqWithDefaults() *SettingsSettingsReq`

NewSettingsSettingsReqWithDefaults instantiates a new SettingsSettingsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SettingsSettingsReq) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SettingsSettingsReq) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SettingsSettingsReq) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SettingsSettingsReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetProduct

`func (o *SettingsSettingsReq) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SettingsSettingsReq) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SettingsSettingsReq) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SettingsSettingsReq) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSecrets

`func (o *SettingsSettingsReq) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SettingsSettingsReq) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SettingsSettingsReq) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *SettingsSettingsReq) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


