-- name: InsertDataTrip :one
INSERT INTO data_trip (
  vehicle_prefix,
  latitude,
  longitude,
  shape_id,
  trip_id,
  trip_status,
  trip_progress,
  trip_delay,
  stop_distance,
  stop_order,
  report_time_diff
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;