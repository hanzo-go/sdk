# TreeJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]TreeEntryJSON**](TreeEntryJSON.md) | Entries are the immediate children, directories before files. | [optional] 

## Methods

### NewTreeJSON

`func NewTreeJSON() *TreeJSON`

NewTreeJSON instantiates a new TreeJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeJSONWithDefaults

`func NewTreeJSONWithDefaults() *TreeJSON`

NewTreeJSONWithDefaults instantiates a new TreeJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *TreeJSON) GetEntries() []TreeEntryJSON`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TreeJSON) GetEntriesOk() (*[]TreeEntryJSON, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TreeJSON) SetEntries(v []TreeEntryJSON)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *TreeJSON) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


