# IamMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to [**IamProvider**](IamProvider.md) |  | [optional] 

## Methods

### NewIamMutationResult

`func NewIamMutationResult() *IamMutationResult`

NewIamMutationResult instantiates a new IamMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamMutationResultWithDefaults

`func NewIamMutationResultWithDefaults() *IamMutationResult`

NewIamMutationResultWithDefaults instantiates a new IamMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *IamMutationResult) GetAffected() bool`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *IamMutationResult) GetAffectedOk() (*bool, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *IamMutationResult) SetAffected(v bool)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *IamMutationResult) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetProvider

`func (o *IamMutationResult) GetProvider() IamProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamMutationResult) GetProviderOk() (*IamProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamMutationResult) SetProvider(v IamProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamMutationResult) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


