# CaptableRoundDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investments** | Pointer to [**[]CaptableRoundInvestment**](CaptableRoundInvestment.md) | Investments is every investment into this round, oldest first. | [optional] 
**Round** | Pointer to [**CaptableRound**](CaptableRound.md) | Round is the round&#39;s own terms — name, type, valuation, target and status — as against the investments beside it. | [optional] 

## Methods

### NewCaptableRoundDetail

`func NewCaptableRoundDetail() *CaptableRoundDetail`

NewCaptableRoundDetail instantiates a new CaptableRoundDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableRoundDetailWithDefaults

`func NewCaptableRoundDetailWithDefaults() *CaptableRoundDetail`

NewCaptableRoundDetailWithDefaults instantiates a new CaptableRoundDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestments

`func (o *CaptableRoundDetail) GetInvestments() []CaptableRoundInvestment`

GetInvestments returns the Investments field if non-nil, zero value otherwise.

### GetInvestmentsOk

`func (o *CaptableRoundDetail) GetInvestmentsOk() (*[]CaptableRoundInvestment, bool)`

GetInvestmentsOk returns a tuple with the Investments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestments

`func (o *CaptableRoundDetail) SetInvestments(v []CaptableRoundInvestment)`

SetInvestments sets Investments field to given value.

### HasInvestments

`func (o *CaptableRoundDetail) HasInvestments() bool`

HasInvestments returns a boolean if a field has been set.

### GetRound

`func (o *CaptableRoundDetail) GetRound() CaptableRound`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *CaptableRoundDetail) GetRoundOk() (*CaptableRound, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *CaptableRoundDetail) SetRound(v CaptableRound)`

SetRound sets Round field to given value.

### HasRound

`func (o *CaptableRoundDetail) HasRound() bool`

HasRound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


