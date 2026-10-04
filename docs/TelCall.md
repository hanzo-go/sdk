# TelCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent names the Hanzo assistant handling the call. Set means the call was answered by that assistant rather than connected to a person. | [optional] 
**From** | Pointer to **string** | From is the calling number in E.164. It must be one this org holds: a carrier refuses an origination from a number nobody proved they own. | [optional] 
**Id** | Pointer to **string** | ID is the carrier&#39;s handle for the call — what a hangup or a lookup names. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the call was placed for or received by. | [optional] 
**Status** | Pointer to **string** | Status is where the call is: \&quot;queued\&quot;, \&quot;ringing\&quot;, \&quot;answered\&quot;, \&quot;completed\&quot; or \&quot;failed\&quot;. Only the last two are terminal. | [optional] 
**To** | Pointer to **string** | To is the called number in E.164. | [optional] 

## Methods

### NewTelCall

`func NewTelCall() *TelCall`

NewTelCall instantiates a new TelCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelCallWithDefaults

`func NewTelCallWithDefaults() *TelCall`

NewTelCallWithDefaults instantiates a new TelCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *TelCall) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *TelCall) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *TelCall) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *TelCall) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFrom

`func (o *TelCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TelCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TelCall) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TelCall) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *TelCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TelCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *TelCall) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TelCall) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TelCall) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TelCall) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStatus

`func (o *TelCall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TelCall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TelCall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TelCall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *TelCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TelCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TelCall) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TelCall) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


