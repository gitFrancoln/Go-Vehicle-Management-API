// Navegación entre secciones
document.querySelectorAll('.menu-item').forEach(item => {
    item.addEventListener('click', function() {
        // Remover clase active de todos los items
        document.querySelectorAll('.menu-item').forEach(i => {
            i.classList.remove('active');
        });
        
        // Agregar clase active al item clickeado
        this.classList.add('active');
        
        // Ocultar todas las secciones
        document.querySelectorAll('.section').forEach(section => {
            section.classList.remove('active');
        });
        
        // Mostrar la sección correspondiente
        const sectionId = this.getAttribute('data-section');
        document.getElementById(sectionId).classList.add('active');
    });
});

// Subir imagen de perfil
const uploadBtn = document.getElementById('upload-btn');
const fileInput = document.getElementById('file-input');
const profileImg = document.getElementById('profile-img');

uploadBtn.addEventListener('click', function() {
    fileInput.click();
});

fileInput.addEventListener('change', function() {
    const file = this.files[0];
    if (file) {
        const reader = new FileReader();
        
        reader.addEventListener('load', function() {
            profileImg.src = reader.result;
        });
        
        reader.readAsDataURL(file);
    }
});

// Guardar cambios del perfil
document.getElementById('save-profile').addEventListener('click', function() {
    const firstName = document.getElementById('first-name').value;
    const lastName = document.getElementById('last-name').value;
    const email = document.getElementById('email').value;
    
    document.getElementById('user-name').textContent = `${firstName} ${lastName}`;
    document.getElementById('user-email').textContent = email;
    
    alert('Cambios guardados correctamente');
});

// Cerrar sesión
document.querySelector('.logout').addEventListener('click', function(e) {
    e.preventDefault();
    if (confirm('¿Estás seguro de que quieres cerrar sesión?')) {
        alert('Sesión cerrada correctamente');
        // Aquí iría la lógica real para cerrar sesión
        // window.location.href = '/login';
    }
});

// Ir a pantalla de suscripciones
document.getElementById('go-to-subscriptions').addEventListener('click', function() {
    alert('Redirigiendo a la pantalla de planes y suscripciones...');
    // Aquí iría la redirección real
    // window.location.href = '/subscriptions';
});

// Modo oscuro
function initDarkMode() {
    // Crear botón de toggle
    const themeToggle = document.createElement('div');
    themeToggle.className = 'theme-toggle';
    themeToggle.innerHTML = `
        <button class="theme-toggle-btn">
            <span class="theme-icon">🌙</span>
            <span class="theme-text">Modo Oscuro</span>
        </button>
    `;
    document.body.appendChild(themeToggle);
    
    // Verificar preferencia guardada
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'dark') {
        document.body.classList.add('dark-mode');
        updateThemeToggle('dark');
    }
    
    // Event listener para el toggle
    themeToggle.addEventListener('click', function() {
        document.body.classList.toggle('dark-mode');
        
        if (document.body.classList.contains('dark-mode')) {
            localStorage.setItem('theme', 'dark');
            updateThemeToggle('dark');
        } else {
            localStorage.setItem('theme', 'light');
            updateThemeToggle('light');
        }
    });
    
    function updateThemeToggle(theme) {
        const themeIcon = themeToggle.querySelector('.theme-icon');
        const themeText = themeToggle.querySelector('.theme-text');
        
        if (theme === 'dark') {
            themeIcon.textContent = '☀️';
            themeText.textContent = 'Modo Claro';
        } else {
            themeIcon.textContent = '🌙';
            themeText.textContent = 'Modo Oscuro';
        }
    }
}

// Inicializar la imagen de perfil por defecto
window.addEventListener('DOMContentLoaded', function() {
    // Crear una imagen de perfil por defecto si no hay ninguna
    if (!profileImg.src) {
        // Crear un canvas para generar una imagen de perfil por defecto
        const canvas = document.createElement('canvas');
        canvas.width = 90;
        canvas.height = 90;
        const ctx = canvas.getContext('2d');
        
        // Fondo de color
        ctx.fillStyle = '#1a3d66';
        ctx.fillRect(0, 0, 90, 90);
        
        // Texto con iniciales
        ctx.fillStyle = 'white';
        ctx.font = 'bold 32px Arial';
        ctx.textAlign = 'center';
        ctx.textBaseline = 'middle';
        ctx.fillText('MF', 45, 45);
        
        // Establecer la imagen generada como src
        profileImg.src = canvas.toDataURL();
    }
    
    // Inicializar modo oscuro
    initDarkMode();
});