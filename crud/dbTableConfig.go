package crud

const baseAirport = "base_airport"
const targetAirport = "target_airport"
const baseDepart = "base_depart"
const targetArrival = "target_arrival"
const baseAirline = "base_airline"
const targetAirline = "target_airline"
const price = "price"


const getAirlineListWithWhereName =
    "SELECT id FROM airline WHERE name=?"
const getAirportListWithWhereCode =
    "SELECT id FROM airport WHERE code=?"
const insertTicketData =
`
INSERT INTO round_trip_ticket_price
(base_airport_id, target_airport_id,
 airline1_id, airline2_id,
 base_airport_depart_time, target_airport_arrival_time,
 price,
 insert_time, last_update_time)
VALUES
(
 ?, ?,
 ?, ?,
 ?, ?,
 ?,
 NOW(), NOW()
)
`

const getAllTicketDataSQL = `
SELECT names1.airport_name base_airport_name, names1.nation_name base_nation_name,
       names2.airport_name target_airport_name, names2.nation_name target_nation_name,
       airline1.name, airline2.name,
       price1.base_airport_depart_time,
       price1.target_airport_arrival_time,
       price1.price, count(*) OVER(),
       price1.target_airport_depart_time,
       price1.base_airport_arrival_time
FROM round_trip_ticket_price price1
    JOIN (SELECT nation.name nation_name, airport.name airport_name, airport.id airport_id
            FROM airport
                JOIN nation
                ON airport.nation_id = nation.id) names1
        ON price1.base_airport_id = names1.airport_id
    JOIN (SELECT nation.name nation_name, airport.name airport_name, airport.id airport_id
            FROM airport
                JOIN nation
                ON airport.nation_id = nation.id) names2
        ON price1.target_airport_id = names2.airport_id
    JOIN airline airline1
        ON price1.airline1_id = airline1.id
    JOIN airline airline2
        ON price1.airline2_id = airline2.id
`
