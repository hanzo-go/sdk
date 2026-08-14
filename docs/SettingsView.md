# SettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when this configuration was first written, RFC 3339 UTC. | [optional] 
**Product** | Pointer to **string** | Product is the catalog slug this configuration belongs to. | [optional] 
**SecretKeys** | Pointer to **[]string** | SecretKeys names the secret fields that ARE set. Their VALUES live only in KMS and are never returned here — the console renders a mask. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when this configuration was last written, RFC 3339 UTC. Empty when nothing has been saved. | [optional] 

## Methods

### NewSettingsView

`func NewSettingsView() *SettingsView`

NewSettingsView instantiates a new SettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsViewWithDefaults

`func NewSettingsViewWithDefaults() *SettingsView`

NewSettingsViewWithDefaults instantiates a new SettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SettingsView) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SettingsView) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SettingsView) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SettingsView) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *SettingsView) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *SettingsView) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetCreatedAt

`func (o *SettingsView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SettingsView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProduct

`func (o *SettingsView) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SettingsView) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SettingsView) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SettingsView) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSecretKeys

`func (o *SettingsView) GetSecretKeys() []string`

GetSecretKeys returns the SecretKeys field if non-nil, zero value otherwise.

### GetSecretKeysOk

`func (o *SettingsView) GetSecretKeysOk() (*[]string, bool)`

GetSecretKeysOk returns a tuple with the SecretKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeys

`func (o *SettingsView) SetSecretKeys(v []string)`

SetSecretKeys sets SecretKeys field to given value.

### HasSecretKeys

`func (o *SettingsView) HasSecretKeys() bool`

HasSecretKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SettingsView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingsView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingsView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


