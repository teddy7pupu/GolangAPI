package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"log"
	"encoding/base64"

	"net/http"

	"api/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/aws/aws-sdk-go/aws"
)

// BugDetail bug資料結構
type BugDetail struct {
	Time     *string 	`json:"time" bson:"time"`        	// 新增時間
	Title    *string 	`json:"title" bson:"title"`      	// 標題
	SubTitle *string 	`json:"subTitle" bson:"subTitle"` 	// 副標題
	Image 	 *string 	`json:"image" bson:"image"` 		// 圖片
	Status   *int    	`json:"status" bson:"status"`     	// 狀態 0: 未處理, 1: 已處理
	ID       string  	`json:"id" bson:"_id"`            	// ID
}

// APIResponse api回傳模型
type APIResponse struct {
	SysCode int         `json:"sysCode"`
	SysMsg  string      `json:"sysMsg"`
	Data    interface{} `json:"data"`
}

// 相關變數
const domainURL = "http://127.0.0.1:3000/api/readImage?image="
const databaseURL = "mongodb://:27017"
const databaseName = "bugDB"
const collectionName = "bug"
const imageList = "./uploaded/"

var database *mongo.Database
var collection *mongo.Collection

func init() {
	ConnetDB()
}

// ConnetDB 連結資料庫
func ConnetDB() {
	clientOptions := options.Client().ApplyURI(databaseURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database = client.Database(databaseName)
	collection = database.Collection(collectionName)

	// fmt.Printf("collection 型別是 %T", collection)
}

// AddBug 新增 bug
func AddBug(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
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

	imageStr := aws.StringValue(bugDetail.Image)
	dist, err := base64.StdEncoding.DecodeString(imageStr)
	if err == nil && bugDetail.Image != nil {
		fileNameStr := bugDetail.ID + ".png"
		fileAddress := imageList + fileNameStr
		fileURL := domainURL + fileNameStr
		f, _ := os.OpenFile(fileAddress, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(dist)
		bugDetail.Image = aws.String(fileURL)
	} else {
		bugDetail.Image = nil
	}

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

	body, err := ioutil.ReadAll(r.Body)
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

	imageStr := aws.StringValue(bugDetail.Image)
	fileNameStr := bugDetail.ID + ".png"
	fileAddress := imageList + fileNameStr
	if bugDetail.Image != nil && bugDetail.Image != aws.String("rm") {
		// 修改圖片
		dist, err := base64.StdEncoding.DecodeString(imageStr)
		if err == nil {
			fileURL := domainURL + fileNameStr
			f, _ := os.OpenFile(fileAddress, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(dist)
			
			update := primitive.E{Key: "image", Value: fileURL}
			updateItem = append(updateItem, update)
		} else {
			os.Remove(fileAddress)
			update := primitive.E{Key: "image", Value: nil}
			updateItem = append(updateItem, update)
		}
	} else {
		// 移除圖片
		update := primitive.E{Key: "image", Value: nil}
		updateItem = append(updateItem, update)
		os.Remove(fileAddress)
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

// ReadImage 取得圖片
func ReadImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	keys := r.URL.Query()
	imageName := keys.Get("image")
	imageAddress := imageList + imageName
	http.ServeFile(w, r, imageAddress)
}

// 统一错误输出接口
func errorHandle(errStr string, w http.ResponseWriter) {
    response := APIResponse{200, errStr, nil}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}