# MqListAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]MqAccount**](MqAccount.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListAccounts200Response

`func NewMqListAccounts200Response() *MqListAccounts200Response`

NewMqListAccounts200Response instantiates a new MqListAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListAccounts200ResponseWithDefaults

`func NewMqListAccounts200ResponseWithDefaults() *MqListAccounts200Response`

NewMqListAccounts200ResponseWithDefaults instantiates a new MqListAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *MqListAccounts200Response) GetAccounts() []MqAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *MqListAccounts200Response) GetAccountsOk() (*[]MqAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *MqListAccounts200Response) SetAccounts(v []MqAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *MqListAccounts200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTotal

`func (o *MqListAccounts200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListAccounts200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListAccounts200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListAccounts200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


