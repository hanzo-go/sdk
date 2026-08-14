# O11yO11yOrgStatsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** | Data are the statistics, keyed by the reporter&#39;s own counter names. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yOrgStatsOut

`func NewO11yO11yOrgStatsOut() *O11yO11yOrgStatsOut`

NewO11yO11yOrgStatsOut instantiates a new O11yO11yOrgStatsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOrgStatsOutWithDefaults

`func NewO11yO11yOrgStatsOutWithDefaults() *O11yO11yOrgStatsOut`

NewO11yO11yOrgStatsOutWithDefaults instantiates a new O11yO11yOrgStatsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yOrgStatsOut) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yOrgStatsOut) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yOrgStatsOut) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yOrgStatsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yOrgStatsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yOrgStatsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yOrgStatsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yOrgStatsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


