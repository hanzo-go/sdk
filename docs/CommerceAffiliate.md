# CommerceAffiliate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Commission** | Pointer to [**CommerceAffiliateCommission**](CommerceAffiliateCommission.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceAffiliate

`func NewCommerceAffiliate() *CommerceAffiliate`

NewCommerceAffiliate instantiates a new CommerceAffiliate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceAffiliateWithDefaults

`func NewCommerceAffiliateWithDefaults() *CommerceAffiliate`

NewCommerceAffiliateWithDefaults instantiates a new CommerceAffiliate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceAffiliate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceAffiliate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceAffiliate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceAffiliate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *CommerceAffiliate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommerceAffiliate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommerceAffiliate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommerceAffiliate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *CommerceAffiliate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommerceAffiliate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommerceAffiliate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommerceAffiliate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CommerceAffiliate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommerceAffiliate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommerceAffiliate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CommerceAffiliate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCommission

`func (o *CommerceAffiliate) GetCommission() CommerceAffiliateCommission`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *CommerceAffiliate) GetCommissionOk() (*CommerceAffiliateCommission, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *CommerceAffiliate) SetCommission(v CommerceAffiliateCommission)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *CommerceAffiliate) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceAffiliate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceAffiliate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceAffiliate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceAffiliate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceAffiliate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceAffiliate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceAffiliate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceAffiliate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


