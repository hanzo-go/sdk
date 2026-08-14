# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the key last changed, as IAM records it. | [optional] 
**Key** | Pointer to **string** | Key is the FULL value, and is present for a publishable key only: it is public by construction and useless to its holder if it cannot be read back. | [optional] 
**Limit** | Pointer to **[]string** | Limit is what this key may reach, as &#x60;kind:name&#x60; entries — &#x60;model:zen5&#x60;, &#x60;project:acme&#x60;, &#x60;product:commerce&#x60;. Absent means the key reaches whatever its holder does, which is what every key minted before limits existed does and must keep doing. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the recognizable, non-secret head of the key — enough to tell two keys apart, never enough to use one. | [optional] 
**Type** | Pointer to **string** | Type is the key class: secret (sk-) or publishable (pk-). | [optional] 

## Methods

### NewApiKey

`func NewApiKey() *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKey

`func (o *ApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *ApiKey) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ApiKey) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ApiKey) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ApiKey) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPrefix

`func (o *ApiKey) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKey) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKey) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiKey) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *ApiKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


