# AutoListAppConnections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoAppConnection**](AutoAppConnection.md) |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoListAppConnections200Response

`func NewAutoListAppConnections200Response() *AutoListAppConnections200Response`

NewAutoListAppConnections200Response instantiates a new AutoListAppConnections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoListAppConnections200ResponseWithDefaults

`func NewAutoListAppConnections200ResponseWithDefaults() *AutoListAppConnections200Response`

NewAutoListAppConnections200ResponseWithDefaults instantiates a new AutoListAppConnections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoListAppConnections200Response) GetData() []AutoAppConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoListAppConnections200Response) GetDataOk() (*[]AutoAppConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoListAppConnections200Response) SetData(v []AutoAppConnection)`

SetData sets Data field to given value.

### HasData

`func (o *AutoListAppConnections200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNext

`func (o *AutoListAppConnections200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AutoListAppConnections200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AutoListAppConnections200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AutoListAppConnections200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *AutoListAppConnections200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AutoListAppConnections200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AutoListAppConnections200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *AutoListAppConnections200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


