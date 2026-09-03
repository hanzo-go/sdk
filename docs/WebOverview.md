# WebOverview

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

### NewWebOverview

`func NewWebOverview() *WebOverview`

NewWebOverview instantiates a new WebOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebOverviewWithDefaults

`func NewWebOverviewWithDefaults() *WebOverview`

NewWebOverviewWithDefaults instantiates a new WebOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *WebOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WebOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WebOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *WebOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPageviews

`func (o *WebOverview) GetPageviews() int64`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *WebOverview) GetPageviewsOk() (*int64, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *WebOverview) SetPageviews(v int64)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *WebOverview) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetReason

`func (o *WebOverview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WebOverview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WebOverview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WebOverview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSessions

`func (o *WebOverview) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *WebOverview) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *WebOverview) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *WebOverview) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSource

`func (o *WebOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisitors

`func (o *WebOverview) GetVisitors() int64`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *WebOverview) GetVisitorsOk() (*int64, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *WebOverview) SetVisitors(v int64)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *WebOverview) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


