package swagger

import(
	"log"
	"os"
	"fmt"
	"encoding/csv"
)


/**
* Crashes collection
*/
var crashes = []Crash{}

/**
* Error handler
* @parm message: error message
* @parm err: error 
**/
func checkCrashError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

/*
* link a flight to a crash
* @parm crashId: crash id
* @parm flightId: fligth id
* return true if ok, false, if someone doesn't exist
*
*/
func linkFlight(crashId string,flightId string) bool{
	var crash = getCrashById(crashId)
	var flight = getFlightById(flightId)
	if crash == nil || flight == nil{
		return false;
	}
	flight.CrashId = crash.CrashId;
	return updateFlight(flight.FlightId,*flight)
}

/*
* link a flight to a crash
* @parm crashId: crash id
* @parm flightId: fligth id
* return true if ok, false, if someone doesn't exist
*
*/
func linkStatistic(crashId string,statisticId string) bool{
	var crash = getCrashById(crashId)
	var statistic = getStatisticById(statisticId)
	if crash == nil || statistic == nil{
		return false;
	}
	statistic.CrashId = crash.CrashId;
	return updateStatistic(statistic.StatisticId,*statistic)
}


/**
* find a Crash index by id
* @id: Crash Id to find
* @return: Crash's index or -1 if not found
*/
func findCrashIndex(id string)int{
	for index, crash := range crashes {
        if id == crash.CrashId {
            return index
        }
    }
    return -1
}

/*
* Retursn crash pointer by id
* @parm id: Crash Id to find
* @return Crash pointer or nil if not found
*
*/
func getCrashById(id string)*Crash{
	var index = findCrashIndex(id)
	if(index == -1){
		return nil
	}
	return &crashes[index];
}

/*
* Add a Crash to the collection
* @parm crash: Crash to add
* @return true if added, false if repeted record
*
*/
func addCrash(crash Crash)bool{
	var aux = findCrashIndex(crash.CrashId);
	if(aux != -1){
		log.Printf("Crash "+crash.CrashId+" not added")
		return false;
	}
	crashes = append(crashes,crash)
	//log.Printf("Crash "+crash.CrashId+" added")
	return true;
}

/*
* return crash Collection
* @return crash collection
*/
func allCrashes()[]Crash{
	return crashes
}


/*
* Remove a Crash from the collection
* @parm crash: Crash to remove
* @return true if removed, false if not found record
*
*/
func removeCrash(id string)bool{
	var index = findCrashIndex(id);
	if(index == -1){
		log.Printf("Crash "+id+" not deleted")
		return false;
	}
	log.Printf("Crash "+id+" deleted")
	crashes = append(crashes[:index], crashes[index+1:]...)
	return true;
}


/*
* Update a Crash from the collection
* @parm crash: Crash to update
* @return true if updated, false if not found record
*/
func updateCrash(id string,crash Crash)bool{
	var index = findCrashIndex(id);
	var index2 = findCrashIndex(crash.CrashId);
	if(index == -1||( index2 == -1 && id != crash.CrashId)){
		log.Printf("Crash "+crash.CrashId+" not updated")
		return false;
	}
	log.Printf("Crash "+crash.CrashId+" updated")
	crashes[index] = crash;
	return true;
}

/**
* Read initial dataset of crashes
*/
func LoadCrashData(){
	fmt.Println("-------------------leyendo registros de accidentes--------------")
	filePath := "crash.csv"
	file, err1 := os.Open(filePath)
	checkCrashError("Unable to read input file "+filePath, err1)
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkCrashError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	for _, record := range records {
		res := addCrash(Crash{
			CrashId : record[0],
			Date    : record[1],
			Time    : record[2],
			Location: record[3],
			Sumary  : record[4]})
		if res == false {
			fmt.Println("Record " +  record[0] + " repeted")
		}
	}
	file.Close()
}




