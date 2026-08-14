# CaptureEvent

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

## Methods

### NewCaptureEvent

`func NewCaptureEvent() *CaptureEvent`

NewCaptureEvent instantiates a new CaptureEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureEventWithDefaults

`func NewCaptureEventWithDefaults() *CaptureEvent`

NewCaptureEventWithDefaults instantiates a new CaptureEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousId

`func (o *CaptureEvent) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *CaptureEvent) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *CaptureEvent) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.

### HasAnonymousId

`func (o *CaptureEvent) HasAnonymousId() bool`

HasAnonymousId returns a boolean if a field has been set.

### GetChannel

`func (o *CaptureEvent) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CaptureEvent) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CaptureEvent) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CaptureEvent) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetClip

`func (o *CaptureEvent) GetClip() ClipBody`

GetClip returns the Clip field if non-nil, zero value otherwise.

### GetClipOk

`func (o *CaptureEvent) GetClipOk() (*ClipBody, bool)`

GetClipOk returns a tuple with the Clip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClip

`func (o *CaptureEvent) SetClip(v ClipBody)`

SetClip sets Clip field to given value.

### HasClip

`func (o *CaptureEvent) HasClip() bool`

HasClip returns a boolean if a field has been set.

### GetCurrency

`func (o *CaptureEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CaptureEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CaptureEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CaptureEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDistinctId

`func (o *CaptureEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CaptureEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CaptureEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CaptureEvent) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEnvironment

`func (o *CaptureEvent) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CaptureEvent) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CaptureEvent) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CaptureEvent) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetError

`func (o *CaptureEvent) GetError() Exception`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CaptureEvent) GetErrorOk() (*Exception, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CaptureEvent) SetError(v Exception)`

SetError sets Error field to given value.

### HasError

`func (o *CaptureEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvent

`func (o *CaptureEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CaptureEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CaptureEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CaptureEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetGroupId

`func (o *CaptureEvent) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CaptureEvent) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CaptureEvent) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CaptureEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupType

`func (o *CaptureEvent) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *CaptureEvent) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *CaptureEvent) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *CaptureEvent) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetKind

`func (o *CaptureEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CaptureEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CaptureEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CaptureEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLevel

`func (o *CaptureEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CaptureEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CaptureEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CaptureEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLibrary

`func (o *CaptureEvent) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *CaptureEvent) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *CaptureEvent) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *CaptureEvent) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *CaptureEvent) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *CaptureEvent) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *CaptureEvent) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *CaptureEvent) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetLog

`func (o *CaptureEvent) GetLog() LogBody`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CaptureEvent) GetLogOk() (*LogBody, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CaptureEvent) SetLog(v LogBody)`

SetLog sets Log field to given value.

### HasLog

`func (o *CaptureEvent) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMessageId

`func (o *CaptureEvent) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *CaptureEvent) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *CaptureEvent) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *CaptureEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetric

`func (o *CaptureEvent) GetMetric() MetricBody`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CaptureEvent) GetMetricOk() (*MetricBody, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CaptureEvent) SetMetric(v MetricBody)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CaptureEvent) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPath

`func (o *CaptureEvent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CaptureEvent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CaptureEvent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CaptureEvent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPersonId

`func (o *CaptureEvent) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *CaptureEvent) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *CaptureEvent) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *CaptureEvent) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetProduct

`func (o *CaptureEvent) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CaptureEvent) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CaptureEvent) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CaptureEvent) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductId

`func (o *CaptureEvent) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CaptureEvent) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CaptureEvent) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CaptureEvent) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProperties

`func (o *CaptureEvent) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CaptureEvent) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CaptureEvent) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CaptureEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *CaptureEvent) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CaptureEvent) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CaptureEvent) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CaptureEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefCode

`func (o *CaptureEvent) GetRefCode() string`

GetRefCode returns the RefCode field if non-nil, zero value otherwise.

### GetRefCodeOk

`func (o *CaptureEvent) GetRefCodeOk() (*string, bool)`

GetRefCodeOk returns a tuple with the RefCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCode

`func (o *CaptureEvent) SetRefCode(v string)`

SetRefCode sets RefCode field to given value.

### HasRefCode

`func (o *CaptureEvent) HasRefCode() bool`

HasRefCode returns a boolean if a field has been set.

### GetReferrer

`func (o *CaptureEvent) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *CaptureEvent) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *CaptureEvent) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *CaptureEvent) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetRelease

`func (o *CaptureEvent) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CaptureEvent) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CaptureEvent) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *CaptureEvent) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetResource

`func (o *CaptureEvent) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CaptureEvent) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CaptureEvent) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CaptureEvent) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRevenue

`func (o *CaptureEvent) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CaptureEvent) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CaptureEvent) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CaptureEvent) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetService

`func (o *CaptureEvent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CaptureEvent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CaptureEvent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CaptureEvent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSessionId

`func (o *CaptureEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CaptureEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CaptureEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CaptureEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSignupWeek

`func (o *CaptureEvent) GetSignupWeek() string`

GetSignupWeek returns the SignupWeek field if non-nil, zero value otherwise.

### GetSignupWeekOk

`func (o *CaptureEvent) GetSignupWeekOk() (*string, bool)`

GetSignupWeekOk returns a tuple with the SignupWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupWeek

`func (o *CaptureEvent) SetSignupWeek(v string)`

SetSignupWeek sets SignupWeek field to given value.

### HasSignupWeek

`func (o *CaptureEvent) HasSignupWeek() bool`

HasSignupWeek returns a boolean if a field has been set.

### GetSite

`func (o *CaptureEvent) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CaptureEvent) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CaptureEvent) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *CaptureEvent) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSpan

`func (o *CaptureEvent) GetSpan() SpanBody`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *CaptureEvent) GetSpanOk() (*SpanBody, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *CaptureEvent) SetSpan(v SpanBody)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *CaptureEvent) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetSpanId

`func (o *CaptureEvent) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *CaptureEvent) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *CaptureEvent) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *CaptureEvent) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CaptureEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CaptureEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CaptureEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CaptureEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *CaptureEvent) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CaptureEvent) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CaptureEvent) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *CaptureEvent) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *CaptureEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptureEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *CaptureEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CaptureEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CaptureEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CaptureEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUtm

`func (o *CaptureEvent) GetUtm() UTM`

GetUtm returns the Utm field if non-nil, zero value otherwise.

### GetUtmOk

`func (o *CaptureEvent) GetUtmOk() (*UTM, bool)`

GetUtmOk returns a tuple with the Utm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtm

`func (o *CaptureEvent) SetUtm(v UTM)`

SetUtm sets Utm field to given value.

### HasUtm

`func (o *CaptureEvent) HasUtm() bool`

HasUtm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


