# CloudCaptableInvestment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount is the cash invested. | [optional] 
**Date** | Pointer to **string** | Date is the ISO date of the investment. | [optional] 
**Id** | Pointer to **string** | ID is the investment id. | [optional] 
**RoundId** | Pointer to **string** | RoundID is the round the cheque went into. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class shares were issued in, for a priced round. | [optional] 
**Shares** | Pointer to **int32** | Shares is how many shares the investment bought; 0 when the round issues no equity at the time of investment. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 

## Methods

### NewCloudCaptableInvestment

`func NewCloudCaptableInvestment() *CloudCaptableInvestment`

NewCloudCaptableInvestment instantiates a new CloudCaptableInvestment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableInvestmentWithDefaults

`func NewCloudCaptableInvestmentWithDefaults() *CloudCaptableInvestment`

NewCloudCaptableInvestmentWithDefaults instantiates a new CloudCaptableInvestment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CloudCaptableInvestment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudCaptableInvestment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudCaptableInvestment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudCaptableInvestment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *CloudCaptableInvestment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CloudCaptableInvestment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CloudCaptableInvestment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CloudCaptableInvestment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableInvestment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableInvestment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableInvestment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableInvestment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoundId

`func (o *CloudCaptableInvestment) GetRoundId() string`

GetRoundId returns the RoundId field if non-nil, zero value otherwise.

### GetRoundIdOk

`func (o *CloudCaptableInvestment) GetRoundIdOk() (*string, bool)`

GetRoundIdOk returns a tuple with the RoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundId

`func (o *CloudCaptableInvestment) SetRoundId(v string)`

SetRoundId sets RoundId field to given value.

### HasRoundId

`func (o *CloudCaptableInvestment) HasRoundId() bool`

HasRoundId returns a boolean if a field has been set.

### GetShareClassId

`func (o *CloudCaptableInvestment) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CloudCaptableInvestment) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CloudCaptableInvestment) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CloudCaptableInvestment) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.

### GetShares

`func (o *CloudCaptableInvestment) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CloudCaptableInvestment) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CloudCaptableInvestment) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CloudCaptableInvestment) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableInvestment) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableInvestment) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableInvestment) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableInvestment) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CloudCaptableInvestment) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CloudCaptableInvestment) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CloudCaptableInvestment) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CloudCaptableInvestment) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


