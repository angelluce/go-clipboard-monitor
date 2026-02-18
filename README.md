# 🛡️ CLIPBOARD MONITOR - Portapapeles Seguro
CLIPBOARD MONITOR es una herramienta ligera diseñada para proteger información sensible 
(IPs, contraseñas, tokens) al interactuar con IAs o chats externos. 
El programa vigila tu portapapeles y reemplaza automáticamente las palabras que tú definas 
antes de que las pegues en cualquier lugar.

## 🚀 Instalación rápida

1. **Descarga** la carpeta correspondiente a tu sistema operativo (Windows o Linux).
2. Asegúrate de que el archivo replacements.json esté en la **misma carpeta** que el ejecutable.
3. **Ejecuta el programa**:
   - **Windows**: Doble clic en bda_clip.exe.
   - **Linux**: Ejecuta ./bda_clip (asegúrate de darle permisos con chmod +x bda_clip).

## 🛠️ Cómo usarlo

### 1. El Vigilante (Monitor)

Simplemente, deja la terminal abierta. Mientras el programa diga [VIGILANTE ACTIVO], 
cualquier texto que copies será procesado. Si copias algo que coincide con tus reglas, 
el programa lo limpiará instantáneamente.

### 2. Agregar nuevas palabras

Tienes dos formas de añadir reglas:

- Desde la terminal: Abre una nueva terminal en la carpeta del programa y escribe:

```bash
# Ejemplo
./bda_clip add -p "10.0.0.45" -r "[IP_PROD]"
```

- Editando el JSON: Abre replacements.json con cualquier editor de texto, añade la palabra 
y el reemplazo, y guarda. El monitor cargará el cambio sin necesidad de reiniciar.

### 3. Ver tus reglas actuales

Si quieres saber qué palabras estás protegiendo:

```bash
./bda_clip list
```

## 📋 Comandos disponibles

| Comando                          | Descripción                             |
|----------------------------------|-----------------------------------------|
| *(ninguno)*                      | Inicia el monitor del portapapeles      |
| add -p "busqueda" -r "reemplazo" | Añade una nueva regla de limpieza       |
| list                             | Muestra la tabla de palabras protegidas |
| help                             | Muestra la guía de ayuda rápida         |

## Notas de seguridad

El programa no envía datos a internet. Todo el proceso ocurre localmente en tu memoria RAM. 
Si cierras la ventana de la terminal, el programa dejará de proteger el portapapeles.

## Autor

Desarrollado con ❤️ por [Angel Lucero](https://www.linkedin.com/in/angellucero/)