# TrustFrameworkRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edition** | Pointer to **string** | Edition is which edition this clause list is taken from. | [optional] 
**Framework** | Pointer to **string** | Framework is the framework id. | [optional] 
**Name** | Pointer to **string** | Name is the published standard&#39;s name. | [optional] 
**Publisher** | Pointer to **string** | Publisher is who publishes it. | [optional] 
**Total** | Pointer to **int64** | Total is how many clauses the standard publishes. | [optional] 
**Unit** | Pointer to **string** | Unit is what one clause is; Units is its plural. | [optional] 
**Units** | Pointer to **string** | Units is Unit&#39;s plural, carried so a caller renders \&quot;12 controls\&quot; without having to pluralise a word it does not know. | [optional] 

## Methods

### NewTrustFrameworkRow

`func NewTrustFrameworkRow() *TrustFrameworkRow`

NewTrustFrameworkRow instantiates a new TrustFrameworkRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustFrameworkRowWithDefaults

`func NewTrustFrameworkRowWithDefaults() *TrustFrameworkRow`

NewTrustFrameworkRowWithDefaults instantiates a new TrustFrameworkRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdition

`func (o *TrustFrameworkRow) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *TrustFrameworkRow) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *TrustFrameworkRow) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *TrustFrameworkRow) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetFramework

`func (o *TrustFrameworkRow) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TrustFrameworkRow) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TrustFrameworkRow) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TrustFrameworkRow) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *TrustFrameworkRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustFrameworkRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustFrameworkRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustFrameworkRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublisher

`func (o *TrustFrameworkRow) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *TrustFrameworkRow) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *TrustFrameworkRow) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *TrustFrameworkRow) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetTotal

`func (o *TrustFrameworkRow) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrustFrameworkRow) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrustFrameworkRow) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrustFrameworkRow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnit

`func (o *TrustFrameworkRow) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TrustFrameworkRow) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TrustFrameworkRow) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TrustFrameworkRow) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnits

`func (o *TrustFrameworkRow) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TrustFrameworkRow) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TrustFrameworkRow) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TrustFrameworkRow) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


