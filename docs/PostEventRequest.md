# PostEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Clip** | Pointer to [**ClipBody**](ClipBody.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**Exception**](Exception.md) |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Library** | Pointer to **string** |  | [optional] 
**LibraryVersion** | Pointer to **string** |  | [optional] 
**Log** | Pointer to [**LogBody**](LogBody.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**MetricBody**](MetricBody.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PersonId** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**RefCode** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to **float32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**SignupWeek** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**SpanBody**](SpanBody.md) |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Utm** | Pointer to [**UTM**](UTM.md) |  | [optional] 
**Batch** | Pointer to [**[]InsightsEvent**](InsightsEvent.md) |  | [optional] 
**Events** | Pointer to [**[]CaptureEvent**](CaptureEvent.md) |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewPostEventRequest

`func NewPostEventRequest() *PostEventRequest`

NewPostEventRequest instantiates a new PostEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEventRequestWithDefaults

`func NewPostEventRequestWithDefaults() *PostEventRequest`

NewPostEventRequestWithDefaults instantiates a new PostEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousId

`func (o *PostEventRequest) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *PostEventRequest) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *PostEventRequest) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.

### HasAnonymousId

`func (o *PostEventRequest) HasAnonymousId() bool`

HasAnonymousId returns a boolean if a field has been set.

### GetChannel

`func (o *PostEventRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostEventRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostEventRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostEventRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetClip

`func (o *PostEventRequest) GetClip() ClipBody`

GetClip returns the Clip field if non-nil, zero value otherwise.

### GetClipOk

`func (o *PostEventRequest) GetClipOk() (*ClipBody, bool)`

GetClipOk returns a tuple with the Clip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClip

`func (o *PostEventRequest) SetClip(v ClipBody)`

SetClip sets Clip field to given value.

### HasClip

`func (o *PostEventRequest) HasClip() bool`

HasClip returns a boolean if a field has been set.

### GetCurrency

`func (o *PostEventRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PostEventRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PostEventRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PostEventRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDistinctId

`func (o *PostEventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PostEventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PostEventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PostEventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEnvironment

`func (o *PostEventRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PostEventRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PostEventRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PostEventRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetError

`func (o *PostEventRequest) GetError() Exception`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostEventRequest) GetErrorOk() (*Exception, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostEventRequest) SetError(v Exception)`

SetError sets Error field to given value.

### HasError

`func (o *PostEventRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvent

`func (o *PostEventRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PostEventRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PostEventRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PostEventRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetGroupId

`func (o *PostEventRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PostEventRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PostEventRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PostEventRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupType

`func (o *PostEventRequest) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *PostEventRequest) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *PostEventRequest) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *PostEventRequest) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetKind

`func (o *PostEventRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PostEventRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PostEventRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PostEventRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLevel

`func (o *PostEventRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PostEventRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PostEventRequest) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *PostEventRequest) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLibrary

`func (o *PostEventRequest) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *PostEventRequest) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *PostEventRequest) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *PostEventRequest) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *PostEventRequest) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *PostEventRequest) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *PostEventRequest) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *PostEventRequest) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetLog

`func (o *PostEventRequest) GetLog() LogBody`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *PostEventRequest) GetLogOk() (*LogBody, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *PostEventRequest) SetLog(v LogBody)`

SetLog sets Log field to given value.

### HasLog

`func (o *PostEventRequest) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMessageId

`func (o *PostEventRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PostEventRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PostEventRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *PostEventRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetric

`func (o *PostEventRequest) GetMetric() MetricBody`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PostEventRequest) GetMetricOk() (*MetricBody, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PostEventRequest) SetMetric(v MetricBody)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *PostEventRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPath

`func (o *PostEventRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PostEventRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PostEventRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PostEventRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPersonId

`func (o *PostEventRequest) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *PostEventRequest) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *PostEventRequest) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *PostEventRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetProduct

