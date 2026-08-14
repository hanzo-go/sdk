# CaptableRoundInvestment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount is the cash invested. | [optional] 
**Comments** | Pointer to **string** | Comments is the note recorded with the cheque, if any. | [optional] 
**Date** | Pointer to **string** | Date is the ISO date of the investment. | [optional] 
**Id** | Pointer to **string** | ID is the investment id. | [optional] 
**Shares** | Pointer to **int32** | Shares is how many shares the investment bought; 0 when the round issues no equity at the time of investment. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 

## Methods

### NewCaptableRoundInvestment

`func NewCaptableRoundInvestment() *CaptableRoundInvestment`

NewCaptableRoundInvestment instantiates a new CaptableRoundInvestment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableRoundInvestmentWithDefaults

`func NewCaptableRoundInvestmentWithDefaults() *CaptableRoundInvestment`

NewCaptableRoundInvestmentWithDefaults instantiates a new CaptableRoundInvestment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CaptableRoundInvestment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CaptableRoundInvestment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CaptableRoundInvestment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CaptableRoundInvestment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetComments

`func (o *CaptableRoundInvestment) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CaptableRoundInvestment) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CaptableRoundInvestment) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CaptableRoundInvestment) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDate

`func (o *CaptableRoundInvestment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CaptableRoundInvestment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CaptableRoundInvestment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CaptableRoundInvestment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *CaptableRoundInvestment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableRoundInvestment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableRoundInvestment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableRoundInvestment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShares

`func (o *CaptableRoundInvestment) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CaptableRoundInvestment) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CaptableRoundInvestment) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CaptableRoundInvestment) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableRoundInvestment) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableRoundInvestment) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableRoundInvestment) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableRoundInvestment) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableRoundInvestment) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableRoundInvestment) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableRoundInvestment) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableRoundInvestment) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


