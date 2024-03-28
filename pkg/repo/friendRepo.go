package repo

import (
	"database/sql"
	"hellowWorldDeploy/pkg/entity"
)

type FriendInterface interface {
	CreateFriendRequest(entity.FriendRequest) error
	AcceptFriendRequest(int64) error
	DeleteFriendRequest(int64) error
	CreateFriends(int64, int64) error
	GetFriendRequestByID(int64) (*entity.FriendRequest, error)
	GetFriendRequestsByRecipientID(int64) ([]entity.FriendRequest, error)
}

func (r *Repository) CreateFriendRequest(request entity.FriendRequest) error {
	stmt, err := r.db.Prepare(`INSERT INTO friend_requests (sender_id,recipient_id) VALUES ($1,$2) RETURNING id`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data CreateFriendRequest(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			r.log.Printf("\nError at the stage of closing stmt CreateFriendRequest(repo): %s\n", err.Error())
		}
	}(stmt)
	err = stmt.QueryRow(request.SenderID, request.RecipientID).Scan(&request.ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateFriendRequest(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) AcceptFriendRequest(requestID int64) error {
	stmt, err := r.db.Prepare(`UPDATE friend_requests SET is_accepted = true WHERE id = $1`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data AcceptFriendRequest(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			r.log.Printf("\nError at the stage of closing stmt AcceptFriendRequest(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(requestID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting AcceptFriendRequest(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) DeleteFriendRequest(requestID int64) error {
	stmt, err := r.db.Prepare(`DELETE FROM friend_requests WHERE id = $1`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data DeleteFriendRequest(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			r.log.Printf("\nError at the stage of closing stmt DeleteFriendRequest(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(requestID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Deleting DeleteFriendRequest(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) CreateFriends(user1ID, user2ID int64) error {
	stmt, err := r.db.Prepare(`INSERT INTO  friends (user_id,friend_id) VALUES ($1,$2), ($2,$1)`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data CreateFriends(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			r.log.Printf("\nError at the stage of closing stmt CreateFriends(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(user1ID, user2ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateFriends(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetFriendRequestByID(request int64) (*entity.FriendRequest, error) {
	result := &entity.FriendRequest{}
	err := r.db.QueryRow(`SELECT * FROM friend_requests WHERE id = $1`, request).Scan(&result.ID, &result.SenderID, &result.RecipientID, &result.IsAccepted)
	if err != nil {
		r.log.Printf("\nError at the stage of data Selecting GetFriendRequestByID(repo): %s\n", err.Error())
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetFriendRequestsByRecipientID(recipientID int64) ([]entity.FriendRequest, error) {
	rows, err := r.db.Query(`SELECT * FROM friend_requests WHERE recipient_id = $1 and is_accepted = false`, recipientID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Selecting GetFriendRequestsByRecipientID(repo): %s\n", err.Error())
		return nil, err
	}
	requestArray := make([]entity.FriendRequest, 0)
	for rows.Next() {
		result := entity.FriendRequest{}
		err = rows.Scan(&result.ID, &result.SenderID, &result.RecipientID, &result.IsAccepted)
		if err != nil {
			r.log.Printf("\nError at the stage of data scanning GetFriendRequestsByRecipientID(repo): %s\n", err.Error())
			return nil, err
		}
		requestArray = append(requestArray, result)
	}
	return requestArray, nil
}
