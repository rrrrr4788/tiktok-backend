// Code generated by hertz generator.

package interaction

import (
	"BiteDans.com/tiktok-backend/biz/dal/model"
	"BiteDans.com/tiktok-backend/pkg/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	interaction "BiteDans.com/tiktok-backend/biz/model/douyin/extra/interaction"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteInteraction .
// @router /douyin/favorite/action/ [POST]
func FavoriteInteraction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentInteraction .
// @router /douyin/comment/action/ [POST]
func CommentInteraction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.DouyinCommentActionResponse)

	var user_id uint
	if user_id, err = utils.GetIdFromToken(req.Token); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Invalid token"
		resp.Comment = nil

		c.JSON(consts.StatusUnauthorized, resp)
		return
	}

	_user := new(model.User)
	if err = model.FindUserById(_user, user_id); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "User id does not exist"
		resp.Comment = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	_video := new(model.Video)
	if err = model.FindVideoById(_video, uint(req.VideoId)); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Video id does not exist"
		resp.Comment = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	if req.ActionType == 1 {
		comment := new(model.Comment)
		comment.UserId = int64(user_id)
		comment.VideoId = req.VideoId
		comment.Content = req.CommentText
		if err = model.CreateComment(comment); err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "Failed to create comment"
			resp.Comment = nil
			c.JSON(consts.StatusInternalServerError, resp)

			hlog.Error("Failed to create comment into database")
			return
		}
		resp.StatusCode = 0
		resp.StatusMsg = "comment on video successfully!"
		resp.Comment = &interaction.Comment{
			ID:         0,
			User:       nil,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.String(),
		}
	} else if req.ActionType == 2 {
		comment := new(model.Comment)
		if comment, err = model.FindCommentById(req.CommentId); err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "Comment id does not exist"
			resp.Comment = nil
			c.JSON(consts.StatusBadRequest, resp)
			return
		}

		if err = model.DeleteComment(comment); err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "Fail to delete comment"
			resp.Comment = nil
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		resp.StatusCode = 0
		resp.StatusMsg = "delete comment successfully"
		resp.Comment = nil
	} else {
		resp.StatusCode = -1
		resp.StatusMsg = "Fail to get action type"
		resp.Comment = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}
