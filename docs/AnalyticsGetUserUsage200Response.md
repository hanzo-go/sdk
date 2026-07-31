# AnalyticsGetUserUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteEventUsage** | Pointer to **int32** |  | [optional] 
**EventDataUsage** | Pointer to **int32** |  | [optional] 
**Websites** | Pointer to [**[]AnalyticsGetUserUsage200ResponseWebsitesInner**](AnalyticsGetUserUsage200ResponseWebsitesInner.md) |  | [optional] 

## Methods

### NewAnalyticsGetUserUsage200Response

`func NewAnalyticsGetUserUsage200Response() *AnalyticsGetUserUsage200Response`

NewAnalyticsGetUserUsage200Response instantiates a new AnalyticsGetUserUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsGetUserUsage200ResponseWithDefaults

`func NewAnalyticsGetUserUsage200ResponseWithDefaults() *AnalyticsGetUserUsage200Response`

NewAnalyticsGetUserUsage200ResponseWithDefaults instantiates a new AnalyticsGetUserUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteEventUsage

`func (o *AnalyticsGetUserUsage200Response) GetWebsiteEventUsage() int32`

GetWebsiteEventUsage returns the WebsiteEventUsage field if non-nil, zero value otherwise.

### GetWebsiteEventUsageOk

`func (o *AnalyticsGetUserUsage200Response) GetWebsiteEventUsageOk() (*int32, bool)`

GetWebsiteEventUsageOk returns a tuple with the WebsiteEventUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteEventUsage

`func (o *AnalyticsGetUserUsage200Response) SetWebsiteEventUsage(v int32)`

SetWebsiteEventUsage sets WebsiteEventUsage field to given value.

### HasWebsiteEventUsage

`func (o *AnalyticsGetUserUsage200Response) HasWebsiteEventUsage() bool`

HasWebsiteEventUsage returns a boolean if a field has been set.

### GetEventDataUsage

`func (o *AnalyticsGetUserUsage200Response) GetEventDataUsage() int32`

GetEventDataUsage returns the EventDataUsage field if non-nil, zero value otherwise.

### GetEventDataUsageOk

`func (o *AnalyticsGetUserUsage200Response) GetEventDataUsageOk() (*int32, bool)`

GetEventDataUsageOk returns a tuple with the EventDataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataUsage

`func (o *AnalyticsGetUserUsage200Response) SetEventDataUsage(v int32)`

SetEventDataUsage sets EventDataUsage field to given value.

### HasEventDataUsage

`func (o *AnalyticsGetUserUsage200Response) HasEventDataUsage() bool`

HasEventDataUsage returns a boolean if a field has been set.

### GetWebsites

`func (o *AnalyticsGetUserUsage200Response) GetWebsites() []AnalyticsGetUserUsage200ResponseWebsitesInner`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *AnalyticsGetUserUsage200Response) GetWebsitesOk() (*[]AnalyticsGetUserUsage200ResponseWebsitesInner, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *AnalyticsGetUserUsage200Response) SetWebsites(v []AnalyticsGetUserUsage200ResponseWebsitesInner)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *AnalyticsGetUserUsage200Response) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


