# O11yO11yDashboardPatchOp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | From is the source JSON Pointer for move and copy; ignored otherwise. | [optional] 
**Op** | Pointer to **string** | Op is the verb: add, remove, replace, move, copy or test. | [optional] 
**Path** | Pointer to **string** | Path is a JSON Pointer into the postable dashboard, e.g. /spec/display/name, /spec/panels/&lt;id&gt;, /tags/-. | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yO11yDashboardPatchOp

`func NewO11yO11yDashboardPatchOp() *O11yO11yDashboardPatchOp`

NewO11yO11yDashboardPatchOp instantiates a new O11yO11yDashboardPatchOp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardPatchOpWithDefaults

`func NewO11yO11yDashboardPatchOpWithDefaults() *O11yO11yDashboardPatchOp`

NewO11yO11yDashboardPatchOpWithDefaults instantiates a new O11yO11yDashboardPatchOp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *O11yO11yDashboardPatchOp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *O11yO11yDashboardPatchOp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *O11yO11yDashboardPatchOp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *O11yO11yDashboardPatchOp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yDashboardPatchOp) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yDashboardPatchOp) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yDashboardPatchOp) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yDashboardPatchOp) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *O11yO11yDashboardPatchOp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *O11yO11yDashboardPatchOp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *O11yO11yDashboardPatchOp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *O11yO11yDashboardPatchOp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yDashboardPatchOp) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yDashboardPatchOp) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yDashboardPatchOp) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yDashboardPatchOp) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *O11yO11yDashboardPatchOp) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *O11yO11yDashboardPatchOp) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


