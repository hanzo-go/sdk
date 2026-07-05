# KvScanKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | Next cursor (0 when scan complete) | [optional] 
**Keys** | Pointer to [**[]KvKeyValue**](KvKeyValue.md) |  | [optional] 

## Methods

### NewKvScanKeys200Response

`func NewKvScanKeys200Response() *KvScanKeys200Response`

NewKvScanKeys200Response instantiates a new KvScanKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvScanKeys200ResponseWithDefaults

`func NewKvScanKeys200ResponseWithDefaults() *KvScanKeys200Response`

NewKvScanKeys200ResponseWithDefaults instantiates a new KvScanKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *KvScanKeys200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *KvScanKeys200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *KvScanKeys200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *KvScanKeys200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetKeys

`func (o *KvScanKeys200Response) GetKeys() []KvKeyValue`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *KvScanKeys200Response) GetKeysOk() (*[]KvKeyValue, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *KvScanKeys200Response) SetKeys(v []KvKeyValue)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *KvScanKeys200Response) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


