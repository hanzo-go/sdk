# CloudCaptableRound

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

### NewCloudCaptableRound

`func NewCloudCaptableRound() *CloudCaptableRound`

NewCloudCaptableRound instantiates a new CloudCaptableRound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableRoundWithDefaults

`func NewCloudCaptableRoundWithDefaults() *CloudCaptableRound`

NewCloudCaptableRoundWithDefaults instantiates a new CloudCaptableRound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *CloudCaptableRound) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CloudCaptableRound) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CloudCaptableRound) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CloudCaptableRound) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCaptableRound) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCaptableRound) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCaptableRound) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCaptableRound) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableRound) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableRound) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableRound) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableRound) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableRound) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableRound) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableRound) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableRound) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreMoneyValuation

`func (o *CloudCaptableRound) GetPreMoneyValuation() float32`

GetPreMoneyValuation returns the PreMoneyValuation field if non-nil, zero value otherwise.

### GetPreMoneyValuationOk

`func (o *CloudCaptableRound) GetPreMoneyValuationOk() (*float32, bool)`

GetPreMoneyValuationOk returns a tuple with the PreMoneyValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreMoneyValuation

`func (o *CloudCaptableRound) SetPreMoneyValuation(v float32)`

SetPreMoneyValuation sets PreMoneyValuation field to given value.

### HasPreMoneyValuation

`func (o *CloudCaptableRound) HasPreMoneyValuation() bool`

HasPreMoneyValuation returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CloudCaptableRound) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CloudCaptableRound) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CloudCaptableRound) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CloudCaptableRound) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetRaisedAmount

`func (o *CloudCaptableRound) GetRaisedAmount() float32`

GetRaisedAmount returns the RaisedAmount field if non-nil, zero value otherwise.

### GetRaisedAmountOk

`func (o *CloudCaptableRound) GetRaisedAmountOk() (*float32, bool)`

GetRaisedAmountOk returns a tuple with the RaisedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaisedAmount

`func (o *CloudCaptableRound) SetRaisedAmount(v float32)`

SetRaisedAmount sets RaisedAmount field to given value.

### HasRaisedAmount

`func (o *CloudCaptableRound) HasRaisedAmount() bool`

HasRaisedAmount returns a boolean if a field has been set.

### GetRoundType

`func (o *CloudCaptableRound) GetRoundType() string`

GetRoundType returns the RoundType field if non-nil, zero value otherwise.

### GetRoundTypeOk

`func (o *CloudCaptableRound) GetRoundTypeOk() (*string, bool)`

GetRoundTypeOk returns a tuple with the RoundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundType

`func (o *CloudCaptableRound) SetRoundType(v string)`

SetRoundType sets RoundType field to given value.

### HasRoundType

`func (o *CloudCaptableRound) HasRoundType() bool`

HasRoundType returns a boolean if a field has been set.

### GetShareClassId

`func (o *CloudCaptableRound) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CloudCaptableRound) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CloudCaptableRound) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CloudCaptableRound) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCaptableRound) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCaptableRound) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCaptableRound) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCaptableRound) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetAmount

`func (o *CloudCaptableRound) GetTargetAmount() float32`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *CloudCaptableRound) GetTargetAmountOk() (*float32, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *CloudCaptableRound) SetTargetAmount(v float32)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *CloudCaptableRound) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


