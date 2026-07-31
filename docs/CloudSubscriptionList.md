# CloudSubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudSubscriptionView**](CloudSubscriptionView.md) | Data holds the repo&#39;s subscriptions. | [optional] 

## Methods

### NewCloudSubscriptionList

`func NewCloudSubscriptionList() *CloudSubscriptionList`

NewCloudSubscriptionList instantiates a new CloudSubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubscriptionListWithDefaults

`func NewCloudSubscriptionListWithDefaults() *CloudSubscriptionList`

NewCloudSubscriptionListWithDefaults instantiates a new CloudSubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSubscriptionList) GetData() []CloudSubscriptionView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSubscriptionList) GetDataOk() (*[]CloudSubscriptionView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSubscriptionList) SetData(v []CloudSubscriptionView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSubscriptionList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


