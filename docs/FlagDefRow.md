# FlagDefRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to **interface{}** |  | [optional] 
**Key** | Pointer to **string** | Key is the flag&#39;s primary key in the caller&#39;s (org, project) store, and the name evaluation looks it up by. On a write it is taken from the URL, never from the body: the stored document&#39;s own \&quot;key\&quot; is forced to match. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the definition was last written, RFC 3339 UTC. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the email of the principal who last wrote it. Empty when the write came from an in-process composer (an experiment registering its own assignment flag) rather than from a signed-in person. | [optional] 
**Version** | Pointer to **int64** | Version is 1 when the key was created and rises by one on every overwrite. It counts writes, not content changes: re-storing an identical document bumps it. | [optional] 

## Methods

### NewFlagDefRow

`func NewFlagDefRow() *FlagDefRow`

NewFlagDefRow instantiates a new FlagDefRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagDefRowWithDefaults

`func NewFlagDefRowWithDefaults() *FlagDefRow`

NewFlagDefRowWithDefaults instantiates a new FlagDefRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *FlagDefRow) GetDefinition() interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *FlagDefRow) GetDefinitionOk() (*interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *FlagDefRow) SetDefinition(v interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *FlagDefRow) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *FlagDefRow) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *FlagDefRow) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetKey

`func (o *FlagDefRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagDefRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagDefRow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagDefRow) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FlagDefRow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlagDefRow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlagDefRow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FlagDefRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FlagDefRow) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FlagDefRow) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FlagDefRow) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FlagDefRow) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *FlagDefRow) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagDefRow) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagDefRow) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FlagDefRow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


