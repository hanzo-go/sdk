# CloudSweepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** | Period is the accrual period as YYYY-MM. Empty takes the current UTC month. | [optional] 
**RevenueCents** | Pointer to **int32** | RevenueCents is the net platform revenue measured for the period, in minor units. Must be &gt;&#x3D; 0. | [optional] 

## Methods

### NewCloudSweepRequest

`func NewCloudSweepRequest() *CloudSweepRequest`

NewCloudSweepRequest instantiates a new CloudSweepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSweepRequestWithDefaults

`func NewCloudSweepRequestWithDefaults() *CloudSweepRequest`

NewCloudSweepRequestWithDefaults instantiates a new CloudSweepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *CloudSweepRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudSweepRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudSweepRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CloudSweepRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRevenueCents

`func (o *CloudSweepRequest) GetRevenueCents() int32`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *CloudSweepRequest) GetRevenueCentsOk() (*int32, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *CloudSweepRequest) SetRevenueCents(v int32)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *CloudSweepRequest) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


