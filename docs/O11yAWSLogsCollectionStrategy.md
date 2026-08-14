# O11yAWSLogsCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]O11yAWSCloudWatchLogsSubscription**](O11yAWSCloudWatchLogsSubscription.md) |  | [optional] 

## Methods

### NewO11yAWSLogsCollectionStrategy

`func NewO11yAWSLogsCollectionStrategy() *O11yAWSLogsCollectionStrategy`

NewO11yAWSLogsCollectionStrategy instantiates a new O11yAWSLogsCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSLogsCollectionStrategyWithDefaults

`func NewO11yAWSLogsCollectionStrategyWithDefaults() *O11yAWSLogsCollectionStrategy`

NewO11yAWSLogsCollectionStrategyWithDefaults instantiates a new O11yAWSLogsCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *O11yAWSLogsCollectionStrategy) GetSubscriptions() []O11yAWSCloudWatchLogsSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *O11yAWSLogsCollectionStrategy) GetSubscriptionsOk() (*[]O11yAWSCloudWatchLogsSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *O11yAWSLogsCollectionStrategy) SetSubscriptions(v []O11yAWSCloudWatchLogsSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *O11yAWSLogsCollectionStrategy) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


