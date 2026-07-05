# StreamListBrokers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brokers** | Pointer to [**[]StreamBrokerInfo**](StreamBrokerInfo.md) |  | [optional] 

## Methods

### NewStreamListBrokers200Response

`func NewStreamListBrokers200Response() *StreamListBrokers200Response`

NewStreamListBrokers200Response instantiates a new StreamListBrokers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamListBrokers200ResponseWithDefaults

`func NewStreamListBrokers200ResponseWithDefaults() *StreamListBrokers200Response`

NewStreamListBrokers200ResponseWithDefaults instantiates a new StreamListBrokers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokers

`func (o *StreamListBrokers200Response) GetBrokers() []StreamBrokerInfo`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *StreamListBrokers200Response) GetBrokersOk() (*[]StreamBrokerInfo, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *StreamListBrokers200Response) SetBrokers(v []StreamBrokerInfo)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *StreamListBrokers200Response) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


