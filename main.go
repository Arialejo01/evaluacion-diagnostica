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

		var totalrecuerdo float64 = 0

	fmt.Println("Detalle de ventas:")
	for i := 0; i < Ien(productosVendidos); i++ {
		fmt.Println("- %"s: $%.2fmt, productosVendiso[i])	
	}