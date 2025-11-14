// Home.js - Combinación de animaciones de landing page y carga de vehículos del backend
document.addEventListener("DOMContentLoaded", () => {
  console.log("✅ Home.js cargado correctamente");

  // ============================================
  // PARTE 1: ANIMACIONES DE SCROLL (Landing Page)
  // ============================================

  // Selecciona todas las secciones que quieres animar
  const sections = document.querySelectorAll(".fade-in-section");

  // Configuración del Intersection Observer
  const observerOptions = {
    root: null, // viewport
    rootMargin: "0px",
    threshold: 0.1, // 10% de la sección visible para activar la animación
  };

  // Callback cuando una sección entra/sale del viewport
  const observerCallback = (entries, observer) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add("visible");
        // Una vez animada, dejamos de observar para mejor rendimiento
        observer.unobserve(entry.target);

        // Si es la sección de productos, anima las tarjetas individualmente
        if (entry.target.querySelector(".products-container")) {
          const productCards = entry.target.querySelectorAll(".product-card");
          productCards.forEach((card, index) => {
            setTimeout(() => {
              card.style.opacity = "1";
              card.style.transform = "translateY(0)";
            }, index * 100); // Retraso progresivo para cada tarjeta
          });
        }
      }
    });
  };

  // Crea el observer
  const observer = new IntersectionObserver(observerCallback, observerOptions);

  // Observa cada sección
  sections.forEach((section) => {
    observer.observe(section);
  });

  // ============================================
  // PARTE 2: SMOOTH SCROLLING
  // ============================================

  // Smooth scrolling for anchor links
  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", function (e) {
      e.preventDefault();
      const targetId = this.getAttribute("href");
      const targetElement = document.querySelector(targetId);

      if (targetElement) {
        window.scrollTo({
          top: targetElement.offsetTop - 80,
          behavior: "smooth",
        });
      }
    });
  });

  // ============================================
  // PARTE 3: MOBILE MENU
  // ============================================

  const mobileMenuBtn = document.getElementById("mobileMenuBtn");
  const mobileMenu = document.getElementById("mobileMenu");

  if (mobileMenuBtn && mobileMenu) {
    mobileMenuBtn.addEventListener("click", () => {
      mobileMenu.classList.toggle("active");
      mobileMenuBtn.classList.toggle("active");
    });

    // Close mobile menu when clicking on a link
    mobileMenu.querySelectorAll("a").forEach((link) => {
      link.addEventListener("click", () => {
        mobileMenu.classList.remove("active");
        mobileMenuBtn.classList.remove("active");
      });
    });
  }
});
