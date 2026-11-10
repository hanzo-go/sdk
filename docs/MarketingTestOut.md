# MarketingTestOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refused** | Pointer to **map[string]string** | Refused are the addresses the suppression list held, with why. | [optional] 
**Sent** | Pointer to **[]string** | Sent are the addresses it went to. | [optional] 

## Methods

### NewMarketingTestOut

`func NewMarketingTestOut() *MarketingTestOut`

NewMarketingTestOut instantiates a new MarketingTestOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingTestOutWithDefaults

`func NewMarketingTestOutWithDefaults() *MarketingTestOut`

NewMarketingTestOutWithDefaults instantiates a new MarketingTestOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefused

`func (o *MarketingTestOut) GetRefused() map[string]string`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *MarketingTestOut) GetRefusedOk() (*map[string]string, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *MarketingTestOut) SetRefused(v map[string]string)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *MarketingTestOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetSent

`func (o *MarketingTestOut) GetSent() []string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MarketingTestOut) GetSentOk() (*[]string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MarketingTestOut) SetSent(v []string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MarketingTestOut) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


