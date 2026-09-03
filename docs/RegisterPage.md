# RegisterPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is how many rows this page holds. | [optional] 
**Formations** | Pointer to [**[]Registration**](Registration.md) | Formations are the rows, newest activity first. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page size that was applied. | [optional] 
**Offset** | Pointer to **int64** | Offset is the offset that was applied. | [optional] 

## Methods

### NewRegisterPage

`func NewRegisterPage() *RegisterPage`

NewRegisterPage instantiates a new RegisterPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPageWithDefaults

`func NewRegisterPageWithDefaults() *RegisterPage`

NewRegisterPageWithDefaults instantiates a new RegisterPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RegisterPage) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RegisterPage) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RegisterPage) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RegisterPage) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFormations

`func (o *RegisterPage) GetFormations() []Registration`

GetFormations returns the Formations field if non-nil, zero value otherwise.

### GetFormationsOk

`func (o *RegisterPage) GetFormationsOk() (*[]Registration, bool)`

GetFormationsOk returns a tuple with the Formations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormations

`func (o *RegisterPage) SetFormations(v []Registration)`

SetFormations sets Formations field to given value.

### HasFormations

`func (o *RegisterPage) HasFormations() bool`

HasFormations returns a boolean if a field has been set.

### GetLimit

`func (o *RegisterPage) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RegisterPage) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RegisterPage) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RegisterPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *RegisterPage) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RegisterPage) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RegisterPage) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RegisterPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


