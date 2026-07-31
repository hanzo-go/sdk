# CloudCaptableSafe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capital** | Pointer to **float32** | Capital is the cash the investor put in. | [optional] 
**DiscountRate** | Pointer to **float32** | DiscountRate is the discount to the next round&#39;s price, if any. | [optional] 
**Id** | Pointer to **string** | ID is the SAFE id. | [optional] 
**IssueDate** | Pointer to **string** | IssueDate is the ISO date the SAFE was signed. | [optional] 
**Mfn** | Pointer to **bool** | MFN is true when the SAFE carries a most-favoured-nation clause. | [optional] 
**ProRata** | Pointer to **bool** | ProRata is true when the SAFE carries pro-rata rights. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the SAFE&#39;s shareable identifier, unique within the company. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the investor. | [optional] 
**StakeholderName** | Pointer to **string** | StakeholderName is that investor&#39;s name. | [optional] 
**Status** | Pointer to **string** | Status is the SAFE&#39;s state, e.g. DRAFT or ACTIVE. | [optional] 
**Type** | Pointer to **string** | Type is POST_MONEY or PRE_MONEY. | [optional] 
**ValuationCap** | Pointer to **float32** | ValuationCap is the valuation cap, if any. | [optional] 

## Methods

### NewCloudCaptableSafe

`func NewCloudCaptableSafe() *CloudCaptableSafe`

NewCloudCaptableSafe instantiates a new CloudCaptableSafe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableSafeWithDefaults

`func NewCloudCaptableSafeWithDefaults() *CloudCaptableSafe`

NewCloudCaptableSafeWithDefaults instantiates a new CloudCaptableSafe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapital

`func (o *CloudCaptableSafe) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *CloudCaptableSafe) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *CloudCaptableSafe) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *CloudCaptableSafe) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetDiscountRate

`func (o *CloudCaptableSafe) GetDiscountRate() float32`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *CloudCaptableSafe) GetDiscountRateOk() (*float32, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *CloudCaptableSafe) SetDiscountRate(v float32)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *CloudCaptableSafe) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableSafe) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableSafe) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableSafe) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableSafe) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *CloudCaptableSafe) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CloudCaptableSafe) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CloudCaptableSafe) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CloudCaptableSafe) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMfn

`func (o *CloudCaptableSafe) GetMfn() bool`

GetMfn returns the Mfn field if non-nil, zero value otherwise.

### GetMfnOk

`func (o *CloudCaptableSafe) GetMfnOk() (*bool, bool)`

GetMfnOk returns a tuple with the Mfn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfn

`func (o *CloudCaptableSafe) SetMfn(v bool)`

SetMfn sets Mfn field to given value.

### HasMfn

`func (o *CloudCaptableSafe) HasMfn() bool`

HasMfn returns a boolean if a field has been set.

### GetProRata

`func (o *CloudCaptableSafe) GetProRata() bool`

GetProRata returns the ProRata field if non-nil, zero value otherwise.

### GetProRataOk

`func (o *CloudCaptableSafe) GetProRataOk() (*bool, bool)`

GetProRataOk returns a tuple with the ProRata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRata

`func (o *CloudCaptableSafe) SetProRata(v bool)`

SetProRata sets ProRata field to given value.

### HasProRata

`func (o *CloudCaptableSafe) HasProRata() bool`

HasProRata returns a boolean if a field has been set.

### GetPublicId

`func (o *CloudCaptableSafe) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CloudCaptableSafe) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CloudCaptableSafe) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CloudCaptableSafe) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableSafe) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableSafe) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableSafe) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableSafe) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.

### GetStakeholderName

`func (o *CloudCaptableSafe) GetStakeholderName() string`

GetStakeholderName returns the StakeholderName field if non-nil, zero value otherwise.

### GetStakeholderNameOk

`func (o *CloudCaptableSafe) GetStakeholderNameOk() (*string, bool)`

GetStakeholderNameOk returns a tuple with the StakeholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderName

`func (o *CloudCaptableSafe) SetStakeholderName(v string)`

SetStakeholderName sets StakeholderName field to given value.

### HasStakeholderName

`func (o *CloudCaptableSafe) HasStakeholderName() bool`

HasStakeholderName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCaptableSafe) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCaptableSafe) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCaptableSafe) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCaptableSafe) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CloudCaptableSafe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCaptableSafe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCaptableSafe) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudCaptableSafe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValuationCap

`func (o *CloudCaptableSafe) GetValuationCap() float32`

GetValuationCap returns the ValuationCap field if non-nil, zero value otherwise.

### GetValuationCapOk

`func (o *CloudCaptableSafe) GetValuationCapOk() (*float32, bool)`

GetValuationCapOk returns a tuple with the ValuationCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuationCap

`func (o *CloudCaptableSafe) SetValuationCap(v float32)`

SetValuationCap sets ValuationCap field to given value.

### HasValuationCap

`func (o *CloudCaptableSafe) HasValuationCap() bool`

HasValuationCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


