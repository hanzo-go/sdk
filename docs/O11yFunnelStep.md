# O11yFunnelStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | step description | [optional] 
**Filters** | Pointer to [**O11yFilterSet**](O11yFilterSet.md) |  | [optional] 
**HasErrors** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**LatencyPointer** | Pointer to **string** |  | [optional] 
**LatencyType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | step name | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**SpanName** | Pointer to **string** |  | [optional] 
**StepOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yFunnelStep

`func NewO11yFunnelStep() *O11yFunnelStep`

NewO11yFunnelStep instantiates a new O11yFunnelStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yFunnelStepWithDefaults

`func NewO11yFunnelStepWithDefaults() *O11yFunnelStep`

NewO11yFunnelStepWithDefaults instantiates a new O11yFunnelStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yFunnelStep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yFunnelStep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yFunnelStep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yFunnelStep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *O11yFunnelStep) GetFilters() O11yFilterSet`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yFunnelStep) GetFiltersOk() (*O11yFilterSet, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yFunnelStep) SetFilters(v O11yFilterSet)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yFunnelStep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetHasErrors

`func (o *O11yFunnelStep) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *O11yFunnelStep) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *O11yFunnelStep) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.

### HasHasErrors

`func (o *O11yFunnelStep) HasHasErrors() bool`

HasHasErrors returns a boolean if a field has been set.

### GetId

`func (o *O11yFunnelStep) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yFunnelStep) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yFunnelStep) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11yFunnelStep) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11yFunnelStep) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11yFunnelStep) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLatencyPointer

`func (o *O11yFunnelStep) GetLatencyPointer() string`

GetLatencyPointer returns the LatencyPointer field if non-nil, zero value otherwise.

### GetLatencyPointerOk

`func (o *O11yFunnelStep) GetLatencyPointerOk() (*string, bool)`

GetLatencyPointerOk returns a tuple with the LatencyPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyPointer

`func (o *O11yFunnelStep) SetLatencyPointer(v string)`

SetLatencyPointer sets LatencyPointer field to given value.

### HasLatencyPointer

`func (o *O11yFunnelStep) HasLatencyPointer() bool`

HasLatencyPointer returns a boolean if a field has been set.

### GetLatencyType

`func (o *O11yFunnelStep) GetLatencyType() string`

GetLatencyType returns the LatencyType field if non-nil, zero value otherwise.

### GetLatencyTypeOk

`func (o *O11yFunnelStep) GetLatencyTypeOk() (*string, bool)`

GetLatencyTypeOk returns a tuple with the LatencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyType

`func (o *O11yFunnelStep) SetLatencyType(v string)`

SetLatencyType sets LatencyType field to given value.

### HasLatencyType

`func (o *O11yFunnelStep) HasLatencyType() bool`

HasLatencyType returns a boolean if a field has been set.

### GetName

`func (o *O11yFunnelStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yFunnelStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yFunnelStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yFunnelStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yFunnelStep) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yFunnelStep) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yFunnelStep) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yFunnelStep) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSpanName

`func (o *O11yFunnelStep) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *O11yFunnelStep) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *O11yFunnelStep) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.

### HasSpanName

`func (o *O11yFunnelStep) HasSpanName() bool`

HasSpanName returns a boolean if a field has been set.

### GetStepOrder

`func (o *O11yFunnelStep) GetStepOrder() int64`

GetStepOrder returns the StepOrder field if non-nil, zero value otherwise.

### GetStepOrderOk

`func (o *O11yFunnelStep) GetStepOrderOk() (*int64, bool)`

GetStepOrderOk returns a tuple with the StepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepOrder

`func (o *O11yFunnelStep) SetStepOrder(v int64)`

SetStepOrder sets StepOrder field to given value.

### HasStepOrder

`func (o *O11yFunnelStep) HasStepOrder() bool`

HasStepOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


