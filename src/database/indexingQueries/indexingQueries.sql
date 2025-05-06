CREATE INDEX user_index ON users (username);
CREATE INDEX devices_index ON devices (device_id, username);
CREATE INDEX broker_index ON brokers (broker_name);
CREATE INDEX staticexchangeconfig_index ON static_exchange_config (exchange, price_type, adv_ret_type_licensed);