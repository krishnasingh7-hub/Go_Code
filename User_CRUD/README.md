# GETTING ERROR 500 WHILE HITTING API
# DB CONNECTION AND COMPILATION IS SUCCESSFUL

 STEPS to RUN :
 1. open cmd
 2. // Then login to DB using command
        mysql -u root -p
 3. create database user
     // change to user db
    use user
 4. DROP TABLE IF EXISTS user;
 5. CREATE TABLE user(
    id         INT AUTO_INCREMENT NOT NULL,
    name      VARCHAR(128) NOT NULL,
    age      INT NOT NULL,
    PRIMARY KEY (`id`)
    );

6.  INSERT INTO user
    (name, age)
    VALUES
    ('Ana Grey', 20),
    ('Jeru Steps',  22),
    ('Charlice Honey',  25),
    ('Harry Potter', 24);
 7. select * from user;
 # Issue might be id not passed in insert command and remove auto increament

 END POINTS
 GET http://localhost:8080/users