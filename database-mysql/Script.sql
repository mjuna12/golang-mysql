-- 1. CREATE TABLE USERS & INSERT DATA
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (name, email, password, created_at, updated_at)
VALUES
('Budi', 'budi@budi.com', '67890', '2019-01-28 05:15:29', '2019-01-28 05:15:29'),
('Caca', 'caca@caca.com', 'abcde', '2019-01-28 05:15:29', '2019-01-28 05:15:29'),
('Deni', 'deni@deni.com', 'fghij', '2019-01-28 05:15:29', '2019-01-28 05:15:29'),
('Euis', 'euis@euis.com', 'klmno', '2019-01-28 05:15:29', '2019-01-28 05:15:29'),
('Fafa', 'fafa@fafa.com', 'pqrst', '2019-01-28 05:15:29', '2019-01-28 05:15:29');

-- 2. CREATE TABLE COURSES & INSERT DATA
CREATE TABLE courses (
    id INT PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    mentor_name VARCHAR(100) NOT NULL,
    mentor_title VARCHAR(50) NOT NULL
);

INSERT INTO courses (course_name, mentor_name, mentor_title)
VALUES
('C++', 'Ari', 'Dr.'),
('C#', 'Ari', 'Dr.'),
('C#', 'Ari', 'Dr.'),
('CSS', 'Cania', 'S.Kom'),
('HTML', 'Cania', 'S.Kom'),
('Javascript', 'Cania', 'S.Kom'),
('Python', 'Cania', 'S.Kom'),
('Micropython', 'Barry', 'S.T.'),
('Java', 'Darren', 'M.T.'),
('Ruby', 'Darren', 'M.T.');

-- 3. CREATE TABLE USER_COURSE & INSERT DATA
CREATE TABLE user_course (
    id_user INT,
    id_course INT,
    PRIMARY KEY (id_user, id_course),
    FOREIGN KEY (id_user) REFERENCES users(id),
    FOREIGN KEY (id_course) REFERENCES courses(id)
);

INSERT INTO user_course (id_user, id_course)
VALUES
(1, 1),
(1, 2),
(3, 7),
(3, 8),
(3, 9),
(3, 1),
(4, 2),
(4, 3),
(4, 4),
(5, 5),
(6, 6),
(6, 7),
(6, 8),
(6, 9),
(2, 3),
(2, 4),
(2, 5),
(3, 6);

-- 4. Menampilkan semua daftar peserta didik dan mata kuliah yang diikuti
SELECT 
    u.name AS peserta_didik, 
    c.course_name AS mata_kuliah, 
    c.mentor_name AS mentor, 
    c.mentor_title AS gelar_mentor
FROM 
    user_course AS uc
JOIN 
    users AS u ON uc.id_user = u.id
JOIN 
    courses AS c ON uc.id_course = c.id;

-- 5. Menampilkan Daftar Peserta Didik dan Mata Kuliah yang Diikutinya dengan Mentor Bergelar Sarjana
SELECT 
    u.name AS peserta_didik, 
    c.course_name AS mata_kuliah, 
    c.mentor_name AS mentor, 
    c.mentor_title AS gelar_mentor
FROM 
    user_course AS uc
JOIN 
    users AS u ON uc.id_user = u.id
JOIN 
    courses AS c ON uc.id_course = c.id
WHERE 
    c.mentor_title IN ('S.Kom', 'S.T.');

-- 6. Daftar peserta didik beserta mata kuliah yang diikutinya, yang mentornya bergelar selain sarjana
SELECT 
    u.name AS peserta_didik, 
    c.course_name AS mata_kuliah, 
    c.mentor_name AS mentor, 
    c.mentor_title AS gelar_mentor
FROM 
    user_course AS uc
JOIN 
    users AS u ON uc.id_user = u.id
JOIN 
    courses AS c ON uc.id_course = c.id
WHERE 
    c.mentor_title NOT IN ('S.Kom', 'S.T.');

-- 7. Menampilkan Jumlah Peserta Didik per Mata Kuliah
SELECT 
    c.course_name AS mata_kuliah, 
    COUNT(uc.id_user) AS jumlah_peserta_didik
FROM 
    user_course AS uc
JOIN 
    courses AS c ON uc.id_course = c.id
GROUP BY 
    c.course_name;

-- 8. Menampilkan jumlah peserta didik dan total fee untuk setiap mentor. Total fee dihitung dengan besaran Rp 2.000.000,- per peserta didik
SELECT 
    c.mentor_name AS mentor, 
    c.mentor_title AS gelar_mentor, 
    COUNT(uc.id_user) AS jumlah_peserta_didik,
    COUNT(uc.id_user) * 2000000 AS total_fee
FROM 
    user_course AS uc
JOIN 
    courses AS c ON uc.id_course = c.id
GROUP BY 
    c.mentor_name, c.mentor_title;
