# Audit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Block** | Pointer to **string** |  | [optional] 
**Block2** | Pointer to **string** |  | [optional] 
**BlockHash** | Pointer to **string** |  | [optional] 
**BlockHash2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsTriggered** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedCommit** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Provider2** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RequestUri** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to **string** |  | [optional] 
**Transaction2** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewAudit

`func NewAudit() *Audit`

NewAudit instantiates a new Audit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWithDefaults

`func NewAuditWithDefaults() *Audit`

NewAuditWithDefaults instantiates a new Audit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Audit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Audit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Audit) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Audit) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlock

`func (o *Audit) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Audit) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Audit) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *Audit) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetBlock2

`func (o *Audit) GetBlock2() string`

GetBlock2 returns the Block2 field if non-nil, zero value otherwise.

### GetBlock2Ok

`func (o *Audit) GetBlock2Ok() (*string, bool)`

GetBlock2Ok returns a tuple with the Block2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock2

`func (o *Audit) SetBlock2(v string)`

SetBlock2 sets Block2 field to given value.

### HasBlock2

`func (o *Audit) HasBlock2() bool`

HasBlock2 returns a boolean if a field has been set.

### GetBlockHash

`func (o *Audit) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Audit) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Audit) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *Audit) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockHash2

`func (o *Audit) GetBlockHash2() string`

GetBlockHash2 returns the BlockHash2 field if non-nil, zero value otherwise.

### GetBlockHash2Ok

`func (o *Audit) GetBlockHash2Ok() (*string, bool)`

GetBlockHash2Ok returns a tuple with the BlockHash2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash2

`func (o *Audit) SetBlockHash2(v string)`

SetBlockHash2 sets BlockHash2 field to given value.

### HasBlockHash2

`func (o *Audit) HasBlockHash2() bool`

HasBlockHash2 returns a boolean if a field has been set.

### GetCity

`func (o *Audit) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Audit) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Audit) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Audit) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetClientIp

`func (o *Audit) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *Audit) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *Audit) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *Audit) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCount

`func (o *Audit) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Audit) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Audit) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Audit) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Audit) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Audit) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Audit) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Audit) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetErrorText

`func (o *Audit) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *Audit) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *Audit) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *Audit) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetId

`func (o *Audit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Audit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Audit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Audit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsTriggered

`func (o *Audit) GetIsTriggered() bool`

GetIsTriggered returns the IsTriggered field if non-nil, zero value otherwise.

### GetIsTriggeredOk

`func (o *Audit) GetIsTriggeredOk() (*bool, bool)`

GetIsTriggeredOk returns a tuple with the IsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTriggered

`func (o *Audit) SetIsTriggered(v bool)`

SetIsTriggered sets IsTriggered field to given value.

### HasIsTriggered

`func (o *Audit) HasIsTriggered() bool`

HasIsTriggered returns a boolean if a field has been set.

### GetLanguage

`func (o *Audit) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Audit) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Audit) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Audit) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMethod

`func (o *Audit) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Audit) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Audit) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Audit) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *Audit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Audit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Audit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Audit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedCommit

`func (o *Audit) GetNeedCommit() bool`

GetNeedCommit returns the NeedCommit field if non-nil, zero value otherwise.

### GetNeedCommitOk

`func (o *Audit) GetNeedCommitOk() (*bool, bool)`

GetNeedCommitOk returns a tuple with the NeedCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedCommit

`func (o *Audit) SetNeedCommit(v bool)`

SetNeedCommit sets NeedCommit field to given value.

### HasNeedCommit

`func (o *Audit) HasNeedCommit() bool`

HasNeedCommit returns a boolean if a field has been set.

### GetObject

`func (o *Audit) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Audit) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Audit) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Audit) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrganization

`func (o *Audit) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Audit) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Audit) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Audit) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *Audit) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Audit) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Audit) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Audit) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProvider

`func (o *Audit) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Audit) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Audit) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Audit) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProvider2

`func (o *Audit) GetProvider2() string`

GetProvider2 returns the Provider2 field if non-nil, zero value otherwise.

### GetProvider2Ok

`func (o *Audit) GetProvider2Ok() (*string, bool)`

GetProvider2Ok returns a tuple with the Provider2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider2

`func (o *Audit) SetProvider2(v string)`

SetProvider2 sets Provider2 field to given value.

### HasProvider2

`func (o *Audit) HasProvider2() bool`

HasProvider2 returns a boolean if a field has been set.

### GetQuery

`func (o *Audit) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Audit) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Audit) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Audit) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRegion

`func (o *Audit) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Audit) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Audit) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Audit) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestUri

`func (o *Audit) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *Audit) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *Audit) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *Audit) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResponse

`func (o *Audit) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Audit) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Audit) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Audit) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSection

`func (o *Audit) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *Audit) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *Audit) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *Audit) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetTransaction

`func (o *Audit) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Audit) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Audit) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *Audit) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetTransaction2

`func (o *Audit) GetTransaction2() string`

GetTransaction2 returns the Transaction2 field if non-nil, zero value otherwise.

### GetTransaction2Ok

`func (o *Audit) GetTransaction2Ok() (*string, bool)`

GetTransaction2Ok returns a tuple with the Transaction2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction2

`func (o *Audit) SetTransaction2(v string)`

SetTransaction2 sets Transaction2 field to given value.

### HasTransaction2

`func (o *Audit) HasTransaction2() bool`

HasTransaction2 returns a boolean if a field has been set.

### GetUnit

`func (o *Audit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Audit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Audit) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Audit) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUser

`func (o *Audit) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Audit) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Audit) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Audit) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserAgent

`func (o *Audit) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Audit) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Audit) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Audit) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


