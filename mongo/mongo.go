package mongo

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "github.com/hdiomede/status-health/domain"
    "github.com/hdiomede/status-health/domain/repository"
)

type requestRepository struct {
	db string
	session *mgo.Session
}

func (repository *requestRepository) Store(r *domain.Request) error {
	r.Id = bson.NewObjectId()
	r.Status = "unknown"

	if r.Environment == "" {
		r.Environment = "DEFAULT"
	}

	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	sess := repository.session.Copy()
	defer sess.Close()

	c := sess.DB(repository.db).C("request")

	err := c.Insert(r)

	return err
}

func NewRequestRepository(db string, session *mgo.Session) (repository.RequestRepository, error) {
	r := &requestRepository{
		db: db,
		session: session,
	}

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:    true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("request")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}

	return r, nil
}
