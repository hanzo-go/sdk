# Tariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency is the ISO code every amount on this quote is denominated in. | [optional] 
**DueNowCents** | Pointer to **int32** | DueNowCents is what is charged to begin: every non-recurring line. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation the filing fee belongs to. | [optional] 
**Lines** | Pointer to [**[]Charge**](Charge.md) | Lines are the charges, in the order a reader should see them. | [optional] 
**Recurring** | Pointer to **string** | Recurring is how often RecurringCents repeats — \&quot;yearly\&quot; for an agent of record. Empty when nothing on this quote recurs. | [optional] 
**RecurringCents** | Pointer to **int32** | RecurringCents is what repeats, and Recurring says how often. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity this prices: c-corp, llc or dao-llc. | [optional] 

## Methods

### NewTariff

`func NewTariff() *Tariff`

NewTariff instantiates a new Tariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffWithDefaults

`func NewTariffWithDefaults() *Tariff`

NewTariffWithDefaults instantiates a new Tariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Tariff) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Tariff) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Tariff) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Tariff) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueNowCents

`func (o *Tariff) GetDueNowCents() int32`

GetDueNowCents returns the DueNowCents field if non-nil, zero value otherwise.

### GetDueNowCentsOk

`func (o *Tariff) GetDueNowCentsOk() (*int32, bool)`

GetDueNowCentsOk returns a tuple with the DueNowCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueNowCents

`func (o *Tariff) SetDueNowCents(v int32)`

SetDueNowCents sets DueNowCents field to given value.

### HasDueNowCents

`func (o *Tariff) HasDueNowCents() bool`

HasDueNowCents returns a boolean if a field has been set.

### GetJurisdiction

`func (o *Tariff) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *Tariff) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *Tariff) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *Tariff) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetLines

`func (o *Tariff) GetLines() []Charge`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Tariff) GetLinesOk() (*[]Charge, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Tariff) SetLines(v []Charge)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Tariff) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetRecurring

`func (o *Tariff) GetRecurring() string`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *Tariff) GetRecurringOk() (*string, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *Tariff) SetRecurring(v string)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *Tariff) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetRecurringCents

`func (o *Tariff) GetRecurringCents() int32`

GetRecurringCents returns the RecurringCents field if non-nil, zero value otherwise.

### GetRecurringCentsOk

`func (o *Tariff) GetRecurringCentsOk() (*int32, bool)`

GetRecurringCentsOk returns a tuple with the RecurringCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCents

`func (o *Tariff) SetRecurringCents(v int32)`

SetRecurringCents sets RecurringCents field to given value.

### HasRecurringCents

`func (o *Tariff) HasRecurringCents() bool`

HasRecurringCents returns a boolean if a field has been set.

### GetStructure

`func (o *Tariff) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *Tariff) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *Tariff) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *Tariff) HasStructure() bool`

HasStructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


