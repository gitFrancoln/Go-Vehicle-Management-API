document.addEventListener("DOMContentLoaded", () => {
  // Animación de paneles
  const signUpButton = document.getElementById("signUp");
  const signInButton = document.getElementById("signIn");
  const container = document.getElementById("container");

  signUpButton.addEventListener("click", () => {
    container.classList.add("right-panel-active");
  });

  signInButton.addEventListener("click", () => {
    container.classList.remove("right-panel-active");
  });

  // LOGIN
  const loginForm = document.getElementById("loginForm");
  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const user_email = document.getElementById("email").value;
    const user_password = document.getElementById("password").value;

    console.log("🔄 Intentando login con:", user_email);

    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ user_email, user_password }),
      });

      console.log("📡 Response status:", response.status);

      const data = await response.json();
      console.log("📦 Response data:", data);

      if (response.ok) {
        console.log("✅ Login exitoso!");
        localStorage.setItem("user", JSON.stringify(data.user));
        console.log("💾 Usuario guardado en localStorage");
        console.log("🔄 Redirigiendo a Home.html...");

        // Redirección inmediata (replace evita que vuelvan atrás con el botón)
        setTimeout(() => {
          window.location.replace("/static/Home.html");
        }, 100);
      } else {
        console.error("❌ Error en login:", data.error);
        alert("❌ " + data.error);
      }
    } catch (err) {
      console.error("⚠️ Error de conexión:", err);
      alert("⚠️ Error de conexión con el servidor");
    }
  });

  // REGISTRO
  const signupForm = document.getElementById("signupForm");
  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const user_name = document.getElementById("signupName").value;
    const user_email = document.getElementById("signupEmail").value;
    const user_password = document.getElementById("signupPassword").value;

    try {
      const response = await fetch("http://localhost:8080/api/users", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ user_name, user_email, user_password }),
      });

      const data = await response.json();

      if (response.ok) {
        alert("✅ Usuario creado correctamente");
        signupForm.reset();
        container.classList.remove("right-panel-active"); // volver a login
      } else {
        alert("❌ " + (data.error || "Revise los datos ingresados"));
      }
    } catch (err) {
      console.error(err);
      alert("⚠️ Error de conexión con el servidor");
    }
  });
});