`func (o *PostEventRequest) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PostEventRequest) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PostEventRequest) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PostEventRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductId

`func (o *PostEventRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PostEventRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PostEventRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PostEventRequest) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProperties

`func (o *PostEventRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PostEventRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PostEventRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PostEventRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *PostEventRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostEventRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostEventRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostEventRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefCode

`func (o *PostEventRequest) GetRefCode() string`

GetRefCode returns the RefCode field if non-nil, zero value otherwise.

### GetRefCodeOk

`func (o *PostEventRequest) GetRefCodeOk() (*string, bool)`

GetRefCodeOk returns a tuple with the RefCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCode

`func (o *PostEventRequest) SetRefCode(v string)`

SetRefCode sets RefCode field to given value.

### HasRefCode

`func (o *PostEventRequest) HasRefCode() bool`

HasRefCode returns a boolean if a field has been set.

### GetReferrer

`func (o *PostEventRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *PostEventRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *PostEventRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *PostEventRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetRelease

`func (o *PostEventRequest) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PostEventRequest) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PostEventRequest) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PostEventRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetResource

`func (o *PostEventRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PostEventRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PostEventRequest) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PostEventRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRevenue

`func (o *PostEventRequest) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *PostEventRequest) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *PostEventRequest) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *PostEventRequest) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetService

`func (o *PostEventRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostEventRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostEventRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PostEventRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSessionId

`func (o *PostEventRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PostEventRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PostEventRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PostEventRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSignupWeek

`func (o *PostEventRequest) GetSignupWeek() string`

GetSignupWeek returns the SignupWeek field if non-nil, zero value otherwise.

### GetSignupWeekOk

`func (o *PostEventRequest) GetSignupWeekOk() (*string, bool)`

GetSignupWeekOk returns a tuple with the SignupWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupWeek

`func (o *PostEventRequest) SetSignupWeek(v string)`

SetSignupWeek sets SignupWeek field to given value.

### HasSignupWeek

`func (o *PostEventRequest) HasSignupWeek() bool`

HasSignupWeek returns a boolean if a field has been set.

### GetSite

`func (o *PostEventRequest) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PostEventRequest) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PostEventRequest) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PostEventRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSpan

`func (o *PostEventRequest) GetSpan() SpanBody`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *PostEventRequest) GetSpanOk() (*SpanBody, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *PostEventRequest) SetSpan(v SpanBody)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *PostEventRequest) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetSpanId

`func (o *PostEventRequest) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *PostEventRequest) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *PostEventRequest) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *PostEventRequest) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PostEventRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PostEventRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PostEventRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PostEventRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *PostEventRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *PostEventRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *PostEventRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *PostEventRequest) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *PostEventRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostEventRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostEventRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostEventRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *PostEventRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostEventRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostEventRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostEventRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUtm

`func (o *PostEventRequest) GetUtm() UTM`

GetUtm returns the Utm field if non-nil, zero value otherwise.

### GetUtmOk

`func (o *PostEventRequest) GetUtmOk() (*UTM, bool)`

GetUtmOk returns a tuple with the Utm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtm

`func (o *PostEventRequest) SetUtm(v UTM)`

SetUtm sets Utm field to given value.

### HasUtm

`func (o *PostEventRequest) HasUtm() bool`

HasUtm returns a boolean if a field has been set.

### GetBatch

`func (o *PostEventRequest) GetBatch() []InsightsEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *PostEventRequest) GetBatchOk() (*[]InsightsEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *PostEventRequest) SetBatch(v []InsightsEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *PostEventRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetEvents

`func (o *PostEventRequest) GetEvents() []CaptureEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostEventRequest) GetEventsOk() (*[]CaptureEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostEventRequest) SetEvents(v []CaptureEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PostEventRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetDistinctId

`func (o *PostEventRequest) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PostEventRequest) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PostEventRequest) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PostEventRequest) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetUuid

`func (o *PostEventRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PostEventRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PostEventRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PostEventRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


