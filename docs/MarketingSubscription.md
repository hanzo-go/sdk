# MarketingSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is \&quot;pending\&quot; from subscribe: a link to confirm is on its way, and the answer is the same whatever the address&#39;s standing. It is \&quot;confirmed\&quot; from confirm: the subscription is recorded. | [optional] 

## Methods

### NewMarketingSubscription

`func NewMarketingSubscription() *MarketingSubscription`

NewMarketingSubscription instantiates a new MarketingSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriptionWithDefaults

`func NewMarketingSubscriptionWithDefaults() *MarketingSubscription`

NewMarketingSubscriptionWithDefaults instantiates a new MarketingSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MarketingSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


