# EsignEsignInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the invitation is for, lower-cased. | [optional] 
**Id** | Pointer to **string** | ID is the new recipient&#39;s id, which is what a field is placed against. | [optional] 
**Name** | Pointer to **string** | Name is the recipient&#39;s display name, empty when none was given. | [optional] 
**Role** | Pointer to **string** | Role is the role they were recorded with — SIGNER unless another was asked for. | [optional] 
**Token** | Pointer to **string** | Token is the crypto-random signing capability for this recipient. It is the entire credential their signing endpoint accepts, so treat it as a secret and hand it only to them: the signing link is built from it. | [optional] 

## Methods

### NewEsignEsignInvite

`func NewEsignEsignInvite() *EsignEsignInvite`

NewEsignEsignInvite instantiates a new EsignEsignInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignInviteWithDefaults

`func NewEsignEsignInviteWithDefaults() *EsignEsignInvite`

NewEsignEsignInviteWithDefaults instantiates a new EsignEsignInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignEsignInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignEsignInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignEsignInvite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignEsignInvite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EsignEsignInvite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignEsignInvite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignEsignInvite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignEsignInvite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EsignEsignInvite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsignEsignInvite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsignEsignInvite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsignEsignInvite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *EsignEsignInvite) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignEsignInvite) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignEsignInvite) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignEsignInvite) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *EsignEsignInvite) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsignEsignInvite) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsignEsignInvite) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsignEsignInvite) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


