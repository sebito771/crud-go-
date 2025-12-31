const API_URL= "http://localhost:8080/jugadores"


 
export async function obtenerJugadores(){
 const response = await fetch(API_URL);
 if(!response.ok){
    throw new Error(">Error al Obtener los Jugadores");
 } 

 const data = await response.json();
 return data;
 
}

export async function crearJugador(nombre, puntaje) {
  const response = await fetch(API_URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      nombre: nombre,
      puntaje: puntaje
    })
  });

  if (!response.ok) {
    throw new Error("Error al crear jugador");
  }

  const data = await response.json();
  return data;
}


export async function borrarJugador(id) {

    const response = await fetch(`${API_URL}/${id}`,{
        method:"DELETE"
    });

    if (!response.ok){  
        throw new Error("Error al Eliminar Jugador");  
    }

    const data = await response.json()
    return data
}


export async function actualizarJugador(id, nombre = null, puntaje = null) {
  const body = {};

  if (nombre !== null) body.nombre = nombre;
  if (puntaje !== null) body.puntaje = puntaje;

  const response = await fetch(`${API_URL}/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(body)
  });

  if (!response.ok) {
    throw new Error("Error al actualizar jugador");
  }

  return await response.json();
}

