# O11yWaterfallSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DbName** | Pointer to **string** | Calculated fields https://o11y.io/docs/traces-management/guides/derived-fields-spans | [optional] 
**DbOperation** | Pointer to **string** |  | [optional] 
**DurationNano** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to [**[]O11yEvent**](O11yEvent.md) |  | [optional] 
**ExternalHttpMethod** | Pointer to **string** |  | [optional] 
**ExternalHttpUrl** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**HasChildren** | Pointer to **bool** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 
**HttpHost** | Pointer to **string** |  | [optional] 
**HttpMethod** | Pointer to **string** |  | [optional] 
**HttpUrl** | Pointer to **string** |  | [optional] 
**IsRemote** | Pointer to **string** |  | [optional] 
**KindString** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]O11yOtelSpanRef**](O11yOtelSpanRef.md) |  | [optional] 
**Resource** | Pointer to **map[string]string** |  | [optional] 
**ResponseStatusCode** | Pointer to **string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**StatusCodeString** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**SubTreeNodeCount** | Pointer to **int32** |  | [optional] 
**TimeUnix** | Pointer to **int32** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**TraceState** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yWaterfallSpan

`func NewO11yWaterfallSpan() *O11yWaterfallSpan`

NewO11yWaterfallSpan instantiates a new O11yWaterfallSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yWaterfallSpanWithDefaults

`func NewO11yWaterfallSpanWithDefaults() *O11yWaterfallSpan`

NewO11yWaterfallSpanWithDefaults instantiates a new O11yWaterfallSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *O11yWaterfallSpan) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11yWaterfallSpan) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11yWaterfallSpan) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11yWaterfallSpan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDbName

`func (o *O11yWaterfallSpan) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *O11yWaterfallSpan) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *O11yWaterfallSpan) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *O11yWaterfallSpan) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbOperation

`func (o *O11yWaterfallSpan) GetDbOperation() string`

GetDbOperation returns the DbOperation field if non-nil, zero value otherwise.

### GetDbOperationOk

`func (o *O11yWaterfallSpan) GetDbOperationOk() (*string, bool)`

GetDbOperationOk returns a tuple with the DbOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOperation

`func (o *O11yWaterfallSpan) SetDbOperation(v string)`

SetDbOperation sets DbOperation field to given value.

### HasDbOperation

`func (o *O11yWaterfallSpan) HasDbOperation() bool`

HasDbOperation returns a boolean if a field has been set.

### GetDurationNano

`func (o *O11yWaterfallSpan) GetDurationNano() int32`

GetDurationNano returns the DurationNano field if non-nil, zero value otherwise.

### GetDurationNanoOk

`func (o *O11yWaterfallSpan) GetDurationNanoOk() (*int32, bool)`

GetDurationNanoOk returns a tuple with the DurationNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNano

`func (o *O11yWaterfallSpan) SetDurationNano(v int32)`

SetDurationNano sets DurationNano field to given value.

### HasDurationNano

`func (o *O11yWaterfallSpan) HasDurationNano() bool`

HasDurationNano returns a boolean if a field has been set.

### GetEvents

`func (o *O11yWaterfallSpan) GetEvents() []O11yEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *O11yWaterfallSpan) GetEventsOk() (*[]O11yEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *O11yWaterfallSpan) SetEvents(v []O11yEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *O11yWaterfallSpan) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExternalHttpMethod

`func (o *O11yWaterfallSpan) GetExternalHttpMethod() string`

GetExternalHttpMethod returns the ExternalHttpMethod field if non-nil, zero value otherwise.

### GetExternalHttpMethodOk

`func (o *O11yWaterfallSpan) GetExternalHttpMethodOk() (*string, bool)`

GetExternalHttpMethodOk returns a tuple with the ExternalHttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpMethod

`func (o *O11yWaterfallSpan) SetExternalHttpMethod(v string)`

SetExternalHttpMethod sets ExternalHttpMethod field to given value.

### HasExternalHttpMethod

`func (o *O11yWaterfallSpan) HasExternalHttpMethod() bool`

