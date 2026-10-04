# CaptableCaptableNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capital** | Pointer to **float64** | Capital is the principal the investor lent. | [optional] 
**ConversionCap** | Pointer to **float64** | ConversionCap is the valuation cap on conversion, if any. | [optional] 
**DiscountRate** | Pointer to **float64** | DiscountRate is the discount to the next round&#39;s price, if any. | [optional] 
**Id** | Pointer to **string** | ID is the note id. | [optional] 
**InterestRate** | Pointer to **float64** | InterestRate is the annual interest rate, if any. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the note was signed. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the note&#39;s shareable identifier, unique within the company. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the note&#39;s state, e.g. DRAFT or ACTIVE. | [optional] 
**Type** | Pointer to **string** | Type is the instrument kind, e.g. NOTE. | [optional] 

## Methods

### NewCaptableCaptableNote

`func NewCaptableCaptableNote() *CaptableCaptableNote`

NewCaptableCaptableNote instantiates a new CaptableCaptableNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableNoteWithDefaults

`func NewCaptableCaptableNoteWithDefaults() *CaptableCaptableNote`

NewCaptableCaptableNoteWithDefaults instantiates a new CaptableCaptableNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapital

`func (o *CaptableCaptableNote) GetCapital() float64`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *CaptableCaptableNote) GetCapitalOk() (*float64, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *CaptableCaptableNote) SetCapital(v float64)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *CaptableCaptableNote) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetConversionCap

`func (o *CaptableCaptableNote) GetConversionCap() float64`

GetConversionCap returns the ConversionCap field if non-nil, zero value otherwise.

### GetConversionCapOk

`func (o *CaptableCaptableNote) GetConversionCapOk() (*float64, bool)`

GetConversionCapOk returns a tuple with the ConversionCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionCap

`func (o *CaptableCaptableNote) SetConversionCap(v float64)`

SetConversionCap sets ConversionCap field to given value.

### HasConversionCap

`func (o *CaptableCaptableNote) HasConversionCap() bool`

HasConversionCap returns a boolean if a field has been set.

### GetDiscountRate

`func (o *CaptableCaptableNote) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *CaptableCaptableNote) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *CaptableCaptableNote) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *CaptableCaptableNote) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetId

`func (o *CaptableCaptableNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCaptableNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCaptableNote) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCaptableNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *CaptableCaptableNote) GetInterestRate() float64`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *CaptableCaptableNote) GetInterestRateOk() (*float64, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *CaptableCaptableNote) SetInterestRate(v float64)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *CaptableCaptableNote) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableCaptableNote) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableCaptableNote) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableCaptableNote) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableCaptableNote) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPublicId

`func (o *CaptableCaptableNote) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CaptableCaptableNote) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CaptableCaptableNote) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CaptableCaptableNote) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableCaptableNote) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableCaptableNote) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableCaptableNote) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableCaptableNote) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableCaptableNote) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableCaptableNote) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableCaptableNote) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableCaptableNote) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableCaptableNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableCaptableNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableCaptableNote) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableCaptableNote) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CaptableCaptableNote) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptableCaptableNote) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptableCaptableNote) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptableCaptableNote) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


