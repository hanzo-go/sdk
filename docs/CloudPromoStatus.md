# CloudPromoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to [**CloudPromo**](CloudPromo.md) |  | [optional] 
**Redeemed** | Pointer to **int32** | Redeemed is how many orgs have taken it, Remaining how many are left under the fleet-wide cap. | [optional] 
**Remaining** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudPromoStatus

`func NewCloudPromoStatus() *CloudPromoStatus`

NewCloudPromoStatus instantiates a new CloudPromoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPromoStatusWithDefaults

`func NewCloudPromoStatusWithDefaults() *CloudPromoStatus`

NewCloudPromoStatusWithDefaults instantiates a new CloudPromoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromo

`func (o *CloudPromoStatus) GetPromo() CloudPromo`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *CloudPromoStatus) GetPromoOk() (*CloudPromo, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *CloudPromoStatus) SetPromo(v CloudPromo)`

SetPromo sets Promo field to given value.

### HasPromo

`func (o *CloudPromoStatus) HasPromo() bool`

HasPromo returns a boolean if a field has been set.

### GetRedeemed

`func (o *CloudPromoStatus) GetRedeemed() int32`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *CloudPromoStatus) GetRedeemedOk() (*int32, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *CloudPromoStatus) SetRedeemed(v int32)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *CloudPromoStatus) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRemaining

`func (o *CloudPromoStatus) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *CloudPromoStatus) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *CloudPromoStatus) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *CloudPromoStatus) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


