# O11yHostFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**O11yFilter**](O11yFilter.md) |  | [optional] 
**FilterByStatus** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yHostFilter

`func NewO11yHostFilter() *O11yHostFilter`

NewO11yHostFilter instantiates a new O11yHostFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHostFilterWithDefaults

`func NewO11yHostFilterWithDefaults() *O11yHostFilter`

NewO11yHostFilterWithDefaults instantiates a new O11yHostFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *O11yHostFilter) GetFilter() O11yFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yHostFilter) GetFilterOk() (*O11yFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yHostFilter) SetFilter(v O11yFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yHostFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterByStatus

`func (o *O11yHostFilter) GetFilterByStatus() interface{}`

GetFilterByStatus returns the FilterByStatus field if non-nil, zero value otherwise.

### GetFilterByStatusOk

`func (o *O11yHostFilter) GetFilterByStatusOk() (*interface{}, bool)`

GetFilterByStatusOk returns a tuple with the FilterByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByStatus

`func (o *O11yHostFilter) SetFilterByStatus(v interface{})`

SetFilterByStatus sets FilterByStatus field to given value.

### HasFilterByStatus

`func (o *O11yHostFilter) HasFilterByStatus() bool`

HasFilterByStatus returns a boolean if a field has been set.

### SetFilterByStatusNil

`func (o *O11yHostFilter) SetFilterByStatusNil(b bool)`

 SetFilterByStatusNil sets the value for FilterByStatus to be an explicit nil

### UnsetFilterByStatus
`func (o *O11yHostFilter) UnsetFilterByStatus()`

UnsetFilterByStatus ensures that no value is present for FilterByStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


