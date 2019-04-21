package crud

const planeTicketDB = "plane_ticket"
const dbTableConnector = "."
const prePlaneTicketDB = planeTicketDB + dbTableConnector

const ticketPriceTableInfo = prePlaneTicketDB + "round_trip_ticket_price"
const nationTableInfo = prePlaneTicketDB + "nation"
const airportTableInfo = prePlaneTicketDB + "airport"

const baseAirportId = "base_airport_id"
const targetAirportId = "target_airport_id"
const airportName = "airportName"

const getAllTicketDataSQL = `
SELECT names1.airport_name base_airport_name, names1.nation_name base_nation_name,
       names2.airport_name target_airport_name, names2.nation_name target_nation_name,
       airline1.name, airline2.name,
       price1.base_airport_depart_time,
       price1.target_airport_arrival_time,
       price1.target_airport_depart_time,
       price1.base_airport_arrival_time,
       price1.price, count(*) OVER()
FROM plane_ticket.round_trip_ticket_price price1
    JOIN (SELECT nation.name nation_name, airport.name airport_name, airport.id airport_id
            FROM plane_ticket.airport
                JOIN plane_ticket.nation
                ON airport.nation_id = nation.id) names1
        ON price1.base_airport_id = names1.airport_id
    JOIN (SELECT nation.name nation_name, airport.name airport_name, airport.id airport_id
            FROM plane_ticket.airport
                JOIN plane_ticket.nation
                ON airport.nation_id = nation.id) names2
        ON price1.target_airport_id = names2.airport_id
    JOIN plane_ticket.airline airline1
        ON price1.airline1_id = airline1.id
    JOIN plane_ticket.airline airline2
        ON price1.airline2_id = airline2.id
`
