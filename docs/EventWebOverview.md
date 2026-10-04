# EventWebOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the product-event table could not be read — the lens is reported missing rather than as zeros that look like real traffic. | [optional] 
**Pageviews** | Pointer to **int64** | Pageviews is how many $pageview events landed in the window. | [optional] 
**Reason** | Pointer to **string** | Reason says why the lens is unavailable. Omitted when it is available. | [optional] 
**Sessions** | Pointer to **int64** | Sessions is how many distinct visits they span. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 
**Visitors** | Pointer to **int64** | Visitors is how many distinct people those pageviews came from. | [optional] 

## Methods

### NewEventWebOverview

`func NewEventWebOverview() *EventWebOverview`

NewEventWebOverview instantiates a new EventWebOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWebOverviewWithDefaults

`func NewEventWebOverviewWithDefaults() *EventWebOverview`

NewEventWebOverviewWithDefaults instantiates a new EventWebOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *EventWebOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EventWebOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EventWebOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *EventWebOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPageviews

`func (o *EventWebOverview) GetPageviews() int64`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *EventWebOverview) GetPageviewsOk() (*int64, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *EventWebOverview) SetPageviews(v int64)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *EventWebOverview) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetReason

`func (o *EventWebOverview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventWebOverview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventWebOverview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventWebOverview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSessions

`func (o *EventWebOverview) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *EventWebOverview) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *EventWebOverview) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *EventWebOverview) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSource

`func (o *EventWebOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventWebOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventWebOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventWebOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisitors

`func (o *EventWebOverview) GetVisitors() int64`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *EventWebOverview) GetVisitorsOk() (*int64, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *EventWebOverview) SetVisitors(v int64)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *EventWebOverview) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


