# CloudAskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowUps** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**MaxQueries** | Pointer to **int32** |  | [optional] 
**MaxSources** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to **[]string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 
**System** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAskRequest

`func NewCloudAskRequest() *CloudAskRequest`

NewCloudAskRequest instantiates a new CloudAskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAskRequestWithDefaults

`func NewCloudAskRequestWithDefaults() *CloudAskRequest`

NewCloudAskRequestWithDefaults instantiates a new CloudAskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowUps

`func (o *CloudAskRequest) GetFollowUps() bool`

GetFollowUps returns the FollowUps field if non-nil, zero value otherwise.

### GetFollowUpsOk

`func (o *CloudAskRequest) GetFollowUpsOk() (*bool, bool)`

GetFollowUpsOk returns a tuple with the FollowUps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUps

`func (o *CloudAskRequest) SetFollowUps(v bool)`

SetFollowUps sets FollowUps field to given value.

### HasFollowUps

`func (o *CloudAskRequest) HasFollowUps() bool`

HasFollowUps returns a boolean if a field has been set.

### GetLanguage

`func (o *CloudAskRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CloudAskRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CloudAskRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CloudAskRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMaxQueries

`func (o *CloudAskRequest) GetMaxQueries() int32`

GetMaxQueries returns the MaxQueries field if non-nil, zero value otherwise.

### GetMaxQueriesOk

`func (o *CloudAskRequest) GetMaxQueriesOk() (*int32, bool)`

GetMaxQueriesOk returns a tuple with the MaxQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueries

`func (o *CloudAskRequest) SetMaxQueries(v int32)`

SetMaxQueries sets MaxQueries field to given value.

### HasMaxQueries

`func (o *CloudAskRequest) HasMaxQueries() bool`

HasMaxQueries returns a boolean if a field has been set.

### GetMaxSources

`func (o *CloudAskRequest) GetMaxSources() int32`

GetMaxSources returns the MaxSources field if non-nil, zero value otherwise.

### GetMaxSourcesOk

`func (o *CloudAskRequest) GetMaxSourcesOk() (*int32, bool)`

GetMaxSourcesOk returns a tuple with the MaxSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSources

`func (o *CloudAskRequest) SetMaxSources(v int32)`

SetMaxSources sets MaxSources field to given value.

### HasMaxSources

`func (o *CloudAskRequest) HasMaxSources() bool`

HasMaxSources returns a boolean if a field has been set.

### GetMode

`func (o *CloudAskRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CloudAskRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CloudAskRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CloudAskRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModel

`func (o *CloudAskRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudAskRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudAskRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudAskRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetQ

`func (o *CloudAskRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *CloudAskRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *CloudAskRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *CloudAskRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetQuestion

`func (o *CloudAskRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *CloudAskRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *CloudAskRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *CloudAskRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetSources

`func (o *CloudAskRequest) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudAskRequest) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudAskRequest) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudAskRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStream

`func (o *CloudAskRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudAskRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudAskRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudAskRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSystem

`func (o *CloudAskRequest) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *CloudAskRequest) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *CloudAskRequest) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *CloudAskRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


