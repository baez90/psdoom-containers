#!/usr/bin/env bash

export PSDOOMPSCMD="psdoom-containers k8s ps"
export PSDOOMRENICECMD="true"
export PSDOOMKILLCMD="psdoom-containers k8s kill"

export DOOMWADPATH=/usr/share/games/doom

psdoom-ng -window -merge -nosound -nosfx -nomusic -nograbmouse
