package swagger

import(
	"log"
	"os"
	"fmt"
	"encoding/csv"
)


/**
* Flightes collection
*/
var flights = []Flight{}

/**
* Error handler
* @parm message: error message
* @parm err: error 
**/
func checkFlightError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}


/**
* find a Flight index by id
* @id: Flight Id to find
* @return: Flight's index or -1 if not found
*/
func findFlightIndex(id string)int{
	for index, flight := range flights {
        if id == flight.FlightId {
            return index
        }
    }
    return -1
}


/*
* Retursn Flight pointer by id
* @parm id: Crash Id to find
* @return Flight pointer or nil if not found
*
*/
func getFlightByCrashId(id string)*Flight{
	for _, flight := range flights {
        if id == flight.CrashId {
            return &flight
        }
    }
    return nil
}



/*
* unlink a flight to a crash
* @parm flightId: fligth id
* return true if ok, false, if flight doesn't exist
*
*/
func unlinkFlight(flightId string) bool{
	var flight = getFlightById(flightId)
	if flight == nil{
		return false;
	}
	flight.CrashId = "-";
	return updateFlight(flight.FlightId,*flight)
}

/*
* Retursn flight pointer by id
* @parm id: Flight Id to find
* @return Flight pointer or nil if not found
*
*/
func getFlightById(id string)*Flight{
	var index = findFlightIndex(id)
	if(index == -1){
		return nil
	}
	return &flights[index];
}

/*
* Add a Flight to the collection
* @parm flight: Flight to add
* @return true if added, false if repeted record
*
*/
func addFlight(flight Flight)bool{
	var aux = findFlightIndex(flight.FlightId);
	if(aux != -1){
		log.Printf("Flight "+flight.FlightId+" not added")
		return false;
	}
	flights = append(flights,flight)
	//log.Printf("Flight "+flight.FlightId+" added")
	return true;
}

/*
* return flight Collection
* @return flight collection
*/
func allFlights()[]Flight{
	return flights
}


/*
* Remove a Flight from the collection
* @parm flight: Flight to remove
* @return true if removed, false if not found record
*
*/
func removeFlight(id string)bool{
	var index = findFlightIndex(id);
	if(index == -1){
		log.Printf("Flight "+id+" not deleted")
		return false;
	}
	log.Printf("Flight "+id+" deleted")
	flights = append(flights[:index], flights[index+1:]...)
	return true;
}


/*
* Update a Flight from the collection
* @parm flight: Flight to update
* @return true if updated, false if not found record
*/
func updateFlight(id string,flight Flight)bool{
	var index = findFlightIndex(id);
	var index2 = findFlightIndex(flight.FlightId);
	if(index == -1||( index2 == -1 && id != flight.FlightId)){
		log.Printf("Flight "+flight.FlightId+" not updated")
		return false;
	}
	log.Printf("Flight "+flight.FlightId+" updated")
	flights[index] = flight;
	return true;
}

/**
* Read initial dataset of flights
*/
func LoadFlightData(){
	fmt.Println("-------------------leyendo registros de vuelos--------------")
	filePath := "flight.csv"
	file, err1 := os.Open(filePath)
	checkFlightError("Unable to read input file "+filePath, err1)
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkFlightError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	for _, record := range records {
		res := addFlight(Flight{
			FlightId      : record[0],
			CrashId       : record[1],
			Operator      : record[2],
			Flight        : record[3],
			Route         : record[4],
			Type_         : record[5],
			Registration  : record[6]})
		if res == false {
			fmt.Println("Record " +  record[0] + " repeted")
		}
	}
	file.Close()
}

