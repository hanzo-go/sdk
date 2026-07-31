# CloudAccountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the ledger account address (\&quot;org:acme:wallet\&quot;, \&quot;fund:reserve\&quot;, …). | [optional] 
**BalanceCents** | Pointer to **int32** | BalanceCents is that account&#39;s signed balance in minor units. | [optional] 

## Methods

### NewCloudAccountView

`func NewCloudAccountView() *CloudAccountView`

NewCloudAccountView instantiates a new CloudAccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountViewWithDefaults

`func NewCloudAccountViewWithDefaults() *CloudAccountView`

NewCloudAccountViewWithDefaults instantiates a new CloudAccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CloudAccountView) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CloudAccountView) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CloudAccountView) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CloudAccountView) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalanceCents

`func (o *CloudAccountView) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *CloudAccountView) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *CloudAccountView) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *CloudAccountView) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


