# CloudNamespaceCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title is the namespace&#39;s display title. Cloudflare mints the id. | [optional] 

## Methods

### NewCloudNamespaceCreateIn

`func NewCloudNamespaceCreateIn() *CloudNamespaceCreateIn`

NewCloudNamespaceCreateIn instantiates a new CloudNamespaceCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNamespaceCreateInWithDefaults

`func NewCloudNamespaceCreateInWithDefaults() *CloudNamespaceCreateIn`

NewCloudNamespaceCreateInWithDefaults instantiates a new CloudNamespaceCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CloudNamespaceCreateIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudNamespaceCreateIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudNamespaceCreateIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudNamespaceCreateIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


