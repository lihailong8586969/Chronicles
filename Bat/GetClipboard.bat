// bat commands

@echo off
::// GetClipboard v1.0 Copyright (C) 2011 By kio (TW) 2011.9.20
::// Only for WinXP,Win7
cd /d "%~dp0"
echo Set WSShell = WScript.CreateObject("WScript.Shell")>temp.txt
echo WSShell.SendKeys "^v">>temp.txt
echo WSShell.SendKeys "^s">>temp.txt
echo WSShell.SendKeys "%%{F4}">>temp.txt
set /p a=<nul>c.txt
start /b cscript temp.txt //E:VBS //NOLOGO
notepad c.txt
for /f "usebackq delims=" %%a in (c.txt) do (echo %%a)
del temp.txt c.txt 2>nul
echo.
pause
