# O11yO11yDomainGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description describes the field, when known. | [optional] 
**FieldContext** | Pointer to **string** | FieldContext says which plane the field lives on, e.g. attribute, resource, span. | [optional] 
**FieldDataType** | Pointer to **string** | FieldDataType is the field&#39;s type: string, int64, float64 or bool. | [optional] 
**Name** | Pointer to **string** | Name is the field&#39;s name. Required. | [optional] 
**Signal** | Pointer to **string** | Signal is the telemetry signal the field belongs to, e.g. traces. | [optional] 
**Unit** | Pointer to **string** | Unit is the field&#39;s unit, when known. | [optional] 

## Methods

### NewO11yO11yDomainGroupBy

`func NewO11yO11yDomainGroupBy() *O11yO11yDomainGroupBy`

NewO11yO11yDomainGroupBy instantiates a new O11yO11yDomainGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDomainGroupByWithDefaults

`func NewO11yO11yDomainGroupByWithDefaults() *O11yO11yDomainGroupBy`

NewO11yO11yDomainGroupByWithDefaults instantiates a new O11yO11yDomainGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yDomainGroupBy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yDomainGroupBy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yDomainGroupBy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yDomainGroupBy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldContext

`func (o *O11yO11yDomainGroupBy) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *O11yO11yDomainGroupBy) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *O11yO11yDomainGroupBy) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *O11yO11yDomainGroupBy) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldDataType

`func (o *O11yO11yDomainGroupBy) GetFieldDataType() string`

GetFieldDataType returns the FieldDataType field if non-nil, zero value otherwise.

### GetFieldDataTypeOk

`func (o *O11yO11yDomainGroupBy) GetFieldDataTypeOk() (*string, bool)`

GetFieldDataTypeOk returns a tuple with the FieldDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDataType

`func (o *O11yO11yDomainGroupBy) SetFieldDataType(v string)`

SetFieldDataType sets FieldDataType field to given value.

### HasFieldDataType

`func (o *O11yO11yDomainGroupBy) HasFieldDataType() bool`

HasFieldDataType returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDomainGroupBy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDomainGroupBy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDomainGroupBy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDomainGroupBy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSignal

`func (o *O11yO11yDomainGroupBy) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *O11yO11yDomainGroupBy) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *O11yO11yDomainGroupBy) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *O11yO11yDomainGroupBy) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yDomainGroupBy) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yDomainGroupBy) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yDomainGroupBy) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yDomainGroupBy) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


