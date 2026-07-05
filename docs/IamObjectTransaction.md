# IamObjectTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**IamObjectTransactionCategory**](IamObjectTransactionCategory.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectTransaction

`func NewIamObjectTransaction() *IamObjectTransaction`

NewIamObjectTransaction instantiates a new IamObjectTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectTransactionWithDefaults

`func NewIamObjectTransactionWithDefaults() *IamObjectTransaction`

NewIamObjectTransactionWithDefaults instantiates a new IamObjectTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IamObjectTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IamObjectTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IamObjectTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *IamObjectTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApplication

`func (o *IamObjectTransaction) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamObjectTransaction) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamObjectTransaction) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamObjectTransaction) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCategory

`func (o *IamObjectTransaction) GetCategory() IamObjectTransactionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IamObjectTransaction) GetCategoryOk() (*IamObjectTransactionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IamObjectTransaction) SetCategory(v IamObjectTransactionCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IamObjectTransaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectTransaction) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectTransaction) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectTransaction) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectTransaction) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *IamObjectTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IamObjectTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IamObjectTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IamObjectTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectTransaction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectTransaction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectTransaction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectTransaction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *IamObjectTransaction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamObjectTransaction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamObjectTransaction) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamObjectTransaction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *IamObjectTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectTransaction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectTransaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectTransaction) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectTransaction) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectTransaction) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectTransaction) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPayment

`func (o *IamObjectTransaction) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *IamObjectTransaction) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *IamObjectTransaction) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *IamObjectTransaction) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetProvider

`func (o *IamObjectTransaction) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamObjectTransaction) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamObjectTransaction) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamObjectTransaction) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *IamObjectTransaction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectTransaction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectTransaction) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectTransaction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubtype

`func (o *IamObjectTransaction) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *IamObjectTransaction) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *IamObjectTransaction) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *IamObjectTransaction) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTag

`func (o *IamObjectTransaction) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamObjectTransaction) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamObjectTransaction) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamObjectTransaction) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *IamObjectTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamObjectTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamObjectTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamObjectTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectTransaction) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectTransaction) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectTransaction) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectTransaction) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


