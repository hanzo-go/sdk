# Restarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the service that was restarted. | [optional] 
**Env** | Pointer to **string** | Env is that namespace&#39;s lifecycle env. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace its Deployment was patched in. | [optional] 
**Ok** | Pointer to **bool** | OK is always true — a failure is an error, not a false here. | [optional] 
**RestartedAt** | Pointer to **string** | RestartedAt is the timestamp stamped onto the pod template, RFC3339 UTC. | [optional] 

## Methods

### NewRestarted

`func NewRestarted() *Restarted`

NewRestarted instantiates a new Restarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartedWithDefaults

`func NewRestartedWithDefaults() *Restarted`

NewRestartedWithDefaults instantiates a new Restarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *Restarted) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Restarted) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Restarted) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *Restarted) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetEnv

`func (o *Restarted) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Restarted) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Restarted) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *Restarted) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNamespace

`func (o *Restarted) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Restarted) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Restarted) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Restarted) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOk

`func (o *Restarted) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *Restarted) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *Restarted) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *Restarted) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetRestartedAt

`func (o *Restarted) GetRestartedAt() string`

GetRestartedAt returns the RestartedAt field if non-nil, zero value otherwise.

### GetRestartedAtOk

`func (o *Restarted) GetRestartedAtOk() (*string, bool)`

GetRestartedAtOk returns a tuple with the RestartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartedAt

`func (o *Restarted) SetRestartedAt(v string)`

SetRestartedAt sets RestartedAt field to given value.

### HasRestartedAt

`func (o *Restarted) HasRestartedAt() bool`

HasRestartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


