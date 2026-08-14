# O11yO11yRoutePoliciesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yGettableRoutePolicy**](O11yGettableRoutePolicy.md) | Data holds the policies. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yRoutePoliciesOut

`func NewO11yO11yRoutePoliciesOut() *O11yO11yRoutePoliciesOut`

NewO11yO11yRoutePoliciesOut instantiates a new O11yO11yRoutePoliciesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoutePoliciesOutWithDefaults

`func NewO11yO11yRoutePoliciesOutWithDefaults() *O11yO11yRoutePoliciesOut`

NewO11yO11yRoutePoliciesOutWithDefaults instantiates a new O11yO11yRoutePoliciesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yRoutePoliciesOut) GetData() []O11yGettableRoutePolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yRoutePoliciesOut) GetDataOk() (*[]O11yGettableRoutePolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yRoutePoliciesOut) SetData(v []O11yGettableRoutePolicy)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yRoutePoliciesOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yRoutePoliciesOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yRoutePoliciesOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yRoutePoliciesOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yRoutePoliciesOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


