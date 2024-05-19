package db

import "errors"

var ErrDBRowUniqueConstraint = errors.New("entry already exists")
