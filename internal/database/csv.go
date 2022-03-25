package database

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/dalot/go-skeleton-mid/pkg/resources"
)

func (s *Store) LoadMessages() error {
	f, err := os.Open("internal/database/messages.csv")
	if err != nil {
		return fmt.Errorf("could not open file: %s", err)
	}
	defer f.Close()

	// Skip first row (which contains the columns)
	row1, err := bufio.NewReader(f).ReadSlice('\n')
	if err != nil {
		return err
	}
	_, err = f.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return err
	}

	reader := csv.NewReader(f)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return fmt.Errorf("could not read the current row: %s", err)
		}
		message, error := fromRow(row)
		if error != nil {
			continue
		}

		s.SetMessage(message)

	}

	return nil
}

func fromRow(row []string) (*resources.Message, error) {
	t, err := time.Parse(time.RFC3339, row[4])
	if err != nil {
		fmt.Println("failed with ", row[0], row[1], row[2], row[3], row[4])
		return nil, fmt.Errorf("could not parse time: %s", err)
	}
	return &resources.Message{
		ID:        row[0],
		Name:      row[1],
		Email:     row[2],
		Text:      row[3],
		CreatedAt: t,
	}, nil
}
