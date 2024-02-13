# Prueba Insive Nicolas Pardo
Este repositorio contiene la prueba para desarrollador para proyecto Insive.

## Software 
go version go1.19.5

## Librerias utilizadas
- bufio
- fmt
- os
- strings

## Paquete principal 
Package main

## Descripcion 

Pureba de desencriptado de texto por medio de XOR. 
En Insive SpA, un desarrollador ha cifrado un texto mediante la aplicación de la puerta lógica XOR a cada letra en su representación ASCII. En la actualidad, enfrenta el problema de haber olvidado la llave secreta utilizada para cifrar el mensaje, cuyo significado resulta crucial para la empresa. El proceso de cifrado se realizó utilizando una llave de longitud 4 con caracteres incluidos por la siguiente regex:

[a-z]

Los caracteres del mensaje encriptado cumple la siguiente regex:

[a-zA-Z0-9\s.,@\-_\/]

