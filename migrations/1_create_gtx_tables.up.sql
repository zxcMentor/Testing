-- 1_create_gtx_tables.up.sql
CREATE TABLE IF NOT EXISTS public.asks (
                                           id SERIAL PRIMARY KEY,
                                           timestamp INT,
                                           price  VARCHAR(100),
                                           volume VARCHAR(100),
                                           amount VARCHAR(100),
                                           factor VARCHAR(100),
                                           type   VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS public.bids (
                                           id SERIAL PRIMARY KEY,
                                           timestamp INT,
                                           price  VARCHAR(100),
                                           volume VARCHAR(100),
                                           amount VARCHAR(100),
                                           factor VARCHAR(100),
                                           type   VARCHAR(100)
);

-- 1_create_gtx_tables.down.sql

DROP TABLE public.asks;
DROP TABLE public.bids;