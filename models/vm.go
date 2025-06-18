// Structures and DB operations
package models

import (
	"encoding/json"

	"github.com/odeeka/go-minicloud-rest-api/db"
)

type VM struct {
	ID          int64             `json:"id"` // DB autoincrement
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	CPU         float64           `json:"cpu"`
	Memory      int               `json:"memory"` // MB
	Ports       []int             `json"ports"`   // [80, 443]
	Env         map[string]string `json:"env"`    // {"Env":"dev"}
	ContainerID string            `json:"container_id"`
}

// Classical CRUD methods
func GetAllVms() ([]VM, error) {
	query := "SELECT * FROM vms"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vms []VM

	for rows.Next() {
		var vm VM
		var portsStr string
		var envStr string
		err := rows.Scan(&vm.ID, &vm.Name, &vm.Image, &vm.CPU, &vm.Memory, &portsStr, &envStr, &vm.ContainerID)

		if err != nil {
			return nil, err
		}

		// Convert from JSON string to Go type
		if err := json.Unmarshal([]byte(portsStr), &vm.Ports); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(envStr), &vm.Env); err != nil {
			return nil, err
		}

		vms = append(vms, vm)
	}

	return vms, nil
}

func (vm *VM) InsertVM() error {

	portsJSON, _ := json.Marshal(vm.Ports)
	envJSON, _ := json.Marshal(vm.Env)

	query := `
	INSERT INTO vms (name, image, cpu, memory, ports, env, container_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(vm.Name, vm.Image, vm.CPU, vm.Memory, string(portsJSON), string(envJSON), vm.ContainerID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	vm.ID = id
	return err
}

func GetVMByID(id int64) (*VM, error) {
	query := "SELECT * FROM vms WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var vm VM

	portsJSON, _ := json.Marshal(&vm.Ports)
	envJSON, _ := json.Marshal(&vm.Env)

	err := row.Scan(&vm.ID, &vm.Name, &vm.Image, &vm.CPU, &vm.Memory, &portsJSON, &envJSON, &vm.ContainerID)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

func (vm *VM) DeleteVMByID() error {
	query := "DELETE FROM vms WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(vm.ID)
	return err
}

func (vm *VM) UpdateVMByID() error {
	query := `
	UPDATE vms
	SET name = ?, image = ?, cpu = ?, memory = ?, ports = ?, env = ?, container_id = ?
	WHERE id = ?
	`

	portsJSON, _ := json.Marshal(vm.Ports)
	envJSON, _ := json.Marshal(vm.Env)

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(vm.Name, vm.Image, vm.CPU, vm.Memory, string(portsJSON), string(envJSON), vm.ContainerID, vm.ID)
	return err
}
