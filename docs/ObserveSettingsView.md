# ObserveSettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Non-secret config document (opaque JSON object), stored verbatim. | [optional] 
**SecretKeys** | Pointer to **[]string** | Names of set secret fields (values masked, never returned). | [optional] 
**UpdatedAt** | Pointer to **string** | RFC3339 timestamp (UTC); empty when unset. | [optional] 
**CreatedAt** | Pointer to **string** | RFC3339 timestamp (UTC); empty when unset. | [optional] 

## Methods

### NewObserveSettingsView

`func NewObserveSettingsView() *ObserveSettingsView`

NewObserveSettingsView instantiates a new ObserveSettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveSettingsViewWithDefaults

`func NewObserveSettingsViewWithDefaults() *ObserveSettingsView`

NewObserveSettingsViewWithDefaults instantiates a new ObserveSettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ObserveSettingsView) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ObserveSettingsView) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ObserveSettingsView) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ObserveSettingsView) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetConfig

`func (o *ObserveSettingsView) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObserveSettingsView) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObserveSettingsView) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ObserveSettingsView) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecretKeys

`func (o *ObserveSettingsView) GetSecretKeys() []string`

GetSecretKeys returns the SecretKeys field if non-nil, zero value otherwise.

### GetSecretKeysOk

`func (o *ObserveSettingsView) GetSecretKeysOk() (*[]string, bool)`

GetSecretKeysOk returns a tuple with the SecretKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeys

`func (o *ObserveSettingsView) SetSecretKeys(v []string)`

SetSecretKeys sets SecretKeys field to given value.

### HasSecretKeys

`func (o *ObserveSettingsView) HasSecretKeys() bool`

HasSecretKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ObserveSettingsView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ObserveSettingsView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ObserveSettingsView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ObserveSettingsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ObserveSettingsView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObserveSettingsView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObserveSettingsView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ObserveSettingsView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


