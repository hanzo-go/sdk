# BillingPayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**ArrivalDate** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**FailureCode** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**ProviderRef** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingPayout

`func NewBillingPayout() *BillingPayout`

NewBillingPayout instantiates a new BillingPayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPayoutWithDefaults

`func NewBillingPayoutWithDefaults() *BillingPayout`

NewBillingPayoutWithDefaults instantiates a new BillingPayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingPayout) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingPayout) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingPayout) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingPayout) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetArrivalDate

`func (o *BillingPayout) GetArrivalDate() string`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *BillingPayout) GetArrivalDateOk() (*string, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *BillingPayout) SetArrivalDate(v string)`

SetArrivalDate sets ArrivalDate field to given value.

### HasArrivalDate

`func (o *BillingPayout) HasArrivalDate() bool`

HasArrivalDate returns a boolean if a field has been set.

### GetCreated

`func (o *BillingPayout) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillingPayout) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillingPayout) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillingPayout) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingPayout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingPayout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingPayout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingPayout) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *BillingPayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingPayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingPayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingPayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationId

`func (o *BillingPayout) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *BillingPayout) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *BillingPayout) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *BillingPayout) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationType

`func (o *BillingPayout) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *BillingPayout) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *BillingPayout) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *BillingPayout) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetFailureCode

`func (o *BillingPayout) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *BillingPayout) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *BillingPayout) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *BillingPayout) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureMessage

`func (o *BillingPayout) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *BillingPayout) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *BillingPayout) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *BillingPayout) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *BillingPayout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPayout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPayout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingPayout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingPayout) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingPayout) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingPayout) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingPayout) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BillingPayout) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BillingPayout) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProviderRef

`func (o *BillingPayout) GetProviderRef() string`

GetProviderRef returns the ProviderRef field if non-nil, zero value otherwise.

### GetProviderRefOk

`func (o *BillingPayout) GetProviderRefOk() (*string, bool)`

GetProviderRefOk returns a tuple with the ProviderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRef

`func (o *BillingPayout) SetProviderRef(v string)`

SetProviderRef sets ProviderRef field to given value.

### HasProviderRef

`func (o *BillingPayout) HasProviderRef() bool`

HasProviderRef returns a boolean if a field has been set.

### GetStatus

`func (o *BillingPayout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingPayout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingPayout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingPayout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


