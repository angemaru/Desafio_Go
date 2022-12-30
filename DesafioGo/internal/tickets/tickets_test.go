package tickets

import (
	//"log"
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Se importa testify

func TestGetTotalTickets(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	destino_1 := "Indonesia"
	destino_2 := "China"
	destino_3 := "Japan"
	impuesto_esperado_1 := 108
	impuesto_esperado_2 := 178
	impuesto_esperado_3 := 14

	// Se ejecuta el test
	resultado_A, err1 := GetTotalTickets(destino_1)
	if err1 != nil {
		log.Fatal(err1)
	}
	resultado_B, err2 := GetTotalTickets(destino_2)
	if err2 != nil {
		log.Fatal(err2)
	}
	resultado_C, err3 := GetTotalTickets(destino_3)
	if err3 != nil {
		log.Fatal(err3)
	}

	// Se validan los resultados
	assert.Equal(t, impuesto_esperado_1, resultado_A, "deben ser iguales")
	assert.Equal(t, impuesto_esperado_2, resultado_B, "deben ser iguales")
	assert.Equal(t, impuesto_esperado_3, resultado_C, "deben ser iguales")
}

func TestAverageDestination(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	destino_1 := "Indonesia"
	destino_2 := "China"
	destino_3 := "Japan"
	promedio_esperado_1 := 10.8
	promedio_esperado_2 := 17.8
	promedio_esperado_3 := 1.4

	// Se ejecuta el test
	resultado_A, err1 := AverageDestination(destino_1, 1000)
	if err1 != nil {
		log.Fatal(err1)
	}
	resultado_B, err2 := AverageDestination(destino_2, 1000)
	if err2 != nil {
		log.Fatal(err2)
	}
	resultado_C, err3 := AverageDestination(destino_3, 1000)
	if err3 != nil {
		log.Fatal(err3)
	}

	// Se validan los resultados
	assert.Equal(t, promedio_esperado_1, resultado_A, "deben ser iguales")
	assert.Equal(t, promedio_esperado_2, resultado_B, "deben ser iguales")
	assert.Equal(t, promedio_esperado_3, math.Round(resultado_C*10)/10, "deben ser iguales")
}

func TestGetCountByPeriod(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	periodo_1 := "noche"
	periodo_2 := "madrugada"
	periodo_3 := "mañana"
	periodo_4 := "tarde"
	cantidad_esperado_1 := 151
	cantidad_esperado_2 := 304
	cantidad_esperado_3 := 256
	cantidad_esperado_4 := 289

	// Se ejecuta el test
	resultado_A, err1 := GetCountByPeriod(periodo_1)
	if err1 != nil {
		log.Fatal(err1)
	}
	resultado_B, err2 := GetCountByPeriod(periodo_2)
	if err2 != nil {
		log.Fatal(err2)
	}
	resultado_C, err3 := GetCountByPeriod(periodo_3)
	if err3 != nil {
		log.Fatal(err3)
	}

	resultado_D, err4 := GetCountByPeriod(periodo_4)
	if err4 != nil {
		log.Fatal(err4)
	}

	// Se validan los resultados
	assert.Equal(t, cantidad_esperado_1, resultado_A, "deben ser iguales")
	assert.Equal(t, cantidad_esperado_2, resultado_B, "deben ser iguales")
	assert.Equal(t, cantidad_esperado_3, resultado_C, "deben ser iguales")
	assert.Equal(t, cantidad_esperado_4, resultado_D, "deben ser iguales")
}
