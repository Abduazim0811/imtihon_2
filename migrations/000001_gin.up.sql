CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS authors (
    author_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    birth_date DATE,
    biography TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS books (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id UUID,
    FOREIGN KEY (author_id) REFERENCES authors(author_id),
    publication_date DATE,
    isbn VARCHAR(20),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INSERT INTO Authors (name, birth_date, biography) VALUES
-- ('Abdulla Qodiriy', '1894-04-10', 'Oʻzbek yozuvchisi, dramaturg, shoir va jamoat arbobi.'),
-- ('Cholpon', '1893-03-06', 'Oʻzbek shoiri, yozuvchisi, tarjimoni va jamoat arbobi.'),
-- ('Abdulla Oripov', '1941-03-21', 'Oʻzbek shoiri va Oʻzbekiston Qahramoni.'),
-- ('Erkin Vohidov', '1936-12-28', 'Oʻzbek shoiri, dramaturg, va jamoat arbobi.'),
-- ('Hamid Olimjon', '1909-12-12', 'Oʻzbek shoiri, dramaturg va jamoat arbobi.'),
-- ('Zulfiya', '1915-03-01', 'Oʻzbek shoira va jurnalisti.'),
-- ('Asqad Muxtor', '1920-02-22', 'Oʻzbek yozuvchisi, dramaturg va jamoat arbobi.'),
-- ('Mirtemir', '1910-05-30', 'Oʻzbek shoiri, yozuvchisi va tarjimoni.'),
-- ('Said Ahmad', '1920-06-10', 'Oʻzbek yozuvchisi va dramaturgi.'),
-- ('Utkir Hoshimov', '1941-08-05', 'Oʻzbek yozuvchisi va jurnalisti.');

-- INSERT INTO Books (title, author_id, publication_date, isbn, description) VALUES
-- ('Oʻtgan kunlar', (SELECT author_id FROM Authors WHERE name = 'Abdulla Qodiriy'), '1926-01-01', '978-9943-472-72-7', 'Oʻzbek adabiyotining klassik romani.'),
-- ('Kecha va kunduz', (SELECT author_id FROM Authors WHERE name = 'Abdulla Qodiriy'), '1936-01-01', '978-9943-472-73-4', 'Oʻzbek xalqining hayotini tasvirlaydigan roman.'),
-- ('Sarvqomat dilbarim', (SELECT author_id FROM Authors WHERE name = 'Cholpon'), '1923-01-01', '978-9943-472-74-1', 'Oʻzbek sheʼriyati namunasi.'),
-- ('Bahor keldi seni soʻroqlab', (SELECT author_id FROM Authors WHERE name = 'Cholpon'), '1925-01-01', '978-9943-472-75-8', 'Oʻzbek sheʼriyati namunasi.'),
-- ('Sadoqat', (SELECT author_id FROM Authors WHERE name = 'Abdulla Oripov'), '1965-01-01', '978-9943-472-76-5', 'Oʻzbek sheʼriyati va dostonlar toʻplami.'),
-- ('Yurtim shamoli', (SELECT author_id FROM Authors WHERE name = 'Abdulla Oripov'), '1971-01-01', '978-9943-472-77-2', 'Oʻzbek sheʼriyati va dostonlar toʻplami.'),
-- ('Ruhlar isyoni', (SELECT author_id FROM Authors WHERE name = 'Erkin Vohidov'), '1960-01-01', '978-9943-472-78-9', 'Oʻzbek sheʼriyati va dostonlar toʻplami.'),
-- ('Qoʻshiqlarim sizga', (SELECT author_id FROM Authors WHERE name = 'Erkin Vohidov'), '1964-01-01', '978-9943-472-79-6', 'Oʻzbek sheʼriyati va dostonlar toʻplami.'),
-- ('Oygul bilan Baxtiyor', (SELECT author_id FROM Authors WHERE name = 'Hamid Olimjon'), '1937-01-01', '978-9943-472-80-2', 'Oʻzbek poeziyasining klassik asari.'),
-- ('Zaynab va Omon', (SELECT author_id FROM Authors WHERE name = 'Hamid Olimjon'), '1938-01-01', '978-9943-472-81-9', 'Oʻzbek poeziyasining klassik asari.');
