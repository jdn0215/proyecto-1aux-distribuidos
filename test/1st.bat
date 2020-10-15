prompt TEST 1: 
set "var=%~1"
echo (1/3) Funciones basicas de registro de accidentes
echo . 
echo .
echo .
pause
echo *******(1/7) Agregar un nuevo registro -X POST ../crash*******
echo . 
@echo on
curl -X POST "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"crashId\":\"test\",\"date\":\"09/17/1908\",\"time\":\"17:18\",\"location\":\"testing\",\"sumary\":\"During a demonstration flight\"}"
@echo off   
echo . 
echo *******Se verifica el dato agregado mediante un get*******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test" -H  "accept: application/json"
@echo off   
echo . 
echo .
echo .
pause
echo . 
echo *******(2/7) Modificar dicho registro PUT (se modifica la fecha) ../crash/test*******
@echo on
curl -X PUT "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/1" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"crashId\":\"test\",\"date\":\"09/17/2010\",\"time\":\"17:18\",\"location\":\"DATO ACTUALIZADO\",\"sumary\":\"During a demonstration flight\"}"
@echo off
echo . 
echo *******Se verifica el dato modificado mediante un get *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test" -H  "accept: application/json"
@echo off
echo . 
echo .
echo .
pause
echo . 
echo *******(3/7) Se liga un vuelo al accidente -X POST /crash/test/flight*******
@echo on
curl -X POST "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/flight" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"id\":\"56\"}"
@echo off
echo . 
echo *******se comprueba consultando mediante una consulta de vuelos por id de accidente -X GET ../crash/test/flight *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/flight" -H  "accept: application/json"
@echo off
echo . 
echo .
echo .
pause
echo . 
echo *******(4/7) Se liga una estadistica al accidente -X POST /crash/test/statistic*******
@echo on
curl -X POST "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/statistic" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"id\":\"56\"}"
@echo off
echo . 
echo *******se comprueba consultando mediante una consulta de vuelos por id de accidente -X GET ../crash/test/statistic *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/statistic" -H  "accept: application/json"
@echo off
echo . 
pause 
echo . 
echo . 
echo . 
echo *******(5/7) se elimina en elnace entre el accidente y el vuelo DELETE ../crash/test/flight/56*******
@echo on
curl -X DELETE "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/flight/56" -H  "accept: */*"
@echo off
echo . 
echo *******se comprueba consultando mediante una consulta de vuelos por id de accidente -X GET ../crash/test/flight *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/flight" -H  "accept: application/json"
@echo off
echo . 
pause
echo . 
echo . 
echo . 
echo *******(6/7) se elimina en elnace entre el accidente y el vuelo DELETE ../crash/test/statistic/56*******
@echo on
curl -X DELETE "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/statistic/56" -H  "accept: */*"
@echo off
echo . 
echo *******se comprueba consultando mediante una consulta de vuelos por id de accidente -X GET ../crash/test/statistic *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test/statistic" -H  "accept: application/json"
@echo off
echo . 
pause 
echo . 
echo . 
echo . 
echo *******(7/7) se elimina el dato de test  DELETE ../crash/test*******
@echo on
curl -X DELETE "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test" -H  "accept: */*"
@echo off
echo . 
echo *******Se verifica el dato eliminado mediante un get *******
echo . 
@echo on
curl -X GET "%var%/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/crash/test" -H  "accept: application/json"
@echo off
echo . 
pause
echo . 
echo . 
echo .