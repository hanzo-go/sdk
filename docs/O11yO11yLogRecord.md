# O11yO11yLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributesBool** | Pointer to **map[string]bool** | AttributesBool are the record&#39;s boolean attributes. | [optional] 
**AttributesFloat** | Pointer to **map[string]float32** | AttributesFloat are the record&#39;s float attributes. | [optional] 
**AttributesInt** | Pointer to **map[string]int32** | AttributesInt are the record&#39;s integer attributes. | [optional] 
**AttributesString** | Pointer to **map[string]string** | AttributesString are the record&#39;s string attributes. | [optional] 
**Body** | Pointer to **string** | Body is the record&#39;s body. | [optional] 
**Id** | Pointer to **string** | ID is the record&#39;s id. | [optional] 
**ResourcesString** | Pointer to **map[string]string** | ResourcesString are the record&#39;s string resource attributes. | [optional] 
**SeverityNumber** | Pointer to **int32** | SeverityNumber is the record&#39;s severity as a number. | [optional] 
**SeverityText** | Pointer to **string** | SeverityText is the record&#39;s severity as text, e.g. ERROR. | [optional] 
**SpanId** | Pointer to **string** | SpanID is the span the record belongs to. | [optional] 
**Timestamp** | Pointer to **int32** | Timestamp is the record&#39;s time as a nanosecond epoch. | [optional] 
**TraceFlags** | Pointer to **int32** | TraceFlags are the record&#39;s trace flags. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the record belongs to. | [optional] 

## Methods

### NewO11yO11yLogRecord

`func NewO11yO11yLogRecord() *O11yO11yLogRecord`

NewO11yO11yLogRecord instantiates a new O11yO11yLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogRecordWithDefaults

`func NewO11yO11yLogRecordWithDefaults() *O11yO11yLogRecord`

NewO11yO11yLogRecordWithDefaults instantiates a new O11yO11yLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributesBool

`func (o *O11yO11yLogRecord) GetAttributesBool() map[string]bool`

GetAttributesBool returns the AttributesBool field if non-nil, zero value otherwise.

### GetAttributesBoolOk

`func (o *O11yO11yLogRecord) GetAttributesBoolOk() (*map[string]bool, bool)`

GetAttributesBoolOk returns a tuple with the AttributesBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesBool

`func (o *O11yO11yLogRecord) SetAttributesBool(v map[string]bool)`

SetAttributesBool sets AttributesBool field to given value.

### HasAttributesBool

`func (o *O11yO11yLogRecord) HasAttributesBool() bool`

HasAttributesBool returns a boolean if a field has been set.

### GetAttributesFloat

`func (o *O11yO11yLogRecord) GetAttributesFloat() map[string]float32`

GetAttributesFloat returns the AttributesFloat field if non-nil, zero value otherwise.

### GetAttributesFloatOk

`func (o *O11yO11yLogRecord) GetAttributesFloatOk() (*map[string]float32, bool)`

GetAttributesFloatOk returns a tuple with the AttributesFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesFloat

`func (o *O11yO11yLogRecord) SetAttributesFloat(v map[string]float32)`

SetAttributesFloat sets AttributesFloat field to given value.

### HasAttributesFloat

`func (o *O11yO11yLogRecord) HasAttributesFloat() bool`

HasAttributesFloat returns a boolean if a field has been set.

### GetAttributesInt

`func (o *O11yO11yLogRecord) GetAttributesInt() map[string]int32`

GetAttributesInt returns the AttributesInt field if non-nil, zero value otherwise.

### GetAttributesIntOk

`func (o *O11yO11yLogRecord) GetAttributesIntOk() (*map[string]int32, bool)`

GetAttributesIntOk returns a tuple with the AttributesInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesInt

`func (o *O11yO11yLogRecord) SetAttributesInt(v map[string]int32)`

SetAttributesInt sets AttributesInt field to given value.

### HasAttributesInt

`func (o *O11yO11yLogRecord) HasAttributesInt() bool`

HasAttributesInt returns a boolean if a field has been set.

### GetAttributesString

`func (o *O11yO11yLogRecord) GetAttributesString() map[string]string`

GetAttributesString returns the AttributesString field if non-nil, zero value otherwise.

### GetAttributesStringOk

`func (o *O11yO11yLogRecord) GetAttributesStringOk() (*map[string]string, bool)`

GetAttributesStringOk returns a tuple with the AttributesString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesString

`func (o *O11yO11yLogRecord) SetAttributesString(v map[string]string)`

SetAttributesString sets AttributesString field to given value.

### HasAttributesString

`func (o *O11yO11yLogRecord) HasAttributesString() bool`

HasAttributesString returns a boolean if a field has been set.

### GetBody

`func (o *O11yO11yLogRecord) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *O11yO11yLogRecord) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *O11yO11yLogRecord) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *O11yO11yLogRecord) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourcesString

`func (o *O11yO11yLogRecord) GetResourcesString() map[string]string`

GetResourcesString returns the ResourcesString field if non-nil, zero value otherwise.

### GetResourcesStringOk

`func (o *O11yO11yLogRecord) GetResourcesStringOk() (*map[string]string, bool)`

GetResourcesStringOk returns a tuple with the ResourcesString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesString

`func (o *O11yO11yLogRecord) SetResourcesString(v map[string]string)`

SetResourcesString sets ResourcesString field to given value.

### HasResourcesString

`func (o *O11yO11yLogRecord) HasResourcesString() bool`

HasResourcesString returns a boolean if a field has been set.

### GetSeverityNumber

`func (o *O11yO11yLogRecord) GetSeverityNumber() int32`

GetSeverityNumber returns the SeverityNumber field if non-nil, zero value otherwise.

### GetSeverityNumberOk

`func (o *O11yO11yLogRecord) GetSeverityNumberOk() (*int32, bool)`

GetSeverityNumberOk returns a tuple with the SeverityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityNumber

`func (o *O11yO11yLogRecord) SetSeverityNumber(v int32)`

SetSeverityNumber sets SeverityNumber field to given value.

### HasSeverityNumber

`func (o *O11yO11yLogRecord) HasSeverityNumber() bool`

HasSeverityNumber returns a boolean if a field has been set.

### GetSeverityText

`func (o *O11yO11yLogRecord) GetSeverityText() string`

GetSeverityText returns the SeverityText field if non-nil, zero value otherwise.

### GetSeverityTextOk

`func (o *O11yO11yLogRecord) GetSeverityTextOk() (*string, bool)`

GetSeverityTextOk returns a tuple with the SeverityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityText

`func (o *O11yO11yLogRecord) SetSeverityText(v string)`

SetSeverityText sets SeverityText field to given value.

### HasSeverityText

`func (o *O11yO11yLogRecord) HasSeverityText() bool`

HasSeverityText returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yO11yLogRecord) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yO11yLogRecord) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yO11yLogRecord) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yO11yLogRecord) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yLogRecord) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yLogRecord) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yLogRecord) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yLogRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceFlags

`func (o *O11yO11yLogRecord) GetTraceFlags() int32`

GetTraceFlags returns the TraceFlags field if non-nil, zero value otherwise.

### GetTraceFlagsOk

`func (o *O11yO11yLogRecord) GetTraceFlagsOk() (*int32, bool)`

GetTraceFlagsOk returns a tuple with the TraceFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceFlags

`func (o *O11yO11yLogRecord) SetTraceFlags(v int32)`

SetTraceFlags sets TraceFlags field to given value.

### HasTraceFlags

`func (o *O11yO11yLogRecord) HasTraceFlags() bool`

HasTraceFlags returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLogRecord) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLogRecord) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLogRecord) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLogRecord) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


