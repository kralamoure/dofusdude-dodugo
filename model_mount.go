/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude - [PHP](https://github.com/dofusdude/dofusdude-php)  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply donate to that cause, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).

API version: 0.8.1
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the Mount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mount{}

// Mount struct for Mount
type Mount struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Effects []EffectsEntry `json:"effects,omitempty"`
}

// NewMount instantiates a new Mount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMount() *Mount {
	this := Mount{}
	return &this
}

// NewMountWithDefaults instantiates a new Mount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountWithDefaults() *Mount {
	this := Mount{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *Mount) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mount) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *Mount) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *Mount) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Mount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Mount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Mount) SetName(v string) {
	o.Name = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *Mount) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mount) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *Mount) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *Mount) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *Mount) GetImageUrls() ImageUrls {
	if o == nil || IsNil(o.ImageUrls) {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mount) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *Mount) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *Mount) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mount) GetEffects() []EffectsEntry {
	if o == nil {
		var ret []EffectsEntry
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mount) GetEffectsOk() ([]EffectsEntry, bool) {
	if o == nil || IsNil(o.Effects) {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *Mount) HasEffects() bool {
	if o != nil && IsNil(o.Effects) {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []EffectsEntry and assigns it to the Effects field.
func (o *Mount) SetEffects(v []EffectsEntry) {
	o.Effects = v
}

func (o Mount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FamilyName) {
		toSerialize["family_name"] = o.FamilyName
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	return toSerialize, nil
}

type NullableMount struct {
	value *Mount
	isSet bool
}

func (v NullableMount) Get() *Mount {
	return v.value
}

func (v *NullableMount) Set(val *Mount) {
	v.value = val
	v.isSet = true
}

func (v NullableMount) IsSet() bool {
	return v.isSet
}

func (v *NullableMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMount(val *Mount) *NullableMount {
	return &NullableMount{value: val, isSet: true}
}

func (v NullableMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


