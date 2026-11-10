# AccountApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**AccountKeyBudget**](AccountKeyBudget.md) | Budget is what the key may spend on models, in cents. Absent means no cap. | [optional] 
**Created** | Pointer to **string** | Created is when the key was minted (RFC 3339). | [optional] 
**Creator** | Pointer to **string** | Creator is the person the key speaks for, &#x60;&lt;org&gt;/&lt;user&gt;&#x60;: who created it. | [optional] 
**Expires** | Pointer to **string** | Expires is when the key stops working (RFC 3339). Absent means never. | [optional] 
**Id** | Pointer to **string** | ID addresses this key in PATCH and DELETE /v1/account/keys/{id}. | [optional] 
**Key** | Pointer to **string** | Key is the full credential. A publishable key always carries it; a secret key carries it once, in the answer to the POST that created it. | [optional] 
**Limit** | Pointer to **[]string** | Limit is what this key may reach, as &#x60;kind:name&#x60; entries — &#x60;model:zen5&#x60;, &#x60;project:acme&#x60;, &#x60;product:train&#x60; (read and write), &#x60;read:billing&#x60; (read only), &#x60;read:*&#x60; (a read-only key). Absent means the key reaches whatever its holder does. | [optional] 
**Name** | Pointer to **string** | Name is the key&#39;s label, chosen by the person who made it. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the head of the credential this key&#39;s holder presents — the sk- of a secret key — enough to tell which string a row is, never enough to use. Absent on a key minted before prefixes were recorded. | [optional] 
**Rate** | Pointer to **int64** | Rate is how many requests a minute the key may make. Absent means no limit of its own. | [optional] 
**Revoked** | Pointer to **string** |  | [optional] 
**Revoker** | Pointer to **string** | Revoker is who revoked the key, &#x60;&lt;org&gt;/&lt;user&gt;&#x60;, and Revoked when. | [optional] 
**Spend** | Pointer to [**AccountKeyBudget**](AccountKeyBudget.md) | Spend is what the key has spent on models, in cents rounded up. Absent when it could not be read just now. | [optional] 
**Status** | Pointer to **string** | Status is active, expired, revoked, or disabled (switched off in IAM). | [optional] 
**Type** | Pointer to **string** | Type is the key class: secret (sk-) or publishable (pk-). | [optional] 
**Used** | Pointer to **string** | Used is when the key was last used (RFC 3339). Absent means never. | [optional] 

## Methods

### NewAccountApiKey

`func NewAccountApiKey() *AccountApiKey`

NewAccountApiKey instantiates a new AccountApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountApiKeyWithDefaults

`func NewAccountApiKeyWithDefaults() *AccountApiKey`

NewAccountApiKeyWithDefaults instantiates a new AccountApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *AccountApiKey) GetBudget() AccountKeyBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AccountApiKey) GetBudgetOk() (*AccountKeyBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AccountApiKey) SetBudget(v AccountKeyBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AccountApiKey) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCreated

`func (o *AccountApiKey) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountApiKey) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountApiKey) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountApiKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *AccountApiKey) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AccountApiKey) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AccountApiKey) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AccountApiKey) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetExpires

`func (o *AccountApiKey) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AccountApiKey) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AccountApiKey) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AccountApiKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetId

`func (o *AccountApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AccountApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AccountApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AccountApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AccountApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *AccountApiKey) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountApiKey) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountApiKey) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccountApiKey) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *AccountApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *AccountApiKey) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AccountApiKey) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AccountApiKey) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AccountApiKey) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRate

`func (o *AccountApiKey) GetRate() int64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AccountApiKey) GetRateOk() (*int64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AccountApiKey) SetRate(v int64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AccountApiKey) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRevoked

`func (o *AccountApiKey) GetRevoked() string`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *AccountApiKey) GetRevokedOk() (*string, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *AccountApiKey) SetRevoked(v string)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *AccountApiKey) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetRevoker

`func (o *AccountApiKey) GetRevoker() string`

GetRevoker returns the Revoker field if non-nil, zero value otherwise.

### GetRevokerOk

`func (o *AccountApiKey) GetRevokerOk() (*string, bool)`

GetRevokerOk returns a tuple with the Revoker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoker

`func (o *AccountApiKey) SetRevoker(v string)`

SetRevoker sets Revoker field to given value.

### HasRevoker

`func (o *AccountApiKey) HasRevoker() bool`

HasRevoker returns a boolean if a field has been set.

### GetSpend

`func (o *AccountApiKey) GetSpend() AccountKeyBudget`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AccountApiKey) GetSpendOk() (*AccountKeyBudget, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AccountApiKey) SetSpend(v AccountKeyBudget)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AccountApiKey) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *AccountApiKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountApiKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountApiKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountApiKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AccountApiKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountApiKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountApiKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountApiKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsed

`func (o *AccountApiKey) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AccountApiKey) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AccountApiKey) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AccountApiKey) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


