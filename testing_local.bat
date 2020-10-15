@echo off
set "url=localhost:8080"
echo . 
echo.
echo ****Cuando presione enter, se abrira otra ventana en la cual arrancara el servicio en el puerto 8080****
echo . 
echo Regrese a esta ventana una vez levantado el servicio para continuar con las pruebas
pause
start run.bat
echo INICIANDO SET DE PRUEBAS **Airplane_Crashes_and_Fatalities_Since_1908** en %url%
echo . 
echo . 
echo .  

call ./test/1st.bat "%url%"
call ./test/1st.bat "%url%"
call ./test/1st.bat "%url%"