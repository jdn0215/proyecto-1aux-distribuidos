prompt TEST 3: 
cls
echo (3/3) Funciones basicas de registro de estadisticas
echo . 
pause
echo . 
echo . 
echo . 
echo *******(1/3) Agregar un nuevo registro -X POST ../statistic/*******
echo . 
@echo on
curl -X POST "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/statistic/" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"statisticId\":\"testing\",\"crashId\":\"1\",\"name\":\"1\",\"aboard\":\"1\",\"fatalities\":\"1\",\"ground\":\"1\"}"
@echo off
echo . 
echo *******Se verifica el dato agregado mediante un get*******
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/Statistic/testing" -H  "accept: application/json"
@echo off   
echo . 
pause
echo . 
echo . 
echo . 
echo *******(2/3) Modificar dicho registro PUT  ../Statistic/testing*******
echo . 
@echo on
curl -X PUT "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/Statistic/testing" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"statisticId\":\"testing\",\"crashId\":\"555\",\"name\":\"555\",\"aboard\":\"555\",\"fatalities\":\"555\",\"ground\":\"555\"}"
echo *******Se verifica el dato modificado mediante un get*******
@echo off
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/Statistic/testing" -H  "accept: application/json"
@echo off   
echo . 
pause
echo . 
echo . 
echo . 
echo *******(3/3) se elimina el dato de test  DELETE ../Statistic/testing*******
@echo on
curl -X DELETE "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/Statistic/testing" -H  "accept: */*"
@echo off
echo . 
echo *******Se verifica el dato eliminado mediante un get *******
TIMEOUT /T 3
echo . 
@echo on
curl -X GET "localhost:8080/1510/Airplane_Crashes_and_Fatalities_Since_1908/1.0.0/Statistic/testing" -H  "accept: application/json"
@echo off
echo . 
pause