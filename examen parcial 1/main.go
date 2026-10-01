package main

import (
	"fmt"
)

var productosVendidos []string
var subtotalesVentas []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	
	productosVendidos = append(productosVendidos, nombre)
	subtotalesVentas = append(subtotalesVentas, subtotal)
	
	fmt.Printf("\n¡Venta registrada con éxito! Producto: %s | Cantidad: %d | Subtotal: $%.2f\n", nombre, cantidad, subtotal)
}

func MostrarEstadisticas() {
	fmt.Println("\n=== Estadísticas de Ventas ===")
	
	if len(productosVendidos) == 0 {
		fmt.Println("No existen ventas registradas hasta el momento.")
		return
	}

	var totalRecaudado float64 = 0

	fmt.Println("Detalle de ventas:")
	for i := 0; i < len(productosVendidos); i++ {
		fmt.Printf("- %s: $%.2f\n", productosVendidos[i], subtotalesVentas[i])
		totalRecaudado += subtotalesVentas[i]
	}

	fmt.Printf("\nTotal recaudado en el día: $%.2f\n", totalRecaudado)
}

func main() {
	var opcion string

	for {
		fmt.Println("\n=== Menú Principal - Tienda ===")
		fmt.Println("1 Registrar una nueva venta")
		fmt.Println("2 Mostrar estadísticas")
		fmt.Println("0 Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case "1":
			fmt.Println("\n--- Productos Disponibles ---")
			fmt.Println("1. Arroz  - $1.25")
			fmt.Println("2. Leche  - $0.95")
			fmt.Println("3. Pan    - $0.50")
			fmt.Print("Seleccione el número del producto: ")

			var productoOp int
			var cantidad int
			fmt.Scan(&productoOp)

			var nombreProducto string
			var precioProducto float64

			switch productoOp {
			case 1:
				nombreProducto = "Arroz"
				precioProducto = 1.25
			case 2:
				nombreProducto = "Leche"
				precioProducto = 0.95
			case 3:
				nombreProducto = "Pan"
				precioProducto = 0.50
			default:
				fmt.Println("Opción de producto inválida. Intente de nuevo.")
				continue
			}

			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			if cantidad <= 0 {
				fmt.Println("La cantidad debe ser mayor a cero.")
				continue
			}

			RegistrarVenta(nombreProducto, precioProducto, cantidad)

		case "2":
			MostrarEstadisticas()

		default:
			fmt.Println("Opción incorrecta, intente de nuevo.")
		}
	}
}