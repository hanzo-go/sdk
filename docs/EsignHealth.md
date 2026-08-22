# EsignHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service names the subsystem that answered, so a probe reading several looks the same on each. | [optional] 
**Status** | Pointer to **string** | Status is ok whenever the subsystem is mounted. It is never anything else: this route is registered before the document host is built, so it is reachability and not a promise that documents can be stored. | [optional] 

## Methods

### NewEsignHealth

`func NewEsignHealth() *EsignHealth`

NewEsignHealth instantiates a new EsignHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignHealthWithDefaults

`func NewEsignHealthWithDefaults() *EsignHealth`

NewEsignHealthWithDefaults instantiates a new EsignHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *EsignHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EsignHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EsignHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *EsignHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *EsignHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


