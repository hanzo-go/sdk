# O11yO11yApdexSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeStatusCodes** | Pointer to **string** | ExcludeStatusCodes are status codes excluded from the score, comma separated. | [optional] 
**Id** | Pointer to **string** | ID is the settings row&#39;s id. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the settings belong to. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service. | [optional] 
**Threshold** | Pointer to **float64** | Threshold is the satisfied-response time in seconds. | [optional] 

## Methods

### NewO11yO11yApdexSettings

`func NewO11yO11yApdexSettings() *O11yO11yApdexSettings`

NewO11yO11yApdexSettings instantiates a new O11yO11yApdexSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yApdexSettingsWithDefaults

`func NewO11yO11yApdexSettingsWithDefaults() *O11yO11yApdexSettings`

NewO11yO11yApdexSettingsWithDefaults instantiates a new O11yO11yApdexSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeStatusCodes

`func (o *O11yO11yApdexSettings) GetExcludeStatusCodes() string`

GetExcludeStatusCodes returns the ExcludeStatusCodes field if non-nil, zero value otherwise.

### GetExcludeStatusCodesOk

`func (o *O11yO11yApdexSettings) GetExcludeStatusCodesOk() (*string, bool)`

GetExcludeStatusCodesOk returns a tuple with the ExcludeStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeStatusCodes

`func (o *O11yO11yApdexSettings) SetExcludeStatusCodes(v string)`

SetExcludeStatusCodes sets ExcludeStatusCodes field to given value.

### HasExcludeStatusCodes

`func (o *O11yO11yApdexSettings) HasExcludeStatusCodes() bool`

HasExcludeStatusCodes returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yApdexSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yApdexSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yApdexSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yApdexSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yApdexSettings) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yApdexSettings) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yApdexSettings) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yApdexSettings) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yApdexSettings) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yApdexSettings) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yApdexSettings) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yApdexSettings) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetThreshold

`func (o *O11yO11yApdexSettings) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *O11yO11yApdexSettings) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *O11yO11yApdexSettings) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *O11yO11yApdexSettings) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


