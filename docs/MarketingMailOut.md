# MarketingMailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it left, RFC 3339 UTC; empty when nothing left. | [optional] 
**Sent** | Pointer to **bool** | Sent is false when the address opted out of the org&#39;s mail, and nothing left. | [optional] 

## Methods

### NewMarketingMailOut

`func NewMarketingMailOut() *MarketingMailOut`

NewMarketingMailOut instantiates a new MarketingMailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingMailOutWithDefaults

`func NewMarketingMailOutWithDefaults() *MarketingMailOut`

NewMarketingMailOutWithDefaults instantiates a new MarketingMailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketingMailOut) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketingMailOut) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketingMailOut) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketingMailOut) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetSent

`func (o *MarketingMailOut) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MarketingMailOut) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MarketingMailOut) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MarketingMailOut) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


