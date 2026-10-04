# PlatformBinarySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Image is the toolchain image the recipe runs in, a Go bookworm image by default. It is the one field the GitHub lane ignores: there the runner IS the toolchain, and a cluster has to be told what a runner already is. | [optional] 
**Ldflags** | Pointer to **string** | Ldflags are the Go linker flags, &#x60;-s -w&#x60; when the recipe names none, on one line. Go lane only. | [optional] 
**Main** | Pointer to **string** | Main is the Go package to build, repo-relative (&#x60;.&#x60; or &#x60;./cmd/x&#x60;), and it selects the GO LANE. Defaults to &#x60;.&#x60; when neither lane is named; declaring it together with &#x60;run&#x60; is refused. | [optional] 
**Name** | Pointer to **string** | Name is the artifact&#39;s base name: the prefix of every file published for this entry, and the name a host later asks for. It must match &#x60;^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$&#x60;, which is what makes it safe as both a filename and a URL path segment. | [optional] 
**Out** | Pointer to **string** | Out is the glob of files &#x60;run&#x60; produced, relative to the repo root; matching nothing FAILS the build rather than publishing an empty entry. It expands unquoted, so it is bounded to path and glob characters. The Go lane names its own files and ignores this. | [optional] 
**Platforms** | Pointer to **[]string** | Platforms are the &#x60;&lt;os&gt;/&lt;arch&gt;&#x60; pairs the Go lane cross-compiles, [linux/amd64] by default. Each one publishes as &#x60;&lt;name&gt;-&lt;os&gt;-&lt;arch&gt;&#x60;, which is the shape a host resolves a binary BY — so the list is what a caller can ask for later. | [optional] 
**Run** | Pointer to **string** | Run is any other toolchain&#39;s build command, run by &#x60;sh -c&#x60; in this entry&#39;s image, and it selects the OTHER LANE. Arbitrary shell is the point — it is the same trust as a Dockerfile RUN — which is why it executes with no object-store credential and no service-account token. It requires &#x60;out&#x60;. | [optional] 

## Methods

### NewPlatformBinarySpec

`func NewPlatformBinarySpec() *PlatformBinarySpec`

NewPlatformBinarySpec instantiates a new PlatformBinarySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformBinarySpecWithDefaults

`func NewPlatformBinarySpecWithDefaults() *PlatformBinarySpec`

NewPlatformBinarySpecWithDefaults instantiates a new PlatformBinarySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PlatformBinarySpec) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformBinarySpec) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformBinarySpec) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformBinarySpec) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLdflags

`func (o *PlatformBinarySpec) GetLdflags() string`

GetLdflags returns the Ldflags field if non-nil, zero value otherwise.

### GetLdflagsOk

`func (o *PlatformBinarySpec) GetLdflagsOk() (*string, bool)`

GetLdflagsOk returns a tuple with the Ldflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdflags

`func (o *PlatformBinarySpec) SetLdflags(v string)`

SetLdflags sets Ldflags field to given value.

### HasLdflags

`func (o *PlatformBinarySpec) HasLdflags() bool`

HasLdflags returns a boolean if a field has been set.

### GetMain

`func (o *PlatformBinarySpec) GetMain() string`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *PlatformBinarySpec) GetMainOk() (*string, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *PlatformBinarySpec) SetMain(v string)`

SetMain sets Main field to given value.

### HasMain

`func (o *PlatformBinarySpec) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetName

`func (o *PlatformBinarySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformBinarySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformBinarySpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformBinarySpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOut

`func (o *PlatformBinarySpec) GetOut() string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *PlatformBinarySpec) GetOutOk() (*string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *PlatformBinarySpec) SetOut(v string)`

SetOut sets Out field to given value.

### HasOut

`func (o *PlatformBinarySpec) HasOut() bool`

HasOut returns a boolean if a field has been set.

### GetPlatforms

`func (o *PlatformBinarySpec) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PlatformBinarySpec) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PlatformBinarySpec) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PlatformBinarySpec) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRun

`func (o *PlatformBinarySpec) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *PlatformBinarySpec) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *PlatformBinarySpec) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *PlatformBinarySpec) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


