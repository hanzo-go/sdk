# AuthzAddPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **bool** | True if the rule was newly added; false if it already existed. | [optional] 

## Methods

### NewAuthzAddPolicyResponse

`func NewAuthzAddPolicyResponse() *AuthzAddPolicyResponse`

NewAuthzAddPolicyResponse instantiates a new AuthzAddPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzAddPolicyResponseWithDefaults

`func NewAuthzAddPolicyResponseWithDefaults() *AuthzAddPolicyResponse`

NewAuthzAddPolicyResponseWithDefaults instantiates a new AuthzAddPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *AuthzAddPolicyResponse) GetAdded() bool`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *AuthzAddPolicyResponse) GetAddedOk() (*bool, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *AuthzAddPolicyResponse) SetAdded(v bool)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *AuthzAddPolicyResponse) HasAdded() bool`

HasAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


