# AdminSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualified** | Pointer to **int32** | Qualified is how many referees have made metered spend. | [optional] 
**Signup** | Pointer to **int32** | Signup is how many are recorded but not yet qualified. | [optional] 
**Total** | Pointer to **int32** | Total is every referral in the directory. | [optional] 

## Methods

### NewAdminSummary

`func NewAdminSummary() *AdminSummary`

NewAdminSummary instantiates a new AdminSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminSummaryWithDefaults

`func NewAdminSummaryWithDefaults() *AdminSummary`

NewAdminSummaryWithDefaults instantiates a new AdminSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualified

`func (o *AdminSummary) GetQualified() int32`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *AdminSummary) GetQualifiedOk() (*int32, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *AdminSummary) SetQualified(v int32)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *AdminSummary) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSignup

`func (o *AdminSummary) GetSignup() int32`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *AdminSummary) GetSignupOk() (*int32, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *AdminSummary) SetSignup(v int32)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *AdminSummary) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTotal

`func (o *AdminSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdminSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdminSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AdminSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


