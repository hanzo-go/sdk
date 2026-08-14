# Funnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertedOrgs** | Pointer to **int32** | ConvertedOrgs is how many distinct referred orgs have produced positive commission at least once — a referral that actually spent. | [optional] 
**RatePct** | Pointer to **float32** | RatePct is convertedOrgs over referredOrgs as a PERCENTAGE, 0–100, and the one non-integer figure on this board. It is 0 when nothing has been referred yet, not undefined. | [optional] 
**ReferredOrgs** | Pointer to **int32** | ReferredOrgs is how many attribution edges exist fleet-wide — one per referred org, first-touch, so it is also the count of distinct referred orgs. | [optional] 

## Methods

### NewFunnel

`func NewFunnel() *Funnel`

NewFunnel instantiates a new Funnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelWithDefaults

`func NewFunnelWithDefaults() *Funnel`

NewFunnelWithDefaults instantiates a new Funnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertedOrgs

`func (o *Funnel) GetConvertedOrgs() int32`

GetConvertedOrgs returns the ConvertedOrgs field if non-nil, zero value otherwise.

### GetConvertedOrgsOk

`func (o *Funnel) GetConvertedOrgsOk() (*int32, bool)`

GetConvertedOrgsOk returns a tuple with the ConvertedOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedOrgs

`func (o *Funnel) SetConvertedOrgs(v int32)`

SetConvertedOrgs sets ConvertedOrgs field to given value.

### HasConvertedOrgs

`func (o *Funnel) HasConvertedOrgs() bool`

HasConvertedOrgs returns a boolean if a field has been set.

### GetRatePct

`func (o *Funnel) GetRatePct() float32`

GetRatePct returns the RatePct field if non-nil, zero value otherwise.

### GetRatePctOk

`func (o *Funnel) GetRatePctOk() (*float32, bool)`

GetRatePctOk returns a tuple with the RatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePct

`func (o *Funnel) SetRatePct(v float32)`

SetRatePct sets RatePct field to given value.

### HasRatePct

`func (o *Funnel) HasRatePct() bool`

HasRatePct returns a boolean if a field has been set.

### GetReferredOrgs

`func (o *Funnel) GetReferredOrgs() int32`

GetReferredOrgs returns the ReferredOrgs field if non-nil, zero value otherwise.

### GetReferredOrgsOk

`func (o *Funnel) GetReferredOrgsOk() (*int32, bool)`

GetReferredOrgsOk returns a tuple with the ReferredOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredOrgs

`func (o *Funnel) SetReferredOrgs(v int32)`

SetReferredOrgs sets ReferredOrgs field to given value.

### HasReferredOrgs

`func (o *Funnel) HasReferredOrgs() bool`

HasReferredOrgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


