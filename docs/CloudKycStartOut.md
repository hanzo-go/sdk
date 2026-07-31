# CloudKycStartOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s incorporation record, with each founder&#39;s session reference and status recorded on it. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired identity-verification provider&#39;s name. | [optional] 
**Sessions** | Pointer to [**[]CloudKycSession**](CloudKycSession.md) | Sessions is one entry per founder, in the order the founders are recorded. | [optional] 

## Methods

### NewCloudKycStartOut

`func NewCloudKycStartOut() *CloudKycStartOut`

NewCloudKycStartOut instantiates a new CloudKycStartOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKycStartOutWithDefaults

`func NewCloudKycStartOutWithDefaults() *CloudKycStartOut`

NewCloudKycStartOutWithDefaults instantiates a new CloudKycStartOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *CloudKycStartOut) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudKycStartOut) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudKycStartOut) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudKycStartOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *CloudKycStartOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudKycStartOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudKycStartOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudKycStartOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSessions

`func (o *CloudKycStartOut) GetSessions() []CloudKycSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *CloudKycStartOut) GetSessionsOk() (*[]CloudKycSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *CloudKycStartOut) SetSessions(v []CloudKycSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *CloudKycStartOut) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


