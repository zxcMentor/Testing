-- 1_create_gtx_tables.up.sql
CREATE TABLE IF NOT EXISTS asks (
                                           id SERIAL PRIMARY KEY,
                                           timestamp INT,
                                           price  VARCHAR(100),
                                           volume VARCHAR(100),
                                           amount VARCHAR(100),
                                           factor VARCHAR(100),
                                           type   VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS bids (
                                           id SERIAL PRIMARY KEY,
                                           timestamp INT,
                                           price  VARCHAR(100),
                                           volume VARCHAR(100),
                                           amount VARCHAR(100),
                                           factor VARCHAR(100),
                                           type   VARCHAR(100)
);
