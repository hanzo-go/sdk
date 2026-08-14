# CustomerTxn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | \&quot;deposit\&quot; (credit) | \&quot;withdraw\&quot; (usage) | [optional] 

## Methods

### NewCustomerTxn

`func NewCustomerTxn() *CustomerTxn`

NewCustomerTxn instantiates a new CustomerTxn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerTxnWithDefaults

`func NewCustomerTxnWithDefaults() *CustomerTxn`

NewCustomerTxnWithDefaults instantiates a new CustomerTxn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *CustomerTxn) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CustomerTxn) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CustomerTxn) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *CustomerTxn) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomerTxn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerTxn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerTxn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerTxn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *CustomerTxn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerTxn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerTxn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerTxn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotes

`func (o *CustomerTxn) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CustomerTxn) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CustomerTxn) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CustomerTxn) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTime

`func (o *CustomerTxn) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CustomerTxn) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CustomerTxn) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CustomerTxn) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *CustomerTxn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerTxn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerTxn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerTxn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


