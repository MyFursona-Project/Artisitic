package handler

import (
	"artistically/api/database"
	"artistically/api/model"
	"artistically/api/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()
var arts = new(model.Art)

func Index(c *fiber.Ctx) error {
	return utils.ResponseMsg(c, 200, "Api is running", nil)
}
func PostArt(c *fiber.Ctx) error {
	arts = new(model.Art)
	var ctx = context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	arts.CreatedAt = time.Now()
	if err := c.BodyParser(arts); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		if r, err := db.Collection("art").InsertOne(ctx, arts); err != nil {
			return utils.ResponseMsg(c, 500, "Inserted data unsuccesfully", err.Error())
		} else {
			return utils.ResponseMsg(c, 200, "Inserted data succesfully", r)
		}
	}
}
func GetArtByID(c *fiber.Ctx) error {
	_id := c.Params("id")
	var ctx = context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	if docID, err := primitive.ObjectIDFromHex(_id); err != nil {
		return utils.ResponseMsg(c, 400, "Get Data unsuccesfully", err.Error())
	} else {
		q := bson.M{"_id": docID}
		art := model.Art{}
		result := db.Collection("art").FindOne(ctx, q)
		result.Decode(&art)
		if result.Err() != nil {
			return utils.ResponseMsg(c, 200, "Get Data unsuccesfully", result.Err().Error())
		} else {
			return utils.ResponseMsg(c, 200, "Get Data Succesfully", art)
		}
	}
}

func GetUserArtsById(c *fiber.Ctx) error {
	var ctx = context.Background()
	_id := c.Params("id")
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("art").Find(ctx, bson.M{"userid": _id})
	if err != nil {
		return utils.ResponseMsg(c, 200, "Get Data unsuccesfully", csr.Err().Error())
	} else {
		return utils.ResponseMsg(c, 200, "Get Data Succesfully", csr)
	}
}
