package db

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
)

type User struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type DB struct {
	db *bolt.DB
}

var DBConnection DB

func NewDB(path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Create(user *User) error {
	err := d.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))

		// Check if user with the same ID already exists
		if bucket.Get([]byte(user.ID)) != nil {
			return errors.New("user with ID already exists")
		}

		// Encode user as JSON and store in bucket
		userBytes, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(user.ID), userBytes)
	})
	return err
}

func (d *DB) Edit(id string, username string, password string) error {
	err := d.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))

		// Check if user with the ID exists
		userBytes := bucket.Get([]byte(id))
		if userBytes == nil {
			return errors.New("user with ID not found")
		}

		// Decode user JSON
		user := &User{}
		err := json.Unmarshal(userBytes, user)
		if err != nil {
			return err
		}

		// Update user fields and store in bucket
		user.Username = username
		user.Password = password
		newUserBytes, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(id), newUserBytes)
	})
	return err
}

func (d *DB) Delete(id string) error {
	err := d.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))

		// Check if user with the ID exists
		if bucket.Get([]byte(id)) == nil {
			return errors.New("user with ID not found")
		}

		// Delete user from bucket
		return bucket.Delete([]byte(id))
	})
	return err
}

func (d *DB) List() ([]User, error) {
	users := []User{}
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))

		// Iterate over all users in bucket and decode JSON
		return bucket.ForEach(func(k, v []byte) error {
			user := &User{}
			err := json.Unmarshal(v, user)
			if err != nil {
				return err
			}
			users = append(users, *user)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DB) Get(id string) (*User, error) {
	var user *User
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))

		// Get user by ID
		userBytes := bucket.Get([]byte(id))
		if userBytes == nil {
			return errors.New("user with ID not found")
		}

		// Decode user JSON
		user = &User{}
		err := json.Unmarshal(userBytes, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
