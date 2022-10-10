package query

import (
	"context"
	"database/sql"
	"social-network/common"
	"social-network/common/application"
	"social-network/db"
	"social-network/model"
)

type FriendsQuery struct {
	UserId int64 `param:"userId"`
	Limit int64  `query:"limit"`
	Offset int64 `query:"offset"`
}

type FriendsQueryResult struct {
	model.PageInfo
	Items []model.Profile
}

type FriendsHandler = application.Handler[FriendsQuery, *FriendsQueryResult]

type friendsHandler struct {
	db *sql.DB
}

func NewFriendsHandler(db *sql.DB) FriendsHandler {
	return &friendsHandler{db}
}

func (receiver *friendsHandler) Handle(ctx context.Context, arg FriendsQuery) (*FriendsQueryResult, error) {

	r := db.NewRepository(receiver.db)

	count,err := r.GetFriendsCountByUserId.Handle(ctx,&db.GetFriendsCountByUserIdQuery{UserId: arg.UserId})

	if err!=nil {
		return nil, err
	}

	profiles,err := r.GetFriendsPageByUserId.Handle(ctx,&db.GetFriendsPageByUserIdQuery{
		UserId: arg.UserId,
		Limit:  arg.Limit,
		Offset: arg.Offset,
		})

	if err!=nil {
		return nil, err
	}


	return &FriendsQueryResult{
		PageInfo: model.PageInfo{
			From:  int(arg.Offset),
			Count: len(profiles),
			Total: int(count[0]),
			},
			Items:    common.Map[db.Profile,model.Profile](profiles, model.NewProfileFromDb),
			},nil

}