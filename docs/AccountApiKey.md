# AccountApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the key last changed, as IAM records it. | [optional] 
**Key** | Pointer to **string** | Key is the FULL value, and is present for a publishable key only: it is public by construction and useless to its holder if it cannot be read back. | [optional] 
**Limit** | Pointer to **[]string** | Limit is what this key may reach, as &#x60;kind:name&#x60; entries — &#x60;model:zen5&#x60;, &#x60;project:acme&#x60;, &#x60;product:commerce&#x60;. Absent means the key reaches whatever its holder does, which is what every key minted before limits existed does and must keep doing. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the recognizable, non-secret head of the key — enough to tell two keys apart, never enough to use one. | [optional] 
**Type** | Pointer to **string** | Type is the key class: secret (sk-) or publishable (pk-). | [optional] 

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

### GetCreatedAt

`func (o *AccountApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


