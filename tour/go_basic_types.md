- `bool`: Representa un valor booleano que puede ser `true` o `false`.

- `string`: Almacena una secuencia de caracteres Unicode.

- `int`, `int8`, `int16`, `int32`, `int64`: Tipos enteros con signo que pueden almacenar números enteros con diferentes tamaños y rangos.

- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`: Tipos enteros sin signo que almacenan números enteros positivos con diferentes tamaños y rangos. `uintptr` se utiliza para almacenar valores que representan direcciones de memoria.

- `byte`: Un alias para `uint8`. Se usa comúnmente para representar datos binarios o valores de 8 bits.

- `rune`: Un alias para `int32`. Se utiliza para representar caracteres Unicode, específicamente puntos de código Unicode.

- `float32`, `float64`: Tipos de punto flotante que pueden almacenar números reales con diferentes precisiones y rangos.

- `complex64`, `complex128`: Tipos que almacenan números complejos con componentes de punto flotante. `complex64` utiliza números de 32 bits para las partes real e imaginaria, mientras que `complex128` utiliza números de 64 bits.

Nota: Los tipos enteros, flotantes y complejos pueden variar en tamaño y rango según la arquitectura del sistema en el que se ejecuta el programa.
