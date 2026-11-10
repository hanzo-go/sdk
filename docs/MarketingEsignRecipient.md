# MarketingEsignRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningOrder** | Pointer to **int64** | SigningOrder is the signer&#39;s turn: 1 is the investor, 2 the countersigner. | [optional] 
**SigningStatus** | Pointer to **string** | SigningStatus is NOT_SIGNED, SIGNED or REJECTED. | [optional] 

## Methods

### NewMarketingEsignRecipient

`func NewMarketingEsignRecipient() *MarketingEsignRecipient`

NewMarketingEsignRecipient instantiates a new MarketingEsignRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEsignRecipientWithDefaults

`func NewMarketingEsignRecipientWithDefaults() *MarketingEsignRecipient`

NewMarketingEsignRecipientWithDefaults instantiates a new MarketingEsignRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningOrder

`func (o *MarketingEsignRecipient) GetSigningOrder() int64`

GetSigningOrder returns the SigningOrder field if non-nil, zero value otherwise.

### GetSigningOrderOk

`func (o *MarketingEsignRecipient) GetSigningOrderOk() (*int64, bool)`

GetSigningOrderOk returns a tuple with the SigningOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOrder

`func (o *MarketingEsignRecipient) SetSigningOrder(v int64)`

SetSigningOrder sets SigningOrder field to given value.

### HasSigningOrder

`func (o *MarketingEsignRecipient) HasSigningOrder() bool`

HasSigningOrder returns a boolean if a field has been set.

### GetSigningStatus

`func (o *MarketingEsignRecipient) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *MarketingEsignRecipient) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *MarketingEsignRecipient) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *MarketingEsignRecipient) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


