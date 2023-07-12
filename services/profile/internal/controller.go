package service

import (
	"context"
	"errors"
	"net/http"
	"social-network/lib/pg"
	"social-network/services/profile/internal/types"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags public
// @Accept  json
// @Produce  json
// @Param body body types.RegisterRequest true "Register request"
// @Success 200 {object} types.Profile
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (o *App) Register(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.RegisterRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth, err := r.NewAuth()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profile, err := r.NewProfile()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fn := func(ctx context.Context, tx pg.Db) error {

		if err := auth.InsertAuth().QueryOne(ctx, tx, &profile.UserId); err != nil {
			return err
		}
		if rowsAffected, err := profile.UpsertProfile().Exec(ctx, tx); err != nil {
			return err
		} else if rowsAffected == 0 {
			return errors.New("profile was not saved")
		}

		return nil
	}

	err = pg.BeginTxFunc(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted}, o.Db, fn)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

// Profiles godoc
// @Summary      Show profiles
// @Description  get page of profiles
// @Tags         public
// @Accept       json
// @Produce      json
// @Param        offset query      int  false  "Page offset"
// @Param        limit  query      int  false  "Page limit"
// @Param        firstName  query      string  true  "First name"
// @Param        lastName  query      string  true  "Last name"
// @Param        age  query      int  true  "Age"
// @Param        gender query      int  true "Gender"
// @Param        city query      string  true "City"
// @Param        hobbies query      string  true "Hobbies"
// @Success      200  {object}  types.PageResponse[types.Profile]
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /profiles [get]
func (o *App) Profiles(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.ProfilesRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var profiles = make([]types.Profile, 0)

	err := r.ReadProfilesPage().Query(ctx, o.Db, &profiles)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var total int64

	err = r.ReadProfilesTotal().QueryOne(ctx, o.Db, &total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := types.NewPageResponse(&r.PageRequest, profiles, total)

	return c.JSON(http.StatusOK, result)
}

// Profile godoc
// @Summary      Show profile
// @Description  get profile by user id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userId query      int  true  "User id"
// @Success      200  {object}  types.Profile
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /profile [get]
// @Security     BasicAuth
func (o *App) Profile(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.UserIdRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := new(types.Profile)

	err := r.ReadProfileByUserId().QueryOne(ctx, o.Db, result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// SaveProfile godoc
// @Summary      Save profile
// @Description  save profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body body types.Profile true "Profile"
// @Success      200  {object}  types.Profile
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /profile [post]
// @Security     BasicAuth
func (o *App) SaveProfile(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Profile)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected, err := r.UpsertProfile().Exec(ctx, o.Db); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else if rowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, "profile was not saved")
	}

	return c.JSON(http.StatusOK, r)
}

// Friends godoc
// @Summary      Show friends
// @Description  get page of friends
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userId query      int  true  "User id"
// @Param        offset query      int  false  "Page offset"
// @Param        limit  query      int  false  "Page limit"
// @Success      200  {object}  types.PageResponse[types.Profile]
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /friends [get]
// @Security     BasicAuth
func (o *App) Friends(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.UserIdRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pr := new(types.PageRequest)

	if err := c.Bind(pr); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var profiles = make([]types.Profile, 0)

	err := r.ReadUserFriendsProfiles(pr).Query(ctx, o.Db, &profiles)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var total int64

	err = r.ReadUserFriendsTotal().QueryOne(ctx, o.Db, &total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := types.NewPageResponse(pr, profiles, total)

	return c.JSON(http.StatusOK, result)
}

// AddFriend godoc
// @Summary      Add friend
// @Description  add friend by user id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userId query      int  true  "User id"
// @Param        friendId query      int  true  "Friend id"
// @Success      200  {object}  types.Friend
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /friend [post]
// @Security     BasicAuth
func (o *App) AddFriend(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Friend)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected, err := r.InsertFriend().Exec(ctx, o.Db); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else if rowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, "friend was not added")
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteFriend godoc
// @Summary      Delete friend
// @Description  delete friend by user id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userId query      int  true  "User id"
// @Param        friendId query      int  true  "Friend id"
// @Success      200  {object}  types.Friend
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /friend [delete]
// @Security     BasicAuth
func (o *App) DeleteFriend(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Friend)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := r.DeleteFriend().Exec(ctx, o.Db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
