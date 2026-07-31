# CloudAdminSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credited** | Pointer to **int32** | Credited is how many have paid both bonuses. | [optional] 
**GrantedCents** | Pointer to **int32** | GrantedCents is the promo credit granted across BOTH sides of every referral, in USD cents — the program&#39;s total liability to date. | [optional] 
**Qualified** | Pointer to **int32** | Qualified is how many have qualified but are not yet credited. | [optional] 
**Signup** | Pointer to **int32** | Signup is how many are recorded but not yet qualified. | [optional] 
**Total** | Pointer to **int32** | Total is every referral in the ledger. | [optional] 

## Methods

### NewCloudAdminSummary

`func NewCloudAdminSummary() *CloudAdminSummary`

NewCloudAdminSummary instantiates a new CloudAdminSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminSummaryWithDefaults

`func NewCloudAdminSummaryWithDefaults() *CloudAdminSummary`

NewCloudAdminSummaryWithDefaults instantiates a new CloudAdminSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredited

`func (o *CloudAdminSummary) GetCredited() int32`

GetCredited returns the Credited field if non-nil, zero value otherwise.

### GetCreditedOk

`func (o *CloudAdminSummary) GetCreditedOk() (*int32, bool)`

GetCreditedOk returns a tuple with the Credited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredited

`func (o *CloudAdminSummary) SetCredited(v int32)`

SetCredited sets Credited field to given value.

### HasCredited

`func (o *CloudAdminSummary) HasCredited() bool`

HasCredited returns a boolean if a field has been set.

### GetGrantedCents

`func (o *CloudAdminSummary) GetGrantedCents() int32`

GetGrantedCents returns the GrantedCents field if non-nil, zero value otherwise.

### GetGrantedCentsOk

`func (o *CloudAdminSummary) GetGrantedCentsOk() (*int32, bool)`

GetGrantedCentsOk returns a tuple with the GrantedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedCents

`func (o *CloudAdminSummary) SetGrantedCents(v int32)`

SetGrantedCents sets GrantedCents field to given value.

### HasGrantedCents

`func (o *CloudAdminSummary) HasGrantedCents() bool`

HasGrantedCents returns a boolean if a field has been set.

### GetQualified

`func (o *CloudAdminSummary) GetQualified() int32`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *CloudAdminSummary) GetQualifiedOk() (*int32, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *CloudAdminSummary) SetQualified(v int32)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *CloudAdminSummary) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSignup

`func (o *CloudAdminSummary) GetSignup() int32`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *CloudAdminSummary) GetSignupOk() (*int32, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *CloudAdminSummary) SetSignup(v int32)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *CloudAdminSummary) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTotal

`func (o *CloudAdminSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudAdminSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudAdminSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudAdminSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


