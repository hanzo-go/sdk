# CloudAnchorOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudAnchorData**](CloudAnchorData.md) | Data is the anchoring status. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. A submit that failed still answers ok with the anchor&#39;s own status set to \&quot;error\&quot; — the attempt is the product. | [optional] 

## Methods

### NewCloudAnchorOut

`func NewCloudAnchorOut() *CloudAnchorOut`

NewCloudAnchorOut instantiates a new CloudAnchorOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnchorOutWithDefaults

`func NewCloudAnchorOutWithDefaults() *CloudAnchorOut`

NewCloudAnchorOutWithDefaults instantiates a new CloudAnchorOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAnchorOut) GetData() CloudAnchorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAnchorOut) GetDataOk() (*CloudAnchorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAnchorOut) SetData(v CloudAnchorData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAnchorOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudAnchorOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudAnchorOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudAnchorOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudAnchorOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAnchorOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAnchorOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAnchorOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAnchorOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


