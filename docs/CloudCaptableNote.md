# CloudCaptableNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capital** | Pointer to **float32** | Capital is the principal the investor lent. | [optional] 
**ConversionCap** | Pointer to **float32** | ConversionCap is the valuation cap on conversion, if any. | [optional] 
**DiscountRate** | Pointer to **float32** | DiscountRate is the discount to the next round&#39;s price, if any. | [optional] 
**Id** | Pointer to **string** | ID is the note id. | [optional] 
**InterestRate** | Pointer to **float32** | InterestRate is the annual interest rate, if any. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the note was signed. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the note&#39;s shareable identifier, unique within the company. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the note&#39;s state, e.g. DRAFT or ACTIVE. | [optional] 
**Type** | Pointer to **string** | Type is the instrument kind, e.g. NOTE. | [optional] 

## Methods

### NewCloudCaptableNote

`func NewCloudCaptableNote() *CloudCaptableNote`

NewCloudCaptableNote instantiates a new CloudCaptableNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableNoteWithDefaults

`func NewCloudCaptableNoteWithDefaults() *CloudCaptableNote`

NewCloudCaptableNoteWithDefaults instantiates a new CloudCaptableNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapital

`func (o *CloudCaptableNote) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *CloudCaptableNote) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *CloudCaptableNote) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *CloudCaptableNote) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetConversionCap

`func (o *CloudCaptableNote) GetConversionCap() float32`

GetConversionCap returns the ConversionCap field if non-nil, zero value otherwise.

### GetConversionCapOk

`func (o *CloudCaptableNote) GetConversionCapOk() (*float32, bool)`

GetConversionCapOk returns a tuple with the ConversionCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionCap

`func (o *CloudCaptableNote) SetConversionCap(v float32)`

SetConversionCap sets ConversionCap field to given value.

### HasConversionCap

`func (o *CloudCaptableNote) HasConversionCap() bool`

HasConversionCap returns a boolean if a field has been set.

### GetDiscountRate

`func (o *CloudCaptableNote) GetDiscountRate() float32`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *CloudCaptableNote) GetDiscountRateOk() (*float32, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *CloudCaptableNote) SetDiscountRate(v float32)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *CloudCaptableNote) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableNote) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *CloudCaptableNote) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *CloudCaptableNote) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *CloudCaptableNote) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *CloudCaptableNote) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIssueDate

`func (o *CloudCaptableNote) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CloudCaptableNote) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CloudCaptableNote) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CloudCaptableNote) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPublicId

`func (o *CloudCaptableNote) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CloudCaptableNote) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CloudCaptableNote) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CloudCaptableNote) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableNote) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableNote) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableNote) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableNote) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CloudCaptableNote) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CloudCaptableNote) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CloudCaptableNote) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CloudCaptableNote) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCaptableNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCaptableNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCaptableNote) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCaptableNote) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CloudCaptableNote) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCaptableNote) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCaptableNote) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCaptableNote) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


