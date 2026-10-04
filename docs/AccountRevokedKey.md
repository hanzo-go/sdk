# AccountRevokedKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** | OK is true when the key was revoked. A failure is an error status, never a false here. | [optional] 
**Type** | Pointer to **string** | Type is the key class that was revoked, resolved — so a caller that named nothing can see it revoked the secret key. | [optional] 

## Methods

### NewAccountRevokedKey

`func NewAccountRevokedKey() *AccountRevokedKey`

NewAccountRevokedKey instantiates a new AccountRevokedKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRevokedKeyWithDefaults

`func NewAccountRevokedKeyWithDefaults() *AccountRevokedKey`

NewAccountRevokedKeyWithDefaults instantiates a new AccountRevokedKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *AccountRevokedKey) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AccountRevokedKey) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AccountRevokedKey) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *AccountRevokedKey) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetType

`func (o *AccountRevokedKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountRevokedKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountRevokedKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountRevokedKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


