package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// MockDB untuk testing category
type MockDBCategory struct {
	QueryFunc    func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRowFunc func(ctx context.Context, sql string, args ...interface{}) pgx.Row
	ExecFunc     func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func (m *MockDBCategory) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(ctx, sql, args...)
	}
	return nil, nil
}

func (m *MockDBCategory) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(ctx, sql, args...)
	}
	return nil
}

func (m *MockDBCategory) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(ctx, sql, arguments...)
	}
	return pgconn.CommandTag{}, nil
}

func (m *MockDBCategory) Ping(ctx context.Context) error {
	return nil
}

func (m *MockDBCategory) Close(ctx context.Context) error {
	return nil
}

// Test GetCategory - Success
func TestGetCategory_Success(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
			return &MockRows{
				data: [][]interface{}{
					{1, "Electronics", "Electronic items"},
					{2, "Furniture", "Office furniture"},
				},
			}, nil
		},
	}

	repo := NewrepoCategory(mockDB)
	categories, err := repo.GetCategory()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(categories))
	}

	if categories[0].Name != "Electronics" {
		t.Errorf("Expected first category name to be 'Electronics', got '%s'", categories[0].Name)
	}
}

// Test GetCategory - Error
func TestGetCategory_Error(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
			return nil, errors.New("database error")
		},
	}

	repo := NewrepoCategory(mockDB)
	_, err := repo.GetCategory()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test AddCategory - Success
func TestAddCategory_Success(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{data: []interface{}{0}, err: nil}
		},
		ExecFunc: func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.AddCategory("New Category", "New Description")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test AddCategory - Duplicate Name
func TestAddCategory_DuplicateName(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{data: []interface{}{1}, err: nil}
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.AddCategory("Existing Category", "Description")

	if err == nil {
		t.Error("Expected an error for duplicate category name, got nil")
	}
}

// Test GetCategoryById - Success
func TestGetCategoryById_Success(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{data: []interface{}{1, "Electronics", "Electronic items"}, err: nil}
		},
	}

	repo := NewrepoCategory(mockDB)
	category, err := repo.GetCategoryById(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if category.Name != "Electronics" {
		t.Errorf("Expected category name to be 'Electronics', got '%s'", category.Name)
	}
}

// Test GetCategoryById - Not Found
func TestGetCategoryById_NotFound(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{err: errors.New("no rows")}
		},
	}

	repo := NewrepoCategory(mockDB)
	_, err := repo.GetCategoryById(999)

	if err == nil {
		t.Error("Expected an error for non-existent category, got nil")
	}
}

// Test UpdateCategory - Success
func TestUpdateCategory_Success(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{data: []interface{}{0}, err: nil}
		},
		ExecFunc: func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.UpdateCategory(1, "Updated Category", "Updated Description")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test UpdateCategory - Duplicate Name
func TestUpdateCategory_DuplicateName(t *testing.T) {
	mockDB := &MockDBCategory{
		QueryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return &MockRow{data: []interface{}{1}, err: nil}
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.UpdateCategory(1, "Existing Category", "Description")

	if err == nil {
		t.Error("Expected an error for duplicate category name, got nil")
	}
}

// Test DeleteCategory - Success
func TestDeleteCategory_Success(t *testing.T) {
	mockDB := &MockDBCategory{
		ExecFunc: func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, nil
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.DeleteCategory(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test DeleteCategory - Error
func TestDeleteCategory_Error(t *testing.T) {
	mockDB := &MockDBCategory{
		ExecFunc: func(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.CommandTag{}, errors.New("delete error")
		},
	}

	repo := NewrepoCategory(mockDB)
	err := repo.DeleteCategory(1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