HasExternalHttpMethod returns a boolean if a field has been set.

### GetExternalHttpUrl

`func (o *O11yWaterfallSpan) GetExternalHttpUrl() string`

GetExternalHttpUrl returns the ExternalHttpUrl field if non-nil, zero value otherwise.

### GetExternalHttpUrlOk

`func (o *O11yWaterfallSpan) GetExternalHttpUrlOk() (*string, bool)`

GetExternalHttpUrlOk returns a tuple with the ExternalHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpUrl

`func (o *O11yWaterfallSpan) SetExternalHttpUrl(v string)`

SetExternalHttpUrl sets ExternalHttpUrl field to given value.

### HasExternalHttpUrl

`func (o *O11yWaterfallSpan) HasExternalHttpUrl() bool`

HasExternalHttpUrl returns a boolean if a field has been set.

### GetFlags

`func (o *O11yWaterfallSpan) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *O11yWaterfallSpan) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *O11yWaterfallSpan) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *O11yWaterfallSpan) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetHasChildren

`func (o *O11yWaterfallSpan) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *O11yWaterfallSpan) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *O11yWaterfallSpan) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *O11yWaterfallSpan) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetHasError

`func (o *O11yWaterfallSpan) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *O11yWaterfallSpan) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *O11yWaterfallSpan) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *O11yWaterfallSpan) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetHttpHost

`func (o *O11yWaterfallSpan) GetHttpHost() string`

GetHttpHost returns the HttpHost field if non-nil, zero value otherwise.

### GetHttpHostOk

`func (o *O11yWaterfallSpan) GetHttpHostOk() (*string, bool)`

GetHttpHostOk returns a tuple with the HttpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHost

`func (o *O11yWaterfallSpan) SetHttpHost(v string)`

SetHttpHost sets HttpHost field to given value.

### HasHttpHost

`func (o *O11yWaterfallSpan) HasHttpHost() bool`

HasHttpHost returns a boolean if a field has been set.

### GetHttpMethod

`func (o *O11yWaterfallSpan) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *O11yWaterfallSpan) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *O11yWaterfallSpan) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *O11yWaterfallSpan) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpUrl

`func (o *O11yWaterfallSpan) GetHttpUrl() string`

GetHttpUrl returns the HttpUrl field if non-nil, zero value otherwise.

### GetHttpUrlOk

`func (o *O11yWaterfallSpan) GetHttpUrlOk() (*string, bool)`

GetHttpUrlOk returns a tuple with the HttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUrl

`func (o *O11yWaterfallSpan) SetHttpUrl(v string)`

SetHttpUrl sets HttpUrl field to given value.

### HasHttpUrl

`func (o *O11yWaterfallSpan) HasHttpUrl() bool`

HasHttpUrl returns a boolean if a field has been set.

### GetIsRemote

`func (o *O11yWaterfallSpan) GetIsRemote() string`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *O11yWaterfallSpan) GetIsRemoteOk() (*string, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *O11yWaterfallSpan) SetIsRemote(v string)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *O11yWaterfallSpan) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetKindString

`func (o *O11yWaterfallSpan) GetKindString() string`

GetKindString returns the KindString field if non-nil, zero value otherwise.

### GetKindStringOk

`func (o *O11yWaterfallSpan) GetKindStringOk() (*string, bool)`

GetKindStringOk returns a tuple with the KindString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindString

`func (o *O11yWaterfallSpan) SetKindString(v string)`

SetKindString sets KindString field to given value.

### HasKindString

`func (o *O11yWaterfallSpan) HasKindString() bool`

HasKindString returns a boolean if a field has been set.

### GetLevel

`func (o *O11yWaterfallSpan) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *O11yWaterfallSpan) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *O11yWaterfallSpan) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *O11yWaterfallSpan) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *O11yWaterfallSpan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yWaterfallSpan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yWaterfallSpan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yWaterfallSpan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentSpanId

`func (o *O11yWaterfallSpan) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *O11yWaterfallSpan) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *O11yWaterfallSpan) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *O11yWaterfallSpan) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetReferences

