package repository

import (
	"context"
	pb "example/go-todolist/task-srv/proto/task"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// 默认数据库名
	DbName = "todolist"
	// 默认表名
	TaskCollection = "task"
	UnFinished     = 0
	Finished       = 1
)

// 定义数据库基本的增删改查操作
type TaskRepository interface {
	InsertOne(ctx context.Context, task *pb.Task) error
	Delete(ctx context.Context, id string) error
	Modify(ctx context.Context, task *pb.Task) error
	Finished(ctx context.Context, task *pb.Task) error
	Count(ctx context.Context, keyword string) (int64, error)
	Search(ctx context.Context, req *pb.SearchRequest) ([]*pb.Task, error)
	FindById(ctx context.Context, id string) (*pb.Task, error)
}

// 数据库操作实现类
type TaskRepositoryImpl struct {
	Conn *mongo.Client
}

// 定义默认的操作表
func (repo *TaskRepositoryImpl) collection() *mongo.Collection {
	return repo.Conn.Database(DbName).Collection(TaskCollection)
}

func (repo *TaskRepositoryImpl) InsertOne(ctx context.Context, task *pb.Task) error {
	_, err := repo.collection().InsertOne(ctx, bson.M{
		"body":       task.Body,
		"startTime":  task.StartTime,
		"endTime":    task.EndTime,
		"isFinished": UnFinished,
		"createTime": time.Now().Unix(),
		"userId":     task.UserId,
	})
	return err
}

func (repo *TaskRepositoryImpl) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.collection().DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (repo *TaskRepositoryImpl) Modify(ctx context.Context, task *pb.Task) error {
	id, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return err
	}
	_, err = repo.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"body":       task.Body,
		"startTime":  task.StartTime,
		"endTime":    task.EndTime,
		"updateTime": time.Now().Unix(),
	}})
	return err
}
func (repo *TaskRepositoryImpl) Finished(ctx context.Context, task *pb.Task) error {
	id, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return errors.WithMessage(err, "parse ID")
	}
	now := time.Now().Unix()
	update := bson.M{
		"isFinished": int32(task.IsFinished),
		"updateTime": now,
	}
	if task.IsFinished == Finished {
		update["finishTime"] = now
	}
	_, err = repo.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (repo *TaskRepositoryImpl) Count(ctx context.Context, keyword string) (int64, error) {
	filter := bson.M{}
	if keyword != "" && strings.TrimSpace(keyword) != "" {
		filter = bson.M{"body": bson.M{"$regex": keyword}}
	}
	count, err := repo.collection().CountDocuments(ctx, filter)
	return count, err
}

func (repo *TaskRepositoryImpl) Search(ctx context.Context, req *pb.SearchRequest) ([]*pb.Task, error) {
	filter := bson.M{}
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		filter = bson.M{"body": bson.M{"$regex": req.Keyword}}
	}

	cursor, err := repo.collection().Find(ctx,
		filter,
		options.Find().SetSkip((req.PageCode-1)*req.PageSize),
		options.Find().SetLimit(req.PageSize),
		options.Find().SetSort(bson.M{req.SortBy: req.Order}))
	if err != nil {
		return nil, errors.WithMessage(err, "search mongo")
	}
	var rows []*pb.Task
	if err := cursor.All(ctx, &rows); err != nil {
		return nil, errors.WithMessage(err, "parse data")
	}
	return rows, nil
}

// 通过ID查询task信息
func (repo *TaskRepositoryImpl) FindById(ctx context.Context, id string) (*pb.Task, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.WithMessage(err, "parse ID")
	}
	result := repo.collection().FindOne(ctx, bson.M{"_id": objectId})
	task := &pb.Task{}
	if err := result.Decode(task); err != nil {
		return nil, errors.WithMessage(err, "search mongo")
	}
	return task, nil
}
