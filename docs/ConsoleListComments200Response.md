# ConsoleListComments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleComment**](ConsoleComment.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListComments200Response

`func NewConsoleListComments200Response() *ConsoleListComments200Response`

NewConsoleListComments200Response instantiates a new ConsoleListComments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListComments200ResponseWithDefaults

`func NewConsoleListComments200ResponseWithDefaults() *ConsoleListComments200Response`

NewConsoleListComments200ResponseWithDefaults instantiates a new ConsoleListComments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListComments200Response) GetData() []ConsoleComment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListComments200Response) GetDataOk() (*[]ConsoleComment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListComments200Response) SetData(v []ConsoleComment)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListComments200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListComments200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListComments200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListComments200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListComments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


