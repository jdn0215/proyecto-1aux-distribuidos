package swagger

import(
	"log"
	"os"
	"fmt"
	"encoding/csv"
)


/**
* Statistices collection
*/
var statistics = []Statistic{}

/**
* Error handler
* @parm message: error message
* @parm err: error 
**/
func checkStatisticError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}


/**
* find a Statistic index by id
* @id: Statistic Id to find
* @return: Statistic's index or -1 if not found
*/
func findStatisticIndex(id string)int{
	for index, statistic := range statistics {
        if id == statistic.StatisticId {
            return index
        }
    }
    return -1
}

/*
* unlink a Statistic from a crash
* @parm StatisticId: Statistic id
* return true if ok, false, if Statistic doesn't exist
*
*/
func unlinkStatistic(StatisticId string) bool{
	var statistic = getStatisticById(StatisticId)
	if statistic == nil{
		return false;
	}
	statistic.CrashId = "-";
	return updateStatistic(statistic.StatisticId,*statistic)
}

/*
* Retursn statistic pointer by id
* @parm id: Statistic Id to find
* @return Statistic pointer or nil if not found
*
*/
func getStatisticById(id string)*Statistic{
	var index = findStatisticIndex(id)
	if(index == -1){
		return nil
	}
	return &statistics[index];
}

/*
* Retursn statistic pointer by id
* @parm id: Crash Id to find
* @return Statistic pointer or nil if not found
*
*/
func getStatisticByCrashId(id string)*Statistic{
	for _, statistic := range statistics {
        if id == statistic.CrashId {
            return &statistic
        }
    }
    return nil
}


/*
* Add a Statistic to the collection
* @parm statistic: Statistic to add
* @return true if added, false if repeted record
*
*/
func addStatistic(statistic Statistic)bool{
	var aux = findStatisticIndex(statistic.StatisticId);
	if(aux != -1){
		log.Printf("Statistic "+statistic.StatisticId+" not added")
		return false;
	}
	statistics = append(statistics,statistic)
	//log.Printf("Statistic "+statistic.StatisticId+" added")
	return true;
}

/*
* return statistic Collection
* @return statistic collection
*/
func allStatistics()[]Statistic{
	return statistics
}


/*
* Remove a Statistic from the collection
* @parm statistic: Statistic to remove
* @return true if removed, false if not found record
*
*/
func removeStatistic(id string)bool{
	var index = findStatisticIndex(id);
	if(index == -1){
		log.Printf("Statistic "+id+" not deleted")
		return false;
	}
	log.Printf("Statistic "+id+" deleted")
	statistics = append(statistics[:index], statistics[index+1:]...)
	return true;
}


/*
* Update a Statistic from the collection
* @parm statistic: Statistic to update
* @return true if updated, false if not found record
*/
func updateStatistic(id string,statistic Statistic)bool{
	var index = findStatisticIndex(id);
	var index2 = findStatisticIndex(statistic.StatisticId);
	if(index == -1||( index2 == -1 && id != statistic.StatisticId)){
		log.Printf("Statistic "+statistic.StatisticId+" not updated")
		return false;
	}
	log.Printf("Statistic "+statistic.StatisticId+" updated")
	statistics[index] = statistic;
	return true;
}

/**
* Read initial dataset of statistics
*/
func LoadStatisticData(){
	fmt.Println("-------------------leyendo registros de estadisticas--------------")
	filePath := "statistics.csv"
	file, err1 := os.Open(filePath)
	checkStatisticError("Unable to read input file "+filePath, err1)
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkStatisticError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	for _, record := range records {
		res := addStatistic(Statistic{
			StatisticId : record[0],
			CrashId     : record[1],
			Aboard      : record[2],
			Fatalities  : record[3],
			Ground      : record[4]})
		if res == false {
			fmt.Println("Record " +  record[0] + " repeted")
		}
	}
	file.Close()
}


