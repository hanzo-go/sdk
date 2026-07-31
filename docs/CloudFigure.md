# CloudFigure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label names the metric, e.g. \&quot;MRR\&quot; or \&quot;Runway\&quot;. | [optional] 
**Period** | Pointer to **string** | Period is the window the figure covers, e.g. \&quot;2026-07\&quot; or \&quot;all-time\&quot;. | [optional] 
**Value** | Pointer to **string** | Value is the figure already formatted through books&#39; own money formatter, so a consumer never re-derives it. | [optional] 

## Methods

### NewCloudFigure

`func NewCloudFigure() *CloudFigure`

NewCloudFigure instantiates a new CloudFigure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFigureWithDefaults

`func NewCloudFigureWithDefaults() *CloudFigure`

NewCloudFigureWithDefaults instantiates a new CloudFigure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CloudFigure) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudFigure) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudFigure) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudFigure) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPeriod

`func (o *CloudFigure) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudFigure) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudFigure) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CloudFigure) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetValue

`func (o *CloudFigure) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudFigure) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudFigure) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudFigure) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


