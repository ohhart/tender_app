-- Создание типа для типа организации
CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

-- Создание таблицы для сотрудников
CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы для организаций
CREATE TABLE organization (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы для ответственных лиц организаций
CREATE TABLE organization_responsible (
    id SERIAL PRIMARY KEY,
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
    user_id INT REFERENCES employee(id) ON DELETE CASCADE
);

-- Создание типа для статуса тендера
CREATE TYPE tender_status AS ENUM (
    'CREATED',
    'PUBLISHED',
    'CLOSED'
);

-- Создание таблицы для тендеров
CREATE TABLE tenders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type VARCHAR(255), -- тип услуги
    status tender_status DEFAULT 'CREATED',
    version INT DEFAULT 1,
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
    creator_id INT REFERENCES employee(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание типа для статуса предложения
CREATE TYPE bid_status AS ENUM (
    'CREATED',
    'PUBLISHED',
    'CANCELED'
);

-- Создание таблицы для предложений
CREATE TABLE bids (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status bid_status DEFAULT 'CREATED',
    tender_id INT REFERENCES tenders(id) ON DELETE CASCADE,
    author_id INT REFERENCES employee(id) ON DELETE CASCADE,
    version INT DEFAULT 1,
    decision VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы для отзывов
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    bid_id INT REFERENCES bids(id) ON DELETE CASCADE,
    tender_id INT REFERENCES tenders(id) ON DELETE CASCADE,
    author_username VARCHAR(50) REFERENCES employee(username),
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE, 
    comment TEXT,
    rating INT CHECK (rating >= 1 AND rating <= 5), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Вставка данных в таблицу employee
INSERT INTO employee (username, first_name, last_name) VALUES
('asmith', 'Alice', 'Smith'),
('jdoe', 'John', 'Doe'),
('mjones', 'Mary', 'Jones'),
('bwilliams', 'Bob', 'Williams'),
('kthomas', 'Karen', 'Thomas');

-- Вставка данных в таблицу organization
INSERT INTO organization (name, description, type) VALUES
('Tech Innovations LLC', 'A technology company focusing on software solutions', 'LLC'),
('Global Enterprises JSC', 'A multinational conglomerate', 'JSC'),
('Smith Consulting IE', 'Consulting services for small businesses', 'IE'),
('Green Energy LLC', 'Renewable energy solutions provider', 'LLC'),
('HealthCare Solutions JSC', 'Innovative healthcare services', 'JSC');

-- Вставка данных в таблицу organization_responsible
INSERT INTO organization_responsible (organization_id, user_id) VALUES
(1, 1), -- Alice Smith ответственная за Tech Innovations LLC
(2, 2), -- John Doe ответственный за Global Enterprises JSC
(3, 3), -- Mary Jones ответственная за Smith Consulting IE
(4, 4), -- Bob Williams ответственный за Green Energy LLC
(5, 5); -- Karen Thomas ответственная за HealthCare Solutions JSC
