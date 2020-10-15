prompt TEST 2: 
cls
echo (2/3) Funciones basicas de registro de vuelos
echo . 
pause
echo . 
echo . 
echo . 
echo *******(1/3) Agregar un nuevo registro -X POST ../flight/*******
echo . 
@echo on
curl -X POST "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"flightId\":\"testing\",\"crashId\":\"1\",\"operator\":\"Military - U.S. Army\",\"flight\":\"na\",\"route\":\"Demonstration\",\"type\":\"Wright Flyer III\",\"registration\":\"III\"}"
@echo off
echo . 
echo *******Se verifica el dato agregado mediante un get*******
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/testing" -H  "accept: application/json"
@echo off   
echo . 
pause
echo . 
echo . 
echo . 
echo *******(2/3) Modificar dicho registro PUT (se modifica la ruta) ../flight/testing*******
echo . 
@echo on
curl -X PUT "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/testing" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"flightId\":\"testing\",\"crashId\":\"1\",\"operator\":\"Military - U.S. Army\",\"flight\":\"na\",\"route\":\"Demonstration de cambio\",\"type\":\"Wright Flyer III\",\"registration\":\"III\"}"
echo *******Se verifica el dato modificado mediante un get*******
@echo off
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/testing" -H  "accept: application/json"
@echo off   
echo . 
pause
echo . 
echo . 
echo . 
echo *******(3/3) se elimina el dato de test  DELETE ../flight/testing*******
@echo on
curl -X DELETE "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/testing" -H  "accept: */*"
@echo off
echo . 
echo *******Se verifica el dato eliminado mediante un get *******
TIMEOUT /T 3
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/flight/testing" -H  "accept: application/json"
@echo off
echo . 
pause
echo . 
echo . 
echo .