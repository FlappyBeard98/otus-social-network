package query

import (
	"context"
	"database/sql"
	"social-network/common"
	"social-network/common/application"
	"social-network/db"
	"social-network/model"
)

type ProfilesByFilterQuery struct {
	FirstName *string `query:"firstName"`
	LastName  *string `query:"lastName"`
	Age       *int32 `query:"age"`
	Gender    *int32 `query:"gender"`
	City      *string `query:"city"`
	Hobbies   *string `query:"hobbies"`
	Limit int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type ProfilesByFilterQueryResult struct {
	model.PageInfo
	Items []model.Profile
}

type ProfilesByFilterHandler = application.Handler[ProfilesByFilterQuery, *ProfilesByFilterQueryResult]

type profilesByFilterHandler struct {
	db *sql.DB
}

func NewProfilesByFilterHandler(db *sql.DB) ProfilesByFilterHandler {
	return &profilesByFilterHandler{db}
}

func (receiver *profilesByFilterHandler) Handle(ctx context.Context, arg ProfilesByFilterQuery) (*ProfilesByFilterQueryResult, error) {
	r := db.NewRepository(receiver.db)

	count,err := r.GetProfilesCountByFilter.Handle(ctx,&db.GetProfilesCountByFilterQuery{
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
	})


	if err!=nil {
		return nil, err
	}

	profiles,err := r.GetProfilesPageByFilter.Handle(ctx,&db.GetProfilesPageByFilterQuery{
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
		Limit:  arg.Limit,
		Offset: arg.Offset,
		})

	if err!=nil {
		return nil, err
	}


	return &ProfilesByFilterQueryResult{
		PageInfo: model.PageInfo{
			From:  int(arg.Offset),
			Count: len(profiles),
			Total: int(count[0]),
			},
			Items:    common.Map[db.Profile,model.Profile](profiles, model.NewProfileFromDb),
			},nil
}
