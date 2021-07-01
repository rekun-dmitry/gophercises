package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
)

var (
	db *bolt.DB
)

func SetCredentials() {
	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("restapi"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.Put([]byte("joe"), []byte("$2a$12$aMfFQpGSiPiYkekov7LOsu63pZFaWzmlfm1T8lvG6JFj2Bh4SZPWS"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("mary"), []byte("$2a$12$l398tX477zeEBP6Se0mAv.ZLR8.LZZehuDgbtw2yoQeMjIyCNCsRW"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// verifyUserPass verifies that username/password is a valid pair matching
// our userPasswords "database".
func VerifyUserPass(username, password string) bool {
	var wantPass []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("restapi"))
		if b == nil {
			return fmt.Errorf("Bucket restapi not found!")
		}
		wantPass = b.Get([]byte(username))
		return nil
	})
	if err != nil || len(wantPass) == 0 {
		return false
	}
	if cmperr := bcrypt.CompareHashAndPassword(wantPass, []byte(password)); cmperr == nil {
		return true
	}
	return false
}
