# StatusCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualified** | Pointer to **int32** | Qualified is how many referees have made metered spend. | [optional] 
**Signup** | Pointer to **int32** | Signup is how many referees have signed up but not yet spent. | [optional] 
**Total** | Pointer to **int32** | Total is every referral this org has made. | [optional] 

## Methods

### NewStatusCounts

`func NewStatusCounts() *StatusCounts`

NewStatusCounts instantiates a new StatusCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCountsWithDefaults

`func NewStatusCountsWithDefaults() *StatusCounts`

NewStatusCountsWithDefaults instantiates a new StatusCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualified

`func (o *StatusCounts) GetQualified() int32`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *StatusCounts) GetQualifiedOk() (*int32, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *StatusCounts) SetQualified(v int32)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *StatusCounts) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSignup

`func (o *StatusCounts) GetSignup() int32`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *StatusCounts) GetSignupOk() (*int32, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *StatusCounts) SetSignup(v int32)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *StatusCounts) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTotal

`func (o *StatusCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StatusCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StatusCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StatusCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


