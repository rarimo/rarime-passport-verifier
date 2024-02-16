/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GistData struct {
	Key
	Attributes GistDataAttributes `json:"attributes"`
}
type GistDataResponse struct {
	Data     GistData `json:"data"`
	Included Included `json:"included"`
}

type GistDataListResponse struct {
	Data     []GistData `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustGistData - returns GistData from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGistData(key Key) *GistData {
	var gistData GistData
	if c.tryFindEntry(key, &gistData) {
		return &gistData
	}
	return nil
}
