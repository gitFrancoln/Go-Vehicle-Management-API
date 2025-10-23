// Animaciones al hacer scroll
document.addEventListener('DOMContentLoaded', function() {
  // Selecciona todas las secciones que quieres animar
  const sections = document.querySelectorAll('.fade-in-section');
  
  // Configuración del Intersection Observer
  const observerOptions = {
      root: null, // viewport
      rootMargin: '0px',
      threshold: 0.1 // 10% de la sección visible para activar la animación
  };
  
  // Callback cuando una sección entra/sale del viewport
  const observerCallback = (entries, observer) => {
      entries.forEach(entry => {
          if (entry.isIntersecting) {
              entry.target.classList.add('visible');
              // Una vez animada, dejamos de observar para mejor rendimiento
              observer.unobserve(entry.target);
              
              // Si es la sección de productos, anima las tarjetas individualmente
              if (entry.target.querySelector('.products-container')) {
                  const productCards = entry.target.querySelectorAll('.product-card');
                  productCards.forEach((card, index) => {
                      setTimeout(() => {
                          card.style.opacity = '1';
                          card.style.transform = 'translateY(0)';
                      }, index * 100); // Retraso progresivo para cada tarjeta
                  });
              }
          }
      });
  };
  
  // Crea el observer
  const observer = new IntersectionObserver(observerCallback, observerOptions);
  
  // Observa cada sección
  sections.forEach(section => {
      observer.observe(section);
  });
  
  // Resto de tu funcionalidad existente
  // Add interactivity for quantity buttons
  document.querySelectorAll('.quantity-btn').forEach(button => {
      button.addEventListener('click', function() {
          const input = this.parentElement.querySelector('.quantity-input');
          let value = parseInt(input.value);
          
          if (this.textContent === '+') {
              value++;
          } else if (this.textContent === '-' && value > 1) {
              value--;
          }
          
          input.value = value;
      });
  });

  // Add to cart functionality
  document.querySelectorAll('.add-to-cart').forEach(button => {
      button.addEventListener('click', function() {
          alert('¡Producto agregado al carrito!');
      });
  });

  // Smooth scrolling for anchor links
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
      anchor.addEventListener('click', function(e) {
          e.preventDefault();
          const targetId = this.getAttribute('href');
          const targetElement = document.querySelector(targetId);
          
          if (targetElement) {
              window.scrollTo({
                  top: targetElement.offsetTop - 80,
                  behavior: 'smooth'
              });
          }
      });
  });

  // Cart counter functionality
  document.querySelector('.cart-icon').addEventListener('click', function() {
      alert('Carrito de compras: 1 artículo');
  });
});