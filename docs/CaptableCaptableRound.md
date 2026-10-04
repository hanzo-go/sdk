# CaptableCaptableRound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseDate** | Pointer to **string** | CloseDate is the ISO date the round closed, once it has. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the round was recorded, in unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the round id. | [optional] 
**Name** | Pointer to **string** | Name is the round name, e.g. \&quot;Series A\&quot;. | [optional] 
**PreMoneyValuation** | Pointer to **float64** | PreMoneyValuation is the pre-money valuation, for a priced round. | [optional] 
**PricePerShare** | Pointer to **float64** | PricePerShare is the price per share, for a priced round. | [optional] 
**RaisedAmount** | Pointer to **float64** | RaisedAmount is how much has been invested so far. | [optional] 
**RoundType** | Pointer to **string** | RoundType is PRICED, SAFE or CONVERTIBLE_NOTE. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class a priced round issues into. | [optional] 
**Status** | Pointer to **string** | Status is OPEN or CLOSED. | [optional] 
**TargetAmount** | Pointer to **float64** | TargetAmount is how much the round set out to raise. | [optional] 

## Methods

### NewCaptableCaptableRound

`func NewCaptableCaptableRound() *CaptableCaptableRound`

NewCaptableCaptableRound instantiates a new CaptableCaptableRound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableRoundWithDefaults

`func NewCaptableCaptableRoundWithDefaults() *CaptableCaptableRound`

NewCaptableCaptableRoundWithDefaults instantiates a new CaptableCaptableRound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *CaptableCaptableRound) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CaptableCaptableRound) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CaptableCaptableRound) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CaptableCaptableRound) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaptableCaptableRound) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptableCaptableRound) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptableCaptableRound) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptableCaptableRound) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CaptableCaptableRound) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableRound) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableRound) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableRound) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CaptableCaptableRound) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableCaptableRound) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableCaptableRound) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableCaptableRound) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreMoneyValuation

`func (o *CaptableCaptableRound) GetPreMoneyValuation() float64`

GetPreMoneyValuation returns the PreMoneyValuation field if non-nil, zero value otherwise.

### GetPreMoneyValuationOk

`func (o *CaptableCaptableRound) GetPreMoneyValuationOk() (*float64, bool)`

GetPreMoneyValuationOk returns a tuple with the PreMoneyValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreMoneyValuation

`func (o *CaptableCaptableRound) SetPreMoneyValuation(v float64)`

SetPreMoneyValuation sets PreMoneyValuation field to given value.

### HasPreMoneyValuation

`func (o *CaptableCaptableRound) HasPreMoneyValuation() bool`

HasPreMoneyValuation returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CaptableCaptableRound) GetPricePerShare() float64`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableCaptableRound) GetPricePerShareOk() (*float64, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableCaptableRound) SetPricePerShare(v float64)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableCaptableRound) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetRaisedAmount

`func (o *CaptableCaptableRound) GetRaisedAmount() float64`

GetRaisedAmount returns the RaisedAmount field if non-nil, zero value otherwise.

### GetRaisedAmountOk

`func (o *CaptableCaptableRound) GetRaisedAmountOk() (*float64, bool)`

GetRaisedAmountOk returns a tuple with the RaisedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisedAmount

`func (o *CaptableCaptableRound) SetRaisedAmount(v float64)`

SetRaisedAmount sets RaisedAmount field to given value.

### HasRaisedAmount

`func (o *CaptableCaptableRound) HasRaisedAmount() bool`

HasRaisedAmount returns a boolean if a field has been set.

### GetRoundType

`func (o *CaptableCaptableRound) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *CaptableCaptableRound) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *CaptableCaptableRound) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *CaptableCaptableRound) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableCaptableRound) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableCaptableRound) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableCaptableRound) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableCaptableRound) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableCaptableRound) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableCaptableRound) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableCaptableRound) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableCaptableRound) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetAmount

`func (o *CaptableCaptableRound) GetTargetAmount() float64`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *CaptableCaptableRound) GetTargetAmountOk() (*float64, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *CaptableCaptableRound) SetTargetAmount(v float64)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *CaptableCaptableRound) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


