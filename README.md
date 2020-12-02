# Tarea2
### Ejecución del código
En todas las maquinas se debe utilizar el comando *make run* y ahí se preguntará
como quiere utilizar el algoritmo.

### Utilización de makefile
- make run: comienza la ejecución de la tarea, para todo tipo de máquina.
- make gen: genera archivos pb.
- make clean: elimina archivos pb.
- make remove: elimina todas las partes de libros del sistema donde se corre.

## Distribución de máquinas:

1. dist41: NameNode
2. dist42: DataNode
3. dist43: DataNode 
4. dist44: DataNode

# Tener en cuenta
- Es necesario de que en todas las maquinas se ejecute el mismo tipo de algoritmo, por lo tanto
no puede existir una combinación de maquinas centralizadas y distribuidas.
- La política de rechazo es estocástica, con un 10% para generar una nueva propuesta.
- Los libros para subir al sistema deben estar dentro de la carpeta llamada **libros**.
- La partes se descargarán dentro de una carpeta llamada **partes**.
- Los libros descargados y unidos quedan en la raiz del programa.
- Se considera que **siempre** están las cuatro máquinas conectadas.
- Si se cierra el NameNode los libros son perdidos. 
- Para subir y bajar libros es necesario que se escriba el nombre identicamente de como está el archivo,
el algorimo no ayudará al usuario a rellenar la información (esto incluye el formato).
