# RateSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the affiliate whose direct rate moves, from the path. | [optional] 
**RateBps** | Pointer to **int32** | RateBps is the direct commission rate, in basis points of Hanzo&#39;s margin; capped so the whole L1+L2+L3 schedule never exceeds the margin. Body-only (&#x60;url:\&quot;-\&quot;&#x60;): a money parameter must never ride the URL into access logs. | [optional] 

## Methods

### NewRateSet

`func NewRateSet() *RateSet`

NewRateSet instantiates a new RateSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateSetWithDefaults

`func NewRateSetWithDefaults() *RateSet`

NewRateSetWithDefaults instantiates a new RateSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RateSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RateSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRateBps

`func (o *RateSet) GetRateBps() int32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *RateSet) GetRateBpsOk() (*int32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *RateSet) SetRateBps(v int32)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *RateSet) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


