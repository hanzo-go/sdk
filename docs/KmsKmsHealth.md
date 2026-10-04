# KmsKmsHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforcing** | Pointer to **bool** | Enforcing is whether a secret value the declaration in force does not grant is refused. False while it says &#x60;enforce: false&#x60; and while none has loaded; every such value is then logged and audited as would-deny. | [optional] 
**Error** | Pointer to **string** | Error is the honest reason readiness is false: no in-process KMS client, or no master key. Absent when ready. | [optional] 
**Policy** | Pointer to **string** | Policy is the sha256 of the grant declaration in force, so a probe can prove which one is live. Absent when none has ever loaded, which measures: the rule each caller had before still decides. | [optional] 
**Ready** | Pointer to **bool** | Ready is whether a secret operation would actually succeed right now. These are exactly the two states in which the secret operations refuse. | [optional] 
**Refused** | Pointer to **string** | Refused is what the MOUNTED declaration reads as — its sha256, or \&quot;absent\&quot; or \&quot;unreadable\&quot; — when it is not the one in force: a document cloud refused to parse, or a mount that went away. The declaration in force stays, so a bad edit changes nothing, and this is where it shows. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering, &#x60;kms&#x60;. | [optional] 
**Signing** | Pointer to **bool** | Signing reports whether signing keys are configured. Absent when there is no in-process client to ask. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;ok&#x60; or &#x60;degraded&#x60;, the one-word form of Ready. | [optional] 

## Methods

### NewKmsKmsHealth

`func NewKmsKmsHealth() *KmsKmsHealth`

NewKmsKmsHealth instantiates a new KmsKmsHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsHealthWithDefaults

`func NewKmsKmsHealthWithDefaults() *KmsKmsHealth`

NewKmsKmsHealthWithDefaults instantiates a new KmsKmsHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforcing

`func (o *KmsKmsHealth) GetEnforcing() bool`

GetEnforcing returns the Enforcing field if non-nil, zero value otherwise.

### GetEnforcingOk

`func (o *KmsKmsHealth) GetEnforcingOk() (*bool, bool)`

GetEnforcingOk returns a tuple with the Enforcing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcing

`func (o *KmsKmsHealth) SetEnforcing(v bool)`

SetEnforcing sets Enforcing field to given value.

### HasEnforcing

`func (o *KmsKmsHealth) HasEnforcing() bool`

HasEnforcing returns a boolean if a field has been set.

### GetError

`func (o *KmsKmsHealth) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KmsKmsHealth) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KmsKmsHealth) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KmsKmsHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPolicy

`func (o *KmsKmsHealth) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KmsKmsHealth) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KmsKmsHealth) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KmsKmsHealth) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetReady

`func (o *KmsKmsHealth) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *KmsKmsHealth) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *KmsKmsHealth) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *KmsKmsHealth) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRefused

`func (o *KmsKmsHealth) GetRefused() string`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *KmsKmsHealth) GetRefusedOk() (*string, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *KmsKmsHealth) SetRefused(v string)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *KmsKmsHealth) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetService

`func (o *KmsKmsHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *KmsKmsHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *KmsKmsHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *KmsKmsHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSigning

`func (o *KmsKmsHealth) GetSigning() bool`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *KmsKmsHealth) GetSigningOk() (*bool, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *KmsKmsHealth) SetSigning(v bool)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *KmsKmsHealth) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetStatus

`func (o *KmsKmsHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KmsKmsHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KmsKmsHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KmsKmsHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


