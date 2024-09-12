package flight

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type flightHandler struct {
	db *sql.DB
}

func NewflightHandler(db *sql.DB) flightHandler {
	return flightHandler{db: db}
}

// UPDATE
func (h flightHandler) UpdateFlightHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	flight := Flight{ID: id, Number: 123456, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"}

	data, _ := UpdateFlight(h.db, flight)
	fmt.Println("Update Flight : ", data)
	c.JSON(200, data)
}

func UpdateFlight(db *sql.DB, f Flight) (sql.Result, error) {
	data, err := db.Exec(`UPDATE flights SET number = $1, airlineCode = $2, destination = $3, arrival = $4 WHERE id = $5;`, f.Number, f.AirlineCode, f.Destination, f.Arrival, f.ID)
	return data, err
}

// INSERT
func (h flightHandler) CreateFlightHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	flight := Flight{ID: id, Number: 3250, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"}

	data, _ := CreateFlight(h.db, flight)
	fmt.Println("Create Flight: ", data)
	c.JSON(200, data)
}

func CreateFlight(db *sql.DB, f Flight) (Flight, error) {
	err := db.QueryRow(`INSERT INTO flights(id, number, airlineCode, destination, arrival) VALUES($1, $2, $3, $4, $5) RETURNING id;`, f.ID, f.Number, f.AirlineCode, f.Destination, f.Arrival).Scan(&f.ID)
	if err != nil {
		return Flight{}, err
	}
	return f, nil
}

// DELETE
func (h flightHandler) DeleteFlightByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	data := DeleteFlightByID(h.db, id)
	if data != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't delete",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func DeleteFlightByID(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM flights WHERE id = $1;`, id)
	return err
}

// GET BY ID
func (h flightHandler) GetByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	data, _ := GetFlightByID(h.db, id)
	fmt.Println("Get Flight by id : ", data)
	c.JSON(200, data)

}

func GetFlightByID(db *sql.DB, id int) (Flight, error) {
	var flight Flight
	row := db.QueryRow(`SELECT id, number, airlineCode, destination, arrival FROM flights WHERE id = $1;`, id)
	err := row.Scan(&flight.ID, &flight.Number, &flight.AirlineCode, &flight.Destination, &flight.Arrival)
	if err != nil {
		return Flight{}, err
	}
	return flight, nil
}

// GET All
func (h flightHandler) GetAllHandler(c *gin.Context) {
	data, _ := GetAllFlights(h.db)
	fmt.Println("Get All Flight: ", data)
	c.JSON(200, data)
}

func GetAllFlights(db *sql.DB) ([]Flight, error) {
	rows, err := db.Query("SELECT id, number, airlineCode, destination, arrival FROM flights")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []Flight
	for rows.Next() {
		var f Flight
		err := rows.Scan(&f.ID, &f.Number, &f.AirlineCode, &f.Destination, &f.Arrival)
		if err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}
