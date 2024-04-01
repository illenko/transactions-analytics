CREATE TABLE IF NOT EXISTS transactions
(
    id       uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    datetime timestamp      not null,
    amount   decimal(20, 2) not null,
    category varchar(256)   not null,
    merchant varchar(256)   not null
);

INSERT INTO transactions (id, datetime, amount, category, merchant)
VALUES ('44bdcdbc-4eae-443d-9bbd-4d1c1b7e628a', '2024-03-28 16:40:31.000000', -13.00, 'taxi', 'uklon'),
       ('64d3e47e-985b-46a0-8366-183693135ce2', '2024-03-27 12:51:29.846253', 20.00, 'p2p', 'stripe'),
       ('64d3e47e-985b-46a0-8366-183693135ceb', '2024-03-27 12:51:29.846000', 55.00, 'p2p', 'paypal'),
       ('cf59219d-5456-4513-a0ab-d6a05189ef98', '2024-01-27 12:45:19.854000', -30.00, 'taxi', 'uber'),
       ('c2a0c7cb-7b3e-4f45-b3a2-82be6afe7d77', '2024-01-27 12:50:44.190000', -24.00, 'house', 'jysk'),
       ('462fddc5-b0bd-4aa5-86cf-ac83272f09c8', '2024-01-27 12:50:01.878000', -100.00, 'food', 'glovo'),
       ('12a0c7cb-7b3e-4f45-b3a2-82be6afe7d77', '2024-01-28 12:51:29.846000', -5.00, 'p2p', 'paypal'),
       ('cd430ce8-92b5-40bb-97aa-ac6b77e6d168', '2024-02-28 13:44:51.000000', -20.05, 'food', 'mcdonalds'),
       ('859d3584-b715-4712-b158-d9dbd13f7366', '2024-02-28 13:45:41.000000', -33.00, 'taxi', 'uklon'),
       ('cd430ce8-92b5-40bb-97aa-3c6b77e6d168', '2024-03-28 13:44:51.000000', -3.20, 'food', 'kfc'),
       ('852d3584-b715-4712-b158-d9dbd13f7366', '2024-01-22 13:45:41.000000', -71.00, 'taxi', 'uklon'),
       ('6caa21c2-ea5c-4d08-ae5e-19860e10cbba', '2023-12-22 17:20:25.000000', -11.00, 'taxi', 'uklon'),
       ('b1b7a396-46f8-4d0b-8c9c-dd4d9e28512f', '2023-11-03 17:21:11.000000', -33.00, 'taxi', 'uklon'),
       ('a507fd0e-89fc-45ab-81ad-c8cf2308a133', '2023-10-13 17:21:44.000000', -11.00, 'taxi', 'uklon'),
       ('fe7976ad-578c-442e-b6fb-33b69f9ec029', '2024-03-30 08:37:17.000000', 10.00, 'p2p', 'revolut'),
       ('64d3e47e-985b-46a0-8366-183693135cef', '2024-03-28 12:51:29.846000', 104.00, 'work', 'upwork');
