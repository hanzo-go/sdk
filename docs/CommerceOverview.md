# CommerceOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aov** | Pointer to **float32** | AOV is average order value — Revenue/Orders, rounded to two places. Zero when there were no orders. | [optional] 
**Available** | Pointer to **bool** | Available is false when the product-event table could not be read — the lens is reported missing rather than as zeros that look like no sales. | [optional] 
**Orders** | Pointer to **int32** | Orders is how many order_completed events landed in the window. | [optional] 
**Reason** | Pointer to **string** | Reason says why the lens is unavailable. Omitted when it is available. | [optional] 
**Revenue** | Pointer to **float32** | Revenue is the total those orders carried, in the events&#39; own currency unit. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewCommerceOverview

`func NewCommerceOverview() *CommerceOverview`

NewCommerceOverview instantiates a new CommerceOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceOverviewWithDefaults

`func NewCommerceOverviewWithDefaults() *CommerceOverview`

NewCommerceOverviewWithDefaults instantiates a new CommerceOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAov

`func (o *CommerceOverview) GetAov() float32`

GetAov returns the Aov field if non-nil, zero value otherwise.

### GetAovOk

`func (o *CommerceOverview) GetAovOk() (*float32, bool)`

GetAovOk returns a tuple with the Aov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAov

`func (o *CommerceOverview) SetAov(v float32)`

SetAov sets Aov field to given value.

### HasAov

`func (o *CommerceOverview) HasAov() bool`

HasAov returns a boolean if a field has been set.

### GetAvailable

`func (o *CommerceOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CommerceOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CommerceOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CommerceOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetOrders

`func (o *CommerceOverview) GetOrders() int32`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *CommerceOverview) GetOrdersOk() (*int32, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *CommerceOverview) SetOrders(v int32)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *CommerceOverview) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetReason

`func (o *CommerceOverview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CommerceOverview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CommerceOverview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CommerceOverview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRevenue

`func (o *CommerceOverview) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CommerceOverview) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CommerceOverview) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CommerceOverview) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSource

`func (o *CommerceOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommerceOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommerceOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommerceOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


