# O11yO11yUpdatableQuickFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]O11yO11yFilterKey**](O11yO11yFilterKey.md) | Filters are the attributes to offer, in the order to offer them. | [optional] 
**Signal** | Pointer to **string** | Signal is the signal whose filters are being replaced. | [optional] 

## Methods

### NewO11yO11yUpdatableQuickFilters

`func NewO11yO11yUpdatableQuickFilters() *O11yO11yUpdatableQuickFilters`

NewO11yO11yUpdatableQuickFilters instantiates a new O11yO11yUpdatableQuickFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yUpdatableQuickFiltersWithDefaults

`func NewO11yO11yUpdatableQuickFiltersWithDefaults() *O11yO11yUpdatableQuickFilters`

NewO11yO11yUpdatableQuickFiltersWithDefaults instantiates a new O11yO11yUpdatableQuickFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *O11yO11yUpdatableQuickFilters) GetFilters() []O11yO11yFilterKey`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yO11yUpdatableQuickFilters) GetFiltersOk() (*[]O11yO11yFilterKey, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yO11yUpdatableQuickFilters) SetFilters(v []O11yO11yFilterKey)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yO11yUpdatableQuickFilters) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSignal

`func (o *O11yO11yUpdatableQuickFilters) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *O11yO11yUpdatableQuickFilters) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *O11yO11yUpdatableQuickFilters) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *O11yO11yUpdatableQuickFilters) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


