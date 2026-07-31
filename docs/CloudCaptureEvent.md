# CloudCaptureEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**CloudException**](CloudException.md) |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Library** | Pointer to **string** |  | [optional] 
**LibraryVersion** | Pointer to **string** |  | [optional] 
**Log** | Pointer to [**CloudLogBody**](CloudLogBody.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**CloudMetricBody**](CloudMetricBody.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PersonId** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**RefCode** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to **float32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**SignupWeek** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CloudSpanBody**](CloudSpanBody.md) |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Utm** | Pointer to [**CloudUTM**](CloudUTM.md) |  | [optional] 

## Methods

### NewCloudCaptureEvent

`func NewCloudCaptureEvent() *CloudCaptureEvent`

NewCloudCaptureEvent instantiates a new CloudCaptureEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptureEventWithDefaults

`func NewCloudCaptureEventWithDefaults() *CloudCaptureEvent`

NewCloudCaptureEventWithDefaults instantiates a new CloudCaptureEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousId

`func (o *CloudCaptureEvent) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *CloudCaptureEvent) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *CloudCaptureEvent) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.

### HasAnonymousId

`func (o *CloudCaptureEvent) HasAnonymousId() bool`

HasAnonymousId returns a boolean if a field has been set.

### GetChannel

`func (o *CloudCaptureEvent) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudCaptureEvent) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudCaptureEvent) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudCaptureEvent) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudCaptureEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudCaptureEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudCaptureEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudCaptureEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDistinctId

`func (o *CloudCaptureEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudCaptureEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudCaptureEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudCaptureEvent) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEnvironment

`func (o *CloudCaptureEvent) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CloudCaptureEvent) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CloudCaptureEvent) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CloudCaptureEvent) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetError

`func (o *CloudCaptureEvent) GetError() CloudException`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudCaptureEvent) GetErrorOk() (*CloudException, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudCaptureEvent) SetError(v CloudException)`

SetError sets Error field to given value.

### HasError

`func (o *CloudCaptureEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvent

`func (o *CloudCaptureEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudCaptureEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudCaptureEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudCaptureEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetGroupId

`func (o *CloudCaptureEvent) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CloudCaptureEvent) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CloudCaptureEvent) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CloudCaptureEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetKind

`func (o *CloudCaptureEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudCaptureEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudCaptureEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudCaptureEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLevel

`func (o *CloudCaptureEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CloudCaptureEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CloudCaptureEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CloudCaptureEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLibrary

`func (o *CloudCaptureEvent) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *CloudCaptureEvent) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *CloudCaptureEvent) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *CloudCaptureEvent) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *CloudCaptureEvent) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *CloudCaptureEvent) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *CloudCaptureEvent) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *CloudCaptureEvent) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetLog

`func (o *CloudCaptureEvent) GetLog() CloudLogBody`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CloudCaptureEvent) GetLogOk() (*CloudLogBody, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CloudCaptureEvent) SetLog(v CloudLogBody)`

SetLog sets Log field to given value.

### HasLog

`func (o *CloudCaptureEvent) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMessageId

`func (o *CloudCaptureEvent) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *CloudCaptureEvent) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *CloudCaptureEvent) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *CloudCaptureEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMetric

`func (o *CloudCaptureEvent) GetMetric() CloudMetricBody`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CloudCaptureEvent) GetMetricOk() (*CloudMetricBody, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CloudCaptureEvent) SetMetric(v CloudMetricBody)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CloudCaptureEvent) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPath

`func (o *CloudCaptureEvent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudCaptureEvent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudCaptureEvent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudCaptureEvent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPersonId

`func (o *CloudCaptureEvent) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *CloudCaptureEvent) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *CloudCaptureEvent) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *CloudCaptureEvent) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetProduct

`func (o *CloudCaptureEvent) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudCaptureEvent) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudCaptureEvent) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudCaptureEvent) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductId

`func (o *CloudCaptureEvent) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CloudCaptureEvent) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CloudCaptureEvent) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CloudCaptureEvent) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProperties

`func (o *CloudCaptureEvent) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CloudCaptureEvent) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CloudCaptureEvent) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CloudCaptureEvent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *CloudCaptureEvent) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CloudCaptureEvent) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CloudCaptureEvent) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CloudCaptureEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefCode

`func (o *CloudCaptureEvent) GetRefCode() string`

GetRefCode returns the RefCode field if non-nil, zero value otherwise.

### GetRefCodeOk

`func (o *CloudCaptureEvent) GetRefCodeOk() (*string, bool)`

GetRefCodeOk returns a tuple with the RefCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCode

`func (o *CloudCaptureEvent) SetRefCode(v string)`

SetRefCode sets RefCode field to given value.

### HasRefCode

`func (o *CloudCaptureEvent) HasRefCode() bool`

HasRefCode returns a boolean if a field has been set.

### GetReferrer

`func (o *CloudCaptureEvent) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *CloudCaptureEvent) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *CloudCaptureEvent) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *CloudCaptureEvent) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetRelease

`func (o *CloudCaptureEvent) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CloudCaptureEvent) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CloudCaptureEvent) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *CloudCaptureEvent) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudCaptureEvent) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudCaptureEvent) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudCaptureEvent) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudCaptureEvent) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetService

`func (o *CloudCaptureEvent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudCaptureEvent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudCaptureEvent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudCaptureEvent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSessionId

`func (o *CloudCaptureEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudCaptureEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudCaptureEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudCaptureEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSignupWeek

`func (o *CloudCaptureEvent) GetSignupWeek() string`

GetSignupWeek returns the SignupWeek field if non-nil, zero value otherwise.

### GetSignupWeekOk

`func (o *CloudCaptureEvent) GetSignupWeekOk() (*string, bool)`

GetSignupWeekOk returns a tuple with the SignupWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupWeek

`func (o *CloudCaptureEvent) SetSignupWeek(v string)`

SetSignupWeek sets SignupWeek field to given value.

### HasSignupWeek

`func (o *CloudCaptureEvent) HasSignupWeek() bool`

HasSignupWeek returns a boolean if a field has been set.

### GetSite

`func (o *CloudCaptureEvent) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CloudCaptureEvent) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CloudCaptureEvent) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *CloudCaptureEvent) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSpan

`func (o *CloudCaptureEvent) GetSpan() CloudSpanBody`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *CloudCaptureEvent) GetSpanOk() (*CloudSpanBody, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *CloudCaptureEvent) SetSpan(v CloudSpanBody)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *CloudCaptureEvent) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetSpanId

`func (o *CloudCaptureEvent) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *CloudCaptureEvent) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *CloudCaptureEvent) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *CloudCaptureEvent) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CloudCaptureEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CloudCaptureEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CloudCaptureEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CloudCaptureEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *CloudCaptureEvent) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CloudCaptureEvent) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CloudCaptureEvent) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *CloudCaptureEvent) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *CloudCaptureEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCaptureEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCaptureEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCaptureEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *CloudCaptureEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudCaptureEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudCaptureEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudCaptureEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUtm

`func (o *CloudCaptureEvent) GetUtm() CloudUTM`

GetUtm returns the Utm field if non-nil, zero value otherwise.

### GetUtmOk

`func (o *CloudCaptureEvent) GetUtmOk() (*CloudUTM, bool)`

GetUtmOk returns a tuple with the Utm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtm

`func (o *CloudCaptureEvent) SetUtm(v CloudUTM)`

SetUtm sets Utm field to given value.

### HasUtm

`func (o *CloudCaptureEvent) HasUtm() bool`

HasUtm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


