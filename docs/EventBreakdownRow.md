# EventBreakdownRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the bucket: a requested path, a referrer domain (\&quot;(direct)\&quot; for none or a same-origin one), or a utm_source (\&quot;(none)\&quot; when absent). | [optional] 
**Pageviews** | Pointer to **int64** | Pageviews is how many $pageview events fell in this bucket. | [optional] 
**Pct** | Pointer to **float64** | Pct is this bucket&#39;s share of ALL in-window pageviews, 0..100, one decimal — not of the returned rows, so a top-N shows the long tail honestly. | [optional] 
**Visitors** | Pointer to **int64** | Visitors is how many distinct people they came from. | [optional] 

## Methods

### NewEventBreakdownRow

`func NewEventBreakdownRow() *EventBreakdownRow`

NewEventBreakdownRow instantiates a new EventBreakdownRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBreakdownRowWithDefaults

`func NewEventBreakdownRowWithDefaults() *EventBreakdownRow`

NewEventBreakdownRowWithDefaults instantiates a new EventBreakdownRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EventBreakdownRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventBreakdownRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventBreakdownRow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventBreakdownRow) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPageviews

`func (o *EventBreakdownRow) GetPageviews() int64`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *EventBreakdownRow) GetPageviewsOk() (*int64, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *EventBreakdownRow) SetPageviews(v int64)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *EventBreakdownRow) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetPct

`func (o *EventBreakdownRow) GetPct() float64`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *EventBreakdownRow) GetPctOk() (*float64, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *EventBreakdownRow) SetPct(v float64)`

SetPct sets Pct field to given value.

### HasPct

`func (o *EventBreakdownRow) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetVisitors

`func (o *EventBreakdownRow) GetVisitors() int64`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *EventBreakdownRow) GetVisitorsOk() (*int64, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *EventBreakdownRow) SetVisitors(v int64)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *EventBreakdownRow) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


