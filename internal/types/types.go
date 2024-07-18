// Code generated by goctl. DO NOT EDIT.
package types

type AddCommentReq struct {
	ArticleId int64  `json:"articleId"`
	Body      string `json:"body"`
}

type AddTagReq struct {
	TagName string `json:"tagName"`
}

type Article struct {
	Id             int64    `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	TagList        []string `json:"tagList"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int64    `json:"favoritesCount"`
	Author         Profile  `json:"author"`
}

type ArticleList struct {
	Articles      []Article `json:"articles"`
	ArticlesCount int64     `json:"articlesCount"`
}

type Comment struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Body      string `json:"body"`
	Author    User   `json:"author"`
}

type CreateArticleReq struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Body        string  `json:"body"`
	TagList     []int64 `json:"tagList"`
}

type DelCommentReq struct {
	ArticleId int64 `json:"articleId"`
	CommentId int64 `json:"commentId"`
}

type DeleteArticleReq struct {
	ArticleId int64 `json:"articleId"`
}

type FavoriteArticleReq struct {
	ArticleId int64 `json:"articleId"`
}

type FollowUserReq struct {
	UserId string `json:"userId"`
}

type GetArticleListReq struct {
	Tag       string `json:"tag"`
	Author    string `json:"author"`
	Favorited string `json:"favorited"`
	Limit     int64  `json:"limit"`
	Offset    int64  `json:"offset"`
}

type GetArticleReq struct {
	ArticleId int64 `json:"articleId"`
}

type GetCommentListReq struct {
	ArticleId int64 `json:"articleId"`
}

type GetProfileReq struct {
	UserId string `json:"userId"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	UserId    int64  `json:"userId"`
	UserName  string `json:"username"`
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
}

type RegistrationReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tag struct {
	Id      int64  `json:"id"`
	TagName string `json:"tagName"`
}

type UnfavoriteArticleReq struct {
	ArticleId int64 `json:"articleId"`
}

type UnfollowUserReq struct {
	UserId string `json:"userId"`
}

type UpdateArticleReq struct {
	ArticleId int64  `json:"articleId"`
	Body      string `json:"body"`
}

type UpdatePasswordReq struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UpdateUserReq struct {
	Email string `json:"email"`
	Bio   string `json:"bio"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	Bio      string `json:"bio"`
}
