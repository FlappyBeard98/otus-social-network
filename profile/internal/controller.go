package service

import (
	"context"
	"database/sql"
	"net/http"
	"social-network/lib/mysql"
	"social-network/profile/internal/types"

	"github.com/labstack/echo/v4"
)

func (o *App) Register(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.RegisterRequest)

	if err := c.Bind(r); err != nil {
    	return c.String(http.StatusBadRequest, err.Error())
  	}

	auth,err := types.NewAuth(r.Auth.Login,r.Auth.Password)

	if err != nil {
    	return c.String(http.StatusBadRequest, err.Error())
  	}

	profile,err := types.NewProfile()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	


	fn := func(ctx context.Context,tx *sql.Tx) error {

		auth.InsertAuth().Exec(ctx,tx)
		return nil
	}

	err = mysql.BeginTxFunc(ctx,nil,o.Db,fn)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}



	password, err := common.Encrypt([]byte(receiver.key), []byte(arg.Password))

	_, err = r.AddAuth.Handle(ctx, &db.AddAuthQuery{
		Login:    arg.Login,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	auth, err := r.GetAuthByLogin.Handle(ctx, &db.GetAuthByLoginQuery{
		Login: arg.Login,
	})

	if err != nil {
		return nil, err
	}

	_, err = r.SaveProfile.Handle(ctx, &db.SaveProfileQuery{
		UserId:    auth[0].UserId,
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
	})

	if err != nil {
		return nil, err
	}

	return &RegisterCommandResult{UserId: auth[0].UserId}, nil
	return c.String(http.StatusOK, "")
}

func (o *App) Profiles(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	count, err := r.GetProfilesCountByFilter.Handle(ctx, &db.GetProfilesCountByFilterQuery{
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
	})

	if err != nil {
		return nil, err
	}

	profiles, err := r.GetProfilesPageByFilter.Handle(ctx, &db.GetProfilesPageByFilterQuery{
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
		Limit:     arg.Limit,
		Offset:    arg.Offset,
	})

	if err != nil {
		return nil, err
	}

	return &ProfilesByFilterQueryResult{
		PageInfo: model.PageInfo{
			From:  int(arg.Offset),
			Count: len(profiles),
			Total: int(count[0]),
		},
		Items: common.Map[db.Profile, model.Profile](profiles, model.NewProfileFromDb),
	}, nil
	return c.String(http.StatusOK, "")
}

func (o *App) Profile(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	profile, err := r.GetProfileByUserId.Handle(ctx, &db.GetProfileByUserIdQuery{UserId: arg.UserId})

	if err != nil {
		return nil, err
	}

	result := model.NewProfileFromDb(profile[0])
	return &result, nil
	return c.String(http.StatusOK, "")
}

func (o *App) SaveProfile(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	_, err := r.SaveProfile.Handle(ctx, &db.SaveProfileQuery{
		UserId:    arg.UserId,
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
	return c.String(http.StatusOK, "")
}

func (o *App) Friends(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	count, err := r.GetFriendsCountByUserId.Handle(ctx, &db.GetFriendsCountByUserIdQuery{UserId: arg.UserId})

	if err != nil {
		return nil, err
	}

	profiles, err := r.GetFriendsPageByUserId.Handle(ctx, &db.GetFriendsPageByUserIdQuery{
		UserId: arg.UserId,
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})

	if err != nil {
		return nil, err
	}

	return &FriendsQueryResult{
		PageInfo: model.PageInfo{
			From:  int(arg.Offset),
			Count: len(profiles),
			Total: int(count[0]),
		},
		Items: common.Map[db.Profile, model.Profile](profiles, model.NewProfileFromDb),
	}, nil
	return c.String(http.StatusOK, "")
}

func (o *App) AddFriend(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	_, err := r.AddFriend.Handle(ctx, &db.AddFriendQuery{
		UserId:       arg.UserId,
		FriendUserId: arg.FriendUserId,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil

	return c.String(http.StatusOK, "")
}

func (o *App) DeleteFriend(c echo.Context) error {
	r := db.NewRepository(receiver.db)

	_, err := r.RemoveFriend.Handle(ctx, &db.RemoveFriendQuery{
		UserId:       arg.UserId,
		FriendUserId: arg.FriendUserId,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
	return c.String(http.StatusOK, "")
}
