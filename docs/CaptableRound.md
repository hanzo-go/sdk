# CaptableRound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseDate** | Pointer to **string** | CloseDate is the ISO date the round closed, once it has. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the round was recorded, in unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the round id. | [optional] 
**Name** | Pointer to **string** | Name is the round name, e.g. \&quot;Series A\&quot;. | [optional] 
**PreMoneyValuation** | Pointer to **float32** | PreMoneyValuation is the pre-money valuation, for a priced round. | [optional] 
**PricePerShare** | Pointer to **float32** | PricePerShare is the price per share, for a priced round. | [optional] 
**RaisedAmount** | Pointer to **float32** | RaisedAmount is how much has been invested so far. | [optional] 
**RoundType** | Pointer to **string** | RoundType is PRICED, SAFE or CONVERTIBLE_NOTE. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class a priced round issues into. | [optional] 
**Status** | Pointer to **string** | Status is OPEN or CLOSED. | [optional] 
**TargetAmount** | Pointer to **float32** | TargetAmount is how much the round set out to raise. | [optional] 

## Methods

### NewCaptableRound

`func NewCaptableRound() *CaptableRound`

NewCaptableRound instantiates a new CaptableRound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableRoundWithDefaults

`func NewCaptableRoundWithDefaults() *CaptableRound`

NewCaptableRoundWithDefaults instantiates a new CaptableRound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *CaptableRound) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CaptableRound) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CaptableRound) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CaptableRound) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaptableRound) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptableRound) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptableRound) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptableRound) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CaptableRound) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableRound) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableRound) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableRound) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CaptableRound) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableRound) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableRound) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableRound) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreMoneyValuation

`func (o *CaptableRound) GetPreMoneyValuation() float32`

GetPreMoneyValuation returns the PreMoneyValuation field if non-nil, zero value otherwise.

### GetPreMoneyValuationOk

`func (o *CaptableRound) GetPreMoneyValuationOk() (*float32, bool)`

GetPreMoneyValuationOk returns a tuple with the PreMoneyValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreMoneyValuation

`func (o *CaptableRound) SetPreMoneyValuation(v float32)`

SetPreMoneyValuation sets PreMoneyValuation field to given value.

### HasPreMoneyValuation

`func (o *CaptableRound) HasPreMoneyValuation() bool`

HasPreMoneyValuation returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CaptableRound) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableRound) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableRound) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableRound) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetRaisedAmount

`func (o *CaptableRound) GetRaisedAmount() float32`

GetRaisedAmount returns the RaisedAmount field if non-nil, zero value otherwise.

### GetRaisedAmountOk

`func (o *CaptableRound) GetRaisedAmountOk() (*float32, bool)`

GetRaisedAmountOk returns a tuple with the RaisedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisedAmount

`func (o *CaptableRound) SetRaisedAmount(v float32)`

SetRaisedAmount sets RaisedAmount field to given value.

### HasRaisedAmount

`func (o *CaptableRound) HasRaisedAmount() bool`

HasRaisedAmount returns a boolean if a field has been set.

### GetRoundType

`func (o *CaptableRound) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *CaptableRound) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *CaptableRound) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *CaptableRound) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableRound) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableRound) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableRound) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableRound) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableRound) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableRound) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableRound) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableRound) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetAmount

`func (o *CaptableRound) GetTargetAmount() float32`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *CaptableRound) GetTargetAmountOk() (*float32, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *CaptableRound) SetTargetAmount(v float32)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *CaptableRound) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


