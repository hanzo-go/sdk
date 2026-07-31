# CloudBalanceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** | cents, display sign | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudBalanceLine

`func NewCloudBalanceLine() *CloudBalanceLine`

NewCloudBalanceLine instantiates a new CloudBalanceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBalanceLineWithDefaults

`func NewCloudBalanceLineWithDefaults() *CloudBalanceLine`

NewCloudBalanceLineWithDefaults instantiates a new CloudBalanceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudBalanceLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudBalanceLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudBalanceLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudBalanceLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *CloudBalanceLine) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudBalanceLine) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudBalanceLine) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudBalanceLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetName

`func (o *CloudBalanceLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBalanceLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBalanceLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBalanceLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CloudBalanceLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudBalanceLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudBalanceLine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudBalanceLine) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


