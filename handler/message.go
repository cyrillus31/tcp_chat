package handler

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

// Marshal Message into []byte
func (m *Message) Marshal() ([]byte, error) {
	JSONdata, err := json.Marshal(m)
	if err != nil {
		return []byte{}, err
	}
	fullData := []byte{}
	fullData = append(fullData, VERSION)
	fullData = append(fullData, READ_JSON)

	lenght := len(JSONdata)
	lenghtData := make([]byte, 2)
	binary.BigEndian.PutUint16(lenghtData, uint16(lenght))

	fullData = append(fullData, lenghtData...)
	fullData = append(fullData, JSONdata...)
	return fullData, nil
}

// Unmarshal Message from []byte
func (m *Message) Unmarshal(data []byte) error {
	if data[0] != VERSION {
		return errors.New("Version is not correct")
	}
	if data[1] != READ_JSON {
		return errors.New("The command is not supported")
	}
	length := int(binary.BigEndian.Uint16(data[2:4]))
	end := HEADER_LENGTH + length
	if len(data) < end {
		return errors.New("Data is too short")
	}
	JSONdata := data[HEADER_LENGTH:HEADER_LENGTH + length]
	println("DEBUG :: ", "length:", length, "end:", end, "total_len:", len(data), "len_jsondata:", len(JSONdata))
	err := json.Unmarshal(JSONdata, m)
	if err != nil {
		return fmt.Errorf("Error during Unmarshal: %w", err)
	}
	return nil
}
