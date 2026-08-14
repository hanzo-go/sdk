# HealthLens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available reports whether that table exists in the warehouse right now. | [optional] 
**Table** | Pointer to **string** | Table is the fully-qualified warehouse table the lens reads. | [optional] 

## Methods

### NewHealthLens

`func NewHealthLens() *HealthLens`

NewHealthLens instantiates a new HealthLens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthLensWithDefaults

`func NewHealthLensWithDefaults() *HealthLens`

NewHealthLensWithDefaults instantiates a new HealthLens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *HealthLens) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *HealthLens) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *HealthLens) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *HealthLens) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetTable

`func (o *HealthLens) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *HealthLens) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *HealthLens) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *HealthLens) HasTable() bool`

HasTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


