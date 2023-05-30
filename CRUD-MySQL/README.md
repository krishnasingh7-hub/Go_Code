Reference Link : https://www.golinuxcloud.com/golang-sql/
 STEPS to RUN :
 1. open cmd
 2. // Then login to DB using command
        mysql -u root -p
 3. create database students
     // change to students db
    use students
 4. DROP TABLE IF EXISTS student;
 5. CREATE TABLE student(
    id         INT AUTO_INCREMENT NOT NULL,
    name      VARCHAR(128) NOT NULL,
    email     VARCHAR(255) NOT NULL,
    age      INT NOT NULL,
    PRIMARY KEY (`id`)
    );
 6. INSERT INTO student
    (name, email, age)
    VALUES
    ('Ana Grey', 'email1@gmail.com', 20),
    ('Jeru Steps', 'email1@gmail.com', 22),
    ('Charlice Honey', 'email1@gmail.com', 25),
    ('Harry Potter', 'email1@gmail.com', 24);
 7. select * from student;