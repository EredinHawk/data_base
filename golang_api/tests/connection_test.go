package db

import (
	"database/sql"
	"golang_api/db"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Connection()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
