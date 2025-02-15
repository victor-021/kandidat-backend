CREATE TABLE Users(
    user_id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    picture bytea,
    rating float4,
    business BOOLEAN NOT NULL
);

CREATE TABLE User_Followers(
    user_followers_id SERIAL PRIMARY KEY,
    fk_user_id INT REFERENCES Users(user_id) NOT NULL,
    fk_followed_id INT REFERENCES Users(user_id) NOT NULL
);

CREATE TABLE Product (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    service BOOLEAN NOT NULL,
    price INT NOT NULL,
    upload_date DATE NOT NULL DEFAULT CURRENT_DATE,
    description VARCHAR,
    picture bytea,
    category VARCHAR,
    fk_user_id INT REFERENCES Users(user_id) NOT NULL,
    fk_buyer_id INT REFERENCES Users(user_id)
);

CREATE TABLE Review (
    review_id SERIAL PRIMARY KEY,
    rating INT NOT NULL,
    content VARCHAR,
    fk_reviewer_id INT REFERENCES Users(user_id) ON UPDATE CASCADE  NOT NULL,
    fk_owner_id INT REFERENCES Users(user_id) ON UPDATE CASCADE NOT NULL

);

CREATE TABLE Pinned_Product (
    fk_product_id INT REFERENCES Product(product_id) ON DELETE CASCADE NOT NULL,
    fk_user_id INT REFERENCES Users(user_id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(fk_product_id, fk_user_id)
);

CREATE TABLE Buying_Product (
    fk_product_id INT REFERENCES Product(product_id) ON DELETE CASCADE NOT NULL,
    fk_user_id INT REFERENCES Users(user_id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(fk_product_id, fk_user_id)
);

CREATE TABLE Chats (
    fk_user_id_1 INT REFERENCES Users(user_id) ON DELETE CASCADE NOT NULL,
    fk_user_id_2 INT REFERENCES Users(user_id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(fk_user_id_1, fk_user_id_2)
);

CREATE TABLE Community (
    community_id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);


CREATE TABLE User_Community (
    user_community_id SERIAL PRIMARY KEY,
    fk_user_id INT REFERENCES Users(user_id) NOT NULL,
    fk_community_id INT REFERENCES Community(community_id) NOT NULL
);


/* test users user_id = 1 & 2 */
INSERT INTO Users (name, phone_number, password, picture, rating, business) VALUES ('Gustav', '+12029182132', '$2a$12$IDEtMuDeOB/m4e.BVwEJ0O/FdUXKNF3sq8BnNHFIQpdf8h/NJCJHi', encode(pg_read_binary_file('/docker-entrypoint-initdb.d/victorkill.jpeg'), 'base64')::bytea, 3,'true');

INSERT INTO USERS (name, phone_number, password, rating,business) VALUES ('Victor', '+12027455483', '$2a$12$IDEtMuDeOB/m4e.BVwEJ0O/FdUXKNF3sq8BnNHFIQpdf8h/NJCJHi', 4,'true');

/* test products product_id = 1 */
INSERT INTO Product (name,service,price,description, fk_user_id ) VALUES ('Soffa','true',1,'Hej',1);
/* test products product_id = 1 & 2*/
INSERT INTO Product (name,service,price,description, fk_user_id ) VALUES ('Couch','true',1,'Couch description',1);

INSERT INTO Product (name,service,price,description, fk_user_id ) VALUES ('Bed','true',1,'Bed description',1);

/* test products product_id = 3 for user_id 2 */
INSERT INTO Product (name,service,price,description, fk_user_id ) VALUES ('Car','true',1,'Car description',2);

/* test review review_id = 1 */
INSERT INTO Review (rating,content, fk_reviewer_id, fk_owner_id) VALUES (2,'SÄMST',1,2);

/* test communities community_id = 1 & 2 & 3 */
INSERT INTO Community (name) VALUES ('Clothes'), ('Politics'), ('Memes');

/* test pinned_product pinnedproduct_id = 1 */
INSERT INTO Pinned_Product (fk_product_id, fk_user_id) VALUES (1,1);

INSERT INTO User_Community(fk_user_id, fk_community_id) VALUES (1,2);

/* test user_followers user_follower_id = 1 */
INSERT INTO User_Followers(fk_user_id, fk_followed_id) VALUES (2, 1);

/* test buying_product */
INSERT INTO Buying_Product (fk_product_id, fk_user_id) VALUES (1,2);

/* test chats */
INSERT INTO Chats (fk_user_id_1, fk_user_id_2) VALUES (1,2);
