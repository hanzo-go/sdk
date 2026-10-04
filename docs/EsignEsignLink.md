# EsignEsignLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address this link is meant for. | [optional] 
**RecipientId** | Pointer to **string** | RecipientID is the recipient the link identifies. | [optional] 
**Role** | Pointer to **string** | Role is their role — only a SIGNER or an APPROVER gets a link, because only they are asked to act. | [optional] 
**SigningPath** | Pointer to **string** | SigningPath is the tail of the address to send them, relative to wherever the signing page is served. | [optional] 
**Token** | Pointer to **string** | Token is the crypto-random signing capability. It is the entire credential, so treat it as a secret and give each one only to the recipient it names. | [optional] 

## Methods

### NewEsignEsignLink

`func NewEsignEsignLink() *EsignEsignLink`

NewEsignEsignLink instantiates a new EsignEsignLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignLinkWithDefaults

`func NewEsignEsignLinkWithDefaults() *EsignEsignLink`

NewEsignEsignLinkWithDefaults instantiates a new EsignEsignLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignEsignLink) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignEsignLink) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignEsignLink) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignEsignLink) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecipientId

`func (o *EsignEsignLink) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignEsignLink) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignEsignLink) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignEsignLink) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetRole

`func (o *EsignEsignLink) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignEsignLink) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignEsignLink) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignEsignLink) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSigningPath

`func (o *EsignEsignLink) GetSigningPath() string`

GetSigningPath returns the SigningPath field if non-nil, zero value otherwise.

### GetSigningPathOk

`func (o *EsignEsignLink) GetSigningPathOk() (*string, bool)`

GetSigningPathOk returns a tuple with the SigningPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPath

`func (o *EsignEsignLink) SetSigningPath(v string)`

SetSigningPath sets SigningPath field to given value.

### HasSigningPath

`func (o *EsignEsignLink) HasSigningPath() bool`

HasSigningPath returns a boolean if a field has been set.

### GetToken

`func (o *EsignEsignLink) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsignEsignLink) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsignEsignLink) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsignEsignLink) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


