# CloudPagesDeploymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityDate** | Pointer to **string** |  | [optional] 
**CompatibilityFlags** | Pointer to **[]string** |  | [optional] 
**D1Databases** | Pointer to [**map[string]CloudPagesD1Binding**](CloudPagesD1Binding.md) |  | [optional] 
**EnvVars** | Pointer to [**map[string]CloudPagesEnvVar**](CloudPagesEnvVar.md) |  | [optional] 
**KvNamespaces** | Pointer to [**map[string]CloudPagesKVBinding**](CloudPagesKVBinding.md) |  | [optional] 
**R2Buckets** | Pointer to [**map[string]CloudPagesR2Binding**](CloudPagesR2Binding.md) |  | [optional] 

## Methods

### NewCloudPagesDeploymentConfig

`func NewCloudPagesDeploymentConfig() *CloudPagesDeploymentConfig`

NewCloudPagesDeploymentConfig instantiates a new CloudPagesDeploymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPagesDeploymentConfigWithDefaults

`func NewCloudPagesDeploymentConfigWithDefaults() *CloudPagesDeploymentConfig`

NewCloudPagesDeploymentConfigWithDefaults instantiates a new CloudPagesDeploymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityDate

`func (o *CloudPagesDeploymentConfig) GetCompatibilityDate() string`

GetCompatibilityDate returns the CompatibilityDate field if non-nil, zero value otherwise.

### GetCompatibilityDateOk

`func (o *CloudPagesDeploymentConfig) GetCompatibilityDateOk() (*string, bool)`

GetCompatibilityDateOk returns a tuple with the CompatibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityDate

`func (o *CloudPagesDeploymentConfig) SetCompatibilityDate(v string)`

SetCompatibilityDate sets CompatibilityDate field to given value.

### HasCompatibilityDate

`func (o *CloudPagesDeploymentConfig) HasCompatibilityDate() bool`

HasCompatibilityDate returns a boolean if a field has been set.

### GetCompatibilityFlags

`func (o *CloudPagesDeploymentConfig) GetCompatibilityFlags() []string`

GetCompatibilityFlags returns the CompatibilityFlags field if non-nil, zero value otherwise.

### GetCompatibilityFlagsOk

`func (o *CloudPagesDeploymentConfig) GetCompatibilityFlagsOk() (*[]string, bool)`

GetCompatibilityFlagsOk returns a tuple with the CompatibilityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityFlags

`func (o *CloudPagesDeploymentConfig) SetCompatibilityFlags(v []string)`

SetCompatibilityFlags sets CompatibilityFlags field to given value.

### HasCompatibilityFlags

`func (o *CloudPagesDeploymentConfig) HasCompatibilityFlags() bool`

HasCompatibilityFlags returns a boolean if a field has been set.

### GetD1Databases

`func (o *CloudPagesDeploymentConfig) GetD1Databases() map[string]CloudPagesD1Binding`

GetD1Databases returns the D1Databases field if non-nil, zero value otherwise.

### GetD1DatabasesOk

`func (o *CloudPagesDeploymentConfig) GetD1DatabasesOk() (*map[string]CloudPagesD1Binding, bool)`

GetD1DatabasesOk returns a tuple with the D1Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD1Databases

`func (o *CloudPagesDeploymentConfig) SetD1Databases(v map[string]CloudPagesD1Binding)`

SetD1Databases sets D1Databases field to given value.

### HasD1Databases

`func (o *CloudPagesDeploymentConfig) HasD1Databases() bool`

HasD1Databases returns a boolean if a field has been set.

### GetEnvVars

`func (o *CloudPagesDeploymentConfig) GetEnvVars() map[string]CloudPagesEnvVar`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *CloudPagesDeploymentConfig) GetEnvVarsOk() (*map[string]CloudPagesEnvVar, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *CloudPagesDeploymentConfig) SetEnvVars(v map[string]CloudPagesEnvVar)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *CloudPagesDeploymentConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetKvNamespaces

`func (o *CloudPagesDeploymentConfig) GetKvNamespaces() map[string]CloudPagesKVBinding`

GetKvNamespaces returns the KvNamespaces field if non-nil, zero value otherwise.

### GetKvNamespacesOk

`func (o *CloudPagesDeploymentConfig) GetKvNamespacesOk() (*map[string]CloudPagesKVBinding, bool)`

GetKvNamespacesOk returns a tuple with the KvNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvNamespaces

`func (o *CloudPagesDeploymentConfig) SetKvNamespaces(v map[string]CloudPagesKVBinding)`

SetKvNamespaces sets KvNamespaces field to given value.

### HasKvNamespaces

`func (o *CloudPagesDeploymentConfig) HasKvNamespaces() bool`

HasKvNamespaces returns a boolean if a field has been set.

### GetR2Buckets

`func (o *CloudPagesDeploymentConfig) GetR2Buckets() map[string]CloudPagesR2Binding`

GetR2Buckets returns the R2Buckets field if non-nil, zero value otherwise.

### GetR2BucketsOk

`func (o *CloudPagesDeploymentConfig) GetR2BucketsOk() (*map[string]CloudPagesR2Binding, bool)`

GetR2BucketsOk returns a tuple with the R2Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR2Buckets

`func (o *CloudPagesDeploymentConfig) SetR2Buckets(v map[string]CloudPagesR2Binding)`

SetR2Buckets sets R2Buckets field to given value.

### HasR2Buckets

`func (o *CloudPagesDeploymentConfig) HasR2Buckets() bool`

HasR2Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


