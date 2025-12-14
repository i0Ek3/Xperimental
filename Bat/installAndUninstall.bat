@echo off

:: black background green text
color 0a


echo=
echo ========================================================
echo   Uninstall OA/VPN/WPS first and then reinstall'em all
echo ========================================================
echo=
echo ========================================================
echo           Now, we've gonna uninstall'em All
echo ========================================================
echo=



:: Set path for given variable
set Base=C:\
set OA=DSOA\
set PAT="Program Files (x86)\"
set VPN="Hillstone\Hillstone Secure Connect\"
set WPS="Kingsoft\WPS Office\11*\utility\"
set time=%time::=%


:: function call
call:removeOA
call:removeVPN
call:removeWPS
call:installAll


echo=
echo ========================================================
echo               All done, please enjoy!
echo ========================================================
echo=


:renameDSOA
cd %Base%
echo Renaming DSOA to deleteMe-%time%......
echo=
rename %OA% deleteMe-%time%
goto:eof


:removeOA
if exist %Base%%OA% (
    cd  %Base%%OA%
    echo Uninstalling OA......
    echo=
    start unins000.exe /verysilent /S
    
    call:renameDSOA
) else (
    echo OA: No software will be removed!
    echo=
)
goto:eof


:removeVPN
if exist %Base%%PAT%%VPN% (
    cd %Base%%PAT%%VPN%
    echo Uninstalling VPN......
    echo=
    start Uninstall.exe /verysilent /S
) else (
    echo VPN: No software will be removed!
    echo=
)
goto:eof


:removeWPS
if exist %Base%%PAT%%WPS% (
    cd %Base%%PAT%%WPS%
    echo Uninstalling WPS......
    echo=
    start uninst.exe /verysilent /S
) else (
    echo WPS: No software will be removed!
    echo=
)
goto:eof


:installAll
echo=
echo ========================================================
echo             Now, we've gonna install'em all
echo ========================================================
echo=

cd C:\Users\%username%\Desktop\inst-all

echo Installing VPN......
echo=
start /wait 0.exe /verysilent sp-

echo Installing OA......
echo=
start /wait 1.exe /verysilent sp-

echo Installing WPS......
echo=
start /wait wpspro.exe /verysilent sp-
goto:eof