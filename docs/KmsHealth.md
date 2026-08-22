# KmsHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the honest reason readiness is false: no in-process KMS client, or no master key. Absent when ready. | [optional] 
**Ready** | Pointer to **bool** | Ready is whether a secret operation would actually succeed right now. These are exactly the two states in which the secret operations refuse. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering, &#x60;kms&#x60;. | [optional] 
**Signing** | Pointer to **bool** | Signing reports whether signing keys are configured. Absent when there is no in-process client to ask. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;ok&#x60; or &#x60;degraded&#x60;, the one-word form of Ready. | [optional] 

## Methods

### NewKmsHealth

`func NewKmsHealth() *KmsHealth`

NewKmsHealth instantiates a new KmsHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsHealthWithDefaults

`func NewKmsHealthWithDefaults() *KmsHealth`

NewKmsHealthWithDefaults instantiates a new KmsHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *KmsHealth) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KmsHealth) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KmsHealth) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KmsHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReady

`func (o *KmsHealth) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *KmsHealth) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *KmsHealth) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *KmsHealth) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetService

`func (o *KmsHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *KmsHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *KmsHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *KmsHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSigning

`func (o *KmsHealth) GetSigning() bool`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *KmsHealth) GetSigningOk() (*bool, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *KmsHealth) SetSigning(v bool)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *KmsHealth) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetStatus

`func (o *KmsHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KmsHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KmsHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KmsHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


