# CloudflarePagesDeploymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityDate** | Pointer to **string** | CompatibilityDate pins which Workers runtime behaviour the functions run under, as a date (\&quot;2024-01-01\&quot;). It is a pin, not a version: the runtime keeps that date&#39;s semantics for code deployed against it. | [optional] 
**CompatibilityFlags** | Pointer to **[]string** | CompatibilityFlags turn individual runtime behaviours on or off ahead of, or behind, the date above (\&quot;nodejs_compat\&quot;). | [optional] 
**D1Databases** | Pointer to [**map[string]CloudflarePagesD1Binding**](CloudflarePagesD1Binding.md) | D1Databases binds D1 databases in, keyed by binding name. | [optional] 
**EnvVars** | Pointer to [**map[string]CloudflarePagesEnvVar**](CloudflarePagesEnvVar.md) | EnvVars are the environment variables the functions see, KEYED BY VARIABLE NAME. The key is the name; the value carries the value and whether it is a secret. | [optional] 
**KvNamespaces** | Pointer to [**map[string]CloudflarePagesKVBinding**](CloudflarePagesKVBinding.md) | KVNamespaces binds KV namespaces into the functions, KEYED BY THE BINDING NAME the code reads (&#x60;env.SESSIONS&#x60;). Same shape for the two below. | [optional] 
**R2Buckets** | Pointer to [**map[string]CloudflarePagesR2Binding**](CloudflarePagesR2Binding.md) | R2Buckets binds R2 buckets in, keyed by binding name. | [optional] 

## Methods

### NewCloudflarePagesDeploymentConfig

`func NewCloudflarePagesDeploymentConfig() *CloudflarePagesDeploymentConfig`

NewCloudflarePagesDeploymentConfig instantiates a new CloudflarePagesDeploymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesDeploymentConfigWithDefaults

`func NewCloudflarePagesDeploymentConfigWithDefaults() *CloudflarePagesDeploymentConfig`

NewCloudflarePagesDeploymentConfigWithDefaults instantiates a new CloudflarePagesDeploymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityDate

`func (o *CloudflarePagesDeploymentConfig) GetCompatibilityDate() string`

GetCompatibilityDate returns the CompatibilityDate field if non-nil, zero value otherwise.

### GetCompatibilityDateOk

`func (o *CloudflarePagesDeploymentConfig) GetCompatibilityDateOk() (*string, bool)`

GetCompatibilityDateOk returns a tuple with the CompatibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityDate

`func (o *CloudflarePagesDeploymentConfig) SetCompatibilityDate(v string)`

SetCompatibilityDate sets CompatibilityDate field to given value.

### HasCompatibilityDate

`func (o *CloudflarePagesDeploymentConfig) HasCompatibilityDate() bool`

HasCompatibilityDate returns a boolean if a field has been set.

### GetCompatibilityFlags

`func (o *CloudflarePagesDeploymentConfig) GetCompatibilityFlags() []string`

GetCompatibilityFlags returns the CompatibilityFlags field if non-nil, zero value otherwise.

### GetCompatibilityFlagsOk

`func (o *CloudflarePagesDeploymentConfig) GetCompatibilityFlagsOk() (*[]string, bool)`

GetCompatibilityFlagsOk returns a tuple with the CompatibilityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityFlags

`func (o *CloudflarePagesDeploymentConfig) SetCompatibilityFlags(v []string)`

SetCompatibilityFlags sets CompatibilityFlags field to given value.

### HasCompatibilityFlags

`func (o *CloudflarePagesDeploymentConfig) HasCompatibilityFlags() bool`

HasCompatibilityFlags returns a boolean if a field has been set.

### GetD1Databases

`func (o *CloudflarePagesDeploymentConfig) GetD1Databases() map[string]CloudflarePagesD1Binding`

GetD1Databases returns the D1Databases field if non-nil, zero value otherwise.

### GetD1DatabasesOk

`func (o *CloudflarePagesDeploymentConfig) GetD1DatabasesOk() (*map[string]CloudflarePagesD1Binding, bool)`

GetD1DatabasesOk returns a tuple with the D1Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD1Databases

`func (o *CloudflarePagesDeploymentConfig) SetD1Databases(v map[string]CloudflarePagesD1Binding)`

SetD1Databases sets D1Databases field to given value.

### HasD1Databases

`func (o *CloudflarePagesDeploymentConfig) HasD1Databases() bool`

HasD1Databases returns a boolean if a field has been set.

### GetEnvVars

`func (o *CloudflarePagesDeploymentConfig) GetEnvVars() map[string]CloudflarePagesEnvVar`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *CloudflarePagesDeploymentConfig) GetEnvVarsOk() (*map[string]CloudflarePagesEnvVar, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *CloudflarePagesDeploymentConfig) SetEnvVars(v map[string]CloudflarePagesEnvVar)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *CloudflarePagesDeploymentConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetKvNamespaces

`func (o *CloudflarePagesDeploymentConfig) GetKvNamespaces() map[string]CloudflarePagesKVBinding`

GetKvNamespaces returns the KvNamespaces field if non-nil, zero value otherwise.

### GetKvNamespacesOk

`func (o *CloudflarePagesDeploymentConfig) GetKvNamespacesOk() (*map[string]CloudflarePagesKVBinding, bool)`

GetKvNamespacesOk returns a tuple with the KvNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvNamespaces

`func (o *CloudflarePagesDeploymentConfig) SetKvNamespaces(v map[string]CloudflarePagesKVBinding)`

SetKvNamespaces sets KvNamespaces field to given value.

### HasKvNamespaces

`func (o *CloudflarePagesDeploymentConfig) HasKvNamespaces() bool`

HasKvNamespaces returns a boolean if a field has been set.

### GetR2Buckets

`func (o *CloudflarePagesDeploymentConfig) GetR2Buckets() map[string]CloudflarePagesR2Binding`

GetR2Buckets returns the R2Buckets field if non-nil, zero value otherwise.

### GetR2BucketsOk

`func (o *CloudflarePagesDeploymentConfig) GetR2BucketsOk() (*map[string]CloudflarePagesR2Binding, bool)`

GetR2BucketsOk returns a tuple with the R2Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR2Buckets

`func (o *CloudflarePagesDeploymentConfig) SetR2Buckets(v map[string]CloudflarePagesR2Binding)`

SetR2Buckets sets R2Buckets field to given value.

### HasR2Buckets

`func (o *CloudflarePagesDeploymentConfig) HasR2Buckets() bool`

HasR2Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


