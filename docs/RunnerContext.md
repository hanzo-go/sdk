# RunnerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionsUrl** | Pointer to **string** | ActionsURL is where act fetches an action repository from when a step names one. A runner that receives nothing here composes an unusable URL and every job dies before its first step, so the runner supplies its own fallback rather than trusting this to be set. | [optional] 
**Actor** | Pointer to **string** | Actor is who caused the run. | [optional] 
**ApiUrl** | Pointer to **string** | APIURL is where the job reaches the forge&#39;s API. | [optional] 
**BaseRef** | Pointer to **string** | BaseRef is the target branch of a pull request, empty otherwise. | [optional] 
**Event** | Pointer to **string** | Event is the webhook payload that triggered the run, as opaque JSON. It stays bytes because its shape belongs to the event and not to us: the runner hands it to act, which evaluates github.event.* against it. | [optional] 
**EventName** | Pointer to **string** | EventName is what triggered the run, e.g. \&quot;push\&quot;. | [optional] 
**HeadRef** | Pointer to **string** | HeadRef is the source branch of a pull request, empty otherwise. | [optional] 
**Job** | Pointer to **string** | Job is which job of the workflow document this task is. | [optional] 
**Ref** | Pointer to **string** | Ref is the full ref the run is for, e.g. \&quot;refs/heads/main\&quot;. | [optional] 
**RefName** | Pointer to **string** | RefName is that ref without its refs/heads/ or refs/tags/ prefix. | [optional] 
**RefType** | Pointer to **string** | RefType is \&quot;branch\&quot; or \&quot;tag\&quot;. | [optional] 
**Repository** | Pointer to **string** | Repository is \&quot;&lt;owner&gt;/&lt;name&gt;\&quot;. | [optional] 
**RepositoryOwner** | Pointer to **string** | RepositoryOwner is the owner half of it. | [optional] 
**RetentionDays** | Pointer to **string** | RetentionDays is how long the run&#39;s artifacts are kept, as a number in a string because that is what an expression reads. | [optional] 
**RunAttempt** | Pointer to **string** | RunAttempt is which attempt of that run this is, counting from 1. | [optional] 
**RunId** | Pointer to **string** | RunID identifies the run this task belongs to. | [optional] 
**RunNumber** | Pointer to **string** | RunNumber is the run&#39;s position in its repository, counting from 1. | [optional] 
**RuntimeToken** | Pointer to **string** | RuntimeToken authorizes the job against the actions runtime: artifacts, the cache. The runner passes it as ACTIONS_RUNTIME_TOKEN and masks it out of the log. | [optional] 
**ServerUrl** | Pointer to **string** | ServerURL is the forge&#39;s own origin, as a link in a log points at it. | [optional] 
**Sha** | Pointer to **string** | Sha is the commit being run, in full. | [optional] 
**Token** | Pointer to **string** | Token is the job&#39;s credential against the forge&#39;s own API — github.token in an expression. | [optional] 

## Methods

### NewRunnerContext

`func NewRunnerContext() *RunnerContext`

NewRunnerContext instantiates a new RunnerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerContextWithDefaults

`func NewRunnerContextWithDefaults() *RunnerContext`

NewRunnerContextWithDefaults instantiates a new RunnerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionsUrl

`func (o *RunnerContext) GetActionsUrl() string`

GetActionsUrl returns the ActionsUrl field if non-nil, zero value otherwise.

### GetActionsUrlOk

`func (o *RunnerContext) GetActionsUrlOk() (*string, bool)`

GetActionsUrlOk returns a tuple with the ActionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsUrl

`func (o *RunnerContext) SetActionsUrl(v string)`

SetActionsUrl sets ActionsUrl field to given value.

### HasActionsUrl

`func (o *RunnerContext) HasActionsUrl() bool`

HasActionsUrl returns a boolean if a field has been set.

### GetActor

`func (o *RunnerContext) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RunnerContext) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RunnerContext) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RunnerContext) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetApiUrl

`func (o *RunnerContext) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *RunnerContext) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *RunnerContext) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *RunnerContext) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetBaseRef

`func (o *RunnerContext) GetBaseRef() string`

GetBaseRef returns the BaseRef field if non-nil, zero value otherwise.

### GetBaseRefOk

`func (o *RunnerContext) GetBaseRefOk() (*string, bool)`

GetBaseRefOk returns a tuple with the BaseRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRef

`func (o *RunnerContext) SetBaseRef(v string)`

SetBaseRef sets BaseRef field to given value.

