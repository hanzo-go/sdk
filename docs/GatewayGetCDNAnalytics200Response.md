# GatewayGetCDNAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRequests** | Pointer to **int32** |  | [optional] 
**CacheHits** | Pointer to **int32** |  | [optional] 
**CacheMisses** | Pointer to **int32** |  | [optional] 
**HitRatio** | Pointer to **float32** |  | [optional] 
**BandwidthBytes** | Pointer to **int32** |  | [optional] 
**Timeseries** | Pointer to [**[]GatewayGetCDNAnalytics200ResponseTimeseriesInner**](GatewayGetCDNAnalytics200ResponseTimeseriesInner.md) |  | [optional] 

## Methods

### NewGatewayGetCDNAnalytics200Response

`func NewGatewayGetCDNAnalytics200Response() *GatewayGetCDNAnalytics200Response`

NewGatewayGetCDNAnalytics200Response instantiates a new GatewayGetCDNAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayGetCDNAnalytics200ResponseWithDefaults

`func NewGatewayGetCDNAnalytics200ResponseWithDefaults() *GatewayGetCDNAnalytics200Response`

NewGatewayGetCDNAnalytics200ResponseWithDefaults instantiates a new GatewayGetCDNAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRequests

`func (o *GatewayGetCDNAnalytics200Response) GetTotalRequests() int32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *GatewayGetCDNAnalytics200Response) GetTotalRequestsOk() (*int32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *GatewayGetCDNAnalytics200Response) SetTotalRequests(v int32)`

SetTotalRequests sets TotalRequests field to given value.

### HasTotalRequests

`func (o *GatewayGetCDNAnalytics200Response) HasTotalRequests() bool`

HasTotalRequests returns a boolean if a field has been set.

### GetCacheHits

`func (o *GatewayGetCDNAnalytics200Response) GetCacheHits() int32`

GetCacheHits returns the CacheHits field if non-nil, zero value otherwise.

### GetCacheHitsOk

`func (o *GatewayGetCDNAnalytics200Response) GetCacheHitsOk() (*int32, bool)`

GetCacheHitsOk returns a tuple with the CacheHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHits

`func (o *GatewayGetCDNAnalytics200Response) SetCacheHits(v int32)`

SetCacheHits sets CacheHits field to given value.

### HasCacheHits

`func (o *GatewayGetCDNAnalytics200Response) HasCacheHits() bool`

HasCacheHits returns a boolean if a field has been set.

### GetCacheMisses

`func (o *GatewayGetCDNAnalytics200Response) GetCacheMisses() int32`

GetCacheMisses returns the CacheMisses field if non-nil, zero value otherwise.

### GetCacheMissesOk

`func (o *GatewayGetCDNAnalytics200Response) GetCacheMissesOk() (*int32, bool)`

GetCacheMissesOk returns a tuple with the CacheMisses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMisses

`func (o *GatewayGetCDNAnalytics200Response) SetCacheMisses(v int32)`

SetCacheMisses sets CacheMisses field to given value.

### HasCacheMisses

`func (o *GatewayGetCDNAnalytics200Response) HasCacheMisses() bool`

HasCacheMisses returns a boolean if a field has been set.

### GetHitRatio

`func (o *GatewayGetCDNAnalytics200Response) GetHitRatio() float32`

GetHitRatio returns the HitRatio field if non-nil, zero value otherwise.

### GetHitRatioOk

`func (o *GatewayGetCDNAnalytics200Response) GetHitRatioOk() (*float32, bool)`

GetHitRatioOk returns a tuple with the HitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRatio

`func (o *GatewayGetCDNAnalytics200Response) SetHitRatio(v float32)`

SetHitRatio sets HitRatio field to given value.

### HasHitRatio

`func (o *GatewayGetCDNAnalytics200Response) HasHitRatio() bool`

HasHitRatio returns a boolean if a field has been set.

### GetBandwidthBytes

`func (o *GatewayGetCDNAnalytics200Response) GetBandwidthBytes() int32`

GetBandwidthBytes returns the BandwidthBytes field if non-nil, zero value otherwise.

### GetBandwidthBytesOk

`func (o *GatewayGetCDNAnalytics200Response) GetBandwidthBytesOk() (*int32, bool)`

GetBandwidthBytesOk returns a tuple with the BandwidthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthBytes

`func (o *GatewayGetCDNAnalytics200Response) SetBandwidthBytes(v int32)`

SetBandwidthBytes sets BandwidthBytes field to given value.

### HasBandwidthBytes

`func (o *GatewayGetCDNAnalytics200Response) HasBandwidthBytes() bool`

HasBandwidthBytes returns a boolean if a field has been set.

### GetTimeseries

`func (o *GatewayGetCDNAnalytics200Response) GetTimeseries() []GatewayGetCDNAnalytics200ResponseTimeseriesInner`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *GatewayGetCDNAnalytics200Response) GetTimeseriesOk() (*[]GatewayGetCDNAnalytics200ResponseTimeseriesInner, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *GatewayGetCDNAnalytics200Response) SetTimeseries(v []GatewayGetCDNAnalytics200ResponseTimeseriesInner)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *GatewayGetCDNAnalytics200Response) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


