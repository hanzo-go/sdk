# O11yMetricsComponentEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedComponent** | Pointer to [**O11yAssociatedComponent**](O11yAssociatedComponent.md) |  | [optional] 
**Metrics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewO11yMetricsComponentEntry

`func NewO11yMetricsComponentEntry() *O11yMetricsComponentEntry`

NewO11yMetricsComponentEntry instantiates a new O11yMetricsComponentEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricsComponentEntryWithDefaults

`func NewO11yMetricsComponentEntryWithDefaults() *O11yMetricsComponentEntry`

NewO11yMetricsComponentEntryWithDefaults instantiates a new O11yMetricsComponentEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedComponent

`func (o *O11yMetricsComponentEntry) GetAssociatedComponent() O11yAssociatedComponent`

GetAssociatedComponent returns the AssociatedComponent field if non-nil, zero value otherwise.

### GetAssociatedComponentOk

`func (o *O11yMetricsComponentEntry) GetAssociatedComponentOk() (*O11yAssociatedComponent, bool)`

GetAssociatedComponentOk returns a tuple with the AssociatedComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedComponent

`func (o *O11yMetricsComponentEntry) SetAssociatedComponent(v O11yAssociatedComponent)`

SetAssociatedComponent sets AssociatedComponent field to given value.

### HasAssociatedComponent

`func (o *O11yMetricsComponentEntry) HasAssociatedComponent() bool`

HasAssociatedComponent returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yMetricsComponentEntry) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yMetricsComponentEntry) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yMetricsComponentEntry) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yMetricsComponentEntry) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


