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

func (m *Message) Marshall() ([]byte, error) {
	JSONdata, err := json.Marshal(m)
	if err != nil {
		return []byte{}, err
	}
	fullData := []byte{}
	fullData = append(fullData, VERSION)
	fullData = append(fullData, READ_JSON)

	lenght := len(JSONdata)
	lenghtData := make([]byte, 2)
	binary.BigEndian.AppendUint16(lenghtData, uint16(lenght))

	fullData = append(fullData, lenghtData...)
	fullData = append(fullData, JSONdata...)
	return fullData, nil
}

func (m *Message) Unmarshall(data []byte) error {
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
	// JSONdata := data[HEADER_LENGTH:HEADER_LENGTH + length]
	JSONdata := data[HEADER_LENGTH:]
	err := json.Unmarshal(JSONdata, m)
	if err != nil {
		return fmt.Errorf("Error during Unmarshal: %w", err)
	}
	return nil
}
