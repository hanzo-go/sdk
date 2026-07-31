# CloudBackfilled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | Pointer to **string** | EntryID is the finance ledger entry created, or \&quot;\&quot; when the balance was non-positive and there was nothing to carry. | [optional] 
**MigratedCents** | Pointer to **int32** | MigratedCents is the balance carried across, read from commerce BEFORE the move. | [optional] 
**Org** | Pointer to **string** | Org is the tenant migrated. | [optional] 

## Methods

### NewCloudBackfilled

`func NewCloudBackfilled() *CloudBackfilled`

NewCloudBackfilled instantiates a new CloudBackfilled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBackfilledWithDefaults

`func NewCloudBackfilledWithDefaults() *CloudBackfilled`

NewCloudBackfilledWithDefaults instantiates a new CloudBackfilled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *CloudBackfilled) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *CloudBackfilled) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *CloudBackfilled) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *CloudBackfilled) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetMigratedCents

`func (o *CloudBackfilled) GetMigratedCents() int32`

GetMigratedCents returns the MigratedCents field if non-nil, zero value otherwise.

### GetMigratedCentsOk

`func (o *CloudBackfilled) GetMigratedCentsOk() (*int32, bool)`

GetMigratedCentsOk returns a tuple with the MigratedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedCents

`func (o *CloudBackfilled) SetMigratedCents(v int32)`

SetMigratedCents sets MigratedCents field to given value.

### HasMigratedCents

`func (o *CloudBackfilled) HasMigratedCents() bool`

HasMigratedCents returns a boolean if a field has been set.

### GetOrg

`func (o *CloudBackfilled) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudBackfilled) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudBackfilled) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudBackfilled) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


