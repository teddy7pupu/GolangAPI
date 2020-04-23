package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"net/http"

	"api/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

// BugDetail bug資料結構
type BugDetail struct {
	Time     *string `json:"time" bson:"time"`         // 新增時間
	Title    *string `json:"title" bson:"title"`       // 標題
	SubTitle *string `json:"subTitle" bson:"subTitle"` // 副標題
	Status   *int    `json:"status" bson:"status"`     // 狀態 0: 未處理, 1: 已處理
	ID       string  `json:"id" bson:"_id"`            // ID
}

// APIResponse api回傳模型
type APIResponse struct {
	SysCode int         `json:"sysCode"`
	SysMsg  string      `json:"sysMsg"`
	Data    interface{} `json:"data"`
}

func init() {
	ConnetDB()
}

// ConnetDB 連結資料庫
func ConnetDB() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database = client.Database("bugDB")

	// fmt.Printf("database 型別是 %T", database)
}

// AddBug 新增 bug
func AddBug(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}

	var bugDetail BugDetail
	json.Unmarshal(body, &bugDetail)
	bugDetail.ID = primitive.NewObjectID().Hex()

	if bugDetail.Time == nil || bugDetail.Title == nil ||
		bugDetail.SubTitle == nil || bugDetail.Status == nil {
		response := APIResponse{200, "缺少必填欄位", nil}
		services.ResponseWithJSONgo(w, http.StatusOK, response)
		return
	}

	collection := database.Collection("bug")
	collection.InsertOne(context.TODO(), bugDetail)

	response := APIResponse{200, "新增成功", bugDetail}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}

// GetBugList 取得 bug 清單
func GetBugList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		fmt.Println(err)
	}

	var bugDetail BugDetail
	json.Unmarshal(body, &bugDetail)

	filters := bson.D{}
	if bugDetail.Title != nil {
		filter := primitive.E{Key: "title", Value: bugDetail.Title}
		filters = append(filters, filter)
	}
	if bugDetail.Time != nil {
		filter := primitive.E{Key: "time", Value: bugDetail.Time}
		filters = append(filters, filter)
	}
	if bugDetail.Status != nil {
		filter := primitive.E{Key: "status", Value: bugDetail.Status}
		filters = append(filters, filter)
	}
	if bugDetail.SubTitle != nil {
		filter := primitive.E{Key: "subTitle", Value: bugDetail.SubTitle}
		filters = append(filters, filter)
	}
	if bugDetail.ID != "" {
		filter := primitive.E{Key: "_id", Value: bugDetail.ID}
		filters = append(filters, filter)
	}

	collection := database.Collection("bug")
	cur, err := collection.Find(context.TODO(), filters)
	if err != nil {
		log.Fatal(err)
	}

	results := []*BugDetail{}
	for cur.Next(context.TODO()) {
		var elem BugDetail
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	response := APIResponse{200, "", results}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}

// UpdateBug 更新 bug
func UpdateBug(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		fmt.Println(err)
	}

	var bugDetail BugDetail
	json.Unmarshal(body, &bugDetail)
	if bugDetail.ID == "" {
		response := APIResponse{200, "請輸入查詢id", nil}
		services.ResponseWithJSONgo(w, http.StatusOK, response)
		return
	}

	collection := database.Collection("bug")
	filter := bson.D{primitive.E{Key: "_id", Value: bugDetail.ID}}
	updateItem := bson.D{}
	opts := options.Update().SetUpsert(true)

	var results *BugDetail
	collection.FindOne(context.TODO(), filter).Decode(&results)
	if results == nil {
		response := APIResponse{200, "查無符合id的資料", nil}
		services.ResponseWithJSONgo(w, http.StatusOK, response)
		return
	}

	if bugDetail.Title != nil {
		update := primitive.E{Key: "title", Value: bugDetail.Title}
		updateItem = append(updateItem, update)
	}
	if bugDetail.SubTitle != nil {
		update := primitive.E{Key: "subTitle", Value: bugDetail.SubTitle}
		updateItem = append(updateItem, update)
	}
	if bugDetail.Status != nil {
		update := primitive.E{Key: "status", Value: bugDetail.Status}
		updateItem = append(updateItem, update)
	}
	if bugDetail.Time != nil {
		update := primitive.E{Key: "time", Value: bugDetail.Time}
		updateItem = append(updateItem, update)
	}
	updateSet := bson.D{primitive.E{Key: "$set", Value: updateItem}}
	collection.UpdateOne(context.TODO(), filter, updateSet, opts)
	collection.FindOne(context.TODO(), filter).Decode(&results)

	response := APIResponse{200, "資料更新成功", results}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}

// DeleteBug 刪除 bug
func DeleteBug(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		fmt.Println(err)
	}

	var bugDetail BugDetail
	json.Unmarshal(body, &bugDetail)
	if bugDetail.ID == "" {
		response := APIResponse{200, "請輸入查詢id", nil}
		services.ResponseWithJSONgo(w, http.StatusOK, response)
		return
	}

	collection := database.Collection("bug")
	filter := bson.D{primitive.E{Key: "_id", Value: bugDetail.ID}}

	var results *BugDetail
	collection.FindOne(context.TODO(), filter).Decode(&results)
	if results == nil {
		response := APIResponse{200, "查無符合id的資料", nil}
		services.ResponseWithJSONgo(w, http.StatusOK, response)
		return
	}

	collection.DeleteOne(context.TODO(), filter)

	response := APIResponse{200, "資料刪除成功", results}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}