### HasBaseRef

`func (o *RunnerContext) HasBaseRef() bool`

HasBaseRef returns a boolean if a field has been set.

### GetEvent

`func (o *RunnerContext) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RunnerContext) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RunnerContext) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *RunnerContext) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventName

`func (o *RunnerContext) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *RunnerContext) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *RunnerContext) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *RunnerContext) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetHeadRef

`func (o *RunnerContext) GetHeadRef() string`

GetHeadRef returns the HeadRef field if non-nil, zero value otherwise.

### GetHeadRefOk

`func (o *RunnerContext) GetHeadRefOk() (*string, bool)`

GetHeadRefOk returns a tuple with the HeadRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRef

`func (o *RunnerContext) SetHeadRef(v string)`

SetHeadRef sets HeadRef field to given value.

### HasHeadRef

`func (o *RunnerContext) HasHeadRef() bool`

HasHeadRef returns a boolean if a field has been set.

### GetJob

`func (o *RunnerContext) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *RunnerContext) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *RunnerContext) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *RunnerContext) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetRef

`func (o *RunnerContext) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RunnerContext) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RunnerContext) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RunnerContext) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRefName

`func (o *RunnerContext) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *RunnerContext) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *RunnerContext) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *RunnerContext) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRefType

`func (o *RunnerContext) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *RunnerContext) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *RunnerContext) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *RunnerContext) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRepository

`func (o *RunnerContext) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RunnerContext) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RunnerContext) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RunnerContext) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryOwner

`func (o *RunnerContext) GetRepositoryOwner() string`

GetRepositoryOwner returns the RepositoryOwner field if non-nil, zero value otherwise.

### GetRepositoryOwnerOk

`func (o *RunnerContext) GetRepositoryOwnerOk() (*string, bool)`

GetRepositoryOwnerOk returns a tuple with the RepositoryOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryOwner

`func (o *RunnerContext) SetRepositoryOwner(v string)`

SetRepositoryOwner sets RepositoryOwner field to given value.

### HasRepositoryOwner

`func (o *RunnerContext) HasRepositoryOwner() bool`

HasRepositoryOwner returns a boolean if a field has been set.

### GetRetentionDays

`func (o *RunnerContext) GetRetentionDays() string`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *RunnerContext) GetRetentionDaysOk() (*string, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *RunnerContext) SetRetentionDays(v string)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *RunnerContext) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetRunAttempt

`func (o *RunnerContext) GetRunAttempt() string`

GetRunAttempt returns the RunAttempt field if non-nil, zero value otherwise.

### GetRunAttemptOk

`func (o *RunnerContext) GetRunAttemptOk() (*string, bool)`

GetRunAttemptOk returns a tuple with the RunAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAttempt

`func (o *RunnerContext) SetRunAttempt(v string)`

SetRunAttempt sets RunAttempt field to given value.

### HasRunAttempt

`func (o *RunnerContext) HasRunAttempt() bool`

HasRunAttempt returns a boolean if a field has been set.

### GetRunId

`func (o *RunnerContext) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *RunnerContext) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *RunnerContext) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *RunnerContext) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunNumber

`func (o *RunnerContext) GetRunNumber() string`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *RunnerContext) GetRunNumberOk() (*string, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *RunnerContext) SetRunNumber(v string)`

SetRunNumber sets RunNumber field to given value.

### HasRunNumber

`func (o *RunnerContext) HasRunNumber() bool`

HasRunNumber returns a boolean if a field has been set.

### GetRuntimeToken

`func (o *RunnerContext) GetRuntimeToken() string`

GetRuntimeToken returns the RuntimeToken field if non-nil, zero value otherwise.

### GetRuntimeTokenOk

`func (o *RunnerContext) GetRuntimeTokenOk() (*string, bool)`

GetRuntimeTokenOk returns a tuple with the RuntimeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeToken

`func (o *RunnerContext) SetRuntimeToken(v string)`

SetRuntimeToken sets RuntimeToken field to given value.

### HasRuntimeToken

`func (o *RunnerContext) HasRuntimeToken() bool`

HasRuntimeToken returns a boolean if a field has been set.

### GetServerUrl

`func (o *RunnerContext) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *RunnerContext) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *RunnerContext) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *RunnerContext) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetSha

`func (o *RunnerContext) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *RunnerContext) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *RunnerContext) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *RunnerContext) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetToken

`func (o *RunnerContext) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RunnerContext) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RunnerContext) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RunnerContext) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


