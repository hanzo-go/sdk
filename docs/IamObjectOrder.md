# IamObjectOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**ProductInfos** | Pointer to [**[]IamObjectProductInfo**](IamObjectProductInfo.md) |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectOrder

`func NewIamObjectOrder() *IamObjectOrder`

NewIamObjectOrder instantiates a new IamObjectOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectOrderWithDefaults

`func NewIamObjectOrderWithDefaults() *IamObjectOrder`

NewIamObjectOrderWithDefaults instantiates a new IamObjectOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamObjectOrder) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectOrder) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectOrder) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectOrder) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamObjectOrder) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamObjectOrder) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamObjectOrder) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamObjectOrder) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectOrder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectOrder) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectOrder) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectOrder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMessage

`func (o *IamObjectOrder) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IamObjectOrder) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IamObjectOrder) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IamObjectOrder) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *IamObjectOrder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectOrder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectOrder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectOrder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectOrder) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectOrder) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectOrder) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectOrder) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPayment

`func (o *IamObjectOrder) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *IamObjectOrder) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *IamObjectOrder) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *IamObjectOrder) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPrice

`func (o *IamObjectOrder) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IamObjectOrder) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IamObjectOrder) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IamObjectOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProductInfos

`func (o *IamObjectOrder) GetProductInfos() []IamObjectProductInfo`

GetProductInfos returns the ProductInfos field if non-nil, zero value otherwise.

### GetProductInfosOk

`func (o *IamObjectOrder) GetProductInfosOk() (*[]IamObjectProductInfo, bool)`

GetProductInfosOk returns a tuple with the ProductInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfos

`func (o *IamObjectOrder) SetProductInfos(v []IamObjectProductInfo)`

SetProductInfos sets ProductInfos field to given value.

### HasProductInfos

`func (o *IamObjectOrder) HasProductInfos() bool`

HasProductInfos returns a boolean if a field has been set.

### GetProducts

`func (o *IamObjectOrder) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *IamObjectOrder) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *IamObjectOrder) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *IamObjectOrder) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetState

`func (o *IamObjectOrder) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectOrder) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectOrder) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectOrder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateTime

`func (o *IamObjectOrder) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *IamObjectOrder) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *IamObjectOrder) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *IamObjectOrder) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectOrder) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectOrder) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectOrder) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectOrder) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


