# MintedKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AccessKey is the same value under its predecessor name, carried so callers written against the older field keep working. One value, two names. | [optional] 
**Key** | Pointer to **string** | Key is the credential, returned ONCE — a secret key is unreadable afterwards. | [optional] 
**Limit** | Pointer to **[]string** | Limit is what the minted key may reach, echoed back so the caller can see the narrowing took. Absent means unrestricted. | [optional] 
**Type** | Pointer to **string** | Type is the class of key that was minted. | [optional] 

## Methods

### NewMintedKey

`func NewMintedKey() *MintedKey`

NewMintedKey instantiates a new MintedKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintedKeyWithDefaults

`func NewMintedKeyWithDefaults() *MintedKey`

NewMintedKeyWithDefaults instantiates a new MintedKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *MintedKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *MintedKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *MintedKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *MintedKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetKey

`func (o *MintedKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MintedKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MintedKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MintedKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *MintedKey) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MintedKey) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MintedKey) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MintedKey) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetType

`func (o *MintedKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MintedKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MintedKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MintedKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


