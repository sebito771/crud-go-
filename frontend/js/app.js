import {
  obtenerJugadores,
  crearJugador,
  borrarJugador,
  actualizarJugador
} from "./api.js";



console.log("app initialized");

// =========================
// Sección de botones menú
// =========================
const seccionBotones = document.getElementById("botones");

const botonAgregarJugador = document.getElementById("agregar-jugador");
const botonObtenerJugadores = document.getElementById("obtener-jugadores");
const botonBorrarJugador = document.getElementById("borrar-jugador");
const botonActualizarJugador = document.getElementById("actualizar-jugador");
const botonRegresar = document.getElementById("regresar");


// =========================
// Secciones principales
// =========================
const seccionAgregar = document.getElementById("seccion-agregar");
const seccionBorrar = document.getElementById("seccion-borrar");
const seccionActualizar = document.getElementById("seccion-actualizar");
const seccionLista = document.getElementById("seccion-lista");

// =========================
// Formularios
// =========================
const formAgregar = document.getElementById("formulario");
const formBorrar = document.getElementById("formulario-borrar");
const formActualizar = document.getElementById("formulario-actualizar");

// =========================
// Inputs
// =========================
// Agregar
const inputNombre = document.getElementById("nombre");
const inputPuntaje = document.getElementById("puntaje");

// Borrar
const inputIdBorrar = document.getElementById("number-borrar");

// Actualizar
const inputIdActualizar = document.getElementById("number-actualizar");
const inputNombreActualizar = document.getElementById("nombre-actualizar");
const inputPuntajeActualizar = document.getElementById("puntaje-actualizar");

// =========================
// Lista
// =========================
const listaJugadores = document.getElementById("jugadores-lista");



function HideSections(){
    seccionActualizar.style.display= "none";
    seccionAgregar.style.display= "none";
    seccionBorrar.style.display= "none";
    seccionLista.style.display="none";
    seccionBotones.style.display="none";
}

HideSections();


//Eventos click al presionar los diferentes botones

botonAgregarJugador.addEventListener("click", () => {
  HideSections();
  seccionAgregar.style.display = "block";
});

botonObtenerJugadores.addEventListener("click", () => {
  HideSections();
  seccionLista.style.display = "block";
});

botonBorrarJugador.addEventListener("click", () => {
  HideSections();
  seccionBorrar.style.display = "block";
});

botonActualizarJugador.addEventListener("click", () => {
  HideSections();
  seccionActualizar.style.display = "block";
});

botonRegresar.addEventListener("click",()=>{
    HideSections();
    seccionBotones.style.display="block"
})


seccionBotones.style.display="block";

//consumo del Api 
formAgregar.addEventListener("submit",async (e)=>{
e.preventDefault();

const nombre= inputNombre.value.trim();
const puntaje= Number(inputPuntaje.value);

 if (!nombre || isNaN(puntaje)) {
    alert("Datos inválidos");
    return;
  }

  try {
    await crearJugador(nombre, puntaje);
    alert("Jugador creado correctamente");

    formAgregar.reset();
    HideSections();
    seccionBotones.style.display = "block";

  } catch (error) {
    alert(error.message);
  }

});



 formActualizar.addEventListener("submit",async(e)=>{
   e.preventDefault();

   const nombre =
    inputNombreActualizar.value.trim() !== ""
      ? inputNombreActualizar.value.trim()
      : null;

  const puntaje =
    inputPuntajeActualizar.value !== ""
      ? Number(inputPuntajeActualizar.value)
      : null;


  try {
    await actualizarJugador(id, nombre, puntaje);
    alert("Jugador actualizado");
    formActualizar.reset();
  }catch (error) {
    alert(error.message);
  }

 });


 formBorrar.addEventListener("submit",async (e)=>{
 e.preventDefault();
 
  const id = Number(inputIdBorrar.value);

  if (isNaN(id)) {
    alert("ID inválido");
    return;
  }
 try {
   await borrarJugador(id)
   formBorrar.reset();
 } catch (error) {
    alert(error.message);
 };


 });


 botonObtenerJugadores.addEventListener("click", async () => {
  HideSections();
  seccionLista.style.display = "block";

  try {
    const jugadores = await obtenerJugadores();

    listaJugadores.innerHTML = "";

    if (jugadores.length === 0) {
      const li = document.createElement("li");
      li.textContent = "No hay jugadores";
      listaJugadores.appendChild(li);
      return;
    }

    jugadores.forEach(j => {
      const li = document.createElement("li");
      li.textContent = `${j.id} - ${j.nombre} (${j.puntaje})`;
      listaJugadores.appendChild(li);
    });

  } catch (error) {
    alert(error.message);
  }
});




