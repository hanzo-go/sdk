# CreditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewCreditEntry

`func NewCreditEntry() *CreditEntry`

NewCreditEntry instantiates a new CreditEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditEntryWithDefaults

`func NewCreditEntryWithDefaults() *CreditEntry`

NewCreditEntryWithDefaults instantiates a new CreditEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CreditEntry) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CreditEntry) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CreditEntry) SetAvailable(v int64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CreditEntry) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditEntry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditEntry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditEntry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreditEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


