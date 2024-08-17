/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Represents a Square Online site, which is an online store for a Square seller.
type Site struct {
	// The Square-assigned ID of the site.
	Id string `json:"id,omitempty"`
	// The title of the site.
	SiteTitle string `json:"site_title,omitempty"`
	// The domain of the site (without the protocol). For example, `mysite1.square.site`.
	Domain string `json:"domain,omitempty"`
	// Indicates whether the site is published.
	IsPublished bool `json:"is_published,omitempty"`
	// The timestamp of when the site was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp of when the site was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
