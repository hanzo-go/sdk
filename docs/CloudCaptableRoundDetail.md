# CloudCaptableRoundDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investments** | Pointer to [**[]CloudCaptableRoundInvestment**](CloudCaptableRoundInvestment.md) | Investments is every investment into this round, oldest first. | [optional] 
**Round** | Pointer to [**CloudCaptableRound**](CloudCaptableRound.md) | Round is the round itself. | [optional] 

## Methods

### NewCloudCaptableRoundDetail

`func NewCloudCaptableRoundDetail() *CloudCaptableRoundDetail`

NewCloudCaptableRoundDetail instantiates a new CloudCaptableRoundDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableRoundDetailWithDefaults

`func NewCloudCaptableRoundDetailWithDefaults() *CloudCaptableRoundDetail`

NewCloudCaptableRoundDetailWithDefaults instantiates a new CloudCaptableRoundDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestments

`func (o *CloudCaptableRoundDetail) GetInvestments() []CloudCaptableRoundInvestment`

GetInvestments returns the Investments field if non-nil, zero value otherwise.

### GetInvestmentsOk

`func (o *CloudCaptableRoundDetail) GetInvestmentsOk() (*[]CloudCaptableRoundInvestment, bool)`

GetInvestmentsOk returns a tuple with the Investments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestments

`func (o *CloudCaptableRoundDetail) SetInvestments(v []CloudCaptableRoundInvestment)`

SetInvestments sets Investments field to given value.

### HasInvestments

`func (o *CloudCaptableRoundDetail) HasInvestments() bool`

HasInvestments returns a boolean if a field has been set.

### GetRound

`func (o *CloudCaptableRoundDetail) GetRound() CloudCaptableRound`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *CloudCaptableRoundDetail) GetRoundOk() (*CloudCaptableRound, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *CloudCaptableRoundDetail) SetRound(v CloudCaptableRound)`

SetRound sets Round field to given value.

### HasRound

`func (o *CloudCaptableRoundDetail) HasRound() bool`

HasRound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


