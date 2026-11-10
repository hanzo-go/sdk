# MarketingSuppressionCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason is what the matching entry was recorded with: \&quot;one-click unsubscribe\&quot;, a bounce, a complaint, an admin&#39;s note. Absent when not suppressed. | [optional] 
**Scope** | Pointer to **string** | Scope is the list that matched: address (the org&#39;s, for the address), domain (the org&#39;s, for the address&#39;s company) or platform (the blocklist, which every org&#39;s gate reads). Absent when not suppressed. | [optional] 
**Suppressed** | Pointer to **bool** | Suppressed is true when the send gate refuses the address on the channel. | [optional] 

## Methods

### NewMarketingSuppressionCheck

`func NewMarketingSuppressionCheck() *MarketingSuppressionCheck`

NewMarketingSuppressionCheck instantiates a new MarketingSuppressionCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSuppressionCheckWithDefaults

`func NewMarketingSuppressionCheckWithDefaults() *MarketingSuppressionCheck`

NewMarketingSuppressionCheckWithDefaults instantiates a new MarketingSuppressionCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *MarketingSuppressionCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketingSuppressionCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketingSuppressionCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketingSuppressionCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScope

`func (o *MarketingSuppressionCheck) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MarketingSuppressionCheck) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MarketingSuppressionCheck) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MarketingSuppressionCheck) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSuppressed

`func (o *MarketingSuppressionCheck) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *MarketingSuppressionCheck) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *MarketingSuppressionCheck) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *MarketingSuppressionCheck) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


