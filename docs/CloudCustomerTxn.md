# CloudCustomerTxn

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

### NewCloudCustomerTxn

`func NewCloudCustomerTxn() *CloudCustomerTxn`

NewCloudCustomerTxn instantiates a new CloudCustomerTxn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCustomerTxnWithDefaults

`func NewCloudCustomerTxnWithDefaults() *CloudCustomerTxn`

NewCloudCustomerTxnWithDefaults instantiates a new CloudCustomerTxn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *CloudCustomerTxn) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CloudCustomerTxn) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CloudCustomerTxn) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *CloudCustomerTxn) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudCustomerTxn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudCustomerTxn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudCustomerTxn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudCustomerTxn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *CloudCustomerTxn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCustomerTxn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCustomerTxn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCustomerTxn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotes

`func (o *CloudCustomerTxn) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CloudCustomerTxn) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CloudCustomerTxn) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CloudCustomerTxn) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTime

`func (o *CloudCustomerTxn) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CloudCustomerTxn) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CloudCustomerTxn) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CloudCustomerTxn) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *CloudCustomerTxn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCustomerTxn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCustomerTxn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCustomerTxn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


