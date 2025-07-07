package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Wallet struct {
	ID       string
	Nombre   string
	Apellido string
}

type Transaction struct {
	Monto     float64
	OrigenID  string
	DestinoID string
	Timestamp time.Time
}

type Block struct {
	Data         Transaction
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Siguiente    *Block
}

type Blockchain struct {
	Head *Block
}

func CrearBilletera(id, nombre, apellido string) Wallet {
	return Wallet{
		ID:       id,
		Nombre:   nombre,
		Apellido: apellido,
	}
}

func NuevaTransaccion(monto float64, origenID, destinoID string) Transaction {
	return Transaction{
		Monto:     monto,
		OrigenID:  origenID,
		DestinoID: destinoID,
		Timestamp: time.Now(),
	}
}

func calcularHash(b *Block) string {
	data := b.PreviousHash +
		b.Data.OrigenID + b.Data.DestinoID +
		strconv.FormatFloat(b.Data.Monto, 'f', 6, 64) +
		b.Data.Timestamp.String() +
		b.Timestamp.String()

	hash := sha256.Sum256([]byte(data))//Calcula el hash SHA-256 del contenido de data
	//Devuelve un valor fijo de 32 bytes 
	return fmt.Sprintf("%x", hash)//lo convierto a hexadecimal para mejor legibilidad
}

func (bc *Blockchain) InsertarBloque(t Transaction) {
	newBlock := &Block{
		Data:      t,
		Timestamp: time.Now(),
	}

	if bc.Head == nil {
		newBlock.PreviousHash = ""
	} else {
		ultimo := bc.Head
		for ultimo.Siguiente != nil {
			ultimo = ultimo.Siguiente
		}
		newBlock.PreviousHash = ultimo.Hash
		ultimo.Siguiente = newBlock
	}

	newBlock.Hash = calcularHash(newBlock) // Calcular el hash del nuevo bloque

	if bc.Head == nil {
		bc.Head = newBlock
	}
}

func (bc Blockchain) ObtenerSaldo(id string) float64 { // Obtiene el saldo de una billetera dada su ID
	saldo := 0.0
	actual := bc.Head
	//Calcula el saldo de la billetera segun el id, recorriendo la blockchain
	for actual != nil { // Recorro la cadena de bloques
		if actual.Data.OrigenID == id {
			saldo -= actual.Data.Monto //Si el id aparece como OrigenID, resto el monto
		}
		if actual.Data.DestinoID == id {
			saldo += actual.Data.Monto //Si aparece como DestinoID, lo sumo
		}
		actual = actual.Siguiente
	}
	return saldo
}

func (bc Blockchain) ValidarCadena() bool {
	actual := bc.Head
	for actual != nil && actual.Siguiente != nil { // Recorre la cadena de bloques
		if actual.Siguiente.PreviousHash != actual.Hash { // Verifica que el hash del bloque siguiente coincida con el hash del bloque actual
			return false
		}
		calculado := calcularHash(actual) // Calcula el hash del bloque actual
		if actual.Hash != calculado {     // Verifica que el hash del bloque actual coincida con el hash calculado
			return false
		}
		actual = actual.Siguiente // Avanza al siguiente bloque
	}
	return true
}

func (bc Blockchain) PuedeTransferir(origenID string, monto float64) bool {
	return bc.ObtenerSaldo(origenID) >= monto // Verifica si la billetera tiene suficiente saldo para realizar la transferencia
}


func main() {
	bc := Blockchain{}

	pedro := CrearBilletera("A1", "Pedro", "Vega")
	lucio := CrearBilletera("B1", "Lucio", "Vaccarini")

	transaccion1 := NuevaTransaccion(70, "SISTEMA", pedro.ID)
	bc.InsertarBloque(transaccion1)

	if bc.PuedeTransferir(pedro.ID, 30) {
		transaccion2 := NuevaTransaccion(30, pedro.ID, lucio.ID)
		bc.InsertarBloque(transaccion2)
	} else {
		fmt.Println("Saldo insuficiente para transferir.")
	}

	fmt.Println("Saldo de Pedro:", bc.ObtenerSaldo(pedro.ID))
	fmt.Println("Saldo de Lucio:", bc.ObtenerSaldo(lucio.ID))
	fmt.Println("Blockchain válida:", bc.ValidarCadena())
}
