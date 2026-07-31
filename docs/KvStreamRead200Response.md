# KvStreamRead200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]KvStreamEntry**](KvStreamEntry.md) |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 

## Methods

### NewKvStreamRead200Response

`func NewKvStreamRead200Response() *KvStreamRead200Response`

NewKvStreamRead200Response instantiates a new KvStreamRead200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStreamRead200ResponseWithDefaults

`func NewKvStreamRead200ResponseWithDefaults() *KvStreamRead200Response`

NewKvStreamRead200ResponseWithDefaults instantiates a new KvStreamRead200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *KvStreamRead200Response) GetEntries() []KvStreamEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *KvStreamRead200Response) GetEntriesOk() (*[]KvStreamEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *KvStreamRead200Response) SetEntries(v []KvStreamEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *KvStreamRead200Response) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetLength

`func (o *KvStreamRead200Response) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *KvStreamRead200Response) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *KvStreamRead200Response) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *KvStreamRead200Response) HasLength() bool`

HasLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


