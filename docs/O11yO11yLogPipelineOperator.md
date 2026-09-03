# O11yO11yLogPipelineOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** | Default is the id of the processor a router falls through to. | [optional] 
**EnableFlattening** | Pointer to **bool** | EnableFlattening flattens parsed JSON one level when true. | [optional] 
**EnablePaths** | Pointer to **bool** | EnablePaths keeps the JSON path in flattened keys when true. | [optional] 
**Enabled** | Pointer to **bool** | Enabled turns the processor on. | [optional] 
**Expr** | Pointer to **string** | Expr is a router route&#39;s expression. | [optional] 
**Field** | Pointer to **string** | Field is the field an add/remove processor works on. | [optional] 
**Fields** | Pointer to **[]string** | Fields are the fields a retain processor keeps. | [optional] 
**From** | Pointer to **string** | From is the source field of a move or copy. | [optional] 
**Id** | Pointer to **string** | ID is the processor&#39;s id, unique within the pipeline. | [optional] 
**If** | Pointer to **string** | If gates the processor on an expression. | [optional] 
**Layout** | Pointer to **string** | Layout is a time parser&#39;s layout. | [optional] 
**LayoutType** | Pointer to **string** | LayoutType is the layout&#39;s kind, e.g. strptime, gotime, epoch. | [optional] 
**Mapping** | Pointer to **map[string][]string** | Mapping maps severity levels (or flattened keys) to the values that mean them. | [optional] 
**Name** | Pointer to **string** | Name is the processor&#39;s display name. | [optional] 
**OnError** | Pointer to **string** | OnError says what to do when the processor fails, e.g. send, drop. | [optional] 
**OrderId** | Pointer to **int64** | OrderID is the processor&#39;s 1-based position in the pipeline. | [optional] 
**Output** | Pointer to **string** | Output is the id of the processor that runs next. | [optional] 
**OverwriteText** | Pointer to **bool** | OverwriteSeverityText rewrites the severity text alongside the number when true. | [optional] 
**ParseFrom** | Pointer to **string** | ParseFrom is where a parser reads from. | [optional] 
**ParseTo** | Pointer to **string** | ParseTo is where a parser writes its result. | [optional] 
**PathPrefix** | Pointer to **string** | PathPrefix prefixes flattened keys. | [optional] 
**Pattern** | Pointer to **string** | Pattern is a grok parser&#39;s pattern. | [optional] 
**Regex** | Pointer to **string** | Regex is a regex parser&#39;s expression. | [optional] 
**Routes** | Pointer to [**[]O11yO11yLogPipelineRoute**](O11yO11yLogPipelineRoute.md) | Routes are a router processor&#39;s routes. | [optional] 
**SpanId** | Pointer to [**O11yO11yLogParseFrom**](O11yO11yLogParseFrom.md) | SpanID says where a trace parser reads the span id from. | [optional] 
**To** | Pointer to **string** | To is the destination field of a move or copy. | [optional] 
**TraceFlags** | Pointer to [**O11yO11yLogParseFrom**](O11yO11yLogParseFrom.md) | TraceFlags says where a trace parser reads the trace flags from. | [optional] 
**TraceId** | Pointer to [**O11yO11yLogParseFrom**](O11yO11yLogParseFrom.md) | TraceID says where a trace parser reads the trace id from. | [optional] 
**Type** | Pointer to **string** | Type is the processor type, e.g. grok_parser, regex_parser, json_parser, trace_parser, time_parser, severity_parser, add, remove, move, copy. | [optional] 
**Value** | Pointer to **string** | Value is the value an add processor writes. | [optional] 

## Methods

### NewO11yO11yLogPipelineOperator

`func NewO11yO11yLogPipelineOperator() *O11yO11yLogPipelineOperator`

NewO11yO11yLogPipelineOperator instantiates a new O11yO11yLogPipelineOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPipelineOperatorWithDefaults

`func NewO11yO11yLogPipelineOperatorWithDefaults() *O11yO11yLogPipelineOperator`

NewO11yO11yLogPipelineOperatorWithDefaults instantiates a new O11yO11yLogPipelineOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *O11yO11yLogPipelineOperator) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *O11yO11yLogPipelineOperator) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *O11yO11yLogPipelineOperator) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *O11yO11yLogPipelineOperator) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnableFlattening

`func (o *O11yO11yLogPipelineOperator) GetEnableFlattening() bool`

GetEnableFlattening returns the EnableFlattening field if non-nil, zero value otherwise.

### GetEnableFlatteningOk

