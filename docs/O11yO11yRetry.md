# O11yO11yRetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **int32** | Delay is how long to wait before retrying, in nanoseconds. | [optional] 

## Methods

### NewO11yO11yRetry

`func NewO11yO11yRetry() *O11yO11yRetry`

NewO11yO11yRetry instantiates a new O11yO11yRetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRetryWithDefaults

`func NewO11yO11yRetryWithDefaults() *O11yO11yRetry`

NewO11yO11yRetryWithDefaults instantiates a new O11yO11yRetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *O11yO11yRetry) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *O11yO11yRetry) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *O11yO11yRetry) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *O11yO11yRetry) HasDelay() bool`

HasDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


