package repository

import (
	"TheBoys/app/model/request"
	"TheBoys/app/model/response"
	"TheBoys/domain"

	"gorm.io/gorm"
)

func NewUserRepository(Db *gorm.DB) domain.UserRepository {
	return &userRepository{Db}
}

type userRepository struct {
	Db *gorm.DB
}

func (r *userRepository) FindUserByEmail(email string) (*response.LoginResponse, error) {
	var data *response.LoginResponse

	err := r.Db.Raw(`WITH "LatestRequests" AS(
       SELECT "userId", Max("createdAt") AS "latestCreatedAt"
	   FROM "UserLoginRequest" 
	   GROUP BY "userId" )
	   SELECT 
	    "User"."id" AS "id",
        "User"."username" AS "name",
        "User"."email" AS "email",
        "User"."mobile" AS "mobile",
        "User"."roleId" AS "roleId",
        "User"."token" AS "webToken",
        "User"."isActive" AS "isActive",
        "Roles"."name" AS "roleName",
        "UserLoginRequest"."createdAt" AS "otpCreatedAt"
	    FROM "User"
        INNER JOIN "Roles" ON "Roles"."id" = "User"."roleId"
        LEFT JOIN "UserLoginRequest" ON "UserLoginRequest"."userId" = "User"."id"
		LEFT JOIN "LatestRequests" ON "UserLoginRequest"."userId" = "LatestRequests"."userId" AND "UserLoginRequest"."createdAt" = "LatestRequests"."latestCreatedAt"
	    WHERE "User"."email" = ? AND "User"."isActive" = ?;`, email, true).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) CreateOtp(otp string, userId int16, email string, mobile string) error {
	err := r.Db.Exec(`INSERT INTO "UserLoginRequest"("userId","email","mobile","otp","createdAt","updatedAt") VALUES(?,?,?,?,NOW(),NOW())`, userId, email, mobile, otp).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindUserByEmailWithOtp(email string, otp string) (bool, error) {
	var count int
	err := r.Db.Raw(`SELECT COUNT(*) FROM "UserLoginRequest" WHERE "email"=? AND "otp" = ? AND "isUsed" = ?`, email, otp, false).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *userRepository) UpdateWebToken(token string, userId uint) error {

	err := r.Db.Exec(`UPDATE "User" SET "token" = ? WHERE "id" = ?`, token, userId).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateLoginOtpStatus(userId uint, otp string) error {
	err := r.Db.Exec(`UPDATE "UserLoginRequest" SET "isUsed" = ? WHERE "userId"=? AND "otp" = ?`, true, userId, otp).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindUserById(userId uint) (*response.LoginResponse, error) {
	var data *response.LoginResponse
	err := r.Db.Raw(`SELECT 
	    "User"."id" AS "id",
        "User"."name" AS "name",
        "User"."email" AS "email",
        "User"."mobile" AS "mobile",
        "User"."roleId" AS "roleId",
        "User"."token" AS "webToken",
        "User"."isActive" AS "isActive",
        "Roles"."name" AS "roleName",
	    FROM "User" 
	    INNER JOIN "Roles" ON "Roles"."id" = "User"."roleId" WHERE "id" = ? `, userId).Scan(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *userRepository) CreateUser(req request.UserCreationRequest) error {
	return r.Db.Transaction(func(tx *gorm.DB) error {

		var userId int64

		user := tx.Raw(`INSERT INTO "User"("name","email","mobile","createdAt","updatedAt") VALUES(?,?,?,NOW(),NOW()) RETURNING "id"`, req.Name, req.Email, req.Mobile).Scan(&userId)

		if user.Error != nil {
			return user.Error
		}

		var addressId int16

		address := tx.Exec(`INSERT INTO "Address"("doorNumber","streetName","pinCode","stateId","createdAt","updatedAt") VALUES(?,?,?,?,NOW(),NOW())RETURNING "id`,
			req.DoorNumber, req.Street, req.Pincode, req.StateId).Scan(&addressId)

		if address.Error != nil {
			return address.Error
		}

		err := tx.Exec(`INSERT INTO UserAddress(addressId,userId,createdAt,updatedAt) VALUES(?,?,NOW(),NOW())`, addressId, userId).Error

		if err != nil {
			return err
		}

		return nil
	})
}
