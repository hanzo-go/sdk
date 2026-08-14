# ProviderCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnCents** | Pointer to **int32** |  | [optional] 
**GrantCents** | Pointer to **int32** |  | [optional] 
**HasCredit** | Pointer to **bool** |  | [optional] 
**IsPaidOnly** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RemainingCents** | Pointer to **int32** |  | [optional] 
**RunwayDays** | Pointer to **float32** | nil when burn is 0 / unknown (never a fabricated infinity) | [optional] 

## Methods

### NewProviderCredit

`func NewProviderCredit() *ProviderCredit`

NewProviderCredit instantiates a new ProviderCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCreditWithDefaults

`func NewProviderCreditWithDefaults() *ProviderCredit`

NewProviderCreditWithDefaults instantiates a new ProviderCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnCents

`func (o *ProviderCredit) GetBurnCents() int32`

GetBurnCents returns the BurnCents field if non-nil, zero value otherwise.

### GetBurnCentsOk

`func (o *ProviderCredit) GetBurnCentsOk() (*int32, bool)`

GetBurnCentsOk returns a tuple with the BurnCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnCents

`func (o *ProviderCredit) SetBurnCents(v int32)`

SetBurnCents sets BurnCents field to given value.

### HasBurnCents

`func (o *ProviderCredit) HasBurnCents() bool`

HasBurnCents returns a boolean if a field has been set.

### GetGrantCents

`func (o *ProviderCredit) GetGrantCents() int32`

GetGrantCents returns the GrantCents field if non-nil, zero value otherwise.

### GetGrantCentsOk

`func (o *ProviderCredit) GetGrantCentsOk() (*int32, bool)`

GetGrantCentsOk returns a tuple with the GrantCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantCents

`func (o *ProviderCredit) SetGrantCents(v int32)`

SetGrantCents sets GrantCents field to given value.

### HasGrantCents

`func (o *ProviderCredit) HasGrantCents() bool`

HasGrantCents returns a boolean if a field has been set.

### GetHasCredit

`func (o *ProviderCredit) GetHasCredit() bool`

GetHasCredit returns the HasCredit field if non-nil, zero value otherwise.

### GetHasCreditOk

`func (o *ProviderCredit) GetHasCreditOk() (*bool, bool)`

GetHasCreditOk returns a tuple with the HasCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCredit

`func (o *ProviderCredit) SetHasCredit(v bool)`

SetHasCredit sets HasCredit field to given value.

### HasHasCredit

`func (o *ProviderCredit) HasHasCredit() bool`

HasHasCredit returns a boolean if a field has been set.

### GetIsPaidOnly

`func (o *ProviderCredit) GetIsPaidOnly() bool`

GetIsPaidOnly returns the IsPaidOnly field if non-nil, zero value otherwise.

### GetIsPaidOnlyOk

`func (o *ProviderCredit) GetIsPaidOnlyOk() (*bool, bool)`

GetIsPaidOnlyOk returns a tuple with the IsPaidOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaidOnly

`func (o *ProviderCredit) SetIsPaidOnly(v bool)`

SetIsPaidOnly sets IsPaidOnly field to given value.

### HasIsPaidOnly

`func (o *ProviderCredit) HasIsPaidOnly() bool`

HasIsPaidOnly returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderCredit) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderCredit) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderCredit) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderCredit) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRemainingCents

`func (o *ProviderCredit) GetRemainingCents() int32`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *ProviderCredit) GetRemainingCentsOk() (*int32, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *ProviderCredit) SetRemainingCents(v int32)`

SetRemainingCents sets RemainingCents field to given value.

### HasRemainingCents

`func (o *ProviderCredit) HasRemainingCents() bool`

HasRemainingCents returns a boolean if a field has been set.

### GetRunwayDays

`func (o *ProviderCredit) GetRunwayDays() float32`

GetRunwayDays returns the RunwayDays field if non-nil, zero value otherwise.

### GetRunwayDaysOk

`func (o *ProviderCredit) GetRunwayDaysOk() (*float32, bool)`

GetRunwayDaysOk returns a tuple with the RunwayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunwayDays

`func (o *ProviderCredit) SetRunwayDays(v float32)`

SetRunwayDays sets RunwayDays field to given value.

### HasRunwayDays

`func (o *ProviderCredit) HasRunwayDays() bool`

HasRunwayDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


