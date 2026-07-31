# DnsQueryAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TotalQueries** | Pointer to **int64** |  | [optional] 
**ByType** | Pointer to **map[string]int64** | Query count by record type | [optional] 
**ByResponseCode** | Pointer to **map[string]int64** | Query count by RCODE (NOERROR, NXDOMAIN, SERVFAIL, etc.) | [optional] 
**ByCountry** | Pointer to **map[string]int64** |  | [optional] 
**TopQueriedNames** | Pointer to [**[]DnsQueryAnalyticsTopQueriedNamesInner**](DnsQueryAnalyticsTopQueriedNamesInner.md) |  | [optional] 

## Methods

### NewDnsQueryAnalytics

`func NewDnsQueryAnalytics() *DnsQueryAnalytics`

NewDnsQueryAnalytics instantiates a new DnsQueryAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsQueryAnalyticsWithDefaults

`func NewDnsQueryAnalyticsWithDefaults() *DnsQueryAnalytics`

NewDnsQueryAnalyticsWithDefaults instantiates a new DnsQueryAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *DnsQueryAnalytics) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DnsQueryAnalytics) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DnsQueryAnalytics) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *DnsQueryAnalytics) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPeriod

`func (o *DnsQueryAnalytics) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DnsQueryAnalytics) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DnsQueryAnalytics) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DnsQueryAnalytics) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalQueries

`func (o *DnsQueryAnalytics) GetTotalQueries() int64`

GetTotalQueries returns the TotalQueries field if non-nil, zero value otherwise.

### GetTotalQueriesOk

`func (o *DnsQueryAnalytics) GetTotalQueriesOk() (*int64, bool)`

GetTotalQueriesOk returns a tuple with the TotalQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQueries

`func (o *DnsQueryAnalytics) SetTotalQueries(v int64)`

SetTotalQueries sets TotalQueries field to given value.

### HasTotalQueries

`func (o *DnsQueryAnalytics) HasTotalQueries() bool`

HasTotalQueries returns a boolean if a field has been set.

### GetByType

`func (o *DnsQueryAnalytics) GetByType() map[string]int64`

GetByType returns the ByType field if non-nil, zero value otherwise.

### GetByTypeOk

`func (o *DnsQueryAnalytics) GetByTypeOk() (*map[string]int64, bool)`

GetByTypeOk returns a tuple with the ByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByType

`func (o *DnsQueryAnalytics) SetByType(v map[string]int64)`

SetByType sets ByType field to given value.

### HasByType

`func (o *DnsQueryAnalytics) HasByType() bool`

HasByType returns a boolean if a field has been set.

### GetByResponseCode

`func (o *DnsQueryAnalytics) GetByResponseCode() map[string]int64`

GetByResponseCode returns the ByResponseCode field if non-nil, zero value otherwise.

### GetByResponseCodeOk

`func (o *DnsQueryAnalytics) GetByResponseCodeOk() (*map[string]int64, bool)`

GetByResponseCodeOk returns a tuple with the ByResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByResponseCode

`func (o *DnsQueryAnalytics) SetByResponseCode(v map[string]int64)`

SetByResponseCode sets ByResponseCode field to given value.

### HasByResponseCode

`func (o *DnsQueryAnalytics) HasByResponseCode() bool`

HasByResponseCode returns a boolean if a field has been set.

### GetByCountry

`func (o *DnsQueryAnalytics) GetByCountry() map[string]int64`

GetByCountry returns the ByCountry field if non-nil, zero value otherwise.

### GetByCountryOk

`func (o *DnsQueryAnalytics) GetByCountryOk() (*map[string]int64, bool)`

GetByCountryOk returns a tuple with the ByCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCountry

`func (o *DnsQueryAnalytics) SetByCountry(v map[string]int64)`

SetByCountry sets ByCountry field to given value.

### HasByCountry

`func (o *DnsQueryAnalytics) HasByCountry() bool`

HasByCountry returns a boolean if a field has been set.

### GetTopQueriedNames

`func (o *DnsQueryAnalytics) GetTopQueriedNames() []DnsQueryAnalyticsTopQueriedNamesInner`

GetTopQueriedNames returns the TopQueriedNames field if non-nil, zero value otherwise.

### GetTopQueriedNamesOk

`func (o *DnsQueryAnalytics) GetTopQueriedNamesOk() (*[]DnsQueryAnalyticsTopQueriedNamesInner, bool)`

GetTopQueriedNamesOk returns a tuple with the TopQueriedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopQueriedNames

`func (o *DnsQueryAnalytics) SetTopQueriedNames(v []DnsQueryAnalyticsTopQueriedNamesInner)`

SetTopQueriedNames sets TopQueriedNames field to given value.

### HasTopQueriedNames

`func (o *DnsQueryAnalytics) HasTopQueriedNames() bool`

HasTopQueriedNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


