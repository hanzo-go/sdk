# MarketReach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is &#x60;read&#x60;, &#x60;unconfigured&#x60;, &#x60;unreachable&#x60; or &#x60;refused&#x60;.  The four values are written out here because this document cannot carry an enum, so the description IS the contract a client reads. Spelling the Go constant names instead would name four identifiers no caller can see. | [optional] 
**Why** | Pointer to **string** | Why is the upstream&#39;s own words, on Unreachable and Refused, and empty otherwise. It is the upstream&#39;s and not ours: a failure reported without its reason sends the reader to look at the venue, which is the one place the fault is not. | [optional] 

## Methods

### NewMarketReach

`func NewMarketReach() *MarketReach`

NewMarketReach instantiates a new MarketReach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketReachWithDefaults

`func NewMarketReachWithDefaults() *MarketReach`

NewMarketReachWithDefaults instantiates a new MarketReach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketReach) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketReach) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketReach) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketReach) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetWhy

`func (o *MarketReach) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *MarketReach) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *MarketReach) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *MarketReach) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


