# ClickCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counted** | Pointer to **bool** | Counted says the in-memory buffer took the ping. It does NOT say the code exists — this is deliberately not a code-existence oracle, and an unknown code simply no-ops at flush time. false means the buffer was full and the ping was dropped, which is harmless: clicks are vanity and move no money. | [optional] 

## Methods

### NewClickCount

`func NewClickCount() *ClickCount`

NewClickCount instantiates a new ClickCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickCountWithDefaults

`func NewClickCountWithDefaults() *ClickCount`

NewClickCountWithDefaults instantiates a new ClickCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounted

`func (o *ClickCount) GetCounted() bool`

GetCounted returns the Counted field if non-nil, zero value otherwise.

### GetCountedOk

`func (o *ClickCount) GetCountedOk() (*bool, bool)`

GetCountedOk returns a tuple with the Counted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounted

`func (o *ClickCount) SetCounted(v bool)`

SetCounted sets Counted field to given value.

### HasCounted

`func (o *ClickCount) HasCounted() bool`

HasCounted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


