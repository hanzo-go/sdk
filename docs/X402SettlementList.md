# X402SettlementList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settlements** | Pointer to [**[]X402Receipt**](X402Receipt.md) | Settlements is at most 1000 receipts. | [optional] 

## Methods

### NewX402SettlementList

`func NewX402SettlementList() *X402SettlementList`

NewX402SettlementList instantiates a new X402SettlementList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX402SettlementListWithDefaults

`func NewX402SettlementListWithDefaults() *X402SettlementList`

NewX402SettlementListWithDefaults instantiates a new X402SettlementList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlements

`func (o *X402SettlementList) GetSettlements() []X402Receipt`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *X402SettlementList) GetSettlementsOk() (*[]X402Receipt, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *X402SettlementList) SetSettlements(v []X402Receipt)`

SetSettlements sets Settlements field to given value.

### HasSettlements

`func (o *X402SettlementList) HasSettlements() bool`

HasSettlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


