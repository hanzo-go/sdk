# FrameworkRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edition** | Pointer to **string** | Edition is which edition this clause list is taken from. | [optional] 
**Framework** | Pointer to **string** | Framework is the framework id. | [optional] 
**Name** | Pointer to **string** | Name is the published standard&#39;s name. | [optional] 
**Publisher** | Pointer to **string** | Publisher is who publishes it. | [optional] 
**Total** | Pointer to **int32** | Total is how many clauses the standard publishes. | [optional] 
**Unit** | Pointer to **string** | Unit is what one clause is; Units is its plural. | [optional] 
**Units** | Pointer to **string** | Units is Unit&#39;s plural, carried so a caller renders \&quot;12 controls\&quot; without having to pluralise a word it does not know. | [optional] 

## Methods

### NewFrameworkRow

`func NewFrameworkRow() *FrameworkRow`

NewFrameworkRow instantiates a new FrameworkRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkRowWithDefaults

`func NewFrameworkRowWithDefaults() *FrameworkRow`

NewFrameworkRowWithDefaults instantiates a new FrameworkRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdition

`func (o *FrameworkRow) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *FrameworkRow) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *FrameworkRow) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *FrameworkRow) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetFramework

`func (o *FrameworkRow) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *FrameworkRow) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *FrameworkRow) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *FrameworkRow) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *FrameworkRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameworkRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameworkRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrameworkRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublisher

`func (o *FrameworkRow) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *FrameworkRow) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *FrameworkRow) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *FrameworkRow) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetTotal

`func (o *FrameworkRow) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FrameworkRow) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FrameworkRow) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FrameworkRow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnit

`func (o *FrameworkRow) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FrameworkRow) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FrameworkRow) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *FrameworkRow) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnits

`func (o *FrameworkRow) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FrameworkRow) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FrameworkRow) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *FrameworkRow) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


