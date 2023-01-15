package socket

import (
	"bufio"
	"net/textproto"
	"strings"
)

func (s *Socket) Input() (bool, string) {
	content := ""

	for s.Connected && content == "" {
		data, err := textproto.NewReader(bufio.NewReader(s.Socket)).ReadLine()

		if err != nil {
			s.Connected = false
			return false, content
		}

		content = strings.TrimSpace(string(data))
		//strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(string(data)), "\n", ""), " ", ""), "\r", "")
	}

	return true, content
}

func (s *Socket) Send(Payload string) bool {
	_, err := s.Socket.Write([]byte(Payload))

	if err != nil {
		s.Connected = false
		return false
	}

	return true
}
