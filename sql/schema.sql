CREATE TABLE data_trip (
  id BIGSERIAL PRIMARY KEY,
  vehicle_prefix text NOT NULL,
  latitude real NOT NULL,
  longitude real NOT NULL,
  shape_id text NOT NULL,
  trip_id text NOT NULL,
  trip_status text NOT NULL,
  trip_progress smallint NOT NULL,
  trip_delay smallint NOT NULL,
  stop_distance real NOT NULL,
  stop_order smallint NOT NULL,
  report_time_diff smallint NOT NULL
);