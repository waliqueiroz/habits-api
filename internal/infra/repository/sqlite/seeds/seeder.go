package seeds

import (
	"database/sql"
	"log"
	"reflect"
)

type Seed struct {
	db *sql.DB
}

// Execute executará o método seeder fornecido
func Execute(db *sql.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Executa todos seeders se nenhum nome de método for fornecido
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// Itera sobre os métodos da struct seeder
		for i := 0; i < seedType.NumMethod(); i++ {
			// Pega o método da posição atual
			method := seedType.Method(i)
			// Executa o método
			seed(s, method.Name)
		}
	}

	// Executa somente os métodos passados como parâmetro
	for _, seedMethodName := range seedMethodNames {
		seed(s, seedMethodName)
	}
}

func seed(s Seed, seedMethodName string) {
	// Busca o método usando reflection
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Interrompe a execução se o método não existir
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Executa o método
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeed")
}