`func (o *O11yWaterfallSpan) GetReferences() []O11yOtelSpanRef`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *O11yWaterfallSpan) GetReferencesOk() (*[]O11yOtelSpanRef, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *O11yWaterfallSpan) SetReferences(v []O11yOtelSpanRef)`

SetReferences sets References field to given value.

### HasReferences

`func (o *O11yWaterfallSpan) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetResource

`func (o *O11yWaterfallSpan) GetResource() map[string]string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *O11yWaterfallSpan) GetResourceOk() (*map[string]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *O11yWaterfallSpan) SetResource(v map[string]string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *O11yWaterfallSpan) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *O11yWaterfallSpan) GetResponseStatusCode() string`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *O11yWaterfallSpan) GetResponseStatusCodeOk() (*string, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *O11yWaterfallSpan) SetResponseStatusCode(v string)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *O11yWaterfallSpan) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yWaterfallSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yWaterfallSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yWaterfallSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yWaterfallSpan) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetStatusCode

`func (o *O11yWaterfallSpan) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *O11yWaterfallSpan) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *O11yWaterfallSpan) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *O11yWaterfallSpan) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusCodeString

`func (o *O11yWaterfallSpan) GetStatusCodeString() string`

GetStatusCodeString returns the StatusCodeString field if non-nil, zero value otherwise.

### GetStatusCodeStringOk

`func (o *O11yWaterfallSpan) GetStatusCodeStringOk() (*string, bool)`

GetStatusCodeStringOk returns a tuple with the StatusCodeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeString

`func (o *O11yWaterfallSpan) SetStatusCodeString(v string)`

SetStatusCodeString sets StatusCodeString field to given value.

### HasStatusCodeString

`func (o *O11yWaterfallSpan) HasStatusCodeString() bool`

HasStatusCodeString returns a boolean if a field has been set.

### GetStatusMessage

`func (o *O11yWaterfallSpan) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *O11yWaterfallSpan) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *O11yWaterfallSpan) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *O11yWaterfallSpan) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSubTreeNodeCount

`func (o *O11yWaterfallSpan) GetSubTreeNodeCount() int32`

GetSubTreeNodeCount returns the SubTreeNodeCount field if non-nil, zero value otherwise.

### GetSubTreeNodeCountOk

`func (o *O11yWaterfallSpan) GetSubTreeNodeCountOk() (*int32, bool)`

GetSubTreeNodeCountOk returns a tuple with the SubTreeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTreeNodeCount

`func (o *O11yWaterfallSpan) SetSubTreeNodeCount(v int32)`

SetSubTreeNodeCount sets SubTreeNodeCount field to given value.

### HasSubTreeNodeCount

`func (o *O11yWaterfallSpan) HasSubTreeNodeCount() bool`

HasSubTreeNodeCount returns a boolean if a field has been set.

### GetTimeUnix

`func (o *O11yWaterfallSpan) GetTimeUnix() int32`

GetTimeUnix returns the TimeUnix field if non-nil, zero value otherwise.

### GetTimeUnixOk

`func (o *O11yWaterfallSpan) GetTimeUnixOk() (*int32, bool)`

GetTimeUnixOk returns a tuple with the TimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnix

`func (o *O11yWaterfallSpan) SetTimeUnix(v int32)`

SetTimeUnix sets TimeUnix field to given value.

### HasTimeUnix

`func (o *O11yWaterfallSpan) HasTimeUnix() bool`

HasTimeUnix returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yWaterfallSpan) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yWaterfallSpan) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yWaterfallSpan) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yWaterfallSpan) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetTraceState

`func (o *O11yWaterfallSpan) GetTraceState() string`

GetTraceState returns the TraceState field if non-nil, zero value otherwise.

### GetTraceStateOk

`func (o *O11yWaterfallSpan) GetTraceStateOk() (*string, bool)`

GetTraceStateOk returns a tuple with the TraceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceState

`func (o *O11yWaterfallSpan) SetTraceState(v string)`

SetTraceState sets TraceState field to given value.

### HasTraceState

`func (o *O11yWaterfallSpan) HasTraceState() bool`

HasTraceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


