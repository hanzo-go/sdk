# PatchSessionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwd** | Pointer to **string** | Cwd is where the session is working NOW.  It was write-once — captured at register and never again — which is right for a run that starts in a directory and stays there, and wrong for a linked shell, which is a place a person moves around in. The console showed the directory &#x60;hanzo link&#x60; happened to be run from and kept showing it after the shell had walked away, so the field answered \&quot;which work is this\&quot; with an answer that was true once. A pointer, so an unchanged path is an omitted field rather than a repeated write. | [optional] 
**Id** | Pointer to **string** | ID is the session to update, from the path. | [optional] 
**Project** | Pointer to **string** | Project tags the product this session built; Published is the author&#39;s decision to let anyone read the story (provenance.go). Both are pointers so \&quot;absent\&quot; and \&quot;cleared\&quot; are different requests. | [optional] 
**Published** | Pointer to **bool** | Published opens the session&#39;s story to the public build route; false withdraws it, and withdrawing is always allowed. PUBLISHING is refused unless the session names a Project — the one set in this same request, or the one already stored — because that route is keyed on (org, project). It widens READ access to what is already there and grants nothing else. | [optional] 
**Status** | Pointer to **string** | Status moves the session to running, paused, done or error. A session that has already finished refuses any change with 409 — done and error are monotonic — and moving INTO one stamps the end time. This is the surface REPORTING what happened; a control command never writes it. | [optional] 
**Target** | Pointer to **string** | Target re-dispatches a session to a run-target (the #48 association). \&quot;\&quot; detaches. | [optional] 
**Terminal** | Pointer to **string** | Terminal publishes (or, with \&quot;\&quot;, withdraws) the URL this session&#39;s live terminal can be watched at. A pointer so \&quot;absent\&quot; and \&quot;withdrawn\&quot; are different requests: a session that stops sharing must be able to say so. | [optional] 
**Title** | Pointer to **string** | Title rewrites the human line, up to 512 characters — usually because the work turned out to be something other than what it was opened as. | [optional] 

## Methods

### NewPatchSessionIn

`func NewPatchSessionIn() *PatchSessionIn`

NewPatchSessionIn instantiates a new PatchSessionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSessionInWithDefaults

`func NewPatchSessionInWithDefaults() *PatchSessionIn`

NewPatchSessionInWithDefaults instantiates a new PatchSessionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwd

`func (o *PatchSessionIn) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *PatchSessionIn) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *PatchSessionIn) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *PatchSessionIn) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetId

`func (o *PatchSessionIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchSessionIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchSessionIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchSessionIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *PatchSessionIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchSessionIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchSessionIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchSessionIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublished

`func (o *PatchSessionIn) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *PatchSessionIn) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *PatchSessionIn) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *PatchSessionIn) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetStatus

`func (o *PatchSessionIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchSessionIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchSessionIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchSessionIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *PatchSessionIn) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchSessionIn) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchSessionIn) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchSessionIn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTerminal

`func (o *PatchSessionIn) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *PatchSessionIn) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *PatchSessionIn) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *PatchSessionIn) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *PatchSessionIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchSessionIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchSessionIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchSessionIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


