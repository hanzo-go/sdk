# IamTokenMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to [**IamToken**](IamToken.md) |  | [optional] 

## Methods

### NewIamTokenMutation

`func NewIamTokenMutation() *IamTokenMutation`

NewIamTokenMutation instantiates a new IamTokenMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTokenMutationWithDefaults

`func NewIamTokenMutationWithDefaults() *IamTokenMutation`

NewIamTokenMutationWithDefaults instantiates a new IamTokenMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *IamTokenMutation) GetAffected() bool`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *IamTokenMutation) GetAffectedOk() (*bool, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *IamTokenMutation) SetAffected(v bool)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *IamTokenMutation) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetToken

`func (o *IamTokenMutation) GetToken() IamToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IamTokenMutation) GetTokenOk() (*IamToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IamTokenMutation) SetToken(v IamToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *IamTokenMutation) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


