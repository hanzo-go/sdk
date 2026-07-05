# IamObjectProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**DisableCustomRecharge** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**IsRecharge** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**ProviderObjs** | Pointer to [**[]IamObjectProvider**](IamObjectProvider.md) |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**RechargeOptions** | Pointer to **[]float64** |  | [optional] 
**Sold** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SuccessUrl** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectProduct

`func NewIamObjectProduct() *IamObjectProduct`

NewIamObjectProduct instantiates a new IamObjectProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectProductWithDefaults

`func NewIamObjectProductWithDefaults() *IamObjectProduct`

NewIamObjectProductWithDefaults instantiates a new IamObjectProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamObjectProduct) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectProduct) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectProduct) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectProduct) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamObjectProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamObjectProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamObjectProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamObjectProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *IamObjectProduct) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *IamObjectProduct) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *IamObjectProduct) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *IamObjectProduct) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDisableCustomRecharge

`func (o *IamObjectProduct) GetDisableCustomRecharge() bool`

GetDisableCustomRecharge returns the DisableCustomRecharge field if non-nil, zero value otherwise.

### GetDisableCustomRechargeOk

`func (o *IamObjectProduct) GetDisableCustomRechargeOk() (*bool, bool)`

GetDisableCustomRechargeOk returns a tuple with the DisableCustomRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCustomRecharge

`func (o *IamObjectProduct) SetDisableCustomRecharge(v bool)`

SetDisableCustomRecharge sets DisableCustomRecharge field to given value.

### HasDisableCustomRecharge

`func (o *IamObjectProduct) HasDisableCustomRecharge() bool`

HasDisableCustomRecharge returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectProduct) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectProduct) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectProduct) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectProduct) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImage

`func (o *IamObjectProduct) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *IamObjectProduct) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *IamObjectProduct) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *IamObjectProduct) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIsRecharge

`func (o *IamObjectProduct) GetIsRecharge() bool`

GetIsRecharge returns the IsRecharge field if non-nil, zero value otherwise.

### GetIsRechargeOk

`func (o *IamObjectProduct) GetIsRechargeOk() (*bool, bool)`

GetIsRechargeOk returns a tuple with the IsRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecharge

`func (o *IamObjectProduct) SetIsRecharge(v bool)`

SetIsRecharge sets IsRecharge field to given value.

### HasIsRecharge

`func (o *IamObjectProduct) HasIsRecharge() bool`

HasIsRecharge returns a boolean if a field has been set.

### GetName

`func (o *IamObjectProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectProduct) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectProduct) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectProduct) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectProduct) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *IamObjectProduct) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IamObjectProduct) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IamObjectProduct) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IamObjectProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProviderObjs

`func (o *IamObjectProduct) GetProviderObjs() []IamObjectProvider`

GetProviderObjs returns the ProviderObjs field if non-nil, zero value otherwise.

### GetProviderObjsOk

`func (o *IamObjectProduct) GetProviderObjsOk() (*[]IamObjectProvider, bool)`

GetProviderObjsOk returns a tuple with the ProviderObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObjs

`func (o *IamObjectProduct) SetProviderObjs(v []IamObjectProvider)`

SetProviderObjs sets ProviderObjs field to given value.

### HasProviderObjs

`func (o *IamObjectProduct) HasProviderObjs() bool`

HasProviderObjs returns a boolean if a field has been set.

### GetProviders

`func (o *IamObjectProduct) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamObjectProduct) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamObjectProduct) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamObjectProduct) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetQuantity

`func (o *IamObjectProduct) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IamObjectProduct) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IamObjectProduct) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IamObjectProduct) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRechargeOptions

`func (o *IamObjectProduct) GetRechargeOptions() []float64`

GetRechargeOptions returns the RechargeOptions field if non-nil, zero value otherwise.

### GetRechargeOptionsOk

`func (o *IamObjectProduct) GetRechargeOptionsOk() (*[]float64, bool)`

GetRechargeOptionsOk returns a tuple with the RechargeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeOptions

`func (o *IamObjectProduct) SetRechargeOptions(v []float64)`

SetRechargeOptions sets RechargeOptions field to given value.

### HasRechargeOptions

`func (o *IamObjectProduct) HasRechargeOptions() bool`

HasRechargeOptions returns a boolean if a field has been set.

### GetSold

`func (o *IamObjectProduct) GetSold() int64`

GetSold returns the Sold field if non-nil, zero value otherwise.

### GetSoldOk

`func (o *IamObjectProduct) GetSoldOk() (*int64, bool)`

GetSoldOk returns a tuple with the Sold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSold

`func (o *IamObjectProduct) SetSold(v int64)`

SetSold sets Sold field to given value.

### HasSold

`func (o *IamObjectProduct) HasSold() bool`

HasSold returns a boolean if a field has been set.

### GetState

`func (o *IamObjectProduct) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectProduct) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectProduct) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectProduct) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *IamObjectProduct) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *IamObjectProduct) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *IamObjectProduct) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *IamObjectProduct) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetTag

`func (o *IamObjectProduct) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamObjectProduct) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamObjectProduct) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamObjectProduct) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


