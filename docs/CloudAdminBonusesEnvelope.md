# CloudAdminBonusesEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudAdminBonusDirectory**](CloudAdminBonusDirectory.md) | Data is the directory itself. | [optional] 
**Msg** | Pointer to **string** | Msg is empty on success; the console surfaces it when status is not \&quot;ok\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewCloudAdminBonusesEnvelope

`func NewCloudAdminBonusesEnvelope() *CloudAdminBonusesEnvelope`

NewCloudAdminBonusesEnvelope instantiates a new CloudAdminBonusesEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminBonusesEnvelopeWithDefaults

`func NewCloudAdminBonusesEnvelopeWithDefaults() *CloudAdminBonusesEnvelope`

NewCloudAdminBonusesEnvelopeWithDefaults instantiates a new CloudAdminBonusesEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAdminBonusesEnvelope) GetData() CloudAdminBonusDirectory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAdminBonusesEnvelope) GetDataOk() (*CloudAdminBonusDirectory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAdminBonusesEnvelope) SetData(v CloudAdminBonusDirectory)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAdminBonusesEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudAdminBonusesEnvelope) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudAdminBonusesEnvelope) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudAdminBonusesEnvelope) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudAdminBonusesEnvelope) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAdminBonusesEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAdminBonusesEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAdminBonusesEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAdminBonusesEnvelope) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


