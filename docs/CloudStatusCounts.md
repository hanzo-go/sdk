# CloudStatusCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credited** | Pointer to **int32** | Credited is how many referrals have paid both bonuses. | [optional] 
**Qualified** | Pointer to **int32** | Qualified is how many referees have spent but are not yet credited. | [optional] 
**Signup** | Pointer to **int32** | Signup is how many referees have signed up but not yet spent. | [optional] 
**Total** | Pointer to **int32** | Total is every referral this org has made. | [optional] 

## Methods

### NewCloudStatusCounts

`func NewCloudStatusCounts() *CloudStatusCounts`

NewCloudStatusCounts instantiates a new CloudStatusCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStatusCountsWithDefaults

`func NewCloudStatusCountsWithDefaults() *CloudStatusCounts`

NewCloudStatusCountsWithDefaults instantiates a new CloudStatusCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredited

`func (o *CloudStatusCounts) GetCredited() int32`

GetCredited returns the Credited field if non-nil, zero value otherwise.

### GetCreditedOk

`func (o *CloudStatusCounts) GetCreditedOk() (*int32, bool)`

GetCreditedOk returns a tuple with the Credited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredited

`func (o *CloudStatusCounts) SetCredited(v int32)`

SetCredited sets Credited field to given value.

### HasCredited

`func (o *CloudStatusCounts) HasCredited() bool`

HasCredited returns a boolean if a field has been set.

### GetQualified

`func (o *CloudStatusCounts) GetQualified() int32`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *CloudStatusCounts) GetQualifiedOk() (*int32, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *CloudStatusCounts) SetQualified(v int32)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *CloudStatusCounts) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSignup

`func (o *CloudStatusCounts) GetSignup() int32`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *CloudStatusCounts) GetSignupOk() (*int32, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *CloudStatusCounts) SetSignup(v int32)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *CloudStatusCounts) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTotal

`func (o *CloudStatusCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudStatusCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudStatusCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudStatusCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


