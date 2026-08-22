# EsignLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address this link is meant for. | [optional] 
**RecipientId** | Pointer to **string** | RecipientID is the recipient the link identifies. | [optional] 
**Role** | Pointer to **string** | Role is their role — only a SIGNER or an APPROVER gets a link, because only they are asked to act. | [optional] 
**SigningPath** | Pointer to **string** | SigningPath is the tail of the address to send them, relative to wherever the signing page is served. | [optional] 
**Token** | Pointer to **string** | Token is the crypto-random signing capability. It is the entire credential, so treat it as a secret and give each one only to the recipient it names. | [optional] 

## Methods

### NewEsignLink

`func NewEsignLink() *EsignLink`

NewEsignLink instantiates a new EsignLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignLinkWithDefaults

`func NewEsignLinkWithDefaults() *EsignLink`

NewEsignLinkWithDefaults instantiates a new EsignLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignLink) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignLink) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignLink) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignLink) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecipientId

`func (o *EsignLink) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignLink) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignLink) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignLink) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetRole

`func (o *EsignLink) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignLink) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignLink) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignLink) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSigningPath

`func (o *EsignLink) GetSigningPath() string`

GetSigningPath returns the SigningPath field if non-nil, zero value otherwise.

### GetSigningPathOk

`func (o *EsignLink) GetSigningPathOk() (*string, bool)`

GetSigningPathOk returns a tuple with the SigningPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPath

`func (o *EsignLink) SetSigningPath(v string)`

SetSigningPath sets SigningPath field to given value.

### HasSigningPath

`func (o *EsignLink) HasSigningPath() bool`

HasSigningPath returns a boolean if a field has been set.

### GetToken

`func (o *EsignLink) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsignLink) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsignLink) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsignLink) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


