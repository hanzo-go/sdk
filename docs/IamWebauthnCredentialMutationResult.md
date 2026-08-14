# IamWebauthnCredentialMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **bool** |  | [optional] 
**WebauthnCredential** | Pointer to [**IamWebauthnCredential**](IamWebauthnCredential.md) |  | [optional] 

## Methods

### NewIamWebauthnCredentialMutationResult

`func NewIamWebauthnCredentialMutationResult() *IamWebauthnCredentialMutationResult`

NewIamWebauthnCredentialMutationResult instantiates a new IamWebauthnCredentialMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamWebauthnCredentialMutationResultWithDefaults

`func NewIamWebauthnCredentialMutationResultWithDefaults() *IamWebauthnCredentialMutationResult`

NewIamWebauthnCredentialMutationResultWithDefaults instantiates a new IamWebauthnCredentialMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *IamWebauthnCredentialMutationResult) GetAffected() bool`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *IamWebauthnCredentialMutationResult) GetAffectedOk() (*bool, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *IamWebauthnCredentialMutationResult) SetAffected(v bool)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *IamWebauthnCredentialMutationResult) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetWebauthnCredential

`func (o *IamWebauthnCredentialMutationResult) GetWebauthnCredential() IamWebauthnCredential`

GetWebauthnCredential returns the WebauthnCredential field if non-nil, zero value otherwise.

### GetWebauthnCredentialOk

`func (o *IamWebauthnCredentialMutationResult) GetWebauthnCredentialOk() (*IamWebauthnCredential, bool)`

GetWebauthnCredentialOk returns a tuple with the WebauthnCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnCredential

`func (o *IamWebauthnCredentialMutationResult) SetWebauthnCredential(v IamWebauthnCredential)`

SetWebauthnCredential sets WebauthnCredential field to given value.

### HasWebauthnCredential

`func (o *IamWebauthnCredentialMutationResult) HasWebauthnCredential() bool`

HasWebauthnCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


