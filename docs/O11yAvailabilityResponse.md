# O11yAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**O11yAvailabilityResponseRange**](O11yAvailabilityResponseRange.md) |  | [optional] 
**Series** | Pointer to [**[]O11yAvailabilityPoint**](O11yAvailabilityPoint.md) | Series is the trend, oldest bucket first. | [optional] 
**Services** | Pointer to [**[]O11yServiceUp**](O11yServiceUp.md) | Services is the current inventory, sorted by name so two reads of an unchanged fleet are byte-identical. | [optional] 
**Total** | Pointer to **int32** | Total is how many services the prober currently watches. | [optional] 
**Up** | Pointer to **int32** | Up is how many services are up right now. | [optional] 

## Methods

### NewO11yAvailabilityResponse

`func NewO11yAvailabilityResponse() *O11yAvailabilityResponse`

NewO11yAvailabilityResponse instantiates a new O11yAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAvailabilityResponseWithDefaults

`func NewO11yAvailabilityResponseWithDefaults() *O11yAvailabilityResponse`

NewO11yAvailabilityResponseWithDefaults instantiates a new O11yAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *O11yAvailabilityResponse) GetRange() O11yAvailabilityResponseRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *O11yAvailabilityResponse) GetRangeOk() (*O11yAvailabilityResponseRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *O11yAvailabilityResponse) SetRange(v O11yAvailabilityResponseRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *O11yAvailabilityResponse) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *O11yAvailabilityResponse) GetSeries() []O11yAvailabilityPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *O11yAvailabilityResponse) GetSeriesOk() (*[]O11yAvailabilityPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *O11yAvailabilityResponse) SetSeries(v []O11yAvailabilityPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *O11yAvailabilityResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetServices

`func (o *O11yAvailabilityResponse) GetServices() []O11yServiceUp`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *O11yAvailabilityResponse) GetServicesOk() (*[]O11yServiceUp, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *O11yAvailabilityResponse) SetServices(v []O11yServiceUp)`

SetServices sets Services field to given value.

### HasServices

`func (o *O11yAvailabilityResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTotal

`func (o *O11yAvailabilityResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yAvailabilityResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yAvailabilityResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yAvailabilityResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUp

`func (o *O11yAvailabilityResponse) GetUp() int32`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *O11yAvailabilityResponse) GetUpOk() (*int32, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *O11yAvailabilityResponse) SetUp(v int32)`

SetUp sets Up field to given value.

### HasUp

`func (o *O11yAvailabilityResponse) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


