# ExecListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | Pointer to **string** | LastModified is when the BYTES last changed, as RFC 3339 in UTC to the second — &#x60;2026-01-02T03:04:05Z&#x60;, the sandbox&#39;s own &#x60;date -u -r&#x60; on the file. It is an mtime and not a creation time, so a file a later run overwrote carries that run&#39;s clock. Never empty: a row exists only because &#x60;find&#x60; stat-ed the file. | [optional] 
**Name** | Pointer to **string** | Name is the file&#39;s IDENTIFIER, &#x60;{session_id}/{fileId}&#x60; whole — never the bare filename, and never URL-escaped. It is exactly what GET /v1/exec/download takes after its prefix, and hanzo.chat matches it as a PREFIX (&#x60;name.startsWith(session + \&quot;/\&quot;)&#x60;) to decide which rows belong to a session it is holding. &#x60;fileId&#x60; is the path RELATIVE to the session&#39;s artifact directory, so it carries &#x60;/&#x60; for anything the run wrote in a sub-directory. | [optional] 

## Methods

### NewExecListing

`func NewExecListing() *ExecListing`

NewExecListing instantiates a new ExecListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecListingWithDefaults

`func NewExecListingWithDefaults() *ExecListing`

NewExecListingWithDefaults instantiates a new ExecListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *ExecListing) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ExecListing) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ExecListing) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ExecListing) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *ExecListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecListing) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


