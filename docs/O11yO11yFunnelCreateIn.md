# O11yO11yFunnelCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelName** | Pointer to **string** | Name is the funnel&#39;s name. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp is when the funnel was created, as a millisecond epoch. Zero takes the runtime&#39;s own clock. | [optional] 

## Methods

### NewO11yO11yFunnelCreateIn

`func NewO11yO11yFunnelCreateIn() *O11yO11yFunnelCreateIn`

NewO11yO11yFunnelCreateIn instantiates a new O11yO11yFunnelCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFunnelCreateInWithDefaults

`func NewO11yO11yFunnelCreateInWithDefaults() *O11yO11yFunnelCreateIn`

NewO11yO11yFunnelCreateInWithDefaults instantiates a new O11yO11yFunnelCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelName

`func (o *O11yO11yFunnelCreateIn) GetFunnelName() string`

GetFunnelName returns the FunnelName field if non-nil, zero value otherwise.

### GetFunnelNameOk

`func (o *O11yO11yFunnelCreateIn) GetFunnelNameOk() (*string, bool)`

GetFunnelNameOk returns a tuple with the FunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelName

`func (o *O11yO11yFunnelCreateIn) SetFunnelName(v string)`

SetFunnelName sets FunnelName field to given value.

### HasFunnelName

`func (o *O11yO11yFunnelCreateIn) HasFunnelName() bool`

HasFunnelName returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yFunnelCreateIn) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yFunnelCreateIn) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yFunnelCreateIn) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yFunnelCreateIn) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


