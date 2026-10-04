# ToolPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** |  | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code, e.g. \&quot;USD\&quot;. Empty means USD. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the payout wallet ref the marketplace seller is paid at. | [optional] 

## Methods

### NewToolPrice

`func NewToolPrice() *ToolPrice`

NewToolPrice instantiates a new ToolPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolPriceWithDefaults

`func NewToolPriceWithDefaults() *ToolPrice`

NewToolPriceWithDefaults instantiates a new ToolPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ToolPrice) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ToolPrice) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ToolPrice) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ToolPrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ToolPrice) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ToolPrice) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *ToolPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ToolPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ToolPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ToolPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecipient

`func (o *ToolPrice) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ToolPrice) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ToolPrice) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ToolPrice) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


