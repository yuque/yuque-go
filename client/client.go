package client

import (
  "encoding/json"

  "github.com/alibabacloud-go/tea/tea"
  "github.com/yuque/yuque-go/baseclient"
)

type Config struct {
  AuthToken *string `json:"authToken" xml:"authToken" require:"true"`
  Domain *string `json:"domain" xml:"domain" require:"true"`
}

func (s Config) String() string {
  return utils.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

func (s *Config) SetAuthToken(v string) *Config {
  s.AuthToken = &v
  return s
}

func (s *Config) SetDomain(v string) *Config {
  s.Domain = &v
  return s
}

type HelloResponse struct {
  Data *Hello `json:"data" xml:"data" require:"true"`
}

func (s HelloResponse) String() string {
  return utils.Prettify(s)
}

func (s HelloResponse) GoString() string {
  return s.String()
}

func (s *HelloResponse) SetData(v *Hello) *HelloResponse {
  s.Data = &v
  return s
}

type Hello struct {
  Message *string `json:"message" xml:"message" require:"true"`
}

func (s Hello) String() string {
  return utils.Prettify(s)
}

func (s Hello) GoString() string {
  return s.String()
}

func (s *Hello) SetMessage(v string) *Hello {
  s.Message = &v
  return s
}

type User struct {
  Id *int `json:"id" xml:"id" require:"true"`
  Type *string `json:"type" xml:"type" require:"true"`
  Login *string `json:"login" xml:"login" require:"true"`
  Name *string `json:"name" xml:"name" require:"true"`
  Description *string `json:"description" xml:"description" require:"true"`
  AvatarUrl *string `json:"avatar_url" xml:"avatar_url" require:"true"`
  CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
  UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}

func (s User) String() string {
  return utils.Prettify(s)
}

func (s User) GoString() string {
  return s.String()
}

func (s *User) SetId(v int) *User {
  s.Id = &v
  return s
}

func (s *User) SetType(v string) *User {
  s.Type = &v
  return s
}

func (s *User) SetLogin(v string) *User {
  s.Login = &v
  return s
}

func (s *User) SetName(v string) *User {
  s.Name = &v
  return s
}

func (s *User) SetDescription(v string) *User {
  s.Description = &v
  return s
}

func (s *User) SetAvatarUrl(v string) *User {
  s.AvatarUrl = &v
  return s
}

func (s *User) SetCreatedAt(v string) *User {
  s.CreatedAt = &v
  return s
}

func (s *User) SetUpdatedAt(v string) *User {
  s.UpdatedAt = &v
  return s
}

type UserResponse struct {
  Data *User `json:"data" xml:"data" require:"true"`
}

func (s UserResponse) String() string {
  return utils.Prettify(s)
}

func (s UserResponse) GoString() string {
  return s.String()
}

func (s *UserResponse) SetData(v *User) *UserResponse {
  s.Data = &v
  return s
}

type Group struct {
  Id *string `json:"id" xml:"id"`
  Login *string `json:"login" xml:"login" require:"true"`
  Name *string `json:"name" xml:"name" require:"true"`
  AvatarUrl *string `json:"avatar_url" xml:"avatar_url"`
  OwnerId *string `json:"owner_id" xml:"owner_id" require:"true"`
  Description *string `json:"description" xml:"description" require:"true"`
  CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
  UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}

func (s Group) String() string {
  return utils.Prettify(s)
}

func (s Group) GoString() string {
  return s.String()
}

func (s *Group) SetId(v string) *Group {
  s.Id = &v
  return s
}

func (s *Group) SetLogin(v string) *Group {
  s.Login = &v
  return s
}

func (s *Group) SetName(v string) *Group {
  s.Name = &v
  return s
}

func (s *Group) SetAvatarUrl(v string) *Group {
  s.AvatarUrl = &v
  return s
}

func (s *Group) SetOwnerId(v string) *Group {
  s.OwnerId = &v
  return s
}

func (s *Group) SetDescription(v string) *Group {
  s.Description = &v
  return s
}

func (s *Group) SetCreatedAt(v string) *Group {
  s.CreatedAt = &v
  return s
}

func (s *Group) SetUpdatedAt(v string) *Group {
  s.UpdatedAt = &v
  return s
}

type GroupsResponse struct {
  Data []*Group `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s GroupsResponse) String() string {
  return utils.Prettify(s)
}

func (s GroupsResponse) GoString() string {
  return s.String()
}

func (s *GroupsResponse) SetData(v []*Group) *GroupsResponse {
  s.Data = v
  return s
}

type UsersResponse struct {
  Data []*User `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s UsersResponse) String() string {
  return utils.Prettify(s)
}

func (s UsersResponse) GoString() string {
  return s.String()
}

func (s *UsersResponse) SetData(v []*User) *UsersResponse {
  s.Data = v
  return s
}

type GroupResponse struct {
  Data *Group `json:"data" xml:"data" require:"true"`
}

func (s GroupResponse) String() string {
  return utils.Prettify(s)
}

func (s GroupResponse) GoString() string {
  return s.String()
}

func (s *GroupResponse) SetData(v *Group) *GroupResponse {
  s.Data = &v
  return s
}

type GroupDetailResponse struct {
  Abilities *GroupDetailResponseAbilities `json:"abilities" xml:"abilities" require:"true" type:"Struct"`
  Data []*Group `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s GroupDetailResponse) String() string {
  return utils.Prettify(s)
}

func (s GroupDetailResponse) GoString() string {
  return s.String()
}

func (s *GroupDetailResponse) SetAbilities(v *GroupDetailResponseAbilities) *GroupDetailResponse {
  s.Abilities = v
  return s
}

func (s *GroupDetailResponse) SetData(v []*Group) *GroupDetailResponse {
  s.Data = v
  return s
}

type GroupDetailResponseAbilities struct {
  Read *bool `json:"read" xml:"read" require:"true"`
  Update *bool `json:"update" xml:"update" require:"true"`
  Destroy *bool `json:"destroy" xml:"destroy" require:"true"`
  GroupUser *GroupDetailResponseAbilitiesGroupUser `json:"group_user" xml:"group_user" require:"true" type:"Struct"`
  Repo *GroupDetailResponseAbilitiesRepo `json:"repo" xml:"repo" require:"true" type:"Struct"`
}

func (s GroupDetailResponseAbilities) String() string {
  return utils.Prettify(s)
}

func (s GroupDetailResponseAbilities) GoString() string {
  return s.String()
}

func (s *GroupDetailResponseAbilities) SetRead(v bool) *GroupDetailResponseAbilities {
  s.Read = &v
  return s
}

func (s *GroupDetailResponseAbilities) SetUpdate(v bool) *GroupDetailResponseAbilities {
  s.Update = &v
  return s
}

func (s *GroupDetailResponseAbilities) SetDestroy(v bool) *GroupDetailResponseAbilities {
  s.Destroy = &v
  return s
}

func (s *GroupDetailResponseAbilities) SetGroupUser(v *GroupDetailResponseAbilitiesGroupUser) *GroupDetailResponseAbilities {
  s.GroupUser = v
  return s
}

func (s *GroupDetailResponseAbilities) SetRepo(v *GroupDetailResponseAbilitiesRepo) *GroupDetailResponseAbilities {
  s.Repo = v
  return s
}

type GroupDetailResponseAbilitiesGroupUser struct {
  Create *bool `json:"create" xml:"create" require:"true"`
  Update *bool `json:"update" xml:"update" require:"true"`
  Destroy *bool `json:"destroy" xml:"destroy" require:"true"`
}

func (s GroupDetailResponseAbilitiesGroupUser) String() string {
  return utils.Prettify(s)
}

func (s GroupDetailResponseAbilitiesGroupUser) GoString() string {
  return s.String()
}

func (s *GroupDetailResponseAbilitiesGroupUser) SetCreate(v bool) *GroupDetailResponseAbilitiesGroupUser {
  s.Create = &v
  return s
}

func (s *GroupDetailResponseAbilitiesGroupUser) SetUpdate(v bool) *GroupDetailResponseAbilitiesGroupUser {
  s.Update = &v
  return s
}

func (s *GroupDetailResponseAbilitiesGroupUser) SetDestroy(v bool) *GroupDetailResponseAbilitiesGroupUser {
  s.Destroy = &v
  return s
}

type GroupDetailResponseAbilitiesRepo struct {
  Create *bool `json:"create" xml:"create" require:"true"`
  Update *bool `json:"update" xml:"update" require:"true"`
  Destroy *bool `json:"destroy" xml:"destroy" require:"true"`
}

func (s GroupDetailResponseAbilitiesRepo) String() string {
  return utils.Prettify(s)
}

func (s GroupDetailResponseAbilitiesRepo) GoString() string {
  return s.String()
}

func (s *GroupDetailResponseAbilitiesRepo) SetCreate(v bool) *GroupDetailResponseAbilitiesRepo {
  s.Create = &v
  return s
}

func (s *GroupDetailResponseAbilitiesRepo) SetUpdate(v bool) *GroupDetailResponseAbilitiesRepo {
  s.Update = &v
  return s
}

func (s *GroupDetailResponseAbilitiesRepo) SetDestroy(v bool) *GroupDetailResponseAbilitiesRepo {
  s.Destroy = &v
  return s
}

type GroupUser struct {
  Id *int `json:"id" xml:"id" require:"true"`
  GroupId *int `json:"group_id" xml:"group_id" require:"true"`
  UserId *int `json:"user_id" xml:"user_id" require:"true"`
  Role *int `json:"role" xml:"role" require:"true"`
  User *User `json:"user" xml:"user" require:"true"`
  CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
  UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}

func (s GroupUser) String() string {
  return utils.Prettify(s)
}

func (s GroupUser) GoString() string {
  return s.String()
}

func (s *GroupUser) SetId(v int) *GroupUser {
  s.Id = &v
  return s
}

func (s *GroupUser) SetGroupId(v int) *GroupUser {
  s.GroupId = &v
  return s
}

func (s *GroupUser) SetUserId(v int) *GroupUser {
  s.UserId = &v
  return s
}

func (s *GroupUser) SetRole(v int) *GroupUser {
  s.Role = &v
  return s
}

func (s *GroupUser) SetUser(v *User) *GroupUser {
  s.User = &v
  return s
}

func (s *GroupUser) SetCreatedAt(v string) *GroupUser {
  s.CreatedAt = &v
  return s
}

func (s *GroupUser) SetUpdatedAt(v string) *GroupUser {
  s.UpdatedAt = &v
  return s
}

type GroupUsersResponse struct {
  Data []*GroupUser `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s GroupUsersResponse) String() string {
  return utils.Prettify(s)
}

func (s GroupUsersResponse) GoString() string {
  return s.String()
}

func (s *GroupUsersResponse) SetData(v []*GroupUser) *GroupUsersResponse {
  s.Data = v
  return s
}

type GroupUserResponse struct {
  Data *GroupUser `json:"data" xml:"data" require:"true"`
}

func (s GroupUserResponse) String() string {
  return utils.Prettify(s)
}

func (s GroupUserResponse) GoString() string {
  return s.String()
}

func (s *GroupUserResponse) SetData(v *GroupUser) *GroupUserResponse {
  s.Data = &v
  return s
}

type Book struct {
  Id *int `json:"id" xml:"id" require:"true"`
  Type *string `json:"type" xml:"type" require:"true"`
  Slug *string `json:"slug" xml:"slug" require:"true"`
  Name *string `json:"name" xml:"name" require:"true"`
  Namespace *string `json:"namespace" xml:"namespace" require:"true"`
  UserId *int `json:"user_id" xml:"user_id" require:"true"`
  User *User `json:"user" xml:"user" require:"true"`
  Description *string `json:"description" xml:"description" require:"true"`
  CreatorId *int `json:"creator_id" xml:"creator_id" require:"true"`
  Public *int `json:"public" xml:"public" require:"true"`
  LikesCount *int `json:"likes_count" xml:"likes_count" require:"true"`
  WatchesCount *int `json:"watches_count" xml:"watches_count" require:"true"`
  CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
  UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}

func (s Book) String() string {
  return utils.Prettify(s)
}

func (s Book) GoString() string {
  return s.String()
}

func (s *Book) SetId(v int) *Book {
  s.Id = &v
  return s
}

func (s *Book) SetType(v string) *Book {
  s.Type = &v
  return s
}

func (s *Book) SetSlug(v string) *Book {
  s.Slug = &v
  return s
}

func (s *Book) SetName(v string) *Book {
  s.Name = &v
  return s
}

func (s *Book) SetNamespace(v string) *Book {
  s.Namespace = &v
  return s
}

func (s *Book) SetUserId(v int) *Book {
  s.UserId = &v
  return s
}

func (s *Book) SetUser(v *User) *Book {
  s.User = &v
  return s
}

func (s *Book) SetDescription(v string) *Book {
  s.Description = &v
  return s
}

func (s *Book) SetCreatorId(v int) *Book {
  s.CreatorId = &v
  return s
}

func (s *Book) SetPublic(v int) *Book {
  s.Public = &v
  return s
}

func (s *Book) SetLikesCount(v int) *Book {
  s.LikesCount = &v
  return s
}

func (s *Book) SetWatchesCount(v int) *Book {
  s.WatchesCount = &v
  return s
}

func (s *Book) SetCreatedAt(v string) *Book {
  s.CreatedAt = &v
  return s
}

func (s *Book) SetUpdatedAt(v string) *Book {
  s.UpdatedAt = &v
  return s
}

type BooksResponse struct {
  Data []*Book `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s BooksResponse) String() string {
  return utils.Prettify(s)
}

func (s BooksResponse) GoString() string {
  return s.String()
}

func (s *BooksResponse) SetData(v []*Book) *BooksResponse {
  s.Data = v
  return s
}

type BookResponse struct {
  Data *Book `json:"data" xml:"data" require:"true"`
}

func (s BookResponse) String() string {
  return utils.Prettify(s)
}

func (s BookResponse) GoString() string {
  return s.String()
}

func (s *BookResponse) SetData(v *Book) *BookResponse {
  s.Data = &v
  return s
}

type BookToc struct {
  Title *string `json:"title" xml:"title" require:"true"`
  Slug *string `json:"slug" xml:"slug" require:"true"`
  Depth *int `json:"depth" xml:"depth" require:"true"`
}

func (s BookToc) String() string {
  return utils.Prettify(s)
}

func (s BookToc) GoString() string {
  return s.String()
}

func (s *BookToc) SetTitle(v string) *BookToc {
  s.Title = &v
  return s
}

func (s *BookToc) SetSlug(v string) *BookToc {
  s.Slug = &v
  return s
}

func (s *BookToc) SetDepth(v int) *BookToc {
  s.Depth = &v
  return s
}

type BookTocsResponse struct {
  Data []*BookToc `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s BookTocsResponse) String() string {
  return utils.Prettify(s)
}

func (s BookTocsResponse) GoString() string {
  return s.String()
}

func (s *BookTocsResponse) SetData(v []*BookToc) *BookTocsResponse {
  s.Data = v
  return s
}

type SearchQuery struct {
  Query *string `json:"query" xml:"query" require:"true"`
  Type *string `json:"type" xml:"type"`
}

func (s SearchQuery) String() string {
  return utils.Prettify(s)
}

func (s SearchQuery) GoString() string {
  return s.String()
}

func (s *SearchQuery) SetQuery(v string) *SearchQuery {
  s.Query = &v
  return s
}

func (s *SearchQuery) SetType(v string) *SearchQuery {
  s.Type = &v
  return s
}

type Doc struct {
  Id *int `json:"id" xml:"id" require:"true"`
  Slug *string `json:"slug" xml:"slug" require:"true"`
  Title *string `json:"title" xml:"title" require:"true"`
  BookId *int `json:"book_id" xml:"book_id" require:"true"`
  Book *Book `json:"book" xml:"book" require:"true"`
  UserId *int `json:"user_id" xml:"user_id" require:"true"`
  User *User `json:"user" xml:"user" require:"true"`
  Format *string `json:"format" xml:"format" require:"true"`
  Body *string `json:"body" xml:"body" require:"true"`
  BodyDraft *string `json:"body_draft" xml:"body_draft" require:"true"`
  BodyHtml *string `json:"body_html" xml:"body_html" require:"true"`
  BodyLake *string `json:"body_lake" xml:"body_lake" require:"true"`
  CreatorId *int `json:"creator_id" xml:"creator_id" require:"true"`
  Public *int `json:"public" xml:"public" require:"true"`
  Status *int `json:"status" xml:"status" require:"true"`
  LikesCount *int `json:"likes_count" xml:"likes_count" require:"true"`
  CommentsCount *int `json:"comments_count" xml:"comments_count" require:"true"`
  ContentUpdatedAt *string `json:"content_updated_at" xml:"content_updated_at" require:"true"`
  DeletedAt *string `json:"deleted_at" xml:"deleted_at" require:"true"`
  CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
  UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}

func (s Doc) String() string {
  return utils.Prettify(s)
}

func (s Doc) GoString() string {
  return s.String()
}

func (s *Doc) SetId(v int) *Doc {
  s.Id = &v
  return s
}

func (s *Doc) SetSlug(v string) *Doc {
  s.Slug = &v
  return s
}

func (s *Doc) SetTitle(v string) *Doc {
  s.Title = &v
  return s
}

func (s *Doc) SetBookId(v int) *Doc {
  s.BookId = &v
  return s
}

func (s *Doc) SetBook(v *Book) *Doc {
  s.Book = &v
  return s
}

func (s *Doc) SetUserId(v int) *Doc {
  s.UserId = &v
  return s
}

func (s *Doc) SetUser(v *User) *Doc {
  s.User = &v
  return s
}

func (s *Doc) SetFormat(v string) *Doc {
  s.Format = &v
  return s
}

func (s *Doc) SetBody(v string) *Doc {
  s.Body = &v
  return s
}

func (s *Doc) SetBodyDraft(v string) *Doc {
  s.BodyDraft = &v
  return s
}

func (s *Doc) SetBodyHtml(v string) *Doc {
  s.BodyHtml = &v
  return s
}

func (s *Doc) SetBodyLake(v string) *Doc {
  s.BodyLake = &v
  return s
}

func (s *Doc) SetCreatorId(v int) *Doc {
  s.CreatorId = &v
  return s
}

func (s *Doc) SetPublic(v int) *Doc {
  s.Public = &v
  return s
}

func (s *Doc) SetStatus(v int) *Doc {
  s.Status = &v
  return s
}

func (s *Doc) SetLikesCount(v int) *Doc {
  s.LikesCount = &v
  return s
}

func (s *Doc) SetCommentsCount(v int) *Doc {
  s.CommentsCount = &v
  return s
}

func (s *Doc) SetContentUpdatedAt(v string) *Doc {
  s.ContentUpdatedAt = &v
  return s
}

func (s *Doc) SetDeletedAt(v string) *Doc {
  s.DeletedAt = &v
  return s
}

func (s *Doc) SetCreatedAt(v string) *Doc {
  s.CreatedAt = &v
  return s
}

func (s *Doc) SetUpdatedAt(v string) *Doc {
  s.UpdatedAt = &v
  return s
}

type DocsResponse struct {
  Data []*Doc `json:"data" xml:"data" require:"true" type:"Repeated"`
}

func (s DocsResponse) String() string {
  return utils.Prettify(s)
}

func (s DocsResponse) GoString() string {
  return s.String()
}

func (s *DocsResponse) SetData(v []*Doc) *DocsResponse {
  s.Data = v
  return s
}

type DocResponse struct {
  Data *Doc `json:"data" xml:"data" require:"true"`
}

func (s DocResponse) String() string {
  return utils.Prettify(s)
}

func (s DocResponse) GoString() string {
  return s.String()
}

func (s *DocResponse) SetData(v *Doc) *DocResponse {
  s.Data = &v
  return s
}

type Client struct {
  baseclient.BaseClient
}

func NewClient(config *Config)(*Client, error){
  client := &Client{}
  input := make(map[string]string)
  byt, _ := json.Marshal(config)
  err := json.Unmarshal(byt, &input)
if err != nil {
    return nil, err
  }
  err = client.InitClient(input)
if err != nil {
    return nil, err
  }
  return client, nil
}
func (client *Client) Get(path string, params map[string]interface{}) (map[string]interface{}, error) {
  request_ := tea.NewRequest()
  request_.Protocol = "https"
  request_.Method = "GET"
  request_.Pathname = /api/v2 + path
  request_.Query = client.ToQuery(params)
  request_.Headers = map[string]string{
    "host": client.Domain,
    "x-auth-token": client.AuthToken,
    "user-agent": "tea/nodejs",
  }
  response_, err := tea.DoRequest(request_, nil)
  if err != nil {
    return nil ,err
  }

  result, err := client.ReadJSON(response_)
  if err != nil {
    return nil ,err
  }
  return result, err
}


func (client *Client) Post(path string, data map[string]interface{}) (map[string]interface{}, error) {
  request_ := tea.NewRequest()
  request_.Method = "POST"
  request_.Pathname = /api/v2 + path
  request_.Headers = map[string]string{
    "host": client.Domain,
    "x-auth-token": client.AuthToken,
    "user-agent": "tea/nodejs",
  }
  request_.Body = client.ToJSON(data)
  response_, err := tea.DoRequest(request_, nil)
  if err != nil {
    return nil ,err
  }

  result, err := client.ReadJSON(response_)
  if err != nil {
    return nil ,err
  }
  return result, err
}


func (client *Client) Put(path string, data map[string]interface{}) (map[string]interface{}, error) {
  request_ := tea.NewRequest()
  request_.Method = "PUT"
  request_.Pathname = /api/v2 + path
  request_.Headers = map[string]string{
    "host": client.Domain,
    "x-auth-token": client.AuthToken,
    "user-agent": "tea/nodejs",
  }
  request_.Body = client.ToJSON(data)
  response_, err := tea.DoRequest(request_, nil)
  if err != nil {
    return nil ,err
  }

  result, err := client.ReadJSON(response_)
  if err != nil {
    return nil ,err
  }
  return result, err
}


func (client *Client) Destroy(path string, data map[string]interface{}) (map[string]interface{}, error) {
  request_ := tea.NewRequest()
  request_.Method = "DELETE"
  request_.Pathname = /api/v2 + path
  request_.Headers = map[string]string{
    "host": client.Domain,
    "x-auth-token": client.AuthToken,
    "user-agent": "tea/nodejs",
  }
  request_.Body = client.ToJSON(data)
  response_, err := tea.DoRequest(request_, nil)
  if err != nil {
    return nil ,err
  }

  result, err := client.ReadJSON(response_)
  if err != nil {
    return nil ,err
  }
  return result, err
}


func (client *Client) Hello () (*HelloResponse, error) {
  _result := &HelloResponse{}
  _body, err := client.get("/hello", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetUserById (login string) (*UserResponse, error) {
  _result := &UserResponse{}
  _body, err := client.get("/users/" + login, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetUser () (*UserResponse, error) {
  _result := &UserResponse{}
  _body, err := client.get("/user", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetGroupsByUser (login string) (*GroupsResponse, error) {
  _result := &GroupsResponse{}
  _body, err := client.get("/users/" + login + "/groups", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetGroups () (*UsersResponse, error) {
  _result := &UsersResponse{}
  _body, err := client.get("/groups", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) CreateGroup (group *Group) (*GroupResponse, error) {
  _result := &GroupResponse{}
  _body, err := client.post("/groups", tea.To(group))
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetGroup (groupid string) (*GroupDetailResponse, error) {
  _result := &GroupDetailResponse{}
  _body, err := client.get("/groups/" + groupid, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) UpdateGroup (groupid string, params map[string]interface{}) (*GroupResponse, error) {
  _result := &GroupResponse{}
  _body, err := client.put("/groups/" + groupid, params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) DestryGroup (groupid string) (map[string]interface{}, error) {
  _result := make(map[string]interface{})
  _body, err := client.destroy("/groups/" + groupid, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetGroupMembers (groupid string) (*GroupUsersResponse, error) {
  _result := &GroupUsersResponse{}
  _body, err := client.get("/groups/" + groupid + "/users", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) UpdateGroupMember (groupid string, userid string, params map[string]interface{}) (*GroupUserResponse, error) {
  _result := &GroupUserResponse{}
  _body, err := client.put("/groups/" + groupid + "/users/" + userid, params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) DestroyGroupMember (groupid string, userid string) (map[string]interface{}, error) {
  _result := make(map[string]interface{})
  _body, err := client.destroy("/groups/" + groupid + "/users/" + userid, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetBooksByUser (userid string, params map[string]interface{}) (*BooksResponse, error) {
  _result := &BooksResponse{}
  _body, err := client.get("/users/" + userid + "/repos", params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) CreateBook (userid string, params map[string]interface{}) (*BooksResponse, error) {
  _result := &BooksResponse{}
  _body, err := client.post("/groups/" + userid + "/repos", params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetBook (namespace string) (*BookResponse, error) {
  _result := &BookResponse{}
  _body, err := client.get("/repos/" + namespace, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) UpdateBook (namespace string, params map[string]interface{}) (*BookResponse, error) {
  _result := &BookResponse{}
  _body, err := client.put("/repos/" + namespace, params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) DestroyBook (namespace string) (map[string]interface{}, error) {
  _result := make(map[string]interface{})
  _body, err := client.destroy("/repos/" + namespace, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetBookToc (namespace string) (*BookTocsResponse, error) {
  _result := &BookTocsResponse{}
  _body, err := client.get("/repos/" + namespace + "/toc", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) SearchBook (params *SearchQuery) (*BooksResponse, error) {
  _result := &BooksResponse{}
  _body, err := client.get("/search/repos", tea.ToObject(params))
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetDocs (namespace string) (*DocsResponse, error) {
  _result := &DocsResponse{}
  _body, err := client.get("/repos/" + namespace + "/docs", map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) GetDoc (namespace string, slug string) (*DocResponse, error) {
  _result := &DocResponse{}
  _body, err := client.get("/repos/" + namespace + "/docs/" + slug, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) CreateDoc (namespace string, params map[string]interface{}) (*DocResponse, error) {
  _result := &DocResponse{}
  _body, err := client.post("/repos/" + namespace + "/docs", params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) UpdateDoc (namespace string, slug string, params map[string]interface{}) (*DocResponse, error) {
  _result := &DocResponse{}
  _body, err := client.put("/repos/" + namespace + "/docs/" + slug, params)
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}


func (client *Client) DestroyDoc (namespace string, slug string) (map[string]interface{}, error) {
  _result := make(map[string]interface{})
  _body, err := client.destroy("/repos/" + namespace + "/docs/" + slug, map[string]interface{}{})
  if err != nil {
    return nil, err
  }
  err = tea.Convert(_body, &_result)
  return _result, err
}

