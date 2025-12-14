package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Mock Row untuk QueryRow
type MockRow struct {
	data []interface{}
	err  error
}

func (m *MockRow) Scan(dest ...interface{}) error {
	if m.err != nil {
		return m.err
	}
	for i, v := range m.data {
		if i < len(dest) {
			switch d := dest[i].(type) {
			case *int:
				if val, ok := v.(int); ok {
					*d = val
				}
			case *string:
				if val, ok := v.(string); ok {
					*d = val
				}
			case *float64:
				if val, ok := v.(float64); ok {
					*d = val
				}
			case *time.Time:
				if val, ok := v.(time.Time); ok {
					*d = val
				}
			}
		}
	}
	return nil
}

// Mock Rows untuk Query
type MockRows struct {
	data     [][]interface{}
	position int
	err      error
}

func (m *MockRows) Next() bool {
	if m.position < len(m.data) {
		m.position++
		return true
	}
	return false
}

func (m *MockRows) Scan(dest ...interface{}) error {
	if m.err != nil {
		return m.err
	}
	if m.position == 0 || m.position > len(m.data) {
		return errors.New("no data")
	}
	row := m.data[m.position-1]
	for i, v := range row {
		if i < len(dest) {
			switch d := dest[i].(type) {
			case *int:
				if val, ok := v.(int); ok {
					*d = val
				}
			case *string:
				if val, ok := v.(string); ok {
					*d = val
				}
			case *float64:
				if val, ok := v.(float64); ok {
					*d = val
				}
			case *time.Time:
				if val, ok := v.(time.Time); ok {
					*d = val
				}
			}
		}
	}
	return nil
}

func (m *MockRows) Close() {}

func (m *MockRows) Err() error {
	return m.err
}

func (m *MockRows) CommandTag() pgconn.CommandTag {
	return pgconn.CommandTag{}
}

func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	return nil
}

func (m *MockRows) Values() ([]interface{}, error) {
	return nil, nil
}

func (m *MockRows) RawValues() [][]byte {
	return nil
}

func (m *MockRows) Conn() *pgx.Conn {
	return nil
}

// Mock DB
type MockDB struct {
	queryResult     pgx.Rows
	queryRowResult  pgx.Row
	execResult      pgconn.CommandTag
	queryError      error
	queryRowError   error
	execError       error
}

func (m *MockDB) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if m.queryError != nil {
		return nil, m.queryError
	}
	return m.queryResult, nil
}

func (m *MockDB) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if m.queryRowResult != nil {
		return m.queryRowResult
	}
	return &MockRow{err: m.queryRowError}
}

func (m *MockDB) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if m.execError != nil {
		return pgconn.CommandTag{}, m.execError
	}
	return m.execResult, nil
}

// Test GetAllItems - Success
func TestGetAllItems_Success(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	mockRows := &MockRows{
		data: [][]interface{}{
			{1, 1, "Elektronik", "MacBook Pro M1", 20000000.0, purchaseDate, 700},
			{2, 2, "Furniture", "Meja Kerja", 1500000.0, purchaseDate, 700},
		},
	}

	mockDB := &MockDB{
		queryResult: mockRows,
	}

	repo := RepositoryManagement{DB: mockDB}
	items, err := repo.GetAllItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Name != "MacBook Pro M1" {
		t.Errorf("Expected 'MacBook Pro M1', got %s", items[0].Name)
	}
}

// Test GetAllItems - Error
func TestGetAllItems_Error(t *testing.T) {
	mockDB := &MockDB{
		queryError: errors.New("database error"),
	}

	repo := RepositoryManagement{DB: mockDB}
	_, err := repo.GetAllItems()

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test AddItem - Success
func TestAddItem_Success(t *testing.T) {
	mockDB := &MockDB{
		execResult: pgconn.CommandTag{},
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.AddItem(1, "Test Item", 1000000.0, "2023-01-15")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test AddItem - Error
func TestAddItem_Error(t *testing.T) {
	mockDB := &MockDB{
		execError: errors.New("insert failed"),
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.AddItem(1, "Test Item", 1000000.0, "2023-01-15")

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test GetItemById - Success
func TestGetItemById_Success(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	mockRow := &MockRow{
		data: []interface{}{1, 1, "Elektronik", "MacBook Pro M1", 20000000.0, purchaseDate, 700},
	}

	mockDB := &MockDB{
		queryRowResult: mockRow,
	}

	repo := RepositoryManagement{DB: mockDB}
	item, err := repo.GetItemById(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item.Name != "MacBook Pro M1" {
		t.Errorf("Expected 'MacBook Pro M1', got %s", item.Name)
	}

	if item.Price != 20000000.0 {
		t.Errorf("Expected 20000000.0, got %f", item.Price)
	}
}

// Test GetItemById - Error
func TestGetItemById_Error(t *testing.T) {
	mockDB := &MockDB{
		queryRowError: errors.New("item not found"),
	}

	repo := RepositoryManagement{DB: mockDB}
	_, err := repo.GetItemById(999)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test UpdateItem - Success
func TestUpdateItem_Success(t *testing.T) {
	mockDB := &MockDB{
		execResult: pgconn.CommandTag{},
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.UpdateItem(1, 2, "Updated Item", 2000000.0, "2023-06-01")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test UpdateItem - Error
func TestUpdateItem_Error(t *testing.T) {
	mockDB := &MockDB{
		execError: errors.New("update failed"),
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.UpdateItem(1, 2, "Updated Item", 2000000.0, "2023-06-01")

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test DeleteItem - Success
func TestDeleteItem_Success(t *testing.T) {
	mockDB := &MockDB{
		execResult: pgconn.CommandTag{},
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.DeleteItem(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test DeleteItem - Error
func TestDeleteItem_Error(t *testing.T) {
	mockDB := &MockDB{
		execError: errors.New("delete failed"),
	}

	repo := RepositoryManagement{DB: mockDB}
	err := repo.DeleteItem(1)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Test SearchItemsByName - Success
func TestSearchItemsByName_Success(t *testing.T) {
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Laptop Dell", 15000000.0, purchaseDate, 30},
				{3, 1, "Electronics", "Laptop HP", 12000000.0, purchaseDate, 45},
			},
		},
	}

	repo := RepositoryManagement{DB: mockDB}
	items, err := repo.SearchItemsByName("Laptop")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Name != "Laptop Dell" {
		t.Errorf("Expected first item name to be 'Laptop Dell', got '%s'", items[0].Name)
	}
}

// Test SearchItemsByName - No Results
func TestSearchItemsByName_NoResults(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{},
		},
	}

	repo := RepositoryManagement{DB: mockDB}
	items, err := repo.SearchItemsByName("NonExistent")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// Test SearchItemsByName - Error
func TestSearchItemsByName_Error(t *testing.T) {
	mockDB := &MockDB{
		queryError: errors.New("search failed"),
	}

	repo := RepositoryManagement{DB: mockDB}
	_, err := repo.SearchItemsByName("Laptop")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test SearchItemsByName - Scan Error
func TestSearchItemsByName_ScanError(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Laptop", 15000000.0, time.Now(), 30},
			},
			err: errors.New("scan error"),
		},
	}

	repo := RepositoryManagement{DB: mockDB}
	_, err := repo.SearchItemsByName("Laptop")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}