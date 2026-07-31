# CloudBreakdownRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the bucket: a requested path, a referrer domain (\&quot;(direct)\&quot; for none or a same-origin one), or a utm_source (\&quot;(none)\&quot; when absent). | [optional] 
**Pageviews** | Pointer to **int32** | Pageviews is how many $pageview events fell in this bucket. | [optional] 
**Pct** | Pointer to **float32** | Pct is this bucket&#39;s share of ALL in-window pageviews, 0..100, one decimal — not of the returned rows, so a top-N shows the long tail honestly. | [optional] 
**Visitors** | Pointer to **int32** | Visitors is how many distinct people they came from. | [optional] 

## Methods

### NewCloudBreakdownRow

`func NewCloudBreakdownRow() *CloudBreakdownRow`

NewCloudBreakdownRow instantiates a new CloudBreakdownRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBreakdownRowWithDefaults

`func NewCloudBreakdownRowWithDefaults() *CloudBreakdownRow`

NewCloudBreakdownRowWithDefaults instantiates a new CloudBreakdownRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CloudBreakdownRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudBreakdownRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudBreakdownRow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudBreakdownRow) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPageviews

`func (o *CloudBreakdownRow) GetPageviews() int32`

GetPageviews returns the Pageviews field if non-nil, zero value otherwise.

### GetPageviewsOk

`func (o *CloudBreakdownRow) GetPageviewsOk() (*int32, bool)`

GetPageviewsOk returns a tuple with the Pageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageviews

`func (o *CloudBreakdownRow) SetPageviews(v int32)`

SetPageviews sets Pageviews field to given value.

### HasPageviews

`func (o *CloudBreakdownRow) HasPageviews() bool`

HasPageviews returns a boolean if a field has been set.

### GetPct

`func (o *CloudBreakdownRow) GetPct() float32`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *CloudBreakdownRow) GetPctOk() (*float32, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *CloudBreakdownRow) SetPct(v float32)`

SetPct sets Pct field to given value.

### HasPct

`func (o *CloudBreakdownRow) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetVisitors

`func (o *CloudBreakdownRow) GetVisitors() int32`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *CloudBreakdownRow) GetVisitorsOk() (*int32, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *CloudBreakdownRow) SetVisitors(v int32)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *CloudBreakdownRow) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


