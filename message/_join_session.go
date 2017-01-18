package message

import (
	"encoding/json"
	"log"

	"github.com/drawr-team/core-server/bolt"
)

const (
	JoinSessionMessageType       = "join-session"
	JoinSessionAckType           = "ack-session"
	JoinSessionDataStatusSuccess = "join-session-success"
	JoinSessionDataStatusFailure = "join-session-failure"
)

// JoinSessionData is the data used to initialize a new Session
type JoinSessionData struct {
	Username  string `json:"username"`
	SessionID string `json:"sessionId"`
}

// JoinSessionAckData is the data returned in the server ack
type JoinSessionAckData struct {
}

// HandleJoinSession handles a `join-session` type
func HandleJoinSession(m GenericMessage, p Provider, db bolt.DBClient) error {
	var data JoinSessionData
	if err := json.Unmarshal(m.Data, &data); err != nil {
		return err
	}

	// TODO: database code here
	// untested !!!
	sessionData, err := db.Get(bolt.SessionBucket, data.SessionID)
	if err != nil {
		// TODO: not like this, please!
		if err == bolt.ErrNotFound {
			log.Println("[db] session not found")
		}
	}

	log.Println(sessionData)

	// create repsonse
	resp, err := CreateMessage(JoinSessionAckType, GenericAck{
		Status: JoinSessionDataStatusSuccess,
		Data:   nil,
	})
	if err != nil {
		return err
	}

	message, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	p.AbsorbTo(data.SessionID, message)

	return nil
}