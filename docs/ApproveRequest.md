# ApproveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the author to approve, from the path. | [optional] 
**ShareBps** | Pointer to **int32** | ShareBps overrides this author&#39;s royalty share, in basis points (0–10000). 0 keeps the platform default. A share change never rewrites history: existing ledger rows keep the share that was applied when they were written. | [optional] 

## Methods

### NewApproveRequest

`func NewApproveRequest() *ApproveRequest`

NewApproveRequest instantiates a new ApproveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveRequestWithDefaults

`func NewApproveRequestWithDefaults() *ApproveRequest`

NewApproveRequestWithDefaults instantiates a new ApproveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApproveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApproveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApproveRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApproveRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShareBps

`func (o *ApproveRequest) GetShareBps() int32`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *ApproveRequest) GetShareBpsOk() (*int32, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *ApproveRequest) SetShareBps(v int32)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *ApproveRequest) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


