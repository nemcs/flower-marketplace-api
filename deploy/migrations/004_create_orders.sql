CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES clients(id),
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    courier_id INTEGER NOT NULL REFERENCES couriers(id),
    address TEXT NOT NULL,
    status TEXT NOT NULL
);
