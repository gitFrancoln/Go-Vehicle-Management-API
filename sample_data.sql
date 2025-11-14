-- Script SQL para insertar datos de ejemplo en la tabla vehicles
-- VehículoMarket - Datos de ejemplo

-- Insertar vehículos de ejemplo
INSERT INTO vehicles (brand, model, year, price, mileage, transmission, fuel_type, color, body_type, status, description, created_at, updated_at) VALUES
('Mercedes-Benz', 'C-Class C 300', 2023, 48999.00, 15000, 'Automática', 'Gasolina', 'Negro', 'Sedán', 'Disponible', 'Mercedes-Benz C-Class 2023 en excelente estado. Equipado con tecnología de última generación, sistema de navegación, asientos de cuero y mucho más.', NOW(), NOW()),

('Ford', 'Mustang GT', 2023, 55000.00, 10000, 'Automática', 'Gasolina', 'Rojo', 'Coupé', 'Disponible', 'Ford Mustang GT 2023, motor V8 potente, diseño deportivo icónico. Perfecto para los amantes de la velocidad.', NOW(), NOW()),

('Honda', 'CR-V', 2023, 32000.00, 15000, 'Automática', 'Gasolina', 'Blanco', 'SUV', 'Disponible', 'Honda CR-V 2023, espacioso y confiable. Ideal para familias, con amplio espacio de carga y excelente consumo de combustible.', NOW(), NOW()),

('Toyota', 'Camry', 2023, 28000.00, 8000, 'Automática', 'Gasolina', 'Gris', 'Sedán', 'Disponible', 'Toyota Camry 2023, sedán familiar con excelente reputación de confiabilidad. Confortable y económico.', NOW(), NOW()),

('BMW', 'X5', 2023, 65000.00, 12000, 'Automática', 'Gasolina', 'Azul', 'SUV', 'Disponible', 'BMW X5 2023, lujo y rendimiento en un solo paquete. SUV premium con tecnología avanzada y acabados de alta calidad.', NOW(), NOW()),

('Volkswagen', 'Golf', 2023, 25000.00, 5000, 'Automática', 'Gasolina', 'Blanco', 'Hatchback', 'Disponible', 'Volkswagen Golf 2023, compacto y eficiente. Perfecto para la ciudad con excelente manejo y consumo de combustible.', NOW(), NOW()),

('Tesla', 'Model 3', 2023, 45000.00, 8000, 'Automática', 'Eléctrico', 'Blanco', 'Sedán', 'Disponible', 'Tesla Model 3 2023, vehículo eléctrico de alto rendimiento. Tecnología de autopilot, cero emisiones y bajo costo de mantenimiento.', NOW(), NOW()),

('Audi', 'A4', 2022, 42000.00, 18000, 'Automática', 'Gasolina', 'Negro', 'Sedán', 'Disponible', 'Audi A4 2022, sedán premium con diseño elegante. Equipado con sistema Quattro y tecnología MMI.', NOW(), NOW()),

('Chevrolet', 'Silverado', 2022, 48000.00, 25000, 'Automática', 'Gasolina', 'Gris', 'Pickup', 'Disponible', 'Chevrolet Silverado 2022, pickup robusta y confiable. Ideal para trabajo pesado y aventuras off-road.', NOW(), NOW()),

('Mazda', 'CX-5', 2023, 34000.00, 12000, 'Automática', 'Gasolina', 'Rojo', 'SUV', 'Disponible', 'Mazda CX-5 2023, SUV compacto con diseño deportivo. Excelente balance entre confort y desempeño.', NOW(), NOW()),

('Hyundai', 'Tucson', 2022, 29000.00, 20000, 'Automática', 'Gasolina', 'Blanco', 'SUV', 'Disponible', 'Hyundai Tucson 2022, SUV familiar con garantía extendida. Espacioso y equipado con características de seguridad avanzadas.', NOW(), NOW()),

('Nissan', 'Altima', 2021, 24000.00, 35000, 'Automática', 'Gasolina', 'Plateado', 'Sedán', 'Disponible', 'Nissan Altima 2021, sedán confiable con bajo kilometraje. Ideal como primer auto o vehículo familiar.', NOW(), NOW()),

('Jeep', 'Wrangler', 2023, 52000.00, 8000, 'Automática', 'Gasolina', 'Verde', 'SUV', 'Disponible', 'Jeep Wrangler 2023, icónico SUV off-road. Perfecto para aventuras y terrenos difíciles, con techo removible.', NOW(), NOW()),

('Subaru', 'Outback', 2022, 36000.00, 22000, 'Automática', 'Gasolina', 'Azul', 'SUV', 'Disponible', 'Subaru Outback 2022, SUV aventurero con tracción AWD estándar. Ideal para viajes largos y condiciones climáticas difíciles.', NOW(), NOW()),

('Kia', 'Sportage', 2023, 31000.00, 10000, 'Automática', 'Gasolina', 'Negro', 'SUV', 'Disponible', 'Kia Sportage 2023, SUV moderno con diseño renovado. Incluye garantía de fábrica y tecnología de conectividad avanzada.', NOW(), NOW()),

('Porsche', '911 Carrera', 2022, 115000.00, 5000, 'Automática', 'Gasolina', 'Amarillo', 'Coupé', 'Disponible', 'Porsche 911 Carrera 2022, icónico deportivo alemán. Performance excepcional y diseño atemporal.', NOW(), NOW()),

('Lexus', 'RX 350', 2023, 58000.00, 7000, 'Automática', 'Gasolina', 'Negro', 'SUV', 'Disponible', 'Lexus RX 350 2023, lujo japonés en su máxima expresión. Confort supremo y confiabilidad legendaria.', NOW(), NOW()),

('Volkswagen', 'Tiguan', 2022, 33000.00, 18000, 'Automática', 'Gasolina', 'Gris', 'SUV', 'Disponible', 'Volkswagen Tiguan 2022, SUV compacto versátil. Espacioso interior y tecnología de seguridad completa.', NOW(), NOW());

-- Verificar inserción
SELECT COUNT(*) as total_vehicles FROM vehicles;
