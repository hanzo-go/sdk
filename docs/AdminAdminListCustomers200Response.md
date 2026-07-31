# AdminAdminListCustomers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AdminCustomerRow**](AdminCustomerRow.md) |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminAdminListCustomers200Response

`func NewAdminAdminListCustomers200Response() *AdminAdminListCustomers200Response`

NewAdminAdminListCustomers200Response instantiates a new AdminAdminListCustomers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAdminListCustomers200ResponseWithDefaults

`func NewAdminAdminListCustomers200ResponseWithDefaults() *AdminAdminListCustomers200Response`

NewAdminAdminListCustomers200ResponseWithDefaults instantiates a new AdminAdminListCustomers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminAdminListCustomers200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminAdminListCustomers200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminAdminListCustomers200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminAdminListCustomers200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AdminAdminListCustomers200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminAdminListCustomers200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminAdminListCustomers200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminAdminListCustomers200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AdminAdminListCustomers200Response) GetData() []AdminCustomerRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminAdminListCustomers200Response) GetDataOk() (*[]AdminCustomerRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminAdminListCustomers200Response) SetData(v []AdminCustomerRow)`

SetData sets Data field to given value.

### HasData

`func (o *AdminAdminListCustomers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AdminAdminListCustomers200Response) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AdminAdminListCustomers200Response) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AdminAdminListCustomers200Response) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AdminAdminListCustomers200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


