# CompanyTariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency is the ISO code every amount on this quote is denominated in. | [optional] 
**DueNowCents** | Pointer to **int64** | DueNowCents is what is charged to begin: every non-recurring line. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation the filing fee belongs to. | [optional] 
**Lines** | Pointer to [**[]CompanyCharge**](CompanyCharge.md) | Lines are the charges, in the order a reader should see them. | [optional] 
**Recurring** | Pointer to **string** | Recurring is how often RecurringCents repeats — \&quot;yearly\&quot; for an agent of record. Empty when nothing on this quote recurs. | [optional] 
**RecurringCents** | Pointer to **int64** | RecurringCents is what repeats, and Recurring says how often. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity this prices: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewCompanyTariff

`func NewCompanyTariff() *CompanyTariff`

NewCompanyTariff instantiates a new CompanyTariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTariffWithDefaults

`func NewCompanyTariffWithDefaults() *CompanyTariff`

NewCompanyTariffWithDefaults instantiates a new CompanyTariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CompanyTariff) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyTariff) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyTariff) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyTariff) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueNowCents

`func (o *CompanyTariff) GetDueNowCents() int64`

GetDueNowCents returns the DueNowCents field if non-nil, zero value otherwise.

### GetDueNowCentsOk

`func (o *CompanyTariff) GetDueNowCentsOk() (*int64, bool)`

GetDueNowCentsOk returns a tuple with the DueNowCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueNowCents

`func (o *CompanyTariff) SetDueNowCents(v int64)`

SetDueNowCents sets DueNowCents field to given value.

### HasDueNowCents

`func (o *CompanyTariff) HasDueNowCents() bool`

HasDueNowCents returns a boolean if a field has been set.

### GetJurisdiction

`func (o *CompanyTariff) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *CompanyTariff) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *CompanyTariff) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *CompanyTariff) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetLines

`func (o *CompanyTariff) GetLines() []CompanyCharge`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *CompanyTariff) GetLinesOk() (*[]CompanyCharge, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *CompanyTariff) SetLines(v []CompanyCharge)`

SetLines sets Lines field to given value.

### HasLines

`func (o *CompanyTariff) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetRecurring

`func (o *CompanyTariff) GetRecurring() string`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CompanyTariff) GetRecurringOk() (*string, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CompanyTariff) SetRecurring(v string)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *CompanyTariff) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetRecurringCents

`func (o *CompanyTariff) GetRecurringCents() int64`

GetRecurringCents returns the RecurringCents field if non-nil, zero value otherwise.

### GetRecurringCentsOk

`func (o *CompanyTariff) GetRecurringCentsOk() (*int64, bool)`

GetRecurringCentsOk returns a tuple with the RecurringCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCents

`func (o *CompanyTariff) SetRecurringCents(v int64)`

SetRecurringCents sets RecurringCents field to given value.

### HasRecurringCents

`func (o *CompanyTariff) HasRecurringCents() bool`

HasRecurringCents returns a boolean if a field has been set.

### GetStructure

`func (o *CompanyTariff) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *CompanyTariff) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *CompanyTariff) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *CompanyTariff) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


