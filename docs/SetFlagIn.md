# SetFlagIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is the switch itself: true enables the flag for every evaluation. | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**Key** | Pointer to **string** | Key is the switch to write, taken from the path (e.g. \&quot;waitlist.chat\&quot;). | [optional] 

## Methods

### NewSetFlagIn

`func NewSetFlagIn() *SetFlagIn`

NewSetFlagIn instantiates a new SetFlagIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFlagInWithDefaults

`func NewSetFlagInWithDefaults() *SetFlagIn`

NewSetFlagInWithDefaults instantiates a new SetFlagIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SetFlagIn) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SetFlagIn) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SetFlagIn) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SetFlagIn) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *SetFlagIn) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SetFlagIn) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SetFlagIn) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SetFlagIn) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *SetFlagIn) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *SetFlagIn) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetKey

`func (o *SetFlagIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SetFlagIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SetFlagIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SetFlagIn) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


