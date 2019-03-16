#!/usr/bin/env bash

export PSDOOMPSCMD="psdoom-containers docker ps"
export PSDOOMRENICECMD="true"
export PSDOOMKILLCMD="psdoom-containers docker kill"

export DOOMWADPATH=/usr/share/games/doom

psdoom-ng -window -merge -nosound -nosfx -nomusic -nograbmouse -devparm