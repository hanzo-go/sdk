# O11yO11ySignalFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]O11yO11yFilterKey**](O11yO11yFilterKey.md) | Filters are the attributes offered, in display order. | [optional] 
**Signal** | Pointer to **string** | Signal is the signal the filters belong to. | [optional] 

## Methods

### NewO11yO11ySignalFilters

`func NewO11yO11ySignalFilters() *O11yO11ySignalFilters`

NewO11yO11ySignalFilters instantiates a new O11yO11ySignalFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySignalFiltersWithDefaults

`func NewO11yO11ySignalFiltersWithDefaults() *O11yO11ySignalFilters`

NewO11yO11ySignalFiltersWithDefaults instantiates a new O11yO11ySignalFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *O11yO11ySignalFilters) GetFilters() []O11yO11yFilterKey`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yO11ySignalFilters) GetFiltersOk() (*[]O11yO11yFilterKey, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yO11ySignalFilters) SetFilters(v []O11yO11yFilterKey)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yO11ySignalFilters) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSignal

`func (o *O11yO11ySignalFilters) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *O11yO11ySignalFilters) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *O11yO11ySignalFilters) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *O11yO11ySignalFilters) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


