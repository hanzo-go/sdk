# PubsubConnInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to **int32** |  | [optional] 
**InMsgs** | Pointer to **int32** |  | [optional] 
**OutMsgs** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubConnInfo

`func NewPubsubConnInfo() *PubsubConnInfo`

NewPubsubConnInfo instantiates a new PubsubConnInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubConnInfoWithDefaults

`func NewPubsubConnInfoWithDefaults() *PubsubConnInfo`

NewPubsubConnInfoWithDefaults instantiates a new PubsubConnInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *PubsubConnInfo) GetCid() int32`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PubsubConnInfo) GetCidOk() (*int32, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PubsubConnInfo) SetCid(v int32)`

SetCid sets Cid field to given value.

### HasCid

`func (o *PubsubConnInfo) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetIp

`func (o *PubsubConnInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PubsubConnInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PubsubConnInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PubsubConnInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *PubsubConnInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PubsubConnInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PubsubConnInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PubsubConnInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetName

`func (o *PubsubConnInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PubsubConnInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PubsubConnInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PubsubConnInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLang

`func (o *PubsubConnInfo) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *PubsubConnInfo) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *PubsubConnInfo) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *PubsubConnInfo) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetVersion

`func (o *PubsubConnInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PubsubConnInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PubsubConnInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PubsubConnInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSubscriptions

`func (o *PubsubConnInfo) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PubsubConnInfo) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PubsubConnInfo) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PubsubConnInfo) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetInMsgs

`func (o *PubsubConnInfo) GetInMsgs() int32`

GetInMsgs returns the InMsgs field if non-nil, zero value otherwise.

### GetInMsgsOk

`func (o *PubsubConnInfo) GetInMsgsOk() (*int32, bool)`

GetInMsgsOk returns a tuple with the InMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsgs

`func (o *PubsubConnInfo) SetInMsgs(v int32)`

SetInMsgs sets InMsgs field to given value.

### HasInMsgs

`func (o *PubsubConnInfo) HasInMsgs() bool`

HasInMsgs returns a boolean if a field has been set.

### GetOutMsgs

`func (o *PubsubConnInfo) GetOutMsgs() int32`

GetOutMsgs returns the OutMsgs field if non-nil, zero value otherwise.

### GetOutMsgsOk

`func (o *PubsubConnInfo) GetOutMsgsOk() (*int32, bool)`

GetOutMsgsOk returns a tuple with the OutMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgs

`func (o *PubsubConnInfo) SetOutMsgs(v int32)`

SetOutMsgs sets OutMsgs field to given value.

### HasOutMsgs

`func (o *PubsubConnInfo) HasOutMsgs() bool`

HasOutMsgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


