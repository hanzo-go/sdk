# PagesDeploymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityDate** | Pointer to **string** |  | [optional] 
**CompatibilityFlags** | Pointer to **[]string** |  | [optional] 
**D1Databases** | Pointer to [**map[string]PagesD1Binding**](PagesD1Binding.md) |  | [optional] 
**EnvVars** | Pointer to [**map[string]PagesEnvVar**](PagesEnvVar.md) |  | [optional] 
**KvNamespaces** | Pointer to [**map[string]PagesKVBinding**](PagesKVBinding.md) |  | [optional] 
**R2Buckets** | Pointer to [**map[string]PagesR2Binding**](PagesR2Binding.md) |  | [optional] 

## Methods

### NewPagesDeploymentConfig

`func NewPagesDeploymentConfig() *PagesDeploymentConfig`

NewPagesDeploymentConfig instantiates a new PagesDeploymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesDeploymentConfigWithDefaults

`func NewPagesDeploymentConfigWithDefaults() *PagesDeploymentConfig`

NewPagesDeploymentConfigWithDefaults instantiates a new PagesDeploymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityDate

`func (o *PagesDeploymentConfig) GetCompatibilityDate() string`

GetCompatibilityDate returns the CompatibilityDate field if non-nil, zero value otherwise.

### GetCompatibilityDateOk

`func (o *PagesDeploymentConfig) GetCompatibilityDateOk() (*string, bool)`

GetCompatibilityDateOk returns a tuple with the CompatibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityDate

`func (o *PagesDeploymentConfig) SetCompatibilityDate(v string)`

SetCompatibilityDate sets CompatibilityDate field to given value.

### HasCompatibilityDate

`func (o *PagesDeploymentConfig) HasCompatibilityDate() bool`

HasCompatibilityDate returns a boolean if a field has been set.

### GetCompatibilityFlags

`func (o *PagesDeploymentConfig) GetCompatibilityFlags() []string`

GetCompatibilityFlags returns the CompatibilityFlags field if non-nil, zero value otherwise.

### GetCompatibilityFlagsOk

`func (o *PagesDeploymentConfig) GetCompatibilityFlagsOk() (*[]string, bool)`

GetCompatibilityFlagsOk returns a tuple with the CompatibilityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityFlags

`func (o *PagesDeploymentConfig) SetCompatibilityFlags(v []string)`

SetCompatibilityFlags sets CompatibilityFlags field to given value.

### HasCompatibilityFlags

`func (o *PagesDeploymentConfig) HasCompatibilityFlags() bool`

HasCompatibilityFlags returns a boolean if a field has been set.

### GetD1Databases

`func (o *PagesDeploymentConfig) GetD1Databases() map[string]PagesD1Binding`

GetD1Databases returns the D1Databases field if non-nil, zero value otherwise.

### GetD1DatabasesOk

`func (o *PagesDeploymentConfig) GetD1DatabasesOk() (*map[string]PagesD1Binding, bool)`

GetD1DatabasesOk returns a tuple with the D1Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD1Databases

`func (o *PagesDeploymentConfig) SetD1Databases(v map[string]PagesD1Binding)`

SetD1Databases sets D1Databases field to given value.

### HasD1Databases

`func (o *PagesDeploymentConfig) HasD1Databases() bool`

HasD1Databases returns a boolean if a field has been set.

### GetEnvVars

`func (o *PagesDeploymentConfig) GetEnvVars() map[string]PagesEnvVar`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *PagesDeploymentConfig) GetEnvVarsOk() (*map[string]PagesEnvVar, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *PagesDeploymentConfig) SetEnvVars(v map[string]PagesEnvVar)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *PagesDeploymentConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetKvNamespaces

`func (o *PagesDeploymentConfig) GetKvNamespaces() map[string]PagesKVBinding`

GetKvNamespaces returns the KvNamespaces field if non-nil, zero value otherwise.

### GetKvNamespacesOk

`func (o *PagesDeploymentConfig) GetKvNamespacesOk() (*map[string]PagesKVBinding, bool)`

GetKvNamespacesOk returns a tuple with the KvNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvNamespaces

`func (o *PagesDeploymentConfig) SetKvNamespaces(v map[string]PagesKVBinding)`

SetKvNamespaces sets KvNamespaces field to given value.

### HasKvNamespaces

`func (o *PagesDeploymentConfig) HasKvNamespaces() bool`

HasKvNamespaces returns a boolean if a field has been set.

### GetR2Buckets

`func (o *PagesDeploymentConfig) GetR2Buckets() map[string]PagesR2Binding`

GetR2Buckets returns the R2Buckets field if non-nil, zero value otherwise.

### GetR2BucketsOk

`func (o *PagesDeploymentConfig) GetR2BucketsOk() (*map[string]PagesR2Binding, bool)`

GetR2BucketsOk returns a tuple with the R2Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR2Buckets

`func (o *PagesDeploymentConfig) SetR2Buckets(v map[string]PagesR2Binding)`

SetR2Buckets sets R2Buckets field to given value.

### HasR2Buckets

`func (o *PagesDeploymentConfig) HasR2Buckets() bool`

HasR2Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


