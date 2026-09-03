# ChannelMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | ExternalID is the provider-side id of the execution the spend belongs to. Absent until the channel has launched. | [optional] 
**Kind** | Pointer to **string** | Kind is which channel this row is: paid, organic or email. It is also the row&#39;s identity — a campaign carries at most one channel per kind. | [optional] 
**Platform** | Pointer to **string** | Platform is the provider the spend was read from: meta, google, x, instagram, or the email provider. | [optional] 
**SpendCents** | Pointer to **int64** | SpendCents is what the provider itself reports this channel spent, in CENTS. 0 when the channel never launched, when no executor is wired for it, or when the read failed — SpendError tells the last case apart from a genuine zero. | [optional] 
**SpendError** | Pointer to **string** | SpendError is why this channel&#39;s spend could not be read (connector not connected, provider error), as one secret-free line. Present only on failure; the campaign total then simply omits this channel rather than failing. | [optional] 
**Status** | Pointer to **string** | Status is the channel&#39;s launch state on the campaign — pending, live, paused, failed or unavailable. Only a live channel is asked for its spend at all. | [optional] 

## Methods

### NewChannelMetric

`func NewChannelMetric() *ChannelMetric`

NewChannelMetric instantiates a new ChannelMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMetricWithDefaults

`func NewChannelMetricWithDefaults() *ChannelMetric`

NewChannelMetricWithDefaults instantiates a new ChannelMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ChannelMetric) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChannelMetric) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChannelMetric) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChannelMetric) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetKind

`func (o *ChannelMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChannelMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChannelMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ChannelMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *ChannelMetric) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ChannelMetric) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ChannelMetric) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ChannelMetric) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpendCents

`func (o *ChannelMetric) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *ChannelMetric) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *ChannelMetric) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *ChannelMetric) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetSpendError

`func (o *ChannelMetric) GetSpendError() string`

GetSpendError returns the SpendError field if non-nil, zero value otherwise.

### GetSpendErrorOk

`func (o *ChannelMetric) GetSpendErrorOk() (*string, bool)`

GetSpendErrorOk returns a tuple with the SpendError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendError

`func (o *ChannelMetric) SetSpendError(v string)`

SetSpendError sets SpendError field to given value.

### HasSpendError

`func (o *ChannelMetric) HasSpendError() bool`

HasSpendError returns a boolean if a field has been set.

### GetStatus

`func (o *ChannelMetric) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChannelMetric) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChannelMetric) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChannelMetric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


