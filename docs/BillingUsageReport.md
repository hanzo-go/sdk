# BillingUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**BillingMoney**](BillingMoney.md) | Amount is what the act costs, as an exact decimal in USD. It must be greater than zero. | [optional] 
**Id** | Pointer to **string** | ID names the act, chosen by the reporting application and stable across its retries: the ledger debits one act once. At most 128 printable characters, no whitespace. Reusing an ID for a different amount is a 409. | [optional] 
**Model** | Pointer to **string** | Model names the unit the amount prices (a machine size, a model id), and is recorded with the debit. | [optional] 
**Org** | Pointer to **string** | Org is the organization this usage is billed to. It must be the org the token acts in — the application&#39;s own, or one that granted it membership and is selected with X-Org-Id — and any other is refused before money moves. | [optional] 
**Project** | Pointer to **string** | Project attributes the debit to one project of the org. Empty is the org&#39;s default project. | [optional] 
**Service** | Pointer to **string** | Service scopes the debit for spend caps and attributes it to a product, as a lowercase slug. Empty takes the reporting application&#39;s name. | [optional] 

## Methods

### NewBillingUsageReport

`func NewBillingUsageReport() *BillingUsageReport`

NewBillingUsageReport instantiates a new BillingUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageReportWithDefaults

`func NewBillingUsageReportWithDefaults() *BillingUsageReport`

NewBillingUsageReportWithDefaults instantiates a new BillingUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingUsageReport) GetAmount() BillingMoney`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingUsageReport) GetAmountOk() (*BillingMoney, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingUsageReport) SetAmount(v BillingMoney)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingUsageReport) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *BillingUsageReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingUsageReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingUsageReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingUsageReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *BillingUsageReport) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BillingUsageReport) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BillingUsageReport) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BillingUsageReport) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *BillingUsageReport) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BillingUsageReport) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BillingUsageReport) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BillingUsageReport) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *BillingUsageReport) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BillingUsageReport) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BillingUsageReport) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BillingUsageReport) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetService

`func (o *BillingUsageReport) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingUsageReport) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingUsageReport) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingUsageReport) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


