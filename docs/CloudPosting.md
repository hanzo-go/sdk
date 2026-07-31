# CloudPosting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCloudPosting

`func NewCloudPosting() *CloudPosting`

NewCloudPosting instantiates a new CloudPosting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPostingWithDefaults

`func NewCloudPostingWithDefaults() *CloudPosting`

NewCloudPostingWithDefaults instantiates a new CloudPosting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudPosting) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudPosting) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudPosting) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudPosting) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *CloudPosting) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudPosting) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudPosting) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudPosting) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *CloudPosting) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *CloudPosting) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


