// Importar package main para el script "Decodificacion XOR"
package main

//Importar librerias que se utlizan en el script
import (
	"fmt"
	"os"
	"strings"
)

// Funcion decodifica el texto con base a la clave, Recibe valores devuelve texto desencriptado o decodificado
func xorDecodificar(textoCifrado []byte, key []byte) []byte {
	textoDecodificado := make([]byte, len(textoCifrado))
	keyIndex := 0
	for i, char := range textoCifrado {
		textoDecodificado[i] = char ^ key[keyIndex]
		keyIndex = (keyIndex + 1) % len(key)
	}
	return textoDecodificado
}

// Funcion Main la cual ejecuta el paquete principal
func main() {
	// Texto cifrado
	textoCifrado := []byte{40, 1, 28, 11, 13, 13, 20, 3, 10, 1, 3, 66, 32, 13, 19, 13, 2, 5, 3, 66, 42, 5, 6, 11, 10, 68, 32, 3, 28, 0, 31, 66, 44, 1, 4, 3, 0, 7, 31, 23, 28, 16, 24, 78, 78, 12, 17, 17, 78, 8, 31, 5, 28, 5, 20, 13, 78, 0, 21, 17, 13, 13, 22, 16, 15, 22, 80, 7, 2, 68, 29, 7, 0, 23, 17, 8, 11, 74, 80, 66, 47, 12, 31, 16, 15, 72, 80, 18, 15, 22, 17, 66, 31, 17, 21, 66, 39, 10, 3, 11, 24, 1, 80, 16, 11, 7, 31, 12, 1, 30, 19, 3, 78, 16, 5, 66, 2, 11, 23, 16, 1, 72, 80, 17, 27, 6, 21, 66, 11, 8, 80, 1, 1, 0, 21, 66, 13, 11, 30, 66, 11, 8, 80, 19, 27, 1, 80, 16, 11, 23, 31, 14, 24, 13, 3, 22, 11, 68, 21, 17, 26, 1, 80, 7, 4, 1, 2, 1, 7, 7, 25, 13, 78, 1, 30, 66, 41, 13, 4, 42, 27, 6, 95, 37, 7, 16, 60, 3, 12, 68, 9, 66, 13, 11, 29, 18, 15, 22, 4, 7, 78, 1, 28, 66, 11, 10, 28, 3, 13, 1, 80, 3, 78, 23, 31, 18, 1, 22, 4, 7, 46, 13, 30, 17, 7, 18, 21, 76, 13, 8, 94}

	// Archivo de salida
	fileName := "resultado.txt"

	// Abrir el archivo para escritura
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Como la clave esta entre [a-z] estos estan comprendidos entre los ASCII 97 a 122, Utiliza ciclos anidados que permite evaluar todas las posibles convinaciones de clave para desencriptacion.
	for a := 97; a <= 122; a++ {
		for b := 97; b <= 122; b++ {
			for c := 97; c <= 122; c++ {
				for d := 97; d <= 122; d++ {
					key := []byte{byte(a), byte(b), byte(c), byte(d)}
					textoDecodificado := xorDecodificar(textoCifrado, key)

					//Convierte de Bytes a string
					str := string(textoDecodificado)

					//Realiza split y posteriormente hace validacion para la decodificacion del texto coherente, e imprime en consola y en archivo resultado.txt
					parts := strings.Split(str, " ")
					line := ("")
					for _, part := range parts {
						if part == "Insive" {
							line = fmt.Sprintf("Key: %s, Texto codificado: %s\n", key, textoDecodificado)
							fmt.Printf("La clave de texto es: %s, y el Texto decodificado es: %s\n", key, textoDecodificado)
						}

					}

					// Escribe la línea en el archivo
					_, err := file.WriteString(line)
					if err != nil {
						fmt.Println("Error al escribir en el archivo:", err)
						return
					}
				}
			}
		}
	}

	fmt.Println("Resultados exportados correctamente a", fileName)
}
