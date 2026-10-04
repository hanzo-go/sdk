# AccountOnboardResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AccessKey is the identifier of the org-scoped credential provisioning minted with the organization. Present on a first run that actually minted one. | [optional] 
**AccessSecret** | Pointer to **string** | AccessSecret is that credential&#39;s confidential half, returned ONCE — on the response that mints it and never again. IAM keeps only its argon2id digest and blanks the plaintext, so this is the single moment it exists in a form its owner can read; a replay of the same provision re-reveals nothing. | [optional] 
**Additional** | Pointer to **bool** | Additional is true when the caller already had an organization and this one was created WITHOUT moving them into it — they reach it via the org switcher. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the organization&#39;s human name. | [optional] 
**Org** | Pointer to **string** | Org is the created organization&#39;s slug, which is what X-Org-Id carries. | [optional] 

## Methods

### NewAccountOnboardResp

`func NewAccountOnboardResp() *AccountOnboardResp`

NewAccountOnboardResp instantiates a new AccountOnboardResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOnboardRespWithDefaults

`func NewAccountOnboardRespWithDefaults() *AccountOnboardResp`

NewAccountOnboardRespWithDefaults instantiates a new AccountOnboardResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AccountOnboardResp) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AccountOnboardResp) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AccountOnboardResp) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AccountOnboardResp) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *AccountOnboardResp) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *AccountOnboardResp) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *AccountOnboardResp) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *AccountOnboardResp) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetAdditional

`func (o *AccountOnboardResp) GetAdditional() bool`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *AccountOnboardResp) GetAdditionalOk() (*bool, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *AccountOnboardResp) SetAdditional(v bool)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *AccountOnboardResp) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccountOnboardResp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountOnboardResp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountOnboardResp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccountOnboardResp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOrg

`func (o *AccountOnboardResp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AccountOnboardResp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AccountOnboardResp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AccountOnboardResp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


