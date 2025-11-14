document.addEventListener("DOMContentLoaded", () => {
  // Verificar si el usuario está autenticado
  const user = JSON.parse(localStorage.getItem("user"));
  if (!user) {
    alert("⚠️ Debes iniciar sesión para acceder");
    window.location.href = "/static/login.html";
    return;
  }

  console.log("✅ Usuario autenticado:", user.Email || user.user_email);

  // Variables globales
  let allVehicles = [];
  let currentPage = 1;
  const vehiclesPerPage = 6;

  // Cargar vehículos al iniciar
  loadVehicles();

  // Cerrar sesión
  const logoutLinks = document.querySelectorAll('a[href="#"]');
  logoutLinks.forEach((logoutBtn) => {
    if (logoutBtn.textContent.includes("Cerrar Sesión")) {
      logoutBtn.addEventListener("click", (e) => {
        e.preventDefault();
        localStorage.removeItem("user");
        console.log("🔓 Sesión cerrada");
        window.location.href = "/static/login.html";
      });
    }
  });

  // Función para cargar vehículos desde el backend
  async function loadVehicles() {
    try {
      console.log("🔄 Cargando vehículos desde el backend...");

      const response = await fetch(
        "http://localhost:8080/api/vehicles/getAll?limit=1000",
        {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        },
      );

      console.log("📡 Response status:", response.status);

      if (response.ok) {
        const data = await response.json();
        console.log("📦 Datos recibidos:", data);

        // El backend devuelve un array directamente
        allVehicles = Array.isArray(data) ? data : data.vehicles || [];
        console.log(`✅ ${allVehicles.length} vehículos cargados`);

        if (allVehicles.length > 0) {
          console.log("📋 Primer vehículo:", allVehicles[0]);
          console.log(
            "🖼️ Imágenes del primer vehículo:",
            allVehicles[0].vehicle_image,
          );
        }

        displayVehicles(allVehicles);
        setupPagination(allVehicles);
      } else {
        console.error("❌ Error al cargar vehículos, status:", response.status);
        alert("⚠️ Error al cargar los vehículos");
      }
    } catch (err) {
      console.error("⚠️ Error de conexión:", err);
      alert("⚠️ Error de conexión con el servidor");
    }
  }

  // Función para mostrar vehículos
  function displayVehicles(vehicles) {
    const vehiclesList = document.querySelector(".vehicles-list");

    if (!vehiclesList) {
      console.error("❌ No se encontró el contenedor .vehicles-list");
      return;
    }

    vehiclesList.innerHTML = "";

    if (vehicles.length === 0) {
      vehiclesList.innerHTML =
        '<p style="text-align: center; width: 100%; padding: 40px; color: #666; font-size: 18px;">📭 No hay vehículos disponibles</p>';
      return;
    }

    const start = (currentPage - 1) * vehiclesPerPage;
    const end = start + vehiclesPerPage;
    const paginatedVehicles = vehicles.slice(start, end);

    console.log(
      `📄 Mostrando vehículos ${start + 1} a ${Math.min(end, vehicles.length)} de ${vehicles.length}`,
    );

    paginatedVehicles.forEach((vehicle) => {
      const vehicleCard = createVehicleCard(vehicle);
      vehiclesList.innerHTML += vehicleCard;
    });

    // Agregar event listeners a los botones
    addVehicleCardListeners();
  }

  // Función para crear tarjeta de vehículo
  function createVehicleCard(vehicle) {
    const vehicleId = vehicle.id || vehicle.ID;
    const vehicleBrand =
      vehicle.vehicle_brand || vehicle.brand || "Marca desconocida";
    const vehicleModel =
      vehicle.vehicle_model || vehicle.model || "Modelo desconocido";
    const vehicleYear = vehicle.vehicle_year || vehicle.year || "";
    const vehiclePrice = vehicle.vehicle_price || vehicle.price || 0;
    const vehicleTitle =
      vehicle.vehicle_title ||
      vehicle.titulo ||
      `${vehicleBrand} ${vehicleModel}`;
    const vehicleState = vehicle.vehicle_state || vehicle.state || "N/A";
    const vehicleVersion = vehicle.vehicle_version || vehicle.version || "N/A";

    // Manejo correcto de las imágenes desde vehicle_image
    let imageUrl = "https://placehold.co/400x240/cccccc/666666?text=Sin+Imagen";

    if (
      vehicle.vehicle_image &&
      Array.isArray(vehicle.vehicle_image) &&
      vehicle.vehicle_image.length > 0
    ) {
      imageUrl = vehicle.vehicle_image[0];
    } else if (
      vehicle.Image &&
      Array.isArray(vehicle.Image) &&
      vehicle.Image.length > 0
    ) {
      imageUrl = vehicle.Image[0];
    }

    console.log(`🖼️ Imagen del vehículo ${vehicleId}:`, imageUrl);

    return `
      <div class="vehicle-card" data-id="${vehicleId}">
        <div class="vehicle-image">
          <img src="${imageUrl}" alt="${vehicleBrand} ${vehicleModel}"
               onerror="this.src='/images/default-car.jpg'">
        </div>
        <div class="vehicle-info">
          <h3>${vehicleBrand} ${vehicleModel} ${vehicleYear}</h3>
          <p class="price">$${formatPrice(vehiclePrice)}</p>
          <div class="vehicle-specs">
            <span>🏷️ Título: ${vehicleTitle}</span>
            <span>📦 Estado: ${vehicleState}</span>
            <span>⚙️ Versión: ${vehicleVersion}</span>
            <span>📅 Año: ${vehicleYear}</span>
          </div>
          <div class="vehicle-actions">
            <button class="btn view-details" data-id="${vehicleId}">Ver Detalles</button>
            <button class="btn contact-seller">Contactar Vendedor</button>
          </div>
        </div>
      </div>
    `;
  }

  // Función para agregar listeners a las tarjetas
  function addVehicleCardListeners() {
    document.querySelectorAll(".view-details").forEach((button) => {
      button.addEventListener("click", function () {
        const vehicleId = this.getAttribute("data-id");
        viewVehicleDetails(vehicleId);
      });
    });

    document.querySelectorAll(".contact-seller").forEach((button) => {
      button.addEventListener("click", function () {
        alert('Funcionalidad de "Contactar Vendedor" en desarrollo');
      });
    });
  }

  // Función para ver detalles de un vehículo
  async function viewVehicleDetails(id) {
    try {
      console.log(`🔍 Obteniendo detalles del vehículo ID: ${id}`);

      const response = await fetch(`http://localhost:8080/api/vehicles/${id}`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      });

      if (response.ok) {
        const data = await response.json();
        console.log("📦 Detalles del vehículo:", data);
        const vehicle = data.vehicle || data;
        showVehicleDetailsModal(vehicle);
      } else {
        console.error("❌ Error al cargar detalles, status:", response.status);
        alert("⚠️ Error al cargar los detalles del vehículo");
      }
    } catch (err) {
      console.error("⚠️ Error:", err);
      alert("⚠️ Error de conexión con el servidor");
    }
  }

  // Función para mostrar modal con detalles
  function showVehicleDetailsModal(vehicle) {
    const vehicleBrand = vehicle.vehicle_brand || vehicle.brand || "N/A";
    const vehicleModel = vehicle.vehicle_model || vehicle.model || "N/A";
    const vehicleYear = vehicle.vehicle_year || vehicle.year || "N/A";
    const vehiclePrice = vehicle.vehicle_price || vehicle.price || 0;
    const vehicleTitle = vehicle.vehicle_title || vehicle.titulo || "N/A";
    const vehicleState = vehicle.vehicle_state || vehicle.state || "N/A";
    const vehicleVersion = vehicle.vehicle_version || vehicle.version || "N/A";

    const details = `
╔════════════════════════════════════════╗
║     DETALLES DEL VEHÍCULO             ║
╚════════════════════════════════════════╝

🏷️ Título: ${vehicleTitle}
🚗 Marca: ${vehicleBrand}
📛 Modelo: ${vehicleModel}
📅 Año: ${vehicleYear}
💰 Precio: $${formatPrice(vehiclePrice)}
⚙️ Versión: ${vehicleVersion}
📦 Estado: ${vehicleState}
    `;
    alert(details);
  }

  // Función para configurar la paginación
  function setupPagination(vehicles) {
    const totalPages = Math.ceil(vehicles.length / vehiclesPerPage);
    const pagination = document.querySelector(".pagination");

    if (!pagination) {
      console.error("❌ No se encontró el contenedor .pagination");
      return;
    }

    pagination.innerHTML = "";

    if (totalPages <= 1) return;

    const prevBtn = document.createElement("button");
    prevBtn.className = "btn pagination-btn";
    prevBtn.textContent = "« Anterior";
    prevBtn.disabled = currentPage === 1;
    prevBtn.addEventListener("click", () => {
      if (currentPage > 1) {
        currentPage--;
        displayVehicles(vehicles);
        setupPagination(vehicles);
        window.scrollTo({ top: 0, behavior: "smooth" });
      }
    });
    pagination.appendChild(prevBtn);

    for (let i = 1; i <= totalPages; i++) {
      const pageBtn = document.createElement("button");
      pageBtn.className = `btn pagination-btn ${i === currentPage ? "active" : ""}`;
      pageBtn.textContent = i;
      pageBtn.addEventListener("click", () => {
        currentPage = i;
        displayVehicles(vehicles);
        setupPagination(vehicles);
        window.scrollTo({ top: 0, behavior: "smooth" });
      });
      pagination.appendChild(pageBtn);
    }

    const nextBtn = document.createElement("button");
    nextBtn.className = "btn pagination-btn";
    nextBtn.textContent = "Siguiente »";
    nextBtn.disabled = currentPage === totalPages;
    nextBtn.addEventListener("click", () => {
      if (currentPage < totalPages) {
        currentPage++;
        displayVehicles(vehicles);
        setupPagination(vehicles);
        window.scrollTo({ top: 0, behavior: "smooth" });
      }
    });
    pagination.appendChild(nextBtn);
  }

  // Funciones auxiliares
  function formatPrice(price) {
    return new Intl.NumberFormat("es-AR").format(price);
  }

  function formatNumber(number) {
    return new Intl.NumberFormat("es-AR").format(number);
  }

  console.log("✅ RequestVehicle.js cargado correctamente");
});
