# CredentialIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID scopes the credential where the provider&#39;s Verify needs one. | [optional] 
**Label** | Pointer to **string** | Label names this connection; empty means \&quot;default\&quot;. | [optional] 
**Oauth** | Pointer to [**OauthBundleIn**](OauthBundleIn.md) | OAuth is a bundle the CLI already obtained through its own local PKCE flow. Present ⇒ the Adopt path; absent ⇒ the Token path. | [optional] 
**Provider** | Pointer to **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | [optional] 
**Token** | Pointer to **string** | Token is the customer-held credential for the Verify path. Read on STDIN by the CLI, never argv; never logged, echoed, or stored outside KMS. | [optional] 

## Methods

### NewCredentialIn

`func NewCredentialIn() *CredentialIn`

NewCredentialIn instantiates a new CredentialIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialInWithDefaults

`func NewCredentialInWithDefaults() *CredentialIn`

NewCredentialInWithDefaults instantiates a new CredentialIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CredentialIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CredentialIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CredentialIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CredentialIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLabel

`func (o *CredentialIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CredentialIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CredentialIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CredentialIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOauth

`func (o *CredentialIn) GetOauth() OauthBundleIn`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *CredentialIn) GetOauthOk() (*OauthBundleIn, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *CredentialIn) SetOauth(v OauthBundleIn)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *CredentialIn) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetProvider

`func (o *CredentialIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CredentialIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CredentialIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CredentialIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetToken

`func (o *CredentialIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CredentialIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CredentialIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CredentialIn) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


