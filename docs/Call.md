# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | queued | ringing | answered | completed | failed | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewCall

`func NewCall() *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *Call) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Call) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Call) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Call) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFrom

`func (o *Call) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Call) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Call) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Call) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *Call) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Call) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Call) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Call) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *Call) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Call) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Call) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Call) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStatus

`func (o *Call) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Call) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Call) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Call) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *Call) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Call) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Call) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Call) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


