# CloudSeedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created is false when this ref had already been seeded — the injection is at-most-once. | [optional] 
**Entry** | Pointer to [**CloudJournalEntry**](CloudJournalEntry.md) | Entry is the journal entry the injection wrote. | [optional] 
**ReserveCents** | Pointer to **int32** | ReserveCents is the fund balance after the injection. | [optional] 

## Methods

### NewCloudSeedData

`func NewCloudSeedData() *CloudSeedData`

NewCloudSeedData instantiates a new CloudSeedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSeedDataWithDefaults

`func NewCloudSeedDataWithDefaults() *CloudSeedData`

NewCloudSeedDataWithDefaults instantiates a new CloudSeedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudSeedData) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudSeedData) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudSeedData) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudSeedData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntry

`func (o *CloudSeedData) GetEntry() CloudJournalEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *CloudSeedData) GetEntryOk() (*CloudJournalEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *CloudSeedData) SetEntry(v CloudJournalEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *CloudSeedData) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetReserveCents

`func (o *CloudSeedData) GetReserveCents() int32`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *CloudSeedData) GetReserveCentsOk() (*int32, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *CloudSeedData) SetReserveCents(v int32)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *CloudSeedData) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


