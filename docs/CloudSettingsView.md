# CloudSettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when this configuration was first written, RFC 3339 UTC. | [optional] 
**Product** | Pointer to **string** | Product is the catalog slug this configuration belongs to. | [optional] 
**SecretKeys** | Pointer to **[]string** | SecretKeys names the secret fields that ARE set. Their VALUES live only in KMS and are never returned here — the console renders a mask. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when this configuration was last written, RFC 3339 UTC. Empty when nothing has been saved. | [optional] 

## Methods

### NewCloudSettingsView

`func NewCloudSettingsView() *CloudSettingsView`

NewCloudSettingsView instantiates a new CloudSettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSettingsViewWithDefaults

`func NewCloudSettingsViewWithDefaults() *CloudSettingsView`

NewCloudSettingsViewWithDefaults instantiates a new CloudSettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CloudSettingsView) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudSettingsView) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudSettingsView) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CloudSettingsView) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *CloudSettingsView) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *CloudSettingsView) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetCreatedAt

`func (o *CloudSettingsView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSettingsView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSettingsView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSettingsView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProduct

`func (o *CloudSettingsView) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudSettingsView) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudSettingsView) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudSettingsView) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSecretKeys

`func (o *CloudSettingsView) GetSecretKeys() []string`

GetSecretKeys returns the SecretKeys field if non-nil, zero value otherwise.

### GetSecretKeysOk

`func (o *CloudSettingsView) GetSecretKeysOk() (*[]string, bool)`

GetSecretKeysOk returns a tuple with the SecretKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeys

`func (o *CloudSettingsView) SetSecretKeys(v []string)`

SetSecretKeys sets SecretKeys field to given value.

### HasSecretKeys

`func (o *CloudSettingsView) HasSecretKeys() bool`

HasSecretKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSettingsView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSettingsView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSettingsView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSettingsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


