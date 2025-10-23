// Generar autos aleatorios
const marcas = ["Toyota", "Ford", "Honda", "BMW", "Mercedes-Benz", "Audi", "Chevrolet", "Nissan", "Volkswagen", "Hyundai", "Peugeot", "Kia"];
const modelos = ["Sedan", "SUV", "Hatchback", "Coupe", "Pickup", "Crossover"];
const combustibles = ["Gasolina", "Diésel", "Eléctrico", "Híbrido"];
const transmisiones = ["Automática", "Manual"];

function generarAutos(cantidad) {
  const autos = [];
  for (let i = 0; i < cantidad; i++) {
    const marca = marcas[Math.floor(Math.random() * marcas.length)];
    const modelo = modelos[Math.floor(Math.random() * modelos.length)];
    const combustible = combustibles[Math.floor(Math.random() * combustibles.length)];
    const transmision = transmisiones[Math.floor(Math.random() * transmisiones.length)];
    const precio = "$" + (20000 + Math.floor(Math.random() * 80000)).toLocaleString();
    const km = Math.floor(Math.random() * 100000).toLocaleString() + " km";
    autos.push({
      nombre: `${marca} ${modelo}`,
      modelo: `${marca} ${modelo} ${2020 + Math.floor(Math.random() * 5)}`,
      precio,
      km,
      transmision,
      combustible
    });
  }
  return autos;
}

// Crear 36 autos (3 páginas de 12)
const vehiculos = generarAutos(36);

const cardsContainer = document.getElementById("cardsContainer");
const searchInput = document.getElementById("searchInput");
const pageIndicator = document.getElementById("pageIndicator");

let currentPage = 1;
const itemsPerPage = 12;

function mostrarVehiculos(lista) {
  cardsContainer.innerHTML = "";
  const start = (currentPage - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  const vehiculosPagina = lista.slice(start, end);

  vehiculosPagina.forEach(v => {
    const card = document.createElement("div");
    card.classList.add("card");
    card.innerHTML = `
      <h2>${v.nombre}</h2>
      <p><strong>${v.modelo}</strong></p>
      <p class="price">${v.precio}</p>
      <p>Kilometraje: ${v.km}</p>
      <p>Transmisión: ${v.transmision}</p>
      <p>Combustible: ${v.combustible}</p>
      <div class="buttons">
        <button class="details">Ver Detalles</button>
        <button class="contact">Contactar Vendedor</button>
      </div>
    `;
    cardsContainer.appendChild(card);
  });

  pageIndicator.textContent = currentPage;
}

function buscarVehiculos() {
  const texto = searchInput.value.toLowerCase();
  const filtrados = vehiculos.filter(v =>
    v.nombre.toLowerCase().includes(texto) ||
    v.modelo.toLowerCase().includes(texto)
  );
  currentPage = 1;
  mostrarVehiculos(filtrados);
}

document.getElementById("prevPage").addEventListener("click", () => {
  if (currentPage > 1) {
    currentPage--;
    mostrarVehiculos(vehiculos);
  }
});

document.getElementById("nextPage").addEventListener("click", () => {
  if (currentPage * itemsPerPage < vehiculos.length) {
    currentPage++;
    mostrarVehiculos(vehiculos);
  }
});

searchInput.addEventListener("input", buscarVehiculos);

mostrarVehiculos(vehiculos);
