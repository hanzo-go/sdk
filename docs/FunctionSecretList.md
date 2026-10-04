# FunctionSecretList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]FunctionSecretView**](FunctionSecretView.md) | Secrets is one row per distinct (namespace, name) a function mounts. Values are NEVER read or returned. | [optional] 

## Methods

### NewFunctionSecretList

`func NewFunctionSecretList() *FunctionSecretList`

NewFunctionSecretList instantiates a new FunctionSecretList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSecretListWithDefaults

`func NewFunctionSecretListWithDefaults() *FunctionSecretList`

NewFunctionSecretListWithDefaults instantiates a new FunctionSecretList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *FunctionSecretList) GetSecrets() []FunctionSecretView`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FunctionSecretList) GetSecretsOk() (*[]FunctionSecretView, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FunctionSecretList) SetSecrets(v []FunctionSecretView)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *FunctionSecretList) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


