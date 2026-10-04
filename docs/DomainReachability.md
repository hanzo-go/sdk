# DomainReachability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Configured is whether the wholesale credentials are present at all. | [optional] 
**Env** | Pointer to **string** | Env is the registrar environment. It is the fact that decides whether money moves: only \&quot;prod\&quot; reaches the live, billable registrar — anything else, including unset, is the sandbox. | [optional] 
**Error** | Pointer to **string** | Error is the blocker, so an operator reads it instead of guessing at it. | [optional] 
**Reachable** | Pointer to **bool** | Reachable is whether the registrar accepted those credentials on a live call made while the caller waited. | [optional] 
**Registrar** | Pointer to **string** | Registrar names the wholesale registrar behind it. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when a live call succeeded, else \&quot;degraded\&quot;. | [optional] 

## Methods

### NewDomainReachability

`func NewDomainReachability() *DomainReachability`

NewDomainReachability instantiates a new DomainReachability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainReachabilityWithDefaults

`func NewDomainReachabilityWithDefaults() *DomainReachability`

NewDomainReachabilityWithDefaults instantiates a new DomainReachability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *DomainReachability) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *DomainReachability) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *DomainReachability) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *DomainReachability) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetEnv

`func (o *DomainReachability) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *DomainReachability) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *DomainReachability) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *DomainReachability) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetError

`func (o *DomainReachability) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DomainReachability) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DomainReachability) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DomainReachability) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReachable

`func (o *DomainReachability) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *DomainReachability) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *DomainReachability) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *DomainReachability) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetRegistrar

`func (o *DomainReachability) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *DomainReachability) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *DomainReachability) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *DomainReachability) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetService

`func (o *DomainReachability) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DomainReachability) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DomainReachability) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DomainReachability) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DomainReachability) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainReachability) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainReachability) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainReachability) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


