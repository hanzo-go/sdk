# CloudOppReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount is the deal value in minor units (cents) of Currency. | [optional] 
**CloseDate** | Pointer to **int32** | CloseDate is the expected close, as a unix second (0 &#x3D; unset). | [optional] 
**CompanyId** | Pointer to **string** | CompanyID links the deal to one of the org&#39;s companies. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code Amount is denominated in; empty defaults to USD. | [optional] 
**Id** | Pointer to **string** | ID names the opportunity to update and comes from the path. A create ignores it: the server mints the id. | [optional] 
**Name** | Pointer to **string** | Name is the deal name. Required. | [optional] 
**PointOfContactId** | Pointer to **string** | PointOfContact links the deal to one of the org&#39;s contacts. | [optional] 
**Stage** | Pointer to **string** | Stage is the pipeline stage: NEW, SCREENING, MEETING, PROPOSAL or CUSTOMER (case-insensitive). Empty defaults to NEW. | [optional] 

## Methods

### NewCloudOppReq

`func NewCloudOppReq() *CloudOppReq`

NewCloudOppReq instantiates a new CloudOppReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOppReqWithDefaults

`func NewCloudOppReqWithDefaults() *CloudOppReq`

NewCloudOppReqWithDefaults instantiates a new CloudOppReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CloudOppReq) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudOppReq) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudOppReq) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudOppReq) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCloseDate

`func (o *CloudOppReq) GetCloseDate() int32`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CloudOppReq) GetCloseDateOk() (*int32, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CloudOppReq) SetCloseDate(v int32)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CloudOppReq) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetCompanyId

`func (o *CloudOppReq) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CloudOppReq) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CloudOppReq) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CloudOppReq) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudOppReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudOppReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudOppReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudOppReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *CloudOppReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudOppReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudOppReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudOppReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudOppReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudOppReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudOppReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudOppReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPointOfContactId

`func (o *CloudOppReq) GetPointOfContactId() string`

GetPointOfContactId returns the PointOfContactId field if non-nil, zero value otherwise.

### GetPointOfContactIdOk

`func (o *CloudOppReq) GetPointOfContactIdOk() (*string, bool)`

GetPointOfContactIdOk returns a tuple with the PointOfContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContactId

`func (o *CloudOppReq) SetPointOfContactId(v string)`

SetPointOfContactId sets PointOfContactId field to given value.

### HasPointOfContactId

`func (o *CloudOppReq) HasPointOfContactId() bool`

HasPointOfContactId returns a boolean if a field has been set.

### GetStage

`func (o *CloudOppReq) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CloudOppReq) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CloudOppReq) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CloudOppReq) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


