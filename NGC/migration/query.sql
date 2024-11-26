-- DDL (Data Definition Language) to create db tables
-- drop existing tables if they exists from the db
DROP TABLE IF EXISTS Heroes;
DROP TABLE IF EXISTS Villain;

-- Create Heroes Table
CREATE TABLE Heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255) NOT NULL,
    Skill VARCHAR(255),
    ImageURL VARCHAR(2083)
);

-- Create Villain Table
CREATE TABLE Villain (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(2083)
);

-- Insert data into Heroes
INSERT INTO Heroes (Name, Universe, Skill, ImageURL) VALUES
('Iron Man', 'Marvel', 'Technology', 'http://example.com/ironman.jpg'),
('Thor', 'Marvel', 'God of Thunder', 'http://example.com/thor.jpg'),
('Batman', 'DC', 'Detective Skills', 'http://example.com/batman.jpg');

-- Insert data into Villain
INSERT INTO Villain (Name, Universe, ImageURL) VALUES
('Loki', 'Marvel', 'http://example.com/loki.jpg'),
('Joker', 'DC', 'http://example.com/joker.jpg'),
('Thanos', 'Marvel', 'http://example.com/thanos.jpg');