# InsightsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Engine names the engine serving the surface: hanzo-analytics. | [optional] 
**Ok** | Pointer to **bool** | OK is always true — reaching this route is the liveness fact it reports. | [optional] 
**Surface** | Pointer to **string** | Surface is the path prefix this status covers: /v1/insights. | [optional] 

## Methods

### NewInsightsStatus

`func NewInsightsStatus() *InsightsStatus`

NewInsightsStatus instantiates a new InsightsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsStatusWithDefaults

`func NewInsightsStatusWithDefaults() *InsightsStatus`

NewInsightsStatusWithDefaults instantiates a new InsightsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *InsightsStatus) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *InsightsStatus) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *InsightsStatus) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *InsightsStatus) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetOk

`func (o *InsightsStatus) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *InsightsStatus) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *InsightsStatus) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *InsightsStatus) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSurface

`func (o *InsightsStatus) GetSurface() string`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *InsightsStatus) GetSurfaceOk() (*string, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *InsightsStatus) SetSurface(v string)`

SetSurface sets Surface field to given value.

### HasSurface

`func (o *InsightsStatus) HasSurface() bool`

HasSurface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


