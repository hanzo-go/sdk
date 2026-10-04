# EventCommerceOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aov** | Pointer to **float64** | AOV is average order value — Revenue/Orders, rounded to two places. Zero when there were no orders. | [optional] 
**Available** | Pointer to **bool** | Available is false when the product-event table could not be read — the lens is reported missing rather than as zeros that look like no sales. | [optional] 
**Orders** | Pointer to **int64** | Orders is how many order_completed events landed in the window. | [optional] 
**Reason** | Pointer to **string** | Reason says why the lens is unavailable. Omitted when it is available. | [optional] 
**Revenue** | Pointer to **float64** | Revenue is the total those orders carried, in the events&#39; own currency unit. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewEventCommerceOverview

`func NewEventCommerceOverview() *EventCommerceOverview`

NewEventCommerceOverview instantiates a new EventCommerceOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCommerceOverviewWithDefaults

`func NewEventCommerceOverviewWithDefaults() *EventCommerceOverview`

NewEventCommerceOverviewWithDefaults instantiates a new EventCommerceOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAov

`func (o *EventCommerceOverview) GetAov() float64`

GetAov returns the Aov field if non-nil, zero value otherwise.

### GetAovOk

`func (o *EventCommerceOverview) GetAovOk() (*float64, bool)`

GetAovOk returns a tuple with the Aov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAov

`func (o *EventCommerceOverview) SetAov(v float64)`

SetAov sets Aov field to given value.

### HasAov

`func (o *EventCommerceOverview) HasAov() bool`

HasAov returns a boolean if a field has been set.

### GetAvailable

`func (o *EventCommerceOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EventCommerceOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EventCommerceOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *EventCommerceOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetOrders

`func (o *EventCommerceOverview) GetOrders() int64`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *EventCommerceOverview) GetOrdersOk() (*int64, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *EventCommerceOverview) SetOrders(v int64)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *EventCommerceOverview) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetReason

`func (o *EventCommerceOverview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventCommerceOverview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventCommerceOverview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventCommerceOverview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRevenue

`func (o *EventCommerceOverview) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *EventCommerceOverview) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *EventCommerceOverview) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *EventCommerceOverview) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSource

`func (o *EventCommerceOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventCommerceOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventCommerceOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventCommerceOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


