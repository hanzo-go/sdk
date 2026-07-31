# CloudCaptableRoundInvestment

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

### NewCloudCaptableRoundInvestment

`func NewCloudCaptableRoundInvestment() *CloudCaptableRoundInvestment`

NewCloudCaptableRoundInvestment instantiates a new CloudCaptableRoundInvestment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableRoundInvestmentWithDefaults

`func NewCloudCaptableRoundInvestmentWithDefaults() *CloudCaptableRoundInvestment`

NewCloudCaptableRoundInvestmentWithDefaults instantiates a new CloudCaptableRoundInvestment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CloudCaptableRoundInvestment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudCaptableRoundInvestment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudCaptableRoundInvestment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudCaptableRoundInvestment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetComments

`func (o *CloudCaptableRoundInvestment) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CloudCaptableRoundInvestment) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CloudCaptableRoundInvestment) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CloudCaptableRoundInvestment) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDate

`func (o *CloudCaptableRoundInvestment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CloudCaptableRoundInvestment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CloudCaptableRoundInvestment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CloudCaptableRoundInvestment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableRoundInvestment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableRoundInvestment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableRoundInvestment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableRoundInvestment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShares

`func (o *CloudCaptableRoundInvestment) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CloudCaptableRoundInvestment) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CloudCaptableRoundInvestment) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CloudCaptableRoundInvestment) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableRoundInvestment) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableRoundInvestment) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableRoundInvestment) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableRoundInvestment) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CloudCaptableRoundInvestment) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CloudCaptableRoundInvestment) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CloudCaptableRoundInvestment) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CloudCaptableRoundInvestment) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


