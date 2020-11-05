/*
 * Copyright (C) 2020  Dmitry Minko
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * MEGA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package iocomponent

import (
	"github.com/nanu-c/qml-go"
	"log"
	"os"
)

type IOComponent struct {
	configFolder string
	cacheFolder string
	appDataFolder string
}

func (io *IOComponent) WriteToFile(fileName string, text string) {
	err := os.Mkdir(io.configFolder, 0777)
	if err != nil {
		log.Println(err)
	}
	file, err := os.Create(io.configFolder + fileName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		log.Println(err)
	}
}

func Register(packageName string) {
	qml.RegisterTypes("GoIOComponent", 0, 1, []qml.TypeSpec{{
		Init: func(v *IOComponent, obj qml.Object) {
			home, err := os.UserHomeDir()
			if err != nil {
				log.Println("GoIOComponent error")
				log.Println(err)
			}
			v.configFolder =  home + "/.config/" + packageName + "/"
			v.cacheFolder =  home + "/.cache/" + packageName + "/"
			v.appDataFolder =  home + "/.local/share/" + packageName + "/"
			log.Println("Init IOComponent")
		},
	}})
}