`func (o *O11yO11yLogPipelineOperator) GetEnableFlatteningOk() (*bool, bool)`

GetEnableFlatteningOk returns a tuple with the EnableFlattening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFlattening

`func (o *O11yO11yLogPipelineOperator) SetEnableFlattening(v bool)`

SetEnableFlattening sets EnableFlattening field to given value.

### HasEnableFlattening

`func (o *O11yO11yLogPipelineOperator) HasEnableFlattening() bool`

HasEnableFlattening returns a boolean if a field has been set.

### GetEnablePaths

`func (o *O11yO11yLogPipelineOperator) GetEnablePaths() bool`

GetEnablePaths returns the EnablePaths field if non-nil, zero value otherwise.

### GetEnablePathsOk

`func (o *O11yO11yLogPipelineOperator) GetEnablePathsOk() (*bool, bool)`

GetEnablePathsOk returns a tuple with the EnablePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePaths

`func (o *O11yO11yLogPipelineOperator) SetEnablePaths(v bool)`

SetEnablePaths sets EnablePaths field to given value.

### HasEnablePaths

`func (o *O11yO11yLogPipelineOperator) HasEnablePaths() bool`

HasEnablePaths returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yO11yLogPipelineOperator) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yO11yLogPipelineOperator) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yO11yLogPipelineOperator) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yO11yLogPipelineOperator) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpr

`func (o *O11yO11yLogPipelineOperator) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *O11yO11yLogPipelineOperator) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *O11yO11yLogPipelineOperator) SetExpr(v string)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *O11yO11yLogPipelineOperator) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### GetField

`func (o *O11yO11yLogPipelineOperator) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *O11yO11yLogPipelineOperator) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *O11yO11yLogPipelineOperator) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *O11yO11yLogPipelineOperator) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFields

`func (o *O11yO11yLogPipelineOperator) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *O11yO11yLogPipelineOperator) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *O11yO11yLogPipelineOperator) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *O11yO11yLogPipelineOperator) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFrom

`func (o *O11yO11yLogPipelineOperator) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *O11yO11yLogPipelineOperator) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *O11yO11yLogPipelineOperator) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *O11yO11yLogPipelineOperator) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLogPipelineOperator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLogPipelineOperator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLogPipelineOperator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLogPipelineOperator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIf

`func (o *O11yO11yLogPipelineOperator) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *O11yO11yLogPipelineOperator) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *O11yO11yLogPipelineOperator) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *O11yO11yLogPipelineOperator) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetLayout

`func (o *O11yO11yLogPipelineOperator) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *O11yO11yLogPipelineOperator) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *O11yO11yLogPipelineOperator) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *O11yO11yLogPipelineOperator) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetLayoutType

`func (o *O11yO11yLogPipelineOperator) GetLayoutType() string`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *O11yO11yLogPipelineOperator) GetLayoutTypeOk() (*string, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *O11yO11yLogPipelineOperator) SetLayoutType(v string)`

SetLayoutType sets LayoutType field to given value.

### HasLayoutType

`func (o *O11yO11yLogPipelineOperator) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### GetMapping

`func (o *O11yO11yLogPipelineOperator) GetMapping() map[string][]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *O11yO11yLogPipelineOperator) GetMappingOk() (*map[string][]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *O11yO11yLogPipelineOperator) SetMapping(v map[string][]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *O11yO11yLogPipelineOperator) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLogPipelineOperator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLogPipelineOperator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLogPipelineOperator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLogPipelineOperator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnError

`func (o *O11yO11yLogPipelineOperator) GetOnError() string`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *O11yO11yLogPipelineOperator) GetOnErrorOk() (*string, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *O11yO11yLogPipelineOperator) SetOnError(v string)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *O11yO11yLogPipelineOperator) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetOrderId

`func (o *O11yO11yLogPipelineOperator) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *O11yO11yLogPipelineOperator) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *O11yO11yLogPipelineOperator) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *O11yO11yLogPipelineOperator) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOutput

