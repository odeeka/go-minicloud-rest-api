// Structures and DB operations
package models

import (
	"github.com/odeeka/go-minicloud-rest-api/db"
)

type Storage struct {
	ID          int64  `json:"id"` // DB autoincrement
	Name        string `json:"name"`
	SizeGB      int    `json:"size_gb"`
	VMID        *int64 `json:"vm_id"` // Nullable without attachment
	ContainerID string `json:"container_id"`
}

// Classical CRUD methods
func GetAllStorages() ([]Storage, error) {
	query := "SELECT * FROM storages"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storages []Storage

	for rows.Next() {
		var storage Storage
		err := rows.Scan(&storage.ID, &storage.Name, &storage.SizeGB, &storage.VMID, &storage.ContainerID)

		if err != nil {
			return nil, err
		}

		storages = append(storages, storage)
	}

	return storages, nil
}

func (storage *Storage) InsertStorage() error {

	query := `
	INSERT INTO storages (name, size_gb, vm_id, container_id) 
	VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(storage.Name, storage.SizeGB, storage.VMID, storage.ContainerID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	storage.ID = id
	return err
}

func GetStorageByID(id int64) (*Storage, error) {
	query := "SELECT * FROM storages WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var storage Storage

	err := row.Scan(&storage.ID, &storage.Name, &storage.SizeGB, &storage.VMID, &storage.ContainerID)
	if err != nil {
		return nil, err
	}

	return &storage, nil
}

func (storage *Storage) DeleteStorageByID() error {
	query := "DELETE FROM storages WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(storage.ID)
	return err
}

func (storage *Storage) UpdateStorageSizeByID() error {
	query := `
	UPDATE storages
	SET size_gb = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(storage.SizeGB, storage.ID)
	return err
}

func (storage *Storage) AttachStorageByID() error {
	query := `UPDATE storages SET vm_id = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(storage.VMID, storage.ID)
	return err
}

func (storage *Storage) DetachStorageByID() error {
	query := `UPDATE storages SET vm_id = -1 WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(storage.ID)
	return err
}
