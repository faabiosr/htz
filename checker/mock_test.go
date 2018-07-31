package checker

import (
	"context"
	"database/sql/driver"
	"errors"
)

type (
	MockDriver     struct{}
	MockDriverConn struct{}
)

func (md *MockDriver) Open(dns string) (driver.Conn, error) {
	return new(MockDriverConn), nil
}

func (mdc *MockDriverConn) Begin() (driver.Tx, error) {
	return nil, nil
}

func (mdc *MockDriverConn) Close() error {
	return nil
}

func (mdc *MockDriverConn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (mdc *MockDriverConn) Ping(ctx context.Context) error {
	return errors.New("Connection error!")
}
