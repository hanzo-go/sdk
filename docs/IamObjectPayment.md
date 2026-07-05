# IamObjectPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**InvoiceRemark** | Pointer to **string** |  | [optional] 
**InvoiceTaxId** | Pointer to **string** |  | [optional] 
**InvoiceTitle** | Pointer to **string** |  | [optional] 
**InvoiceType** | Pointer to **string** |  | [optional] 
**InvoiceUrl** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**OrderObj** | Pointer to [**IamObjectOrder**](IamObjectOrder.md) |  | [optional] 
**OutOrderId** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PayUrl** | Pointer to **string** |  | [optional] 
**PersonEmail** | Pointer to **string** |  | [optional] 
**PersonIdCard** | Pointer to **string** |  | [optional] 
**PersonName** | Pointer to **string** |  | [optional] 
**PersonPhone** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**ProductsDisplayName** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**IamPpPaymentState**](IamPpPaymentState.md) |  | [optional] 
**SuccessUrl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectPayment

`func NewIamObjectPayment() *IamObjectPayment`

NewIamObjectPayment instantiates a new IamObjectPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectPaymentWithDefaults

`func NewIamObjectPaymentWithDefaults() *IamObjectPayment`

NewIamObjectPaymentWithDefaults instantiates a new IamObjectPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamObjectPayment) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectPayment) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectPayment) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectPayment) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamObjectPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamObjectPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamObjectPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamObjectPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDetail

`func (o *IamObjectPayment) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *IamObjectPayment) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *IamObjectPayment) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *IamObjectPayment) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectPayment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectPayment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectPayment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectPayment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInvoiceRemark

`func (o *IamObjectPayment) GetInvoiceRemark() string`

GetInvoiceRemark returns the InvoiceRemark field if non-nil, zero value otherwise.

### GetInvoiceRemarkOk

`func (o *IamObjectPayment) GetInvoiceRemarkOk() (*string, bool)`

GetInvoiceRemarkOk returns a tuple with the InvoiceRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRemark

`func (o *IamObjectPayment) SetInvoiceRemark(v string)`

SetInvoiceRemark sets InvoiceRemark field to given value.

### HasInvoiceRemark

`func (o *IamObjectPayment) HasInvoiceRemark() bool`

HasInvoiceRemark returns a boolean if a field has been set.

### GetInvoiceTaxId

`func (o *IamObjectPayment) GetInvoiceTaxId() string`

GetInvoiceTaxId returns the InvoiceTaxId field if non-nil, zero value otherwise.

### GetInvoiceTaxIdOk

`func (o *IamObjectPayment) GetInvoiceTaxIdOk() (*string, bool)`

GetInvoiceTaxIdOk returns a tuple with the InvoiceTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaxId

`func (o *IamObjectPayment) SetInvoiceTaxId(v string)`

SetInvoiceTaxId sets InvoiceTaxId field to given value.

### HasInvoiceTaxId

`func (o *IamObjectPayment) HasInvoiceTaxId() bool`

HasInvoiceTaxId returns a boolean if a field has been set.

### GetInvoiceTitle

`func (o *IamObjectPayment) GetInvoiceTitle() string`

GetInvoiceTitle returns the InvoiceTitle field if non-nil, zero value otherwise.

### GetInvoiceTitleOk

`func (o *IamObjectPayment) GetInvoiceTitleOk() (*string, bool)`

GetInvoiceTitleOk returns a tuple with the InvoiceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTitle

`func (o *IamObjectPayment) SetInvoiceTitle(v string)`

SetInvoiceTitle sets InvoiceTitle field to given value.

### HasInvoiceTitle

`func (o *IamObjectPayment) HasInvoiceTitle() bool`

HasInvoiceTitle returns a boolean if a field has been set.

### GetInvoiceType

`func (o *IamObjectPayment) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *IamObjectPayment) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *IamObjectPayment) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *IamObjectPayment) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetInvoiceUrl

`func (o *IamObjectPayment) GetInvoiceUrl() string`

GetInvoiceUrl returns the InvoiceUrl field if non-nil, zero value otherwise.

### GetInvoiceUrlOk

`func (o *IamObjectPayment) GetInvoiceUrlOk() (*string, bool)`

GetInvoiceUrlOk returns a tuple with the InvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrl

`func (o *IamObjectPayment) SetInvoiceUrl(v string)`

SetInvoiceUrl sets InvoiceUrl field to given value.

### HasInvoiceUrl

`func (o *IamObjectPayment) HasInvoiceUrl() bool`

HasInvoiceUrl returns a boolean if a field has been set.

### GetMessage

