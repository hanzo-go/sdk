# CloudApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the key last changed, as IAM records it. | [optional] 
**Key** | Pointer to **string** | Key is the FULL value, and is present for a publishable key only: it is public by construction and useless to its holder if it cannot be read back. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the recognizable, non-secret head of the key — enough to tell two keys apart, never enough to use one. | [optional] 
**Type** | Pointer to **string** | Type is the key class: secret (sk-) or publishable (pk-). | [optional] 

## Methods

### NewCloudApiKey

`func NewCloudApiKey() *CloudApiKey`

NewCloudApiKey instantiates a new CloudApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudApiKeyWithDefaults

`func NewCloudApiKeyWithDefaults() *CloudApiKey`

NewCloudApiKeyWithDefaults instantiates a new CloudApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKey

`func (o *CloudApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPrefix

`func (o *CloudApiKey) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudApiKey) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudApiKey) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CloudApiKey) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *CloudApiKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudApiKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudApiKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudApiKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


