# BillingCapVerdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | **bool** | Whether the request is permitted. | 
**Reason** | [**BillingCapReason**](BillingCapReason.md) |  | 
**CapCents** | Pointer to **int64** | The tightest violated cap&#39;s ceiling in USD cents (0 when allowed). | [optional] 
**SpentCents** | Pointer to **int64** | Period spend for the scope in USD cents at decision time. | [optional] 
**WarnPct** | Pointer to **int32** | Utilization percent of the most-utilized covering cap when at/over its soft threshold (0 otherwise, and 0 on a deny). | [optional] 

## Methods

### NewBillingCapVerdict

`func NewBillingCapVerdict(allow bool, reason BillingCapReason, ) *BillingCapVerdict`

NewBillingCapVerdict instantiates a new BillingCapVerdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCapVerdictWithDefaults

`func NewBillingCapVerdictWithDefaults() *BillingCapVerdict`

NewBillingCapVerdictWithDefaults instantiates a new BillingCapVerdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *BillingCapVerdict) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *BillingCapVerdict) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *BillingCapVerdict) SetAllow(v bool)`

SetAllow sets Allow field to given value.


### GetReason

`func (o *BillingCapVerdict) GetReason() BillingCapReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BillingCapVerdict) GetReasonOk() (*BillingCapReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BillingCapVerdict) SetReason(v BillingCapReason)`

SetReason sets Reason field to given value.


### GetCapCents

`func (o *BillingCapVerdict) GetCapCents() int64`

GetCapCents returns the CapCents field if non-nil, zero value otherwise.

### GetCapCentsOk

`func (o *BillingCapVerdict) GetCapCentsOk() (*int64, bool)`

GetCapCentsOk returns a tuple with the CapCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapCents

`func (o *BillingCapVerdict) SetCapCents(v int64)`

SetCapCents sets CapCents field to given value.

### HasCapCents

`func (o *BillingCapVerdict) HasCapCents() bool`

HasCapCents returns a boolean if a field has been set.

### GetSpentCents

`func (o *BillingCapVerdict) GetSpentCents() int64`

GetSpentCents returns the SpentCents field if non-nil, zero value otherwise.

### GetSpentCentsOk

`func (o *BillingCapVerdict) GetSpentCentsOk() (*int64, bool)`

GetSpentCentsOk returns a tuple with the SpentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentCents

`func (o *BillingCapVerdict) SetSpentCents(v int64)`

SetSpentCents sets SpentCents field to given value.

### HasSpentCents

`func (o *BillingCapVerdict) HasSpentCents() bool`

HasSpentCents returns a boolean if a field has been set.

### GetWarnPct

`func (o *BillingCapVerdict) GetWarnPct() int32`

GetWarnPct returns the WarnPct field if non-nil, zero value otherwise.

### GetWarnPctOk

`func (o *BillingCapVerdict) GetWarnPctOk() (*int32, bool)`

GetWarnPctOk returns a tuple with the WarnPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnPct

`func (o *BillingCapVerdict) SetWarnPct(v int32)`

SetWarnPct sets WarnPct field to given value.

### HasWarnPct

`func (o *BillingCapVerdict) HasWarnPct() bool`

HasWarnPct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


