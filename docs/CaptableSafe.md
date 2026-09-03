# CaptableSafe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capital** | Pointer to **float64** | Capital is the cash the investor put in. | [optional] 
**DiscountRate** | Pointer to **float64** | DiscountRate is the discount to the next round&#39;s price, if any. | [optional] 
**Id** | Pointer to **string** | ID is the SAFE id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the SAFE was signed. | [optional] 
**Mfn** | Pointer to **bool** | MFN is true when the SAFE carries a most-favoured-nation clause. | [optional] 
**ProRata** | Pointer to **bool** | ProRata is true when the SAFE carries pro-rata rights. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the SAFE&#39;s shareable identifier, unique within the company. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the SAFE&#39;s state, e.g. DRAFT or ACTIVE. | [optional] 
**Type** | Pointer to **string** | Type is POST_MONEY or PRE_MONEY. | [optional] 
**ValuationCap** | Pointer to **float64** | ValuationCap is the valuation cap, if any. | [optional] 

## Methods

### NewCaptableSafe

`func NewCaptableSafe() *CaptableSafe`

NewCaptableSafe instantiates a new CaptableSafe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableSafeWithDefaults

`func NewCaptableSafeWithDefaults() *CaptableSafe`

NewCaptableSafeWithDefaults instantiates a new CaptableSafe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapital

`func (o *CaptableSafe) GetCapital() float64`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *CaptableSafe) GetCapitalOk() (*float64, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *CaptableSafe) SetCapital(v float64)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *CaptableSafe) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetDiscountRate

`func (o *CaptableSafe) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *CaptableSafe) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *CaptableSafe) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *CaptableSafe) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetId

`func (o *CaptableSafe) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableSafe) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableSafe) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableSafe) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CaptableSafe) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaptableSafe) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaptableSafe) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaptableSafe) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMfn

`func (o *CaptableSafe) GetMfn() bool`

GetMfn returns the Mfn field if non-nil, zero value otherwise.

### GetMfnOk

`func (o *CaptableSafe) GetMfnOk() (*bool, bool)`

GetMfnOk returns a tuple with the Mfn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfn

`func (o *CaptableSafe) SetMfn(v bool)`

SetMfn sets Mfn field to given value.

### HasMfn

`func (o *CaptableSafe) HasMfn() bool`

HasMfn returns a boolean if a field has been set.

### GetProRata

`func (o *CaptableSafe) GetProRata() bool`

GetProRata returns the ProRata field if non-nil, zero value otherwise.

### GetProRataOk

`func (o *CaptableSafe) GetProRataOk() (*bool, bool)`

GetProRataOk returns a tuple with the ProRata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRata

`func (o *CaptableSafe) SetProRata(v bool)`

SetProRata sets ProRata field to given value.

### HasProRata

`func (o *CaptableSafe) HasProRata() bool`

HasProRata returns a boolean if a field has been set.

### GetPublicId

`func (o *CaptableSafe) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CaptableSafe) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CaptableSafe) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CaptableSafe) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableSafe) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableSafe) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableSafe) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableSafe) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CaptableSafe) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CaptableSafe) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CaptableSafe) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CaptableSafe) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CaptableSafe) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptableSafe) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptableSafe) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptableSafe) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CaptableSafe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptableSafe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptableSafe) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CaptableSafe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValuationCap

`func (o *CaptableSafe) GetValuationCap() float64`

GetValuationCap returns the ValuationCap field if non-nil, zero value otherwise.

### GetValuationCapOk

`func (o *CaptableSafe) GetValuationCapOk() (*float64, bool)`

GetValuationCapOk returns a tuple with the ValuationCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuationCap

`func (o *CaptableSafe) SetValuationCap(v float64)`

SetValuationCap sets ValuationCap field to given value.

### HasValuationCap

`func (o *CaptableSafe) HasValuationCap() bool`

HasValuationCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


