# MarketplaceReceived

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corrected** | Pointer to **bool** | Corrected is true for a statement that corrects an earlier one. | [optional] 
**Furnished** | Pointer to **int64** | Furnished is when it was delivered, unix seconds. | [optional] 
**Id** | Pointer to **string** | ID is the form id the payer issued it under. | [optional] 
**Kind** | Pointer to **string** | Kind is 1099-NEC or 1099-MISC. | [optional] 
**Payer** | Pointer to **string** | Payer is the org that furnished it. | [optional] 
**Superseded** | Pointer to **bool** | Superseded is true for one a later correction replaced. | [optional] 
**Year** | Pointer to **int64** | Year is the calendar year of the payments. | [optional] 

## Methods

### NewMarketplaceReceived

`func NewMarketplaceReceived() *MarketplaceReceived`

NewMarketplaceReceived instantiates a new MarketplaceReceived object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceReceivedWithDefaults

`func NewMarketplaceReceivedWithDefaults() *MarketplaceReceived`

NewMarketplaceReceivedWithDefaults instantiates a new MarketplaceReceived object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrected

`func (o *MarketplaceReceived) GetCorrected() bool`

GetCorrected returns the Corrected field if non-nil, zero value otherwise.

### GetCorrectedOk

`func (o *MarketplaceReceived) GetCorrectedOk() (*bool, bool)`

GetCorrectedOk returns a tuple with the Corrected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrected

`func (o *MarketplaceReceived) SetCorrected(v bool)`

SetCorrected sets Corrected field to given value.

### HasCorrected

`func (o *MarketplaceReceived) HasCorrected() bool`

HasCorrected returns a boolean if a field has been set.

### GetFurnished

`func (o *MarketplaceReceived) GetFurnished() int64`

GetFurnished returns the Furnished field if non-nil, zero value otherwise.

### GetFurnishedOk

`func (o *MarketplaceReceived) GetFurnishedOk() (*int64, bool)`

GetFurnishedOk returns a tuple with the Furnished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurnished

`func (o *MarketplaceReceived) SetFurnished(v int64)`

SetFurnished sets Furnished field to given value.

### HasFurnished

`func (o *MarketplaceReceived) HasFurnished() bool`

HasFurnished returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceReceived) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceReceived) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceReceived) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceReceived) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *MarketplaceReceived) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplaceReceived) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplaceReceived) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplaceReceived) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPayer

`func (o *MarketplaceReceived) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *MarketplaceReceived) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *MarketplaceReceived) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *MarketplaceReceived) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetSuperseded

`func (o *MarketplaceReceived) GetSuperseded() bool`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *MarketplaceReceived) GetSupersededOk() (*bool, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *MarketplaceReceived) SetSuperseded(v bool)`

SetSuperseded sets Superseded field to given value.

### HasSuperseded

`func (o *MarketplaceReceived) HasSuperseded() bool`

HasSuperseded returns a boolean if a field has been set.

### GetYear

`func (o *MarketplaceReceived) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MarketplaceReceived) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MarketplaceReceived) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *MarketplaceReceived) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


