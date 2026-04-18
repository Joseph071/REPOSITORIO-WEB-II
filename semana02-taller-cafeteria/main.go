package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}
type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("no hay clientes")
		return

	}
	for _, c := range clientes {
		fmt.Printf("[%d] %s | Carrera: %s | Saldo: %.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)

	}

}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1 // convención: -1 significa "no encontrado"
}

func agregarCliente(clientes []Cliente, cliente Cliente) []Cliente {
	return append(clientes, cliente)
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func ListarProducto(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS REGISTRADOS ===")
	if len(productos) == 0 {
		fmt.Println("no hay productos")
		return

	}
	for _, c := range productos {
		fmt.Printf("[%d] PRODUCTO: %s | PRECIO: %.2F | STOCK: %.d | CATEGORIA: %s |\n", c.ID, c.Nombre, c.Precio, c.Stock, c.Categoria)

	}

}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, c := range productos {
		if c.ID == id {
			return i
		}
	}
	return -1 // convención: -1 significa "no encontrado"
}

func agregarProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", id)
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func DescontarStock(producto *Producto, cantidad int) error {

	if producto.Stock < cantidad {
		return fmt.Errorf("stock insuficiente en %s (hay %d, solicita %d)",
			producto.Nombre, producto.Stock, cantidad)
	}
	producto.Stock -= cantidad
	return nil
}

func DescontarSaldo(cliente *Cliente, total float64) error {

	if cliente.Saldo < total {
		return fmt.Errorf("saldo insuficiente en %s (hay %.2f, solicita %.2f)",
			cliente.Nombre, cliente.Saldo, total)
	}
	cliente.Saldo -= total
	return nil
}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {

	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		return pedidos, err
	}

	nuevoID := len(pedidos) + 1
	nueva := Pedido{
		ID:         nuevoID,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nueva)

	return pedidos, nil

	/* (OPCIONAL avanzado: si falla aquí, deberías devolver el stock que
	descontaste en el paso 4) */
}

func PedidosDeCliente(
	pedidos []Pedido,
	clientes []Cliente,
	productos []Producto,
	clienteID int,
) {

	fmt.Printf("\n=== REPORTE: Pedidos del ciente con ID: %d ===\n", clienteID)

}

func mostrarMenu() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  CAFETERIA MANTA                         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  1. Listar productos                     ║")
	fmt.Println("║  2. Listar clientes                      ║")
	fmt.Println("║  3. Agregar cliente                      ║")
	fmt.Println("║  4. Agregar productos                    ║")
	fmt.Println("║  5. Registrar nuevo Pedido               ║")
	fmt.Println("║  6. Ver pedidos de un cliente (reporte)  ║")
	fmt.Println("║  0. Salir                                ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea) // elimina el \n y espacios sobrantes
}

func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1 // convención: -1 si la conversión falla
	}
	return n
}

func leerFloat(lector *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	f, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return -1
	}
	return f
}

func main() {

	clientes := []Cliente{
		{1, "Jostin", "Ti", 500.40},
		{2, "Joseph", "Ti", 100.40},
		{3, "Lisbeth", "Derechos", 500.40},
	}
	productos := []Producto{
		{1, "Cafe", 0.75, 100, "Bebidas"},
		{2, "Capuchino", 1.5, 50, "Bebidas"},
		{3, "Tostada", 1.0, 30, "Comestible"},
		{4, "Batidos", 1.50, 20, "Bebidas"},
	}

	pedidos := []Pedido{}

	var err error

	lector := bufio.NewReader(os.Stdin)
	for {
		mostrarMenu()
		opcion := leerEntero(lector, "Elige una opción: ")
		switch opcion {
		case 1:
			ListarProducto(productos)
		case 2:
			ListarClientes(clientes)

		case 3:
			fmt.Println("\n--- Agregar cliente ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo ($): ")
			clientes = agregarCliente(clientes, Cliente{
				id, nombre, carrera, saldo,
			})
			fmt.Println("✓ Cliente agregado.")

		case 4:
			fmt.Println("\n--- Agregar producto ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio ($): ")
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoria: ")
			categoria := leerLinea(lector)

			productos = agregarProducto(productos, Producto{
				id, nombre, precio, stock, categoria,
			})
			fmt.Println("✓ Producto agregado.")

		case 5:
			fmt.Println("\n--- Registrar pedido ---")
			cid := leerEntero(lector, "ID del cliente: ")
			pid := leerEntero(lector, "ID del producto: ")
			cantidad := leerEntero(lector, "Cantidad: ")
			fmt.Print("Fecha: ")
			fecha := leerLinea(lector)

			pedidos, err = RegistrarPedido(clientes, productos, pedidos, cid, pid, cantidad, fecha)
			if err != nil {
				fmt.Println("⚠ No se pudo registrar:", err)
			} else {
				fmt.Println("✓ Pedido registrada con éxito.")
			}

		case 6:
			fmt.Println("\n--- Reporte Pedido ---")
			id := leerEntero(lector, "ID del Cliente: ")

			PedidosDeCliente(pedidos, clientes, productos, id)

		case 0:
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}

}
