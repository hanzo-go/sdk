# Funnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available separates \&quot;this org has no traffic\&quot; from \&quot;we could not ask\&quot;. False means the warehouse was unreachable or the org has emitted nothing at all, and every count below is then a placeholder zero rather than a measurement — a caller must read this before reading any of them. | [optional] 
**Orders** | Pointer to **int64** | Orders counts completed orders in the window — purchases, not carts started. | [optional] 
**Pageviews** | Pointer to **int64** | Pageviews counts page events in the window, one per view rather than per person, so a single visitor reading ten pages counts ten. | [optional] 
**Revenue** | Pointer to **float64** | Revenue is the sum of the amounts those orders reported, in whatever currency the beacon stamped on them (major units, e.g. 49.5 for $49.50) — NOT cents, and not converted to a single currency. Contrast revenueCents on the profile, which is the money of record. | [optional] 
**Signups** | Pointer to **int64** | Signups counts completed signups in the window, the step where an anonymous visitor becomes somebody with an account. | [optional] 
**Visitors** | Pointer to **int64** | Visitors is the number of DISTINCT people seen in the window, counted by the beacon&#39;s distinct id — so it is unique visitors, not sessions and not views. | [optional] 
**WindowDays** | Pointer to **int64** | WindowDays is the length of the trailing window every count covers, so a reader knows whether 40 signups is a month or a day. | [optional] 

## Methods

### NewFunnel

`func NewFunnel() *Funnel`

NewFunnel instantiates a new Funnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelWithDefaults

`func NewFunnelWithDefaults() *Funnel`

NewFunnelWithDefaults instantiates a new Funnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Funnel) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Funnel) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Funnel) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Funnel) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetOrders

`func (o *Funnel) GetOrders() int64`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Funnel) GetOrdersOk() (*int64, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Funnel) SetOrders(v int64)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Funnel) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetPageviews

`func (o *Funnel) GetPageviews() int64`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *Funnel) GetPageviewsOk() (*int64, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *Funnel) SetPageviews(v int64)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *Funnel) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetRevenue

`func (o *Funnel) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *Funnel) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *Funnel) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *Funnel) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSignups

`func (o *Funnel) GetSignups() int64`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *Funnel) GetSignupsOk() (*int64, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *Funnel) SetSignups(v int64)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *Funnel) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetVisitors

`func (o *Funnel) GetVisitors() int64`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *Funnel) GetVisitorsOk() (*int64, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *Funnel) SetVisitors(v int64)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *Funnel) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.

### GetWindowDays

`func (o *Funnel) GetWindowDays() int64`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *Funnel) GetWindowDaysOk() (*int64, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *Funnel) SetWindowDays(v int64)`

SetWindowDays sets WindowDays field to given value.

### HasWindowDays

`func (o *Funnel) HasWindowDays() bool`

HasWindowDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


