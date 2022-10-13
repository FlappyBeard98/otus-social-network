package query

import (
	"context"
	"database/sql"
	"social-network/common/application"
	"social-network/db"
	"social-network/model"
)

type ProfileQuery struct {
	UserId int64 `param:"userId"`
}

type ProfileHandler = application.Handler[ProfileQuery, *model.Profile]

type profileHandler struct {
	db *sql.DB
}

func NewProfileHandler(db *sql.DB) ProfileHandler {
	return &profileHandler{db}
}

func (receiver *profileHandler) Handle(ctx context.Context, arg ProfileQuery) (*model.Profile, error) {

	r := db.NewRepository(receiver.db)

	profile, err := r.GetProfileByUserId.Handle(ctx, &db.GetProfileByUserIdQuery{UserId: arg.UserId})

	if err != nil {
		return nil, err
	}

	result := model.NewProfileFromDb(profile[0])
	return &result, nil
}
