# OnboardResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AccessKey is the identifier of the org-scoped credential provisioning minted with the organization. Present on a first run that actually minted one. | [optional] 
**AccessSecret** | Pointer to **string** | AccessSecret is that credential&#39;s confidential half, returned ONCE — on the response that mints it and never again. IAM keeps only its argon2id digest and blanks the plaintext, so this is the single moment it exists in a form its owner can read; a replay of the same provision re-reveals nothing. | [optional] 
**Additional** | Pointer to **bool** | Additional is true when the caller already had an organization and this one was created WITHOUT moving them into it — they reach it via the org switcher. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the organization&#39;s human name. | [optional] 
**Org** | Pointer to **string** | Org is the created organization&#39;s slug, which is what X-Org-Id carries. | [optional] 

## Methods

### NewOnboardResp

`func NewOnboardResp() *OnboardResp`

NewOnboardResp instantiates a new OnboardResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardRespWithDefaults

`func NewOnboardRespWithDefaults() *OnboardResp`

NewOnboardRespWithDefaults instantiates a new OnboardResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *OnboardResp) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OnboardResp) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OnboardResp) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *OnboardResp) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *OnboardResp) GetAccessSecret() string`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *OnboardResp) GetAccessSecretOk() (*string, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *OnboardResp) SetAccessSecret(v string)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *OnboardResp) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetAdditional

`func (o *OnboardResp) GetAdditional() bool`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *OnboardResp) GetAdditionalOk() (*bool, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *OnboardResp) SetAdditional(v bool)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *OnboardResp) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDisplayName

`func (o *OnboardResp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OnboardResp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OnboardResp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OnboardResp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOrg

`func (o *OnboardResp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *OnboardResp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *OnboardResp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *OnboardResp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


