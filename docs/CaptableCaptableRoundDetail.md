# CaptableCaptableRoundDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investments** | Pointer to [**[]CaptableCaptableRoundInvestment**](CaptableCaptableRoundInvestment.md) | Investments is every investment into this round, oldest first. | [optional] 
**Round** | Pointer to [**CaptableCaptableRound**](CaptableCaptableRound.md) | Round is the round&#39;s own terms — name, type, valuation, target and status — as against the investments beside it. | [optional] 

## Methods

### NewCaptableCaptableRoundDetail

`func NewCaptableCaptableRoundDetail() *CaptableCaptableRoundDetail`

NewCaptableCaptableRoundDetail instantiates a new CaptableCaptableRoundDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableRoundDetailWithDefaults

`func NewCaptableCaptableRoundDetailWithDefaults() *CaptableCaptableRoundDetail`

NewCaptableCaptableRoundDetailWithDefaults instantiates a new CaptableCaptableRoundDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestments

`func (o *CaptableCaptableRoundDetail) GetInvestments() []CaptableCaptableRoundInvestment`

GetInvestments returns the Investments field if non-nil, zero value otherwise.

### GetInvestmentsOk

`func (o *CaptableCaptableRoundDetail) GetInvestmentsOk() (*[]CaptableCaptableRoundInvestment, bool)`

GetInvestmentsOk returns a tuple with the Investments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestments

`func (o *CaptableCaptableRoundDetail) SetInvestments(v []CaptableCaptableRoundInvestment)`

SetInvestments sets Investments field to given value.

### HasInvestments

`func (o *CaptableCaptableRoundDetail) HasInvestments() bool`

HasInvestments returns a boolean if a field has been set.

### GetRound

`func (o *CaptableCaptableRoundDetail) GetRound() CaptableCaptableRound`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *CaptableCaptableRoundDetail) GetRoundOk() (*CaptableCaptableRound, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *CaptableCaptableRoundDetail) SetRound(v CaptableCaptableRound)`

SetRound sets Round field to given value.

### HasRound

`func (o *CaptableCaptableRoundDetail) HasRound() bool`

HasRound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


