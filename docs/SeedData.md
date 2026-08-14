# SeedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created is false when this ref had already been seeded — the injection is at-most-once. | [optional] 
**Entry** | Pointer to [**JournalEntry**](JournalEntry.md) | Entry is the journal entry the injection wrote. | [optional] 
**ReserveCents** | Pointer to **int32** | ReserveCents is the fund balance after the injection. | [optional] 

## Methods

### NewSeedData

`func NewSeedData() *SeedData`

NewSeedData instantiates a new SeedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeedDataWithDefaults

`func NewSeedDataWithDefaults() *SeedData`

NewSeedDataWithDefaults instantiates a new SeedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SeedData) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SeedData) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SeedData) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SeedData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntry

`func (o *SeedData) GetEntry() JournalEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *SeedData) GetEntryOk() (*JournalEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *SeedData) SetEntry(v JournalEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *SeedData) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetReserveCents

`func (o *SeedData) GetReserveCents() int32`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *SeedData) GetReserveCentsOk() (*int32, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *SeedData) SetReserveCents(v int32)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *SeedData) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