`func (o *O11yO11yLogPipelineOperator) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *O11yO11yLogPipelineOperator) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *O11yO11yLogPipelineOperator) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *O11yO11yLogPipelineOperator) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOverwriteText

`func (o *O11yO11yLogPipelineOperator) GetOverwriteText() bool`

GetOverwriteText returns the OverwriteText field if non-nil, zero value otherwise.

### GetOverwriteTextOk

`func (o *O11yO11yLogPipelineOperator) GetOverwriteTextOk() (*bool, bool)`

GetOverwriteTextOk returns a tuple with the OverwriteText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteText

`func (o *O11yO11yLogPipelineOperator) SetOverwriteText(v bool)`

SetOverwriteText sets OverwriteText field to given value.

### HasOverwriteText

`func (o *O11yO11yLogPipelineOperator) HasOverwriteText() bool`

HasOverwriteText returns a boolean if a field has been set.

### GetParseFrom

`func (o *O11yO11yLogPipelineOperator) GetParseFrom() string`

GetParseFrom returns the ParseFrom field if non-nil, zero value otherwise.

### GetParseFromOk

`func (o *O11yO11yLogPipelineOperator) GetParseFromOk() (*string, bool)`

GetParseFromOk returns a tuple with the ParseFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseFrom

`func (o *O11yO11yLogPipelineOperator) SetParseFrom(v string)`

SetParseFrom sets ParseFrom field to given value.

### HasParseFrom

`func (o *O11yO11yLogPipelineOperator) HasParseFrom() bool`

HasParseFrom returns a boolean if a field has been set.

### GetParseTo

`func (o *O11yO11yLogPipelineOperator) GetParseTo() string`

GetParseTo returns the ParseTo field if non-nil, zero value otherwise.

### GetParseToOk

`func (o *O11yO11yLogPipelineOperator) GetParseToOk() (*string, bool)`

GetParseToOk returns a tuple with the ParseTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseTo

`func (o *O11yO11yLogPipelineOperator) SetParseTo(v string)`

SetParseTo sets ParseTo field to given value.

### HasParseTo

`func (o *O11yO11yLogPipelineOperator) HasParseTo() bool`

HasParseTo returns a boolean if a field has been set.

### GetPathPrefix

`func (o *O11yO11yLogPipelineOperator) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *O11yO11yLogPipelineOperator) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *O11yO11yLogPipelineOperator) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *O11yO11yLogPipelineOperator) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPattern

`func (o *O11yO11yLogPipelineOperator) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *O11yO11yLogPipelineOperator) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *O11yO11yLogPipelineOperator) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *O11yO11yLogPipelineOperator) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRegex

`func (o *O11yO11yLogPipelineOperator) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *O11yO11yLogPipelineOperator) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *O11yO11yLogPipelineOperator) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *O11yO11yLogPipelineOperator) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetRoutes

`func (o *O11yO11yLogPipelineOperator) GetRoutes() []O11yO11yLogPipelineRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *O11yO11yLogPipelineOperator) GetRoutesOk() (*[]O11yO11yLogPipelineRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *O11yO11yLogPipelineOperator) SetRoutes(v []O11yO11yLogPipelineRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *O11yO11yLogPipelineOperator) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yO11yLogPipelineOperator) GetSpanId() O11yO11yLogParseFrom`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yO11yLogPipelineOperator) GetSpanIdOk() (*O11yO11yLogParseFrom, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yO11yLogPipelineOperator) SetSpanId(v O11yO11yLogParseFrom)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yO11yLogPipelineOperator) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTo

`func (o *O11yO11yLogPipelineOperator) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *O11yO11yLogPipelineOperator) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *O11yO11yLogPipelineOperator) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *O11yO11yLogPipelineOperator) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTraceFlags

`func (o *O11yO11yLogPipelineOperator) GetTraceFlags() O11yO11yLogParseFrom`

GetTraceFlags returns the TraceFlags field if non-nil, zero value otherwise.

### GetTraceFlagsOk

`func (o *O11yO11yLogPipelineOperator) GetTraceFlagsOk() (*O11yO11yLogParseFrom, bool)`

GetTraceFlagsOk returns a tuple with the TraceFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceFlags

`func (o *O11yO11yLogPipelineOperator) SetTraceFlags(v O11yO11yLogParseFrom)`

SetTraceFlags sets TraceFlags field to given value.

### HasTraceFlags

`func (o *O11yO11yLogPipelineOperator) HasTraceFlags() bool`

HasTraceFlags returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLogPipelineOperator) GetTraceId() O11yO11yLogParseFrom`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLogPipelineOperator) GetTraceIdOk() (*O11yO11yLogParseFrom, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLogPipelineOperator) SetTraceId(v O11yO11yLogParseFrom)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLogPipelineOperator) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yLogPipelineOperator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yLogPipelineOperator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yLogPipelineOperator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yLogPipelineOperator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yLogPipelineOperator) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yLogPipelineOperator) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yLogPipelineOperator) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yLogPipelineOperator) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