`func (o *IamObjectPayment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IamObjectPayment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IamObjectPayment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IamObjectPayment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *IamObjectPayment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectPayment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectPayment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectPayment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *IamObjectPayment) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *IamObjectPayment) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *IamObjectPayment) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *IamObjectPayment) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderObj

`func (o *IamObjectPayment) GetOrderObj() IamObjectOrder`

GetOrderObj returns the OrderObj field if non-nil, zero value otherwise.

### GetOrderObjOk

`func (o *IamObjectPayment) GetOrderObjOk() (*IamObjectOrder, bool)`

GetOrderObjOk returns a tuple with the OrderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderObj

`func (o *IamObjectPayment) SetOrderObj(v IamObjectOrder)`

SetOrderObj sets OrderObj field to given value.

### HasOrderObj

`func (o *IamObjectPayment) HasOrderObj() bool`

HasOrderObj returns a boolean if a field has been set.

### GetOutOrderId

`func (o *IamObjectPayment) GetOutOrderId() string`

GetOutOrderId returns the OutOrderId field if non-nil, zero value otherwise.

### GetOutOrderIdOk

`func (o *IamObjectPayment) GetOutOrderIdOk() (*string, bool)`

GetOutOrderIdOk returns a tuple with the OutOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOrderId

`func (o *IamObjectPayment) SetOutOrderId(v string)`

SetOutOrderId sets OutOrderId field to given value.

### HasOutOrderId

`func (o *IamObjectPayment) HasOutOrderId() bool`

HasOutOrderId returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectPayment) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectPayment) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectPayment) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectPayment) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPayUrl

`func (o *IamObjectPayment) GetPayUrl() string`

GetPayUrl returns the PayUrl field if non-nil, zero value otherwise.

### GetPayUrlOk

`func (o *IamObjectPayment) GetPayUrlOk() (*string, bool)`

GetPayUrlOk returns a tuple with the PayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayUrl

`func (o *IamObjectPayment) SetPayUrl(v string)`

SetPayUrl sets PayUrl field to given value.

### HasPayUrl

`func (o *IamObjectPayment) HasPayUrl() bool`

HasPayUrl returns a boolean if a field has been set.

### GetPersonEmail

`func (o *IamObjectPayment) GetPersonEmail() string`

GetPersonEmail returns the PersonEmail field if non-nil, zero value otherwise.

### GetPersonEmailOk

`func (o *IamObjectPayment) GetPersonEmailOk() (*string, bool)`

GetPersonEmailOk returns a tuple with the PersonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonEmail

`func (o *IamObjectPayment) SetPersonEmail(v string)`

SetPersonEmail sets PersonEmail field to given value.

### HasPersonEmail

`func (o *IamObjectPayment) HasPersonEmail() bool`

HasPersonEmail returns a boolean if a field has been set.

### GetPersonIdCard

`func (o *IamObjectPayment) GetPersonIdCard() string`

GetPersonIdCard returns the PersonIdCard field if non-nil, zero value otherwise.

### GetPersonIdCardOk

`func (o *IamObjectPayment) GetPersonIdCardOk() (*string, bool)`

GetPersonIdCardOk returns a tuple with the PersonIdCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIdCard

`func (o *IamObjectPayment) SetPersonIdCard(v string)`

SetPersonIdCard sets PersonIdCard field to given value.

### HasPersonIdCard

`func (o *IamObjectPayment) HasPersonIdCard() bool`

HasPersonIdCard returns a boolean if a field has been set.

### GetPersonName

`func (o *IamObjectPayment) GetPersonName() string`

GetPersonName returns the PersonName field if non-nil, zero value otherwise.

### GetPersonNameOk

`func (o *IamObjectPayment) GetPersonNameOk() (*string, bool)`

GetPersonNameOk returns a tuple with the PersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonName

`func (o *IamObjectPayment) SetPersonName(v string)`

SetPersonName sets PersonName field to given value.

### HasPersonName

`func (o *IamObjectPayment) HasPersonName() bool`

HasPersonName returns a boolean if a field has been set.

### GetPersonPhone

`func (o *IamObjectPayment) GetPersonPhone() string`

GetPersonPhone returns the PersonPhone field if non-nil, zero value otherwise.

### GetPersonPhoneOk

`func (o *IamObjectPayment) GetPersonPhoneOk() (*string, bool)`

GetPersonPhoneOk returns a tuple with the PersonPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonPhone

`func (o *IamObjectPayment) SetPersonPhone(v string)`

SetPersonPhone sets PersonPhone field to given value.

### HasPersonPhone

`func (o *IamObjectPayment) HasPersonPhone() bool`

HasPersonPhone returns a boolean if a field has been set.

### GetPrice

`func (o *IamObjectPayment) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IamObjectPayment) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IamObjectPayment) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IamObjectPayment) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProducts

`func (o *IamObjectPayment) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *IamObjectPayment) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *IamObjectPayment) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *IamObjectPayment) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetProductsDisplayName

`func (o *IamObjectPayment) GetProductsDisplayName() string`

GetProductsDisplayName returns the ProductsDisplayName field if non-nil, zero value otherwise.

### GetProductsDisplayNameOk

`func (o *IamObjectPayment) GetProductsDisplayNameOk() (*string, bool)`

GetProductsDisplayNameOk returns a tuple with the ProductsDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsDisplayName

`func (o *IamObjectPayment) SetProductsDisplayName(v string)`

SetProductsDisplayName sets ProductsDisplayName field to given value.

### HasProductsDisplayName

`func (o *IamObjectPayment) HasProductsDisplayName() bool`

HasProductsDisplayName returns a boolean if a field has been set.

### GetProvider

`func (o *IamObjectPayment) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamObjectPayment) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamObjectPayment) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamObjectPayment) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *IamObjectPayment) GetState() IamPpPaymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectPayment) GetStateOk() (*IamPpPaymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectPayment) SetState(v IamPpPaymentState)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectPayment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *IamObjectPayment) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *IamObjectPayment) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *IamObjectPayment) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *IamObjectPayment) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetType

`func (o *IamObjectPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamObjectPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamObjectPayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamObjectPayment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectPayment) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectPayment) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectPayment) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectPayment) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


