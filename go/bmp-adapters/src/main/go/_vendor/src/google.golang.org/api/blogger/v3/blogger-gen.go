// Package blogger provides access to the Blogger API.
//
// See https://developers.google.com/blogger/docs/3.0/getting_started
//
// Usage example:
//
//   import "google.golang.org/api/blogger/v3"
//   ...
//   bloggerService, err := blogger.New(oauthHttpClient)
package blogger // import "google.golang.org/api/blogger/v3"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/internal"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = internal.MarshalJSON

const apiId = "blogger:v3"
const apiName = "blogger"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/blogger/v3/"

// OAuth2 scopes used by this API.
const (
	// Manage your Blogger account
	BloggerScope = "https://www.googleapis.com/auth/blogger"

	// View your Blogger account
	BloggerReadonlyScope = "https://www.googleapis.com/auth/blogger.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.BlogUserInfos = NewBlogUserInfosService(s)
	s.Blogs = NewBlogsService(s)
	s.Comments = NewCommentsService(s)
	s.PageViews = NewPageViewsService(s)
	s.Pages = NewPagesService(s)
	s.PostUserInfos = NewPostUserInfosService(s)
	s.Posts = NewPostsService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BlogUserInfos *BlogUserInfosService

	Blogs *BlogsService

	Comments *CommentsService

	PageViews *PageViewsService

	Pages *PagesService

	PostUserInfos *PostUserInfosService

	Posts *PostsService

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBlogUserInfosService(s *Service) *BlogUserInfosService {
	rs := &BlogUserInfosService{s: s}
	return rs
}

type BlogUserInfosService struct {
	s *Service
}

func NewBlogsService(s *Service) *BlogsService {
	rs := &BlogsService{s: s}
	return rs
}

type BlogsService struct {
	s *Service
}

func NewCommentsService(s *Service) *CommentsService {
	rs := &CommentsService{s: s}
	return rs
}

type CommentsService struct {
	s *Service
}

func NewPageViewsService(s *Service) *PageViewsService {
	rs := &PageViewsService{s: s}
	return rs
}

type PageViewsService struct {
	s *Service
}

func NewPagesService(s *Service) *PagesService {
	rs := &PagesService{s: s}
	return rs
}

type PagesService struct {
	s *Service
}

func NewPostUserInfosService(s *Service) *PostUserInfosService {
	rs := &PostUserInfosService{s: s}
	return rs
}

type PostUserInfosService struct {
	s *Service
}

func NewPostsService(s *Service) *PostsService {
	rs := &PostsService{s: s}
	return rs
}

type PostsService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type Blog struct {
	// CustomMetaData: The JSON custom meta-data for the Blog
	CustomMetaData string `json:"customMetaData,omitempty"`

	// Description: The description of this blog. This is displayed
	// underneath the title.
	Description string `json:"description,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entry. Always blogger#blog
	Kind string `json:"kind,omitempty"`

	// Locale: The locale this Blog is set to.
	Locale *BlogLocale `json:"locale,omitempty"`

	// Name: The name of this blog. This is displayed as the title.
	Name string `json:"name,omitempty"`

	// Pages: The container of pages in this blog.
	Pages *BlogPages `json:"pages,omitempty"`

	// Posts: The container of posts in this blog.
	Posts *BlogPosts `json:"posts,omitempty"`

	// Published: RFC 3339 date-time when this blog was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The status of the blog.
	Status string `json:"status,omitempty"`

	// Updated: RFC 3339 date-time when this blog was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL where this blog is published.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomMetaData") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Blog) MarshalJSON() ([]byte, error) {
	type noMethod Blog
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// BlogLocale: The locale this Blog is set to.
type BlogLocale struct {
	// Country: The country this blog's locale is set to.
	Country string `json:"country,omitempty"`

	// Language: The language this blog is authored in.
	Language string `json:"language,omitempty"`

	// Variant: The language variant this blog is authored in.
	Variant string `json:"variant,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Country") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogLocale) MarshalJSON() ([]byte, error) {
	type noMethod BlogLocale
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// BlogPages: The container of pages in this blog.
type BlogPages struct {
	// SelfLink: The URL of the container for pages in this blog.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of pages in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SelfLink") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogPages) MarshalJSON() ([]byte, error) {
	type noMethod BlogPages
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// BlogPosts: The container of posts in this blog.
type BlogPosts struct {
	// Items: The List of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`

	// SelfLink: The URL of the container for posts in this blog.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of posts in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogPosts) MarshalJSON() ([]byte, error) {
	type noMethod BlogPosts
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type BlogList struct {
	// BlogUserInfos: Admin level list of blog per-user information
	BlogUserInfos []*BlogUserInfo `json:"blogUserInfos,omitempty"`

	// Items: The list of Blogs this user has Authorship or Admin rights
	// over.
	Items []*Blog `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#blogList
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BlogUserInfos") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogList) MarshalJSON() ([]byte, error) {
	type noMethod BlogList
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type BlogPerUserInfo struct {
	// BlogId: ID of the Blog resource
	BlogId string `json:"blogId,omitempty"`

	// HasAdminAccess: True if the user has Admin level access to the blog.
	HasAdminAccess bool `json:"hasAdminAccess,omitempty"`

	// Kind: The kind of this entity. Always blogger#blogPerUserInfo
	Kind string `json:"kind,omitempty"`

	// PhotosAlbumKey: The Photo Album Key for the user when adding photos
	// to the blog
	PhotosAlbumKey string `json:"photosAlbumKey,omitempty"`

	// Role: Access permissions that the user has for the blog (ADMIN,
	// AUTHOR, or READER).
	Role string `json:"role,omitempty"`

	// UserId: ID of the User
	UserId string `json:"userId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BlogId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogPerUserInfo) MarshalJSON() ([]byte, error) {
	type noMethod BlogPerUserInfo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type BlogUserInfo struct {
	// Blog: The Blog resource.
	Blog *Blog `json:"blog,omitempty"`

	// BlogUserInfo: Information about a User for the Blog.
	BlogUserInfo *BlogPerUserInfo `json:"blog_user_info,omitempty"`

	// Kind: The kind of this entity. Always blogger#blogUserInfo
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Blog") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlogUserInfo) MarshalJSON() ([]byte, error) {
	type noMethod BlogUserInfo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type Comment struct {
	// Author: The author of this Comment.
	Author *CommentAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this comment.
	Blog *CommentBlog `json:"blog,omitempty"`

	// Content: The actual content of the comment. May include HTML markup.
	Content string `json:"content,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// InReplyTo: Data about the comment this is in reply to.
	InReplyTo *CommentInReplyTo `json:"inReplyTo,omitempty"`

	// Kind: The kind of this entry. Always blogger#comment
	Kind string `json:"kind,omitempty"`

	// Post: Data about the post containing this comment.
	Post *CommentPost `json:"post,omitempty"`

	// Published: RFC 3339 date-time when this comment was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The status of the comment (only populated for admin users)
	Status string `json:"status,omitempty"`

	// Updated: RFC 3339 date-time when this comment was last updated.
	Updated string `json:"updated,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Author") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Comment) MarshalJSON() ([]byte, error) {
	type noMethod Comment
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// CommentAuthor: The author of this Comment.
type CommentAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Comment creator.
	Id string `json:"id,omitempty"`

	// Image: The comment creator's avatar.
	Image *CommentAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Comment creator's Profile page.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentAuthor) MarshalJSON() ([]byte, error) {
	type noMethod CommentAuthor
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// CommentAuthorImage: The comment creator's avatar.
type CommentAuthorImage struct {
	// Url: The comment creator's avatar URL.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentAuthorImage) MarshalJSON() ([]byte, error) {
	type noMethod CommentAuthorImage
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// CommentBlog: Data about the blog containing this comment.
type CommentBlog struct {
	// Id: The identifier of the blog containing this comment.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentBlog) MarshalJSON() ([]byte, error) {
	type noMethod CommentBlog
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// CommentInReplyTo: Data about the comment this is in reply to.
type CommentInReplyTo struct {
	// Id: The identified of the parent of this comment.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentInReplyTo) MarshalJSON() ([]byte, error) {
	type noMethod CommentInReplyTo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// CommentPost: Data about the post containing this comment.
type CommentPost struct {
	// Id: The identifier of the post containing this comment.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentPost) MarshalJSON() ([]byte, error) {
	type noMethod CommentPost
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type CommentList struct {
	// Etag: Etag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: The List of Comments for a Post.
	Items []*Comment `json:"items,omitempty"`

	// Kind: The kind of this entry. Always blogger#commentList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PrevPageToken: Pagination token to fetch the previous page, if one
	// exists.
	PrevPageToken string `json:"prevPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommentList) MarshalJSON() ([]byte, error) {
	type noMethod CommentList
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type Page struct {
	// Author: The author of this Page.
	Author *PageAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this Page.
	Blog *PageBlog `json:"blog,omitempty"`

	// Content: The body content of this Page, in HTML.
	Content string `json:"content,omitempty"`

	// Etag: Etag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entity. Always blogger#page
	Kind string `json:"kind,omitempty"`

	// Published: RFC 3339 date-time when this Page was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The status of the page for admin resources (either LIVE or
	// DRAFT).
	Status string `json:"status,omitempty"`

	// Title: The title of this entity. This is the name displayed in the
	// Admin user interface.
	Title string `json:"title,omitempty"`

	// Updated: RFC 3339 date-time when this Page was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL that this Page is displayed at.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Author") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Page) MarshalJSON() ([]byte, error) {
	type noMethod Page
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PageAuthor: The author of this Page.
type PageAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Page creator.
	Id string `json:"id,omitempty"`

	// Image: The page author's avatar.
	Image *PageAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Page creator's Profile page.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PageAuthor) MarshalJSON() ([]byte, error) {
	type noMethod PageAuthor
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PageAuthorImage: The page author's avatar.
type PageAuthorImage struct {
	// Url: The page author's avatar URL.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PageAuthorImage) MarshalJSON() ([]byte, error) {
	type noMethod PageAuthorImage
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PageBlog: Data about the blog containing this Page.
type PageBlog struct {
	// Id: The identifier of the blog containing this page.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PageBlog) MarshalJSON() ([]byte, error) {
	type noMethod PageBlog
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PageList struct {
	// Etag: Etag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: The list of Pages for a Blog.
	Items []*Page `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#pageList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PageList) MarshalJSON() ([]byte, error) {
	type noMethod PageList
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type Pageviews struct {
	// BlogId: Blog Id
	BlogId string `json:"blogId,omitempty"`

	// Counts: The container of posts in this blog.
	Counts []*PageviewsCounts `json:"counts,omitempty"`

	// Kind: The kind of this entry. Always blogger#page_views
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BlogId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Pageviews) MarshalJSON() ([]byte, error) {
	type noMethod Pageviews
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PageviewsCounts struct {
	// Count: Count of page views for the given time range
	Count int64 `json:"count,omitempty,string"`

	// TimeRange: Time range the given count applies to
	TimeRange string `json:"timeRange,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PageviewsCounts) MarshalJSON() ([]byte, error) {
	type noMethod PageviewsCounts
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type Post struct {
	// Author: The author of this Post.
	Author *PostAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this Post.
	Blog *PostBlog `json:"blog,omitempty"`

	// Content: The content of the Post. May contain HTML markup.
	Content string `json:"content,omitempty"`

	// CustomMetaData: The JSON meta-data for the Post.
	CustomMetaData string `json:"customMetaData,omitempty"`

	// Etag: Etag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: The identifier of this Post.
	Id string `json:"id,omitempty"`

	// Images: Display image for the Post.
	Images []*PostImages `json:"images,omitempty"`

	// Kind: The kind of this entity. Always blogger#post
	Kind string `json:"kind,omitempty"`

	// Labels: The list of labels this Post was tagged with.
	Labels []string `json:"labels,omitempty"`

	// Location: The location for geotagged posts.
	Location *PostLocation `json:"location,omitempty"`

	// Published: RFC 3339 date-time when this Post was published.
	Published string `json:"published,omitempty"`

	// ReaderComments: Comment control and display setting for readers of
	// this post.
	ReaderComments string `json:"readerComments,omitempty"`

	// Replies: The container of comments on this Post.
	Replies *PostReplies `json:"replies,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: Status of the post. Only set for admin-level requests
	Status string `json:"status,omitempty"`

	// Title: The title of the Post.
	Title string `json:"title,omitempty"`

	// TitleLink: The title link URL, similar to atom's related link.
	TitleLink string `json:"titleLink,omitempty"`

	// Updated: RFC 3339 date-time when this Post was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL where this Post is displayed.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Author") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Post) MarshalJSON() ([]byte, error) {
	type noMethod Post
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PostAuthor: The author of this Post.
type PostAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Post creator.
	Id string `json:"id,omitempty"`

	// Image: The Post author's avatar.
	Image *PostAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Post creator's Profile page.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostAuthor) MarshalJSON() ([]byte, error) {
	type noMethod PostAuthor
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PostAuthorImage: The Post author's avatar.
type PostAuthorImage struct {
	// Url: The Post author's avatar URL.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostAuthorImage) MarshalJSON() ([]byte, error) {
	type noMethod PostAuthorImage
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PostBlog: Data about the blog containing this Post.
type PostBlog struct {
	// Id: The identifier of the Blog that contains this Post.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostBlog) MarshalJSON() ([]byte, error) {
	type noMethod PostBlog
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PostImages struct {
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostImages) MarshalJSON() ([]byte, error) {
	type noMethod PostImages
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PostLocation: The location for geotagged posts.
type PostLocation struct {
	// Lat: Location's latitude.
	Lat float64 `json:"lat,omitempty"`

	// Lng: Location's longitude.
	Lng float64 `json:"lng,omitempty"`

	// Name: Location name.
	Name string `json:"name,omitempty"`

	// Span: Location's viewport span. Can be used when rendering a map
	// preview.
	Span string `json:"span,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Lat") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostLocation) MarshalJSON() ([]byte, error) {
	type noMethod PostLocation
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PostReplies: The container of comments on this Post.
type PostReplies struct {
	// Items: The List of Comments for this Post.
	Items []*Comment `json:"items,omitempty"`

	// SelfLink: The URL of the comments on this post.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of comments on this post.
	TotalItems int64 `json:"totalItems,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostReplies) MarshalJSON() ([]byte, error) {
	type noMethod PostReplies
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PostList struct {
	// Etag: Etag of the response.
	Etag string `json:"etag,omitempty"`

	// Items: The list of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#postList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostList) MarshalJSON() ([]byte, error) {
	type noMethod PostList
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PostPerUserInfo struct {
	// BlogId: ID of the Blog that the post resource belongs to.
	BlogId string `json:"blogId,omitempty"`

	// HasEditAccess: True if the user has Author level access to the post.
	HasEditAccess bool `json:"hasEditAccess,omitempty"`

	// Kind: The kind of this entity. Always blogger#postPerUserInfo
	Kind string `json:"kind,omitempty"`

	// PostId: ID of the Post resource.
	PostId string `json:"postId,omitempty"`

	// UserId: ID of the User.
	UserId string `json:"userId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BlogId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostPerUserInfo) MarshalJSON() ([]byte, error) {
	type noMethod PostPerUserInfo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PostUserInfo struct {
	// Kind: The kind of this entity. Always blogger#postUserInfo
	Kind string `json:"kind,omitempty"`

	// Post: The Post resource.
	Post *Post `json:"post,omitempty"`

	// PostUserInfo: Information about a User for the Post.
	PostUserInfo *PostPerUserInfo `json:"post_user_info,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostUserInfo) MarshalJSON() ([]byte, error) {
	type noMethod PostUserInfo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type PostUserInfosList struct {
	// Items: The list of Posts with User information for the post, for this
	// Blog.
	Items []*PostUserInfo `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#postList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PostUserInfosList) MarshalJSON() ([]byte, error) {
	type noMethod PostUserInfosList
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type User struct {
	// About: Profile summary information.
	About string `json:"about,omitempty"`

	// Blogs: The container of blogs for this user.
	Blogs *UserBlogs `json:"blogs,omitempty"`

	// Created: The timestamp of when this profile was created, in seconds
	// since epoch.
	Created string `json:"created,omitempty"`

	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier for this User.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entity. Always blogger#user
	Kind string `json:"kind,omitempty"`

	// Locale: This user's locale
	Locale *UserLocale `json:"locale,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Url: The user's profile page.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "About") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *User) MarshalJSON() ([]byte, error) {
	type noMethod User
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UserBlogs: The container of blogs for this user.
type UserBlogs struct {
	// SelfLink: The URL of the Blogs for this user.
	SelfLink string `json:"selfLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SelfLink") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UserBlogs) MarshalJSON() ([]byte, error) {
	type noMethod UserBlogs
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UserLocale: This user's locale
type UserLocale struct {
	// Country: The user's country setting.
	Country string `json:"country,omitempty"`

	// Language: The user's language setting.
	Language string `json:"language,omitempty"`

	// Variant: The user's language variant setting.
	Variant string `json:"variant,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Country") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UserLocale) MarshalJSON() ([]byte, error) {
	type noMethod UserLocale
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// method id "blogger.blogUserInfos.get":

type BlogUserInfosGetCall struct {
	s      *Service
	userId string
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Gets one blog and user info pair by blogId and userId.
func (r *BlogUserInfosService) Get(userId string, blogId string) *BlogUserInfosGetCall {
	c := &BlogUserInfosGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.blogId = blogId
	return c
}

// MaxPosts sets the optional parameter "maxPosts": Maximum number of
// posts to pull back with the blog.
func (c *BlogUserInfosGetCall) MaxPosts(maxPosts int64) *BlogUserInfosGetCall {
	c.opt_["maxPosts"] = maxPosts
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BlogUserInfosGetCall) Fields(s ...googleapi.Field) *BlogUserInfosGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *BlogUserInfosGetCall) Context(ctx context.Context) *BlogUserInfosGetCall {
	c.ctx_ = ctx
	return c
}

func (c *BlogUserInfosGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxPosts"]; ok {
		params.Set("maxPosts", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{userId}/blogs/{blogId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *BlogUserInfosGetCall) Do() (*BlogUserInfo, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BlogUserInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one blog and user info pair by blogId and userId.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogUserInfos.get",
	//   "parameterOrder": [
	//     "userId",
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxPosts": {
	//       "description": "Maximum number of posts to pull back with the blog.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose blogs are to be fetched. Either the word 'self' (sans quote marks) or the user's profile identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/blogs/{blogId}",
	//   "response": {
	//     "$ref": "BlogUserInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.blogs.get":

type BlogsGetCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Gets one blog by ID.
func (r *BlogsService) Get(blogId string) *BlogsGetCall {
	c := &BlogsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// MaxPosts sets the optional parameter "maxPosts": Maximum number of
// posts to pull back with the blog.
func (c *BlogsGetCall) MaxPosts(maxPosts int64) *BlogsGetCall {
	c.opt_["maxPosts"] = maxPosts
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the blog. Note that some fields require elevated access.
//
// Possible values:
//   "ADMIN" - Admin level detail.
//   "AUTHOR" - Author level detail.
//   "READER" - Reader level detail.
func (c *BlogsGetCall) View(view string) *BlogsGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BlogsGetCall) Fields(s ...googleapi.Field) *BlogsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *BlogsGetCall) Context(ctx context.Context) *BlogsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *BlogsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxPosts"]; ok {
		params.Set("maxPosts", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *BlogsGetCall) Do() (*Blog, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Blog
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one blog by ID.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.get",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxPosts": {
	//       "description": "Maximum number of posts to pull back with the blog.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the blog. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail.",
	//         "Author level detail.",
	//         "Reader level detail."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}",
	//   "response": {
	//     "$ref": "Blog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.blogs.getByUrl":

type BlogsGetByUrlCall struct {
	s    *Service
	url  string
	opt_ map[string]interface{}
	ctx_ context.Context
}

// GetByUrl: Retrieve a Blog by URL.
func (r *BlogsService) GetByUrl(url string) *BlogsGetByUrlCall {
	c := &BlogsGetByUrlCall{s: r.s, opt_: make(map[string]interface{})}
	c.url = url
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the blog. Note that some fields require elevated access.
//
// Possible values:
//   "ADMIN" - Admin level detail.
//   "AUTHOR" - Author level detail.
//   "READER" - Reader level detail.
func (c *BlogsGetByUrlCall) View(view string) *BlogsGetByUrlCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BlogsGetByUrlCall) Fields(s ...googleapi.Field) *BlogsGetByUrlCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *BlogsGetByUrlCall) Context(ctx context.Context) *BlogsGetByUrlCall {
	c.ctx_ = ctx
	return c
}

func (c *BlogsGetByUrlCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("url", fmt.Sprintf("%v", c.url))
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/byurl")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *BlogsGetByUrlCall) Do() (*Blog, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Blog
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Blog by URL.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.getByUrl",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "url": {
	//       "description": "The URL of the blog to retrieve.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the blog. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail.",
	//         "Author level detail.",
	//         "Reader level detail."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/byurl",
	//   "response": {
	//     "$ref": "Blog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.blogs.listByUser":

type BlogsListByUserCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// ListByUser: Retrieves a list of blogs, possibly filtered.
func (r *BlogsService) ListByUser(userId string) *BlogsListByUserCall {
	c := &BlogsListByUserCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// FetchUserInfo sets the optional parameter "fetchUserInfo": Whether
// the response is a list of blogs with per-user information instead of
// just blogs.
func (c *BlogsListByUserCall) FetchUserInfo(fetchUserInfo bool) *BlogsListByUserCall {
	c.opt_["fetchUserInfo"] = fetchUserInfo
	return c
}

// Role sets the optional parameter "role": User access types for blogs
// to include in the results, e.g. AUTHOR will return blogs where the
// user has author level access. If no roles are specified, defaults to
// ADMIN and AUTHOR roles.
//
// Possible values:
//   "ADMIN" - Admin role - Blogs where the user has Admin level access.
//   "AUTHOR" - Author role - Blogs where the user has Author level
// access.
//   "READER" - Reader role - Blogs where the user has Reader level
// access (to a private blog).
func (c *BlogsListByUserCall) Role(role string) *BlogsListByUserCall {
	c.opt_["role"] = role
	return c
}

// Status sets the optional parameter "status": Blog statuses to include
// in the result (default: Live blogs only). Note that ADMIN access is
// required to view deleted blogs.
//
// Possible values:
//   "DELETED" - Blog has been deleted by an administrator.
//   "LIVE" (default) - Blog is currently live.
func (c *BlogsListByUserCall) Status(status string) *BlogsListByUserCall {
	c.opt_["status"] = status
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the blogs. Note that some fields require elevated access.
//
// Possible values:
//   "ADMIN" - Admin level detail.
//   "AUTHOR" - Author level detail.
//   "READER" - Reader level detail.
func (c *BlogsListByUserCall) View(view string) *BlogsListByUserCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BlogsListByUserCall) Fields(s ...googleapi.Field) *BlogsListByUserCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *BlogsListByUserCall) Context(ctx context.Context) *BlogsListByUserCall {
	c.ctx_ = ctx
	return c
}

func (c *BlogsListByUserCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchUserInfo"]; ok {
		params.Set("fetchUserInfo", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["role"]; ok {
		params.Set("role", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{userId}/blogs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *BlogsListByUserCall) Do() (*BlogList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BlogList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of blogs, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.listByUser",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "fetchUserInfo": {
	//       "description": "Whether the response is a list of blogs with per-user information instead of just blogs.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "role": {
	//       "description": "User access types for blogs to include in the results, e.g. AUTHOR will return blogs where the user has author level access. If no roles are specified, defaults to ADMIN and AUTHOR roles.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin role - Blogs where the user has Admin level access.",
	//         "Author role - Blogs where the user has Author level access.",
	//         "Reader role - Blogs where the user has Reader level access (to a private blog)."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "status": {
	//       "default": "LIVE",
	//       "description": "Blog statuses to include in the result (default: Live blogs only). Note that ADMIN access is required to view deleted blogs.",
	//       "enum": [
	//         "DELETED",
	//         "LIVE"
	//       ],
	//       "enumDescriptions": [
	//         "Blog has been deleted by an administrator.",
	//         "Blog is currently live."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user whose blogs are to be fetched. Either the word 'self' (sans quote marks) or the user's profile identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the blogs. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail.",
	//         "Author level detail.",
	//         "Reader level detail."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/blogs",
	//   "response": {
	//     "$ref": "BlogList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.approve":

type CommentsApproveCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// Approve: Marks a comment as not spam.
func (r *CommentsService) Approve(blogId string, postId string, commentId string) *CommentsApproveCall {
	c := &CommentsApproveCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsApproveCall) Fields(s ...googleapi.Field) *CommentsApproveCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsApproveCall) Context(ctx context.Context) *CommentsApproveCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsApproveCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments/{commentId}/approve")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId":    c.blogId,
		"postId":    c.postId,
		"commentId": c.commentId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsApproveCall) Do() (*Comment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Comment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks a comment as not spam.",
	//   "httpMethod": "POST",
	//   "id": "blogger.comments.approve",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to mark as not spam.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}/approve",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.comments.delete":

type CommentsDeleteCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// Delete: Delete a comment by ID.
func (r *CommentsService) Delete(blogId string, postId string, commentId string) *CommentsDeleteCall {
	c := &CommentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsDeleteCall) Fields(s ...googleapi.Field) *CommentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsDeleteCall) Context(ctx context.Context) *CommentsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments/{commentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId":    c.blogId,
		"postId":    c.postId,
		"commentId": c.commentId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete a comment by ID.",
	//   "httpMethod": "DELETE",
	//   "id": "blogger.comments.delete",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.comments.get":

type CommentsGetCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// Get: Gets one comment by ID.
func (r *CommentsService) Get(blogId string, postId string, commentId string) *CommentsGetCall {
	c := &CommentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

// View sets the optional parameter "view": Access level for the
// requested comment (default: READER). Note that some comments will
// require elevated permissions, for example comments where the parent
// posts which is in a draft state, or comments that are pending
// moderation.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Admin level detail
func (c *CommentsGetCall) View(view string) *CommentsGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsGetCall) Fields(s ...googleapi.Field) *CommentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsGetCall) Context(ctx context.Context) *CommentsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments/{commentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId":    c.blogId,
		"postId":    c.postId,
		"commentId": c.commentId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsGetCall) Do() (*Comment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Comment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one comment by ID.",
	//   "httpMethod": "GET",
	//   "id": "blogger.comments.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to containing the comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the post to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level for the requested comment (default: READER). Note that some comments will require elevated permissions, for example comments where the parent posts which is in a draft state, or comments that are pending moderation.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Admin level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.list":

type CommentsListCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// List: Retrieves the comments for a post, possibly filtered.
func (r *CommentsService) List(blogId string, postId string) *CommentsListCall {
	c := &CommentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// EndDate sets the optional parameter "endDate": Latest date of comment
// to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListCall) EndDate(endDate string) *CommentsListCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of the comments is included.
func (c *CommentsListCall) FetchBodies(fetchBodies bool) *CommentsListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of comments to include in the result.
func (c *CommentsListCall) MaxResults(maxResults int64) *CommentsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if request is paged.
func (c *CommentsListCall) PageToken(pageToken string) *CommentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest date of
// comment to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListCall) StartDate(startDate string) *CommentsListCall {
	c.opt_["startDate"] = startDate
	return c
}

// Status sets the optional parameter "status":
//
// Possible values:
//   "emptied" - Comments that have had their content removed
//   "live" - Comments that are publicly visible
//   "pending" - Comments that are awaiting administrator approval
//   "spam" - Comments marked as spam by the administrator
func (c *CommentsListCall) Status(status string) *CommentsListCall {
	c.opt_["status"] = status
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require elevated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *CommentsListCall) View(view string) *CommentsListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsListCall) Fields(s ...googleapi.Field) *CommentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsListCall) Context(ctx context.Context) *CommentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsListCall) Do() (*CommentList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CommentList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the comments for a post, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.comments.list",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch comments from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether the body content of the comments is included.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of comments to include in the result.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the post to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "emptied",
	//         "live",
	//         "pending",
	//         "spam"
	//       ],
	//       "enumDescriptions": [
	//         "Comments that have had their content removed",
	//         "Comments that are publicly visible",
	//         "Comments that are awaiting administrator approval",
	//         "Comments marked as spam by the administrator"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments",
	//   "response": {
	//     "$ref": "CommentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.listByBlog":

type CommentsListByBlogCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// ListByBlog: Retrieves the comments for a blog, across all posts,
// possibly filtered.
func (r *CommentsService) ListByBlog(blogId string) *CommentsListByBlogCall {
	c := &CommentsListByBlogCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// EndDate sets the optional parameter "endDate": Latest date of comment
// to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListByBlogCall) EndDate(endDate string) *CommentsListByBlogCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of the comments is included.
func (c *CommentsListByBlogCall) FetchBodies(fetchBodies bool) *CommentsListByBlogCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of comments to include in the result.
func (c *CommentsListByBlogCall) MaxResults(maxResults int64) *CommentsListByBlogCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if request is paged.
func (c *CommentsListByBlogCall) PageToken(pageToken string) *CommentsListByBlogCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest date of
// comment to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListByBlogCall) StartDate(startDate string) *CommentsListByBlogCall {
	c.opt_["startDate"] = startDate
	return c
}

// Status sets the optional parameter "status":
//
// Possible values:
//   "emptied" - Comments that have had their content removed
//   "live" - Comments that are publicly visible
//   "pending" - Comments that are awaiting administrator approval
//   "spam" - Comments marked as spam by the administrator
func (c *CommentsListByBlogCall) Status(status string) *CommentsListByBlogCall {
	c.opt_["status"] = status
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsListByBlogCall) Fields(s ...googleapi.Field) *CommentsListByBlogCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsListByBlogCall) Context(ctx context.Context) *CommentsListByBlogCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsListByBlogCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/comments")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsListByBlogCall) Do() (*CommentList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CommentList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the comments for a blog, across all posts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.comments.listByBlog",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch comments from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether the body content of the comments is included.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of comments to include in the result.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "emptied",
	//         "live",
	//         "pending",
	//         "spam"
	//       ],
	//       "enumDescriptions": [
	//         "Comments that have had their content removed",
	//         "Comments that are publicly visible",
	//         "Comments that are awaiting administrator approval",
	//         "Comments marked as spam by the administrator"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/comments",
	//   "response": {
	//     "$ref": "CommentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.markAsSpam":

type CommentsMarkAsSpamCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// MarkAsSpam: Marks a comment as spam.
func (r *CommentsService) MarkAsSpam(blogId string, postId string, commentId string) *CommentsMarkAsSpamCall {
	c := &CommentsMarkAsSpamCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsMarkAsSpamCall) Fields(s ...googleapi.Field) *CommentsMarkAsSpamCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsMarkAsSpamCall) Context(ctx context.Context) *CommentsMarkAsSpamCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsMarkAsSpamCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments/{commentId}/spam")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId":    c.blogId,
		"postId":    c.postId,
		"commentId": c.commentId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsMarkAsSpamCall) Do() (*Comment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Comment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks a comment as spam.",
	//   "httpMethod": "POST",
	//   "id": "blogger.comments.markAsSpam",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to mark as spam.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}/spam",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.comments.removeContent":

type CommentsRemoveContentCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// RemoveContent: Removes the content of a comment.
func (r *CommentsService) RemoveContent(blogId string, postId string, commentId string) *CommentsRemoveContentCall {
	c := &CommentsRemoveContentCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CommentsRemoveContentCall) Fields(s ...googleapi.Field) *CommentsRemoveContentCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *CommentsRemoveContentCall) Context(ctx context.Context) *CommentsRemoveContentCall {
	c.ctx_ = ctx
	return c
}

func (c *CommentsRemoveContentCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/comments/{commentId}/removecontent")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId":    c.blogId,
		"postId":    c.postId,
		"commentId": c.commentId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *CommentsRemoveContentCall) Do() (*Comment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Comment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes the content of a comment.",
	//   "httpMethod": "POST",
	//   "id": "blogger.comments.removeContent",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to delete content from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}/removecontent",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pageViews.get":

type PageViewsGetCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Retrieve pageview stats for a Blog.
func (r *PageViewsService) Get(blogId string) *PageViewsGetCall {
	c := &PageViewsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// Range sets the optional parameter "range":
//
// Possible values:
//   "30DAYS" - Page view counts from the last thirty days.
//   "7DAYS" - Page view counts from the last seven days.
//   "all" - Total page view counts from all time.
func (c *PageViewsGetCall) Range(range_ string) *PageViewsGetCall {
	c.opt_["range"] = range_
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PageViewsGetCall) Fields(s ...googleapi.Field) *PageViewsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PageViewsGetCall) Context(ctx context.Context) *PageViewsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *PageViewsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["range"]; ok {
		params.Set("range", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pageviews")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PageViewsGetCall) Do() (*Pageviews, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Pageviews
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve pageview stats for a Blog.",
	//   "httpMethod": "GET",
	//   "id": "blogger.pageViews.get",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "range": {
	//       "enum": [
	//         "30DAYS",
	//         "7DAYS",
	//         "all"
	//       ],
	//       "enumDescriptions": [
	//         "Page view counts from the last thirty days.",
	//         "Page view counts from the last seven days.",
	//         "Total page view counts from all time."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pageviews",
	//   "response": {
	//     "$ref": "Pageviews"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.delete":

type PagesDeleteCall struct {
	s      *Service
	blogId string
	pageId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Delete: Delete a page by ID.
func (r *PagesService) Delete(blogId string, pageId string) *PagesDeleteCall {
	c := &PagesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesDeleteCall) Fields(s ...googleapi.Field) *PagesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesDeleteCall) Context(ctx context.Context) *PagesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete a page by ID.",
	//   "httpMethod": "DELETE",
	//   "id": "blogger.pages.delete",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the Page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.get":

type PagesGetCall struct {
	s      *Service
	blogId string
	pageId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Gets one blog page by ID.
func (r *PagesService) Get(blogId string, pageId string) *PagesGetCall {
	c := &PagesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	return c
}

// View sets the optional parameter "view":
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PagesGetCall) View(view string) *PagesGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesGetCall) Fields(s ...googleapi.Field) *PagesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesGetCall) Context(ctx context.Context) *PagesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesGetCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one blog page by ID.",
	//   "httpMethod": "GET",
	//   "id": "blogger.pages.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog containing the page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the page to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}",
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.pages.insert":

type PagesInsertCall struct {
	s      *Service
	blogId string
	page   *Page
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Insert: Add a page.
func (r *PagesService) Insert(blogId string, page *Page) *PagesInsertCall {
	c := &PagesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.page = page
	return c
}

// IsDraft sets the optional parameter "isDraft": Whether to create the
// page as a draft (default: false).
func (c *PagesInsertCall) IsDraft(isDraft bool) *PagesInsertCall {
	c.opt_["isDraft"] = isDraft
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesInsertCall) Fields(s ...googleapi.Field) *PagesInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesInsertCall) Context(ctx context.Context) *PagesInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.page)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["isDraft"]; ok {
		params.Set("isDraft", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesInsertCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a page.",
	//   "httpMethod": "POST",
	//   "id": "blogger.pages.insert",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to add the page to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "isDraft": {
	//       "description": "Whether to create the page as a draft (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages",
	//   "request": {
	//     "$ref": "Page"
	//   },
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.list":

type PagesListCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// List: Retrieves the pages for a blog, optionally including non-LIVE
// statuses.
func (r *PagesService) List(blogId string) *PagesListCall {
	c := &PagesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether to
// retrieve the Page bodies.
func (c *PagesListCall) FetchBodies(fetchBodies bool) *PagesListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of Pages to fetch.
func (c *PagesListCall) MaxResults(maxResults int64) *PagesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if the request is paged.
func (c *PagesListCall) PageToken(pageToken string) *PagesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Status sets the optional parameter "status":
//
// Possible values:
//   "draft" - Draft (unpublished) Pages
//   "live" - Pages that are publicly visible
func (c *PagesListCall) Status(status string) *PagesListCall {
	c.opt_["status"] = status
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require elevated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PagesListCall) View(view string) *PagesListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesListCall) Fields(s ...googleapi.Field) *PagesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesListCall) Context(ctx context.Context) *PagesListCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesListCall) Do() (*PageList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PageList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the pages for a blog, optionally including non-LIVE statuses.",
	//   "httpMethod": "GET",
	//   "id": "blogger.pages.list",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch Pages from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether to retrieve the Page bodies.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of Pages to fetch.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if the request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "draft",
	//         "live"
	//       ],
	//       "enumDescriptions": [
	//         "Draft (unpublished) Pages",
	//         "Pages that are publicly visible"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages",
	//   "response": {
	//     "$ref": "PageList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.pages.patch":

type PagesPatchCall struct {
	s      *Service
	blogId string
	pageId string
	page   *Page
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Patch: Update a page. This method supports patch semantics.
func (r *PagesService) Patch(blogId string, pageId string, page *Page) *PagesPatchCall {
	c := &PagesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	c.page = page
	return c
}

// Publish sets the optional parameter "publish": Whether a publish
// action should be performed when the page is updated (default: false).
func (c *PagesPatchCall) Publish(publish bool) *PagesPatchCall {
	c.opt_["publish"] = publish
	return c
}

// Revert sets the optional parameter "revert": Whether a revert action
// should be performed when the page is updated (default: false).
func (c *PagesPatchCall) Revert(revert bool) *PagesPatchCall {
	c.opt_["revert"] = revert
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesPatchCall) Fields(s ...googleapi.Field) *PagesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesPatchCall) Context(ctx context.Context) *PagesPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.page)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["publish"]; ok {
		params.Set("publish", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revert"]; ok {
		params.Set("revert", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesPatchCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a page. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "blogger.pages.patch",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the Page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publish": {
	//       "description": "Whether a publish action should be performed when the page is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "revert": {
	//       "description": "Whether a revert action should be performed when the page is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}",
	//   "request": {
	//     "$ref": "Page"
	//   },
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.publish":

type PagesPublishCall struct {
	s      *Service
	blogId string
	pageId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Publish: Publishes a draft page.
func (r *PagesService) Publish(blogId string, pageId string) *PagesPublishCall {
	c := &PagesPublishCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesPublishCall) Fields(s ...googleapi.Field) *PagesPublishCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesPublishCall) Context(ctx context.Context) *PagesPublishCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesPublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}/publish")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesPublishCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Publishes a draft page.",
	//   "httpMethod": "POST",
	//   "id": "blogger.pages.publish",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}/publish",
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.revert":

type PagesRevertCall struct {
	s      *Service
	blogId string
	pageId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Revert: Revert a published or scheduled page to draft state.
func (r *PagesService) Revert(blogId string, pageId string) *PagesRevertCall {
	c := &PagesRevertCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesRevertCall) Fields(s ...googleapi.Field) *PagesRevertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesRevertCall) Context(ctx context.Context) *PagesRevertCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesRevertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}/revert")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesRevertCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Revert a published or scheduled page to draft state.",
	//   "httpMethod": "POST",
	//   "id": "blogger.pages.revert",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}/revert",
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.pages.update":

type PagesUpdateCall struct {
	s      *Service
	blogId string
	pageId string
	page   *Page
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Update: Update a page.
func (r *PagesService) Update(blogId string, pageId string, page *Page) *PagesUpdateCall {
	c := &PagesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	c.page = page
	return c
}

// Publish sets the optional parameter "publish": Whether a publish
// action should be performed when the page is updated (default: false).
func (c *PagesUpdateCall) Publish(publish bool) *PagesUpdateCall {
	c.opt_["publish"] = publish
	return c
}

// Revert sets the optional parameter "revert": Whether a revert action
// should be performed when the page is updated (default: false).
func (c *PagesUpdateCall) Revert(revert bool) *PagesUpdateCall {
	c.opt_["revert"] = revert
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagesUpdateCall) Fields(s ...googleapi.Field) *PagesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PagesUpdateCall) Context(ctx context.Context) *PagesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *PagesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.page)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["publish"]; ok {
		params.Set("publish", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revert"]; ok {
		params.Set("revert", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/pages/{pageId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"pageId": c.pageId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PagesUpdateCall) Do() (*Page, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Page
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a page.",
	//   "httpMethod": "PUT",
	//   "id": "blogger.pages.update",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the Page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publish": {
	//       "description": "Whether a publish action should be performed when the page is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "revert": {
	//       "description": "Whether a revert action should be performed when the page is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}",
	//   "request": {
	//     "$ref": "Page"
	//   },
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.postUserInfos.get":

type PostUserInfosGetCall struct {
	s      *Service
	userId string
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Gets one post and user info pair, by post ID and user ID. The
// post user info contains per-user information about the post, such as
// access rights, specific to the user.
func (r *PostUserInfosService) Get(userId string, blogId string, postId string) *PostUserInfosGetCall {
	c := &PostUserInfosGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.blogId = blogId
	c.postId = postId
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to pull back on a post.
func (c *PostUserInfosGetCall) MaxComments(maxComments int64) *PostUserInfosGetCall {
	c.opt_["maxComments"] = maxComments
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostUserInfosGetCall) Fields(s ...googleapi.Field) *PostUserInfosGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostUserInfosGetCall) Context(ctx context.Context) *PostUserInfosGetCall {
	c.ctx_ = ctx
	return c
}

func (c *PostUserInfosGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{userId}/blogs/{blogId}/posts/{postId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostUserInfosGetCall) Do() (*PostUserInfo, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PostUserInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one post and user info pair, by post ID and user ID. The post user info contains per-user information about the post, such as access rights, specific to the user.",
	//   "httpMethod": "GET",
	//   "id": "blogger.postUserInfos.get",
	//   "parameterOrder": [
	//     "userId",
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to pull back on a post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "The ID of the post to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user for the per-user information to be fetched. Either the word 'self' (sans quote marks) or the user's profile identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/blogs/{blogId}/posts/{postId}",
	//   "response": {
	//     "$ref": "PostUserInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.postUserInfos.list":

type PostUserInfosListCall struct {
	s      *Service
	userId string
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// List: Retrieves a list of post and post user info pairs, possibly
// filtered. The post user info contains per-user information about the
// post, such as access rights, specific to the user.
func (r *PostUserInfosService) List(userId string, blogId string) *PostUserInfosListCall {
	c := &PostUserInfosListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.blogId = blogId
	return c
}

// EndDate sets the optional parameter "endDate": Latest post date to
// fetch, a date-time with RFC 3339 formatting.
func (c *PostUserInfosListCall) EndDate(endDate string) *PostUserInfosListCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of posts is included. Default is false.
func (c *PostUserInfosListCall) FetchBodies(fetchBodies bool) *PostUserInfosListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// Labels sets the optional parameter "labels": Comma-separated list of
// labels to search for.
func (c *PostUserInfosListCall) Labels(labels string) *PostUserInfosListCall {
	c.opt_["labels"] = labels
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of posts to fetch.
func (c *PostUserInfosListCall) MaxResults(maxResults int64) *PostUserInfosListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort order applied to
// search results. Default is published.
//
// Possible values:
//   "published" - Order by the date the post was published
//   "updated" - Order by the date the post was last updated
func (c *PostUserInfosListCall) OrderBy(orderBy string) *PostUserInfosListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if the request is paged.
func (c *PostUserInfosListCall) PageToken(pageToken string) *PostUserInfosListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest post date
// to fetch, a date-time with RFC 3339 formatting.
func (c *PostUserInfosListCall) StartDate(startDate string) *PostUserInfosListCall {
	c.opt_["startDate"] = startDate
	return c
}

// Status sets the optional parameter "status":
//
// Possible values:
//   "draft" - Draft posts
//   "live" - Published posts
//   "scheduled" - Posts that are scheduled to publish in future.
func (c *PostUserInfosListCall) Status(status string) *PostUserInfosListCall {
	c.opt_["status"] = status
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require elevated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PostUserInfosListCall) View(view string) *PostUserInfosListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostUserInfosListCall) Fields(s ...googleapi.Field) *PostUserInfosListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostUserInfosListCall) Context(ctx context.Context) *PostUserInfosListCall {
	c.ctx_ = ctx
	return c
}

func (c *PostUserInfosListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{userId}/blogs/{blogId}/posts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostUserInfosListCall) Do() (*PostUserInfosList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PostUserInfosList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of post and post user info pairs, possibly filtered. The post user info contains per-user information about the post, such as access rights, specific to the user.",
	//   "httpMethod": "GET",
	//   "id": "blogger.postUserInfos.list",
	//   "parameterOrder": [
	//     "userId",
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "default": "false",
	//       "description": "Whether the body content of posts is included. Default is false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "labels": {
	//       "description": "Comma-separated list of labels to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of posts to fetch.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "default": "PUBLISHED",
	//       "description": "Sort order applied to search results. Default is published.",
	//       "enum": [
	//         "published",
	//         "updated"
	//       ],
	//       "enumDescriptions": [
	//         "Order by the date the post was published",
	//         "Order by the date the post was last updated"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if the request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "draft",
	//         "live",
	//         "scheduled"
	//       ],
	//       "enumDescriptions": [
	//         "Draft posts",
	//         "Published posts",
	//         "Posts that are scheduled to publish in future."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "ID of the user for the per-user information to be fetched. Either the word 'self' (sans quote marks) or the user's profile identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/blogs/{blogId}/posts",
	//   "response": {
	//     "$ref": "PostUserInfosList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.delete":

type PostsDeleteCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Delete: Delete a post by ID.
func (r *PostsService) Delete(blogId string, postId string) *PostsDeleteCall {
	c := &PostsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsDeleteCall) Fields(s ...googleapi.Field) *PostsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsDeleteCall) Context(ctx context.Context) *PostsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete a post by ID.",
	//   "httpMethod": "DELETE",
	//   "id": "blogger.posts.delete",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.get":

type PostsGetCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Get a post by ID.
func (r *PostsService) Get(blogId string, postId string) *PostsGetCall {
	c := &PostsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// FetchBody sets the optional parameter "fetchBody": Whether the body
// content of the post is included (default: true). This should be set
// to false when the post bodies are not required, to help minimize
// traffic.
func (c *PostsGetCall) FetchBody(fetchBody bool) *PostsGetCall {
	c.opt_["fetchBody"] = fetchBody
	return c
}

// FetchImages sets the optional parameter "fetchImages": Whether image
// URL metadata for each post is included (default: false).
func (c *PostsGetCall) FetchImages(fetchImages bool) *PostsGetCall {
	c.opt_["fetchImages"] = fetchImages
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to pull back on a post.
func (c *PostsGetCall) MaxComments(maxComments int64) *PostsGetCall {
	c.opt_["maxComments"] = maxComments
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require elevated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PostsGetCall) View(view string) *PostsGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsGetCall) Fields(s ...googleapi.Field) *PostsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsGetCall) Context(ctx context.Context) *PostsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchBody"]; ok {
		params.Set("fetchBody", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchImages"]; ok {
		params.Set("fetchImages", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsGetCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a post by ID.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBody": {
	//       "default": "true",
	//       "description": "Whether the body content of the post is included (default: true). This should be set to false when the post bodies are not required, to help minimize traffic.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "fetchImages": {
	//       "description": "Whether image URL metadata for each post is included (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to pull back on a post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "The ID of the post",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.getByPath":

type PostsGetByPathCall struct {
	s      *Service
	blogId string
	path   string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// GetByPath: Retrieve a Post by Path.
func (r *PostsService) GetByPath(blogId string, path string) *PostsGetByPathCall {
	c := &PostsGetByPathCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.path = path
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to pull back on a post.
func (c *PostsGetByPathCall) MaxComments(maxComments int64) *PostsGetByPathCall {
	c.opt_["maxComments"] = maxComments
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require elevated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PostsGetByPathCall) View(view string) *PostsGetByPathCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsGetByPathCall) Fields(s ...googleapi.Field) *PostsGetByPathCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsGetByPathCall) Context(ctx context.Context) *PostsGetByPathCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsGetByPathCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("path", fmt.Sprintf("%v", c.path))
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/bypath")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsGetByPathCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Post by Path.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.getByPath",
	//   "parameterOrder": [
	//     "blogId",
	//     "path"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to pull back on a post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "path": {
	//       "description": "Path of the Post to retrieve.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require elevated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/bypath",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.insert":

type PostsInsertCall struct {
	s      *Service
	blogId string
	post   *Post
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Insert: Add a post.
func (r *PostsService) Insert(blogId string, post *Post) *PostsInsertCall {
	c := &PostsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.post = post
	return c
}

// FetchBody sets the optional parameter "fetchBody": Whether the body
// content of the post is included with the result (default: true).
func (c *PostsInsertCall) FetchBody(fetchBody bool) *PostsInsertCall {
	c.opt_["fetchBody"] = fetchBody
	return c
}

// FetchImages sets the optional parameter "fetchImages": Whether image
// URL metadata for each post is included in the returned result
// (default: false).
func (c *PostsInsertCall) FetchImages(fetchImages bool) *PostsInsertCall {
	c.opt_["fetchImages"] = fetchImages
	return c
}

// IsDraft sets the optional parameter "isDraft": Whether to create the
// post as a draft (default: false).
func (c *PostsInsertCall) IsDraft(isDraft bool) *PostsInsertCall {
	c.opt_["isDraft"] = isDraft
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsInsertCall) Fields(s ...googleapi.Field) *PostsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsInsertCall) Context(ctx context.Context) *PostsInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchBody"]; ok {
		params.Set("fetchBody", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchImages"]; ok {
		params.Set("fetchImages", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["isDraft"]; ok {
		params.Set("isDraft", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsInsertCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a post.",
	//   "httpMethod": "POST",
	//   "id": "blogger.posts.insert",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to add the post to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBody": {
	//       "default": "true",
	//       "description": "Whether the body content of the post is included with the result (default: true).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "fetchImages": {
	//       "description": "Whether image URL metadata for each post is included in the returned result (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "isDraft": {
	//       "description": "Whether to create the post as a draft (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.list":

type PostsListCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// List: Retrieves a list of posts, possibly filtered.
func (r *PostsService) List(blogId string) *PostsListCall {
	c := &PostsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// EndDate sets the optional parameter "endDate": Latest post date to
// fetch, a date-time with RFC 3339 formatting.
func (c *PostsListCall) EndDate(endDate string) *PostsListCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of posts is included (default: true). This should be set
// to false when the post bodies are not required, to help minimize
// traffic.
func (c *PostsListCall) FetchBodies(fetchBodies bool) *PostsListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// FetchImages sets the optional parameter "fetchImages": Whether image
// URL metadata for each post is included.
func (c *PostsListCall) FetchImages(fetchImages bool) *PostsListCall {
	c.opt_["fetchImages"] = fetchImages
	return c
}

// Labels sets the optional parameter "labels": Comma-separated list of
// labels to search for.
func (c *PostsListCall) Labels(labels string) *PostsListCall {
	c.opt_["labels"] = labels
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of posts to fetch.
func (c *PostsListCall) MaxResults(maxResults int64) *PostsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort search results
//
// Possible values:
//   "published" - Order by the date the post was published
//   "updated" - Order by the date the post was last updated
func (c *PostsListCall) OrderBy(orderBy string) *PostsListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if the request is paged.
func (c *PostsListCall) PageToken(pageToken string) *PostsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest post date
// to fetch, a date-time with RFC 3339 formatting.
func (c *PostsListCall) StartDate(startDate string) *PostsListCall {
	c.opt_["startDate"] = startDate
	return c
}

// Status sets the optional parameter "status": Statuses to include in
// the results.
//
// Possible values:
//   "draft" - Draft (non-published) posts.
//   "live" - Published posts
//   "scheduled" - Posts that are scheduled to publish in the future.
func (c *PostsListCall) Status(status string) *PostsListCall {
	c.opt_["status"] = status
	return c
}

// View sets the optional parameter "view": Access level with which to
// view the returned result. Note that some fields require escalated
// access.
//
// Possible values:
//   "ADMIN" - Admin level detail
//   "AUTHOR" - Author level detail
//   "READER" - Reader level detail
func (c *PostsListCall) View(view string) *PostsListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsListCall) Fields(s ...googleapi.Field) *PostsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsListCall) Context(ctx context.Context) *PostsListCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchImages"]; ok {
		params.Set("fetchImages", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["status"]; ok {
		params.Set("status", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsListCall) Do() (*PostList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PostList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of posts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.list",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "default": "true",
	//       "description": "Whether the body content of posts is included (default: true). This should be set to false when the post bodies are not required, to help minimize traffic.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "fetchImages": {
	//       "description": "Whether image URL metadata for each post is included.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "labels": {
	//       "description": "Comma-separated list of labels to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of posts to fetch.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "default": "PUBLISHED",
	//       "description": "Sort search results",
	//       "enum": [
	//         "published",
	//         "updated"
	//       ],
	//       "enumDescriptions": [
	//         "Order by the date the post was published",
	//         "Order by the date the post was last updated"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if the request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "Statuses to include in the results.",
	//       "enum": [
	//         "draft",
	//         "live",
	//         "scheduled"
	//       ],
	//       "enumDescriptions": [
	//         "Draft (non-published) posts.",
	//         "Published posts",
	//         "Posts that are scheduled to publish in the future."
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Access level with which to view the returned result. Note that some fields require escalated access.",
	//       "enum": [
	//         "ADMIN",
	//         "AUTHOR",
	//         "READER"
	//       ],
	//       "enumDescriptions": [
	//         "Admin level detail",
	//         "Author level detail",
	//         "Reader level detail"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts",
	//   "response": {
	//     "$ref": "PostList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.patch":

type PostsPatchCall struct {
	s      *Service
	blogId string
	postId string
	post   *Post
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Patch: Update a post. This method supports patch semantics.
func (r *PostsService) Patch(blogId string, postId string, post *Post) *PostsPatchCall {
	c := &PostsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.post = post
	return c
}

// FetchBody sets the optional parameter "fetchBody": Whether the body
// content of the post is included with the result (default: true).
func (c *PostsPatchCall) FetchBody(fetchBody bool) *PostsPatchCall {
	c.opt_["fetchBody"] = fetchBody
	return c
}

// FetchImages sets the optional parameter "fetchImages": Whether image
// URL metadata for each post is included in the returned result
// (default: false).
func (c *PostsPatchCall) FetchImages(fetchImages bool) *PostsPatchCall {
	c.opt_["fetchImages"] = fetchImages
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to retrieve with the returned post.
func (c *PostsPatchCall) MaxComments(maxComments int64) *PostsPatchCall {
	c.opt_["maxComments"] = maxComments
	return c
}

// Publish sets the optional parameter "publish": Whether a publish
// action should be performed when the post is updated (default: false).
func (c *PostsPatchCall) Publish(publish bool) *PostsPatchCall {
	c.opt_["publish"] = publish
	return c
}

// Revert sets the optional parameter "revert": Whether a revert action
// should be performed when the post is updated (default: false).
func (c *PostsPatchCall) Revert(revert bool) *PostsPatchCall {
	c.opt_["revert"] = revert
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsPatchCall) Fields(s ...googleapi.Field) *PostsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsPatchCall) Context(ctx context.Context) *PostsPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchBody"]; ok {
		params.Set("fetchBody", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchImages"]; ok {
		params.Set("fetchImages", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publish"]; ok {
		params.Set("publish", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revert"]; ok {
		params.Set("revert", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsPatchCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a post. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "blogger.posts.patch",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBody": {
	//       "default": "true",
	//       "description": "Whether the body content of the post is included with the result (default: true).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "fetchImages": {
	//       "description": "Whether image URL metadata for each post is included in the returned result (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to retrieve with the returned post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publish": {
	//       "description": "Whether a publish action should be performed when the post is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "revert": {
	//       "description": "Whether a revert action should be performed when the post is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.publish":

type PostsPublishCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Publish: Publishes a draft post, optionally at the specific time of
// the given publishDate parameter.
func (r *PostsService) Publish(blogId string, postId string) *PostsPublishCall {
	c := &PostsPublishCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// PublishDate sets the optional parameter "publishDate": Optional date
// and time to schedule the publishing of the Blog. If no publishDate
// parameter is given, the post is either published at the a previously
// saved schedule date (if present), or the current time. If a future
// date is given, the post will be scheduled to be published.
func (c *PostsPublishCall) PublishDate(publishDate string) *PostsPublishCall {
	c.opt_["publishDate"] = publishDate
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsPublishCall) Fields(s ...googleapi.Field) *PostsPublishCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsPublishCall) Context(ctx context.Context) *PostsPublishCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsPublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["publishDate"]; ok {
		params.Set("publishDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/publish")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsPublishCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Publishes a draft post, optionally at the specific time of the given publishDate parameter.",
	//   "httpMethod": "POST",
	//   "id": "blogger.posts.publish",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publishDate": {
	//       "description": "Optional date and time to schedule the publishing of the Blog. If no publishDate parameter is given, the post is either published at the a previously saved schedule date (if present), or the current time. If a future date is given, the post will be scheduled to be published.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/publish",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.revert":

type PostsRevertCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Revert: Revert a published or scheduled post to draft state.
func (r *PostsService) Revert(blogId string, postId string) *PostsRevertCall {
	c := &PostsRevertCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsRevertCall) Fields(s ...googleapi.Field) *PostsRevertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsRevertCall) Context(ctx context.Context) *PostsRevertCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsRevertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}/revert")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsRevertCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Revert a published or scheduled post to draft state.",
	//   "httpMethod": "POST",
	//   "id": "blogger.posts.revert",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/revert",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.search":

type PostsSearchCall struct {
	s      *Service
	blogId string
	q      string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Search: Search for a post.
func (r *PostsService) Search(blogId string, q string) *PostsSearchCall {
	c := &PostsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.q = q
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of posts is included (default: true). This should be set
// to false when the post bodies are not required, to help minimize
// traffic.
func (c *PostsSearchCall) FetchBodies(fetchBodies bool) *PostsSearchCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort search results
//
// Possible values:
//   "published" - Order by the date the post was published
//   "updated" - Order by the date the post was last updated
func (c *PostsSearchCall) OrderBy(orderBy string) *PostsSearchCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsSearchCall) Fields(s ...googleapi.Field) *PostsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsSearchCall) Context(ctx context.Context) *PostsSearchCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsSearchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("q", fmt.Sprintf("%v", c.q))
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsSearchCall) Do() (*PostList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PostList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Search for a post.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.search",
	//   "parameterOrder": [
	//     "blogId",
	//     "q"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "default": "true",
	//       "description": "Whether the body content of posts is included (default: true). This should be set to false when the post bodies are not required, to help minimize traffic.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "orderBy": {
	//       "default": "PUBLISHED",
	//       "description": "Sort search results",
	//       "enum": [
	//         "published",
	//         "updated"
	//       ],
	//       "enumDescriptions": [
	//         "Order by the date the post was published",
	//         "Order by the date the post was last updated"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query terms to search this blog for matching posts.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/search",
	//   "response": {
	//     "$ref": "PostList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.update":

type PostsUpdateCall struct {
	s      *Service
	blogId string
	postId string
	post   *Post
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Update: Update a post.
func (r *PostsService) Update(blogId string, postId string, post *Post) *PostsUpdateCall {
	c := &PostsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.post = post
	return c
}

// FetchBody sets the optional parameter "fetchBody": Whether the body
// content of the post is included with the result (default: true).
func (c *PostsUpdateCall) FetchBody(fetchBody bool) *PostsUpdateCall {
	c.opt_["fetchBody"] = fetchBody
	return c
}

// FetchImages sets the optional parameter "fetchImages": Whether image
// URL metadata for each post is included in the returned result
// (default: false).
func (c *PostsUpdateCall) FetchImages(fetchImages bool) *PostsUpdateCall {
	c.opt_["fetchImages"] = fetchImages
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to retrieve with the returned post.
func (c *PostsUpdateCall) MaxComments(maxComments int64) *PostsUpdateCall {
	c.opt_["maxComments"] = maxComments
	return c
}

// Publish sets the optional parameter "publish": Whether a publish
// action should be performed when the post is updated (default: false).
func (c *PostsUpdateCall) Publish(publish bool) *PostsUpdateCall {
	c.opt_["publish"] = publish
	return c
}

// Revert sets the optional parameter "revert": Whether a revert action
// should be performed when the post is updated (default: false).
func (c *PostsUpdateCall) Revert(revert bool) *PostsUpdateCall {
	c.opt_["revert"] = revert
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostsUpdateCall) Fields(s ...googleapi.Field) *PostsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *PostsUpdateCall) Context(ctx context.Context) *PostsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *PostsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fetchBody"]; ok {
		params.Set("fetchBody", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchImages"]; ok {
		params.Set("fetchImages", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["publish"]; ok {
		params.Set("publish", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revert"]; ok {
		params.Set("revert", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "blogs/{blogId}/posts/{postId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"blogId": c.blogId,
		"postId": c.postId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *PostsUpdateCall) Do() (*Post, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Post
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a post.",
	//   "httpMethod": "PUT",
	//   "id": "blogger.posts.update",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBody": {
	//       "default": "true",
	//       "description": "Whether the body content of the post is included with the result (default: true).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "fetchImages": {
	//       "description": "Whether image URL metadata for each post is included in the returned result (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to retrieve with the returned post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publish": {
	//       "description": "Whether a publish action should be performed when the post is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "revert": {
	//       "description": "Whether a revert action should be performed when the post is updated (default: false).",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.users.get":

type UsersGetCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
	ctx_   context.Context
}

// Get: Gets one user by ID.
func (r *UsersService) Get(userId string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *UsersGetCall) Context(ctx context.Context) *UsersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *UsersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *UsersGetCall) Do() (*User, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *User
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one user by ID.",
	//   "httpMethod": "GET",
	//   "id": "blogger.users.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "The ID of the user to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}