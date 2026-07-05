# ConsoleListAnnotationQueues200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleAnnotationQueue**](ConsoleAnnotationQueue.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListAnnotationQueues200Response

`func NewConsoleListAnnotationQueues200Response() *ConsoleListAnnotationQueues200Response`

NewConsoleListAnnotationQueues200Response instantiates a new ConsoleListAnnotationQueues200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListAnnotationQueues200ResponseWithDefaults

`func NewConsoleListAnnotationQueues200ResponseWithDefaults() *ConsoleListAnnotationQueues200Response`

NewConsoleListAnnotationQueues200ResponseWithDefaults instantiates a new ConsoleListAnnotationQueues200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListAnnotationQueues200Response) GetData() []ConsoleAnnotationQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListAnnotationQueues200Response) GetDataOk() (*[]ConsoleAnnotationQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListAnnotationQueues200Response) SetData(v []ConsoleAnnotationQueue)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListAnnotationQueues200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListAnnotationQueues200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListAnnotationQueues200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListAnnotationQueues200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListAnnotationQueues200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


