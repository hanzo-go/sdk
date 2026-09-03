# CaptableInvestment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | Amount is the cash invested. | [optional] 
**Date** | Pointer to **string** | Date is the ISO date of the investment. | [optional] 
**Id** | Pointer to **string** | ID is the investment id. | [optional] 
**RoundId** | Pointer to **string** | RoundID is the round the cheque went into. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class shares were issued in, for a priced round. | [optional] 
**Shares** | Pointer to **int64** | Shares is how many shares the investment bought; 0 when the round issues no equity at the time of investment. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 

## Methods

### NewCaptableInvestment

`func NewCaptableInvestment() *CaptableInvestment`

NewCaptableInvestment instantiates a new CaptableInvestment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableInvestmentWithDefaults

`func NewCaptableInvestmentWithDefaults() *CaptableInvestment`

NewCaptableInvestmentWithDefaults instantiates a new CaptableInvestment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CaptableInvestment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CaptableInvestment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CaptableInvestment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CaptableInvestment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *CaptableInvestment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CaptableInvestment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CaptableInvestment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CaptableInvestment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *CaptableInvestment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableInvestment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableInvestment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableInvestment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoundId

`func (o *CaptableInvestment) GetRoundId() string`

GetRoundId returns the RoundId field if non-nil, zero value otherwise.

### GetRoundIdOk

`func (o *CaptableInvestment) GetRoundIdOk() (*string, bool)`

GetRoundIdOk returns a tuple with the RoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundId

`func (o *CaptableInvestment) SetRoundId(v string)`

SetRoundId sets RoundId field to given value.

### HasRoundId

`func (o *CaptableInvestment) HasRoundId() bool`

HasRoundId returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableInvestment) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableInvestment) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableInvestment) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableInvestment) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetShares

`func (o *CaptableInvestment) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CaptableInvestment) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CaptableInvestment) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CaptableInvestment) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableInvestment) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableInvestment) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableInvestment) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableInvestment) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableInvestment) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableInvestment) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableInvestment) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableInvestment) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


