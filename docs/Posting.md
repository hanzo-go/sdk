# Posting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPosting

`func NewPosting() *Posting`

NewPosting instantiates a new Posting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostingWithDefaults

`func NewPostingWithDefaults() *Posting`

NewPostingWithDefaults instantiates a new Posting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Posting) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Posting) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Posting) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Posting) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *Posting) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Posting) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Posting) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Posting) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *Posting) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *Posting) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


