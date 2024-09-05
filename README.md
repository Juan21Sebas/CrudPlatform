# CrudPlatform

La API RESTful desarrollada en Go es una plataforma robusta para la gestión eficiente de tareas. Ofrece a los usuarios la capacidad de realizar operaciones básicas de CRUD (Crear, Leer, Actualizar, Eliminar) sobre tareas, brindando así una experiencia completa de gestión de tareas.

Creación de tareas: Los usuarios pueden crear nuevas tareas proporcionando la información necesaria, como el título, la descripción, la fecha de vencimiento, etc. La API valida los datos proporcionados y crea una nueva tarea en la base de datos.

Lectura de tareas: La API permite a los usuarios acceder a la lista de tareas existentes. Esto les permite ver todas las tareas disponibles, junto con sus detalles, como el título, la descripción y el estado.

Actualización de tareas: Los usuarios pueden modificar los detalles de una tarea existente, como su título, descripción, estado o fecha de vencimiento. La API garantiza que los cambios se reflejen correctamente en la base de datos y se actualicen de manera coherente.

Eliminación de tareas: Los usuarios tienen la capacidad de eliminar tareas que ya no son relevantes o necesarias. Al hacerlo, la API elimina la tarea correspondiente de la base de datos de manera segura.

Además de estas operaciones básicas, la API está diseñada para manejar eficientemente múltiples solicitudes concurrentes. Esto se logra mediante técnicas de concurrencia y goroutines en Go, lo que garantiza un rendimiento óptimo incluso bajo cargas de trabajo intensas. 

## Documentación detallada en Swagger
En el repositorio existe una carpeta llamada swagger, dentro hay un archivo en formato yaml que contiene la documentación detallada de la API, para acceder a ella se recomienda utilizar el siguiente editor: https://editor.swagger.io/

## Documentacion Detallada, explicacion de Prubas de carga con jmeter, digrama de componentes 

https://www.canva.com/design/DAF_ZsVnFRc/saI1oziJxArq_-Lfk_tsjw/edit?utm_content=DAF_ZsVnFRc&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton

## Autorizacion

Recuerde que para la autorizacion se creo un middleware que valida o simula el uso de un token, para que no tenga problemas con el uso de la peticion recuerde en las cabeceras agregar Authorization y en value dejar mi_token_secreto, de lo contrario le saldra un mensaje de Unauthorized

## EndPoints 
POST    users/
GET     users/:id
PUT     users/:id
DELETE  users/:id


### Request example POST:
json
Copy code
{
    "title": "Título de la tarea",
    "description": "Descripción de la tarea",
    "status": "Completadas"
}

### Request example PUT:
json
Copy code
{
    "title": "Cambio",
    "description": "Cambio",
    "status": "Completada"
}


## Local execution
To run this service locally, follow these steps:

- Clone this repository to your local machine.
- Make sure you have Go installed and configured on your system.
- Run go build to compile the project.
- Run ./your-binary-name to start the server.
- Access http://lcambiar to interact with the service locally.

### Autor
#### Juan Sebastian Sanchez Arteta

