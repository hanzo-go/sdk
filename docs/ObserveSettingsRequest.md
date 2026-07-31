# ObserveSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Non-secret config; stored verbatim (bounded at 64KiB). | [optional] 
**Secrets** | Pointer to **map[string]string** | Secret fields; VALUES routed to KMS, never SQLite. Each key must match &#x60;^[a-z0-9][a-z0-9._-]{0,62}$&#x60;; each value is bounded at 8KiB. An empty value or the mask sentinel (&#x60;••••••••&#x60;) means \&quot;unchanged\&quot;. Max 64 fields.  | [optional] 

## Methods

### NewObserveSettingsRequest

`func NewObserveSettingsRequest() *ObserveSettingsRequest`

NewObserveSettingsRequest instantiates a new ObserveSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveSettingsRequestWithDefaults

`func NewObserveSettingsRequestWithDefaults() *ObserveSettingsRequest`

NewObserveSettingsRequestWithDefaults instantiates a new ObserveSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ObserveSettingsRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObserveSettingsRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObserveSettingsRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ObserveSettingsRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecrets

`func (o *ObserveSettingsRequest) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ObserveSettingsRequest) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ObserveSettingsRequest) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ObserveSettingsRequest) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


