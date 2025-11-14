document.addEventListener("DOMContentLoaded", () => {
  const inputs = document.querySelectorAll(".input");

  function focusFunc() {
    let parent = this.parentNode;
    parent.classList.add("focus");
  }

  function blurFunc() {
    let parent = this.parentNode;
    if (this.value == "") {
      parent.classList.remove("focus");
    }
  }

  inputs.forEach((input) => {
    input.addEventListener("focus", focusFunc);
    input.addEventListener("blur", blurFunc);
  });

  // Manejo del formulario de contacto
  const contactForm = document.querySelector("form");

  contactForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    // Obtener valores del formulario
    const name = document.querySelector('input[name="name"]').value.trim();
    const email = document.querySelector('input[name="email"]').value.trim();
    const phone = document.querySelector('input[name="phone"]').value.trim();
    const message = document
      .querySelector('textarea[name="message"]')
      .value.trim();

    // Validación básica
    if (!name || !email || !message) {
      alert(
        "⚠️ Por favor completa todos los campos obligatorios (Nombre, Email y Mensaje)",
      );
      return;
    }

    // Validación de email
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      alert("⚠️ Por favor ingresa un email válido");
      return;
    }

    try {
      // Mostrar indicador de carga
      const submitButton = contactForm.querySelector(".btn");
      const originalText = submitButton.value;
      submitButton.value = "Enviando...";
      submitButton.disabled = true;

      const response = await fetch("http://localhost:8080/api/contact", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name,
          email,
          phone,
          message,
        }),
      });

      const data = await response.json();

      if (response.ok) {
        alert("✅ " + data.message);
        contactForm.reset();

        // Remover la clase focus de todos los inputs
        document.querySelectorAll(".input-container").forEach((container) => {
          container.classList.remove("focus");
        });
      } else {
        alert("❌ " + (data.error || "Error al enviar el mensaje"));
      }

      // Restaurar botón
      submitButton.value = originalText;
      submitButton.disabled = false;
    } catch (err) {
      console.error("Error de conexión:", err);
      alert(
        "⚠️ Error de conexión con el servidor. Por favor, intenta nuevamente.",
      );

      // Restaurar botón en caso de error
      const submitButton = contactForm.querySelector(".btn");
      submitButton.value = "Enviar";
      submitButton.disabled = false;
    }
  });
});
