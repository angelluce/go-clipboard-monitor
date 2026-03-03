# 🛡️ Clipboard Monitor - Portapapeles seguro

Clipboard Monitor es una herramienta diseñada para proteger información sensible 
(IPs, contraseñas, tokens) al interactuar con IAs o chats externos. 
El programa vigila tu portapapeles y reemplaza automáticamente las palabras que 
tú definas antes de que las pegues en cualquier lugar.


## 🚀 Compilación rápida

```bash
# Windows
set GOOS=windows
set GOARCH=amd64
go build -o clipboard_monitor.exe main.go

# Linux
set GOOS=linux
set GOARCH=amd64
go build -o clipboard_monitor main.go
```


## 🛠️ Cómo usarlo

### 1. El vigilante (Modo interactivo)

Al ejecutar el programa, entrarás en una consola protegida. El monitor se 
activa automáticamente en segundo plano. No necesitas abrir otras terminales; 
puedes escribir comandos directamente mientras el programa sigue vigilando tu 
portapapeles.

### 2. Agregar nuevas palabras

Ahora puedes agregar palabras simples o frases completas usando comillas. 
Escribe el comando directamente en la consola del programa:

- **Palabras simples:** add IP_LOCAL 127.0.0.1
- **Frases con espacios:** add "Mi empresa segura" "trabajo"

Nota: El uso de comillas " es obligatorio si tu búsqueda o tu reemplazo contienen espacios.

### 3. Ver y gestionar reglas

Desde la misma consola, puedes interactuar con el diccionario de protección:

- Para ver lo que estás protegiendo: Escribe list.
- Para editar manualmente: Abre el archivo `replacements.json`. Los cambios se cargan en tiempo real sin reiniciar.

### 4. Escanear archivos

Puedes analizar y sanitizar archivos aplicando las mismas reglas de protección:

```bash
scan ruta/de/mi/archivo.txt
scan logs/aplicacion.log
```

El comando generará un nuevo archivo con el sufijo `_sanitized` (ej: `archivo_sanitized.txt`) 
manteniendo el original intacto. Actualmente soporta archivos `.txt` y `.log`.


## 📋 Comandos disponibles

| Comando | Formato                  | Descripción                               |
|---------|--------------------------|-------------------------------------------|
| add     | add "buscar" "reemplazo" | Registra una nueva regla (soporta frases) |
| list    | list                     | Muestra la tabla de reglas activas        |
| stats   | stats                    | Muestra las estadísticas de la sesión     |
| scan    | scan ruta/archivo.txt    | Escanea y sanitiza archivos (.txt, .log)  |
| help    | help                     | Muestra la guía rápida de comandos        |


## Notas de seguridad

El programa no envía datos a internet. Todo el proceso ocurre localmente en tu memoria RAM. 
Si cierras la ventana de la terminal, el programa dejará de proteger el portapapeles.


## Autor

Desarrollado con ❤️ por [Angel Lucero](https://www.linkedin.com/in/angellucero/)