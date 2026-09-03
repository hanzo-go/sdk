# O11yO11yApdexSetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeStatusCodes** | Pointer to **string** | ExcludeStatusCodes are status codes excluded from the score, comma separated. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service the threshold applies to. | [optional] 
**Threshold** | Pointer to **float64** | Threshold is the satisfied-response time in seconds. | [optional] 

## Methods

### NewO11yO11yApdexSetIn

`func NewO11yO11yApdexSetIn() *O11yO11yApdexSetIn`

NewO11yO11yApdexSetIn instantiates a new O11yO11yApdexSetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yApdexSetInWithDefaults

`func NewO11yO11yApdexSetInWithDefaults() *O11yO11yApdexSetIn`

NewO11yO11yApdexSetInWithDefaults instantiates a new O11yO11yApdexSetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeStatusCodes

`func (o *O11yO11yApdexSetIn) GetExcludeStatusCodes() string`

GetExcludeStatusCodes returns the ExcludeStatusCodes field if non-nil, zero value otherwise.

### GetExcludeStatusCodesOk

`func (o *O11yO11yApdexSetIn) GetExcludeStatusCodesOk() (*string, bool)`

GetExcludeStatusCodesOk returns a tuple with the ExcludeStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeStatusCodes

`func (o *O11yO11yApdexSetIn) SetExcludeStatusCodes(v string)`

SetExcludeStatusCodes sets ExcludeStatusCodes field to given value.

### HasExcludeStatusCodes

`func (o *O11yO11yApdexSetIn) HasExcludeStatusCodes() bool`

HasExcludeStatusCodes returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yApdexSetIn) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yApdexSetIn) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yApdexSetIn) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yApdexSetIn) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetThreshold

`func (o *O11yO11yApdexSetIn) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *O11yO11yApdexSetIn) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *O11yO11yApdexSetIn) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *O11yO11yApdexSetIn) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


