# CloudSetFlagIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is the switch itself: true enables the flag for every evaluation. | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**Key** | Pointer to **string** | Key is the switch to write, taken from the path (e.g. \&quot;waitlist.chat\&quot;). | [optional] 

## Methods

### NewCloudSetFlagIn

`func NewCloudSetFlagIn() *CloudSetFlagIn`

NewCloudSetFlagIn instantiates a new CloudSetFlagIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSetFlagInWithDefaults

`func NewCloudSetFlagInWithDefaults() *CloudSetFlagIn`

NewCloudSetFlagInWithDefaults instantiates a new CloudSetFlagIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudSetFlagIn) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudSetFlagIn) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudSetFlagIn) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudSetFlagIn) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *CloudSetFlagIn) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CloudSetFlagIn) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CloudSetFlagIn) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CloudSetFlagIn) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *CloudSetFlagIn) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *CloudSetFlagIn) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetKey

`func (o *CloudSetFlagIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudSetFlagIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudSetFlagIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudSetFlagIn) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


