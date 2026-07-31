# CloudMintedKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AccessKey is the same value under its predecessor name, carried so callers written against the older field keep working. One value, two names. | [optional] 
**Key** | Pointer to **string** | Key is the credential, returned ONCE — a secret key is unreadable afterwards. | [optional] 
**Type** | Pointer to **string** | Type is the class of key that was minted. | [optional] 

## Methods

### NewCloudMintedKey

`func NewCloudMintedKey() *CloudMintedKey`

NewCloudMintedKey instantiates a new CloudMintedKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMintedKeyWithDefaults

`func NewCloudMintedKeyWithDefaults() *CloudMintedKey`

NewCloudMintedKeyWithDefaults instantiates a new CloudMintedKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *CloudMintedKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CloudMintedKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CloudMintedKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CloudMintedKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetKey

`func (o *CloudMintedKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudMintedKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudMintedKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudMintedKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *CloudMintedKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudMintedKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudMintedKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudMintedKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


