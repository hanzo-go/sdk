# ObserveStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**LatencyMs** | Pointer to **float32** | Health-probe round trip; -1 when the probe did not complete. | [optional] 
**HttpCode** | Pointer to **int32** | Health-probe status code (0 &#x3D; no response). | [optional] 
**ScrapeUp** | Pointer to **NullableInt32** | VictoriaMetrics up{service}: 1/0, or null when no scrape target. | [optional] 
**ProbeUrl** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | Which signal set &#x60;up&#x60;. | [optional] 
**CheckedAt** | Pointer to **string** | RFC3339 timestamp (UTC). | [optional] 

## Methods

### NewObserveStatusResponse

`func NewObserveStatusResponse() *ObserveStatusResponse`

NewObserveStatusResponse instantiates a new ObserveStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveStatusResponseWithDefaults

`func NewObserveStatusResponseWithDefaults() *ObserveStatusResponse`

NewObserveStatusResponseWithDefaults instantiates a new ObserveStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ObserveStatusResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ObserveStatusResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ObserveStatusResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ObserveStatusResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUp

`func (o *ObserveStatusResponse) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ObserveStatusResponse) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ObserveStatusResponse) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *ObserveStatusResponse) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetLatencyMs

`func (o *ObserveStatusResponse) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *ObserveStatusResponse) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *ObserveStatusResponse) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *ObserveStatusResponse) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetHttpCode

`func (o *ObserveStatusResponse) GetHttpCode() int32`

GetHttpCode returns the HttpCode field if non-nil, zero value otherwise.

### GetHttpCodeOk

`func (o *ObserveStatusResponse) GetHttpCodeOk() (*int32, bool)`

GetHttpCodeOk returns a tuple with the HttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCode

`func (o *ObserveStatusResponse) SetHttpCode(v int32)`

SetHttpCode sets HttpCode field to given value.

### HasHttpCode

`func (o *ObserveStatusResponse) HasHttpCode() bool`

HasHttpCode returns a boolean if a field has been set.

### GetScrapeUp

`func (o *ObserveStatusResponse) GetScrapeUp() int32`

GetScrapeUp returns the ScrapeUp field if non-nil, zero value otherwise.

### GetScrapeUpOk

`func (o *ObserveStatusResponse) GetScrapeUpOk() (*int32, bool)`

GetScrapeUpOk returns a tuple with the ScrapeUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeUp

`func (o *ObserveStatusResponse) SetScrapeUp(v int32)`

SetScrapeUp sets ScrapeUp field to given value.

### HasScrapeUp

`func (o *ObserveStatusResponse) HasScrapeUp() bool`

HasScrapeUp returns a boolean if a field has been set.

### SetScrapeUpNil

`func (o *ObserveStatusResponse) SetScrapeUpNil(b bool)`

 SetScrapeUpNil sets the value for ScrapeUp to be an explicit nil

### UnsetScrapeUp
`func (o *ObserveStatusResponse) UnsetScrapeUp()`

UnsetScrapeUp ensures that no value is present for ScrapeUp, not even an explicit nil
### GetProbeUrl

`func (o *ObserveStatusResponse) GetProbeUrl() string`

GetProbeUrl returns the ProbeUrl field if non-nil, zero value otherwise.

### GetProbeUrlOk

`func (o *ObserveStatusResponse) GetProbeUrlOk() (*string, bool)`

GetProbeUrlOk returns a tuple with the ProbeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeUrl

`func (o *ObserveStatusResponse) SetProbeUrl(v string)`

SetProbeUrl sets ProbeUrl field to given value.

### HasProbeUrl

`func (o *ObserveStatusResponse) HasProbeUrl() bool`

HasProbeUrl returns a boolean if a field has been set.

### GetSource

`func (o *ObserveStatusResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ObserveStatusResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ObserveStatusResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ObserveStatusResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCheckedAt

`func (o *ObserveStatusResponse) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *ObserveStatusResponse) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *ObserveStatusResponse) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *ObserveStatusResponse) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


